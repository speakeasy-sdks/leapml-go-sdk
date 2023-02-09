package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type SamplesControllerCreateURLPathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type SamplesControllerCreateURLSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type SamplesControllerCreateURLRequest struct {
	PathParams SamplesControllerCreateURLPathParams
	Request    shared.UploadSamplesViaURLDto `request:"mediaType=application/json"`
	Security   SamplesControllerCreateURLSecurity
}

type SamplesControllerCreateURLResponse struct {
	ContentType          string
	StatusCode           int64
	TrainingSampleEntity *shared.TrainingSampleEntity
}