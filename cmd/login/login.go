package main

import (
	"encoding/json"
	"fmt"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Request is what the client will be sending.
type Request struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
}

// Login User

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

	// Define PK for query.
	user := platform.NewUser()
	user.UserID = cr.UserID

	// Get User from db.
	if err := platform.GetUser(user); err != nil {
		return platform.Response{
			StatusCode: "500",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	// Check password from client against hash in database, get a JWT.
	jwt, err := auth.Login(user, cr.Password)
	if err != nil {
		return platform.Response{
			StatusCode: "401",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	// Return token to client.
	body, _ := json.Marshal(struct {
		AccessToken string `json:"accessToken"`
		TokenType   string `json:"tokenType"`
	}{
		AccessToken: jwt,
		TokenType:   "Bearer",
	})

	return platform.Response{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

// Handle is the entrypoint for the shim.
func main() {}
