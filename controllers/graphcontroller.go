package controllers

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ArifulProtik/gograph-notes/auth"
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
	logger      log.Logger
	dbclient    *ent.Client
	config      *config.Config
	echocontext echo.Context
}

func NewGraphController(logger log.Logger, dbclient *ent.Client, cfg *config.Config, ctx echo.Context) GraphController {
	return &graphController{
		logger:      logger,
		dbclient:    dbclient,
		config:      cfg,
		echocontext: ctx,
	}
}

func (s *graphController) Handlequery(c echo.Context) error {
	cc := c.(*auth.CustomContext)
	req := cc.Request()
	res := cc.Response()
	srv := handler.NewDefaultServer(graph.NewSchema(s.logger, s.dbclient, s.config, s.echocontext))
	srv.ServeHTTP(res, req)
	return nil
}
func (s *graphController) Handleplayground(c echo.Context) error {
	cc := c.(*auth.CustomContext)
	req := cc.Request()
	res := cc.Response()
	playground.Handler("GraphQL", "/query").ServeHTTP(res, req)
	return nil
}
