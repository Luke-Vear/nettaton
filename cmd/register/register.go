package main

import (
	"encoding/json"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

// Request is what the client will be sending.
type Request struct {
	cpf.User
}

// Handle is the entrypoint for the shim.
func Handle(evt *cpf.Event, ctx *cpf.Context) (interface{}, error) {

	var cr Request
	if err := json.Unmarshal([]byte(evt.Body), &cr); err != nil {
		return cpf.NewResponse("400", "", err)
	}

	// Check required fields.
	if cr.ID == "" || cr.ClearTextPassword == "" {
		return cpf.NewResponse("400", "", cpf.ErrRequiredFieldNotInRequest)
	}

	// Create *User object.
	user := cpf.NewUser(cr.ID)

	// Attempt to serialize new User to database. Return conflict if exists.
	err := user.Create()
	if err != nil && err != cpf.ErrUserAlreadyExists {
		return cpf.NewResponse("500", "", err)
	}
	if err == cpf.ErrUserAlreadyExists {
		return cpf.NewResponse("409", "", err)
	}

	// Return user ID to client.
	body, _ := json.Marshal(struct {
		ID string `json:"id"`
	}{
		ID: cr.ID,
	})
	return cpf.NewResponse("201", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
