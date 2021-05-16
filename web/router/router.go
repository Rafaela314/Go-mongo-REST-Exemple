package router

import (
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/controllers/health"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/controllers/planet"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"
)

type Router struct {
	s *server.Server
}

func New(s *server.Server) Router {
	return Router{s: s}
}

func (r *Router) Setup() {
	health.New().SetupRouter(r.s)
	planet.New().SetupRouter(r.s)
}
