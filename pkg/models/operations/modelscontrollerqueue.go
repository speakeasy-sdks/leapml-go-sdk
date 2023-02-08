package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type ModelsControllerQueuePathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type ModelsControllerQueueSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type ModelsControllerQueueRequest struct {
	PathParams ModelsControllerQueuePathParams
	Request    *shared.TrainModelDto `request:"mediaType=application/json"`
	Security   ModelsControllerQueueSecurity
}

type ModelsControllerQueueResponse struct {
	ContentType        string
	ModelVersionEntity *shared.ModelVersionEntity
	StatusCode         int64
}
