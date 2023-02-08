package shared

import (
	"time"
)

type InferenceEntity struct {
	CreatedAt      time.Time `json:"createdAt"`
	Height         float64   `json:"height"`
	ID             string    `json:"id"`
	Images         []string  `json:"images"`
	ModelID        []string  `json:"modelId"`
	NegativePrompt string    `json:"negativePrompt"`
	NumberOfImages float64   `json:"numberOfImages"`
	Prompt         string    `json:"prompt"`
	PromptStrength float64   `json:"promptStrength"`
	Seed           float64   `json:"seed"`
	State          float64   `json:"state"`
	Steps          float64   `json:"steps"`
	Width          float64   `json:"width"`
}
