package cloudplatform

import "testing"

func TestNewResponse(t *testing.T) {
	type inputs struct {
		statusCode string
		body       string
		err        error
	}
	type returns struct {
		resp Response
		err  error
	}
	tt := []struct {
		description string
		input       inputs
		expected    returns
	}{
		{
			"all fine",

			inputs{
				statusCode: "200",
				body:       "hello",
				err:        nil,
			},

			returns{
				resp: Response{
					Headers:    map[string]string{"Content-Type": "application/json"},
					StatusCode: "200",
					Body:       "hello",
				},
				err: nil,
			},
		},
	}
	_ = tt
}
