# GetContact200ResponseChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformIdentifier** | Pointer to **string** |  | [optional] 
**DisplayIdentifier** | Pointer to **string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**LastActiveAt** | Pointer to **NullableTime** | Most recent message (either direction) in this channel&#39;s conversation, or null if none. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetContact200ResponseChannelsInner

`func NewGetContact200ResponseChannelsInner() *GetContact200ResponseChannelsInner`

NewGetContact200ResponseChannelsInner instantiates a new GetContact200ResponseChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContact200ResponseChannelsInnerWithDefaults

`func NewGetContact200ResponseChannelsInnerWithDefaults() *GetContact200ResponseChannelsInner`

NewGetContact200ResponseChannelsInnerWithDefaults instantiates a new GetContact200ResponseChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetContact200ResponseChannelsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetContact200ResponseChannelsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetContact200ResponseChannelsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetContact200ResponseChannelsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetContact200ResponseChannelsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetContact200ResponseChannelsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetContact200ResponseChannelsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetContact200ResponseChannelsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetContact200ResponseChannelsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetContact200ResponseChannelsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetContact200ResponseChannelsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetContact200ResponseChannelsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformIdentifier

`func (o *GetContact200ResponseChannelsInner) GetPlatformIdentifier() string`

GetPlatformIdentifier returns the PlatformIdentifier field if non-nil, zero value otherwise.

### GetPlatformIdentifierOk

`func (o *GetContact200ResponseChannelsInner) GetPlatformIdentifierOk() (*string, bool)`

GetPlatformIdentifierOk returns a tuple with the PlatformIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentifier

`func (o *GetContact200ResponseChannelsInner) SetPlatformIdentifier(v string)`

SetPlatformIdentifier sets PlatformIdentifier field to given value.

### HasPlatformIdentifier

`func (o *GetContact200ResponseChannelsInner) HasPlatformIdentifier() bool`

HasPlatformIdentifier returns a boolean if a field has been set.

### GetDisplayIdentifier

`func (o *GetContact200ResponseChannelsInner) GetDisplayIdentifier() string`

GetDisplayIdentifier returns the DisplayIdentifier field if non-nil, zero value otherwise.

### GetDisplayIdentifierOk

`func (o *GetContact200ResponseChannelsInner) GetDisplayIdentifierOk() (*string, bool)`

GetDisplayIdentifierOk returns a tuple with the DisplayIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIdentifier

`func (o *GetContact200ResponseChannelsInner) SetDisplayIdentifier(v string)`

SetDisplayIdentifier sets DisplayIdentifier field to given value.

### HasDisplayIdentifier

`func (o *GetContact200ResponseChannelsInner) HasDisplayIdentifier() bool`

HasDisplayIdentifier returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *GetContact200ResponseChannelsInner) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *GetContact200ResponseChannelsInner) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *GetContact200ResponseChannelsInner) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *GetContact200ResponseChannelsInner) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetConversationId

`func (o *GetContact200ResponseChannelsInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetContact200ResponseChannelsInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetContact200ResponseChannelsInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *GetContact200ResponseChannelsInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetLastActiveAt

`func (o *GetContact200ResponseChannelsInner) GetLastActiveAt() time.Time`

GetLastActiveAt returns the LastActiveAt field if non-nil, zero value otherwise.

### GetLastActiveAtOk

`func (o *GetContact200ResponseChannelsInner) GetLastActiveAtOk() (*time.Time, bool)`

GetLastActiveAtOk returns a tuple with the LastActiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveAt

`func (o *GetContact200ResponseChannelsInner) SetLastActiveAt(v time.Time)`

SetLastActiveAt sets LastActiveAt field to given value.

### HasLastActiveAt

`func (o *GetContact200ResponseChannelsInner) HasLastActiveAt() bool`

HasLastActiveAt returns a boolean if a field has been set.

### SetLastActiveAtNil

`func (o *GetContact200ResponseChannelsInner) SetLastActiveAtNil(b bool)`

 SetLastActiveAtNil sets the value for LastActiveAt to be an explicit nil

### UnsetLastActiveAt
`func (o *GetContact200ResponseChannelsInner) UnsetLastActiveAt()`

UnsetLastActiveAt ensures that no value is present for LastActiveAt, not even an explicit nil
### GetCreatedAt

`func (o *GetContact200ResponseChannelsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetContact200ResponseChannelsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetContact200ResponseChannelsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetContact200ResponseChannelsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


