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

// checks if the ApiBadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiBadRequestResponse{}

// ApiBadRequestResponse struct for ApiBadRequestResponse
type ApiBadRequestResponse struct {
	Errors []ApiErrorResponse `json:"errors,omitempty"`
}

// NewApiBadRequestResponse instantiates a new ApiBadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiBadRequestResponse() *ApiBadRequestResponse {
	this := ApiBadRequestResponse{}
	return &this
}

// NewApiBadRequestResponseWithDefaults instantiates a new ApiBadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiBadRequestResponseWithDefaults() *ApiBadRequestResponse {
	this := ApiBadRequestResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ApiBadRequestResponse) GetErrors() []ApiErrorResponse {
	if o == nil || IsNil(o.Errors) {
		var ret []ApiErrorResponse
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiBadRequestResponse) GetErrorsOk() ([]ApiErrorResponse, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ApiBadRequestResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ApiErrorResponse and assigns it to the Errors field.
func (o *ApiBadRequestResponse) SetErrors(v []ApiErrorResponse) {
	o.Errors = v
}

func (o ApiBadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiBadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableApiBadRequestResponse struct {
	value *ApiBadRequestResponse
	isSet bool
}

func (v NullableApiBadRequestResponse) Get() *ApiBadRequestResponse {
	return v.value
}

func (v *NullableApiBadRequestResponse) Set(val *ApiBadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiBadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiBadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiBadRequestResponse(val *ApiBadRequestResponse) *NullableApiBadRequestResponse {
	return &NullableApiBadRequestResponse{value: val, isSet: true}
}

func (v NullableApiBadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiBadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


