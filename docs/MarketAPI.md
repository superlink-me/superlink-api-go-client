# \MarketAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketPurchase**](MarketAPI.md#MarketPurchase) | **Post** /v1/market/purchase | Purchase returns the payment details required by Stripe
[**MarketSearch**](MarketAPI.md#MarketSearch) | **Get** /v1/market/search/{query} | Returns market listings



## MarketPurchase

> ApiMarketPurchaseResponse MarketPurchase(ctx).Request(request).Execute()

Purchase returns the payment details required by Stripe



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
    request := *openapiclient.NewApiPurchaseRequest() // ApiPurchaseRequest | purchase request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketAPI.MarketPurchase(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarketPurchase`: ApiMarketPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketPurchase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiPurchaseRequest**](ApiPurchaseRequest.md) | purchase request | 

### Return type

[**ApiMarketPurchaseResponse**](ApiMarketPurchaseResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketSearch

> ApiMarketSearchResponse MarketSearch(ctx, query).Execute()

Returns market listings



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
    query := "query_example" // string | johndoe

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MarketAPI.MarketSearch(context.Background(), query).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MarketSearch`: ApiMarketSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**query** | **string** | johndoe | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiMarketSearchResponse**](ApiMarketSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

