/*
Superlink

API for Superlink

API version: v0.3.0
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiAccessTokenDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAccessTokenDeleteRequest{}

// ApiAccessTokenDeleteRequest struct for ApiAccessTokenDeleteRequest
type ApiAccessTokenDeleteRequest struct {
	Id *string `json:"id,omitempty"`
}

// NewApiAccessTokenDeleteRequest instantiates a new ApiAccessTokenDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccessTokenDeleteRequest() *ApiAccessTokenDeleteRequest {
	this := ApiAccessTokenDeleteRequest{}
	return &this
}

// NewApiAccessTokenDeleteRequestWithDefaults instantiates a new ApiAccessTokenDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccessTokenDeleteRequestWithDefaults() *ApiAccessTokenDeleteRequest {
	this := ApiAccessTokenDeleteRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAccessTokenDeleteRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenDeleteRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAccessTokenDeleteRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAccessTokenDeleteRequest) SetId(v string) {
	o.Id = &v
}

func (o ApiAccessTokenDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAccessTokenDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableApiAccessTokenDeleteRequest struct {
	value *ApiAccessTokenDeleteRequest
	isSet bool
}

func (v NullableApiAccessTokenDeleteRequest) Get() *ApiAccessTokenDeleteRequest {
	return v.value
}

func (v *NullableApiAccessTokenDeleteRequest) Set(val *ApiAccessTokenDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccessTokenDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccessTokenDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccessTokenDeleteRequest(val *ApiAccessTokenDeleteRequest) *NullableApiAccessTokenDeleteRequest {
	return &NullableApiAccessTokenDeleteRequest{value: val, isSet: true}
}

func (v NullableApiAccessTokenDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccessTokenDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


