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

// checks if the ApiPurchaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPurchaseResponse{}

// ApiPurchaseResponse struct for ApiPurchaseResponse
type ApiPurchaseResponse struct {
	Currency *string `json:"currency,omitempty"`
	Domain *string `json:"domain,omitempty"`
	NameService *string `json:"nameService,omitempty"`
	OwnerAddress *string `json:"ownerAddress,omitempty"`
	PartnerId *string `json:"partnerId,omitempty"`
	Price *float32 `json:"price,omitempty"`
}

// NewApiPurchaseResponse instantiates a new ApiPurchaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPurchaseResponse() *ApiPurchaseResponse {
	this := ApiPurchaseResponse{}
	return &this
}

// NewApiPurchaseResponseWithDefaults instantiates a new ApiPurchaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPurchaseResponseWithDefaults() *ApiPurchaseResponse {
	this := ApiPurchaseResponse{}
	return &this
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ApiPurchaseResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ApiPurchaseResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ApiPurchaseResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiPurchaseResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiPurchaseResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiPurchaseResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetNameService returns the NameService field value if set, zero value otherwise.
func (o *ApiPurchaseResponse) GetNameService() string {
	if o == nil || IsNil(o.NameService) {
		var ret string
		return ret
	}
	return *o.NameService
}

// GetNameServiceOk returns a tuple with the NameService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseResponse) GetNameServiceOk() (*string, bool) {
	if o == nil || IsNil(o.NameService) {
		return nil, false
	}
	return o.NameService, true
}

// HasNameService returns a boolean if a field has been set.
func (o *ApiPurchaseResponse) HasNameService() bool {
	if o != nil && !IsNil(o.NameService) {
		return true
	}

	return false
}

// SetNameService gets a reference to the given string and assigns it to the NameService field.
func (o *ApiPurchaseResponse) SetNameService(v string) {
	o.NameService = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ApiPurchaseResponse) GetOwnerAddress() string {
	if o == nil || IsNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseResponse) GetOwnerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ApiPurchaseResponse) HasOwnerAddress() bool {
	if o != nil && !IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ApiPurchaseResponse) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetPartnerId returns the PartnerId field value if set, zero value otherwise.
func (o *ApiPurchaseResponse) GetPartnerId() string {
	if o == nil || IsNil(o.PartnerId) {
		var ret string
		return ret
	}
	return *o.PartnerId
}

// GetPartnerIdOk returns a tuple with the PartnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseResponse) GetPartnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartnerId) {
		return nil, false
	}
	return o.PartnerId, true
}

// HasPartnerId returns a boolean if a field has been set.
func (o *ApiPurchaseResponse) HasPartnerId() bool {
	if o != nil && !IsNil(o.PartnerId) {
		return true
	}

	return false
}

// SetPartnerId gets a reference to the given string and assigns it to the PartnerId field.
func (o *ApiPurchaseResponse) SetPartnerId(v string) {
	o.PartnerId = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ApiPurchaseResponse) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseResponse) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ApiPurchaseResponse) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ApiPurchaseResponse) SetPrice(v float32) {
	o.Price = &v
}

func (o ApiPurchaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPurchaseResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OwnerAddress) {
		toSerialize["ownerAddress"] = o.OwnerAddress
	}
	if !IsNil(o.PartnerId) {
		toSerialize["partnerId"] = o.PartnerId
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	return toSerialize, nil
}

type NullableApiPurchaseResponse struct {
	value *ApiPurchaseResponse
	isSet bool
}

func (v NullableApiPurchaseResponse) Get() *ApiPurchaseResponse {
	return v.value
}

func (v *NullableApiPurchaseResponse) Set(val *ApiPurchaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPurchaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPurchaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPurchaseResponse(val *ApiPurchaseResponse) *NullableApiPurchaseResponse {
	return &NullableApiPurchaseResponse{value: val, isSet: true}
}

func (v NullableApiPurchaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPurchaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


