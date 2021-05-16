package dao

import (
	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
)

// Planets collection
const (
	Planets = "planets"
)

// Planet dao interface
type Planet interface {
	InsertPlanet(planet models.Planet) error
	//GetPlanet(id int) (*models.Planet, error)
	//GetPlanetByName(name string) (*models.Planet, error)
}

type planet struct {
	dao
}

// NewPlanet dao
func NewPlanet() Planet {
	init, _ := New(Planets)

	return &planet{
		dao{name: Planets, db: init.db}}
}

func (d *planet) InsertPlanet(planet models.Planet) error {

	return d.Save(planet, Planets)
}
