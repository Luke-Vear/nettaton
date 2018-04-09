package nettaton

import (
	"fmt"

	"github.com/Luke-Vear/nettaton/internal/state"

	"github.com/Luke-Vear/nettaton/internal/platform"
	"github.com/Luke-Vear/nettaton/internal/quiz"
)

var stateGateway *state.Gateway

func SetStateGateway(sgw *state.Gateway) {
	stateGateway = sgw
}

type QuestionOperation func(r *platform.Request) (*quiz.Question, error)

// CreateQuestion ...
func CreateQuestion(r *platform.Request) (*quiz.Question, error) {
	var kind string
	if k, ok := r.QueryStringParameters["kind"]; ok {
		if !quiz.ValidQuestionKind(k) {
			return nil, fmt.Errorf("")
		}
		kind = k
	}

	q := quiz.NewQuestion("", "", kind)

	err := stateGateway.UpdateQuestion(q)
	if err != nil {
		return nil, err
	}

	return q, nil
}

// ReadQuestion ...
func ReadQuestion(r *platform.Request) (*quiz.Question, error) {
	return nil, nil
}

// AnswerQuestion ...
func AnswerQuestion(r *platform.Request) (*quiz.Question, error) {
	return nil, nil
}
