# ListAdCampaigns200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaigns** | Pointer to [**[]AdCampaign**](AdCampaign.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewListAdCampaigns200Response

`func NewListAdCampaigns200Response() *ListAdCampaigns200Response`

NewListAdCampaigns200Response instantiates a new ListAdCampaigns200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAdCampaigns200ResponseWithDefaults

`func NewListAdCampaigns200ResponseWithDefaults() *ListAdCampaigns200Response`

NewListAdCampaigns200ResponseWithDefaults instantiates a new ListAdCampaigns200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaigns

`func (o *ListAdCampaigns200Response) GetCampaigns() []AdCampaign`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *ListAdCampaigns200Response) GetCampaignsOk() (*[]AdCampaign, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *ListAdCampaigns200Response) SetCampaigns(v []AdCampaign)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *ListAdCampaigns200Response) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetPagination

`func (o *ListAdCampaigns200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAdCampaigns200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAdCampaigns200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAdCampaigns200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


