# ApiResolveWalletAddressByDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Records** | Pointer to [**[]ApiDNSRecord**](ApiDNSRecord.md) |  | [optional] 
**Wallets** | Pointer to [**[]ApiWalletData**](ApiWalletData.md) |  | [optional] 

## Methods

### NewApiResolveWalletAddressByDomainResponse

`func NewApiResolveWalletAddressByDomainResponse() *ApiResolveWalletAddressByDomainResponse`

NewApiResolveWalletAddressByDomainResponse instantiates a new ApiResolveWalletAddressByDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResolveWalletAddressByDomainResponseWithDefaults

`func NewApiResolveWalletAddressByDomainResponseWithDefaults() *ApiResolveWalletAddressByDomainResponse`

NewApiResolveWalletAddressByDomainResponseWithDefaults instantiates a new ApiResolveWalletAddressByDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ApiResolveWalletAddressByDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiResolveWalletAddressByDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiResolveWalletAddressByDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiResolveWalletAddressByDomainResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRecords

`func (o *ApiResolveWalletAddressByDomainResponse) GetRecords() []ApiDNSRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ApiResolveWalletAddressByDomainResponse) GetRecordsOk() (*[]ApiDNSRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ApiResolveWalletAddressByDomainResponse) SetRecords(v []ApiDNSRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ApiResolveWalletAddressByDomainResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetWallets

`func (o *ApiResolveWalletAddressByDomainResponse) GetWallets() []ApiWalletData`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *ApiResolveWalletAddressByDomainResponse) GetWalletsOk() (*[]ApiWalletData, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *ApiResolveWalletAddressByDomainResponse) SetWallets(v []ApiWalletData)`

SetWallets sets Wallets field to given value.

### HasWallets

`func (o *ApiResolveWalletAddressByDomainResponse) HasWallets() bool`

HasWallets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


