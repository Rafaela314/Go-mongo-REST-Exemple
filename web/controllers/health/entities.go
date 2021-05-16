package health

import (
	"fmt"

	"github.com/Rafaela314/Go-mongo-REST-Exemple/models"
)

type ping struct {
	models.Ping
	Version string `json:"version"`
}

func (p *ping) ToPing() models.Ping {
	p.Ping.Name = fmt.Sprintf("Name: %s", p.Ping.Name)
	return p.Ping
}

type pong struct {
	Msg       string `json:"msg"`
	Timestamp int64
}
