package nettaton

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Luke-Vear/nettaton/internal/data"
	pf "github.com/Luke-Vear/nettaton/internal/platform"
	"github.com/Luke-Vear/nettaton/internal/quiz"
)

var (
	ErrLenIdZero           = fmt.Errorf("length of id path parameter is zero")
	ErrInvalidQuestionKind = fmt.Errorf("question kind is invalid")
)

// Nexus ...
type Nexus struct {
	ds datastore
}

// datastore ...
type datastore interface {
	GetQuestion(string) (*quiz.Question, error)
	UpdateQuestion(*quiz.Question) error
	DeleteQuestion(questionID string) error
}

// NewNexus ...
func NewNexus(ds datastore) *Nexus {
	return &Nexus{
		ds: ds,
	}
}

// CreateQuestion ...
func (n *Nexus) CreateQuestion(r *pf.Request) (*pf.Response, error) {
	var kind string
	if k, ok := r.QueryStringParameters["kind"]; ok {
		if !quiz.ValidQuestionKind(k) {
			err := fmt.Errorf("%s: %s", ErrInvalidQuestionKind.Error(), k)
			return pf.NewResponse(http.StatusBadRequest, "", err)
		}
		kind = k
	}
	q := quiz.NewQuestion("", "", kind)

	err := n.ds.UpdateQuestion(q)
	if err != nil {
		return pf.NewResponse(http.StatusInternalServerError, "", err)
	}

	body, _ := json.Marshal(q)
	return pf.NewResponse(http.StatusCreated, string(body), nil)
}

// ReadQuestion ...
func (n *Nexus) ReadQuestion(r *pf.Request) (*pf.Response, error) {
	id, _ := r.PathParameters["id"]
	if len(id) == 0 {
		return pf.NewResponse(http.StatusBadRequest, "", ErrLenIdZero)
	}

	q, err := n.ds.GetQuestion(id)
	switch {
	case err == data.ErrQuestionNotFound:
		errNotFound := fmt.Errorf("%s: %s", err.Error(), id)
		return pf.NewResponse(http.StatusNotFound, "", errNotFound)
	case err != nil:
		return pf.NewResponse(http.StatusInternalServerError, "", err)
	}

	body, _ := json.Marshal(q)
	return pf.NewResponse(http.StatusOK, string(body), nil)
}

// AnswerQuestion ...
func (n *Nexus) AnswerQuestion(r *pf.Request) (*pf.Response, error) {
	id, _ := r.PathParameters["id"]
	if len(id) == 0 {
		return pf.NewResponse(http.StatusBadRequest, "", ErrLenIdZero)
	}

	proffered := &struct {
		Answer string `json:"answer"`
	}{}
	if err := json.Unmarshal([]byte(r.Body), &proffered); err != nil {
		return pf.NewResponse(http.StatusBadRequest, "", err)
	}

	q, err := n.ds.GetQuestion(id)
	switch {
	case err == data.ErrQuestionNotFound:
		errNotFound := fmt.Errorf("%s: %s", err.Error(), id)
		return pf.NewResponse(http.StatusNotFound, "", errNotFound)
	case err != nil:
		return pf.NewResponse(http.StatusInternalServerError, "", err)
	}

	var correct bool
	if proffered.Answer == q.Solution() {
		defer n.ds.DeleteQuestion(id)
		correct = true
	}

	body, _ := json.Marshal(struct {
		Correct bool `json:"correct"`
	}{
		Correct: correct,
	})
	return pf.NewResponse(http.StatusOK, string(body), nil)
}
