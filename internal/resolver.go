package internal

import (
	"github.com/Luke-Vear/nettaton/internal/data"
	"github.com/Luke-Vear/nettaton/internal/nettaton"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DependencyResolver ...
type DependencyResolver struct {
	config Config
}

// NewDependencyResolver ...
func NewDependencyResolver(config Config) *DependencyResolver {
	return &DependencyResolver{
		config: config,
	}
}

// ResolveNettatonNexus ...
func (r *DependencyResolver) ResolveNettatonNexus() *nettaton.Nexus {
	return nettaton.NewNexus(r.ResolveDatastore())
}

// ResolveDatastore ...
func (r *DependencyResolver) ResolveDatastore() *data.Store {
	dynamo := dynamodb.New(session.New())
	return data.NewStore(r.config.DB.Table, dynamo)
}
