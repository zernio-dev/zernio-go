# GetWhatsAppBlockedUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedUsers** | Pointer to [**[]GetWhatsAppBlockedUsers200ResponseBlockedUsersInner**](GetWhatsAppBlockedUsers200ResponseBlockedUsersInner.md) |  | [optional] 
**NextCursor** | Pointer to **NullableString** | Pass as &#x60;after&#x60; to fetch the next page. Null when there are no more pages. | [optional] 

## Methods

### NewGetWhatsAppBlockedUsers200Response

`func NewGetWhatsAppBlockedUsers200Response() *GetWhatsAppBlockedUsers200Response`

NewGetWhatsAppBlockedUsers200Response instantiates a new GetWhatsAppBlockedUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppBlockedUsers200ResponseWithDefaults

`func NewGetWhatsAppBlockedUsers200ResponseWithDefaults() *GetWhatsAppBlockedUsers200Response`

NewGetWhatsAppBlockedUsers200ResponseWithDefaults instantiates a new GetWhatsAppBlockedUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedUsers

`func (o *GetWhatsAppBlockedUsers200Response) GetBlockedUsers() []GetWhatsAppBlockedUsers200ResponseBlockedUsersInner`

GetBlockedUsers returns the BlockedUsers field if non-nil, zero value otherwise.

### GetBlockedUsersOk

`func (o *GetWhatsAppBlockedUsers200Response) GetBlockedUsersOk() (*[]GetWhatsAppBlockedUsers200ResponseBlockedUsersInner, bool)`

GetBlockedUsersOk returns a tuple with the BlockedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedUsers

`func (o *GetWhatsAppBlockedUsers200Response) SetBlockedUsers(v []GetWhatsAppBlockedUsers200ResponseBlockedUsersInner)`

SetBlockedUsers sets BlockedUsers field to given value.

### HasBlockedUsers

`func (o *GetWhatsAppBlockedUsers200Response) HasBlockedUsers() bool`

HasBlockedUsers returns a boolean if a field has been set.

### GetNextCursor

`func (o *GetWhatsAppBlockedUsers200Response) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *GetWhatsAppBlockedUsers200Response) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *GetWhatsAppBlockedUsers200Response) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.

### HasNextCursor

`func (o *GetWhatsAppBlockedUsers200Response) HasNextCursor() bool`

HasNextCursor returns a boolean if a field has been set.

### SetNextCursorNil

`func (o *GetWhatsAppBlockedUsers200Response) SetNextCursorNil(b bool)`

 SetNextCursorNil sets the value for NextCursor to be an explicit nil

### UnsetNextCursor
`func (o *GetWhatsAppBlockedUsers200Response) UnsetNextCursor()`

UnsetNextCursor ensures that no value is present for NextCursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


