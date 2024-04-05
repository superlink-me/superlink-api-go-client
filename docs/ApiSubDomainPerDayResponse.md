# ApiSubDomainPerDayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aggregates** | Pointer to [**[]ApiSubDomainPerDayCount**](ApiSubDomainPerDayCount.md) |  | [optional] 
**NextPaginationToken** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiSubDomainPerDayResponse

`func NewApiSubDomainPerDayResponse() *ApiSubDomainPerDayResponse`

NewApiSubDomainPerDayResponse instantiates a new ApiSubDomainPerDayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubDomainPerDayResponseWithDefaults

`func NewApiSubDomainPerDayResponseWithDefaults() *ApiSubDomainPerDayResponse`

NewApiSubDomainPerDayResponseWithDefaults instantiates a new ApiSubDomainPerDayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregates

`func (o *ApiSubDomainPerDayResponse) GetAggregates() []ApiSubDomainPerDayCount`

GetAggregates returns the Aggregates field if non-nil, zero value otherwise.

### GetAggregatesOk

`func (o *ApiSubDomainPerDayResponse) GetAggregatesOk() (*[]ApiSubDomainPerDayCount, bool)`

GetAggregatesOk returns a tuple with the Aggregates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregates

`func (o *ApiSubDomainPerDayResponse) SetAggregates(v []ApiSubDomainPerDayCount)`

SetAggregates sets Aggregates field to given value.

### HasAggregates

`func (o *ApiSubDomainPerDayResponse) HasAggregates() bool`

HasAggregates returns a boolean if a field has been set.

### GetNextPaginationToken

`func (o *ApiSubDomainPerDayResponse) GetNextPaginationToken() string`

GetNextPaginationToken returns the NextPaginationToken field if non-nil, zero value otherwise.

### GetNextPaginationTokenOk

`func (o *ApiSubDomainPerDayResponse) GetNextPaginationTokenOk() (*string, bool)`

GetNextPaginationTokenOk returns a tuple with the NextPaginationToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPaginationToken

`func (o *ApiSubDomainPerDayResponse) SetNextPaginationToken(v string)`

SetNextPaginationToken sets NextPaginationToken field to given value.

### HasNextPaginationToken

`func (o *ApiSubDomainPerDayResponse) HasNextPaginationToken() bool`

HasNextPaginationToken returns a boolean if a field has been set.

### GetTotal

`func (o *ApiSubDomainPerDayResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ApiSubDomainPerDayResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ApiSubDomainPerDayResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *ApiSubDomainPerDayResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


