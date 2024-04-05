# \PartnerAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](PartnerAPI.md#CreateAccessToken) | **Post** /v1/partner/access-token | Creates an admin token for a partner
[**DeleteAccessToken**](PartnerAPI.md#DeleteAccessToken) | **Delete** /v1/partner/access-token | Deletes an access token
[**ListAccessTokens**](PartnerAPI.md#ListAccessTokens) | **Get** /v1/partner/access-token | Lists access tokens for a partner
[**PartnerPurchases**](PartnerAPI.md#PartnerPurchases) | **Get** /v1/partner/purchases | Returns a list of all purchases



## CreateAccessToken

> ApiAccessTokenResponse CreateAccessToken(ctx).Request(request).Execute()

Creates an admin token for a partner



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
	resp, r, err := apiClient.PartnerAPI.CreateAccessToken(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerAPI.CreateAccessToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessToken`: ApiAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PartnerAPI.CreateAccessToken`: %v\n", resp)
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
	r, err := apiClient.PartnerAPI.DeleteAccessToken(context.Background()).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerAPI.DeleteAccessToken``: %v\n", err)
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

> ApiAccessTokenResponse ListAccessTokens(ctx).Execute()

Lists access tokens for a partner



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
	resp, r, err := apiClient.PartnerAPI.ListAccessTokens(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerAPI.ListAccessTokens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessTokens`: ApiAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `PartnerAPI.ListAccessTokens`: %v\n", resp)
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


## PartnerPurchases

> ApiPurchaseListResponse PartnerPurchases(ctx).Execute()

Returns a list of all purchases



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
	resp, r, err := apiClient.PartnerAPI.PartnerPurchases(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PartnerAPI.PartnerPurchases``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartnerPurchases`: ApiPurchaseListResponse
	fmt.Fprintf(os.Stdout, "Response from `PartnerAPI.PartnerPurchases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPartnerPurchasesRequest struct via the builder pattern


### Return type

[**ApiPurchaseListResponse**](ApiPurchaseListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

