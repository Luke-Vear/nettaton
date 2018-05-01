package platform

import (
	"encoding/json"
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
		// TODO: log the err
		errJSON, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{
			Error: http.StatusText(statusCode),
		})
		body = string(errJSON)
	}
	return &Response{
		Headers:    map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"},
		StatusCode: statusCode,
		Body:       body,
	}, nil
}
