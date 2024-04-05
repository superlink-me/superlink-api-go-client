# \SubdomainAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ParentdomainPurchase**](SubdomainAPI.md#ParentdomainPurchase) | **Post** /v1/parentdomains/{parentDomain}/buy | Purchases and sets up a parent domain for use with ens subdomains
[**ParentdomainSearch**](SubdomainAPI.md#ParentdomainSearch) | **Get** /v1/parentdomains/{parentDomain}/search | Returns a list of available parent-domains provided the preferred parent ddomain
[**ParentdomainValidation**](SubdomainAPI.md#ParentdomainValidation) | **Get** /v1/parentdomains/{parentDomain} | Validates if parent domain is correctly configured for use with ens subdomains
[**SubdomainAvailable**](SubdomainAPI.md#SubdomainAvailable) | **Get** /v1/parentdomains/{parentDomain}/subdomains/{subDomainName} | Returns subdomain availability
[**SubdomainClaimed**](SubdomainAPI.md#SubdomainClaimed) | **Get** /v1/parentdomains/{parentDomain}/claimed/{ethAddress} | Returns subdomain availability
[**SubdomainList**](SubdomainAPI.md#SubdomainList) | **Get** /v1/parentdomains/{parentDomain}/list | Paginates over all subdomains in descending order of the creation date
[**SubdomainMint**](SubdomainAPI.md#SubdomainMint) | **Post** /v1/parentdomains/{parentDomain}/subdomains/{subDomainName} | Creates a subdomain for provided parentdomain
[**SubdomainMintSig**](SubdomainAPI.md#SubdomainMintSig) | **Post** /v1/parentdomains/{parentDomain}/subdomains/{subDomainName}/mint-with-sig | Creates a subdomain for provided parentdomain with signature
[**SubdomainPerDay**](SubdomainAPI.md#SubdomainPerDay) | **Get** /v1/parentdomains/{parentDomain}/per-day | Paginates over per-day aggregated counts for subdomains created given a parentdomain
[**SubdomainResolve**](SubdomainAPI.md#SubdomainResolve) | **Get** /v1/parentdomains/{parentDomain}/subdomains/{subDomainName}/resolve | Returns resolution of a provided subdomain
[**SubdomainReverseResolve**](SubdomainAPI.md#SubdomainReverseResolve) | **Get** /v1/parentdomains/{parentDomain}/addresses/{ethAddress}/reverse-resolve | Returns reverse-resolution of a provided eth address limited to provided parentdomain
[**SubdomainRevoke**](SubdomainAPI.md#SubdomainRevoke) | **Delete** /v1/parentdomains/{parentDomain}/subdomains/{subDomainName} | Deletes a subdomain for provided parentdomain (Omnipotent)
[**SubdomainTotal**](SubdomainAPI.md#SubdomainTotal) | **Get** /v1/parentdomains/{parentDomain}/total | Returns the total number of subdomains registered to parentdomain



## ParentdomainPurchase

> ParentdomainPurchase(ctx, parentDomain).Request(request).Execute()

Purchases and sets up a parent domain for use with ens subdomains



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
	parentDomain := "parentDomain_example" // string | superlink.me
	request := *openapiclient.NewApiParentDomainPurchaseRequest() // ApiParentDomainPurchaseRequest | purchase parent domain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubdomainAPI.ParentdomainPurchase(context.Background(), parentDomain).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.ParentdomainPurchase``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiParentdomainPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ApiParentDomainPurchaseRequest**](ApiParentDomainPurchaseRequest.md) | purchase parent domain | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentdomainSearch

> ApiParentDomainSearchResponse ParentdomainSearch(ctx, parentDomain).Execute()

Returns a list of available parent-domains provided the preferred parent ddomain



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
	parentDomain := "parentDomain_example" // string | superlink.me

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.ParentdomainSearch(context.Background(), parentDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.ParentdomainSearch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentdomainSearch`: ApiParentDomainSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.ParentdomainSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiParentdomainSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiParentDomainSearchResponse**](ApiParentDomainSearchResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParentdomainValidation

> ApiParentDomainValidationResponse ParentdomainValidation(ctx, parentDomain).Execute()

Validates if parent domain is correctly configured for use with ens subdomains



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
	parentDomain := "parentDomain_example" // string | superlink.me

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.ParentdomainValidation(context.Background(), parentDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.ParentdomainValidation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParentdomainValidation`: ApiParentDomainValidationResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.ParentdomainValidation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiParentdomainValidationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiParentDomainValidationResponse**](ApiParentDomainValidationResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainAvailable

> ApiSubDomainAvailableResponse SubdomainAvailable(ctx, parentDomain, subDomainName).Execute()

Returns subdomain availability



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
	parentDomain := "parentDomain_example" // string | superlink.me
	subDomainName := "subDomainName_example" // string | test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainAvailable(context.Background(), parentDomain, subDomainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainAvailable`: ApiSubDomainAvailableResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**subDomainName** | **string** | test | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSubDomainAvailableResponse**](ApiSubDomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainClaimed

> ApiSubDomainAvailableResponse SubdomainClaimed(ctx, parentDomain, ethAddress).Execute()

Returns subdomain availability



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
	parentDomain := "parentDomain_example" // string | superlink.me
	ethAddress := "ethAddress_example" // string | 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainClaimed(context.Background(), parentDomain, ethAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainClaimed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainClaimed`: ApiSubDomainAvailableResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainClaimed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**ethAddress** | **string** | 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainClaimedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSubDomainAvailableResponse**](ApiSubDomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainList

> ApiSubDomainListResponse SubdomainList(ctx, parentDomain).PaginationToken(paginationToken).PageSize(pageSize).Execute()

Paginates over all subdomains in descending order of the creation date



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
	parentDomain := "parentDomain_example" // string | superlink.me
	paginationToken := "paginationToken_example" // string | KCJuYW1lIjoiYWJjLmRlZi5naGkiLCJkYXRlIjoiMjAyMS0xMS0xMSIp (optional)
	pageSize := int32(56) // int32 | 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainList(context.Background(), parentDomain).PaginationToken(paginationToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainList`: ApiSubDomainListResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationToken** | **string** | KCJuYW1lIjoiYWJjLmRlZi5naGkiLCJkYXRlIjoiMjAyMS0xMS0xMSIp | 
 **pageSize** | **int32** | 100 | 

### Return type

[**ApiSubDomainListResponse**](ApiSubDomainListResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainMint

> SubdomainMint(ctx, parentDomain, subDomainName).Request(request).Execute()

Creates a subdomain for provided parentdomain



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
	parentDomain := "parentDomain_example" // string | superlink.me
	subDomainName := "subDomainName_example" // string | test
	request := *openapiclient.NewApiSubDomainMintRequest() // ApiSubDomainMintRequest | create subdomain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubdomainAPI.SubdomainMint(context.Background(), parentDomain, subDomainName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainMint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**subDomainName** | **string** | test | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainMintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ApiSubDomainMintRequest**](ApiSubDomainMintRequest.md) | create subdomain | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainMintSig

> SubdomainMintSig(ctx, parentDomain, subDomainName).Request(request).Execute()

Creates a subdomain for provided parentdomain with signature



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
	parentDomain := "parentDomain_example" // string | superlink.me
	subDomainName := "subDomainName_example" // string | test
	request := *openapiclient.NewApiSubDomainMintWithSigRequest() // ApiSubDomainMintWithSigRequest | create subdomain

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubdomainAPI.SubdomainMintSig(context.Background(), parentDomain, subDomainName).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainMintSig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**subDomainName** | **string** | test | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainMintSigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**ApiSubDomainMintWithSigRequest**](ApiSubDomainMintWithSigRequest.md) | create subdomain | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainPerDay

> ApiSubDomainPerDayResponse SubdomainPerDay(ctx, parentDomain).PaginationToken(paginationToken).PageSize(pageSize).Execute()

Paginates over per-day aggregated counts for subdomains created given a parentdomain



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
	parentDomain := "parentDomain_example" // string | superlink.me
	paginationToken := "paginationToken_example" // string | KCJuYW1lIjoiYWJjLmRlZi5naGkiLCJkYXRlIjoiMjAyMS0xMS0xMSIp (optional)
	pageSize := int32(56) // int32 | 100 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainPerDay(context.Background(), parentDomain).PaginationToken(paginationToken).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainPerDay``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainPerDay`: ApiSubDomainPerDayResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainPerDay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainPerDayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **paginationToken** | **string** | KCJuYW1lIjoiYWJjLmRlZi5naGkiLCJkYXRlIjoiMjAyMS0xMS0xMSIp | 
 **pageSize** | **int32** | 100 | 

### Return type

[**ApiSubDomainPerDayResponse**](ApiSubDomainPerDayResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainResolve

> ApiSubDomainResolveResponse SubdomainResolve(ctx, parentDomain, subDomainName).Execute()

Returns resolution of a provided subdomain



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
	parentDomain := "parentDomain_example" // string | superlink.me
	subDomainName := "subDomainName_example" // string | test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainResolve(context.Background(), parentDomain, subDomainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainResolve`: ApiSubDomainResolveResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**subDomainName** | **string** | test | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSubDomainResolveResponse**](ApiSubDomainResolveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainReverseResolve

> ApiSubDomainReverseResolveResponse SubdomainReverseResolve(ctx, parentDomain, ethAddress).Execute()

Returns reverse-resolution of a provided eth address limited to provided parentdomain



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
	parentDomain := "parentDomain_example" // string | superlink.me
	ethAddress := "ethAddress_example" // string | 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainReverseResolve(context.Background(), parentDomain, ethAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainReverseResolve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainReverseResolve`: ApiSubDomainReverseResolveResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainReverseResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**ethAddress** | **string** | 0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266 | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainReverseResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiSubDomainReverseResolveResponse**](ApiSubDomainReverseResolveResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubdomainRevoke

> SubdomainRevoke(ctx, parentDomain, subDomainName).Execute()

Deletes a subdomain for provided parentdomain (Omnipotent)



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
	parentDomain := "parentDomain_example" // string | superlink.me
	subDomainName := "subDomainName_example" // string | test

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SubdomainAPI.SubdomainRevoke(context.Background(), parentDomain, subDomainName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainRevoke``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 
**subDomainName** | **string** | test | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## SubdomainTotal

> ApiSubDomainTotalResponse SubdomainTotal(ctx, parentDomain).Execute()

Returns the total number of subdomains registered to parentdomain



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
	parentDomain := "parentDomain_example" // string | superlink.me

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubdomainAPI.SubdomainTotal(context.Background(), parentDomain).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubdomainAPI.SubdomainTotal``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubdomainTotal`: ApiSubDomainTotalResponse
	fmt.Fprintf(os.Stdout, "Response from `SubdomainAPI.SubdomainTotal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parentDomain** | **string** | superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubdomainTotalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiSubDomainTotalResponse**](ApiSubDomainTotalResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

