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

// checks if the ApiHNSCreateDomainWithWalletRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHNSCreateDomainWithWalletRequest{}

// ApiHNSCreateDomainWithWalletRequest struct for ApiHNSCreateDomainWithWalletRequest
type ApiHNSCreateDomainWithWalletRequest struct {
	Addresses *map[string]string `json:"addresses,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Message *string `json:"message,omitempty"`
	OwnerAddress *string `json:"ownerAddress,omitempty"`
	SignedMessage *string `json:"signedMessage,omitempty"`
}

// NewApiHNSCreateDomainWithWalletRequest instantiates a new ApiHNSCreateDomainWithWalletRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHNSCreateDomainWithWalletRequest() *ApiHNSCreateDomainWithWalletRequest {
	this := ApiHNSCreateDomainWithWalletRequest{}
	return &this
}

// NewApiHNSCreateDomainWithWalletRequestWithDefaults instantiates a new ApiHNSCreateDomainWithWalletRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHNSCreateDomainWithWalletRequestWithDefaults() *ApiHNSCreateDomainWithWalletRequest {
	this := ApiHNSCreateDomainWithWalletRequest{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ApiHNSCreateDomainWithWalletRequest) GetAddresses() map[string]string {
	if o == nil || IsNil(o.Addresses) {
		var ret map[string]string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) GetAddressesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given map[string]string and assigns it to the Addresses field.
func (o *ApiHNSCreateDomainWithWalletRequest) SetAddresses(v map[string]string) {
	o.Addresses = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiHNSCreateDomainWithWalletRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiHNSCreateDomainWithWalletRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiHNSCreateDomainWithWalletRequest) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiHNSCreateDomainWithWalletRequest) SetMessage(v string) {
	o.Message = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ApiHNSCreateDomainWithWalletRequest) GetOwnerAddress() string {
	if o == nil || IsNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) GetOwnerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) HasOwnerAddress() bool {
	if o != nil && !IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ApiHNSCreateDomainWithWalletRequest) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetSignedMessage returns the SignedMessage field value if set, zero value otherwise.
func (o *ApiHNSCreateDomainWithWalletRequest) GetSignedMessage() string {
	if o == nil || IsNil(o.SignedMessage) {
		var ret string
		return ret
	}
	return *o.SignedMessage
}

// GetSignedMessageOk returns a tuple with the SignedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) GetSignedMessageOk() (*string, bool) {
	if o == nil || IsNil(o.SignedMessage) {
		return nil, false
	}
	return o.SignedMessage, true
}

// HasSignedMessage returns a boolean if a field has been set.
func (o *ApiHNSCreateDomainWithWalletRequest) HasSignedMessage() bool {
	if o != nil && !IsNil(o.SignedMessage) {
		return true
	}

	return false
}

// SetSignedMessage gets a reference to the given string and assigns it to the SignedMessage field.
func (o *ApiHNSCreateDomainWithWalletRequest) SetSignedMessage(v string) {
	o.SignedMessage = &v
}

func (o ApiHNSCreateDomainWithWalletRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHNSCreateDomainWithWalletRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.OwnerAddress) {
		toSerialize["ownerAddress"] = o.OwnerAddress
	}
	if !IsNil(o.SignedMessage) {
		toSerialize["signedMessage"] = o.SignedMessage
	}
	return toSerialize, nil
}

type NullableApiHNSCreateDomainWithWalletRequest struct {
	value *ApiHNSCreateDomainWithWalletRequest
	isSet bool
}

func (v NullableApiHNSCreateDomainWithWalletRequest) Get() *ApiHNSCreateDomainWithWalletRequest {
	return v.value
}

func (v *NullableApiHNSCreateDomainWithWalletRequest) Set(val *ApiHNSCreateDomainWithWalletRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHNSCreateDomainWithWalletRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHNSCreateDomainWithWalletRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHNSCreateDomainWithWalletRequest(val *ApiHNSCreateDomainWithWalletRequest) *NullableApiHNSCreateDomainWithWalletRequest {
	return &NullableApiHNSCreateDomainWithWalletRequest{value: val, isSet: true}
}

func (v NullableApiHNSCreateDomainWithWalletRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHNSCreateDomainWithWalletRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

