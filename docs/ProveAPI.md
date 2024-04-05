# \ProveAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEthWalletMessage**](ProveAPI.md#GetEthWalletMessage) | **Get** /v1/prove/wallet/eth/{ownerAddress} | Get message which should be signed by an eth wallet as proof of ownership



## GetEthWalletMessage

> ApiWalletMessageResponse GetEthWalletMessage(ctx, ownerAddress).Execute()

Get message which should be signed by an eth wallet as proof of ownership



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProveAPI.GetEthWalletMessage(context.Background(), ownerAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProveAPI.GetEthWalletMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEthWalletMessage`: ApiWalletMessageResponse
	fmt.Fprintf(os.Stdout, "Response from `ProveAPI.GetEthWalletMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ownerAddress** | **string** | 0xA5D70E12348Fef6A123EBD1231b123c51235E321 | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEthWalletMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApiWalletMessageResponse**](ApiWalletMessageResponse.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

