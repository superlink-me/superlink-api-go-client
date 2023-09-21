/*
Superlink

API for Superlink

API version: v0.3.19
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiMarketSearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMarketSearchResponse{}

// ApiMarketSearchResponse struct for ApiMarketSearchResponse
type ApiMarketSearchResponse struct {
	MarketListings []ApiMarketListing `json:"marketListings,omitempty"`
}

// NewApiMarketSearchResponse instantiates a new ApiMarketSearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMarketSearchResponse() *ApiMarketSearchResponse {
	this := ApiMarketSearchResponse{}
	return &this
}

// NewApiMarketSearchResponseWithDefaults instantiates a new ApiMarketSearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMarketSearchResponseWithDefaults() *ApiMarketSearchResponse {
	this := ApiMarketSearchResponse{}
	return &this
}

// GetMarketListings returns the MarketListings field value if set, zero value otherwise.
func (o *ApiMarketSearchResponse) GetMarketListings() []ApiMarketListing {
	if o == nil || IsNil(o.MarketListings) {
		var ret []ApiMarketListing
		return ret
	}
	return o.MarketListings
}

// GetMarketListingsOk returns a tuple with the MarketListings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketSearchResponse) GetMarketListingsOk() ([]ApiMarketListing, bool) {
	if o == nil || IsNil(o.MarketListings) {
		return nil, false
	}
	return o.MarketListings, true
}

// HasMarketListings returns a boolean if a field has been set.
func (o *ApiMarketSearchResponse) HasMarketListings() bool {
	if o != nil && !IsNil(o.MarketListings) {
		return true
	}

	return false
}

// SetMarketListings gets a reference to the given []ApiMarketListing and assigns it to the MarketListings field.
func (o *ApiMarketSearchResponse) SetMarketListings(v []ApiMarketListing) {
	o.MarketListings = v
}

func (o ApiMarketSearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMarketSearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MarketListings) {
		toSerialize["marketListings"] = o.MarketListings
	}
	return toSerialize, nil
}

type NullableApiMarketSearchResponse struct {
	value *ApiMarketSearchResponse
	isSet bool
}

func (v NullableApiMarketSearchResponse) Get() *ApiMarketSearchResponse {
	return v.value
}

func (v *NullableApiMarketSearchResponse) Set(val *ApiMarketSearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMarketSearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMarketSearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMarketSearchResponse(val *ApiMarketSearchResponse) *NullableApiMarketSearchResponse {
	return &NullableApiMarketSearchResponse{value: val, isSet: true}
}

func (v NullableApiMarketSearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMarketSearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


