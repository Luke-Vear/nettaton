package state

import (
	"github.com/Luke-Vear/nettaton/internal/quiz"
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
	UpdateQuestion(*quiz.Question) error
	GetQuestion(string) (*quiz.Question, error)
}

// UpdateQuestion ...
func (sgw *Gateway) UpdateQuestion(qq *quiz.Question) error {
	return sgw.datastore.UpdateQuestion(qq)
}

// GetQuestion ...
func (sgw *Gateway) GetQuestion(questionID string) (*quiz.Question, error) {
	return sgw.datastore.GetQuestion(questionID)
}
