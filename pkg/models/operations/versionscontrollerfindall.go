package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type VersionsControllerFindAllPathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type VersionsControllerFindAllSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type VersionsControllerFindAllRequest struct {
	PathParams VersionsControllerFindAllPathParams
	Security   VersionsControllerFindAllSecurity
}

type VersionsControllerFindAllResponse struct {
	ContentType          string
	ModelVersionEntities []shared.ModelVersionEntity
	StatusCode           int64
}
