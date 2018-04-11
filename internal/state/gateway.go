package state

import (
	"github.com/Luke-Vear/nettaton/internal/quiz"
	"github.com/Luke-Vear/nettaton/internal/state/store"
)

var (
	ErrQuestionNotFound = store.ErrQuestionNotFound
)

// Gateway ...
type Gateway struct {
	datastore datastore
}

// NewGateway ...
func NewGateway(datastore datastore) *Gateway {
	return &Gateway{
		datastore: datastore,
	}
}

// datastore ...
type datastore interface {
	GetQuestion(string) (*quiz.Question, error)
	UpdateQuestion(*quiz.Question) error
	DeleteQuestion(questionID string) error
}

// GetQuestion ...
func (sgw *Gateway) GetQuestion(questionID string) (*quiz.Question, error) {
	return sgw.datastore.GetQuestion(questionID)
}

// UpdateQuestion ...
func (sgw *Gateway) UpdateQuestion(qq *quiz.Question) error {
	return sgw.datastore.UpdateQuestion(qq)
}

// DeleteQuestion ...
func (sgw *Gateway) DeleteQuestion(questionID string) error {
	return sgw.datastore.DeleteQuestion(questionID)
}
