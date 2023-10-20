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

// checks if the ApiCryptoPurchaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCryptoPurchaseRequest{}

// ApiCryptoPurchaseRequest struct for ApiCryptoPurchaseRequest
type ApiCryptoPurchaseRequest struct {
	Currency *string `json:"currency,omitempty"`
	Demo *bool `json:"demo,omitempty"`
	Domain *string `json:"domain,omitempty"`
	OwnerAddress *string `json:"ownerAddress,omitempty"`
	OwnerEmail *string `json:"ownerEmail,omitempty"`
	Years *int32 `json:"years,omitempty"`
}

// NewApiCryptoPurchaseRequest instantiates a new ApiCryptoPurchaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCryptoPurchaseRequest() *ApiCryptoPurchaseRequest {
	this := ApiCryptoPurchaseRequest{}
	return &this
}

// NewApiCryptoPurchaseRequestWithDefaults instantiates a new ApiCryptoPurchaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCryptoPurchaseRequestWithDefaults() *ApiCryptoPurchaseRequest {
	this := ApiCryptoPurchaseRequest{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ApiCryptoPurchaseRequest) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCryptoPurchaseRequest) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ApiCryptoPurchaseRequest) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ApiCryptoPurchaseRequest) SetCurrency(v string) {
	o.Currency = &v
}

// GetDemo returns the Demo field value if set, zero value otherwise.
func (o *ApiCryptoPurchaseRequest) GetDemo() bool {
	if o == nil || IsNil(o.Demo) {
		var ret bool
		return ret
	}
	return *o.Demo
}

// GetDemoOk returns a tuple with the Demo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCryptoPurchaseRequest) GetDemoOk() (*bool, bool) {
	if o == nil || IsNil(o.Demo) {
		return nil, false
	}
	return o.Demo, true
}

// HasDemo returns a boolean if a field has been set.
func (o *ApiCryptoPurchaseRequest) HasDemo() bool {
	if o != nil && !IsNil(o.Demo) {
		return true
	}

	return false
}

// SetDemo gets a reference to the given bool and assigns it to the Demo field.
func (o *ApiCryptoPurchaseRequest) SetDemo(v bool) {
	o.Demo = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiCryptoPurchaseRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCryptoPurchaseRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiCryptoPurchaseRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiCryptoPurchaseRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ApiCryptoPurchaseRequest) GetOwnerAddress() string {
	if o == nil || IsNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCryptoPurchaseRequest) GetOwnerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ApiCryptoPurchaseRequest) HasOwnerAddress() bool {
	if o != nil && !IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ApiCryptoPurchaseRequest) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetOwnerEmail returns the OwnerEmail field value if set, zero value otherwise.
func (o *ApiCryptoPurchaseRequest) GetOwnerEmail() string {
	if o == nil || IsNil(o.OwnerEmail) {
		var ret string
		return ret
	}
	return *o.OwnerEmail
}

// GetOwnerEmailOk returns a tuple with the OwnerEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCryptoPurchaseRequest) GetOwnerEmailOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerEmail) {
		return nil, false
	}
	return o.OwnerEmail, true
}

// HasOwnerEmail returns a boolean if a field has been set.
func (o *ApiCryptoPurchaseRequest) HasOwnerEmail() bool {
	if o != nil && !IsNil(o.OwnerEmail) {
		return true
	}

	return false
}

// SetOwnerEmail gets a reference to the given string and assigns it to the OwnerEmail field.
func (o *ApiCryptoPurchaseRequest) SetOwnerEmail(v string) {
	o.OwnerEmail = &v
}

// GetYears returns the Years field value if set, zero value otherwise.
func (o *ApiCryptoPurchaseRequest) GetYears() int32 {
	if o == nil || IsNil(o.Years) {
		var ret int32
		return ret
	}
	return *o.Years
}

// GetYearsOk returns a tuple with the Years field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiCryptoPurchaseRequest) GetYearsOk() (*int32, bool) {
	if o == nil || IsNil(o.Years) {
		return nil, false
	}
	return o.Years, true
}

// HasYears returns a boolean if a field has been set.
func (o *ApiCryptoPurchaseRequest) HasYears() bool {
	if o != nil && !IsNil(o.Years) {
		return true
	}

	return false
}

// SetYears gets a reference to the given int32 and assigns it to the Years field.
func (o *ApiCryptoPurchaseRequest) SetYears(v int32) {
	o.Years = &v
}

func (o ApiCryptoPurchaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCryptoPurchaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Demo) {
		toSerialize["demo"] = o.Demo
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.OwnerAddress) {
		toSerialize["ownerAddress"] = o.OwnerAddress
	}
	if !IsNil(o.OwnerEmail) {
		toSerialize["ownerEmail"] = o.OwnerEmail
	}
	if !IsNil(o.Years) {
		toSerialize["years"] = o.Years
	}
	return toSerialize, nil
}

type NullableApiCryptoPurchaseRequest struct {
	value *ApiCryptoPurchaseRequest
	isSet bool
}

func (v NullableApiCryptoPurchaseRequest) Get() *ApiCryptoPurchaseRequest {
	return v.value
}

func (v *NullableApiCryptoPurchaseRequest) Set(val *ApiCryptoPurchaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCryptoPurchaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCryptoPurchaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCryptoPurchaseRequest(val *ApiCryptoPurchaseRequest) *NullableApiCryptoPurchaseRequest {
	return &NullableApiCryptoPurchaseRequest{value: val, isSet: true}
}

func (v NullableApiCryptoPurchaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCryptoPurchaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


