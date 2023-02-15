package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type ListAllModelsSecurity struct {
	Bearer  *shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
	Bearer1 *shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type ListAllModelsRequest struct {
	Security ListAllModelsSecurity
}

type ListAllModelsResponse struct {
	ContentType   string
	ModelEntities []shared.ModelEntity
	StatusCode    int64
}
