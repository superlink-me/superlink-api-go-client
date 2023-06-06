# \HealthAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HealthCheck**](HealthAPI.md#HealthCheck) | **Get** /v1/health | Checks the health of the API



## HealthCheck

> HealthCheck(ctx).Execute()

Checks the health of the API



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/superlink-me/resolution-sdk-go"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.HealthAPI.HealthCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthAPI.HealthCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiHealthCheckRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

