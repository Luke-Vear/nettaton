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

	// Attempt to parse ip address and snq.
	nip, cidr, err := snq.Parse(cr.IPAddress, cr.Network)
	if err != nil {
		return cpf.NewResponse("400", "", err)
	}

	// Test if question type is valid, then resolve answer.
	if _, ok := snq.QuestionFuncMap[cr.QuestionKind]; !ok {
		return cpf.NewResponse("400", "", cpf.ErrInvalidQuestionKind)
	}
	actualAnswer := snq.QuestionFuncMap[cr.QuestionKind](nip, cidr)

	// Extract jwt from headers (if exists), parse user claim.
	user := cpf.NewUser("")
	if jwtString, ok := evt.Headers["Authorization"]; ok && jwtString != "" {

		userID, err := cpf.UserID(jwtString)
		if err != nil {
			return cpf.NewResponse("401", "", err)
		}

		// Define PK for query.
		user.UserID = userID

		// Get User from db.
		if err := cpf.GetUser(user); err != nil {
			return cpf.NewResponse("500", "", err)
		}

		// Increment scores.
		if actualAnswer == cr.Answer {
			user.Scores[cr.QuestionKind].Correct++
		}
		user.Scores[cr.QuestionKind].Attempts++

		// Put modified User back into db.
		if err := cpf.PutUser(user); err != nil {
			return cpf.NewResponse("500", "", err)
		}
	}

	// Send actualAnswer back to client.
	body, _ := json.Marshal(struct {
		UserAnswer   string                        `json:"userAnswer"`
		ActualAnswer string                        `json:"actualAnswer"`
		Scores       map[string]*cpf.QuestionScore `json:"scores"`
	}{
		UserAnswer:   cr.Answer,
		ActualAnswer: actualAnswer,
		Scores:       user.Scores,
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
