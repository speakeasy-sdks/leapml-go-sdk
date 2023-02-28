package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type SamplesControllerCreatePathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type SamplesControllerCreateRequestBodyFiles struct {
	Content []byte `multipartForm:"content"`
	Files   string `multipartForm:"name=files"`
}

type SamplesControllerCreateRequestBody struct {
	Files *SamplesControllerCreateRequestBodyFiles `multipartForm:"file"`
}

type SamplesControllerCreateSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type SamplesControllerCreateRequest struct {
	PathParams SamplesControllerCreatePathParams
	Request    SamplesControllerCreateRequestBody `request:"mediaType=multipart/form-data"`
	Security   SamplesControllerCreateSecurity
}

type SamplesControllerCreateResponse struct {
	ContentType          string
	StatusCode           int
	TrainingSampleEntity *shared.TrainingSampleEntity
}
