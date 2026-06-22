# ListFormLeads200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Leads** | Pointer to [**[]ListFormLeads200ResponseLeadsInner**](ListFormLeads200ResponseLeadsInner.md) |  | [optional] 
**Pagination** | Pointer to [**GetInboxPostComments200ResponsePagination**](GetInboxPostComments200ResponsePagination.md) |  | [optional] 

## Methods

### NewListFormLeads200Response

`func NewListFormLeads200Response() *ListFormLeads200Response`

NewListFormLeads200Response instantiates a new ListFormLeads200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListFormLeads200ResponseWithDefaults

`func NewListFormLeads200ResponseWithDefaults() *ListFormLeads200Response`

NewListFormLeads200ResponseWithDefaults instantiates a new ListFormLeads200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListFormLeads200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListFormLeads200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListFormLeads200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListFormLeads200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLeads

`func (o *ListFormLeads200Response) GetLeads() []ListFormLeads200ResponseLeadsInner`

GetLeads returns the Leads field if non-nil, zero value otherwise.

### GetLeadsOk

`func (o *ListFormLeads200Response) GetLeadsOk() (*[]ListFormLeads200ResponseLeadsInner, bool)`

GetLeadsOk returns a tuple with the Leads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeads

`func (o *ListFormLeads200Response) SetLeads(v []ListFormLeads200ResponseLeadsInner)`

SetLeads sets Leads field to given value.

### HasLeads

`func (o *ListFormLeads200Response) HasLeads() bool`

HasLeads returns a boolean if a field has been set.

### GetPagination

`func (o *ListFormLeads200Response) GetPagination() GetInboxPostComments200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListFormLeads200Response) GetPaginationOk() (*GetInboxPostComments200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListFormLeads200Response) SetPagination(v GetInboxPostComments200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListFormLeads200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


