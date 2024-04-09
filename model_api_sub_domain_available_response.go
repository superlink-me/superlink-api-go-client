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

// checks if the ApiSubDomainAvailableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubDomainAvailableResponse{}

// ApiSubDomainAvailableResponse struct for ApiSubDomainAvailableResponse
type ApiSubDomainAvailableResponse struct {
	Domain *string `json:"domain,omitempty"`
	Status *ApiSubdomainStatus `json:"status,omitempty"`
}

// NewApiSubDomainAvailableResponse instantiates a new ApiSubDomainAvailableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubDomainAvailableResponse() *ApiSubDomainAvailableResponse {
	this := ApiSubDomainAvailableResponse{}
	return &this
}

// NewApiSubDomainAvailableResponseWithDefaults instantiates a new ApiSubDomainAvailableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubDomainAvailableResponseWithDefaults() *ApiSubDomainAvailableResponse {
	this := ApiSubDomainAvailableResponse{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiSubDomainAvailableResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainAvailableResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiSubDomainAvailableResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiSubDomainAvailableResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiSubDomainAvailableResponse) GetStatus() ApiSubdomainStatus {
	if o == nil || IsNil(o.Status) {
		var ret ApiSubdomainStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainAvailableResponse) GetStatusOk() (*ApiSubdomainStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiSubDomainAvailableResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiSubdomainStatus and assigns it to the Status field.
func (o *ApiSubDomainAvailableResponse) SetStatus(v ApiSubdomainStatus) {
	o.Status = &v
}

func (o ApiSubDomainAvailableResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubDomainAvailableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiSubDomainAvailableResponse struct {
	value *ApiSubDomainAvailableResponse
	isSet bool
}

func (v NullableApiSubDomainAvailableResponse) Get() *ApiSubDomainAvailableResponse {
	return v.value
}

func (v *NullableApiSubDomainAvailableResponse) Set(val *ApiSubDomainAvailableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubDomainAvailableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubDomainAvailableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubDomainAvailableResponse(val *ApiSubDomainAvailableResponse) *NullableApiSubDomainAvailableResponse {
	return &NullableApiSubDomainAvailableResponse{value: val, isSet: true}
}

func (v NullableApiSubDomainAvailableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubDomainAvailableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

