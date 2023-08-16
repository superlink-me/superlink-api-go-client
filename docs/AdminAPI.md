# \AdminAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminCreateAccessToken**](AdminAPI.md#AdminCreateAccessToken) | **Post** /v1/admin/access-token | Creates an admin token
[**AdminDeleteAccessToken**](AdminAPI.md#AdminDeleteAccessToken) | **Delete** /v1/admin/access-token | Deletes an access token
[**AdminListAccessTokens**](AdminAPI.md#AdminListAccessTokens) | **Get** /v1/admin/access-token | Lists access tokens
[**AdminListPartners**](AdminAPI.md#AdminListPartners) | **Get** /v1/admin/partner | Lists partners
[**AdminPartnerCreate**](AdminAPI.md#AdminPartnerCreate) | **Post** /v1/admin/partner | Creates a partner
[**RemoveReverseResolutionAddress**](AdminAPI.md#RemoveReverseResolutionAddress) | **Delete** /v1/admin/reverse | Removes a reverse resolution address from a domain
[**SetReverseResolutionAddress**](AdminAPI.md#SetReverseResolutionAddress) | **Post** /v1/admin/reverse | Assigns an address to a domain for reverse resolution



## AdminCreateAccessToken

> ApiAccessTokenResponse AdminCreateAccessToken(ctx).Request(request).Execute()

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
    request := *openapiclient.NewApiAdminAccessTokenCreateRequest() // ApiAdminAccessTokenCreateRequest | access token create request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.AdminCreateAccessToken(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminCreateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminCreateAccessToken`: ApiAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminCreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiAdminAccessTokenCreateRequest**](ApiAdminAccessTokenCreateRequest.md) | access token create request | 

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


## AdminDeleteAccessToken

> AdminDeleteAccessToken(ctx).Request(request).Execute()

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
    r, err := apiClient.AdminAPI.AdminDeleteAccessToken(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminDeleteAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminDeleteAccessTokenRequest struct via the builder pattern


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


## AdminListAccessTokens

> ApiAccessTokenResponse AdminListAccessTokens(ctx).Execute()

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
    resp, r, err := apiClient.AdminAPI.AdminListAccessTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminListAccessTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminListAccessTokens`: ApiAccessTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminListAccessTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListAccessTokensRequest struct via the builder pattern


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


## AdminListPartners

> ApiPartnerResponse AdminListPartners(ctx).Execute()

Lists partners



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
    resp, r, err := apiClient.AdminAPI.AdminListPartners(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminListPartners``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminListPartners`: ApiPartnerResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminListPartners`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiAdminListPartnersRequest struct via the builder pattern


### Return type

[**ApiPartnerResponse**](ApiPartnerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AdminPartnerCreate

> ApiPartnerResponse AdminPartnerCreate(ctx).Request(request).Execute()

Creates a partner



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
    request := *openapiclient.NewApiAdminPartnerCreateRequest() // ApiAdminPartnerCreateRequest | partner create request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminAPI.AdminPartnerCreate(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.AdminPartnerCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminPartnerCreate`: ApiPartnerResponse
    fmt.Fprintf(os.Stdout, "Response from `AdminAPI.AdminPartnerCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAdminPartnerCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiAdminPartnerCreateRequest**](ApiAdminPartnerCreateRequest.md) | partner create request | 

### Return type

[**ApiPartnerResponse**](ApiPartnerResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveReverseResolutionAddress

> RemoveReverseResolutionAddress(ctx).Request(request).Execute()

Removes a reverse resolution address from a domain



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
    request := *openapiclient.NewApiReverseResolutionDeleteRequest() // ApiReverseResolutionDeleteRequest | reverse address delete request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.RemoveReverseResolutionAddress(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.RemoveReverseResolutionAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveReverseResolutionAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiReverseResolutionDeleteRequest**](ApiReverseResolutionDeleteRequest.md) | reverse address delete request | 

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


## SetReverseResolutionAddress

> SetReverseResolutionAddress(ctx).Request(request).Execute()

Assigns an address to a domain for reverse resolution



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
    request := *openapiclient.NewApiSetReverseAddressRequest() // ApiSetReverseAddressRequest | set reverse address request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AdminAPI.SetReverseResolutionAddress(context.Background()).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminAPI.SetReverseResolutionAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetReverseResolutionAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **request** | [**ApiSetReverseAddressRequest**](ApiSetReverseAddressRequest.md) | set reverse address request | 

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

