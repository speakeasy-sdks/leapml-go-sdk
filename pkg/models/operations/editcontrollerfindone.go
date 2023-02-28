package operations

import (
	"github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
)

type EditControllerFindOnePathParams struct {
	EditID string `pathParam:"style=simple,explode=false,name=editId"`
}

type EditControllerFindOneSecurity struct {
	Bearer shared.SchemeBearer `security:"scheme,type=http,subtype=bearer"`
}

type EditControllerFindOneRequest struct {
	PathParams EditControllerFindOnePathParams
	Security   EditControllerFindOneSecurity
}

type EditControllerFindOneResponse struct {
	ContentType string
	EditEntity  *shared.EditEntity
	StatusCode  int
}
