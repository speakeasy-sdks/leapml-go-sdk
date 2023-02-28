package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type SamplesControllerFindAllPathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type SamplesControllerFindAllSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type SamplesControllerFindAllRequest struct {
	PathParams SamplesControllerFindAllPathParams
	Security   SamplesControllerFindAllSecurity
}

type SamplesControllerFindAllResponse struct {
	ContentType            string
	StatusCode             int
	TrainingSampleEntities []shared.TrainingSampleEntity
}
