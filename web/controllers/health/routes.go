package health

import "github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"

func (c *controller) SetupRouter(s *server.Server) {
	//system test
	s.E.GET("/system/alive", c.Health)
	s.E.POST("/system/ping", c.Ping)
	c.s = s
}
