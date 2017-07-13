package main

import (
	"encoding/json"
	"fmt"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Handle is invoked by the shim.
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

	headers := map[string]string{"Content-Type": "application/json"}

	// Extract jwt from headers.
	jwtString := platform.JWTFromEvt(evt)
	if jwtString == "" {
		return platform.Response{
			StatusCode: "401",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"You must login to retrieve scores.\"}"),
		}, nil
	}

	// Parse UserID claim.
	userID, err := auth.UserID(jwtString)
	if err != nil {
		return platform.Response{
			StatusCode: "401",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	// Define PK for query.
	user := platform.NewUser()
	user.UserID = userID

	// Get User from db.
	if err := platform.GetUser(user); err != nil {
		return platform.Response{
			StatusCode: "500",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	// Return only scores in response.
	body, _ := json.Marshal(struct {
		Scores map[string]*platform.QuestionScore `json:"scores"`
	}{
		Scores: user.Scores,
	})

	return platform.Response{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

// Handle is invoked by the shim.
func main() {}
