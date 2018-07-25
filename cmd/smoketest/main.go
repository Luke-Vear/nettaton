package main

import (
	"encoding/json"
	"fmt"

	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/google/uuid"

	"github.com/Luke-Vear/nettaton/internal/nettaton"
)

func main() {

	nc := nettaton.NewClient("api.dev.nettaton.com")

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

	rq, err := nc.ReadQuestion(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", rq)

	aq, err := nc.AnswerQuestion(u, qq.Solution())
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", aq)
}
