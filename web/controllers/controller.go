package controllers

import "github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"

type Controller interface {
	SetupRouter(s *server.Server)
}
