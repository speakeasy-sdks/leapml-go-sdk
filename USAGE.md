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
    
    req := operations.ModelsControllerCreateRequest{
        Security: operations.ModelsControllerCreateSecurity{
            Bearer: shared.SchemeBearer{
                Authorization: "Bearer YOUR_BEARER_TOKEN_HERE",
            },
        },
        Request: shared.CreateModelDto{
            SubjectIdentifier: "unde",
            SubjectKeyword: "deserunt",
            Title: "porro",
        },
    }
    
    res, err := s.FineTuning.ModelsControllerCreate(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModelEntity != nil {
        // handle response
    }
```
<!-- End SDK Example Usage -->