# ApiResponseBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ApiResponseError**](ApiResponseError.md) |  | [optional] 

## Methods

### NewApiResponseBadRequest

`func NewApiResponseBadRequest() *ApiResponseBadRequest`

NewApiResponseBadRequest instantiates a new ApiResponseBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResponseBadRequestWithDefaults

`func NewApiResponseBadRequestWithDefaults() *ApiResponseBadRequest`

NewApiResponseBadRequestWithDefaults instantiates a new ApiResponseBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ApiResponseBadRequest) GetErrors() []ApiResponseError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiResponseBadRequest) GetErrorsOk() (*[]ApiResponseError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiResponseBadRequest) SetErrors(v []ApiResponseError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiResponseBadRequest) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


