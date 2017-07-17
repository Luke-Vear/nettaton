package main

import (
	"encoding/json"
	"math/rand"
	"time"

	cpf "github.com/Luke-Vear/nettaton/pkg/cloudplatform"
	snq "github.com/Luke-Vear/nettaton/pkg/subnetquiz"
)

// Handle is the entrypoint for the shim.
func Handle(evt *cpf.Event, ctx *cpf.Context) (interface{}, error) {

	// Need to seed for random question generation below.
	rand.Seed(time.Now().UTC().UnixNano())

	// Generate and marshal random IP, network and question into response.
	body, _ := json.Marshal(struct {
		IPAddress    string `json:"ipAddress"`
		Network      string `json:"network"`
		QuestionKind string `json:"questionKind"`
	}{
		IPAddress:    snq.RandomIP(),
		Network:      snq.RandomNetwork(),
		QuestionKind: snq.RandomQuestionKind(),
	})
	return cpf.NewResponse("200", string(body), nil)
}

// Handle is the entrypoint for the shim.
func main() {}
