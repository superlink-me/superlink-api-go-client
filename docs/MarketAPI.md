# \MarketAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MarketCryptoEstimate**](MarketAPI.md#MarketCryptoEstimate) | **Get** /v1/market/crypto/estimate | CryptoEstimate returns the estimated conversion rates for supported crypto payment options
[**MarketCryptoPurchase**](MarketAPI.md#MarketCryptoPurchase) | **Post** /v1/market/crypto/purchase | CryptoPurchase returns the payment details required for crypto payment
[**MarketOrder**](MarketAPI.md#MarketOrder) | **Get** /v1/market/order/{orderID} | Returns order information
[**MarketPurchase**](MarketAPI.md#MarketPurchase) | **Post** /v1/market/purchase | Purchase returns the payment details required by Stripe
[**MarketSearch**](MarketAPI.md#MarketSearch) | **Post** /v1/market/search | Returns market listings
[**MarketSuggestion**](MarketAPI.md#MarketSuggestion) | **Post** /v1/market/suggest | Returns market listings for suggestions



## MarketCryptoEstimate

> ApiMarketCryptoEstimationResponse MarketCryptoEstimate(ctx).Execute()

CryptoEstimate returns the estimated conversion rates for supported crypto payment options



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
	resp, r, err := apiClient.MarketAPI.MarketCryptoEstimate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketCryptoEstimate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketCryptoEstimate`: ApiMarketCryptoEstimationResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketCryptoEstimate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiMarketCryptoEstimateRequest struct via the builder pattern


### Return type

[**ApiMarketCryptoEstimationResponse**](ApiMarketCryptoEstimationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketCryptoPurchase

> ApiMarketCryptoPurchaseResponse MarketCryptoPurchase(ctx).Request(request).Execute()

CryptoPurchase returns the payment details required for crypto payment



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
	request := *openapiclient.NewApiCryptoPurchaseRequest() // ApiCryptoPurchaseRequest | crypto purchase request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.MarketCryptoPurchase(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketCryptoPurchase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketCryptoPurchase`: ApiMarketCryptoPurchaseResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketCryptoPurchase`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketCryptoPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiCryptoPurchaseRequest**](ApiCryptoPurchaseRequest.md) | crypto purchase request | 

### Return type

[**ApiMarketCryptoPurchaseResponse**](ApiMarketCryptoPurchaseResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketOrder

> ApiMarketplaceOrderResponse MarketOrder(ctx, orderID).Execute()

Returns order information



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
	orderID := "orderID_example" // string | 92456d2b-c315-4b2b-b234-c674490b7324

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.MarketOrder(context.Background(), orderID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketOrder`: ApiMarketplaceOrderResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderID** | **string** | 92456d2b-c315-4b2b-b234-c674490b7324 | 

### Other Parameters

Other parameters are passed through a pointer to a apiMarketOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiMarketplaceOrderResponse**](ApiMarketplaceOrderResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketSearch

> ApiMarketSearchResponse MarketSearch(ctx).Request(request).Execute()

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
	request := *openapiclient.NewApiMarketSearchRequest() // ApiMarketSearchRequest | market search request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.MarketSearch(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketSearch`: ApiMarketSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiMarketSearchRequest**](ApiMarketSearchRequest.md) | market search request | 

### Return type

[**ApiMarketSearchResponse**](ApiMarketSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MarketSuggestion

> ApiMarketSearchResponse MarketSuggestion(ctx).Request(request).Execute()

Returns market listings for suggestions



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
	request := *openapiclient.NewApiMarketSuggestRequest() // ApiMarketSuggestRequest | market suggest request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MarketAPI.MarketSuggestion(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MarketAPI.MarketSuggestion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MarketSuggestion`: ApiMarketSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `MarketAPI.MarketSuggestion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMarketSuggestionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiMarketSuggestRequest**](ApiMarketSuggestRequest.md) | market suggest request | 

### Return type

[**ApiMarketSearchResponse**](ApiMarketSearchResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

