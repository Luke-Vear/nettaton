package main

import (
	"encoding/json"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// LamdaResponse is a specific JSON response required in order for Lambda Proxy to work with API Gateway.
type LamdaResponse struct {
	StatusCode string            `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

// ClientRequest is what the front end will be sending.
type ClientRequest struct {
	Answer       string `json:"answer"`
	IPAddress    string `json:"ipAddress"`
	Network      string `json:"network"`
	QuestionKind string `json:"questionKind"`
}

var funcMap = map[string]func(string, string) string{
	"first":     First,
	"last":      Last,
	"broadcast": Broadcast,
	"range":     Range,
}

// Handle is the entrypoint for the shim.
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

	headers := map[string]string{"Content-Type": "application/json"}

	var cr *ClientRequest
	if err := json.Unmarshal(evt, cr); err != nil {
		return LamdaResponse{
			StatusCode: "400",
			Headers:    headers,
			Body:       "{\"Error\": \"Problem Unmarshalling JSON from client.\"}",
		}, err
	}

	if err := validateClientRequest(cr); err != nil {
		return LamdaResponse{
			StatusCode: "400",
			Headers:    headers,
			Body:       "{\"Error\": \"Problem with validateClientRequest from client.\"}",
		}, err
	}

	ans := funcMap[cr.QuestionKind](cr.IPAddress, cr.Network)
	if ans == cr.Answer {
		increaseScoreDynamo()
	}

	body, _ := json.Marshal(struct {
		UserAnswer   string `json:"userAnswer"`
		ActualAnswer string `json:"actualAnswer"`
	}{
		UserAnswer:   cr.Answer,
		ActualAnswer: ans,
	})

	return LamdaResponse{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

func validateClientRequest(cr *ClientRequest) error {
	// TODO: intentially send errors
	// return error("This is an invalidated client request")
	// TODO: implement validation
	return nil
}

func increaseScoreDynamo() {}

// First returns.
func First(ip string, net string) string {
	return ""
}

// Last returns.
func Last(ip string, net string) string {
	return ""
}

// Broadcast returns.
func Broadcast(ip string, net string) string {
	return ""
}

// Range returns.
func Range(ip string, net string) string {
	return ""
}

// Handle is the entrypoint for the shim.
func main() {}
