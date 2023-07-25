# \NftAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTokenImageByDomain**](NftAPI.md#GetTokenImageByDomain) | **Get** /v1/card-image/{domain}.svg | Returns a SVG image for a Superlink NFT
[**GetTokenMetadataByDomain**](NftAPI.md#GetTokenMetadataByDomain) | **Get** /v1/metadata/{domain} | Returns metadata usually associated with NFTs uri



## GetTokenImageByDomain

> string GetTokenImageByDomain(ctx, domain).Execute()

Returns a SVG image for a Superlink NFT



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
    domain := "domain_example" // string | firstname.lastname

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftAPI.GetTokenImageByDomain(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftAPI.GetTokenImageByDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenImageByDomain`: string
    fmt.Fprintf(os.Stdout, "Response from `NftAPI.GetTokenImageByDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | firstname.lastname | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenImageByDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTokenMetadataByDomain

> ApiDomainMetadataResponse GetTokenMetadataByDomain(ctx, domain).Execute()

Returns metadata usually associated with NFTs uri



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
    domain := "domain_example" // string | firstname.lastname

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NftAPI.GetTokenMetadataByDomain(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NftAPI.GetTokenMetadataByDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTokenMetadataByDomain`: ApiDomainMetadataResponse
    fmt.Fprintf(os.Stdout, "Response from `NftAPI.GetTokenMetadataByDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | firstname.lastname | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenMetadataByDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiDomainMetadataResponse**](ApiDomainMetadataResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

