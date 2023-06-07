# ApiBadRequestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ApiErrorResponse**](ApiErrorResponse.md) |  | [optional] 

## Methods

### NewApiBadRequestResponse

`func NewApiBadRequestResponse() *ApiBadRequestResponse`

NewApiBadRequestResponse instantiates a new ApiBadRequestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiBadRequestResponseWithDefaults

`func NewApiBadRequestResponseWithDefaults() *ApiBadRequestResponse`

NewApiBadRequestResponseWithDefaults instantiates a new ApiBadRequestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *ApiBadRequestResponse) GetErrors() []ApiErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ApiBadRequestResponse) GetErrorsOk() (*[]ApiErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ApiBadRequestResponse) SetErrors(v []ApiErrorResponse)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ApiBadRequestResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


