# GetContact200ResponseContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**IsBlocked** | Pointer to **bool** |  | [optional] 
**MessagesSentCount** | Pointer to **int32** | Messages sent to the contact, derived live from message history across all linked conversations. | [optional] 
**MessagesReceivedCount** | Pointer to **int32** | Messages received from the contact, derived live from message history across all linked conversations. | [optional] 
**LastMessageSentAt** | Pointer to **NullableTime** | Timestamp of the most recent outgoing message, or null if none. | [optional] 
**LastMessageReceivedAt** | Pointer to **NullableTime** | Timestamp of the most recent incoming message, or null if none. | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ConversationIds** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetContact200ResponseContact

`func NewGetContact200ResponseContact() *GetContact200ResponseContact`

NewGetContact200ResponseContact instantiates a new GetContact200ResponseContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContact200ResponseContactWithDefaults

`func NewGetContact200ResponseContactWithDefaults() *GetContact200ResponseContact`

NewGetContact200ResponseContactWithDefaults instantiates a new GetContact200ResponseContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetContact200ResponseContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetContact200ResponseContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetContact200ResponseContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetContact200ResponseContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetContact200ResponseContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetContact200ResponseContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetContact200ResponseContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetContact200ResponseContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *GetContact200ResponseContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetContact200ResponseContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetContact200ResponseContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetContact200ResponseContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *GetContact200ResponseContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GetContact200ResponseContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GetContact200ResponseContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GetContact200ResponseContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *GetContact200ResponseContact) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *GetContact200ResponseContact) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *GetContact200ResponseContact) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *GetContact200ResponseContact) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetTags

`func (o *GetContact200ResponseContact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetContact200ResponseContact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetContact200ResponseContact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetContact200ResponseContact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *GetContact200ResponseContact) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *GetContact200ResponseContact) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *GetContact200ResponseContact) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *GetContact200ResponseContact) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsBlocked

`func (o *GetContact200ResponseContact) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *GetContact200ResponseContact) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *GetContact200ResponseContact) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *GetContact200ResponseContact) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetMessagesSentCount

`func (o *GetContact200ResponseContact) GetMessagesSentCount() int32`

GetMessagesSentCount returns the MessagesSentCount field if non-nil, zero value otherwise.

### GetMessagesSentCountOk

`func (o *GetContact200ResponseContact) GetMessagesSentCountOk() (*int32, bool)`

GetMessagesSentCountOk returns a tuple with the MessagesSentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentCount

`func (o *GetContact200ResponseContact) SetMessagesSentCount(v int32)`

SetMessagesSentCount sets MessagesSentCount field to given value.

### HasMessagesSentCount

`func (o *GetContact200ResponseContact) HasMessagesSentCount() bool`

HasMessagesSentCount returns a boolean if a field has been set.

### GetMessagesReceivedCount

`func (o *GetContact200ResponseContact) GetMessagesReceivedCount() int32`

GetMessagesReceivedCount returns the MessagesReceivedCount field if non-nil, zero value otherwise.

### GetMessagesReceivedCountOk

`func (o *GetContact200ResponseContact) GetMessagesReceivedCountOk() (*int32, bool)`

GetMessagesReceivedCountOk returns a tuple with the MessagesReceivedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesReceivedCount

`func (o *GetContact200ResponseContact) SetMessagesReceivedCount(v int32)`

SetMessagesReceivedCount sets MessagesReceivedCount field to given value.

### HasMessagesReceivedCount

`func (o *GetContact200ResponseContact) HasMessagesReceivedCount() bool`

HasMessagesReceivedCount returns a boolean if a field has been set.

### GetLastMessageSentAt

`func (o *GetContact200ResponseContact) GetLastMessageSentAt() time.Time`

GetLastMessageSentAt returns the LastMessageSentAt field if non-nil, zero value otherwise.

### GetLastMessageSentAtOk

`func (o *GetContact200ResponseContact) GetLastMessageSentAtOk() (*time.Time, bool)`

GetLastMessageSentAtOk returns a tuple with the LastMessageSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageSentAt

`func (o *GetContact200ResponseContact) SetLastMessageSentAt(v time.Time)`

SetLastMessageSentAt sets LastMessageSentAt field to given value.

### HasLastMessageSentAt

`func (o *GetContact200ResponseContact) HasLastMessageSentAt() bool`

HasLastMessageSentAt returns a boolean if a field has been set.

### SetLastMessageSentAtNil

`func (o *GetContact200ResponseContact) SetLastMessageSentAtNil(b bool)`

 SetLastMessageSentAtNil sets the value for LastMessageSentAt to be an explicit nil

### UnsetLastMessageSentAt
`func (o *GetContact200ResponseContact) UnsetLastMessageSentAt()`

UnsetLastMessageSentAt ensures that no value is present for LastMessageSentAt, not even an explicit nil
### GetLastMessageReceivedAt

`func (o *GetContact200ResponseContact) GetLastMessageReceivedAt() time.Time`

GetLastMessageReceivedAt returns the LastMessageReceivedAt field if non-nil, zero value otherwise.

### GetLastMessageReceivedAtOk

`func (o *GetContact200ResponseContact) GetLastMessageReceivedAtOk() (*time.Time, bool)`

GetLastMessageReceivedAtOk returns a tuple with the LastMessageReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageReceivedAt

`func (o *GetContact200ResponseContact) SetLastMessageReceivedAt(v time.Time)`

SetLastMessageReceivedAt sets LastMessageReceivedAt field to given value.

### HasLastMessageReceivedAt

`func (o *GetContact200ResponseContact) HasLastMessageReceivedAt() bool`

HasLastMessageReceivedAt returns a boolean if a field has been set.

### SetLastMessageReceivedAtNil

`func (o *GetContact200ResponseContact) SetLastMessageReceivedAtNil(b bool)`

 SetLastMessageReceivedAtNil sets the value for LastMessageReceivedAt to be an explicit nil

### UnsetLastMessageReceivedAt
`func (o *GetContact200ResponseContact) UnsetLastMessageReceivedAt()`

UnsetLastMessageReceivedAt ensures that no value is present for LastMessageReceivedAt, not even an explicit nil
### GetCustomFields

`func (o *GetContact200ResponseContact) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *GetContact200ResponseContact) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *GetContact200ResponseContact) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *GetContact200ResponseContact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetNotes

`func (o *GetContact200ResponseContact) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetContact200ResponseContact) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetContact200ResponseContact) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetContact200ResponseContact) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetConversationIds

`func (o *GetContact200ResponseContact) GetConversationIds() []string`

GetConversationIds returns the ConversationIds field if non-nil, zero value otherwise.

### GetConversationIdsOk

`func (o *GetContact200ResponseContact) GetConversationIdsOk() (*[]string, bool)`

GetConversationIdsOk returns a tuple with the ConversationIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationIds

`func (o *GetContact200ResponseContact) SetConversationIds(v []string)`

SetConversationIds sets ConversationIds field to given value.

### HasConversationIds

`func (o *GetContact200ResponseContact) HasConversationIds() bool`

HasConversationIds returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetContact200ResponseContact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetContact200ResponseContact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetContact200ResponseContact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetContact200ResponseContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetContact200ResponseContact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetContact200ResponseContact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetContact200ResponseContact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetContact200ResponseContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


