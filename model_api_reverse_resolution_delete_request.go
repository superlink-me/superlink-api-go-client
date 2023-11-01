/*
Superlink

API for Superlink

API version: v0.3.33
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiReverseResolutionDeleteRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiReverseResolutionDeleteRequest{}

// ApiReverseResolutionDeleteRequest struct for ApiReverseResolutionDeleteRequest
type ApiReverseResolutionDeleteRequest struct {
	Domain *string `json:"domain,omitempty"`
}

// NewApiReverseResolutionDeleteRequest instantiates a new ApiReverseResolutionDeleteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiReverseResolutionDeleteRequest() *ApiReverseResolutionDeleteRequest {
	this := ApiReverseResolutionDeleteRequest{}
	return &this
}

// NewApiReverseResolutionDeleteRequestWithDefaults instantiates a new ApiReverseResolutionDeleteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiReverseResolutionDeleteRequestWithDefaults() *ApiReverseResolutionDeleteRequest {
	this := ApiReverseResolutionDeleteRequest{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiReverseResolutionDeleteRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiReverseResolutionDeleteRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiReverseResolutionDeleteRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiReverseResolutionDeleteRequest) SetDomain(v string) {
	o.Domain = &v
}

func (o ApiReverseResolutionDeleteRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiReverseResolutionDeleteRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	return toSerialize, nil
}

type NullableApiReverseResolutionDeleteRequest struct {
	value *ApiReverseResolutionDeleteRequest
	isSet bool
}

func (v NullableApiReverseResolutionDeleteRequest) Get() *ApiReverseResolutionDeleteRequest {
	return v.value
}

func (v *NullableApiReverseResolutionDeleteRequest) Set(val *ApiReverseResolutionDeleteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiReverseResolutionDeleteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiReverseResolutionDeleteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiReverseResolutionDeleteRequest(val *ApiReverseResolutionDeleteRequest) *NullableApiReverseResolutionDeleteRequest {
	return &NullableApiReverseResolutionDeleteRequest{value: val, isSet: true}
}

func (v NullableApiReverseResolutionDeleteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiReverseResolutionDeleteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


