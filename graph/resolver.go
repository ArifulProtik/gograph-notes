package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/graph/generated"
	"github.com/ArifulProtik/gograph-notes/log"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger   log.Logger
	dbclient *ent.Client
}

func NewSchema(logger log.Logger, dbclient *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{
			logger:   logger,
			dbclient: dbclient,
		},
	})
}
