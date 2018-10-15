package main

import (
	"encoding/json"
	"flag"
	"fmt"

	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/google/uuid"

	"github.com/Luke-Vear/nettaton/internal/nettaton"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "", "The environment to smoketest.")
	flag.Parse()

	if env == "prod" {
		env = ""
	} else {
		env = "." + env
	}

	endpoint := "api" + env + ".nettaton.com"
	fmt.Println(endpoint)

	nc := nettaton.NewClient(endpoint)

	q, err := nc.CreateQuestion()
	if err != nil {
		panic(err)
	}

	var qq *quiz.Question
	err = json.Unmarshal([]byte(q), &qq)
	if err != nil {
		panic(err)
	}

	u, err := uuid.Parse(qq.ID)
	if err != nil {
		panic(err)
	}

	_, err = nc.ReadQuestion(u)
	if err != nil {
		panic(err)
	}

	aq, err := nc.AnswerQuestion(u, qq.Solution())
	if err != nil {
		panic(err)
	}

	fmt.Println("expect: { \"correct\": true }")
	fmt.Println("actual:", aq)
}
