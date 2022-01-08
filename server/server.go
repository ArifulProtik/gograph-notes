package server

import (
	"github.com/ArifulProtik/gograph-notes/config"
	"github.com/ArifulProtik/gograph-notes/log"
	"github.com/gofiber/fiber/v2"
)

type Server interface {
	Run()
}
type fiberServer struct {
	config *config.Config
	fiber  *fiber.App
	logger log.Logger
}

func NewServer(cfg *config.Config, logger log.Logger) Server {
	return &fiberServer{
		config: cfg,
		fiber:  fiber.New(),
		logger: logger,
	}
}

// Run Stats a Server and Handles all the service of a server
func (s *fiberServer) Run() {
	s.RouteMapper()

	s.logger.Fatal(s.fiber.Listen(s.config.Port))
}
