package main

import (
	"encoding/json"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
	snq "github.com/Luke-Vear/nettaton/pkg/subnetquiz"
)

// Request is what the client will be sending.
type Request struct {
	Answer       string `json:"answer"`
	IPAddress    string `json:"ipAddress"`
	Network      string `json:"network"`
	QuestionKind string `json:"questionKind"`
}

// responseBody is the body of the response we want to returnt to the client.
// If the client is not logged in (they have no JWT), they won't have any
// marks to update.
type responseBody struct {
	UserAnswer   string                `json:"userAnswer"`
	ActualAnswer string                `json:"actualAnswer"`
	Marks        map[string]*cpf.Marks `json:"marks"`
}

// Handle is the entrypoint for the shim.
func Handle(evt *cpf.Event, ctx *cpf.Context) (interface{}, error) {

	var cr Request
	if err := json.Unmarshal([]byte(evt.Body), &cr); err != nil {
		return cpf.NewResponse("400", "", err)
	}

	// Check required fields.
	if cr.Answer == "" || cr.IPAddress == "" || cr.Network == "" || cr.QuestionKind == "" {
		return cpf.NewResponse("400", "", cpf.ErrRequiredFieldNotInRequest)
	}

	// Attempt to parse ip address and QuestionKind.
	nip, cidr, err := snq.Parse(cr.IPAddress, cr.Network)
	if err != nil {
		return cpf.NewResponse("400", "", err)
	}

	// Test if question type is valid, if valid then resolve answer.
	if _, ok := snq.Questions[cr.QuestionKind]; !ok {
		return cpf.NewResponse("400", "", cpf.ErrInvalidQuestionKind)
	}
	actualAnswer := snq.Questions[cr.QuestionKind](nip, cidr)

	// If empty JWT, not logged in, happy return.
	if jwtString, ok := evt.Headers["Authorization"]; !ok || jwtString == "" {
		body, _ := json.Marshal(responseBody{
			UserAnswer:   cr.Answer,
			ActualAnswer: actualAnswer,
		})
		return cpf.NewResponse("200", string(body), nil)
	}

	// Else request has JWT, parse id.
	id, err := cpf.IDFromToken(evt.Headers["Authorization"])
	if err != nil {
		return cpf.NewResponse("401", "", err)
	}

	// Create *User object.
	user := cpf.NewUser(id)

	// Get User from db.
	if err := user.Read(); err != nil {
		return cpf.NewResponse("500", "", err)
	}
	if user.Status == "" {
		return cpf.NewResponse("418", "", cpf.ErrValidJwtButNoUserInDb)
	}

	// Increment marks.
	if actualAnswer == cr.Answer {
		user.Marks[cr.QuestionKind].Correct++
	}
	user.Marks[cr.QuestionKind].Attempts++

	// Put modified User back into db.
	if err := user.Update(); err != nil {
		return cpf.NewResponse("500", "", err)
	}

	// Send actualAnswer and Marks back to client.
	body, _ := json.Marshal(responseBody{
		UserAnswer:   cr.Answer,
		ActualAnswer: actualAnswer,
		Marks:        user.Marks,
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
