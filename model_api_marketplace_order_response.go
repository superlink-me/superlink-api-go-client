/*
Superlink

API for Superlink

API version: v0.4.2
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiMarketplaceOrderResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiMarketplaceOrderResponse{}

// ApiMarketplaceOrderResponse struct for ApiMarketplaceOrderResponse
type ApiMarketplaceOrderResponse struct {
	BaseCurrency *string `json:"baseCurrency,omitempty"`
	BasePrice *float32 `json:"basePrice,omitempty"`
	CreatedAt *string `json:"createdAt,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Domain *string `json:"domain,omitempty"`
	Id *string `json:"id,omitempty"`
	NameService *string `json:"nameService,omitempty"`
	OrderStatus *string `json:"orderStatus,omitempty"`
	OrderStatusReason *string `json:"orderStatusReason,omitempty"`
	OwnerAddress *string `json:"ownerAddress,omitempty"`
	PaymentReferenceId *string `json:"paymentReferenceId,omitempty"`
	PaymentType *string `json:"paymentType,omitempty"`
	Price *float32 `json:"price,omitempty"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewApiMarketplaceOrderResponse instantiates a new ApiMarketplaceOrderResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiMarketplaceOrderResponse() *ApiMarketplaceOrderResponse {
	this := ApiMarketplaceOrderResponse{}
	return &this
}

// NewApiMarketplaceOrderResponseWithDefaults instantiates a new ApiMarketplaceOrderResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiMarketplaceOrderResponseWithDefaults() *ApiMarketplaceOrderResponse {
	this := ApiMarketplaceOrderResponse{}
	return &this
}

// GetBaseCurrency returns the BaseCurrency field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetBaseCurrency() string {
	if o == nil || IsNil(o.BaseCurrency) {
		var ret string
		return ret
	}
	return *o.BaseCurrency
}

// GetBaseCurrencyOk returns a tuple with the BaseCurrency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetBaseCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.BaseCurrency) {
		return nil, false
	}
	return o.BaseCurrency, true
}

// HasBaseCurrency returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasBaseCurrency() bool {
	if o != nil && !IsNil(o.BaseCurrency) {
		return true
	}

	return false
}

// SetBaseCurrency gets a reference to the given string and assigns it to the BaseCurrency field.
func (o *ApiMarketplaceOrderResponse) SetBaseCurrency(v string) {
	o.BaseCurrency = &v
}

// GetBasePrice returns the BasePrice field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetBasePrice() float32 {
	if o == nil || IsNil(o.BasePrice) {
		var ret float32
		return ret
	}
	return *o.BasePrice
}

// GetBasePriceOk returns a tuple with the BasePrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetBasePriceOk() (*float32, bool) {
	if o == nil || IsNil(o.BasePrice) {
		return nil, false
	}
	return o.BasePrice, true
}

// HasBasePrice returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasBasePrice() bool {
	if o != nil && !IsNil(o.BasePrice) {
		return true
	}

	return false
}

// SetBasePrice gets a reference to the given float32 and assigns it to the BasePrice field.
func (o *ApiMarketplaceOrderResponse) SetBasePrice(v float32) {
	o.BasePrice = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ApiMarketplaceOrderResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// HasCurrency returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *ApiMarketplaceOrderResponse) SetCurrency(v string) {
	o.Currency = &v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiMarketplaceOrderResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiMarketplaceOrderResponse) SetId(v string) {
	o.Id = &v
}

// GetNameService returns the NameService field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetNameService() string {
	if o == nil || IsNil(o.NameService) {
		var ret string
		return ret
	}
	return *o.NameService
}

// GetNameServiceOk returns a tuple with the NameService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetNameServiceOk() (*string, bool) {
	if o == nil || IsNil(o.NameService) {
		return nil, false
	}
	return o.NameService, true
}

// HasNameService returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasNameService() bool {
	if o != nil && !IsNil(o.NameService) {
		return true
	}

	return false
}

// SetNameService gets a reference to the given string and assigns it to the NameService field.
func (o *ApiMarketplaceOrderResponse) SetNameService(v string) {
	o.NameService = &v
}

// GetOrderStatus returns the OrderStatus field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetOrderStatus() string {
	if o == nil || IsNil(o.OrderStatus) {
		var ret string
		return ret
	}
	return *o.OrderStatus
}

// GetOrderStatusOk returns a tuple with the OrderStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetOrderStatusOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatus) {
		return nil, false
	}
	return o.OrderStatus, true
}

// HasOrderStatus returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasOrderStatus() bool {
	if o != nil && !IsNil(o.OrderStatus) {
		return true
	}

	return false
}

// SetOrderStatus gets a reference to the given string and assigns it to the OrderStatus field.
func (o *ApiMarketplaceOrderResponse) SetOrderStatus(v string) {
	o.OrderStatus = &v
}

// GetOrderStatusReason returns the OrderStatusReason field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetOrderStatusReason() string {
	if o == nil || IsNil(o.OrderStatusReason) {
		var ret string
		return ret
	}
	return *o.OrderStatusReason
}

// GetOrderStatusReasonOk returns a tuple with the OrderStatusReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetOrderStatusReasonOk() (*string, bool) {
	if o == nil || IsNil(o.OrderStatusReason) {
		return nil, false
	}
	return o.OrderStatusReason, true
}

// HasOrderStatusReason returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasOrderStatusReason() bool {
	if o != nil && !IsNil(o.OrderStatusReason) {
		return true
	}

	return false
}

// SetOrderStatusReason gets a reference to the given string and assigns it to the OrderStatusReason field.
func (o *ApiMarketplaceOrderResponse) SetOrderStatusReason(v string) {
	o.OrderStatusReason = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetOwnerAddress() string {
	if o == nil || IsNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetOwnerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasOwnerAddress() bool {
	if o != nil && !IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ApiMarketplaceOrderResponse) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetPaymentReferenceId returns the PaymentReferenceId field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetPaymentReferenceId() string {
	if o == nil || IsNil(o.PaymentReferenceId) {
		var ret string
		return ret
	}
	return *o.PaymentReferenceId
}

// GetPaymentReferenceIdOk returns a tuple with the PaymentReferenceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetPaymentReferenceIdOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentReferenceId) {
		return nil, false
	}
	return o.PaymentReferenceId, true
}

// HasPaymentReferenceId returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasPaymentReferenceId() bool {
	if o != nil && !IsNil(o.PaymentReferenceId) {
		return true
	}

	return false
}

// SetPaymentReferenceId gets a reference to the given string and assigns it to the PaymentReferenceId field.
func (o *ApiMarketplaceOrderResponse) SetPaymentReferenceId(v string) {
	o.PaymentReferenceId = &v
}

// GetPaymentType returns the PaymentType field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetPaymentType() string {
	if o == nil || IsNil(o.PaymentType) {
		var ret string
		return ret
	}
	return *o.PaymentType
}

// GetPaymentTypeOk returns a tuple with the PaymentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetPaymentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentType) {
		return nil, false
	}
	return o.PaymentType, true
}

// HasPaymentType returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasPaymentType() bool {
	if o != nil && !IsNil(o.PaymentType) {
		return true
	}

	return false
}

// SetPaymentType gets a reference to the given string and assigns it to the PaymentType field.
func (o *ApiMarketplaceOrderResponse) SetPaymentType(v string) {
	o.PaymentType = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// HasPrice returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *ApiMarketplaceOrderResponse) SetPrice(v float32) {
	o.Price = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ApiMarketplaceOrderResponse) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiMarketplaceOrderResponse) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ApiMarketplaceOrderResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *ApiMarketplaceOrderResponse) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o ApiMarketplaceOrderResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiMarketplaceOrderResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BaseCurrency) {
		toSerialize["baseCurrency"] = o.BaseCurrency
	}
	if !IsNil(o.BasePrice) {
		toSerialize["basePrice"] = o.BasePrice
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.NameService) {
		toSerialize["nameService"] = o.NameService
	}
	if !IsNil(o.OrderStatus) {
		toSerialize["orderStatus"] = o.OrderStatus
	}
	if !IsNil(o.OrderStatusReason) {
		toSerialize["orderStatusReason"] = o.OrderStatusReason
	}
	if !IsNil(o.OwnerAddress) {
		toSerialize["ownerAddress"] = o.OwnerAddress
	}
	if !IsNil(o.PaymentReferenceId) {
		toSerialize["paymentReferenceId"] = o.PaymentReferenceId
	}
	if !IsNil(o.PaymentType) {
		toSerialize["paymentType"] = o.PaymentType
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableApiMarketplaceOrderResponse struct {
	value *ApiMarketplaceOrderResponse
	isSet bool
}

func (v NullableApiMarketplaceOrderResponse) Get() *ApiMarketplaceOrderResponse {
	return v.value
}

func (v *NullableApiMarketplaceOrderResponse) Set(val *ApiMarketplaceOrderResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiMarketplaceOrderResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiMarketplaceOrderResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiMarketplaceOrderResponse(val *ApiMarketplaceOrderResponse) *NullableApiMarketplaceOrderResponse {
	return &NullableApiMarketplaceOrderResponse{value: val, isSet: true}
}

func (v NullableApiMarketplaceOrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiMarketplaceOrderResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


