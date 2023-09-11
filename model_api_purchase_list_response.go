/*
Superlink

API for Superlink

API version: v0.3.8
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiPurchaseListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPurchaseListResponse{}

// ApiPurchaseListResponse struct for ApiPurchaseListResponse
type ApiPurchaseListResponse struct {
	Purchases []ApiPurchaseResponse `json:"purchases,omitempty"`
}

// NewApiPurchaseListResponse instantiates a new ApiPurchaseListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPurchaseListResponse() *ApiPurchaseListResponse {
	this := ApiPurchaseListResponse{}
	return &this
}

// NewApiPurchaseListResponseWithDefaults instantiates a new ApiPurchaseListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPurchaseListResponseWithDefaults() *ApiPurchaseListResponse {
	this := ApiPurchaseListResponse{}
	return &this
}

// GetPurchases returns the Purchases field value if set, zero value otherwise.
func (o *ApiPurchaseListResponse) GetPurchases() []ApiPurchaseResponse {
	if o == nil || IsNil(o.Purchases) {
		var ret []ApiPurchaseResponse
		return ret
	}
	return o.Purchases
}

// GetPurchasesOk returns a tuple with the Purchases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseListResponse) GetPurchasesOk() ([]ApiPurchaseResponse, bool) {
	if o == nil || IsNil(o.Purchases) {
		return nil, false
	}
	return o.Purchases, true
}

// HasPurchases returns a boolean if a field has been set.
func (o *ApiPurchaseListResponse) HasPurchases() bool {
	if o != nil && !IsNil(o.Purchases) {
		return true
	}

	return false
}

// SetPurchases gets a reference to the given []ApiPurchaseResponse and assigns it to the Purchases field.
func (o *ApiPurchaseListResponse) SetPurchases(v []ApiPurchaseResponse) {
	o.Purchases = v
}

func (o ApiPurchaseListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPurchaseListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Purchases) {
		toSerialize["purchases"] = o.Purchases
	}
	return toSerialize, nil
}

type NullableApiPurchaseListResponse struct {
	value *ApiPurchaseListResponse
	isSet bool
}

func (v NullableApiPurchaseListResponse) Get() *ApiPurchaseListResponse {
	return v.value
}

func (v *NullableApiPurchaseListResponse) Set(val *ApiPurchaseListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPurchaseListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPurchaseListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPurchaseListResponse(val *ApiPurchaseListResponse) *NullableApiPurchaseListResponse {
	return &NullableApiPurchaseListResponse{value: val, isSet: true}
}

func (v NullableApiPurchaseListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPurchaseListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


