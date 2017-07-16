package platform

import (
	"fmt"
)

// Response is a specific JSON response required in order for Lambda Proxy to work with API Gateway.
type Response struct {
	Headers    map[string]string `json:"headers"`
	StatusCode string            `json:"statusCode"`
	Body       string            `json:"body"`
}

// NewResponse returns a properly formatted Response.
func NewResponse(statusCode string, body string, err error) (Response, error) {

	if err != nil {
		body = fmt.Sprintf(`{"Error": "%v"}`, err)
	}

	return Response{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       body,
	}, err
}
