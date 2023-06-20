# ApiWalletData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional] 
**Coin** | Pointer to [**ApiCoin**](ApiCoin.md) |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewApiWalletData

`func NewApiWalletData() *ApiWalletData`

NewApiWalletData instantiates a new ApiWalletData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWalletDataWithDefaults

`func NewApiWalletDataWithDefaults() *ApiWalletData`

NewApiWalletDataWithDefaults instantiates a new ApiWalletData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *ApiWalletData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *ApiWalletData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *ApiWalletData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *ApiWalletData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCoin

`func (o *ApiWalletData) GetCoin() ApiCoin`

GetCoin returns the Coin field if non-nil, zero value otherwise.

### GetCoinOk

`func (o *ApiWalletData) GetCoinOk() (*ApiCoin, bool)`

GetCoinOk returns a tuple with the Coin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoin

`func (o *ApiWalletData) SetCoin(v ApiCoin)`

SetCoin sets Coin field to given value.

### HasCoin

`func (o *ApiWalletData) HasCoin() bool`

HasCoin returns a boolean if a field has been set.

### GetNetwork

`func (o *ApiWalletData) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ApiWalletData) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ApiWalletData) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ApiWalletData) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetTag

`func (o *ApiWalletData) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *ApiWalletData) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *ApiWalletData) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *ApiWalletData) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetVersion

`func (o *ApiWalletData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiWalletData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiWalletData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiWalletData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


