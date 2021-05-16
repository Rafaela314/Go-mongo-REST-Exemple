package models

import "gopkg.in/mgo.v2/bson"

// Planet represents a planet that appears in star wars movies.
type Planet struct {
	ID          bson.ObjectId `db:"_id" json:"id" swaggertype:"string"`
	Name        string        `db:"name" json:"name" swaggertype:"string"`
	Weather     string        `db:"weather" json:"weather" swaggertype:"string"`
	Land        float64       `db:"land" json:"land" swaggertype:"number"`
	Appearances int           `db:"appearances" json:"appearances" swaggertype:"number"`
}

// PlanetInfosResponse represents additional information send by swapi api.
type PlanetResponse struct {
	Name           string   `json:"name"`
	RotationPeriod string   `json:"rotation_period"`
	OrbitalPeriod  string   `json:"orbital_period"`
	Diameter       string   `json:"diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   string   `json:"surface_water"`
	Population     string   `json:"population"`
	Residents      []string `json:"residents"`
	Films          []string `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            string   `json:"url"`
}

type InfoResponse struct {
	Count   int              `json:"count"`
	Results []PlanetResponse `json:"results"`
}
