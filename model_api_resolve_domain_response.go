/*
Superlink

API for Superlink

API version: v0.3.3
Contact: support@superlink.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package superlink

import (
	"encoding/json"
)

// checks if the ApiResolveDomainResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiResolveDomainResponse{}

// ApiResolveDomainResponse struct for ApiResolveDomainResponse
type ApiResolveDomainResponse struct {
	ContentHash *string `json:"contentHash,omitempty"`
	DnsRecords []ApiDNSRecord `json:"dnsRecords,omitempty"`
	Domain *string `json:"domain,omitempty"`
	NameService *ApiNameService `json:"nameService,omitempty"`
	OwnerAddress *string `json:"ownerAddress,omitempty"`
	TxtRecords []ApiTXTRecord `json:"txtRecords,omitempty"`
	Wallets []ApiWalletData `json:"wallets,omitempty"`
}

// NewApiResolveDomainResponse instantiates a new ApiResolveDomainResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResolveDomainResponse() *ApiResolveDomainResponse {
	this := ApiResolveDomainResponse{}
	return &this
}

// NewApiResolveDomainResponseWithDefaults instantiates a new ApiResolveDomainResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResolveDomainResponseWithDefaults() *ApiResolveDomainResponse {
	this := ApiResolveDomainResponse{}
	return &this
}

// GetContentHash returns the ContentHash field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetContentHash() string {
	if o == nil || IsNil(o.ContentHash) {
		var ret string
		return ret
	}
	return *o.ContentHash
}

// GetContentHashOk returns a tuple with the ContentHash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetContentHashOk() (*string, bool) {
	if o == nil || IsNil(o.ContentHash) {
		return nil, false
	}
	return o.ContentHash, true
}

// HasContentHash returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasContentHash() bool {
	if o != nil && !IsNil(o.ContentHash) {
		return true
	}

	return false
}

// SetContentHash gets a reference to the given string and assigns it to the ContentHash field.
func (o *ApiResolveDomainResponse) SetContentHash(v string) {
	o.ContentHash = &v
}

// GetDnsRecords returns the DnsRecords field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetDnsRecords() []ApiDNSRecord {
	if o == nil || IsNil(o.DnsRecords) {
		var ret []ApiDNSRecord
		return ret
	}
	return o.DnsRecords
}

// GetDnsRecordsOk returns a tuple with the DnsRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetDnsRecordsOk() ([]ApiDNSRecord, bool) {
	if o == nil || IsNil(o.DnsRecords) {
		return nil, false
	}
	return o.DnsRecords, true
}

// HasDnsRecords returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasDnsRecords() bool {
	if o != nil && !IsNil(o.DnsRecords) {
		return true
	}

	return false
}

// SetDnsRecords gets a reference to the given []ApiDNSRecord and assigns it to the DnsRecords field.
func (o *ApiResolveDomainResponse) SetDnsRecords(v []ApiDNSRecord) {
	o.DnsRecords = v
}

// GetDomain returns the Domain field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetDomain() string {
	if o == nil || IsNil(o.Domain) {
		var ret string
		return ret
	}
	return *o.Domain
}

// GetDomainOk returns a tuple with the Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetDomainOk() (*string, bool) {
	if o == nil || IsNil(o.Domain) {
		return nil, false
	}
	return o.Domain, true
}

// HasDomain returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasDomain() bool {
	if o != nil && !IsNil(o.Domain) {
		return true
	}

	return false
}

// SetDomain gets a reference to the given string and assigns it to the Domain field.
func (o *ApiResolveDomainResponse) SetDomain(v string) {
	o.Domain = &v
}

// GetNameService returns the NameService field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetNameService() ApiNameService {
	if o == nil || IsNil(o.NameService) {
		var ret ApiNameService
		return ret
	}
	return *o.NameService
}

// GetNameServiceOk returns a tuple with the NameService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetNameServiceOk() (*ApiNameService, bool) {
	if o == nil || IsNil(o.NameService) {
		return nil, false
	}
	return o.NameService, true
}

// HasNameService returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasNameService() bool {
	if o != nil && !IsNil(o.NameService) {
		return true
	}

	return false
}

// SetNameService gets a reference to the given ApiNameService and assigns it to the NameService field.
func (o *ApiResolveDomainResponse) SetNameService(v ApiNameService) {
	o.NameService = &v
}

// GetOwnerAddress returns the OwnerAddress field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetOwnerAddress() string {
	if o == nil || IsNil(o.OwnerAddress) {
		var ret string
		return ret
	}
	return *o.OwnerAddress
}

// GetOwnerAddressOk returns a tuple with the OwnerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetOwnerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerAddress) {
		return nil, false
	}
	return o.OwnerAddress, true
}

// HasOwnerAddress returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasOwnerAddress() bool {
	if o != nil && !IsNil(o.OwnerAddress) {
		return true
	}

	return false
}

// SetOwnerAddress gets a reference to the given string and assigns it to the OwnerAddress field.
func (o *ApiResolveDomainResponse) SetOwnerAddress(v string) {
	o.OwnerAddress = &v
}

// GetTxtRecords returns the TxtRecords field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetTxtRecords() []ApiTXTRecord {
	if o == nil || IsNil(o.TxtRecords) {
		var ret []ApiTXTRecord
		return ret
	}
	return o.TxtRecords
}

// GetTxtRecordsOk returns a tuple with the TxtRecords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetTxtRecordsOk() ([]ApiTXTRecord, bool) {
	if o == nil || IsNil(o.TxtRecords) {
		return nil, false
	}
	return o.TxtRecords, true
}

// HasTxtRecords returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasTxtRecords() bool {
	if o != nil && !IsNil(o.TxtRecords) {
		return true
	}

	return false
}

// SetTxtRecords gets a reference to the given []ApiTXTRecord and assigns it to the TxtRecords field.
func (o *ApiResolveDomainResponse) SetTxtRecords(v []ApiTXTRecord) {
	o.TxtRecords = v
}

// GetWallets returns the Wallets field value if set, zero value otherwise.
func (o *ApiResolveDomainResponse) GetWallets() []ApiWalletData {
	if o == nil || IsNil(o.Wallets) {
		var ret []ApiWalletData
		return ret
	}
	return o.Wallets
}

// GetWalletsOk returns a tuple with the Wallets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResolveDomainResponse) GetWalletsOk() ([]ApiWalletData, bool) {
	if o == nil || IsNil(o.Wallets) {
		return nil, false
	}
	return o.Wallets, true
}

// HasWallets returns a boolean if a field has been set.
func (o *ApiResolveDomainResponse) HasWallets() bool {
	if o != nil && !IsNil(o.Wallets) {
		return true
	}

	return false
}

// SetWallets gets a reference to the given []ApiWalletData and assigns it to the Wallets field.
func (o *ApiResolveDomainResponse) SetWallets(v []ApiWalletData) {
	o.Wallets = v
}

func (o ApiResolveDomainResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiResolveDomainResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ContentHash) {
		toSerialize["contentHash"] = o.ContentHash
	}
	if !IsNil(o.DnsRecords) {
		toSerialize["dnsRecords"] = o.DnsRecords
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
	if !IsNil(o.TxtRecords) {
		toSerialize["txtRecords"] = o.TxtRecords
	}
	if !IsNil(o.Wallets) {
		toSerialize["wallets"] = o.Wallets
	}
	return toSerialize, nil
}

type NullableApiResolveDomainResponse struct {
	value *ApiResolveDomainResponse
	isSet bool
}

func (v NullableApiResolveDomainResponse) Get() *ApiResolveDomainResponse {
	return v.value
}

func (v *NullableApiResolveDomainResponse) Set(val *ApiResolveDomainResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResolveDomainResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResolveDomainResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResolveDomainResponse(val *ApiResolveDomainResponse) *NullableApiResolveDomainResponse {
	return &NullableApiResolveDomainResponse{value: val, isSet: true}
}

func (v NullableApiResolveDomainResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResolveDomainResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


