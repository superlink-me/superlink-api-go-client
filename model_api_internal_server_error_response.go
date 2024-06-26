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

// checks if the ApiInternalServerErrorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiInternalServerErrorResponse{}

// ApiInternalServerErrorResponse struct for ApiInternalServerErrorResponse
type ApiInternalServerErrorResponse struct {
	Message *string `json:"message,omitempty"`
}

// NewApiInternalServerErrorResponse instantiates a new ApiInternalServerErrorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiInternalServerErrorResponse() *ApiInternalServerErrorResponse {
	this := ApiInternalServerErrorResponse{}
	return &this
}

// NewApiInternalServerErrorResponseWithDefaults instantiates a new ApiInternalServerErrorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiInternalServerErrorResponseWithDefaults() *ApiInternalServerErrorResponse {
	this := ApiInternalServerErrorResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiInternalServerErrorResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiInternalServerErrorResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiInternalServerErrorResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiInternalServerErrorResponse) SetMessage(v string) {
	o.Message = &v
}

func (o ApiInternalServerErrorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiInternalServerErrorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableApiInternalServerErrorResponse struct {
	value *ApiInternalServerErrorResponse
	isSet bool
}

func (v NullableApiInternalServerErrorResponse) Get() *ApiInternalServerErrorResponse {
	return v.value
}

func (v *NullableApiInternalServerErrorResponse) Set(val *ApiInternalServerErrorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiInternalServerErrorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiInternalServerErrorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiInternalServerErrorResponse(val *ApiInternalServerErrorResponse) *NullableApiInternalServerErrorResponse {
	return &NullableApiInternalServerErrorResponse{value: val, isSet: true}
}

func (v NullableApiInternalServerErrorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiInternalServerErrorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


