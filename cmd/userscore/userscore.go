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

	// Read User from db.
	if err := user.Read(); err != nil {
		return cpf.NewResponse("404", "", err)
	}
	if user.Status == "" {
		return cpf.NewResponse("404", "", cpf.ErrUserNotFoundInDatabase)
	}

	// Return only marks in response.
	body, _ := json.Marshal(struct {
		Marks map[string]*cpf.Marks `json:"marks"`
	}{
		Marks: user.Marks,
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is invoked by the shim.
func main() {}
