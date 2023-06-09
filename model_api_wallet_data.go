/*
Superlink

API for Superlink

API version: v0.1.1
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
	Network *string `json:"network,omitempty"`
	Tag *string `json:"tag,omitempty"`
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

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ApiWalletData) GetNetwork() string {
	if o == nil || IsNil(o.Network) {
		var ret string
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletData) GetNetworkOk() (*string, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// HasNetwork returns a boolean if a field has been set.
func (o *ApiWalletData) HasNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given string and assigns it to the Network field.
func (o *ApiWalletData) SetNetwork(v string) {
	o.Network = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *ApiWalletData) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiWalletData) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *ApiWalletData) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *ApiWalletData) SetTag(v string) {
	o.Tag = &v
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
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
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

