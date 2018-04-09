package state

import (
	"github.com/Luke-Vear/nettaton/internal/quiz"
)

type Gateway struct {
	datastore datastore
}

type datastore interface {
	UpdateQuestion(*quiz.Question) error
	GetQuestion(string) (*quiz.Question, error)
}

func NewGateway(datastore datastore) *Gateway {
	return &Gateway{
		datastore: datastore,
	}
}

func (sgw *Gateway) UpdateQuestion(qq *quiz.Question) error {
	return sgw.datastore.UpdateQuestion(qq)
}

func (sgw *Gateway) GetQuestion(questionID string) (*quiz.Question, error) {
	return sgw.datastore.GetQuestion(questionID)
}
