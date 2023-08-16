# ApiPurchaseListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Purchases** | Pointer to [**[]ApiPurchaseResponse**](ApiPurchaseResponse.md) |  | [optional] 

## Methods

### NewApiPurchaseListResponse

`func NewApiPurchaseListResponse() *ApiPurchaseListResponse`

NewApiPurchaseListResponse instantiates a new ApiPurchaseListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPurchaseListResponseWithDefaults

`func NewApiPurchaseListResponseWithDefaults() *ApiPurchaseListResponse`

NewApiPurchaseListResponseWithDefaults instantiates a new ApiPurchaseListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPurchases

`func (o *ApiPurchaseListResponse) GetPurchases() []ApiPurchaseResponse`

GetPurchases returns the Purchases field if non-nil, zero value otherwise.

### GetPurchasesOk

`func (o *ApiPurchaseListResponse) GetPurchasesOk() (*[]ApiPurchaseResponse, bool)`

GetPurchasesOk returns a tuple with the Purchases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchases

`func (o *ApiPurchaseListResponse) SetPurchases(v []ApiPurchaseResponse)`

SetPurchases sets Purchases field to given value.

### HasPurchases

`func (o *ApiPurchaseListResponse) HasPurchases() bool`

HasPurchases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


