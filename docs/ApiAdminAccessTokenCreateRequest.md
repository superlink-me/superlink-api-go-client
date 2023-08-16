# ApiAdminAccessTokenCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | Pointer to **string** |  | [optional] 
**PartnerId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DataAccessTokenType**](DataAccessTokenType.md) |  | [optional] 
**ValidFrom** | Pointer to **string** |  | [optional] 
**ValidTill** | Pointer to **string** |  | [optional] 

## Methods

### NewApiAdminAccessTokenCreateRequest

`func NewApiAdminAccessTokenCreateRequest() *ApiAdminAccessTokenCreateRequest`

NewApiAdminAccessTokenCreateRequest instantiates a new ApiAdminAccessTokenCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAdminAccessTokenCreateRequestWithDefaults

`func NewApiAdminAccessTokenCreateRequestWithDefaults() *ApiAdminAccessTokenCreateRequest`

NewApiAdminAccessTokenCreateRequestWithDefaults instantiates a new ApiAdminAccessTokenCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ApiAdminAccessTokenCreateRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiAdminAccessTokenCreateRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiAdminAccessTokenCreateRequest) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApiAdminAccessTokenCreateRequest) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPartnerId

`func (o *ApiAdminAccessTokenCreateRequest) GetPartnerId() string`

GetPartnerId returns the PartnerId field if non-nil, zero value otherwise.

### GetPartnerIdOk

`func (o *ApiAdminAccessTokenCreateRequest) GetPartnerIdOk() (*string, bool)`

GetPartnerIdOk returns a tuple with the PartnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartnerId

`func (o *ApiAdminAccessTokenCreateRequest) SetPartnerId(v string)`

SetPartnerId sets PartnerId field to given value.

### HasPartnerId

`func (o *ApiAdminAccessTokenCreateRequest) HasPartnerId() bool`

HasPartnerId returns a boolean if a field has been set.

### GetType

`func (o *ApiAdminAccessTokenCreateRequest) GetType() DataAccessTokenType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiAdminAccessTokenCreateRequest) GetTypeOk() (*DataAccessTokenType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiAdminAccessTokenCreateRequest) SetType(v DataAccessTokenType)`

SetType sets Type field to given value.

### HasType

`func (o *ApiAdminAccessTokenCreateRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValidFrom

`func (o *ApiAdminAccessTokenCreateRequest) GetValidFrom() string`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *ApiAdminAccessTokenCreateRequest) GetValidFromOk() (*string, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *ApiAdminAccessTokenCreateRequest) SetValidFrom(v string)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *ApiAdminAccessTokenCreateRequest) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTill

`func (o *ApiAdminAccessTokenCreateRequest) GetValidTill() string`

GetValidTill returns the ValidTill field if non-nil, zero value otherwise.

### GetValidTillOk

`func (o *ApiAdminAccessTokenCreateRequest) GetValidTillOk() (*string, bool)`

GetValidTillOk returns a tuple with the ValidTill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTill

`func (o *ApiAdminAccessTokenCreateRequest) SetValidTill(v string)`

SetValidTill sets ValidTill field to given value.

### HasValidTill

`func (o *ApiAdminAccessTokenCreateRequest) HasValidTill() bool`

HasValidTill returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


