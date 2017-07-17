package main

import (
	"encoding/json"

	"github.com/Luke-Vear/nettaton/pkg/do"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

// Handle is invoked by the shim.
func Handle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	// Extract user from path parameters and define PK for query.
	if userID, ok := evt.PathParameters["userID"]; !ok || userID == "" {
		return platform.NewResponse("400", "", platform.ErrUserNotSpecified)
	}
	user := do.NewUser()
	user.UserID = evt.PathParameters["userID"]

	// Get User from db.
	if err := platform.GetUser(user); err != nil {
		return platform.NewResponse("404", "", err)
	}

	// Return only scores in response.
	body, _ := json.Marshal(struct {
		Scores map[string]*do.QuestionScore `json:"scores"`
	}{
		Scores: user.Scores,
	})
	return platform.NewResponse("200", string(body), nil)
}

// Handle is invoked by the shim.
func main() {}
