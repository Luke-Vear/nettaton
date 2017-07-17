package main

import (
	"encoding/json"

	"github.com/Luke-Vear/nettaton/pkg/do"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

// Request is what the client will be sending.
type Request struct {
	UserID   string `json:"userID"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Handle is the entrypoint for the shim.
func Handle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	var cr Request
	if err := json.Unmarshal([]byte(evt.Body), &cr); err != nil {
		return platform.NewResponse("400", "", err)
	}

	// Check required fields. TODO: max size.
	if cr.UserID == "" || cr.Password == "" || cr.Email == "" {
		return platform.NewResponse("400", "", platform.ErrRequiredFieldNotInRequest)
	}

	// Define PK for query.
	user := do.NewUser()
	user.UserID = cr.UserID

	// Try and get user from database, return if there is an error other than user not found.
	err := platform.GetUser(user)
	if err != platform.ErrUserNotFoundInDatabase && err != nil {
		return platform.NewResponse("500", "", err)
	}

	// If there is not a user not found error, return.
	if err != platform.ErrUserNotFoundInDatabase {
		return platform.NewResponse("409", "", platform.ErrUserAlreadyExists)
	}

	// Create user.
	user.Email = cr.Email
	user.Password = cr.Password
	if err := platform.PutUser(user); err != nil {
		platform.NewResponse("500", "", err)
	}

	// Generate and marshal random IP, network and question into response.
	body, _ := json.Marshal(struct {
		UserID string `json:"userID"`
	}{
		UserID: user.UserID,
	})
	return platform.NewResponse("201", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
