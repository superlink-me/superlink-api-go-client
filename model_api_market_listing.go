/*
Superlink

API for Superlink

API version: v0.3.15
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiMarketListing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMarketListing{}

// ApiMarketListing struct for ApiMarketListing
type ApiMarketListing struct {
	Currency *string `json:"currency,omitempty"`
	Domain *string `json:"domain,omitempty"`
	NameService *string `json:"nameService,omitempty"`
	Price *map[string]float32 `json:"price,omitempty"`
}

// NewApiMarketListing instantiates a new ApiMarketListing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMarketListing() *ApiMarketListing {
	this := ApiMarketListing{}
	return &this
}

// NewApiMarketListingWithDefaults instantiates a new ApiMarketListing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMarketListingWithDefaults() *ApiMarketListing {
	this := ApiMarketListing{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ApiMarketListing) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketListing) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ApiMarketListing) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ApiMarketListing) SetCurrency(v string) {
	o.Currency = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiMarketListing) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketListing) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiMarketListing) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiMarketListing) SetDomain(v string) {
	o.Domain = &v
}

// GetNameService returns the NameService field value if set, zero value otherwise.
func (o *ApiMarketListing) GetNameService() string {
	if o == nil || IsNil(o.NameService) {
		var ret string
		return ret
	}
	return *o.NameService
}

// GetNameServiceOk returns a tuple with the NameService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketListing) GetNameServiceOk() (*string, bool) {
	if o == nil || IsNil(o.NameService) {
		return nil, false
	}
	return o.NameService, true
}

// HasNameService returns a boolean if a field has been set.
func (o *ApiMarketListing) HasNameService() bool {
	if o != nil && !IsNil(o.NameService) {
		return true
	}

	return false
}

// SetNameService gets a reference to the given string and assigns it to the NameService field.
func (o *ApiMarketListing) SetNameService(v string) {
	o.NameService = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ApiMarketListing) GetPrice() map[string]float32 {
	if o == nil || IsNil(o.Price) {
		var ret map[string]float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketListing) GetPriceOk() (*map[string]float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ApiMarketListing) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given map[string]float32 and assigns it to the Price field.
func (o *ApiMarketListing) SetPrice(v map[string]float32) {
	o.Price = &v
}

func (o ApiMarketListing) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMarketListing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.NameService) {
		toSerialize["nameService"] = o.NameService
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableApiMarketListing struct {
	value *ApiMarketListing
	isSet bool
}

func (v NullableApiMarketListing) Get() *ApiMarketListing {
	return v.value
}

func (v *NullableApiMarketListing) Set(val *ApiMarketListing) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMarketListing) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMarketListing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMarketListing(val *ApiMarketListing) *NullableApiMarketListing {
	return &NullableApiMarketListing{value: val, isSet: true}
}

func (v NullableApiMarketListing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMarketListing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


