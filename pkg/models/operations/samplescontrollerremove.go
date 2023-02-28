package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type SamplesControllerRemovePathParams struct {
	ModelID  string `pathParam:"style=simple,explode=false,name=modelId"`
	SampleID string `pathParam:"style=simple,explode=false,name=sampleId"`
}

type SamplesControllerRemoveSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type SamplesControllerRemoveRequest struct {
	PathParams SamplesControllerRemovePathParams
	Security   SamplesControllerRemoveSecurity
}

type SamplesControllerRemoveResponse struct {
	ContentType          string
	StatusCode           int
	TrainingSampleEntity *shared.TrainingSampleEntity
}
