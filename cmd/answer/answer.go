package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/Luke-Vear/nettaton/pkg/subnet"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Request is what the client will be sending.
type Request struct {
	Answer       string `json:"answer"`
	IPAddress    string `json:"ipAddress"`
	Network      string `json:"network"`
	QuestionKind string `json:"questionKind"`
}

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

	// Attempt to parse ip address and subnet.
	nip, cidr, err := subnet.Parse(cr.IPAddress, cr.Network)
	if err != nil {
		return platform.Response{
			StatusCode: "400",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	// Test if question type is valid, then resolve answer.
	if _, ok := subnet.QuestionFuncMap[cr.QuestionKind]; !ok {
		return platform.Response{
			StatusCode: "400",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"Invalid questionKind %v\"}", cr.QuestionKind),
		}, err
	}
	actualAnswer := subnet.QuestionFuncMap[cr.QuestionKind](nip, cidr)

	// Extract jwt from headers (if exists), parse user claim, update user scores.
	if jwtString := platform.JWTFromEvt(evt); jwtString != "" {

		userID, err := auth.UserID(jwtString, os.Getenv("SECRET"))
		if err != nil {
			return platform.Response{
				StatusCode: "401",
				Headers:    headers,
				Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
			}, err
		}

		if err := platform.UpdateUserScore(cr.QuestionKind, userID, actualAnswer == cr.Answer); err != nil {
			return platform.Response{
				StatusCode: "500",
				Headers:    headers,
				Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
			}, err
		}
	}

	// Send actualAnswer back to client.
	body, _ := json.Marshal(struct {
		UserAnswer   string `json:"userAnswer"`
		ActualAnswer string `json:"actualAnswer"`
	}{
		UserAnswer:   cr.Answer,
		ActualAnswer: actualAnswer,
	})

	return platform.Response{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

// Handle is the entrypoint for the shim.
func main() {}
