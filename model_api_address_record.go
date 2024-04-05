/*
Superlink

API for Superlink

API version: v0.4.2
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiAddressRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAddressRecord{}

// ApiAddressRecord struct for ApiAddressRecord
type ApiAddressRecord struct {
	Addr *string `json:"addr,omitempty"`
	CoinId *int32 `json:"coinId,omitempty"`
}

// NewApiAddressRecord instantiates a new ApiAddressRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAddressRecord() *ApiAddressRecord {
	this := ApiAddressRecord{}
	return &this
}

// NewApiAddressRecordWithDefaults instantiates a new ApiAddressRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAddressRecordWithDefaults() *ApiAddressRecord {
	this := ApiAddressRecord{}
	return &this
}

// GetAddr returns the Addr field value if set, zero value otherwise.
func (o *ApiAddressRecord) GetAddr() string {
	if o == nil || IsNil(o.Addr) {
		var ret string
		return ret
	}
	return *o.Addr
}

// GetAddrOk returns a tuple with the Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAddressRecord) GetAddrOk() (*string, bool) {
	if o == nil || IsNil(o.Addr) {
		return nil, false
	}
	return o.Addr, true
}

// HasAddr returns a boolean if a field has been set.
func (o *ApiAddressRecord) HasAddr() bool {
	if o != nil && !IsNil(o.Addr) {
		return true
	}

	return false
}

// SetAddr gets a reference to the given string and assigns it to the Addr field.
func (o *ApiAddressRecord) SetAddr(v string) {
	o.Addr = &v
}

// GetCoinId returns the CoinId field value if set, zero value otherwise.
func (o *ApiAddressRecord) GetCoinId() int32 {
	if o == nil || IsNil(o.CoinId) {
		var ret int32
		return ret
	}
	return *o.CoinId
}

// GetCoinIdOk returns a tuple with the CoinId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAddressRecord) GetCoinIdOk() (*int32, bool) {
	if o == nil || IsNil(o.CoinId) {
		return nil, false
	}
	return o.CoinId, true
}

// HasCoinId returns a boolean if a field has been set.
func (o *ApiAddressRecord) HasCoinId() bool {
	if o != nil && !IsNil(o.CoinId) {
		return true
	}

	return false
}

// SetCoinId gets a reference to the given int32 and assigns it to the CoinId field.
func (o *ApiAddressRecord) SetCoinId(v int32) {
	o.CoinId = &v
}

func (o ApiAddressRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAddressRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Addr) {
		toSerialize["addr"] = o.Addr
	}
	if !IsNil(o.CoinId) {
		toSerialize["coinId"] = o.CoinId
	}
	return toSerialize, nil
}

type NullableApiAddressRecord struct {
	value *ApiAddressRecord
	isSet bool
}

func (v NullableApiAddressRecord) Get() *ApiAddressRecord {
	return v.value
}

func (v *NullableApiAddressRecord) Set(val *ApiAddressRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAddressRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAddressRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAddressRecord(val *ApiAddressRecord) *NullableApiAddressRecord {
	return &NullableApiAddressRecord{value: val, isSet: true}
}

func (v NullableApiAddressRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAddressRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


