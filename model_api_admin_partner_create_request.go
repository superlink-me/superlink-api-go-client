/*
Superlink

API for Superlink

API version: v0.3.27
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiAdminPartnerCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAdminPartnerCreateRequest{}

// ApiAdminPartnerCreateRequest struct for ApiAdminPartnerCreateRequest
type ApiAdminPartnerCreateRequest struct {
	Name *string `json:"name,omitempty"`
}

// NewApiAdminPartnerCreateRequest instantiates a new ApiAdminPartnerCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAdminPartnerCreateRequest() *ApiAdminPartnerCreateRequest {
	this := ApiAdminPartnerCreateRequest{}
	return &this
}

// NewApiAdminPartnerCreateRequestWithDefaults instantiates a new ApiAdminPartnerCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAdminPartnerCreateRequestWithDefaults() *ApiAdminPartnerCreateRequest {
	this := ApiAdminPartnerCreateRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiAdminPartnerCreateRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminPartnerCreateRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiAdminPartnerCreateRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiAdminPartnerCreateRequest) SetName(v string) {
	o.Name = &v
}

func (o ApiAdminPartnerCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAdminPartnerCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiAdminPartnerCreateRequest struct {
	value *ApiAdminPartnerCreateRequest
	isSet bool
}

func (v NullableApiAdminPartnerCreateRequest) Get() *ApiAdminPartnerCreateRequest {
	return v.value
}

func (v *NullableApiAdminPartnerCreateRequest) Set(val *ApiAdminPartnerCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAdminPartnerCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAdminPartnerCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAdminPartnerCreateRequest(val *ApiAdminPartnerCreateRequest) *NullableApiAdminPartnerCreateRequest {
	return &NullableApiAdminPartnerCreateRequest{value: val, isSet: true}
}

func (v NullableApiAdminPartnerCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAdminPartnerCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


