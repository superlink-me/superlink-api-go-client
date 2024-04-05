# \HnsAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HnsDomainAvailable**](HnsAPI.md#HnsDomainAvailable) | **Get** /v1/hns/tld/{tld}/sld/{sld}/available | checks if a handshake domain is available to mint
[**HnsDomainCreate**](HnsAPI.md#HnsDomainCreate) | **Post** /v1/hns/tld/{tld}/sld/mint | HNSCreateDomain creates a handshake domain
[**HnsDomainCreateSig**](HnsAPI.md#HnsDomainCreateSig) | **Post** /v1/hns/tld/{tld}/sld/mint-with-sig | creates a handshake domain authenticated by a signed message
[**HnsDomainDelete**](HnsAPI.md#HnsDomainDelete) | **Delete** /v1/hns/tld/{tld}/sld/{sld} | deletes a handshake domain
[**HnsTldCheckClaimed**](HnsAPI.md#HnsTldCheckClaimed) | **Get** /v1/hns/tld/{tld}/check-if-claimed/{address} | tests if the wallet has already claimed a domain for this TLD.
[**HnsTldValidate**](HnsAPI.md#HnsTldValidate) | **Get** /v1/hns/tld/{tld}/validate | HNSTLDValidation tests if the TLD blockchain DNS records are valid.



## HnsDomainAvailable

> ApiHNSDomainAvailableResponse HnsDomainAvailable(ctx, tld, sld).Execute()

checks if a handshake domain is available to mint



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
	tld := "tld_example" // string | com
	sld := "sld_example" // string | google

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HnsAPI.HnsDomainAvailable(context.Background(), tld, sld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HnsAPI.HnsDomainAvailable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HnsDomainAvailable`: ApiHNSDomainAvailableResponse
	fmt.Fprintf(os.Stdout, "Response from `HnsAPI.HnsDomainAvailable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | com | 
**sld** | **string** | google | 

### Other Parameters

Other parameters are passed through a pointer to a apiHnsDomainAvailableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**ApiHNSDomainAvailableResponse**](ApiHNSDomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HnsDomainCreate

> ApiHNSDomainAvailableResponse HnsDomainCreate(ctx, tld).Request(request).Execute()

HNSCreateDomain creates a handshake domain



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
	tld := "tld_example" // string | com
	request := *openapiclient.NewApiHNSCreateDomainRequest() // ApiHNSCreateDomainRequest | register hns tld request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HnsAPI.HnsDomainCreate(context.Background(), tld).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HnsAPI.HnsDomainCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HnsDomainCreate`: ApiHNSDomainAvailableResponse
	fmt.Fprintf(os.Stdout, "Response from `HnsAPI.HnsDomainCreate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | com | 

### Other Parameters

Other parameters are passed through a pointer to a apiHnsDomainCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ApiHNSCreateDomainRequest**](ApiHNSCreateDomainRequest.md) | register hns tld request | 

### Return type

[**ApiHNSDomainAvailableResponse**](ApiHNSDomainAvailableResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HnsDomainCreateSig

> ApiHNSDomainAvailableResponse HnsDomainCreateSig(ctx, tld).Request(request).Execute()

creates a handshake domain authenticated by a signed message



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
	tld := "tld_example" // string | com
	request := *openapiclient.NewApiHNSCreateDomainWithWalletRequest() // ApiHNSCreateDomainWithWalletRequest | register hns domain request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HnsAPI.HnsDomainCreateSig(context.Background(), tld).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HnsAPI.HnsDomainCreateSig``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HnsDomainCreateSig`: ApiHNSDomainAvailableResponse
	fmt.Fprintf(os.Stdout, "Response from `HnsAPI.HnsDomainCreateSig`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | com | 

### Other Parameters

Other parameters are passed through a pointer to a apiHnsDomainCreateSigRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ApiHNSCreateDomainWithWalletRequest**](ApiHNSCreateDomainWithWalletRequest.md) | register hns domain request | 

### Return type

[**ApiHNSDomainAvailableResponse**](ApiHNSDomainAvailableResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HnsDomainDelete

> HnsDomainDelete(ctx, tld, sld).Execute()

deletes a handshake domain



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
	tld := "tld_example" // string | com
	sld := "sld_example" // string | google

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HnsAPI.HnsDomainDelete(context.Background(), tld, sld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HnsAPI.HnsDomainDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | com | 
**sld** | **string** | google | 

### Other Parameters

Other parameters are passed through a pointer to a apiHnsDomainDeleteRequest struct via the builder pattern


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


## HnsTldCheckClaimed

> HnsTldCheckClaimed(ctx, tld, address).Execute()

tests if the wallet has already claimed a domain for this TLD.



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
	tld := "tld_example" // string | com
	address := "address_example" // string | 0x123456789

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.HnsAPI.HnsTldCheckClaimed(context.Background(), tld, address).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HnsAPI.HnsTldCheckClaimed``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | com | 
**address** | **string** | 0x123456789 | 

### Other Parameters

Other parameters are passed through a pointer to a apiHnsTldCheckClaimedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HnsTldValidate

> ApiMarketPurchaseResponse HnsTldValidate(ctx, tld).Execute()

HNSTLDValidation tests if the TLD blockchain DNS records are valid.



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
	tld := "tld_example" // string | com

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HnsAPI.HnsTldValidate(context.Background(), tld).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HnsAPI.HnsTldValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HnsTldValidate`: ApiMarketPurchaseResponse
	fmt.Fprintf(os.Stdout, "Response from `HnsAPI.HnsTldValidate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tld** | **string** | com | 

### Other Parameters

Other parameters are passed through a pointer to a apiHnsTldValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiMarketPurchaseResponse**](ApiMarketPurchaseResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

