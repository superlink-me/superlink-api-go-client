# ApiPurchaseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutId** | Pointer to **string** |  | [optional] 
**Domain** | Pointer to **string** |  | [optional] 
**ExternalUserId** | Pointer to **string** |  | [optional] 
**OwnerAddress** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 
**StripeConnectedAccountId** | Pointer to **string** |  | [optional] 
**WalletAddrs** | Pointer to [**[]ApiAddressRecord**](ApiAddressRecord.md) |  | [optional] 
**Years** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiPurchaseRequest

`func NewApiPurchaseRequest() *ApiPurchaseRequest`

NewApiPurchaseRequest instantiates a new ApiPurchaseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPurchaseRequestWithDefaults

`func NewApiPurchaseRequestWithDefaults() *ApiPurchaseRequest`

NewApiPurchaseRequestWithDefaults instantiates a new ApiPurchaseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutId

`func (o *ApiPurchaseRequest) GetCheckoutId() string`

GetCheckoutId returns the CheckoutId field if non-nil, zero value otherwise.

### GetCheckoutIdOk

`func (o *ApiPurchaseRequest) GetCheckoutIdOk() (*string, bool)`

GetCheckoutIdOk returns a tuple with the CheckoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutId

`func (o *ApiPurchaseRequest) SetCheckoutId(v string)`

SetCheckoutId sets CheckoutId field to given value.

### HasCheckoutId

`func (o *ApiPurchaseRequest) HasCheckoutId() bool`

HasCheckoutId returns a boolean if a field has been set.

### GetDomain

`func (o *ApiPurchaseRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *ApiPurchaseRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *ApiPurchaseRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *ApiPurchaseRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetExternalUserId

`func (o *ApiPurchaseRequest) GetExternalUserId() string`

GetExternalUserId returns the ExternalUserId field if non-nil, zero value otherwise.

### GetExternalUserIdOk

`func (o *ApiPurchaseRequest) GetExternalUserIdOk() (*string, bool)`

GetExternalUserIdOk returns a tuple with the ExternalUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalUserId

`func (o *ApiPurchaseRequest) SetExternalUserId(v string)`

SetExternalUserId sets ExternalUserId field to given value.

### HasExternalUserId

`func (o *ApiPurchaseRequest) HasExternalUserId() bool`

HasExternalUserId returns a boolean if a field has been set.

### GetOwnerAddress

`func (o *ApiPurchaseRequest) GetOwnerAddress() string`

GetOwnerAddress returns the OwnerAddress field if non-nil, zero value otherwise.

### GetOwnerAddressOk

`func (o *ApiPurchaseRequest) GetOwnerAddressOk() (*string, bool)`

GetOwnerAddressOk returns a tuple with the OwnerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerAddress

`func (o *ApiPurchaseRequest) SetOwnerAddress(v string)`

SetOwnerAddress sets OwnerAddress field to given value.

### HasOwnerAddress

`func (o *ApiPurchaseRequest) HasOwnerAddress() bool`

HasOwnerAddress returns a boolean if a field has been set.

### GetPartnerId

`func (o *ApiPurchaseRequest) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ApiPurchaseRequest) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ApiPurchaseRequest) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ApiPurchaseRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetStripeConnectedAccountId

`func (o *ApiPurchaseRequest) GetStripeConnectedAccountId() string`

GetStripeConnectedAccountId returns the StripeConnectedAccountId field if non-nil, zero value otherwise.

### GetStripeConnectedAccountIdOk

`func (o *ApiPurchaseRequest) GetStripeConnectedAccountIdOk() (*string, bool)`

GetStripeConnectedAccountIdOk returns a tuple with the StripeConnectedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStripeConnectedAccountId

`func (o *ApiPurchaseRequest) SetStripeConnectedAccountId(v string)`

SetStripeConnectedAccountId sets StripeConnectedAccountId field to given value.

### HasStripeConnectedAccountId

`func (o *ApiPurchaseRequest) HasStripeConnectedAccountId() bool`

HasStripeConnectedAccountId returns a boolean if a field has been set.

### GetWalletAddrs

`func (o *ApiPurchaseRequest) GetWalletAddrs() []ApiAddressRecord`

GetWalletAddrs returns the WalletAddrs field if non-nil, zero value otherwise.

### GetWalletAddrsOk

`func (o *ApiPurchaseRequest) GetWalletAddrsOk() (*[]ApiAddressRecord, bool)`

GetWalletAddrsOk returns a tuple with the WalletAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletAddrs

`func (o *ApiPurchaseRequest) SetWalletAddrs(v []ApiAddressRecord)`

SetWalletAddrs sets WalletAddrs field to given value.

### HasWalletAddrs

`func (o *ApiPurchaseRequest) HasWalletAddrs() bool`

HasWalletAddrs returns a boolean if a field has been set.

### GetYears

`func (o *ApiPurchaseRequest) GetYears() int32`

GetYears returns the Years field if non-nil, zero value otherwise.

### GetYearsOk

`func (o *ApiPurchaseRequest) GetYearsOk() (*int32, bool)`

GetYearsOk returns a tuple with the Years field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYears

`func (o *ApiPurchaseRequest) SetYears(v int32)`

SetYears sets Years field to given value.

### HasYears

`func (o *ApiPurchaseRequest) HasYears() bool`

HasYears returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


