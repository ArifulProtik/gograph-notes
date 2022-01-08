package server

func (s *fiberServer) RouteMapper() {
	s.fiber.Group("/api/v1")

	// Every Controllers For This Router Goes Here

}
