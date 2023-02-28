package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type CreateModelSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type CreateModelRequest struct {
	Request  shared.CreateModelDto `request:"mediaType=application/json"`
	Security CreateModelSecurity
}

type CreateModelResponse struct {
	ContentType string
	ModelEntity *shared.ModelEntity
	StatusCode  int
}
