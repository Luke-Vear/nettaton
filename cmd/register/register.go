package main

import (
	"encoding/json"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

// Request is what the client will be sending.
type Request struct {
	cpf.Credentials
}

// Handle is the entrypoint for the shim.
func Handle(evt *cpf.Event, ctx *cpf.Context) (interface{}, error) {

	var cr Request
	if err := json.Unmarshal([]byte(evt.Body), &cr); err != nil {
		return cpf.NewResponse("400", "", err)
	}

	// Check required fields. TODO: max size.
	if cr.UserID == "" || cr.ClearTextPassword == "" {
		return cpf.NewResponse("400", "", cpf.ErrRequiredFieldNotInRequest)
	}

	// Define PK for query.
	user := cpf.NewUser(cr.UserID)

	// Try and get user from database, return if there is an error other than user not found.
	err := cpf.GetUser(user)
	if err != cpf.ErrUserNotFoundInDatabase && err != nil {
		return cpf.NewResponse("500", "", err)
	}
	// If there is not a ErrUserNotFoundInDatabase error, return.
	if err != cpf.ErrUserNotFoundInDatabase {
		return cpf.NewResponse("409", "", cpf.ErrUserAlreadyExists)
	}

	// Create user.
	user.Password = cr.ClearTextPassword
	if err := cpf.PutUser(user); err != nil {
		cpf.NewResponse("500", "", err)
	}

	// Generate and marshal random IP, network and question into response.
	body, _ := json.Marshal(struct {
		UserID string `json:"userID"`
	}{
		UserID: user.UserID,
	})
	return cpf.NewResponse("201", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
