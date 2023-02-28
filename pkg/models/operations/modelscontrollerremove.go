package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type ModelsControllerRemovePathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type ModelsControllerRemoveSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type ModelsControllerRemoveRequest struct {
	PathParams ModelsControllerRemovePathParams
	Security   ModelsControllerRemoveSecurity
}

type ModelsControllerRemoveResponse struct {
	ContentType string
	StatusCode  int
}
