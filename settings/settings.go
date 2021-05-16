package settings

import "log"

// Settings for server struct
type Settings struct {
	SwapiURL string
	MongoURL string
}

var Setting Settings

// NewSettings returns a Settings struct
func NewSettings(mongoURL, swapiURL string) *Settings {

	if swapiURL == "" {
		log.Fatal("Configuring Application: config param SWAPI_URL not found!", swapiURL)
	}

	if mongoURL == "" {
		log.Fatal("Configuring Application: config param MONGO_URL not found!", mongoURL)
	}

	Setting = Settings{
		SwapiURL: swapiURL,
		MongoURL: mongoURL,
	}
	return &Setting
}
