package planet

import (
	"net/http"

	services "github.com/Rafaela314/Go-mongo-REST-Exemple/services/planet"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/controllers"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type controller struct {
}

// New controller
func New() controllers.Controller {
	return &controller{}
}

// @Summary Create
// @Description Inserts a new planet
// @Tags planet
// @Produce json
// @Success 200 {object} "Returns a json object with the newly created planet."
// @Router /planets [post]

func (c *controller) Create(ec echo.Context) error {

	var planet planetPayload

	err := ec.Bind(planet)
	if err != nil {
		return xerrors.Errorf("On binding request body : %v", err)
	}

	newplanet, err := services.New().Create(planet.Planet)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	return ec.JSON(http.StatusCreated, newplanet)
}
