package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type InferencesControllerCreatePathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type InferencesControllerCreateSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type InferencesControllerCreateRequest struct {
	PathParams InferencesControllerCreatePathParams
	Request    shared.CreateInferenceDto `request:"mediaType=application/json"`
	Security   InferencesControllerCreateSecurity
}

type InferencesControllerCreateResponse struct {
	ContentType     string
	InferenceEntity *shared.InferenceEntity
	StatusCode      int
}
