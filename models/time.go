package models

import (
	"time"
)

type Timestamp struct {
	Unix  int64     `json:"unix"`
	UTC   time.Time `json:"utc"`
	Local time.Time `json:"local"`
}
