package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type VersionsControllerFindOnePathParams struct {
	ModelID   string `pathParam:"style=simple,explode=false,name=modelId"`
	VersionID string `pathParam:"style=simple,explode=false,name=versionId"`
}

type VersionsControllerFindOneSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type VersionsControllerFindOneRequest struct {
	PathParams VersionsControllerFindOnePathParams
	Security   VersionsControllerFindOneSecurity
}

type VersionsControllerFindOneResponse struct {
	ContentType        string
	ModelVersionEntity *shared.ModelVersionEntity
	StatusCode         int64
}
