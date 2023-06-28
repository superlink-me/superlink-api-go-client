/*
Superlink

API for Superlink

API version: v0.1.10
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiErrorResponse{}

// ApiErrorResponse struct for ApiErrorResponse
type ApiErrorResponse struct {
	Field *string `json:"field,omitempty"`
	Info *string `json:"info,omitempty"`
}

// NewApiErrorResponse instantiates a new ApiErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiErrorResponse() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// NewApiErrorResponseWithDefaults instantiates a new ApiErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiErrorResponseWithDefaults() *ApiErrorResponse {
	this := ApiErrorResponse{}
	return &this
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *ApiErrorResponse) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *ApiErrorResponse) SetField(v string) {
	o.Field = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *ApiErrorResponse) GetInfo() string {
	if o == nil || IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiErrorResponse) GetInfoOk() (*string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *ApiErrorResponse) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *ApiErrorResponse) SetInfo(v string) {
	o.Info = &v
}

func (o ApiErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	return toSerialize, nil
}

type NullableApiErrorResponse struct {
	value *ApiErrorResponse
	isSet bool
}

func (v NullableApiErrorResponse) Get() *ApiErrorResponse {
	return v.value
}

func (v *NullableApiErrorResponse) Set(val *ApiErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiErrorResponse(val *ApiErrorResponse) *NullableApiErrorResponse {
	return &NullableApiErrorResponse{value: val, isSet: true}
}

func (v NullableApiErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


