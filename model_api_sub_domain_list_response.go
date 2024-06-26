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

// checks if the ApiSubDomainListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiSubDomainListResponse{}

// ApiSubDomainListResponse struct for ApiSubDomainListResponse
type ApiSubDomainListResponse struct {
	Names []ApiSubDomainItem `json:"names,omitempty"`
	NextPaginationToken *string `json:"nextPaginationToken,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewApiSubDomainListResponse instantiates a new ApiSubDomainListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiSubDomainListResponse() *ApiSubDomainListResponse {
	this := ApiSubDomainListResponse{}
	return &this
}

// NewApiSubDomainListResponseWithDefaults instantiates a new ApiSubDomainListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiSubDomainListResponseWithDefaults() *ApiSubDomainListResponse {
	this := ApiSubDomainListResponse{}
	return &this
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *ApiSubDomainListResponse) GetNames() []ApiSubDomainItem {
	if o == nil || IsNil(o.Names) {
		var ret []ApiSubDomainItem
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainListResponse) GetNamesOk() ([]ApiSubDomainItem, bool) {
	if o == nil || IsNil(o.Names) {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *ApiSubDomainListResponse) HasNames() bool {
	if o != nil && !IsNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []ApiSubDomainItem and assigns it to the Names field.
func (o *ApiSubDomainListResponse) SetNames(v []ApiSubDomainItem) {
	o.Names = v
}

// GetNextPaginationToken returns the NextPaginationToken field value if set, zero value otherwise.
func (o *ApiSubDomainListResponse) GetNextPaginationToken() string {
	if o == nil || IsNil(o.NextPaginationToken) {
		var ret string
		return ret
	}
	return *o.NextPaginationToken
}

// GetNextPaginationTokenOk returns a tuple with the NextPaginationToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainListResponse) GetNextPaginationTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPaginationToken) {
		return nil, false
	}
	return o.NextPaginationToken, true
}

// HasNextPaginationToken returns a boolean if a field has been set.
func (o *ApiSubDomainListResponse) HasNextPaginationToken() bool {
	if o != nil && !IsNil(o.NextPaginationToken) {
		return true
	}

	return false
}

// SetNextPaginationToken gets a reference to the given string and assigns it to the NextPaginationToken field.
func (o *ApiSubDomainListResponse) SetNextPaginationToken(v string) {
	o.NextPaginationToken = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *ApiSubDomainListResponse) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiSubDomainListResponse) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *ApiSubDomainListResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *ApiSubDomainListResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o ApiSubDomainListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiSubDomainListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !IsNil(o.NextPaginationToken) {
		toSerialize["nextPaginationToken"] = o.NextPaginationToken
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableApiSubDomainListResponse struct {
	value *ApiSubDomainListResponse
	isSet bool
}

func (v NullableApiSubDomainListResponse) Get() *ApiSubDomainListResponse {
	return v.value
}

func (v *NullableApiSubDomainListResponse) Set(val *ApiSubDomainListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiSubDomainListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiSubDomainListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiSubDomainListResponse(val *ApiSubDomainListResponse) *NullableApiSubDomainListResponse {
	return &NullableApiSubDomainListResponse{value: val, isSet: true}
}

func (v NullableApiSubDomainListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiSubDomainListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


