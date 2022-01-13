package controllers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/graph"
	"github.com/ArifulProtik/gograph-notes/log"

	"github.com/labstack/echo/v4"
)

type GraphController interface {
	Handlequery(echo.Context) error
	Handleplayground(echo.Context) error
}

type graphController struct {
	logger   log.Logger
	dbclient *ent.Client
	config   *config.Config
}

func NewGraphController(logger log.Logger, dbclient *ent.Client, cfg *config.Config) GraphController {
	return &graphController{
		logger:   logger,
		dbclient: dbclient,
		config:   cfg,
	}
}

func (s *graphController) Handlequery(c echo.Context) error {
	srv := handler.NewDefaultServer(graph.NewSchema(s.logger, s.dbclient, s.config))
	srv.ServeHTTP(c.Response(), c.Request())
	return nil
}
func (s *graphController) Handleplayground(c echo.Context) error {
	playground.Handler("GraphQL", "/query").ServeHTTP(c.Response(), c.Request())
	return nil
}
