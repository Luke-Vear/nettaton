package main

import (
	"encoding/json"
	"fmt"

	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Request is what the client will be sending.
type Request struct {
	// Answer       string `json:"answer"`
	// IPAddress    string `json:"ipAddress"`
	// Network      string `json:"network"`
	// QuestionKind string `json:"questionKind"`
}

// Register User
// Login User

// Handle is the entrypoint for the shim.
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

	headers := map[string]string{"Content-Type": "application/json"}

	var cr *Request
	if err := json.Unmarshal(evt, cr); err != nil {
		return platform.Response{
			StatusCode: "400",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	body, _ := json.Marshal(struct{}{})

	return platform.Response{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

// Handle is the entrypoint for the shim.
func main() {}
