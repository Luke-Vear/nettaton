package main

import (
	"reflect"
	"testing"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

var (
	headers        = map[string]string{"Content-Type": "application/json"}
	badToken       = map[string]string{"Authorization": "Bearer abcd"}
	oldToken       = map[string]string{"Authorization": "Bearer: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDEyNTg2NjgsIm5iZiI6MTUwMDcxODY2OCwic3ViIjoibmljayJ9.lNSCEYI9Ij7XLlOB4yOq8Ezd1pQeMojmuqeOa4f3LwY"}
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
			description: "missing fields",
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
			description: "good request, no login",
			inputEvt: &cpf.Event{
				Body: `{"answer": "10.10.10.1", "ipAddress": "10.10.10.10", "network": "255.255.255.0", "questionKind": "first" }`,
			},
			inputCtx: &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "200",
				Body:       `{"userAnswer":"10.10.10.1","actualAnswer":"10.10.10.1","marks":null}`,
			},
			expectErr: nil,
		},
		{
			description: "bad token",
			inputEvt: &cpf.Event{
				Body:           `{"answer": "10.10.10.1", "ipAddress": "10.10.10.10", "network": "255.255.255.0", "questionKind": "first" }`,
				PathParameters: pathParameters,
				Headers:        badToken,
			},
			inputCtx: &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "401",
				Body:       `{"Error": "token contains an invalid number of segments"}`,
			},
			expectErr: nil,
		},
		{
			description: "old token",
			inputEvt: &cpf.Event{
				Body:           `{"answer": "10.10.10.1", "ipAddress": "10.10.10.10", "network": "255.255.255.0", "questionKind": "first" }`,
				PathParameters: pathParameters,
				Headers:        oldToken,
			},
			inputCtx: &cpf.Context{},
			expectResp: &cpf.Response{
				Headers:    headers,
				StatusCode: "401",
				Body:       `{"Error": "signature is invalid"}`,
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
