/*
Superlink

API for Superlink

API version: v0.3.21
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiMarketPurchaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMarketPurchaseResponse{}

// ApiMarketPurchaseResponse struct for ApiMarketPurchaseResponse
type ApiMarketPurchaseResponse struct {
	CheckoutId *string `json:"checkoutId,omitempty"`
	StripeCustomerId *string `json:"stripeCustomerId,omitempty"`
	StripeEphemeralKey *string `json:"stripeEphemeralKey,omitempty"`
	StripePaymentIntent *string `json:"stripePaymentIntent,omitempty"`
	StripePublishableKey *string `json:"stripePublishableKey,omitempty"`
}

// NewApiMarketPurchaseResponse instantiates a new ApiMarketPurchaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMarketPurchaseResponse() *ApiMarketPurchaseResponse {
	this := ApiMarketPurchaseResponse{}
	return &this
}

// NewApiMarketPurchaseResponseWithDefaults instantiates a new ApiMarketPurchaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMarketPurchaseResponseWithDefaults() *ApiMarketPurchaseResponse {
	this := ApiMarketPurchaseResponse{}
	return &this
}

// GetCheckoutId returns the CheckoutId field value if set, zero value otherwise.
func (o *ApiMarketPurchaseResponse) GetCheckoutId() string {
	if o == nil || IsNil(o.CheckoutId) {
		var ret string
		return ret
	}
	return *o.CheckoutId
}

// GetCheckoutIdOk returns a tuple with the CheckoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketPurchaseResponse) GetCheckoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.CheckoutId) {
		return nil, false
	}
	return o.CheckoutId, true
}

// HasCheckoutId returns a boolean if a field has been set.
func (o *ApiMarketPurchaseResponse) HasCheckoutId() bool {
	if o != nil && !IsNil(o.CheckoutId) {
		return true
	}

	return false
}

// SetCheckoutId gets a reference to the given string and assigns it to the CheckoutId field.
func (o *ApiMarketPurchaseResponse) SetCheckoutId(v string) {
	o.CheckoutId = &v
}

// GetStripeCustomerId returns the StripeCustomerId field value if set, zero value otherwise.
func (o *ApiMarketPurchaseResponse) GetStripeCustomerId() string {
	if o == nil || IsNil(o.StripeCustomerId) {
		var ret string
		return ret
	}
	return *o.StripeCustomerId
}

// GetStripeCustomerIdOk returns a tuple with the StripeCustomerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketPurchaseResponse) GetStripeCustomerIdOk() (*string, bool) {
	if o == nil || IsNil(o.StripeCustomerId) {
		return nil, false
	}
	return o.StripeCustomerId, true
}

// HasStripeCustomerId returns a boolean if a field has been set.
func (o *ApiMarketPurchaseResponse) HasStripeCustomerId() bool {
	if o != nil && !IsNil(o.StripeCustomerId) {
		return true
	}

	return false
}

// SetStripeCustomerId gets a reference to the given string and assigns it to the StripeCustomerId field.
func (o *ApiMarketPurchaseResponse) SetStripeCustomerId(v string) {
	o.StripeCustomerId = &v
}

// GetStripeEphemeralKey returns the StripeEphemeralKey field value if set, zero value otherwise.
func (o *ApiMarketPurchaseResponse) GetStripeEphemeralKey() string {
	if o == nil || IsNil(o.StripeEphemeralKey) {
		var ret string
		return ret
	}
	return *o.StripeEphemeralKey
}

// GetStripeEphemeralKeyOk returns a tuple with the StripeEphemeralKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketPurchaseResponse) GetStripeEphemeralKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StripeEphemeralKey) {
		return nil, false
	}
	return o.StripeEphemeralKey, true
}

// HasStripeEphemeralKey returns a boolean if a field has been set.
func (o *ApiMarketPurchaseResponse) HasStripeEphemeralKey() bool {
	if o != nil && !IsNil(o.StripeEphemeralKey) {
		return true
	}

	return false
}

// SetStripeEphemeralKey gets a reference to the given string and assigns it to the StripeEphemeralKey field.
func (o *ApiMarketPurchaseResponse) SetStripeEphemeralKey(v string) {
	o.StripeEphemeralKey = &v
}

// GetStripePaymentIntent returns the StripePaymentIntent field value if set, zero value otherwise.
func (o *ApiMarketPurchaseResponse) GetStripePaymentIntent() string {
	if o == nil || IsNil(o.StripePaymentIntent) {
		var ret string
		return ret
	}
	return *o.StripePaymentIntent
}

// GetStripePaymentIntentOk returns a tuple with the StripePaymentIntent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketPurchaseResponse) GetStripePaymentIntentOk() (*string, bool) {
	if o == nil || IsNil(o.StripePaymentIntent) {
		return nil, false
	}
	return o.StripePaymentIntent, true
}

// HasStripePaymentIntent returns a boolean if a field has been set.
func (o *ApiMarketPurchaseResponse) HasStripePaymentIntent() bool {
	if o != nil && !IsNil(o.StripePaymentIntent) {
		return true
	}

	return false
}

// SetStripePaymentIntent gets a reference to the given string and assigns it to the StripePaymentIntent field.
func (o *ApiMarketPurchaseResponse) SetStripePaymentIntent(v string) {
	o.StripePaymentIntent = &v
}

// GetStripePublishableKey returns the StripePublishableKey field value if set, zero value otherwise.
func (o *ApiMarketPurchaseResponse) GetStripePublishableKey() string {
	if o == nil || IsNil(o.StripePublishableKey) {
		var ret string
		return ret
	}
	return *o.StripePublishableKey
}

// GetStripePublishableKeyOk returns a tuple with the StripePublishableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketPurchaseResponse) GetStripePublishableKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StripePublishableKey) {
		return nil, false
	}
	return o.StripePublishableKey, true
}

// HasStripePublishableKey returns a boolean if a field has been set.
func (o *ApiMarketPurchaseResponse) HasStripePublishableKey() bool {
	if o != nil && !IsNil(o.StripePublishableKey) {
		return true
	}

	return false
}

// SetStripePublishableKey gets a reference to the given string and assigns it to the StripePublishableKey field.
func (o *ApiMarketPurchaseResponse) SetStripePublishableKey(v string) {
	o.StripePublishableKey = &v
}

func (o ApiMarketPurchaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMarketPurchaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CheckoutId) {
		toSerialize["checkoutId"] = o.CheckoutId
	}
	if !IsNil(o.StripeCustomerId) {
		toSerialize["stripeCustomerId"] = o.StripeCustomerId
	}
	if !IsNil(o.StripeEphemeralKey) {
		toSerialize["stripeEphemeralKey"] = o.StripeEphemeralKey
	}
	if !IsNil(o.StripePaymentIntent) {
		toSerialize["stripePaymentIntent"] = o.StripePaymentIntent
	}
	if !IsNil(o.StripePublishableKey) {
		toSerialize["stripePublishableKey"] = o.StripePublishableKey
	}
	return toSerialize, nil
}

type NullableApiMarketPurchaseResponse struct {
	value *ApiMarketPurchaseResponse
	isSet bool
}

func (v NullableApiMarketPurchaseResponse) Get() *ApiMarketPurchaseResponse {
	return v.value
}

func (v *NullableApiMarketPurchaseResponse) Set(val *ApiMarketPurchaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMarketPurchaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMarketPurchaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMarketPurchaseResponse(val *ApiMarketPurchaseResponse) *NullableApiMarketPurchaseResponse {
	return &NullableApiMarketPurchaseResponse{value: val, isSet: true}
}

func (v NullableApiMarketPurchaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMarketPurchaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


