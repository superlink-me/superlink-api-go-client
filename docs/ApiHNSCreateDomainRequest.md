# ApiHNSCreateDomainRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addresses** | Pointer to **map[string]string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**OwnerAddress** | Pointer to **string** |  | [optional] 

## Methods

### NewApiHNSCreateDomainRequest

`func NewApiHNSCreateDomainRequest() *ApiHNSCreateDomainRequest`

NewApiHNSCreateDomainRequest instantiates a new ApiHNSCreateDomainRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiHNSCreateDomainRequestWithDefaults

`func NewApiHNSCreateDomainRequestWithDefaults() *ApiHNSCreateDomainRequest`

NewApiHNSCreateDomainRequestWithDefaults instantiates a new ApiHNSCreateDomainRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddresses

`func (o *ApiHNSCreateDomainRequest) GetAddresses() map[string]string`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *ApiHNSCreateDomainRequest) GetAddressesOk() (*map[string]string, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *ApiHNSCreateDomainRequest) SetAddresses(v map[string]string)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *ApiHNSCreateDomainRequest) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.

### GetDomain

`func (o *ApiHNSCreateDomainRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiHNSCreateDomainRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiHNSCreateDomainRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiHNSCreateDomainRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *ApiHNSCreateDomainRequest) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *ApiHNSCreateDomainRequest) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *ApiHNSCreateDomainRequest) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *ApiHNSCreateDomainRequest) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


