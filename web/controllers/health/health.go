package health

import (
	"fmt"
	"net/http"

	pingservice "github.com/Rafaela314/Go-mongo-REST-Exemple/services/ping"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/controllers"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/server"
	"github.com/labstack/echo/v4"
)

type controller struct {
	s *server.Server
}

// New controller
func New() controllers.Controller {
	return &controller{s: nil}
}

func (c *controller) Health(ec echo.Context) error {
	return ec.String(http.StatusOK, "Alive.")
}

func (c *controller) Ping(ec echo.Context) error {
	v := new(ping)
	if err := ec.Bind(v); err != nil {
		return err
	}

	pingService := pingservice.New()
	p := v.ToPing()

	p = pingService.InjectTimestamp(p)

	return ec.JSON(http.StatusOK, pong{
		Msg:       fmt.Sprintf("Pong: %s", p.Name),
		Timestamp: p.Timestamp,
	})
}
