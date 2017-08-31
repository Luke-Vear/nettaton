package cloudplatform

import (
	"reflect"
	"testing"
)

func TestNewResponse(t *testing.T) {
	type input struct {
		statusCode string
		body       string
		err        error
	}
	type expected struct {
		resp Response
		err  error
	}
	tt := []struct {
		description string
		input       input
		expected    expected
	}{
		{
			"all fine",

			input{
				statusCode: "200",
				body:       "hello",
				err:        nil,
			},

			expected{
				resp: Response{
					Headers:    map[string]string{"Content-Type": "application/json"},
					StatusCode: "200",
					Body:       "hello",
				},
				err: nil,
			},
		},
		{
			"an error!",

			input{
				statusCode: "503",
				body:       "you should not see this",
				err:        ErrUserAlreadyExists,
			},

			expected{
				resp: Response{
					Headers:    map[string]string{"Content-Type": "application/json"},
					StatusCode: "503",
					Body:       `{"Error": "user already exists in database"}`,
				},
				err: nil,
			},
		},
	}
	for _, tc := range tt {
		actualResponse, actualErr := NewResponse(tc.input.statusCode, tc.input.body, tc.input.err)
		if !reflect.DeepEqual(actualResponse, tc.expected.resp) {
			t.Errorf("test: %v\nactualResponse: %v\nexpectedResponse: %v\n", tc.description, actualResponse, tc.expected.resp)
		}
		if !reflect.DeepEqual(actualErr, tc.expected.err) {
			t.Errorf("test: %v\actualErr: %v\nexpectedErr: %v\n", tc.description, actualResponse, tc.expected.resp)
		}
	}
}
