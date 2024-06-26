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

// checks if the ApiHNSRegisterTLDRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiHNSRegisterTLDRequest{}

// ApiHNSRegisterTLDRequest struct for ApiHNSRegisterTLDRequest
type ApiHNSRegisterTLDRequest struct {
	PartnerId *string `json:"partnerId,omitempty"`
	Tld *string `json:"tld,omitempty"`
}

// NewApiHNSRegisterTLDRequest instantiates a new ApiHNSRegisterTLDRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiHNSRegisterTLDRequest() *ApiHNSRegisterTLDRequest {
	this := ApiHNSRegisterTLDRequest{}
	return &this
}

// NewApiHNSRegisterTLDRequestWithDefaults instantiates a new ApiHNSRegisterTLDRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiHNSRegisterTLDRequestWithDefaults() *ApiHNSRegisterTLDRequest {
	this := ApiHNSRegisterTLDRequest{}
	return &this
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ApiHNSRegisterTLDRequest) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSRegisterTLDRequest) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ApiHNSRegisterTLDRequest) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ApiHNSRegisterTLDRequest) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetTld returns the Tld field value if set, zero value otherwise.
func (o *ApiHNSRegisterTLDRequest) GetTld() string {
	if o == nil || IsNil(o.Tld) {
		var ret string
		return ret
	}
	return *o.Tld
}

// GetTldOk returns a tuple with the Tld field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiHNSRegisterTLDRequest) GetTldOk() (*string, bool) {
	if o == nil || IsNil(o.Tld) {
		return nil, false
	}
	return o.Tld, true
}

// HasTld returns a boolean if a field has been set.
func (o *ApiHNSRegisterTLDRequest) HasTld() bool {
	if o != nil && !IsNil(o.Tld) {
		return true
	}

	return false
}

// SetTld gets a reference to the given string and assigns it to the Tld field.
func (o *ApiHNSRegisterTLDRequest) SetTld(v string) {
	o.Tld = &v
}

func (o ApiHNSRegisterTLDRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiHNSRegisterTLDRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PartnerId) {
		toSerialize["partnerId"] = o.PartnerId
	}
	if !IsNil(o.Tld) {
		toSerialize["tld"] = o.Tld
	}
	return toSerialize, nil
}

type NullableApiHNSRegisterTLDRequest struct {
	value *ApiHNSRegisterTLDRequest
	isSet bool
}

func (v NullableApiHNSRegisterTLDRequest) Get() *ApiHNSRegisterTLDRequest {
	return v.value
}

func (v *NullableApiHNSRegisterTLDRequest) Set(val *ApiHNSRegisterTLDRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiHNSRegisterTLDRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiHNSRegisterTLDRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiHNSRegisterTLDRequest(val *ApiHNSRegisterTLDRequest) *NullableApiHNSRegisterTLDRequest {
	return &NullableApiHNSRegisterTLDRequest{value: val, isSet: true}
}

func (v NullableApiHNSRegisterTLDRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiHNSRegisterTLDRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


