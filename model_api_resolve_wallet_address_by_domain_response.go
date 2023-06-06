/*
Resolution API

The Resolution API resolve Superlink domains

API version: 1.0
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ResolutionAPI

import (
	"encoding/json"
)

// checks if the ApiResolveWalletAddressByDomainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiResolveWalletAddressByDomainResponse{}

// ApiResolveWalletAddressByDomainResponse struct for ApiResolveWalletAddressByDomainResponse
type ApiResolveWalletAddressByDomainResponse struct {
	Domain *string `json:"domain,omitempty"`
	Records []ApiDNSRecord `json:"records,omitempty"`
	Wallets *map[string]string `json:"wallets,omitempty"`
}

// NewApiResolveWalletAddressByDomainResponse instantiates a new ApiResolveWalletAddressByDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResolveWalletAddressByDomainResponse() *ApiResolveWalletAddressByDomainResponse {
	this := ApiResolveWalletAddressByDomainResponse{}
	return &this
}

// NewApiResolveWalletAddressByDomainResponseWithDefaults instantiates a new ApiResolveWalletAddressByDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResolveWalletAddressByDomainResponseWithDefaults() *ApiResolveWalletAddressByDomainResponse {
	this := ApiResolveWalletAddressByDomainResponse{}
	return &this
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiResolveWalletAddressByDomainResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveWalletAddressByDomainResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiResolveWalletAddressByDomainResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiResolveWalletAddressByDomainResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetRecords returns the Records field value if set, zero value otherwise.
func (o *ApiResolveWalletAddressByDomainResponse) GetRecords() []ApiDNSRecord {
	if o == nil || IsNil(o.Records) {
		var ret []ApiDNSRecord
		return ret
	}
	return o.Records
}

// GetRecordsOk returns a tuple with the Records field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveWalletAddressByDomainResponse) GetRecordsOk() ([]ApiDNSRecord, bool) {
	if o == nil || IsNil(o.Records) {
		return nil, false
	}
	return o.Records, true
}

// HasRecords returns a boolean if a field has been set.
func (o *ApiResolveWalletAddressByDomainResponse) HasRecords() bool {
	if o != nil && !IsNil(o.Records) {
		return true
	}

	return false
}

// SetRecords gets a reference to the given []ApiDNSRecord and assigns it to the Records field.
func (o *ApiResolveWalletAddressByDomainResponse) SetRecords(v []ApiDNSRecord) {
	o.Records = v
}

// GetWallets returns the Wallets field value if set, zero value otherwise.
func (o *ApiResolveWalletAddressByDomainResponse) GetWallets() map[string]string {
	if o == nil || IsNil(o.Wallets) {
		var ret map[string]string
		return ret
	}
	return *o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveWalletAddressByDomainResponse) GetWalletsOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Wallets) {
		return nil, false
	}
	return o.Wallets, true
}

// HasWallets returns a boolean if a field has been set.
func (o *ApiResolveWalletAddressByDomainResponse) HasWallets() bool {
	if o != nil && !IsNil(o.Wallets) {
		return true
	}

	return false
}

// SetWallets gets a reference to the given map[string]string and assigns it to the Wallets field.
func (o *ApiResolveWalletAddressByDomainResponse) SetWallets(v map[string]string) {
	o.Wallets = &v
}

func (o ApiResolveWalletAddressByDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiResolveWalletAddressByDomainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Domain) {
		toSerialize["domain"] = o.Domain
	}
	if !IsNil(o.Records) {
		toSerialize["records"] = o.Records
	}
	if !IsNil(o.Wallets) {
		toSerialize["wallets"] = o.Wallets
	}
	return toSerialize, nil
}

type NullableApiResolveWalletAddressByDomainResponse struct {
	value *ApiResolveWalletAddressByDomainResponse
	isSet bool
}

func (v NullableApiResolveWalletAddressByDomainResponse) Get() *ApiResolveWalletAddressByDomainResponse {
	return v.value
}

func (v *NullableApiResolveWalletAddressByDomainResponse) Set(val *ApiResolveWalletAddressByDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResolveWalletAddressByDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResolveWalletAddressByDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResolveWalletAddressByDomainResponse(val *ApiResolveWalletAddressByDomainResponse) *NullableApiResolveWalletAddressByDomainResponse {
	return &NullableApiResolveWalletAddressByDomainResponse{value: val, isSet: true}
}

func (v NullableApiResolveWalletAddressByDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResolveWalletAddressByDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


