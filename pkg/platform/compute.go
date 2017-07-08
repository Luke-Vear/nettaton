package platform

import (
	"encoding/json"
	"strings"
)

// Response is a specific JSON response required in order for Lambda Proxy to work with API Gateway.
type Response struct {
	StatusCode string            `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

// JWTFromEvt extracts the JWT from an API Gateway integrated Lambda invocation's event.
func JWTFromEvt(evt json.RawMessage) string {

	var event struct {
		Headers struct {
			Authorization string
		}
	}

	cleanedEvt := strings.Replace(string(evt), `\`, "", -1)

	if err := json.Unmarshal([]byte(cleanedEvt), &event); err != nil {
		return ""
	}

	return strings.Split(event.Headers.Authorization, " ")[1]
}
