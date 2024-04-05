# ApiHNSDomainAvailableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**ApiSubdomainStatus**](ApiSubdomainStatus.md) |  | [optional] 

## Methods

### NewApiHNSDomainAvailableResponse

`func NewApiHNSDomainAvailableResponse() *ApiHNSDomainAvailableResponse`

NewApiHNSDomainAvailableResponse instantiates a new ApiHNSDomainAvailableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHNSDomainAvailableResponseWithDefaults

`func NewApiHNSDomainAvailableResponseWithDefaults() *ApiHNSDomainAvailableResponse`

NewApiHNSDomainAvailableResponseWithDefaults instantiates a new ApiHNSDomainAvailableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ApiHNSDomainAvailableResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiHNSDomainAvailableResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiHNSDomainAvailableResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiHNSDomainAvailableResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetStatus

`func (o *ApiHNSDomainAvailableResponse) GetStatus() ApiSubdomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiHNSDomainAvailableResponse) GetStatusOk() (*ApiSubdomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiHNSDomainAvailableResponse) SetStatus(v ApiSubdomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiHNSDomainAvailableResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


