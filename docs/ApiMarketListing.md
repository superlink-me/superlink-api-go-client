# ApiMarketListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**NameService** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **map[string]float32** |  | [optional] 

## Methods

### NewApiMarketListing

`func NewApiMarketListing() *ApiMarketListing`

NewApiMarketListing instantiates a new ApiMarketListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMarketListingWithDefaults

`func NewApiMarketListingWithDefaults() *ApiMarketListing`

NewApiMarketListingWithDefaults instantiates a new ApiMarketListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *ApiMarketListing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ApiMarketListing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ApiMarketListing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ApiMarketListing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDomain

`func (o *ApiMarketListing) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiMarketListing) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiMarketListing) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiMarketListing) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNameService

`func (o *ApiMarketListing) GetNameService() string`

GetNameService returns the NameService field if non-nil, zero value otherwise.

### GetNameServiceOk

`func (o *ApiMarketListing) GetNameServiceOk() (*string, bool)`

GetNameServiceOk returns a tuple with the NameService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameService

`func (o *ApiMarketListing) SetNameService(v string)`

SetNameService sets NameService field to given value.

### HasNameService

`func (o *ApiMarketListing) HasNameService() bool`

HasNameService returns a boolean if a field has been set.

### GetPrice

`func (o *ApiMarketListing) GetPrice() map[string]float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ApiMarketListing) GetPriceOk() (*map[string]float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ApiMarketListing) SetPrice(v map[string]float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ApiMarketListing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


