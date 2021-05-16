package models

// Planet represents a planet that appears in star wars movies.
type Planet struct {
	Name        string  `db:"name" json:"name" swaggertype:"string"`
	Weather     string  `db:"weather" json:"weather" swaggertype:"string"`
	Land        float64 `db:"land" json:"land" swaggertype:"number"`
	Appearances int     `db:"appearances" json:"appearances" swaggertype:"number"`
}
