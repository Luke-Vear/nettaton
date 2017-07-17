package main

import (
	"encoding/json"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
)

// Handle is invoked by the shim.
func Handle(evt *cpf.Event, ctx *cpf.Context) (interface{}, error) {

	// Extract user from path parameters and define PK for query.
	if userID, ok := evt.PathParameters["userID"]; !ok || userID == "" {
		return cpf.NewResponse("400", "", cpf.ErrUserNotSpecified)
	}
	user := cpf.NewUser(evt.PathParameters["userID"])

	// Get User from db.
	if err := cpf.GetUser(user); err != nil {
		return cpf.NewResponse("404", "", err)
	}

	// Return only scores in response.
	body, _ := json.Marshal(struct {
		Scores map[string]*cpf.QuestionScore `json:"scores"`
	}{
		Scores: user.Scores,
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is invoked by the shim.
func main() {}
