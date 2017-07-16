package main

import (
	"encoding/json"

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

	// Check required fields.
	if cr.UserID == "" || cr.Password == "" || cr.Email == "" {
		return platform.NewResponse("400", "", platform.ErrRequiredFieldNotInRequest)
	}

	return platform.NewResponse("200", "", nil)
}

// Handle is the entrypoint for the shim.
func main() {}
