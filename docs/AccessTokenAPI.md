# \AccessTokenAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AccessTokenAPI.md#CreateAccessToken) | **Post** /v1/access-token | Creates an admin token
[**ListAccessTokens**](AccessTokenAPI.md#ListAccessTokens) | **Get** /v1/access-token | Lists access tokens



## CreateAccessToken

> ApiAccessTokenResponse CreateAccessToken(ctx).Request(request).Execute()

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
    // response from `CreateAccessToken`: ApiAccessTokenResponse
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

[**ApiAccessTokenResponse**](ApiAccessTokenResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessTokens

> ApiAccessTokenResponse ListAccessTokens(ctx).Execute()

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
    // response from `ListAccessTokens`: ApiAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AccessTokenAPI.ListAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessTokensRequest struct via the builder pattern


### Return type

[**ApiAccessTokenResponse**](ApiAccessTokenResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

