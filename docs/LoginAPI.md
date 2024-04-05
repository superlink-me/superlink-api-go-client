# \LoginAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginWithLink**](LoginAPI.md#LoginWithLink) | **Get** /v1/login/link/{loginToken} | Login via a magic link
[**LoginWithWallet**](LoginAPI.md#LoginWithWallet) | **Post** /v1/login/wallet/{ownerAddress} | Login with a valid wallet
[**SendLoginEmail**](LoginAPI.md#SendLoginEmail) | **Post** /v1/login/email/{emailAddress} | Sends the e-mail that contains the magic link to login a specfic user



## LoginWithLink

> ApiAccessTokenResponse LoginWithLink(ctx, loginToken).ReturnUrl(returnUrl).Execute()

Login via a magic link



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
	loginToken := "loginToken_example" // string | ycA4dEverAzydH0Df2si5QAqoFFeYJxDdszP2WIBEQ6rx%2FT98aJQ7HZSAl7abC1f%2BDYBlL6bQxOrr%2FVHgabyqC4ln1c8L%2BmkDRFC9QsR67kcW%2BZc92KAv9eGKdOyHgIjarjUw3FmQSeQpKsAtGytVRBzhVHLHQ%3D%3D
	returnUrl := "returnUrl_example" // string | /superlinks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LoginAPI.LoginWithLink(context.Background(), loginToken).ReturnUrl(returnUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.LoginWithLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LoginWithLink`: ApiAccessTokenResponse
	fmt.Fprintf(os.Stdout, "Response from `LoginAPI.LoginWithLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**loginToken** | **string** | ycA4dEverAzydH0Df2si5QAqoFFeYJxDdszP2WIBEQ6rx%2FT98aJQ7HZSAl7abC1f%2BDYBlL6bQxOrr%2FVHgabyqC4ln1c8L%2BmkDRFC9QsR67kcW%2BZc92KAv9eGKdOyHgIjarjUw3FmQSeQpKsAtGytVRBzhVHLHQ%3D%3D | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoginWithLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnUrl** | **string** | /superlinks | 

### Return type

[**ApiAccessTokenResponse**](ApiAccessTokenResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LoginWithWallet

> LoginWithWallet(ctx, ownerAddress).Request(request).Execute()

Login with a valid wallet



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
	ownerAddress := "ownerAddress_example" // string | 0xA5D70E12348Fef6A123EBD1231b123c51235E321
	request := *openapiclient.NewApiWalletLoginRequest() // ApiWalletLoginRequest | wallet login request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoginAPI.LoginWithWallet(context.Background(), ownerAddress).Request(request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.LoginWithWallet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerAddress** | **string** | 0xA5D70E12348Fef6A123EBD1231b123c51235E321 | 

### Other Parameters

Other parameters are passed through a pointer to a apiLoginWithWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**ApiWalletLoginRequest**](ApiWalletLoginRequest.md) | wallet login request | 

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


## SendLoginEmail

> SendLoginEmail(ctx, emailAddress).ReturnUrl(returnUrl).Execute()

Sends the e-mail that contains the magic link to login a specfic user



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
	emailAddress := "emailAddress_example" // string | test@superlink.me
	returnUrl := "returnUrl_example" // string | /superlinks (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LoginAPI.SendLoginEmail(context.Background(), emailAddress).ReturnUrl(returnUrl).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LoginAPI.SendLoginEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**emailAddress** | **string** | test@superlink.me | 

### Other Parameters

Other parameters are passed through a pointer to a apiSendLoginEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnUrl** | **string** | /superlinks | 

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

