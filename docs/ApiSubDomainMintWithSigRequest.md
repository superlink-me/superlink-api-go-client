# ApiSubDomainMintWithSigRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**NameData** | Pointer to [**ApiSubDomainNameData**](ApiSubDomainNameData.md) |  | [optional] 
**SignedMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewApiSubDomainMintWithSigRequest

`func NewApiSubDomainMintWithSigRequest() *ApiSubDomainMintWithSigRequest`

NewApiSubDomainMintWithSigRequest instantiates a new ApiSubDomainMintWithSigRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSubDomainMintWithSigRequestWithDefaults

`func NewApiSubDomainMintWithSigRequestWithDefaults() *ApiSubDomainMintWithSigRequest`

NewApiSubDomainMintWithSigRequestWithDefaults instantiates a new ApiSubDomainMintWithSigRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiSubDomainMintWithSigRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiSubDomainMintWithSigRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiSubDomainMintWithSigRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiSubDomainMintWithSigRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetNameData

`func (o *ApiSubDomainMintWithSigRequest) GetNameData() ApiSubDomainNameData`

GetNameData returns the NameData field if non-nil, zero value otherwise.

### GetNameDataOk

`func (o *ApiSubDomainMintWithSigRequest) GetNameDataOk() (*ApiSubDomainNameData, bool)`

GetNameDataOk returns a tuple with the NameData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameData

`func (o *ApiSubDomainMintWithSigRequest) SetNameData(v ApiSubDomainNameData)`

SetNameData sets NameData field to given value.

### HasNameData

`func (o *ApiSubDomainMintWithSigRequest) HasNameData() bool`

HasNameData returns a boolean if a field has been set.

### GetSignedMessage

`func (o *ApiSubDomainMintWithSigRequest) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *ApiSubDomainMintWithSigRequest) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *ApiSubDomainMintWithSigRequest) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.

### HasSignedMessage

`func (o *ApiSubDomainMintWithSigRequest) HasSignedMessage() bool`

HasSignedMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


