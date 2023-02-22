<!-- Start SDK Example Usage -->
```go
package main

import (
    "log"
    "github.com/speakeasy-sdks/leapml-go-sdk"
    "github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/shared"
    "github.com/speakeasy-sdks/leapml-go-sdk/pkg/models/operations"
)

func main() {
    s := leapml.New()
    
    req := operations.ModelsControllerRemoveRequest{
        Security: operations.ModelsControllerRemoveSecurity{
            Bearer: shared.SchemeBearer{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
        PathParams: operations.ModelsControllerRemovePathParams{
            ModelID: "unde",
        },
    }
    
    res, err := s.FineTuning.ModelsControllerRemove(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
```
<!-- End SDK Example Usage -->