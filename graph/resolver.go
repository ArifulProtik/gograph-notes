package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/graph/directives"
	"github.com/ArifulProtik/gograph-notes/graph/generated"
	"github.com/ArifulProtik/gograph-notes/log"
	"github.com/labstack/echo/v4"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	logger      log.Logger
	dbclient    *ent.Client
	config      *config.Config
	echocontext echo.Context
}

func NewSchema(logger log.Logger, dbclient *ent.Client, cfg *config.Config, ctx echo.Context) graphql.ExecutableSchema {

	c := generated.Config{
		Resolvers: &Resolver{
			logger:      logger,
			dbclient:    dbclient,
			config:      cfg,
			echocontext: ctx,
		},
	}
	c.Directives.Auth = directives.Auth

	return generated.NewExecutableSchema(c)
}
