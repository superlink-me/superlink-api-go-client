/*
Superlink

API for Superlink

API version: v0.5.0
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiWalletLoginRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiWalletLoginRequest{}

// ApiWalletLoginRequest struct for ApiWalletLoginRequest
type ApiWalletLoginRequest struct {
	Message *string `json:"message,omitempty"`
	RedirectUrl *string `json:"redirectUrl,omitempty"`
	SignedMessage *string `json:"signedMessage,omitempty"`
}

// NewApiWalletLoginRequest instantiates a new ApiWalletLoginRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWalletLoginRequest() *ApiWalletLoginRequest {
	this := ApiWalletLoginRequest{}
	return &this
}

// NewApiWalletLoginRequestWithDefaults instantiates a new ApiWalletLoginRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWalletLoginRequestWithDefaults() *ApiWalletLoginRequest {
	this := ApiWalletLoginRequest{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiWalletLoginRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletLoginRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiWalletLoginRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiWalletLoginRequest) SetMessage(v string) {
	o.Message = &v
}

// GetRedirectUrl returns the RedirectUrl field value if set, zero value otherwise.
func (o *ApiWalletLoginRequest) GetRedirectUrl() string {
	if o == nil || IsNil(o.RedirectUrl) {
		var ret string
		return ret
	}
	return *o.RedirectUrl
}

// GetRedirectUrlOk returns a tuple with the RedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletLoginRequest) GetRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectUrl) {
		return nil, false
	}
	return o.RedirectUrl, true
}

// HasRedirectUrl returns a boolean if a field has been set.
func (o *ApiWalletLoginRequest) HasRedirectUrl() bool {
	if o != nil && !IsNil(o.RedirectUrl) {
		return true
	}

	return false
}

// SetRedirectUrl gets a reference to the given string and assigns it to the RedirectUrl field.
func (o *ApiWalletLoginRequest) SetRedirectUrl(v string) {
	o.RedirectUrl = &v
}

// GetSignedMessage returns the SignedMessage field value if set, zero value otherwise.
func (o *ApiWalletLoginRequest) GetSignedMessage() string {
	if o == nil || IsNil(o.SignedMessage) {
		var ret string
		return ret
	}
	return *o.SignedMessage
}

// GetSignedMessageOk returns a tuple with the SignedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletLoginRequest) GetSignedMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SignedMessage) {
		return nil, false
	}
	return o.SignedMessage, true
}

// HasSignedMessage returns a boolean if a field has been set.
func (o *ApiWalletLoginRequest) HasSignedMessage() bool {
	if o != nil && !IsNil(o.SignedMessage) {
		return true
	}

	return false
}

// SetSignedMessage gets a reference to the given string and assigns it to the SignedMessage field.
func (o *ApiWalletLoginRequest) SetSignedMessage(v string) {
	o.SignedMessage = &v
}

func (o ApiWalletLoginRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiWalletLoginRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.RedirectUrl) {
		toSerialize["redirectUrl"] = o.RedirectUrl
	}
	if !IsNil(o.SignedMessage) {
		toSerialize["signedMessage"] = o.SignedMessage
	}
	return toSerialize, nil
}

type NullableApiWalletLoginRequest struct {
	value *ApiWalletLoginRequest
	isSet bool
}

func (v NullableApiWalletLoginRequest) Get() *ApiWalletLoginRequest {
	return v.value
}

func (v *NullableApiWalletLoginRequest) Set(val *ApiWalletLoginRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWalletLoginRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWalletLoginRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWalletLoginRequest(val *ApiWalletLoginRequest) *NullableApiWalletLoginRequest {
	return &NullableApiWalletLoginRequest{value: val, isSet: true}
}

func (v NullableApiWalletLoginRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWalletLoginRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


