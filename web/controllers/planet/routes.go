package planet

import "github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"

func (c *controller) SetupRouter(s *server.Server) {
	s.E.POST("/planets", c.Create)
	s.E.GET("/planets/:id", c.Get)
	s.E.POST("/planets/name/:name", c.GetByName)
}
