package shared

import (
	"time"
)

type EditEntity struct {
	CreatedAt          time.Time `json:"createdAt"`
	EditedImageURI     string    `json:"editedImageUri"`
	ID                 string    `json:"id"`
	ImageGuidanceScale float64   `json:"imageGuidanceScale"`
	ProjectID          string    `json:"projectId"`
	Prompt             string    `json:"prompt"`
	SourceImageURI     string    `json:"sourceImageUri"`
	Status             string    `json:"status"`
	Steps              float64   `json:"steps"`
	TextGuidanceScale  float64   `json:"textGuidanceScale"`
}
