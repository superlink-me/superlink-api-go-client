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

// checks if the ApiSubDomainNameData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubDomainNameData{}

// ApiSubDomainNameData struct for ApiSubDomainNameData
type ApiSubDomainNameData struct {
	Addresses *map[string]string `json:"addresses,omitempty"`
	Contenthash *string `json:"contenthash,omitempty"`
	Text *map[string]string `json:"text,omitempty"`
}

// NewApiSubDomainNameData instantiates a new ApiSubDomainNameData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubDomainNameData() *ApiSubDomainNameData {
	this := ApiSubDomainNameData{}
	return &this
}

// NewApiSubDomainNameDataWithDefaults instantiates a new ApiSubDomainNameData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubDomainNameDataWithDefaults() *ApiSubDomainNameData {
	this := ApiSubDomainNameData{}
	return &this
}

// GetAddresses returns the Addresses field value if set, zero value otherwise.
func (o *ApiSubDomainNameData) GetAddresses() map[string]string {
	if o == nil || IsNil(o.Addresses) {
		var ret map[string]string
		return ret
	}
	return *o.Addresses
}

// GetAddressesOk returns a tuple with the Addresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainNameData) GetAddressesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Addresses) {
		return nil, false
	}
	return o.Addresses, true
}

// HasAddresses returns a boolean if a field has been set.
func (o *ApiSubDomainNameData) HasAddresses() bool {
	if o != nil && !IsNil(o.Addresses) {
		return true
	}

	return false
}

// SetAddresses gets a reference to the given map[string]string and assigns it to the Addresses field.
func (o *ApiSubDomainNameData) SetAddresses(v map[string]string) {
	o.Addresses = &v
}

// GetContenthash returns the Contenthash field value if set, zero value otherwise.
func (o *ApiSubDomainNameData) GetContenthash() string {
	if o == nil || IsNil(o.Contenthash) {
		var ret string
		return ret
	}
	return *o.Contenthash
}

// GetContenthashOk returns a tuple with the Contenthash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainNameData) GetContenthashOk() (*string, bool) {
	if o == nil || IsNil(o.Contenthash) {
		return nil, false
	}
	return o.Contenthash, true
}

// HasContenthash returns a boolean if a field has been set.
func (o *ApiSubDomainNameData) HasContenthash() bool {
	if o != nil && !IsNil(o.Contenthash) {
		return true
	}

	return false
}

// SetContenthash gets a reference to the given string and assigns it to the Contenthash field.
func (o *ApiSubDomainNameData) SetContenthash(v string) {
	o.Contenthash = &v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *ApiSubDomainNameData) GetText() map[string]string {
	if o == nil || IsNil(o.Text) {
		var ret map[string]string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainNameData) GetTextOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *ApiSubDomainNameData) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given map[string]string and assigns it to the Text field.
func (o *ApiSubDomainNameData) SetText(v map[string]string) {
	o.Text = &v
}

func (o ApiSubDomainNameData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubDomainNameData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addresses) {
		toSerialize["addresses"] = o.Addresses
	}
	if !IsNil(o.Contenthash) {
		toSerialize["contenthash"] = o.Contenthash
	}
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableApiSubDomainNameData struct {
	value *ApiSubDomainNameData
	isSet bool
}

func (v NullableApiSubDomainNameData) Get() *ApiSubDomainNameData {
	return v.value
}

func (v *NullableApiSubDomainNameData) Set(val *ApiSubDomainNameData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubDomainNameData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubDomainNameData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubDomainNameData(val *ApiSubDomainNameData) *NullableApiSubDomainNameData {
	return &NullableApiSubDomainNameData{value: val, isSet: true}
}

func (v NullableApiSubDomainNameData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubDomainNameData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


