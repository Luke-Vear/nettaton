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
	if cr.ClearTextPassword == "" {
		return cpf.NewResponse("400", "", cpf.ErrRequiredFieldNotInRequest)
	}

	// Extract user from path parameters and create *User object.
	if id, ok := evt.PathParameters["id"]; !ok || id == "" {
		return cpf.NewResponse("400", "", cpf.ErrUserNotSpecified)
	}
	user := cpf.NewUser(evt.PathParameters["id"])

	// Deserialize user from db into User.
	err := user.Read()
	if err != nil && err != cpf.ErrUserNotFoundInDatabase {
		return cpf.NewResponse("500", "", err)
	}
	if err == cpf.ErrUserNotFoundInDatabase {
		return cpf.NewResponse("404", "", err)
	}

	// Check password from client against hash in database, get a JWT.
	jwt, err := user.Login(cr.ClearTextPassword)
	if err != nil {
		return cpf.NewResponse("401", "", err)
	}

	// Return token to client.
	body, _ := json.Marshal(struct {
		AccessToken string `json:"accessToken"`
		TokenType   string `json:"tokenType"`
	}{
		AccessToken: jwt,
		TokenType:   "Bearer",
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
