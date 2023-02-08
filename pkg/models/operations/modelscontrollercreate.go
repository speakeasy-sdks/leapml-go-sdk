package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type ModelsControllerCreateSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type ModelsControllerCreateRequest struct {
	Request  shared.CreateModelDto `request:"mediaType=application/json"`
	Security ModelsControllerCreateSecurity
}

type ModelsControllerCreateResponse struct {
	ContentType string
	ModelEntity *shared.ModelEntity
	StatusCode  int64
}
