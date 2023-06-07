# Go API client for superlink

API for Superlink

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v0.0.1-alpha.6
- Package version: v0.0.1-alpha.6
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
*AccessTokenAPI* | [**CreateAccessToken**](docs/AccessTokenAPI.md#createaccesstoken) | **Post** /v1/access-token | Creates an admin token
*AccessTokenAPI* | [**DeleteAccessToken**](docs/AccessTokenAPI.md#deleteaccesstoken) | **Delete** /v1/access-token | Deletes an access token
*AccessTokenAPI* | [**ListAccessTokens**](docs/AccessTokenAPI.md#listaccesstokens) | **Get** /v1/access-token | Lists access tokens
*HealthAPI* | [**HealthCheck**](docs/HealthAPI.md#healthcheck) | **Get** /v1/health | Checks the health of the API
*ResolutionAPI* | [**ResolveDataByDomain**](docs/ResolutionAPI.md#resolvedatabydomain) | **Get** /v1/resolve/{domain} | Resolves wallets and DNS records for a domain


## Documentation For Models

 - [ApiAccessTokenCreateRequest](docs/ApiAccessTokenCreateRequest.md)
 - [ApiAccessTokenDeleteRequest](docs/ApiAccessTokenDeleteRequest.md)
 - [ApiAccessTokenResponse](docs/ApiAccessTokenResponse.md)
 - [ApiBadRequestResponse](docs/ApiBadRequestResponse.md)
 - [ApiDNSRecord](docs/ApiDNSRecord.md)
 - [ApiErrorResponse](docs/ApiErrorResponse.md)
 - [ApiInternalServerErrorResponse](docs/ApiInternalServerErrorResponse.md)
 - [ApiResolveWalletAddressByDomainResponse](docs/ApiResolveWalletAddressByDomainResponse.md)
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

