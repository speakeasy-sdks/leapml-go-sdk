package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type EditControllerCreateRequestBodyBody struct {
	ImageGuidanceScale *float64 `json:"imageGuidanceScale,omitempty"`
	Prompt             string   `json:"prompt"`
	Seed               *float64 `json:"seed,omitempty"`
	Steps              *float64 `json:"steps,omitempty"`
	TextGuidanceScale  *float64 `json:"textGuidanceScale,omitempty"`
	WebhookURL         *string  `json:"webhookUrl,omitempty"`
}

type EditControllerCreateRequestBodyFiles struct {
	Content []byte `multipartForm:"content"`
	Files   string `multipartForm:"name=files"`
}

type EditControllerCreateRequestBody struct {
	Body  EditControllerCreateRequestBodyBody  `multipartForm:"name=body,json"`
	Files EditControllerCreateRequestBodyFiles `multipartForm:"file"`
}

type EditControllerCreateSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type EditControllerCreateRequest struct {
	Request  EditControllerCreateRequestBody `request:"mediaType=multipart/form-data"`
	Security EditControllerCreateSecurity
}

type EditControllerCreateResponse struct {
	ContentType string
	EditEntity  *shared.EditEntity
	StatusCode  int64
}
