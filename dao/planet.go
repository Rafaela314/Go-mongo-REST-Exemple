package dao

import (
	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/xerrors"
	"gopkg.in/mgo.v2/bson"
)

// Planets collection
const (
	Planets = "planets"
)

// Planet dao interface
type Planet interface {
	InsertPlanet(planet models.Planet) error
	GetPlanet(id string) (*models.Planet, error)
	GetPlanetByName(name string) ([]models.Planet, error)
	DeletePlanet(planetID string) error
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

func (d *planet) GetPlanet(id string) (*models.Planet, error) {

	var result models.Planet

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, xerrors.Errorf("Converting string to primitive bson id: %v", err)
	}

	_ = d.GetCollection().FindOne(d.ctx, bson.M{"_id": oid}).Decode(&result)
	if err != nil {
		return nil, xerrors.Errorf("Unable to use mongo: %v", err)
	}

	return &result, err
}

func (d *planet) GetPlanetByName(name string) ([]models.Planet, error) {

	filter, err := d.GetCollection().Find(d.ctx, bson.M{"name": name})
	if err != nil {
		return nil, xerrors.Errorf("Unable to filter by param on mongo: %v", err)
	}

	var result []models.Planet

	err = filter.All(d.ctx, &result)
	if err != nil {
		return nil, xerrors.Errorf("Unable find planet by name: %v", err)
	}

	return result, nil
}

func (d *planet) DeletePlanet(planetID string) error {

	oid, err := primitive.ObjectIDFromHex(planetID)
	if err != nil {
		return xerrors.Errorf("Converting string to primitive bson id: %v", err)
	}

	_, err = d.GetCollection().DeleteMany(d.ctx, bson.M{"_id": oid})
	if err != nil {
		return xerrors.Errorf("Deleting planet: %v", err)
	}

	return nil
}
