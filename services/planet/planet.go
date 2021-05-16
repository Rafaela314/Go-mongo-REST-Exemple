package planet

import "github.com/Rafaela314/Go-mongo-REST-Exemple/models"

// Planet public interface
type Planet interface {
	Create(planet models.Planet) (*models.Planet, error)
}

type service struct {
}

// NewPlanet initializes interface
func New() Planet {
	return &service{}
}

func (s *service) Create(planet models.Planet) (*models.Planet, error) {
	var response models.Planet

	response.Name = "teste"

	return &response, nil
}
