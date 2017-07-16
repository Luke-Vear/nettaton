package main

import (
	"encoding/json"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

// Request is what the client will be sending.
type Request struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

// Handle is the entrypoint for the shim.
func Handle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	var cr Request
	if err := json.Unmarshal([]byte(evt.Body), &cr); err != nil {
		return platform.NewResponse("400", "", err)
	}

	// Define PK for query.
	user := platform.NewUser()
	user.UserID = cr.UserID

	// Get User from db.
	if err := platform.GetUser(user); err != nil {
		return platform.NewResponse("500", "", err)
	}

	// Check password from client against hash in database, get a JWT.
	jwt, err := auth.Login(user, cr.Password)
	if err != nil {
		return platform.NewResponse("401", "", err)
	}

	// Return token to client.
	body, _ := json.Marshal(struct {
		AccessToken string `json:"accessToken"`
		TokenType   string `json:"tokenType"`
	}{
		AccessToken: jwt,
		TokenType:   "Bearer",
	})
	return platform.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
