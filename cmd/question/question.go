package main

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/Luke-Vear/nettaton/pkg/subnet"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Handle is the entrypoint for the shim.
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

	// Need to seed for random question generation below.
	rand.Seed(time.Now().UTC().UnixNano())

	// Generate and marshal random IP, network and question into response.
	body, _ := json.Marshal(struct {
		IPAddress    string `json:"ipAddress"`
		Network      string `json:"network"`
		QuestionKind string `json:"questionKind"`
	}{
		IPAddress:    subnet.RandomIP(),
		Network:      subnet.RandomNetwork(),
		QuestionKind: subnet.RandomQuestionKind(),
	})
	return platform.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
