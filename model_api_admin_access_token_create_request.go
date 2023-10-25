/*
Superlink

API for Superlink

API version: v0.3.28
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiAdminAccessTokenCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAdminAccessTokenCreateRequest{}

// ApiAdminAccessTokenCreateRequest struct for ApiAdminAccessTokenCreateRequest
type ApiAdminAccessTokenCreateRequest struct {
	Label *string `json:"label,omitempty"`
	PartnerId *string `json:"partnerId,omitempty"`
	Type *DataAccessTokenType `json:"type,omitempty"`
	ValidFrom *string `json:"validFrom,omitempty"`
	ValidTill *string `json:"validTill,omitempty"`
}

// NewApiAdminAccessTokenCreateRequest instantiates a new ApiAdminAccessTokenCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAdminAccessTokenCreateRequest() *ApiAdminAccessTokenCreateRequest {
	this := ApiAdminAccessTokenCreateRequest{}
	return &this
}

// NewApiAdminAccessTokenCreateRequestWithDefaults instantiates a new ApiAdminAccessTokenCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAdminAccessTokenCreateRequestWithDefaults() *ApiAdminAccessTokenCreateRequest {
	this := ApiAdminAccessTokenCreateRequest{}
	return &this
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ApiAdminAccessTokenCreateRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminAccessTokenCreateRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ApiAdminAccessTokenCreateRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ApiAdminAccessTokenCreateRequest) SetLabel(v string) {
	o.Label = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ApiAdminAccessTokenCreateRequest) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminAccessTokenCreateRequest) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ApiAdminAccessTokenCreateRequest) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ApiAdminAccessTokenCreateRequest) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ApiAdminAccessTokenCreateRequest) GetType() DataAccessTokenType {
	if o == nil || IsNil(o.Type) {
		var ret DataAccessTokenType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminAccessTokenCreateRequest) GetTypeOk() (*DataAccessTokenType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ApiAdminAccessTokenCreateRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DataAccessTokenType and assigns it to the Type field.
func (o *ApiAdminAccessTokenCreateRequest) SetType(v DataAccessTokenType) {
	o.Type = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *ApiAdminAccessTokenCreateRequest) GetValidFrom() string {
	if o == nil || IsNil(o.ValidFrom) {
		var ret string
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminAccessTokenCreateRequest) GetValidFromOk() (*string, bool) {
	if o == nil || IsNil(o.ValidFrom) {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *ApiAdminAccessTokenCreateRequest) HasValidFrom() bool {
	if o != nil && !IsNil(o.ValidFrom) {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given string and assigns it to the ValidFrom field.
func (o *ApiAdminAccessTokenCreateRequest) SetValidFrom(v string) {
	o.ValidFrom = &v
}

// GetValidTill returns the ValidTill field value if set, zero value otherwise.
func (o *ApiAdminAccessTokenCreateRequest) GetValidTill() string {
	if o == nil || IsNil(o.ValidTill) {
		var ret string
		return ret
	}
	return *o.ValidTill
}

// GetValidTillOk returns a tuple with the ValidTill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAdminAccessTokenCreateRequest) GetValidTillOk() (*string, bool) {
	if o == nil || IsNil(o.ValidTill) {
		return nil, false
	}
	return o.ValidTill, true
}

// HasValidTill returns a boolean if a field has been set.
func (o *ApiAdminAccessTokenCreateRequest) HasValidTill() bool {
	if o != nil && !IsNil(o.ValidTill) {
		return true
	}

	return false
}

// SetValidTill gets a reference to the given string and assigns it to the ValidTill field.
func (o *ApiAdminAccessTokenCreateRequest) SetValidTill(v string) {
	o.ValidTill = &v
}

func (o ApiAdminAccessTokenCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAdminAccessTokenCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partnerId"] = o.PartnerId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.ValidFrom) {
		toSerialize["validFrom"] = o.ValidFrom
	}
	if !IsNil(o.ValidTill) {
		toSerialize["validTill"] = o.ValidTill
	}
	return toSerialize, nil
}

type NullableApiAdminAccessTokenCreateRequest struct {
	value *ApiAdminAccessTokenCreateRequest
	isSet bool
}

func (v NullableApiAdminAccessTokenCreateRequest) Get() *ApiAdminAccessTokenCreateRequest {
	return v.value
}

func (v *NullableApiAdminAccessTokenCreateRequest) Set(val *ApiAdminAccessTokenCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAdminAccessTokenCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAdminAccessTokenCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAdminAccessTokenCreateRequest(val *ApiAdminAccessTokenCreateRequest) *NullableApiAdminAccessTokenCreateRequest {
	return &NullableApiAdminAccessTokenCreateRequest{value: val, isSet: true}
}

func (v NullableApiAdminAccessTokenCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAdminAccessTokenCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


