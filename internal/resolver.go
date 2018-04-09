package internal

import (
	"github.com/Luke-Vear/nettaton/internal/nettaton"
	"github.com/Luke-Vear/nettaton/internal/state"
	"github.com/Luke-Vear/nettaton/internal/state/store"
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
	return nettaton.NewNexus(r.ResolveGateway())
}

// ResolveGateway ...
func (r *DependencyResolver) ResolveGateway() *state.Gateway {
	return state.NewGateway(r.ResolveDB())
}

// ResolveDB ...
func (r *DependencyResolver) ResolveDB() *store.DB {
	dynamo := dynamodb.New(session.New())
	return store.NewDB(r.config.DB.Table, dynamo)
}
