package planet

import "github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"

func (c *controller) SetupRouter(s *server.Server) {
	s.E.POST("/planets", c.Create)
	s.E.GET("/planets/:id", c.Get)
	s.E.GET("/planets/name/:name", c.GetByName)
	s.E.DELETE("/planets/:id", c.Delete)
}
