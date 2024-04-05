/*
Superlink

API for Superlink

API version: v0.4.1
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiTXTRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTXTRecord{}

// ApiTXTRecord struct for ApiTXTRecord
type ApiTXTRecord struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewApiTXTRecord instantiates a new ApiTXTRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTXTRecord() *ApiTXTRecord {
	this := ApiTXTRecord{}
	return &this
}

// NewApiTXTRecordWithDefaults instantiates a new ApiTXTRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTXTRecordWithDefaults() *ApiTXTRecord {
	this := ApiTXTRecord{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ApiTXTRecord) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTXTRecord) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ApiTXTRecord) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ApiTXTRecord) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiTXTRecord) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTXTRecord) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiTXTRecord) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiTXTRecord) SetValue(v string) {
	o.Value = &v
}

func (o ApiTXTRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTXTRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiTXTRecord struct {
	value *ApiTXTRecord
	isSet bool
}

func (v NullableApiTXTRecord) Get() *ApiTXTRecord {
	return v.value
}

func (v *NullableApiTXTRecord) Set(val *ApiTXTRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTXTRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTXTRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTXTRecord(val *ApiTXTRecord) *NullableApiTXTRecord {
	return &NullableApiTXTRecord{value: val, isSet: true}
}

func (v NullableApiTXTRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTXTRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


