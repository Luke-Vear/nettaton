package nettaton

import (
	"encoding/json"
	"fmt"
	"net/http"

	pf "github.com/Luke-Vear/nettaton/internal/platform"
	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/Luke-Vear/nettaton/internal/state"
)

var (
	ErrLenIdZero = fmt.Errorf("length of id path parameter is zero")
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
func (n *Nexus) CreateQuestion(r *pf.Request) (*pf.Response, error) {
	var kind string
	if k, ok := r.QueryStringParameters["kind"]; ok {
		if !quiz.ValidQuestionKind(k) {
			return pf.NewResponse(http.StatusBadRequest, "kind invalid", nil)
		}
		kind = k
	}
	q := quiz.NewQuestion("", "", kind)

	err := n.sgw.UpdateQuestion(q)
	if err != nil {
		return pf.NewResponse(http.StatusInternalServerError, err.Error(), err)
	}

	body, _ := json.Marshal(q)
	return pf.NewResponse(http.StatusCreated, string(body), nil)
}

// ReadQuestion ...
func (n *Nexus) ReadQuestion(r *pf.Request) (*pf.Response, error) {
	id, _ := r.PathParameters["id"]
	if len(id) == 0 {
		return pf.NewResponse(http.StatusBadRequest, ErrLenIdZero.Error(), ErrLenIdZero)
	}

	q, err := n.sgw.GetQuestion(id)
	switch {
	case err == state.ErrQuestionNotFound:
		return pf.NewResponse(http.StatusNotFound, id+" not found", err)
	case err != nil:
		return pf.NewResponse(http.StatusInternalServerError, err.Error(), err)
	}

	body, _ := json.Marshal(q)
	return pf.NewResponse(http.StatusOK, string(body), nil)
}

// AnswerQuestion ...
func (n *Nexus) AnswerQuestion(r *pf.Request) (*pf.Response, error) {
	id, _ := r.PathParameters["id"]
	if len(id) == 0 {
		return pf.NewResponse(http.StatusBadRequest, ErrLenIdZero.Error(), ErrLenIdZero)
	}

	proffered := &struct {
		Answer string `json:"answer"`
	}{}
	if err := json.Unmarshal([]byte(r.Body), &proffered); err != nil {
		return pf.NewResponse(http.StatusBadRequest, "malformed request body", err)
	}

	q, err := n.sgw.GetQuestion(id)
	switch {
	case err == state.ErrQuestionNotFound:
		return pf.NewResponse(http.StatusNotFound, id+" not found", err)
	case err != nil:
		return pf.NewResponse(http.StatusInternalServerError, err.Error(), err)
	}

	var correct bool
	if proffered.Answer == q.Solution() {
		defer n.sgw.DeleteQuestion(id)
		correct = true
	}

	body, _ := json.Marshal(struct {
		Correct bool `json:"correct"`
	}{
		Correct: correct,
	})
	return pf.NewResponse(http.StatusOK, string(body), nil)
}
