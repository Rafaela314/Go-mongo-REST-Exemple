package ping

import (
	"time"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
)

type Ping interface {
	InjectTimestamp(models.Ping) models.Ping
}

type service struct{}

func New() Ping {
	return &service{}
}

func (s *service) InjectTimestamp(p models.Ping) models.Ping {
	t := time.Now()
	p.Timestamp = t.Unix()
	return p
}
