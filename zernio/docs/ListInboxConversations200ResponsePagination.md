# ListInboxConversations200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**NextCursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListInboxConversations200ResponsePagination

`func NewListInboxConversations200ResponsePagination() *ListInboxConversations200ResponsePagination`

NewListInboxConversations200ResponsePagination instantiates a new ListInboxConversations200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversations200ResponsePaginationWithDefaults

`func NewListInboxConversations200ResponsePaginationWithDefaults() *ListInboxConversations200ResponsePagination`

NewListInboxConversations200ResponsePaginationWithDefaults instantiates a new ListInboxConversations200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ListInboxConversations200ResponsePagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListInboxConversations200ResponsePagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListInboxConversations200ResponsePagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListInboxConversations200ResponsePagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetNextCursor

`func (o *ListInboxConversations200ResponsePagination) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListInboxConversations200ResponsePagination) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListInboxConversations200ResponsePagination) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListInboxConversations200ResponsePagination) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListInboxConversations200ResponsePagination) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListInboxConversations200ResponsePagination) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


