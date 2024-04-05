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

// checks if the ApiSubDomainTotalResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubDomainTotalResponse{}

// ApiSubDomainTotalResponse struct for ApiSubDomainTotalResponse
type ApiSubDomainTotalResponse struct {
	Count *int32 `json:"count,omitempty"`
}

// NewApiSubDomainTotalResponse instantiates a new ApiSubDomainTotalResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubDomainTotalResponse() *ApiSubDomainTotalResponse {
	this := ApiSubDomainTotalResponse{}
	return &this
}

// NewApiSubDomainTotalResponseWithDefaults instantiates a new ApiSubDomainTotalResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubDomainTotalResponseWithDefaults() *ApiSubDomainTotalResponse {
	this := ApiSubDomainTotalResponse{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ApiSubDomainTotalResponse) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainTotalResponse) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ApiSubDomainTotalResponse) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *ApiSubDomainTotalResponse) SetCount(v int32) {
	o.Count = &v
}

func (o ApiSubDomainTotalResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubDomainTotalResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	return toSerialize, nil
}

type NullableApiSubDomainTotalResponse struct {
	value *ApiSubDomainTotalResponse
	isSet bool
}

func (v NullableApiSubDomainTotalResponse) Get() *ApiSubDomainTotalResponse {
	return v.value
}

func (v *NullableApiSubDomainTotalResponse) Set(val *ApiSubDomainTotalResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubDomainTotalResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubDomainTotalResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubDomainTotalResponse(val *ApiSubDomainTotalResponse) *NullableApiSubDomainTotalResponse {
	return &NullableApiSubDomainTotalResponse{value: val, isSet: true}
}

func (v NullableApiSubDomainTotalResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubDomainTotalResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


