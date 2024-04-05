# ApiMarketCryptoEstimationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Estimations** | Pointer to [**[]ApiMarketCryptoEstimation**](ApiMarketCryptoEstimation.md) |  | [optional] 
**FiatCurrency** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMarketCryptoEstimationResponse

`func NewApiMarketCryptoEstimationResponse() *ApiMarketCryptoEstimationResponse`

NewApiMarketCryptoEstimationResponse instantiates a new ApiMarketCryptoEstimationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMarketCryptoEstimationResponseWithDefaults

`func NewApiMarketCryptoEstimationResponseWithDefaults() *ApiMarketCryptoEstimationResponse`

NewApiMarketCryptoEstimationResponseWithDefaults instantiates a new ApiMarketCryptoEstimationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEstimations

`func (o *ApiMarketCryptoEstimationResponse) GetEstimations() []ApiMarketCryptoEstimation`

GetEstimations returns the Estimations field if non-nil, zero value otherwise.

### GetEstimationsOk

`func (o *ApiMarketCryptoEstimationResponse) GetEstimationsOk() (*[]ApiMarketCryptoEstimation, bool)`

GetEstimationsOk returns a tuple with the Estimations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimations

`func (o *ApiMarketCryptoEstimationResponse) SetEstimations(v []ApiMarketCryptoEstimation)`

SetEstimations sets Estimations field to given value.

### HasEstimations

`func (o *ApiMarketCryptoEstimationResponse) HasEstimations() bool`

HasEstimations returns a boolean if a field has been set.

### GetFiatCurrency

`func (o *ApiMarketCryptoEstimationResponse) GetFiatCurrency() string`

GetFiatCurrency returns the FiatCurrency field if non-nil, zero value otherwise.

### GetFiatCurrencyOk

`func (o *ApiMarketCryptoEstimationResponse) GetFiatCurrencyOk() (*string, bool)`

GetFiatCurrencyOk returns a tuple with the FiatCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiatCurrency

`func (o *ApiMarketCryptoEstimationResponse) SetFiatCurrency(v string)`

SetFiatCurrency sets FiatCurrency field to given value.

### HasFiatCurrency

`func (o *ApiMarketCryptoEstimationResponse) HasFiatCurrency() bool`

HasFiatCurrency returns a boolean if a field has been set.

### GetRegion

`func (o *ApiMarketCryptoEstimationResponse) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ApiMarketCryptoEstimationResponse) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ApiMarketCryptoEstimationResponse) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ApiMarketCryptoEstimationResponse) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


