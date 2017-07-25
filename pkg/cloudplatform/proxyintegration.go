package cloudplatform

import (
	"fmt"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

type (
	// Event represents an Amazon API Gateway Proxy Event.
	Event apigatewayproxyevt.Event

	// Context provides information about Lambda execution environment.
	Context runtime.Context
)

// Response is a specific JSON response required for compatibility API Gateway Lambda Proxy.
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
	}, nil
}
