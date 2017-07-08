package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Luke-Vear/nettaton/pkg/auth"
	"github.com/Luke-Vear/nettaton/pkg/platform"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
)

// Handle is invoked by the shim.
func Handle(evt json.RawMessage, ctx *runtime.Context) (interface{}, error) {

	headers := map[string]string{"Content-Type": "application/json"}

	jwtString := platform.JWTFromEvt(evt)

	if jwtString == "" {
		return platform.Response{
			StatusCode: "401",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"You must login to retrieve scores.\"}"),
		}, nil
	}

	userID, err := auth.UserID(jwtString, os.Getenv("SECRET"))
	if err != nil {
		return platform.Response{
			StatusCode: "401",
			Headers:    headers,
			Body:       fmt.Sprintf("{\"Error\": \"%v\"}", err),
		}, err
	}

	// get scores from DB
	_ = userID

	// marshall into response
	body, _ := json.Marshal(struct {
		Score1 string `json:"score1"`
		Score2 string `json:"score2"`
	}{
		Score1: "",
		Score2: "",
	})

	return platform.Response{
		StatusCode: "200",
		Headers:    headers,
		Body:       string(body),
	}, nil
}

func main() {}
