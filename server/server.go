package server

import (
	"github.com/ArifulProtik/gograph-notes/auth"
	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/ent"
	"github.com/ArifulProtik/gograph-notes/log"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Run()
}
type EchoServer struct {
	config *config.Config
	echo   *echo.Echo
	logger log.Logger
	ent    *ent.Client
}

func NewServer(cfg *config.Config, logger log.Logger, dbclient *ent.Client) Server {
	return &EchoServer{
		config: cfg,
		echo:   echo.New(),
		logger: logger,
		ent:    dbclient,
	}
}

// Run Stats a Server and Handles all the service of a server
func (s *EchoServer) Run() {
	s.echo.Use(auth.JWTMiddleware())

	// s.echo.Use(middleware.Recover())

	s.RouteMapper()

	s.logger.Fatal(s.echo.Start(s.config.Port))
}
