package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type RetrieveSingleModelPathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type RetrieveSingleModelSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type RetrieveSingleModelRequest struct {
	PathParams RetrieveSingleModelPathParams
	Security   RetrieveSingleModelSecurity
}

type RetrieveSingleModelResponse struct {
	ContentType string
	ModelEntity *shared.ModelEntity
	StatusCode  int
}
