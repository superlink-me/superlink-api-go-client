# Go API client for superlink

API for Superlink

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.3.13
- Package version: v0.3.13
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://superlink.me/about#contact-us](https://superlink.me/about#contact-us)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import superlink "github.com/superlink-me/superlink-api-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), superlink.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), superlink.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), superlink.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), superlink.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.superlink.me*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AdminAPI* | [**AdminCreateAccessToken**](docs/AdminAPI.md#admincreateaccesstoken) | **Post** /v1/admin/access-token | Creates an admin token
*AdminAPI* | [**AdminDeleteAccessToken**](docs/AdminAPI.md#admindeleteaccesstoken) | **Delete** /v1/admin/access-token | Deletes an access token
*AdminAPI* | [**AdminListAccessTokens**](docs/AdminAPI.md#adminlistaccesstokens) | **Get** /v1/admin/access-token | Lists access tokens
*AdminAPI* | [**AdminListPartners**](docs/AdminAPI.md#adminlistpartners) | **Get** /v1/admin/partner | Lists partners
*AdminAPI* | [**AdminPartnerCreate**](docs/AdminAPI.md#adminpartnercreate) | **Post** /v1/admin/partner | Creates a partner
*AdminAPI* | [**RemoveReverseResolutionAddress**](docs/AdminAPI.md#removereverseresolutionaddress) | **Delete** /v1/admin/reverse | Removes a reverse resolution address from a domain
*AdminAPI* | [**SetReverseResolutionAddress**](docs/AdminAPI.md#setreverseresolutionaddress) | **Post** /v1/admin/reverse | Assigns an address to a domain for reverse resolution
*DefaultAPI* | [**HealthCheck**](docs/DefaultAPI.md#healthcheck) | **Get** /v1/health | Checks the health of the API
*MarketAPI* | [**MarketPurchase**](docs/MarketAPI.md#marketpurchase) | **Post** /v1/market/purchase | Purchase returns the payment details required by Stripe
*MarketAPI* | [**MarketSearch**](docs/MarketAPI.md#marketsearch) | **Get** /v1/market/search/{query} | Returns market listings
*NftAPI* | [**GetTokenImageByDomain**](docs/NftAPI.md#gettokenimagebydomain) | **Get** /v1/card-image/{domain}.svg | Returns a SVG image for a Superlink NFT
*NftAPI* | [**GetTokenMetadataByDomain**](docs/NftAPI.md#gettokenmetadatabydomain) | **Get** /v1/metadata/{domain} | Returns metadata usually associated with NFTs uri
*PartnerAPI* | [**CreateAccessToken**](docs/PartnerAPI.md#createaccesstoken) | **Post** /v1/partner/access-token | Creates an admin token for a partner
*PartnerAPI* | [**DeleteAccessToken**](docs/PartnerAPI.md#deleteaccesstoken) | **Delete** /v1/partner/access-token | Deletes an access token
*PartnerAPI* | [**ListAccessTokens**](docs/PartnerAPI.md#listaccesstokens) | **Get** /v1/partner/access-token | Lists access tokens for a partner
*PartnerAPI* | [**PartnerPurchases**](docs/PartnerAPI.md#partnerpurchases) | **Get** /v1/partner/purchases | Returns a list of all purchases
*ResolutionAPI* | [**ResolveDataByAddress**](docs/ResolutionAPI.md#resolvedatabyaddress) | **Get** /v1/reverse/{address} | Resolves wallets and DNS records for an address
*ResolutionAPI* | [**ResolveDataByDomain**](docs/ResolutionAPI.md#resolvedatabydomain) | **Get** /v1/resolve/{domain} | Resolves wallets and DNS records for a domain


## Documentation For Models

 - [ApiAccessTokenCreateRequest](docs/ApiAccessTokenCreateRequest.md)
 - [ApiAccessTokenDeleteRequest](docs/ApiAccessTokenDeleteRequest.md)
 - [ApiAccessTokenResponse](docs/ApiAccessTokenResponse.md)
 - [ApiAdminAccessTokenCreateRequest](docs/ApiAdminAccessTokenCreateRequest.md)
 - [ApiAdminPartnerCreateRequest](docs/ApiAdminPartnerCreateRequest.md)
 - [ApiBadRequestResponse](docs/ApiBadRequestResponse.md)
 - [ApiCoin](docs/ApiCoin.md)
 - [ApiDNSRecord](docs/ApiDNSRecord.md)
 - [ApiDomainMetadataAttribute](docs/ApiDomainMetadataAttribute.md)
 - [ApiDomainMetadataResponse](docs/ApiDomainMetadataResponse.md)
 - [ApiErrorResponse](docs/ApiErrorResponse.md)
 - [ApiInternalServerErrorResponse](docs/ApiInternalServerErrorResponse.md)
 - [ApiMarketListing](docs/ApiMarketListing.md)
 - [ApiMarketPurchaseResponse](docs/ApiMarketPurchaseResponse.md)
 - [ApiMarketSearchResponse](docs/ApiMarketSearchResponse.md)
 - [ApiNameService](docs/ApiNameService.md)
 - [ApiPartnerResponse](docs/ApiPartnerResponse.md)
 - [ApiPurchaseListResponse](docs/ApiPurchaseListResponse.md)
 - [ApiPurchaseRequest](docs/ApiPurchaseRequest.md)
 - [ApiPurchaseResponse](docs/ApiPurchaseResponse.md)
 - [ApiResolveDomainResponse](docs/ApiResolveDomainResponse.md)
 - [ApiReverseResolutionDeleteRequest](docs/ApiReverseResolutionDeleteRequest.md)
 - [ApiSetReverseAddressRequest](docs/ApiSetReverseAddressRequest.md)
 - [ApiTXTRecord](docs/ApiTXTRecord.md)
 - [ApiWalletData](docs/ApiWalletData.md)
 - [DataAccessTokenType](docs/DataAccessTokenType.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### BearerAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

support@superlink.me

