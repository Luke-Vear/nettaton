package main

import (
	"encoding/json"
	"math/rand"
	"strconv"
	"time"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

var (
	// Too verbose, I use lots.
	s, r = strconv.Itoa, rand.Intn
)

// LamdaResponse is a specific JSON response required in order for Lambda Proxy to work with API Gateway.
type LamdaResponse struct {
	StatusCode string            `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

// Handle is the entrypoint for the shim.
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

	rand.Seed(time.Now().UTC().UnixNano())

	headers := map[string]string{"Content-Type": "application/json"}

	body, _ := json.Marshal(struct {
		IPAddress    string `json:"ipAddress"`
		Network      string `json:"network"`
		QuestionKind string `json:"questionKind"`
	}{
		IPAddress:    getRandomIP(),
		Network:      getRandomNetwork(),
		QuestionKind: getRandomQuestionKind(),
	})

	return LamdaResponse{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

// TODO: more real internal IP ranges.
func getRandomIP() string {
	return "10" + "." + s(r(253)+1) + "." + s(r(253)+1) + "." + s(r(253)+1)
}

func getRandomNetwork() string {
	switch r(2) {
	case 0:
		return networks[r(len(networks))].netmask
	}
	return networks[r(len(networks))].prefix
}

// networks is used for above func.
var networks = []struct {
	netmask string
	prefix  string
}{
	{netmask: "255.255.255.252", prefix: "/30"},
	{netmask: "255.255.255.248", prefix: "/29"},
	{netmask: "255.255.255.240", prefix: "/28"},
	{netmask: "255.255.255.224", prefix: "/27"},
	{netmask: "255.255.255.192", prefix: "/26"},
	{netmask: "255.255.255.128", prefix: "/25"},
	{netmask: "255.255.255.0", prefix: "/24"},
	{netmask: "255.255.254.0", prefix: "/23"},
	{netmask: "255.255.252.0", prefix: "/22"},
	{netmask: "255.255.248.0", prefix: "/21"},
	{netmask: "255.255.240.0", prefix: "/20"},
	{netmask: "255.255.224.0", prefix: "/19"},
	{netmask: "255.255.192.0", prefix: "/18"},
	{netmask: "255.255.128.0", prefix: "/17"},
	{netmask: "255.255.0.0", prefix: "/16"},
	{netmask: "255.254.0.0", prefix: "/15"},
	{netmask: "255.252.0.0", prefix: "/14"},
	{netmask: "255.248.0.0", prefix: "/13"},
	{netmask: "255.240.0.0", prefix: "/12"},
}

// TODO: parse incoming json for question types the front end wants.
func getRandomQuestionKind() string {
	return qkinds[r(len(qkinds))]
}

// qkinds used for above func.
var qkinds = []string{
	"first",
	"last",
	"broadcast",
	"range",
}

// Handle is the entrypoint for the shim.
func main() {}
