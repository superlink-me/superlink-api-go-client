# ApiMarketSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MarketListings** | Pointer to [**[]ApiMarketListing**](ApiMarketListing.md) |  | [optional] 

## Methods

### NewApiMarketSearchResponse

`func NewApiMarketSearchResponse() *ApiMarketSearchResponse`

NewApiMarketSearchResponse instantiates a new ApiMarketSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMarketSearchResponseWithDefaults

`func NewApiMarketSearchResponseWithDefaults() *ApiMarketSearchResponse`

NewApiMarketSearchResponseWithDefaults instantiates a new ApiMarketSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMarketListings

`func (o *ApiMarketSearchResponse) GetMarketListings() []ApiMarketListing`

GetMarketListings returns the MarketListings field if non-nil, zero value otherwise.

### GetMarketListingsOk

`func (o *ApiMarketSearchResponse) GetMarketListingsOk() (*[]ApiMarketListing, bool)`

GetMarketListingsOk returns a tuple with the MarketListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketListings

`func (o *ApiMarketSearchResponse) SetMarketListings(v []ApiMarketListing)`

SetMarketListings sets MarketListings field to given value.

### HasMarketListings

`func (o *ApiMarketSearchResponse) HasMarketListings() bool`

HasMarketListings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


