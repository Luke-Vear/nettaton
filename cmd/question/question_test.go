package main

import (
	"reflect"
	"testing"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

func TestHandle(t *testing.T) {

	evt, ctx := &cpf.Event{}, &cpf.Context{}
	expectResp := cpf.Response{
		StatusCode: "200",
		Headers:    map[string]string{"Content-Type": "application/json"},
	}

	resp, err := Handle(evt, ctx)
	if err != nil {
		t.Errorf("actual: %v, error should be nil", err)
	}
	r := resp.(*cpf.Response)
	if r.Body == "" || !reflect.DeepEqual(r.Headers, expectResp.Headers) || r.StatusCode != expectResp.StatusCode {
		t.Errorf("malformed response: %v", r)
	}
}
