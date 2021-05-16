package planet

import (
	"encoding/json"
	"fmt"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
	swapi "github.com/Rafaela314/Go-mongo-REST-Exemple/remote"
	"golang.org/x/xerrors"
)

const PLANET string = "planets"

// Planet public interface
type Planet interface {
	Create(planet models.Planet) (*models.Planet, error)
	GetPlanetInfos(name string) (*models.InfoResponse, error)
}

type service struct {
}

// NewPlanet initializes interface
func New() Planet {
	return &service{}
}

func (s *service) Create(planet models.Planet) (*models.Planet, error) {
	var response models.Planet

	info, err := s.GetPlanetInfos(planet.Name)

	if err != nil {
		return nil, xerrors.Errorf("could not find planet in swapi: %w", err)
	}

	response.Appearances = len(info.Results[0].Films)

	return &response, nil
}

func (s *service) GetPlanetInfos(name string) (*models.InfoResponse, error) {
	var response models.InfoResponse

	query := `?search=`

	sub := fmt.Sprintf("%s%s", query, name)

	info, err := swapi.New().Get(sub, PLANET)

	err = json.Unmarshal(info, &response)
	if err != nil {
		return nil, xerrors.Errorf("On body unmarshal: %w", err)
	}

	return &response, nil

}
