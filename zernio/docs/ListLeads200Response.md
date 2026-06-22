# ListLeads200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Leads** | Pointer to [**[]ListLeads200ResponseLeadsInner**](ListLeads200ResponseLeadsInner.md) |  | [optional] 
**Pagination** | Pointer to [**GetInboxPostComments200ResponsePagination**](GetInboxPostComments200ResponsePagination.md) |  | [optional] 

## Methods

### NewListLeads200Response

`func NewListLeads200Response() *ListLeads200Response`

NewListLeads200Response instantiates a new ListLeads200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLeads200ResponseWithDefaults

`func NewListLeads200ResponseWithDefaults() *ListLeads200Response`

NewListLeads200ResponseWithDefaults instantiates a new ListLeads200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListLeads200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLeads200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLeads200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLeads200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLeads

`func (o *ListLeads200Response) GetLeads() []ListLeads200ResponseLeadsInner`

GetLeads returns the Leads field if non-nil, zero value otherwise.

### GetLeadsOk

`func (o *ListLeads200Response) GetLeadsOk() (*[]ListLeads200ResponseLeadsInner, bool)`

GetLeadsOk returns a tuple with the Leads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeads

`func (o *ListLeads200Response) SetLeads(v []ListLeads200ResponseLeadsInner)`

SetLeads sets Leads field to given value.

### HasLeads

`func (o *ListLeads200Response) HasLeads() bool`

HasLeads returns a boolean if a field has been set.

### GetPagination

`func (o *ListLeads200Response) GetPagination() GetInboxPostComments200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListLeads200Response) GetPaginationOk() (*GetInboxPostComments200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListLeads200Response) SetPagination(v GetInboxPostComments200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListLeads200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


