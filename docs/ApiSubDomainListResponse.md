# ApiSubDomainListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to [**[]ApiSubDomainItem**](ApiSubDomainItem.md) |  | [optional] 
**NextPaginationToken** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiSubDomainListResponse

`func NewApiSubDomainListResponse() *ApiSubDomainListResponse`

NewApiSubDomainListResponse instantiates a new ApiSubDomainListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubDomainListResponseWithDefaults

`func NewApiSubDomainListResponseWithDefaults() *ApiSubDomainListResponse`

NewApiSubDomainListResponseWithDefaults instantiates a new ApiSubDomainListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *ApiSubDomainListResponse) GetNames() []ApiSubDomainItem`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *ApiSubDomainListResponse) GetNamesOk() (*[]ApiSubDomainItem, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *ApiSubDomainListResponse) SetNames(v []ApiSubDomainItem)`

SetNames sets Names field to given value.

### HasNames

`func (o *ApiSubDomainListResponse) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetNextPaginationToken

`func (o *ApiSubDomainListResponse) GetNextPaginationToken() string`

GetNextPaginationToken returns the NextPaginationToken field if non-nil, zero value otherwise.

### GetNextPaginationTokenOk

`func (o *ApiSubDomainListResponse) GetNextPaginationTokenOk() (*string, bool)`

GetNextPaginationTokenOk returns a tuple with the NextPaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaginationToken

`func (o *ApiSubDomainListResponse) SetNextPaginationToken(v string)`

SetNextPaginationToken sets NextPaginationToken field to given value.

### HasNextPaginationToken

`func (o *ApiSubDomainListResponse) HasNextPaginationToken() bool`

HasNextPaginationToken returns a boolean if a field has been set.

### GetTotal

`func (o *ApiSubDomainListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiSubDomainListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiSubDomainListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiSubDomainListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


