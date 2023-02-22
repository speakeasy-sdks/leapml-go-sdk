package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type EditControllerCreateWithURLRequestBody struct {
	ImageGuidanceScale *float64 `json:"imageGuidanceScale,omitempty"`
	ImageURL           string   `json:"imageUrl"`
	Prompt             string   `json:"prompt"`
	Seed               *float64 `json:"seed,omitempty"`
	Steps              *float64 `json:"steps,omitempty"`
	TextGuidanceScale  *float64 `json:"textGuidanceScale,omitempty"`
	WebhookURL         *string  `json:"webhookUrl,omitempty"`
}

type EditControllerCreateWithURLSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type EditControllerCreateWithURLRequest struct {
	Request  EditControllerCreateWithURLRequestBody `request:"mediaType=application/json"`
	Security EditControllerCreateWithURLSecurity
}

type EditControllerCreateWithURLResponse struct {
	ContentType string
	EditEntity  *shared.EditEntity
	StatusCode  int
}
