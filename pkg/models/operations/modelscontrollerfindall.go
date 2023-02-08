package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type ModelsControllerFindAllSecurity struct {
	Bearer  *shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
	Bearer1 *shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type ModelsControllerFindAllRequest struct {
	Security ModelsControllerFindAllSecurity
}

type ModelsControllerFindAllResponse struct {
	ContentType   string
	ModelEntities []shared.ModelEntity
	StatusCode    int64
}
