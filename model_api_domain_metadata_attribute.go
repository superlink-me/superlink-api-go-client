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

// checks if the ApiDomainMetadataAttribute type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiDomainMetadataAttribute{}

// ApiDomainMetadataAttribute struct for ApiDomainMetadataAttribute
type ApiDomainMetadataAttribute struct {
	TraitType *string `json:"trait_type,omitempty"`
	// This can be string or int
	Value map[string]interface{} `json:"value,omitempty"`
}

// NewApiDomainMetadataAttribute instantiates a new ApiDomainMetadataAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiDomainMetadataAttribute() *ApiDomainMetadataAttribute {
	this := ApiDomainMetadataAttribute{}
	return &this
}

// NewApiDomainMetadataAttributeWithDefaults instantiates a new ApiDomainMetadataAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiDomainMetadataAttributeWithDefaults() *ApiDomainMetadataAttribute {
	this := ApiDomainMetadataAttribute{}
	return &this
}

// GetTraitType returns the TraitType field value if set, zero value otherwise.
func (o *ApiDomainMetadataAttribute) GetTraitType() string {
	if o == nil || IsNil(o.TraitType) {
		var ret string
		return ret
	}
	return *o.TraitType
}

// GetTraitTypeOk returns a tuple with the TraitType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataAttribute) GetTraitTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TraitType) {
		return nil, false
	}
	return o.TraitType, true
}

// HasTraitType returns a boolean if a field has been set.
func (o *ApiDomainMetadataAttribute) HasTraitType() bool {
	if o != nil && !IsNil(o.TraitType) {
		return true
	}

	return false
}

// SetTraitType gets a reference to the given string and assigns it to the TraitType field.
func (o *ApiDomainMetadataAttribute) SetTraitType(v string) {
	o.TraitType = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiDomainMetadataAttribute) GetValue() map[string]interface{} {
	if o == nil || IsNil(o.Value) {
		var ret map[string]interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiDomainMetadataAttribute) GetValueOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return map[string]interface{}{}, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiDomainMetadataAttribute) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *ApiDomainMetadataAttribute) SetValue(v map[string]interface{}) {
	o.Value = v
}

func (o ApiDomainMetadataAttribute) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiDomainMetadataAttribute) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TraitType) {
		toSerialize["trait_type"] = o.TraitType
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiDomainMetadataAttribute struct {
	value *ApiDomainMetadataAttribute
	isSet bool
}

func (v NullableApiDomainMetadataAttribute) Get() *ApiDomainMetadataAttribute {
	return v.value
}

func (v *NullableApiDomainMetadataAttribute) Set(val *ApiDomainMetadataAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableApiDomainMetadataAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableApiDomainMetadataAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiDomainMetadataAttribute(val *ApiDomainMetadataAttribute) *NullableApiDomainMetadataAttribute {
	return &NullableApiDomainMetadataAttribute{value: val, isSet: true}
}

func (v NullableApiDomainMetadataAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiDomainMetadataAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


