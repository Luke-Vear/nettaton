package platform

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

type (
	// Request is what we expect from the AWS API Gateway Proxy Integration.
	Request = events.APIGatewayProxyRequest
	// Response is what the AWS API Gateway Proxy Integration expects from us.
	Response = events.APIGatewayProxyResponse
)

// headers required so far are:
// "Content-Type" - this is required for AWS API Gateway Proxy integration.
// "Access-Control-Allow-Origin" - other people can consume API.
var headers = map[string]string{"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"}

// NewResponse returns a properly formatted Response.
func NewResponse(statusCode int, body string, err error) (*Response, error) {

	if statusCode >= 500 && statusCode < 600 {
		log.Println(err)
	}

	if err != nil {
		errJSON, _ := json.Marshal(struct {
			Error string `json:"error"`
		}{
			Error: http.StatusText(statusCode),
		})
		body = string(errJSON)
	}

	return &Response{
		Headers:    headers,
		StatusCode: statusCode,
		Body:       body,
	}, nil
}
