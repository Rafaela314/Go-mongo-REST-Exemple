package settings

import "log"

// Settings for server struct
type Settings struct {
	SwapiURL string
}

var Setting Settings

// NewSettings returns a Settings struct
func NewSettings(swapiURL string) *Settings {

	if swapiURL == "" {
		log.Fatal("Configuring Application: config param PAYMENT_URL not found!", swapiURL)
	}

	Setting = Settings{
		SwapiURL: swapiURL,
	}
	return &Setting
}
