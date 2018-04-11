package platform

import (
	"log"
	"net/http"
	"os"

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
		log.New(os.Stderr, "ERROR: ", log.Llongfile).Println(body)
		body = http.StatusText(statusCode)
	}
	return &Response{
		Headers:    map[string]string{"Content-Type": "application/json"},
		StatusCode: statusCode,
		Body:       body,
	}, nil
}
