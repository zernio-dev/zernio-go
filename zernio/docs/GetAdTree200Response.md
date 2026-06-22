# GetAdTree200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Campaigns** | Pointer to [**[]AdTreeCampaign**](AdTreeCampaign.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 

## Methods

### NewGetAdTree200Response

`func NewGetAdTree200Response() *GetAdTree200Response`

NewGetAdTree200Response instantiates a new GetAdTree200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdTree200ResponseWithDefaults

`func NewGetAdTree200ResponseWithDefaults() *GetAdTree200Response`

NewGetAdTree200ResponseWithDefaults instantiates a new GetAdTree200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaigns

`func (o *GetAdTree200Response) GetCampaigns() []AdTreeCampaign`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *GetAdTree200Response) GetCampaignsOk() (*[]AdTreeCampaign, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *GetAdTree200Response) SetCampaigns(v []AdTreeCampaign)`

SetCampaigns sets Campaigns field to given value.

### HasCampaigns

`func (o *GetAdTree200Response) HasCampaigns() bool`

HasCampaigns returns a boolean if a field has been set.

### GetPagination

`func (o *GetAdTree200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAdTree200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAdTree200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAdTree200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


