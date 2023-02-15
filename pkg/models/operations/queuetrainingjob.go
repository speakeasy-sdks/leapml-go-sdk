package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type QueueTrainingJobPathParams struct {
	ModelID string `pathParam:"style=simple,explode=false,name=modelId"`
}

type QueueTrainingJobSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type QueueTrainingJobRequest struct {
	PathParams QueueTrainingJobPathParams
	Request    *shared.TrainModelDto `request:"mediaType=application/json"`
	Security   QueueTrainingJobSecurity
}

type QueueTrainingJobResponse struct {
	ContentType        string
	ModelVersionEntity *shared.ModelVersionEntity
	StatusCode         int64
}
