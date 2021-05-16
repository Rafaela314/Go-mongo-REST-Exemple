package planet

import (
	"encoding/json"
	"fmt"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/dao"
	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
	swapi "github.com/Rafaela314/Go-mongo-REST-Exemple/remote"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	info, err := s.GetPlanetInfos(planet.Name)
	if err != nil {
		return nil, xerrors.Errorf("could not find planet in swapi: %w", err)
	}

	planet.Appearances = len(info.Results[0].Films)

	planet.ID = primitive.NewObjectID()

	err = dao.NewPlanet().InsertPlanet(planet)
	if err != nil {
		return nil, xerrors.Errorf("Saving planet on db: %v", err)
	}
	return &planet, nil
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
