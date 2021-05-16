package planet

import "github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"

func (c *controller) SetupRouter(s *server.Server) {
	s.E.POST("/planets", c.Create)
}
