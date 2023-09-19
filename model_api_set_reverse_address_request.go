/*
Superlink

API for Superlink

API version: v0.3.17
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiSetReverseAddressRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSetReverseAddressRequest{}

// ApiSetReverseAddressRequest struct for ApiSetReverseAddressRequest
type ApiSetReverseAddressRequest struct {
	Address *string `json:"address,omitempty"`
	Domain *string `json:"domain,omitempty"`
	SignedMessage *string `json:"signedMessage,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewApiSetReverseAddressRequest instantiates a new ApiSetReverseAddressRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSetReverseAddressRequest() *ApiSetReverseAddressRequest {
	this := ApiSetReverseAddressRequest{}
	return &this
}

// NewApiSetReverseAddressRequestWithDefaults instantiates a new ApiSetReverseAddressRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSetReverseAddressRequestWithDefaults() *ApiSetReverseAddressRequest {
	this := ApiSetReverseAddressRequest{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApiSetReverseAddressRequest) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSetReverseAddressRequest) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApiSetReverseAddressRequest) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ApiSetReverseAddressRequest) SetAddress(v string) {
	o.Address = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiSetReverseAddressRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSetReverseAddressRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiSetReverseAddressRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiSetReverseAddressRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetSignedMessage returns the SignedMessage field value if set, zero value otherwise.
func (o *ApiSetReverseAddressRequest) GetSignedMessage() string {
	if o == nil || IsNil(o.SignedMessage) {
		var ret string
		return ret
	}
	return *o.SignedMessage
}

// GetSignedMessageOk returns a tuple with the SignedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSetReverseAddressRequest) GetSignedMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SignedMessage) {
		return nil, false
	}
	return o.SignedMessage, true
}

// HasSignedMessage returns a boolean if a field has been set.
func (o *ApiSetReverseAddressRequest) HasSignedMessage() bool {
	if o != nil && !IsNil(o.SignedMessage) {
		return true
	}

	return false
}

// SetSignedMessage gets a reference to the given string and assigns it to the SignedMessage field.
func (o *ApiSetReverseAddressRequest) SetSignedMessage(v string) {
	o.SignedMessage = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ApiSetReverseAddressRequest) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSetReverseAddressRequest) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ApiSetReverseAddressRequest) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *ApiSetReverseAddressRequest) SetUserId(v string) {
	o.UserId = &v
}

func (o ApiSetReverseAddressRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSetReverseAddressRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.SignedMessage) {
		toSerialize["signedMessage"] = o.SignedMessage
	}
	if !IsNil(o.UserId) {
		toSerialize["userId"] = o.UserId
	}
	return toSerialize, nil
}

type NullableApiSetReverseAddressRequest struct {
	value *ApiSetReverseAddressRequest
	isSet bool
}

func (v NullableApiSetReverseAddressRequest) Get() *ApiSetReverseAddressRequest {
	return v.value
}

func (v *NullableApiSetReverseAddressRequest) Set(val *ApiSetReverseAddressRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSetReverseAddressRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSetReverseAddressRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSetReverseAddressRequest(val *ApiSetReverseAddressRequest) *NullableApiSetReverseAddressRequest {
	return &NullableApiSetReverseAddressRequest{value: val, isSet: true}
}

func (v NullableApiSetReverseAddressRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSetReverseAddressRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


