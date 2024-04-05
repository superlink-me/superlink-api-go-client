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

// checks if the ApiParentDomainPurchaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiParentDomainPurchaseRequest{}

// ApiParentDomainPurchaseRequest struct for ApiParentDomainPurchaseRequest
type ApiParentDomainPurchaseRequest struct {
	DomainContactDetails *ApiParentDomainPurchaseRequestDomainContactDetails `json:"domainContactDetails,omitempty"`
}

// NewApiParentDomainPurchaseRequest instantiates a new ApiParentDomainPurchaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiParentDomainPurchaseRequest() *ApiParentDomainPurchaseRequest {
	this := ApiParentDomainPurchaseRequest{}
	return &this
}

// NewApiParentDomainPurchaseRequestWithDefaults instantiates a new ApiParentDomainPurchaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiParentDomainPurchaseRequestWithDefaults() *ApiParentDomainPurchaseRequest {
	this := ApiParentDomainPurchaseRequest{}
	return &this
}

// GetDomainContactDetails returns the DomainContactDetails field value if set, zero value otherwise.
func (o *ApiParentDomainPurchaseRequest) GetDomainContactDetails() ApiParentDomainPurchaseRequestDomainContactDetails {
	if o == nil || IsNil(o.DomainContactDetails) {
		var ret ApiParentDomainPurchaseRequestDomainContactDetails
		return ret
	}
	return *o.DomainContactDetails
}

// GetDomainContactDetailsOk returns a tuple with the DomainContactDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiParentDomainPurchaseRequest) GetDomainContactDetailsOk() (*ApiParentDomainPurchaseRequestDomainContactDetails, bool) {
	if o == nil || IsNil(o.DomainContactDetails) {
		return nil, false
	}
	return o.DomainContactDetails, true
}

// HasDomainContactDetails returns a boolean if a field has been set.
func (o *ApiParentDomainPurchaseRequest) HasDomainContactDetails() bool {
	if o != nil && !IsNil(o.DomainContactDetails) {
		return true
	}

	return false
}

// SetDomainContactDetails gets a reference to the given ApiParentDomainPurchaseRequestDomainContactDetails and assigns it to the DomainContactDetails field.
func (o *ApiParentDomainPurchaseRequest) SetDomainContactDetails(v ApiParentDomainPurchaseRequestDomainContactDetails) {
	o.DomainContactDetails = &v
}

func (o ApiParentDomainPurchaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiParentDomainPurchaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DomainContactDetails) {
		toSerialize["domainContactDetails"] = o.DomainContactDetails
	}
	return toSerialize, nil
}

type NullableApiParentDomainPurchaseRequest struct {
	value *ApiParentDomainPurchaseRequest
	isSet bool
}

func (v NullableApiParentDomainPurchaseRequest) Get() *ApiParentDomainPurchaseRequest {
	return v.value
}

func (v *NullableApiParentDomainPurchaseRequest) Set(val *ApiParentDomainPurchaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiParentDomainPurchaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiParentDomainPurchaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiParentDomainPurchaseRequest(val *ApiParentDomainPurchaseRequest) *NullableApiParentDomainPurchaseRequest {
	return &NullableApiParentDomainPurchaseRequest{value: val, isSet: true}
}

func (v NullableApiParentDomainPurchaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiParentDomainPurchaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


