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

// checks if the ApiAccessTokenCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAccessTokenCreateRequest{}

// ApiAccessTokenCreateRequest struct for ApiAccessTokenCreateRequest
type ApiAccessTokenCreateRequest struct {
	Label *string `json:"label,omitempty"`
	ValidFrom *string `json:"validFrom,omitempty"`
	ValidTill *string `json:"validTill,omitempty"`
}

// NewApiAccessTokenCreateRequest instantiates a new ApiAccessTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccessTokenCreateRequest() *ApiAccessTokenCreateRequest {
	this := ApiAccessTokenCreateRequest{}
	return &this
}

// NewApiAccessTokenCreateRequestWithDefaults instantiates a new ApiAccessTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccessTokenCreateRequestWithDefaults() *ApiAccessTokenCreateRequest {
	this := ApiAccessTokenCreateRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ApiAccessTokenCreateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenCreateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ApiAccessTokenCreateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ApiAccessTokenCreateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ApiAccessTokenCreateRequest) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenCreateRequest) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ApiAccessTokenCreateRequest) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ApiAccessTokenCreateRequest) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTill returns the ValidTill field value if set, zero value otherwise.
func (o *ApiAccessTokenCreateRequest) GetValidTill() string {
	if o == nil || IsNil(o.ValidTill) {
		var ret string
		return ret
	}
	return *o.ValidTill
}

// GetValidTillOk returns a tuple with the ValidTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenCreateRequest) GetValidTillOk() (*string, bool) {
	if o == nil || IsNil(o.ValidTill) {
		return nil, false
	}
	return o.ValidTill, true
}

// HasValidTill returns a boolean if a field has been set.
func (o *ApiAccessTokenCreateRequest) HasValidTill() bool {
	if o != nil && !IsNil(o.ValidTill) {
		return true
	}

	return false
}

// SetValidTill gets a reference to the given string and assigns it to the ValidTill field.
func (o *ApiAccessTokenCreateRequest) SetValidTill(v string) {
	o.ValidTill = &v
}

func (o ApiAccessTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAccessTokenCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if !IsNil(o.ValidTill) {
		toSerialize["validTill"] = o.ValidTill
	}
	return toSerialize, nil
}

type NullableApiAccessTokenCreateRequest struct {
	value *ApiAccessTokenCreateRequest
	isSet bool
}

func (v NullableApiAccessTokenCreateRequest) Get() *ApiAccessTokenCreateRequest {
	return v.value
}

func (v *NullableApiAccessTokenCreateRequest) Set(val *ApiAccessTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccessTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccessTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccessTokenCreateRequest(val *ApiAccessTokenCreateRequest) *NullableApiAccessTokenCreateRequest {
	return &NullableApiAccessTokenCreateRequest{value: val, isSet: true}
}

func (v NullableApiAccessTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccessTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


