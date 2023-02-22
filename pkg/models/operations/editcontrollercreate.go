package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type EditControllerCreateRequestBodyFiles struct {
	Content []byte `multipartForm:"content"`
	Files   string `multipartForm:"name=files"`
}

type EditControllerCreateRequestBody struct {
	Files              EditControllerCreateRequestBodyFiles `multipartForm:"file"`
	ImageGuidanceScale *float64                             `multipartForm:"name=imageGuidanceScale"`
	Prompt             string                               `multipartForm:"name=prompt"`
	Seed               *float64                             `multipartForm:"name=seed"`
	Steps              *float64                             `multipartForm:"name=steps"`
	TextGuidanceScale  *float64                             `multipartForm:"name=textGuidanceScale"`
	WebhookURL         *string                              `multipartForm:"name=webhookUrl"`
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
	StatusCode  int
}
