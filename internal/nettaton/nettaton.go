package nettaton

import (
	"encoding/json"
	"fmt"
	"net/http"

	pf "github.com/Luke-Vear/nettaton/internal/platform"
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
func (n *Nexus) CreateQuestion(r *pf.Request) (*pf.Response, error) {
	var kind string
	if k, ok := r.QueryStringParameters["kind"]; ok {

		if !quiz.ValidQuestionKind(k) {
			err := fmt.Errorf("")
			return pf.NewResponse(http.StatusBadRequest, "", err)
		}

		kind = k
	}
	q := quiz.NewQuestion("", "", kind)

	err := n.sgw.UpdateQuestion(q)
	if err != nil {
		return pf.NewResponse(http.StatusInternalServerError, "", err)
	}

	body, _ := json.Marshal(q)
	return pf.NewResponse(http.StatusCreated, string(body), nil)
}

// ReadQuestion ...
func (n *Nexus) ReadQuestion(r *pf.Request) (*pf.Response, error) {
	// TODO: Not needed? Can path parameters be len 0?
	id, _ := r.PathParameters["id"]
	if len(id) == 0 {
		return pf.NewResponse(http.StatusBadRequest, "", nil)
	}

	q, err := n.sgw.GetQuestion(id)
	switch {
	case err == state.ErrQuestionNotFound:
		return pf.NewResponse(http.StatusNotFound, "", err)
	case err != nil:
		return pf.NewResponse(http.StatusInternalServerError, "", err)
	}

	body, _ := json.Marshal(q)
	return pf.NewResponse(http.StatusOK, string(body), nil)
}

// AnswerQuestion ...
func (n *Nexus) AnswerQuestion(r *pf.Request) (*pf.Response, error) {
	// TODO: Not needed? Can path parameters be len 0?
	id, _ := r.PathParameters["id"]
	if len(id) == 0 {
		return pf.NewResponse(http.StatusBadRequest, "", nil)
	}

	proffered := &struct {
		Answer string `json:"answer"`
	}{}
	if err := json.Unmarshal([]byte(r.Body), &proffered); err != nil {
		return pf.NewResponse(http.StatusBadRequest, "", err)
	}

	q, err := n.sgw.GetQuestion(id)
	if err == state.ErrQuestionNotFound {
		return pf.NewResponse(http.StatusNotFound, "", err)
	}
	if err != nil {
		return pf.NewResponse(http.StatusInternalServerError, "", err)
	}

	var correct bool
	if proffered.Answer == q.Solution() {
		defer n.sgw.DeleteQuestion(id)
		correct = true
	}

	body, _ := json.Marshal(struct {
		correct bool
	}{
		correct: correct,
	})
	return pf.NewResponse(http.StatusOK, string(body), nil)
}
