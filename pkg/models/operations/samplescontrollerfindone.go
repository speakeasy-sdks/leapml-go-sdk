package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type SamplesControllerFindOnePathParams struct {
	ModelID  string `pathParam:"style=simple,explode=false,name=modelId"`
	SampleID string `pathParam:"style=simple,explode=false,name=sampleId"`
}

type SamplesControllerFindOneSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type SamplesControllerFindOneRequest struct {
	PathParams SamplesControllerFindOnePathParams
	Security   SamplesControllerFindOneSecurity
}

type SamplesControllerFindOneResponse struct {
	ContentType            string
	StatusCode             int
	TrainingSampleEntities []shared.TrainingSampleEntity
}
