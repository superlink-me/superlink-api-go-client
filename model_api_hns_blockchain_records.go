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

// checks if the ApiHNSBlockchainRecords type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHNSBlockchainRecords{}

// ApiHNSBlockchainRecords struct for ApiHNSBlockchainRecords
type ApiHNSBlockchainRecords struct {
	A *string `json:"a,omitempty"`
	Ds1 *string `json:"ds1,omitempty"`
	Ds2 *string `json:"ds2,omitempty"`
	Ds4 *string `json:"ds4,omitempty"`
}

// NewApiHNSBlockchainRecords instantiates a new ApiHNSBlockchainRecords object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHNSBlockchainRecords() *ApiHNSBlockchainRecords {
	this := ApiHNSBlockchainRecords{}
	return &this
}

// NewApiHNSBlockchainRecordsWithDefaults instantiates a new ApiHNSBlockchainRecords object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHNSBlockchainRecordsWithDefaults() *ApiHNSBlockchainRecords {
	this := ApiHNSBlockchainRecords{}
	return &this
}

// GetA returns the A field value if set, zero value otherwise.
func (o *ApiHNSBlockchainRecords) GetA() string {
	if o == nil || IsNil(o.A) {
		var ret string
		return ret
	}
	return *o.A
}

// GetAOk returns a tuple with the A field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSBlockchainRecords) GetAOk() (*string, bool) {
	if o == nil || IsNil(o.A) {
		return nil, false
	}
	return o.A, true
}

// HasA returns a boolean if a field has been set.
func (o *ApiHNSBlockchainRecords) HasA() bool {
	if o != nil && !IsNil(o.A) {
		return true
	}

	return false
}

// SetA gets a reference to the given string and assigns it to the A field.
func (o *ApiHNSBlockchainRecords) SetA(v string) {
	o.A = &v
}

// GetDs1 returns the Ds1 field value if set, zero value otherwise.
func (o *ApiHNSBlockchainRecords) GetDs1() string {
	if o == nil || IsNil(o.Ds1) {
		var ret string
		return ret
	}
	return *o.Ds1
}

// GetDs1Ok returns a tuple with the Ds1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSBlockchainRecords) GetDs1Ok() (*string, bool) {
	if o == nil || IsNil(o.Ds1) {
		return nil, false
	}
	return o.Ds1, true
}

// HasDs1 returns a boolean if a field has been set.
func (o *ApiHNSBlockchainRecords) HasDs1() bool {
	if o != nil && !IsNil(o.Ds1) {
		return true
	}

	return false
}

// SetDs1 gets a reference to the given string and assigns it to the Ds1 field.
func (o *ApiHNSBlockchainRecords) SetDs1(v string) {
	o.Ds1 = &v
}

// GetDs2 returns the Ds2 field value if set, zero value otherwise.
func (o *ApiHNSBlockchainRecords) GetDs2() string {
	if o == nil || IsNil(o.Ds2) {
		var ret string
		return ret
	}
	return *o.Ds2
}

// GetDs2Ok returns a tuple with the Ds2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSBlockchainRecords) GetDs2Ok() (*string, bool) {
	if o == nil || IsNil(o.Ds2) {
		return nil, false
	}
	return o.Ds2, true
}

// HasDs2 returns a boolean if a field has been set.
func (o *ApiHNSBlockchainRecords) HasDs2() bool {
	if o != nil && !IsNil(o.Ds2) {
		return true
	}

	return false
}

// SetDs2 gets a reference to the given string and assigns it to the Ds2 field.
func (o *ApiHNSBlockchainRecords) SetDs2(v string) {
	o.Ds2 = &v
}

// GetDs4 returns the Ds4 field value if set, zero value otherwise.
func (o *ApiHNSBlockchainRecords) GetDs4() string {
	if o == nil || IsNil(o.Ds4) {
		var ret string
		return ret
	}
	return *o.Ds4
}

// GetDs4Ok returns a tuple with the Ds4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSBlockchainRecords) GetDs4Ok() (*string, bool) {
	if o == nil || IsNil(o.Ds4) {
		return nil, false
	}
	return o.Ds4, true
}

// HasDs4 returns a boolean if a field has been set.
func (o *ApiHNSBlockchainRecords) HasDs4() bool {
	if o != nil && !IsNil(o.Ds4) {
		return true
	}

	return false
}

// SetDs4 gets a reference to the given string and assigns it to the Ds4 field.
func (o *ApiHNSBlockchainRecords) SetDs4(v string) {
	o.Ds4 = &v
}

func (o ApiHNSBlockchainRecords) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHNSBlockchainRecords) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.A) {
		toSerialize["a"] = o.A
	}
	if !IsNil(o.Ds1) {
		toSerialize["ds1"] = o.Ds1
	}
	if !IsNil(o.Ds2) {
		toSerialize["ds2"] = o.Ds2
	}
	if !IsNil(o.Ds4) {
		toSerialize["ds4"] = o.Ds4
	}
	return toSerialize, nil
}

type NullableApiHNSBlockchainRecords struct {
	value *ApiHNSBlockchainRecords
	isSet bool
}

func (v NullableApiHNSBlockchainRecords) Get() *ApiHNSBlockchainRecords {
	return v.value
}

func (v *NullableApiHNSBlockchainRecords) Set(val *ApiHNSBlockchainRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHNSBlockchainRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHNSBlockchainRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHNSBlockchainRecords(val *ApiHNSBlockchainRecords) *NullableApiHNSBlockchainRecords {
	return &NullableApiHNSBlockchainRecords{value: val, isSet: true}
}

func (v NullableApiHNSBlockchainRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHNSBlockchainRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


