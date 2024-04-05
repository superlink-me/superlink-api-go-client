# \DomainInfoAPI

All URIs are relative to *https://api.superlink.me*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllDomains**](DomainInfoAPI.md#GetAllDomains) | **Get** /v1/domain | Returns all the domains owned by a given ownerAddress



## GetAllDomains

> ApiGetAllDomainsResponse GetAllDomains(ctx).Execute()

Returns all the domains owned by a given ownerAddress



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
	resp, r, err := apiClient.DomainInfoAPI.GetAllDomains(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DomainInfoAPI.GetAllDomains``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDomains`: ApiGetAllDomainsResponse
	fmt.Fprintf(os.Stdout, "Response from `DomainInfoAPI.GetAllDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDomainsRequest struct via the builder pattern


### Return type

[**ApiGetAllDomainsResponse**](ApiGetAllDomainsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

