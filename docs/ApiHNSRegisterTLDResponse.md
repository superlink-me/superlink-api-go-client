# ApiHNSRegisterTLDResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**ApiHNSBlockchainRecords**](ApiHNSBlockchainRecords.md) |  | [optional] 
**Tld** | Pointer to **string** |  | [optional] 

## Methods

### NewApiHNSRegisterTLDResponse

`func NewApiHNSRegisterTLDResponse() *ApiHNSRegisterTLDResponse`

NewApiHNSRegisterTLDResponse instantiates a new ApiHNSRegisterTLDResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHNSRegisterTLDResponseWithDefaults

`func NewApiHNSRegisterTLDResponseWithDefaults() *ApiHNSRegisterTLDResponse`

NewApiHNSRegisterTLDResponseWithDefaults instantiates a new ApiHNSRegisterTLDResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *ApiHNSRegisterTLDResponse) GetRecords() ApiHNSBlockchainRecords`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *ApiHNSRegisterTLDResponse) GetRecordsOk() (*ApiHNSBlockchainRecords, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *ApiHNSRegisterTLDResponse) SetRecords(v ApiHNSBlockchainRecords)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *ApiHNSRegisterTLDResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTld

`func (o *ApiHNSRegisterTLDResponse) GetTld() string`

GetTld returns the Tld field if non-nil, zero value otherwise.

### GetTldOk

`func (o *ApiHNSRegisterTLDResponse) GetTldOk() (*string, bool)`

GetTldOk returns a tuple with the Tld field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTld

`func (o *ApiHNSRegisterTLDResponse) SetTld(v string)`

SetTld sets Tld field to given value.

### HasTld

`func (o *ApiHNSRegisterTLDResponse) HasTld() bool`

HasTld returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


