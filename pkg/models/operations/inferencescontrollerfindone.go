package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type InferencesControllerFindOnePathParams struct {
	InferenceID string `pathParam:"style=simple,explode=false,name=inferenceId"`
	ModelID     string `pathParam:"style=simple,explode=false,name=modelId"`
}

type InferencesControllerFindOneSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type InferencesControllerFindOneRequest struct {
	PathParams InferencesControllerFindOnePathParams
	Security   InferencesControllerFindOneSecurity
}

type InferencesControllerFindOneResponse struct {
	ContentType     string
	InferenceEntity *shared.InferenceEntity
	StatusCode      int64
}
