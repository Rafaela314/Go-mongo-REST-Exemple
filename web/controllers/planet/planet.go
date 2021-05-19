package planet

import (
	"net/http"

	services "github.com/Rafaela314/Go-mongo-REST-Exemple/services/planet"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/settings"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/controllers"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/web/errors"

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
		return ec.JSON(http.StatusBadRequest, errors.TypedError{Code: errors.ErrorCreatingPlanet, Message: "Error creating a new planet. Check https://swapi.dev/api/planets/ to get the list of allwed planets ", Stack: planet})
	}

	return ec.JSON(http.StatusCreated, newplanet)
}

// @Summary GetPlanetByName
// @Description Returns a planet
// @Tags planet
// @Produce json
// @Param name path string true "The name of the planet"
// @Success 200 {object} Planet "Returns a json object with the requested planet."
// @Router /planets/name/{name} [get]

func (c *controller) GetByName(ec echo.Context) error {

	//get params
	name := ec.Param("name")

	planet, err := services.New().GetPlanetByName(name)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, errors.TypedError{Code: errors.ErrorFindingPlanetByName, Message: "No planet was found with this name, please verify planet name.", Stack: name})
	}

	return ec.JSON(http.StatusOK, planet)
}

// @Summary GetPlanet
// @Description Returns a planet
// @Tags planet
// @Produce json
// @Param id path string true "The ID of the planet"
// @Success 200 {object} Planet "Returns a json object with the requested planet."
// @Router /api/planet/{id} [get]
func (c *controller) Get(ec echo.Context) error {

	planetID := ec.Param("id")

	planet, err := services.New().GetPlanet(planetID)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, errors.TypedError{Code: errors.ErrorFindingPlanetById, Message: "No planet was found with this id, please verify planet id.", Stack: planetID})
	}

	return ec.JSON(http.StatusOK, planet)
}

// @Summary Delete
// @Description Deletes a planet
// @Tags planet
// @Produce json
// @Param id path string true "The ID of the planet to be deleted"
// @Param Authorization header string false "Should be 'Bearer (token)'"
// @Success 204 "Returns a 'no content' response."
// @Router /api/planet/{id} [delete]
func (c *controller) Delete(ec echo.Context) error {

	planetID := ec.Param("id")

	err := services.New().DeletePlanet(planetID)
	if err != nil {
		return ec.JSON(http.StatusBadRequest, errors.TypedError{Code: errors.ErrorDeletingPlanet, Message: "Could not find planet to delete. please verify the id.", Stack: planetID})
	}

	return ec.JSON(http.StatusOK, planetID)
}
