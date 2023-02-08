package shared

import (
	"time"
)

type TrainingSampleEntity struct {
	CreatedAt time.Time `json:"createdAt"`
	ID        string    `json:"id"`
	URI       string    `json:"uri"`
}
