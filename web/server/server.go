package server

import (
	"net/http"

	settings "github.com/Rafaela314/Go-mongo-REST-Exemple/settings"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/errors"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/middleware"
	"github.com/labstack/echo/v4"
)

//Server - wrap echo.Echo
type Server struct {
	E        *echo.Echo
	Settings *settings.Settings
}

//NewServer returns a new instance
func NewServer(e *echo.Echo, settings *settings.Settings) *Server {
	s := &Server{E: e, Settings: settings}
	e.HTTPErrorHandler = s.errorHandler

	return s
}

//Start initialize the instance
func (s *Server) Start(address string) error {
	return s.E.Start(address)
}

//ServeHHTP implements Handler interface
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.E.ServeHTTP(w, r)
}

func (s *Server) errorHandler(err error, c echo.Context) {
	if appEr, ok := err.(errors.AppError); ok {
		if err := c.JSON(appEr.HTTPCode, appEr); err != nil {
			s.E.Logger.Error(err)
		}

		return
	}

	c.Logger().Error(err)
	s.E.DefaultHTTPErrorHandler(err, c)
}

//SetCustomContextKey sets key for echo.Context
func (s *Server) SetCustomContextKey(key string, val interface{}) {
	s.E.Use(middleware.CustomContextMiddelware(key, val))
}
