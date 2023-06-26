/*
Superlink

API for Superlink

API version: v0.1.6-alpha.2
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiWalletData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiWalletData{}

// ApiWalletData struct for ApiWalletData
type ApiWalletData struct {
	Address *string `json:"address,omitempty"`
	Coin *ApiCoin `json:"coin,omitempty"`
	Version *string `json:"version,omitempty"`
}

// NewApiWalletData instantiates a new ApiWalletData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiWalletData() *ApiWalletData {
	this := ApiWalletData{}
	return &this
}

// NewApiWalletDataWithDefaults instantiates a new ApiWalletData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiWalletDataWithDefaults() *ApiWalletData {
	this := ApiWalletData{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApiWalletData) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletData) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApiWalletData) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ApiWalletData) SetAddress(v string) {
	o.Address = &v
}

// GetCoin returns the Coin field value if set, zero value otherwise.
func (o *ApiWalletData) GetCoin() ApiCoin {
	if o == nil || IsNil(o.Coin) {
		var ret ApiCoin
		return ret
	}
	return *o.Coin
}

// GetCoinOk returns a tuple with the Coin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletData) GetCoinOk() (*ApiCoin, bool) {
	if o == nil || IsNil(o.Coin) {
		return nil, false
	}
	return o.Coin, true
}

// HasCoin returns a boolean if a field has been set.
func (o *ApiWalletData) HasCoin() bool {
	if o != nil && !IsNil(o.Coin) {
		return true
	}

	return false
}

// SetCoin gets a reference to the given ApiCoin and assigns it to the Coin field.
func (o *ApiWalletData) SetCoin(v ApiCoin) {
	o.Coin = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApiWalletData) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletData) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApiWalletData) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApiWalletData) SetVersion(v string) {
	o.Version = &v
}

func (o ApiWalletData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiWalletData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Coin) {
		toSerialize["coin"] = o.Coin
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableApiWalletData struct {
	value *ApiWalletData
	isSet bool
}

func (v NullableApiWalletData) Get() *ApiWalletData {
	return v.value
}

func (v *NullableApiWalletData) Set(val *ApiWalletData) {
	v.value = val
	v.isSet = true
}

func (v NullableApiWalletData) IsSet() bool {
	return v.isSet
}

func (v *NullableApiWalletData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiWalletData(val *ApiWalletData) *NullableApiWalletData {
	return &NullableApiWalletData{value: val, isSet: true}
}

func (v NullableApiWalletData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiWalletData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


