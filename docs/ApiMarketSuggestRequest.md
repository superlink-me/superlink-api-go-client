# ApiMarketSuggestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addrs** | Pointer to [**[]ApiAddressRecord**](ApiAddressRecord.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 

## Methods

### NewApiMarketSuggestRequest

`func NewApiMarketSuggestRequest() *ApiMarketSuggestRequest`

NewApiMarketSuggestRequest instantiates a new ApiMarketSuggestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMarketSuggestRequestWithDefaults

`func NewApiMarketSuggestRequestWithDefaults() *ApiMarketSuggestRequest`

NewApiMarketSuggestRequestWithDefaults instantiates a new ApiMarketSuggestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddrs

`func (o *ApiMarketSuggestRequest) GetAddrs() []ApiAddressRecord`

GetAddrs returns the Addrs field if non-nil, zero value otherwise.

### GetAddrsOk

`func (o *ApiMarketSuggestRequest) GetAddrsOk() (*[]ApiAddressRecord, bool)`

GetAddrsOk returns a tuple with the Addrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddrs

`func (o *ApiMarketSuggestRequest) SetAddrs(v []ApiAddressRecord)`

SetAddrs sets Addrs field to given value.

### HasAddrs

`func (o *ApiMarketSuggestRequest) HasAddrs() bool`

HasAddrs returns a boolean if a field has been set.

### GetQuery

`func (o *ApiMarketSuggestRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ApiMarketSuggestRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ApiMarketSuggestRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ApiMarketSuggestRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


