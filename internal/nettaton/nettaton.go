package nettaton

import (
	"fmt"

	"github.com/Luke-Vear/nettaton/internal/platform"
	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/Luke-Vear/nettaton/internal/state"
)

// Nexus ...
type Nexus struct {
	sgw *state.Gateway
}

// NewNexus ...
func NewNexus(sgw *state.Gateway) *Nexus {
	return &Nexus{
		sgw: sgw,
	}
}

// CreateQuestion ...
func (n *Nexus) CreateQuestion(r *platform.Request) (*quiz.Question, error) {

	var kind string
	if k, ok := r.QueryStringParameters["kind"]; ok {
		if !quiz.ValidQuestionKind(k) {
			return nil, fmt.Errorf("")
		}
		kind = k
	}

	q := quiz.NewQuestion("", "", kind)

	err := n.sgw.UpdateQuestion(q)
	if err != nil {
		return nil, err
	}

	return q, nil
}

// ReadQuestion ...
func (n *Nexus) ReadQuestion(r *platform.Request) (*quiz.Question, error) {
	return nil, nil
}

// AnswerQuestion ...
func (n *Nexus) AnswerQuestion(r *platform.Request) (*quiz.Question, error) {
	return nil, nil
}
