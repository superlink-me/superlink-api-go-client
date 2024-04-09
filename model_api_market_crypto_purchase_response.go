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

// checks if the ApiMarketCryptoPurchaseResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMarketCryptoPurchaseResponse{}

// ApiMarketCryptoPurchaseResponse struct for ApiMarketCryptoPurchaseResponse
type ApiMarketCryptoPurchaseResponse struct {
	Address *string `json:"address,omitempty"`
	// PaymentDetails     CryptoPaymentDetails `json:\"paymentDetails\"`
	Amount *float32 `json:"amount,omitempty"`
	ExpiryDateEpoch *int32 `json:"expiryDateEpoch,omitempty"`
	OrderId *string `json:"orderId,omitempty"`
	PaymentId *string `json:"paymentId,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Uri *string `json:"uri,omitempty"`
}

// NewApiMarketCryptoPurchaseResponse instantiates a new ApiMarketCryptoPurchaseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMarketCryptoPurchaseResponse() *ApiMarketCryptoPurchaseResponse {
	this := ApiMarketCryptoPurchaseResponse{}
	return &this
}

// NewApiMarketCryptoPurchaseResponseWithDefaults instantiates a new ApiMarketCryptoPurchaseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMarketCryptoPurchaseResponseWithDefaults() *ApiMarketCryptoPurchaseResponse {
	this := ApiMarketCryptoPurchaseResponse{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ApiMarketCryptoPurchaseResponse) SetAddress(v string) {
	o.Address = &v
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *ApiMarketCryptoPurchaseResponse) SetAmount(v float32) {
	o.Amount = &v
}

// GetExpiryDateEpoch returns the ExpiryDateEpoch field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetExpiryDateEpoch() int32 {
	if o == nil || IsNil(o.ExpiryDateEpoch) {
		var ret int32
		return ret
	}
	return *o.ExpiryDateEpoch
}

// GetExpiryDateEpochOk returns a tuple with the ExpiryDateEpoch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetExpiryDateEpochOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiryDateEpoch) {
		return nil, false
	}
	return o.ExpiryDateEpoch, true
}

// HasExpiryDateEpoch returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasExpiryDateEpoch() bool {
	if o != nil && !IsNil(o.ExpiryDateEpoch) {
		return true
	}

	return false
}

// SetExpiryDateEpoch gets a reference to the given int32 and assigns it to the ExpiryDateEpoch field.
func (o *ApiMarketCryptoPurchaseResponse) SetExpiryDateEpoch(v int32) {
	o.ExpiryDateEpoch = &v
}

// GetOrderId returns the OrderId field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetOrderId() string {
	if o == nil || IsNil(o.OrderId) {
		var ret string
		return ret
	}
	return *o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetOrderIdOk() (*string, bool) {
	if o == nil || IsNil(o.OrderId) {
		return nil, false
	}
	return o.OrderId, true
}

// HasOrderId returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasOrderId() bool {
	if o != nil && !IsNil(o.OrderId) {
		return true
	}

	return false
}

// SetOrderId gets a reference to the given string and assigns it to the OrderId field.
func (o *ApiMarketCryptoPurchaseResponse) SetOrderId(v string) {
	o.OrderId = &v
}

// GetPaymentId returns the PaymentId field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetPaymentId() string {
	if o == nil || IsNil(o.PaymentId) {
		var ret string
		return ret
	}
	return *o.PaymentId
}

// GetPaymentIdOk returns a tuple with the PaymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetPaymentIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentId) {
		return nil, false
	}
	return o.PaymentId, true
}

// HasPaymentId returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasPaymentId() bool {
	if o != nil && !IsNil(o.PaymentId) {
		return true
	}

	return false
}

// SetPaymentId gets a reference to the given string and assigns it to the PaymentId field.
func (o *ApiMarketCryptoPurchaseResponse) SetPaymentId(v string) {
	o.PaymentId = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ApiMarketCryptoPurchaseResponse) SetProtocol(v string) {
	o.Protocol = &v
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *ApiMarketCryptoPurchaseResponse) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketCryptoPurchaseResponse) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *ApiMarketCryptoPurchaseResponse) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *ApiMarketCryptoPurchaseResponse) SetUri(v string) {
	o.Uri = &v
}

func (o ApiMarketCryptoPurchaseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMarketCryptoPurchaseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.ExpiryDateEpoch) {
		toSerialize["expiryDateEpoch"] = o.ExpiryDateEpoch
	}
	if !IsNil(o.OrderId) {
		toSerialize["orderId"] = o.OrderId
	}
	if !IsNil(o.PaymentId) {
		toSerialize["paymentId"] = o.PaymentId
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	return toSerialize, nil
}

type NullableApiMarketCryptoPurchaseResponse struct {
	value *ApiMarketCryptoPurchaseResponse
	isSet bool
}

func (v NullableApiMarketCryptoPurchaseResponse) Get() *ApiMarketCryptoPurchaseResponse {
	return v.value
}

func (v *NullableApiMarketCryptoPurchaseResponse) Set(val *ApiMarketCryptoPurchaseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMarketCryptoPurchaseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMarketCryptoPurchaseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMarketCryptoPurchaseResponse(val *ApiMarketCryptoPurchaseResponse) *NullableApiMarketCryptoPurchaseResponse {
	return &NullableApiMarketCryptoPurchaseResponse{value: val, isSet: true}
}

func (v NullableApiMarketCryptoPurchaseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMarketCryptoPurchaseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


