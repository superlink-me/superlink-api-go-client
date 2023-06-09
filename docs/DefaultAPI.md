# \DefaultAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAccessToken**](DefaultAPI.md#DeleteAccessToken) | **Delete** /v1/access-token | Deletes an access token
[**HealthCheck**](DefaultAPI.md#HealthCheck) | **Get** /v1/health | Checks the health of the API



## DeleteAccessToken

> DeleteAccessToken(ctx).Request(request).Execute()

Deletes an access token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/superlink-me/superlink-api-go-client"
)

func main() {
    request := *openapiclient.NewApiAccessTokenDeleteRequest() // ApiAccessTokenDeleteRequest | access token delete request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.DeleteAccessToken(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.DeleteAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiAccessTokenDeleteRequest**](ApiAccessTokenDeleteRequest.md) | access token delete request | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    openapiclient "github.com/superlink-me/superlink-api-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DefaultAPI.HealthCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.HealthCheck``: %v\n", err)
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

