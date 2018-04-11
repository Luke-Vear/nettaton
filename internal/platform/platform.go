package platform

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type (
	// Request ...
	Request = events.APIGatewayProxyRequest
	// Response ...
	Response = events.APIGatewayProxyResponse
)

// NewResponse returns a properly formatted Response.
func NewResponse(statusCode int, body string, err error) (*Response, error) {
	if err != nil {
		// TODO: Log error.
		body = http.StatusText(http.StatusInternalServerError)
	}
	return &Response{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       body,
	}, nil
}
