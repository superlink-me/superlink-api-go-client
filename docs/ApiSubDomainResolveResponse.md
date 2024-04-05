# ApiSubDomainResolveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**NameData** | Pointer to [**ApiSubDomainNameData**](ApiSubDomainNameData.md) |  | [optional] 

## Methods

### NewApiSubDomainResolveResponse

`func NewApiSubDomainResolveResponse() *ApiSubDomainResolveResponse`

NewApiSubDomainResolveResponse instantiates a new ApiSubDomainResolveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubDomainResolveResponseWithDefaults

`func NewApiSubDomainResolveResponseWithDefaults() *ApiSubDomainResolveResponse`

NewApiSubDomainResolveResponseWithDefaults instantiates a new ApiSubDomainResolveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *ApiSubDomainResolveResponse) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiSubDomainResolveResponse) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiSubDomainResolveResponse) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiSubDomainResolveResponse) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetNameData

`func (o *ApiSubDomainResolveResponse) GetNameData() ApiSubDomainNameData`

GetNameData returns the NameData field if non-nil, zero value otherwise.

### GetNameDataOk

`func (o *ApiSubDomainResolveResponse) GetNameDataOk() (*ApiSubDomainNameData, bool)`

GetNameDataOk returns a tuple with the NameData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameData

`func (o *ApiSubDomainResolveResponse) SetNameData(v ApiSubDomainNameData)`

SetNameData sets NameData field to given value.

### HasNameData

`func (o *ApiSubDomainResolveResponse) HasNameData() bool`

HasNameData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


