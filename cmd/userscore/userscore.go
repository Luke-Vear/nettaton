package main

import (
	"encoding/json"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

// Handle is invoked by the shim.
func Handle(evt *cpf.Event, ctx *cpf.Context) (interface{}, error) {

	// Extract user from path parameters and create User.
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

	// Return only marks in response.
	body, _ := json.Marshal(struct {
		Marks map[string]*cpf.Marks `json:"marks"`
	}{
		Marks: user.ListMarks(),
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is invoked by the shim.
func main() {}
