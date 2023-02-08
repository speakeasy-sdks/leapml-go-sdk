package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type InferencesControllerFindAllPathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type InferencesControllerFindAllSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type InferencesControllerFindAllRequest struct {
	PathParams InferencesControllerFindAllPathParams
	Security   InferencesControllerFindAllSecurity
}

type InferencesControllerFindAllResponse struct {
	ContentType string
	StatusCode  int64
}
