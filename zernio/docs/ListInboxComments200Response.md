# ListInboxComments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ListInboxComments200ResponseDataInner**](ListInboxComments200ResponseDataInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListInboxConversations200ResponsePagination**](ListInboxConversations200ResponsePagination.md) |  | [optional] 
**Meta** | Pointer to [**ListInboxConversations200ResponseMeta**](ListInboxConversations200ResponseMeta.md) |  | [optional] 

## Methods

### NewListInboxComments200Response

`func NewListInboxComments200Response() *ListInboxComments200Response`

NewListInboxComments200Response instantiates a new ListInboxComments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxComments200ResponseWithDefaults

`func NewListInboxComments200ResponseWithDefaults() *ListInboxComments200Response`

NewListInboxComments200ResponseWithDefaults instantiates a new ListInboxComments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListInboxComments200Response) GetData() []ListInboxComments200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListInboxComments200Response) GetDataOk() (*[]ListInboxComments200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListInboxComments200Response) SetData(v []ListInboxComments200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListInboxComments200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListInboxComments200Response) GetPagination() ListInboxConversations200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInboxComments200Response) GetPaginationOk() (*ListInboxConversations200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInboxComments200Response) SetPagination(v ListInboxConversations200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListInboxComments200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetMeta

`func (o *ListInboxComments200Response) GetMeta() ListInboxConversations200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInboxComments200Response) GetMetaOk() (*ListInboxConversations200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInboxComments200Response) SetMeta(v ListInboxConversations200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInboxComments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


