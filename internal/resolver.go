package internal

import (
	"github.com/Luke-Vear/nettaton/internal/nettaton"
	"github.com/Luke-Vear/nettaton/internal/state"
	"github.com/Luke-Vear/nettaton/internal/state/store"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type DependencyResolver struct {
	config Config
}

func NewDependencyResolver(config Config) *DependencyResolver {
	return &DependencyResolver{
		config: config,
	}
}

func (r *DependencyResolver) ResolveGateway() *state.Gateway {
	return state.NewGateway(r.ResolveDB())
}

func (r *DependencyResolver) ResolveDB() *store.DB {
	dynamo := dynamodb.New(session.New())
	return store.NewDB(r.config.DBConfig.Table, dynamo)
}

func (r *DependencyResolver) ResolveQuestionOperation() nettaton.QuestionOperation {

	nettaton.SetStateGateway(r.ResolveGateway())

	switch r.config.QO.Operation {
	case "read":
		return nettaton.ReadQuestion
	case "answer":
		return nettaton.AnswerQuestion
	default:
		return nettaton.CreateQuestion
	}
}
