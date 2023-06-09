/*
Superlink

API for Superlink

API version: v0.1.2
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiAccessTokenResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAccessTokenResponse{}

// ApiAccessTokenResponse struct for ApiAccessTokenResponse
type ApiAccessTokenResponse struct {
	CreatedAt *string `json:"createdAt,omitempty"`
	Id *string `json:"id,omitempty"`
	Label *string `json:"label,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
	ValidFrom *string `json:"validFrom,omitempty"`
	ValidTill *string `json:"validTill,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewApiAccessTokenResponse instantiates a new ApiAccessTokenResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccessTokenResponse() *ApiAccessTokenResponse {
	this := ApiAccessTokenResponse{}
	return &this
}

// NewApiAccessTokenResponseWithDefaults instantiates a new ApiAccessTokenResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccessTokenResponseWithDefaults() *ApiAccessTokenResponse {
	this := ApiAccessTokenResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApiAccessTokenResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiAccessTokenResponse) SetId(v string) {
	o.Id = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ApiAccessTokenResponse) SetLabel(v string) {
	o.Label = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ApiAccessTokenResponse) SetType(v string) {
	o.Type = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ApiAccessTokenResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ApiAccessTokenResponse) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTill returns the ValidTill field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetValidTill() string {
	if o == nil || IsNil(o.ValidTill) {
		var ret string
		return ret
	}
	return *o.ValidTill
}

// GetValidTillOk returns a tuple with the ValidTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetValidTillOk() (*string, bool) {
	if o == nil || IsNil(o.ValidTill) {
		return nil, false
	}
	return o.ValidTill, true
}

// HasValidTill returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasValidTill() bool {
	if o != nil && !IsNil(o.ValidTill) {
		return true
	}

	return false
}

// SetValidTill gets a reference to the given string and assigns it to the ValidTill field.
func (o *ApiAccessTokenResponse) SetValidTill(v string) {
	o.ValidTill = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ApiAccessTokenResponse) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccessTokenResponse) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ApiAccessTokenResponse) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ApiAccessTokenResponse) SetValue(v string) {
	o.Value = &v
}

func (o ApiAccessTokenResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAccessTokenResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if !IsNil(o.ValidTill) {
		toSerialize["validTill"] = o.ValidTill
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableApiAccessTokenResponse struct {
	value *ApiAccessTokenResponse
	isSet bool
}

func (v NullableApiAccessTokenResponse) Get() *ApiAccessTokenResponse {
	return v.value
}

func (v *NullableApiAccessTokenResponse) Set(val *ApiAccessTokenResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccessTokenResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccessTokenResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccessTokenResponse(val *ApiAccessTokenResponse) *NullableApiAccessTokenResponse {
	return &NullableApiAccessTokenResponse{value: val, isSet: true}
}

func (v NullableApiAccessTokenResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccessTokenResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


