# GetLinkedInPostReactions200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** |  | [optional] 
**Cursor** | Pointer to **string** | Offset for next page | [optional] 
**Total** | Pointer to **int32** | Total number of reactions (when available) | [optional] 

## Methods

### NewGetLinkedInPostReactions200ResponsePagination

`func NewGetLinkedInPostReactions200ResponsePagination() *GetLinkedInPostReactions200ResponsePagination`

NewGetLinkedInPostReactions200ResponsePagination instantiates a new GetLinkedInPostReactions200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkedInPostReactions200ResponsePaginationWithDefaults

`func NewGetLinkedInPostReactions200ResponsePaginationWithDefaults() *GetLinkedInPostReactions200ResponsePagination`

NewGetLinkedInPostReactions200ResponsePaginationWithDefaults instantiates a new GetLinkedInPostReactions200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *GetLinkedInPostReactions200ResponsePagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetLinkedInPostReactions200ResponsePagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetLinkedInPostReactions200ResponsePagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetLinkedInPostReactions200ResponsePagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetCursor

`func (o *GetLinkedInPostReactions200ResponsePagination) GetCursor() string`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *GetLinkedInPostReactions200ResponsePagination) GetCursorOk() (*string, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *GetLinkedInPostReactions200ResponsePagination) SetCursor(v string)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *GetLinkedInPostReactions200ResponsePagination) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### GetTotal

`func (o *GetLinkedInPostReactions200ResponsePagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetLinkedInPostReactions200ResponsePagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetLinkedInPostReactions200ResponsePagination) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetLinkedInPostReactions200ResponsePagination) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


