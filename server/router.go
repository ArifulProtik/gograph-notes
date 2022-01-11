package server

import "github.com/ArifulProtik/gograph-notes/controllers"

func (s *EchoServer) RouteMapper() {
	r := s.echo.Group("")

	// Every Controllers For This Router Goes Here
	graphcontroller := controllers.NewGraphController(s.logger, s.ent)
	r.POST("/query", graphcontroller.Handlequery)
	r.GET("/playground", graphcontroller.Handleplayground)
}
