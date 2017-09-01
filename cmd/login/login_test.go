package main

import (
	"reflect"
	"testing"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

var (
	headers        = map[string]string{"Content-Type": "application/json"}
	pathParameters = map[string]string{"id": "dave"}
)

func TestHandle(t *testing.T) {
	tt := []struct {
		description string
		inputEvt    *cpf.Event
		inputCtx    *cpf.Context
		expectResp  *cpf.Response
		expectErr   error
	}{
		{
			description: "no body",
			inputEvt:    &cpf.Event{},
			inputCtx:    &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "400",
				Body:       `{"Error": "unexpected end of JSON input"}`,
			},
			expectErr: nil,
		},
		{
			description: "missing password",
			inputEvt: &cpf.Event{
				Body: `{"ID": "dave"}`,
			},
			inputCtx: &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "400",
				Body:       `{"Error": "required request field empty"}`,
			},
			expectErr: nil,
		},
		{
			description: "no user",
			inputEvt: &cpf.Event{
				Body: `{"id": "dave", "clearTextPassword": "abc123"}`,
			},
			inputCtx: &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "400",
				Body:       `{"Error": "user not specified"}`,
			},
			expectErr: nil,
		},
		{
			description: "no region",
			inputEvt: &cpf.Event{
				Body:           `{"id": "dave", "clearTextPassword": "abc123"}`,
				PathParameters: pathParameters,
			},
			inputCtx: &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "500",
				Body:       `{"Error": "MissingRegion: could not find region configuration"}`,
			},
			expectErr: nil,
		},
	}
	for _, tc := range tt {
		actualResp, actualErr := Handle(tc.inputEvt, tc.inputCtx)
		if !reflect.DeepEqual(actualResp, tc.expectResp) {
			t.Errorf("test: %v, actualResp: %v, expectResp: %v", tc.description, actualResp, tc.expectResp)
		}
		if !reflect.DeepEqual(actualErr, tc.expectErr) {
			t.Errorf("test: %v, actualErr: %v, error should be: %v", tc.description, actualErr, tc.expectErr)
		}
	}

}
