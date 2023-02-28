package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type InferencesControllerRemovePathParams struct {
	InferenceID string `pathParam:"style=simple,explode=false,name=inferenceId"`
	ModelID     string `pathParam:"style=simple,explode=false,name=modelId"`
}

type InferencesControllerRemoveSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type InferencesControllerRemoveRequest struct {
	PathParams InferencesControllerRemovePathParams
	Security   InferencesControllerRemoveSecurity
}

type InferencesControllerRemoveResponse struct {
	ContentType string
	StatusCode  int
}
