# ApiWalletLoginRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 
**SignedMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewApiWalletLoginRequest

`func NewApiWalletLoginRequest() *ApiWalletLoginRequest`

NewApiWalletLoginRequest instantiates a new ApiWalletLoginRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiWalletLoginRequestWithDefaults

`func NewApiWalletLoginRequestWithDefaults() *ApiWalletLoginRequest`

NewApiWalletLoginRequestWithDefaults instantiates a new ApiWalletLoginRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ApiWalletLoginRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiWalletLoginRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiWalletLoginRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiWalletLoginRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *ApiWalletLoginRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *ApiWalletLoginRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *ApiWalletLoginRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *ApiWalletLoginRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetSignedMessage

`func (o *ApiWalletLoginRequest) GetSignedMessage() string`

GetSignedMessage returns the SignedMessage field if non-nil, zero value otherwise.

### GetSignedMessageOk

`func (o *ApiWalletLoginRequest) GetSignedMessageOk() (*string, bool)`

GetSignedMessageOk returns a tuple with the SignedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedMessage

`func (o *ApiWalletLoginRequest) SetSignedMessage(v string)`

SetSignedMessage sets SignedMessage field to given value.

### HasSignedMessage

`func (o *ApiWalletLoginRequest) HasSignedMessage() bool`

HasSignedMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


