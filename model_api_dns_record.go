/*
Superlink

API for Superlink

API version: v0.0.4-alpha.8
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiDNSRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDNSRecord{}

// ApiDNSRecord struct for ApiDNSRecord
type ApiDNSRecord struct {
	Name *string `json:"name,omitempty"`
	Ttl *int32 `json:"ttl,omitempty"`
	Type *string `json:"type,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewApiDNSRecord instantiates a new ApiDNSRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDNSRecord() *ApiDNSRecord {
	this := ApiDNSRecord{}
	return &this
}

// NewApiDNSRecordWithDefaults instantiates a new ApiDNSRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDNSRecordWithDefaults() *ApiDNSRecord {
	this := ApiDNSRecord{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiDNSRecord) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDNSRecord) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiDNSRecord) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiDNSRecord) SetName(v string) {
	o.Name = &v
}

// GetTtl returns the Ttl field value if set, zero value otherwise.
func (o *ApiDNSRecord) GetTtl() int32 {
	if o == nil || IsNil(o.Ttl) {
		var ret int32
		return ret
	}
	return *o.Ttl
}

// GetTtlOk returns a tuple with the Ttl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDNSRecord) GetTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.Ttl) {
		return nil, false
	}
	return o.Ttl, true
}

// HasTtl returns a boolean if a field has been set.
func (o *ApiDNSRecord) HasTtl() bool {
	if o != nil && !IsNil(o.Ttl) {
		return true
	}

	return false
}

// SetTtl gets a reference to the given int32 and assigns it to the Ttl field.
func (o *ApiDNSRecord) SetTtl(v int32) {
	o.Ttl = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiDNSRecord) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDNSRecord) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiDNSRecord) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiDNSRecord) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiDNSRecord) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDNSRecord) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiDNSRecord) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiDNSRecord) SetValue(v string) {
	o.Value = &v
}

func (o ApiDNSRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDNSRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Ttl) {
		toSerialize["ttl"] = o.Ttl
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiDNSRecord struct {
	value *ApiDNSRecord
	isSet bool
}

func (v NullableApiDNSRecord) Get() *ApiDNSRecord {
	return v.value
}

func (v *NullableApiDNSRecord) Set(val *ApiDNSRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDNSRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDNSRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDNSRecord(val *ApiDNSRecord) *NullableApiDNSRecord {
	return &NullableApiDNSRecord{value: val, isSet: true}
}

func (v NullableApiDNSRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDNSRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


