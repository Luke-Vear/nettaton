package main

import (
	"encoding/json"
	"errors"

	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

// Handle is invoked by the shim.
func Handle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	// Extract user from headers.
	if _, ok := evt.PathParameters["userID"]; !ok {
		return platform.NewResponse("400", "", errors.New("no userID path parameter"))
	}

	// Define PK for query.
	user := platform.NewUser()
	user.UserID = evt.PathParameters["userID"]

	// Get User from db.
	if err := platform.GetUser(user); err != nil {
		return platform.NewResponse("404", "", err)
	}

	// Return only scores in response.
	body, _ := json.Marshal(struct {
		Scores map[string]*platform.QuestionScore `json:"scores"`
	}{
		Scores: user.Scores,
	})
	return platform.NewResponse("200", string(body), nil)
}

// Handle is invoked by the shim.
func main() {}
