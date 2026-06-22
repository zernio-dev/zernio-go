# ListAds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ads** | Pointer to [**[]Ad**](Ad.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAds200Response

`func NewListAds200Response() *ListAds200Response`

NewListAds200Response instantiates a new ListAds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAds200ResponseWithDefaults

`func NewListAds200ResponseWithDefaults() *ListAds200Response`

NewListAds200ResponseWithDefaults instantiates a new ListAds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAds

`func (o *ListAds200Response) GetAds() []Ad`

GetAds returns the Ads field if non-nil, zero value otherwise.

### GetAdsOk

`func (o *ListAds200Response) GetAdsOk() (*[]Ad, bool)`

GetAdsOk returns a tuple with the Ads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAds

`func (o *ListAds200Response) SetAds(v []Ad)`

SetAds sets Ads field to given value.

### HasAds

`func (o *ListAds200Response) HasAds() bool`

HasAds returns a boolean if a field has been set.

### GetPagination

`func (o *ListAds200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAds200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAds200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAds200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


