# ApiDNSRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewApiDNSRecord

`func NewApiDNSRecord() *ApiDNSRecord`

NewApiDNSRecord instantiates a new ApiDNSRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiDNSRecordWithDefaults

`func NewApiDNSRecordWithDefaults() *ApiDNSRecord`

NewApiDNSRecordWithDefaults instantiates a new ApiDNSRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiDNSRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiDNSRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiDNSRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiDNSRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTtl

`func (o *ApiDNSRecord) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ApiDNSRecord) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ApiDNSRecord) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ApiDNSRecord) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *ApiDNSRecord) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiDNSRecord) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiDNSRecord) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiDNSRecord) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ApiDNSRecord) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiDNSRecord) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiDNSRecord) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiDNSRecord) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


