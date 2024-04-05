/*
Superlink

API for Superlink

API version: v0.4.2
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiSubDomainItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubDomainItem{}

// ApiSubDomainItem struct for ApiSubDomainItem
type ApiSubDomainItem struct {
	CreatedTimestamp *string `json:"createdTimestamp,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiSubDomainItem instantiates a new ApiSubDomainItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubDomainItem() *ApiSubDomainItem {
	this := ApiSubDomainItem{}
	return &this
}

// NewApiSubDomainItemWithDefaults instantiates a new ApiSubDomainItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubDomainItemWithDefaults() *ApiSubDomainItem {
	this := ApiSubDomainItem{}
	return &this
}

// GetCreatedTimestamp returns the CreatedTimestamp field value if set, zero value otherwise.
func (o *ApiSubDomainItem) GetCreatedTimestamp() string {
	if o == nil || IsNil(o.CreatedTimestamp) {
		var ret string
		return ret
	}
	return *o.CreatedTimestamp
}

// GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainItem) GetCreatedTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedTimestamp) {
		return nil, false
	}
	return o.CreatedTimestamp, true
}

// HasCreatedTimestamp returns a boolean if a field has been set.
func (o *ApiSubDomainItem) HasCreatedTimestamp() bool {
	if o != nil && !IsNil(o.CreatedTimestamp) {
		return true
	}

	return false
}

// SetCreatedTimestamp gets a reference to the given string and assigns it to the CreatedTimestamp field.
func (o *ApiSubDomainItem) SetCreatedTimestamp(v string) {
	o.CreatedTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiSubDomainItem) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainItem) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiSubDomainItem) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiSubDomainItem) SetName(v string) {
	o.Name = &v
}

func (o ApiSubDomainItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubDomainItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedTimestamp) {
		toSerialize["createdTimestamp"] = o.CreatedTimestamp
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableApiSubDomainItem struct {
	value *ApiSubDomainItem
	isSet bool
}

func (v NullableApiSubDomainItem) Get() *ApiSubDomainItem {
	return v.value
}

func (v *NullableApiSubDomainItem) Set(val *ApiSubDomainItem) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubDomainItem) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubDomainItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubDomainItem(val *ApiSubDomainItem) *NullableApiSubDomainItem {
	return &NullableApiSubDomainItem{value: val, isSet: true}
}

func (v NullableApiSubDomainItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubDomainItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


