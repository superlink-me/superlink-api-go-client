# \AccessTokenAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessTokenAPI.md#CreateAccessToken) | **Post** /v1/access-token | Creates an admin token
[**DeleteAccessToken**](AccessTokenAPI.md#DeleteAccessToken) | **Delete** /v1/access-token | Deletes an access token
[**ListAccessTokens**](AccessTokenAPI.md#ListAccessTokens) | **Get** /v1/access-token | Lists access tokens



## CreateAccessToken

> ApiAccessToken CreateAccessToken(ctx).Request(request).Execute()

Creates an admin token



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
    request := *openapiclient.NewApiAccessTokenCreateRequest() // ApiAccessTokenCreateRequest | access token create request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccessTokenAPI.CreateAccessToken(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.CreateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessToken`: ApiAccessToken
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiAccessTokenCreateRequest**](ApiAccessTokenCreateRequest.md) | access token create request | 

### Return type

[**ApiAccessToken**](ApiAccessToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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
    r, err := apiClient.AccessTokenAPI.DeleteAccessToken(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.DeleteAccessToken``: %v\n", err)
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


## ListAccessTokens

> ApiAccessToken ListAccessTokens(ctx).Execute()

Lists access tokens



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
    resp, r, err := apiClient.AccessTokenAPI.ListAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessTokenAPI.ListAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessTokens`: ApiAccessToken
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.ListAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessTokensRequest struct via the builder pattern


### Return type

[**ApiAccessToken**](ApiAccessToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

