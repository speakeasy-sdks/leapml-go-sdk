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
    
    req := operations.SamplesControllerCreateRequest{
        Security: operations.SamplesControllerCreateSecurity{
            Bearer: shared.SchemeBearer{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
        PathParams: operations.SamplesControllerCreatePathParams{
            ModelID: "unde",
        },
        Request: operations.SamplesControllerCreateRequestBody{
            Files: &operations.SamplesControllerCreateRequestBodyFiles{
                Content: []byte("deserunt"),
                Files: "porro",
            },
        },
    }
    
    res, err := s.FineTuning.SamplesControllerCreate(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.TrainingSampleEntity != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->