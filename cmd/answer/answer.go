package main

import (
	"encoding/json"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/do"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/Luke-Vear/nettaton/pkg/subnet"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/eawsy/aws-lambda-go-event/service/lambda/runtime/event/apigatewayproxyevt"
)

// Request is what the client will be sending.
type Request struct {
	Answer       string `json:"answer"`
	IPAddress    string `json:"ipAddress"`
	Network      string `json:"network"`
	QuestionKind string `json:"questionKind"`
}

// Handle is the entrypoint for the shim.
func Handle(evt *apigatewayproxyevt.Event, ctx *runtime.Context) (interface{}, error) {

	var cr Request
	if err := json.Unmarshal([]byte(evt.Body), &cr); err != nil {
		return platform.NewResponse("400", "", err)
	}

	// Check required fields.
	if cr.Answer == "" || cr.IPAddress == "" || cr.Network == "" || cr.QuestionKind == "" {
		return platform.NewResponse("400", "", platform.ErrRequiredFieldNotInRequest)
	}

	// Attempt to parse ip address and subnet.
	nip, cidr, err := subnet.Parse(cr.IPAddress, cr.Network)
	if err != nil {
		return platform.NewResponse("400", "", err)
	}

	// Test if question type is valid, then resolve answer.
	if _, ok := subnet.QuestionFuncMap[cr.QuestionKind]; !ok {
		return platform.NewResponse("400", "", platform.ErrInvalidQuestionKind)
	}
	actualAnswer := subnet.QuestionFuncMap[cr.QuestionKind](nip, cidr)

	// Extract jwt from headers (if exists), parse user claim.
	user := do.NewUser()
	if jwtString, ok := evt.Headers["Authorization"]; ok && jwtString != "" {

		userID, err := auth.UserID(jwtString)
		if err != nil {
			return platform.NewResponse("401", "", err)
		}

		// Define PK for query.
		user.UserID = userID

		// Get User from db.
		if err := platform.GetUser(user); err != nil {
			return platform.NewResponse("500", "", err)
		}

		// Increment scores.
		if actualAnswer == cr.Answer {
			user.Scores[cr.QuestionKind].Correct++
		}
		user.Scores[cr.QuestionKind].Attempts++

		// Put modified User back into db.
		if err := platform.PutUser(user); err != nil {
			return platform.NewResponse("500", "", err)
		}
	}

	// Send actualAnswer back to client.
	body, _ := json.Marshal(struct {
		UserAnswer   string                       `json:"userAnswer"`
		ActualAnswer string                       `json:"actualAnswer"`
		Scores       map[string]*do.QuestionScore `json:"scores"`
	}{
		UserAnswer:   cr.Answer,
		ActualAnswer: actualAnswer,
		Scores:       user.Scores,
	})
	return platform.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
