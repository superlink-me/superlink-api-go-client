/*
Superlink

API for Superlink

API version: v0.3.0
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiPurchaseRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPurchaseRequest{}

// ApiPurchaseRequest struct for ApiPurchaseRequest
type ApiPurchaseRequest struct {
	CheckoutId *string `json:"checkoutId,omitempty"`
	Domain *string `json:"domain,omitempty"`
	ExternalUserId *string `json:"externalUserId,omitempty"`
	OwnerAddress *string `json:"ownerAddress,omitempty"`
	StripeConnectedAccountId *string `json:"stripeConnectedAccountId,omitempty"`
}

// NewApiPurchaseRequest instantiates a new ApiPurchaseRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPurchaseRequest() *ApiPurchaseRequest {
	this := ApiPurchaseRequest{}
	return &this
}

// NewApiPurchaseRequestWithDefaults instantiates a new ApiPurchaseRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPurchaseRequestWithDefaults() *ApiPurchaseRequest {
	this := ApiPurchaseRequest{}
	return &this
}

// GetCheckoutId returns the CheckoutId field value if set, zero value otherwise.
func (o *ApiPurchaseRequest) GetCheckoutId() string {
	if o == nil || IsNil(o.CheckoutId) {
		var ret string
		return ret
	}
	return *o.CheckoutId
}

// GetCheckoutIdOk returns a tuple with the CheckoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseRequest) GetCheckoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutId) {
		return nil, false
	}
	return o.CheckoutId, true
}

// HasCheckoutId returns a boolean if a field has been set.
func (o *ApiPurchaseRequest) HasCheckoutId() bool {
	if o != nil && !IsNil(o.CheckoutId) {
		return true
	}

	return false
}

// SetCheckoutId gets a reference to the given string and assigns it to the CheckoutId field.
func (o *ApiPurchaseRequest) SetCheckoutId(v string) {
	o.CheckoutId = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiPurchaseRequest) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseRequest) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiPurchaseRequest) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiPurchaseRequest) SetDomain(v string) {
	o.Domain = &v
}

// GetExternalUserId returns the ExternalUserId field value if set, zero value otherwise.
func (o *ApiPurchaseRequest) GetExternalUserId() string {
	if o == nil || IsNil(o.ExternalUserId) {
		var ret string
		return ret
	}
	return *o.ExternalUserId
}

// GetExternalUserIdOk returns a tuple with the ExternalUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseRequest) GetExternalUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUserId) {
		return nil, false
	}
	return o.ExternalUserId, true
}

// HasExternalUserId returns a boolean if a field has been set.
func (o *ApiPurchaseRequest) HasExternalUserId() bool {
	if o != nil && !IsNil(o.ExternalUserId) {
		return true
	}

	return false
}

// SetExternalUserId gets a reference to the given string and assigns it to the ExternalUserId field.
func (o *ApiPurchaseRequest) SetExternalUserId(v string) {
	o.ExternalUserId = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ApiPurchaseRequest) GetOwnerAddress() string {
	if o == nil || IsNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseRequest) GetOwnerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ApiPurchaseRequest) HasOwnerAddress() bool {
	if o != nil && !IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ApiPurchaseRequest) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetStripeConnectedAccountId returns the StripeConnectedAccountId field value if set, zero value otherwise.
func (o *ApiPurchaseRequest) GetStripeConnectedAccountId() string {
	if o == nil || IsNil(o.StripeConnectedAccountId) {
		var ret string
		return ret
	}
	return *o.StripeConnectedAccountId
}

// GetStripeConnectedAccountIdOk returns a tuple with the StripeConnectedAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiPurchaseRequest) GetStripeConnectedAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripeConnectedAccountId) {
		return nil, false
	}
	return o.StripeConnectedAccountId, true
}

// HasStripeConnectedAccountId returns a boolean if a field has been set.
func (o *ApiPurchaseRequest) HasStripeConnectedAccountId() bool {
	if o != nil && !IsNil(o.StripeConnectedAccountId) {
		return true
	}

	return false
}

// SetStripeConnectedAccountId gets a reference to the given string and assigns it to the StripeConnectedAccountId field.
func (o *ApiPurchaseRequest) SetStripeConnectedAccountId(v string) {
	o.StripeConnectedAccountId = &v
}

func (o ApiPurchaseRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPurchaseRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckoutId) {
		toSerialize["checkoutId"] = o.CheckoutId
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.ExternalUserId) {
		toSerialize["externalUserId"] = o.ExternalUserId
	}
	if !IsNil(o.OwnerAddress) {
		toSerialize["ownerAddress"] = o.OwnerAddress
	}
	if !IsNil(o.StripeConnectedAccountId) {
		toSerialize["stripeConnectedAccountId"] = o.StripeConnectedAccountId
	}
	return toSerialize, nil
}

type NullableApiPurchaseRequest struct {
	value *ApiPurchaseRequest
	isSet bool
}

func (v NullableApiPurchaseRequest) Get() *ApiPurchaseRequest {
	return v.value
}

func (v *NullableApiPurchaseRequest) Set(val *ApiPurchaseRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPurchaseRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPurchaseRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPurchaseRequest(val *ApiPurchaseRequest) *NullableApiPurchaseRequest {
	return &NullableApiPurchaseRequest{value: val, isSet: true}
}

func (v NullableApiPurchaseRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPurchaseRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

