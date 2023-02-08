package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type ModelsControllerFindOnePathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type ModelsControllerFindOneSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type ModelsControllerFindOneRequest struct {
	PathParams ModelsControllerFindOnePathParams
	Security   ModelsControllerFindOneSecurity
}

type ModelsControllerFindOneResponse struct {
	ContentType string
	ModelEntity *shared.ModelEntity
	StatusCode  int64
}
