# \ResolutionApiAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResolveDataByDomain**](ResolutionApiAPI.md#ResolveDataByDomain) | **Get** /v1/resolve/{domain} | Resolves wallets and records given a domain



## ResolveDataByDomain

> ApiResolveWalletAddressByDomainResponse ResolveDataByDomain(ctx, domain).Execute()

Resolves wallets and records given a domain



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
    domain := "domain_example" // string | firstname.lastname

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResolutionApiAPI.ResolveDataByDomain(context.Background(), domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResolutionApiAPI.ResolveDataByDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveDataByDomain`: ApiResolveWalletAddressByDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ResolutionApiAPI.ResolveDataByDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domain** | **string** | firstname.lastname | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveDataByDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiResolveWalletAddressByDomainResponse**](ApiResolveWalletAddressByDomainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

