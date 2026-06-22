# GetInboxConversationMessages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Pagination** | Pointer to [**ListInboxConversations200ResponsePagination**](ListInboxConversations200ResponsePagination.md) |  | [optional] 
**SortOrderApplied** | Pointer to **string** | Sort order actually applied to the returned page. May differ from the requested &#x60;sortOrder&#x60; for Facebook and Bluesky (always &#x60;desc&#x60; regardless of request).  | [optional] 
**Messages** | Pointer to [**[]GetInboxConversationMessages200ResponseMessagesInner**](GetInboxConversationMessages200ResponseMessagesInner.md) |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetInboxConversationMessages200Response

`func NewGetInboxConversationMessages200Response() *GetInboxConversationMessages200Response`

NewGetInboxConversationMessages200Response instantiates a new GetInboxConversationMessages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversationMessages200ResponseWithDefaults

`func NewGetInboxConversationMessages200ResponseWithDefaults() *GetInboxConversationMessages200Response`

NewGetInboxConversationMessages200ResponseWithDefaults instantiates a new GetInboxConversationMessages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetInboxConversationMessages200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInboxConversationMessages200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInboxConversationMessages200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInboxConversationMessages200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPagination

`func (o *GetInboxConversationMessages200Response) GetPagination() ListInboxConversations200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetInboxConversationMessages200Response) GetPaginationOk() (*ListInboxConversations200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetInboxConversationMessages200Response) SetPagination(v ListInboxConversations200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetInboxConversationMessages200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetSortOrderApplied

`func (o *GetInboxConversationMessages200Response) GetSortOrderApplied() string`

GetSortOrderApplied returns the SortOrderApplied field if non-nil, zero value otherwise.

### GetSortOrderAppliedOk

`func (o *GetInboxConversationMessages200Response) GetSortOrderAppliedOk() (*string, bool)`

GetSortOrderAppliedOk returns a tuple with the SortOrderApplied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrderApplied

`func (o *GetInboxConversationMessages200Response) SetSortOrderApplied(v string)`

SetSortOrderApplied sets SortOrderApplied field to given value.

### HasSortOrderApplied

`func (o *GetInboxConversationMessages200Response) HasSortOrderApplied() bool`

HasSortOrderApplied returns a boolean if a field has been set.

### GetMessages

`func (o *GetInboxConversationMessages200Response) GetMessages() []GetInboxConversationMessages200ResponseMessagesInner`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetInboxConversationMessages200Response) GetMessagesOk() (*[]GetInboxConversationMessages200ResponseMessagesInner, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetInboxConversationMessages200Response) SetMessages(v []GetInboxConversationMessages200ResponseMessagesInner)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *GetInboxConversationMessages200Response) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInboxConversationMessages200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInboxConversationMessages200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInboxConversationMessages200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInboxConversationMessages200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


