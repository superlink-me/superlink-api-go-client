/*
Superlink

API for Superlink

API version: v0.3.25
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiMarketCryptoEstimationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMarketCryptoEstimationResponse{}

// ApiMarketCryptoEstimationResponse struct for ApiMarketCryptoEstimationResponse
type ApiMarketCryptoEstimationResponse struct {
	Estimations []ApiMarketCryptoEstimation `json:"estimations,omitempty"`
}

// NewApiMarketCryptoEstimationResponse instantiates a new ApiMarketCryptoEstimationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMarketCryptoEstimationResponse() *ApiMarketCryptoEstimationResponse {
	this := ApiMarketCryptoEstimationResponse{}
	return &this
}

// NewApiMarketCryptoEstimationResponseWithDefaults instantiates a new ApiMarketCryptoEstimationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMarketCryptoEstimationResponseWithDefaults() *ApiMarketCryptoEstimationResponse {
	this := ApiMarketCryptoEstimationResponse{}
	return &this
}

// GetEstimations returns the Estimations field value if set, zero value otherwise.
func (o *ApiMarketCryptoEstimationResponse) GetEstimations() []ApiMarketCryptoEstimation {
	if o == nil || IsNil(o.Estimations) {
		var ret []ApiMarketCryptoEstimation
		return ret
	}
	return o.Estimations
}

// GetEstimationsOk returns a tuple with the Estimations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoEstimationResponse) GetEstimationsOk() ([]ApiMarketCryptoEstimation, bool) {
	if o == nil || IsNil(o.Estimations) {
		return nil, false
	}
	return o.Estimations, true
}

// HasEstimations returns a boolean if a field has been set.
func (o *ApiMarketCryptoEstimationResponse) HasEstimations() bool {
	if o != nil && !IsNil(o.Estimations) {
		return true
	}

	return false
}

// SetEstimations gets a reference to the given []ApiMarketCryptoEstimation and assigns it to the Estimations field.
func (o *ApiMarketCryptoEstimationResponse) SetEstimations(v []ApiMarketCryptoEstimation) {
	o.Estimations = v
}

func (o ApiMarketCryptoEstimationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMarketCryptoEstimationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Estimations) {
		toSerialize["estimations"] = o.Estimations
	}
	return toSerialize, nil
}

type NullableApiMarketCryptoEstimationResponse struct {
	value *ApiMarketCryptoEstimationResponse
	isSet bool
}

func (v NullableApiMarketCryptoEstimationResponse) Get() *ApiMarketCryptoEstimationResponse {
	return v.value
}

func (v *NullableApiMarketCryptoEstimationResponse) Set(val *ApiMarketCryptoEstimationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMarketCryptoEstimationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMarketCryptoEstimationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMarketCryptoEstimationResponse(val *ApiMarketCryptoEstimationResponse) *NullableApiMarketCryptoEstimationResponse {
	return &NullableApiMarketCryptoEstimationResponse{value: val, isSet: true}
}

func (v NullableApiMarketCryptoEstimationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMarketCryptoEstimationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


