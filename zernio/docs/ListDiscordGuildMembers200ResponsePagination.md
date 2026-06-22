# ListDiscordGuildMembers200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NextCursor** | Pointer to **NullableString** | Pass as &#x60;after&#x60; on the next call. Null when there are no more pages. | [optional] 
**HasMore** | Pointer to **bool** |  | [optional] 

## Methods

### NewListDiscordGuildMembers200ResponsePagination

`func NewListDiscordGuildMembers200ResponsePagination() *ListDiscordGuildMembers200ResponsePagination`

NewListDiscordGuildMembers200ResponsePagination instantiates a new ListDiscordGuildMembers200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscordGuildMembers200ResponsePaginationWithDefaults

`func NewListDiscordGuildMembers200ResponsePaginationWithDefaults() *ListDiscordGuildMembers200ResponsePagination`

NewListDiscordGuildMembers200ResponsePaginationWithDefaults instantiates a new ListDiscordGuildMembers200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNextCursor

`func (o *ListDiscordGuildMembers200ResponsePagination) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *ListDiscordGuildMembers200ResponsePagination) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *ListDiscordGuildMembers200ResponsePagination) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *ListDiscordGuildMembers200ResponsePagination) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *ListDiscordGuildMembers200ResponsePagination) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *ListDiscordGuildMembers200ResponsePagination) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil
### GetHasMore

`func (o *ListDiscordGuildMembers200ResponsePagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListDiscordGuildMembers200ResponsePagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListDiscordGuildMembers200ResponsePagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListDiscordGuildMembers200ResponsePagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


