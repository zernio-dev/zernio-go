# GetInboxPostComments200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Cursor** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetInboxPostComments200ResponsePagination

`func NewGetInboxPostComments200ResponsePagination() *GetInboxPostComments200ResponsePagination`

NewGetInboxPostComments200ResponsePagination instantiates a new GetInboxPostComments200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxPostComments200ResponsePaginationWithDefaults

`func NewGetInboxPostComments200ResponsePaginationWithDefaults() *GetInboxPostComments200ResponsePagination`

NewGetInboxPostComments200ResponsePaginationWithDefaults instantiates a new GetInboxPostComments200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetInboxPostComments200ResponsePagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetInboxPostComments200ResponsePagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetInboxPostComments200ResponsePagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetInboxPostComments200ResponsePagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetCursor

`func (o *GetInboxPostComments200ResponsePagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetInboxPostComments200ResponsePagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetInboxPostComments200ResponsePagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetInboxPostComments200ResponsePagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *GetInboxPostComments200ResponsePagination) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *GetInboxPostComments200ResponsePagination) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


