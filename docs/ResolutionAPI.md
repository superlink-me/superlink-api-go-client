# \ResolutionAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ResolveDataByAddress**](ResolutionAPI.md#ResolveDataByAddress) | **Get** /v1/reverse/{address} | Resolves wallets and DNS records for an address
[**ResolveDataByDomain**](ResolutionAPI.md#ResolveDataByDomain) | **Get** /v1/resolve/{domain} | Resolves wallets and DNS records for a domain



## ResolveDataByAddress

> ApiResolveDomainResponse ResolveDataByAddress(ctx, address).Nameservice(nameservice).Execute()

Resolves wallets and DNS records for an address



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
    address := "address_example" // string | 0x1234561234556
    nameservice := "nameservice_example" // string | superlink (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResolutionAPI.ResolveDataByAddress(context.Background(), address).Nameservice(nameservice).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResolutionAPI.ResolveDataByAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveDataByAddress`: ApiResolveDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ResolutionAPI.ResolveDataByAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**address** | **string** | 0x1234561234556 | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveDataByAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nameservice** | **string** | superlink | 

### Return type

[**ApiResolveDomainResponse**](ApiResolveDomainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveDataByDomain

> ApiResolveDomainResponse ResolveDataByDomain(ctx, domain).Nameservices(nameservices).Coins(coins).Execute()

Resolves wallets and DNS records for a domain



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
    nameservices := []string{"Nameservices_example"} // []string | superlink,ens,ud (optional)
    coins := []string{"Coins_example"} // []string | BTC,ETH,MATIC (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ResolutionAPI.ResolveDataByDomain(context.Background(), domain).Nameservices(nameservices).Coins(coins).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ResolutionAPI.ResolveDataByDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResolveDataByDomain`: ApiResolveDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `ResolutionAPI.ResolveDataByDomain`: %v\n", resp)
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

 **nameservices** | **[]string** | superlink,ens,ud | 
 **coins** | **[]string** | BTC,ETH,MATIC | 

### Return type

[**ApiResolveDomainResponse**](ApiResolveDomainResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

