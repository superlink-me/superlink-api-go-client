# ApiAddressRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Addr** | Pointer to **string** |  | [optional] 
**CoinId** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiAddressRecord

`func NewApiAddressRecord() *ApiAddressRecord`

NewApiAddressRecord instantiates a new ApiAddressRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAddressRecordWithDefaults

`func NewApiAddressRecordWithDefaults() *ApiAddressRecord`

NewApiAddressRecordWithDefaults instantiates a new ApiAddressRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddr

`func (o *ApiAddressRecord) GetAddr() string`

GetAddr returns the Addr field if non-nil, zero value otherwise.

### GetAddrOk

`func (o *ApiAddressRecord) GetAddrOk() (*string, bool)`

GetAddrOk returns a tuple with the Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddr

`func (o *ApiAddressRecord) SetAddr(v string)`

SetAddr sets Addr field to given value.

### HasAddr

`func (o *ApiAddressRecord) HasAddr() bool`

HasAddr returns a boolean if a field has been set.

### GetCoinId

`func (o *ApiAddressRecord) GetCoinId() int32`

GetCoinId returns the CoinId field if non-nil, zero value otherwise.

### GetCoinIdOk

`func (o *ApiAddressRecord) GetCoinIdOk() (*int32, bool)`

GetCoinIdOk returns a tuple with the CoinId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoinId

`func (o *ApiAddressRecord) SetCoinId(v int32)`

SetCoinId sets CoinId field to given value.

### HasCoinId

`func (o *ApiAddressRecord) HasCoinId() bool`

HasCoinId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


