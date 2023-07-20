# ApiResolveDomainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentHash** | Pointer to **string** |  | [optional] 
**DnsRecords** | Pointer to [**[]ApiDNSRecord**](ApiDNSRecord.md) |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**NameService** | Pointer to [**ApiNameService**](ApiNameService.md) |  | [optional] 
**OwnerAddress** | Pointer to **string** |  | [optional] 
**TxtRecords** | Pointer to [**[]ApiTXTRecord**](ApiTXTRecord.md) |  | [optional] 
**Wallets** | Pointer to [**[]ApiWalletData**](ApiWalletData.md) |  | [optional] 

## Methods

### NewApiResolveDomainResponse

`func NewApiResolveDomainResponse() *ApiResolveDomainResponse`

NewApiResolveDomainResponse instantiates a new ApiResolveDomainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResolveDomainResponseWithDefaults

`func NewApiResolveDomainResponseWithDefaults() *ApiResolveDomainResponse`

NewApiResolveDomainResponseWithDefaults instantiates a new ApiResolveDomainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentHash

`func (o *ApiResolveDomainResponse) GetContentHash() string`

GetContentHash returns the ContentHash field if non-nil, zero value otherwise.

### GetContentHashOk

`func (o *ApiResolveDomainResponse) GetContentHashOk() (*string, bool)`

GetContentHashOk returns a tuple with the ContentHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentHash

`func (o *ApiResolveDomainResponse) SetContentHash(v string)`

SetContentHash sets ContentHash field to given value.

### HasContentHash

`func (o *ApiResolveDomainResponse) HasContentHash() bool`

HasContentHash returns a boolean if a field has been set.

### GetDnsRecords

`func (o *ApiResolveDomainResponse) GetDnsRecords() []ApiDNSRecord`

GetDnsRecords returns the DnsRecords field if non-nil, zero value otherwise.

### GetDnsRecordsOk

`func (o *ApiResolveDomainResponse) GetDnsRecordsOk() (*[]ApiDNSRecord, bool)`

GetDnsRecordsOk returns a tuple with the DnsRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsRecords

`func (o *ApiResolveDomainResponse) SetDnsRecords(v []ApiDNSRecord)`

SetDnsRecords sets DnsRecords field to given value.

### HasDnsRecords

`func (o *ApiResolveDomainResponse) HasDnsRecords() bool`

HasDnsRecords returns a boolean if a field has been set.

### GetDomain

`func (o *ApiResolveDomainResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiResolveDomainResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiResolveDomainResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiResolveDomainResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNameService

`func (o *ApiResolveDomainResponse) GetNameService() ApiNameService`

GetNameService returns the NameService field if non-nil, zero value otherwise.

### GetNameServiceOk

`func (o *ApiResolveDomainResponse) GetNameServiceOk() (*ApiNameService, bool)`

GetNameServiceOk returns a tuple with the NameService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameService

`func (o *ApiResolveDomainResponse) SetNameService(v ApiNameService)`

SetNameService sets NameService field to given value.

### HasNameService

`func (o *ApiResolveDomainResponse) HasNameService() bool`

HasNameService returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *ApiResolveDomainResponse) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *ApiResolveDomainResponse) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *ApiResolveDomainResponse) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *ApiResolveDomainResponse) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetTxtRecords

`func (o *ApiResolveDomainResponse) GetTxtRecords() []ApiTXTRecord`

GetTxtRecords returns the TxtRecords field if non-nil, zero value otherwise.

### GetTxtRecordsOk

`func (o *ApiResolveDomainResponse) GetTxtRecordsOk() (*[]ApiTXTRecord, bool)`

GetTxtRecordsOk returns a tuple with the TxtRecords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxtRecords

`func (o *ApiResolveDomainResponse) SetTxtRecords(v []ApiTXTRecord)`

SetTxtRecords sets TxtRecords field to given value.

### HasTxtRecords

`func (o *ApiResolveDomainResponse) HasTxtRecords() bool`

HasTxtRecords returns a boolean if a field has been set.

### GetWallets

`func (o *ApiResolveDomainResponse) GetWallets() []ApiWalletData`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *ApiResolveDomainResponse) GetWalletsOk() (*[]ApiWalletData, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *ApiResolveDomainResponse) SetWallets(v []ApiWalletData)`

SetWallets sets Wallets field to given value.

### HasWallets

`func (o *ApiResolveDomainResponse) HasWallets() bool`

HasWallets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


