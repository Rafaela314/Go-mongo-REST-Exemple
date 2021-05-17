package planet

import (
	"net/http"

	services "github.com/Rafaela314/Go-mongo-REST-Exemple/services/planet"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/settings"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/controllers"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
)

type controller struct {
	MongoURL string
}

// New controller
func New() controllers.Controller {
	return &controller{MongoURL: settings.Setting.MongoURL}
}

// @Summary Create
// @Description Inserts a new planet
// @Tags planet
// @Produce json
// @Success 200 {object} "Returns a json object with the newly created planet."
// @Router /planets [post]

func (c *controller) Create(ec echo.Context) error {

	planet := new(planetPayload)

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

// @Summary List
// @Description Returns all planet
// @Tags planet
// @Produce json
// @Success 200 {object} Planet "Returns a json object with the requested planet list."
// @Router /planets [get]
func (c *controller) List(ec echo.Context) error {

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

// @Summary GetPlanetByName
// @Description Returns a planet
// @Tags planet
// @Produce json
// @Param name path string true "The name of the planet"
// @Success 200 {object} Planet "Returns a json object with the requested planet."
// @Router /planets/{name} [get]

func (c *controller) GetByName(ec echo.Context) error {

	//get params
	name := ec.Param("name")

	planet, err := services.New().GetPlanetByName(name)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	return ec.JSON(http.StatusCreated, planet)
}

// @Summary GetPlanet
// @Description Returns a planet
// @Tags planet
// @Produce json
// @Param id path string true "The ID of the planet"
// @Success 200 {object} Planet "Returns a json object with the requested planet."
// @Router /api/company/{id} [get]
func (c *controller) Get(ec echo.Context) error {

	id := ec.Param("id")

	planet, err := services.New().GetPlanet(id)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, err)
	}

	return ec.JSON(http.StatusCreated, planet)
}

// @Summary Delete
// @Description Deletes a company
// @Tags company
// @Produce json
// @Param id path string true "The ID of the company to be deleted"
// @Param Authorization header string false "Should be 'Bearer (token)'"
// @Success 204 "Returns a 'no content' response."
// @Router /api/company/{id} [delete]
func (c *controller) Delete(ec echo.Context) error {

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
