package shared

import (
	"time"
)

type ModelVersionEntityStatusEnum string

const (
	ModelVersionEntityStatusEnumQueued     ModelVersionEntityStatusEnum = "queued"
	ModelVersionEntityStatusEnumProcessing ModelVersionEntityStatusEnum = "processing"
	ModelVersionEntityStatusEnumFinished   ModelVersionEntityStatusEnum = "finished"
	ModelVersionEntityStatusEnumFailed     ModelVersionEntityStatusEnum = "failed"
)

type ModelVersionEntity struct {
	CreatedAt time.Time                    `json:"createdAt"`
	ID        string                       `json:"id"`
	Model     ModelEntity                  `json:"model"`
	Status    ModelVersionEntityStatusEnum `json:"status"`
	Weights   WeightsEntity                `json:"weights"`
}
