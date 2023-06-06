# ApiResponseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **string** |  | [optional] 
**Info** | Pointer to **string** |  | [optional] 

## Methods

### NewApiResponseError

`func NewApiResponseError() *ApiResponseError`

NewApiResponseError instantiates a new ApiResponseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponseErrorWithDefaults

`func NewApiResponseErrorWithDefaults() *ApiResponseError`

NewApiResponseErrorWithDefaults instantiates a new ApiResponseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ApiResponseError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ApiResponseError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ApiResponseError) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *ApiResponseError) HasField() bool`

HasField returns a boolean if a field has been set.

### GetInfo

`func (o *ApiResponseError) GetInfo() string`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *ApiResponseError) GetInfoOk() (*string, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *ApiResponseError) SetInfo(v string)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *ApiResponseError) HasInfo() bool`

HasInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


