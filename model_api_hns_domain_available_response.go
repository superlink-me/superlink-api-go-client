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

// checks if the ApiHNSDomainAvailableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHNSDomainAvailableResponse{}

// ApiHNSDomainAvailableResponse struct for ApiHNSDomainAvailableResponse
type ApiHNSDomainAvailableResponse struct {
	Domain *string `json:"domain,omitempty"`
	Status *ApiSubdomainStatus `json:"status,omitempty"`
}

// NewApiHNSDomainAvailableResponse instantiates a new ApiHNSDomainAvailableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHNSDomainAvailableResponse() *ApiHNSDomainAvailableResponse {
	this := ApiHNSDomainAvailableResponse{}
	return &this
}

// NewApiHNSDomainAvailableResponseWithDefaults instantiates a new ApiHNSDomainAvailableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHNSDomainAvailableResponseWithDefaults() *ApiHNSDomainAvailableResponse {
	this := ApiHNSDomainAvailableResponse{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiHNSDomainAvailableResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSDomainAvailableResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiHNSDomainAvailableResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiHNSDomainAvailableResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiHNSDomainAvailableResponse) GetStatus() ApiSubdomainStatus {
	if o == nil || IsNil(o.Status) {
		var ret ApiSubdomainStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSDomainAvailableResponse) GetStatusOk() (*ApiSubdomainStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiHNSDomainAvailableResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given ApiSubdomainStatus and assigns it to the Status field.
func (o *ApiHNSDomainAvailableResponse) SetStatus(v ApiSubdomainStatus) {
	o.Status = &v
}

func (o ApiHNSDomainAvailableResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHNSDomainAvailableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableApiHNSDomainAvailableResponse struct {
	value *ApiHNSDomainAvailableResponse
	isSet bool
}

func (v NullableApiHNSDomainAvailableResponse) Get() *ApiHNSDomainAvailableResponse {
	return v.value
}

func (v *NullableApiHNSDomainAvailableResponse) Set(val *ApiHNSDomainAvailableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHNSDomainAvailableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHNSDomainAvailableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHNSDomainAvailableResponse(val *ApiHNSDomainAvailableResponse) *NullableApiHNSDomainAvailableResponse {
	return &NullableApiHNSDomainAvailableResponse{value: val, isSet: true}
}

func (v NullableApiHNSDomainAvailableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHNSDomainAvailableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


