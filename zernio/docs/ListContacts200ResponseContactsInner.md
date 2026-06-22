# ListContacts200ResponseContactsInner

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
**LastMessageSentAt** | Pointer to **time.Time** |  | [optional] 
**LastMessageReceivedAt** | Pointer to **time.Time** |  | [optional] 
**MessagesSentCount** | Pointer to **int32** |  | [optional] 
**MessagesReceivedCount** | Pointer to **int32** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformIdentifier** | Pointer to **string** |  | [optional] 
**DisplayIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewListContacts200ResponseContactsInner

`func NewListContacts200ResponseContactsInner() *ListContacts200ResponseContactsInner`

NewListContacts200ResponseContactsInner instantiates a new ListContacts200ResponseContactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContacts200ResponseContactsInnerWithDefaults

`func NewListContacts200ResponseContactsInnerWithDefaults() *ListContacts200ResponseContactsInner`

NewListContacts200ResponseContactsInnerWithDefaults instantiates a new ListContacts200ResponseContactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListContacts200ResponseContactsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListContacts200ResponseContactsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListContacts200ResponseContactsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListContacts200ResponseContactsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListContacts200ResponseContactsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListContacts200ResponseContactsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListContacts200ResponseContactsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListContacts200ResponseContactsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *ListContacts200ResponseContactsInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListContacts200ResponseContactsInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListContacts200ResponseContactsInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListContacts200ResponseContactsInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *ListContacts200ResponseContactsInner) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *ListContacts200ResponseContactsInner) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *ListContacts200ResponseContactsInner) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *ListContacts200ResponseContactsInner) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *ListContacts200ResponseContactsInner) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *ListContacts200ResponseContactsInner) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *ListContacts200ResponseContactsInner) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *ListContacts200ResponseContactsInner) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetTags

`func (o *ListContacts200ResponseContactsInner) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListContacts200ResponseContactsInner) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListContacts200ResponseContactsInner) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListContacts200ResponseContactsInner) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *ListContacts200ResponseContactsInner) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *ListContacts200ResponseContactsInner) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *ListContacts200ResponseContactsInner) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *ListContacts200ResponseContactsInner) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsBlocked

`func (o *ListContacts200ResponseContactsInner) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *ListContacts200ResponseContactsInner) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *ListContacts200ResponseContactsInner) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *ListContacts200ResponseContactsInner) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetLastMessageSentAt

`func (o *ListContacts200ResponseContactsInner) GetLastMessageSentAt() time.Time`

GetLastMessageSentAt returns the LastMessageSentAt field if non-nil, zero value otherwise.

### GetLastMessageSentAtOk

`func (o *ListContacts200ResponseContactsInner) GetLastMessageSentAtOk() (*time.Time, bool)`

GetLastMessageSentAtOk returns a tuple with the LastMessageSentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageSentAt

`func (o *ListContacts200ResponseContactsInner) SetLastMessageSentAt(v time.Time)`

SetLastMessageSentAt sets LastMessageSentAt field to given value.

### HasLastMessageSentAt

`func (o *ListContacts200ResponseContactsInner) HasLastMessageSentAt() bool`

HasLastMessageSentAt returns a boolean if a field has been set.

### GetLastMessageReceivedAt

`func (o *ListContacts200ResponseContactsInner) GetLastMessageReceivedAt() time.Time`

GetLastMessageReceivedAt returns the LastMessageReceivedAt field if non-nil, zero value otherwise.

### GetLastMessageReceivedAtOk

`func (o *ListContacts200ResponseContactsInner) GetLastMessageReceivedAtOk() (*time.Time, bool)`

GetLastMessageReceivedAtOk returns a tuple with the LastMessageReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageReceivedAt

`func (o *ListContacts200ResponseContactsInner) SetLastMessageReceivedAt(v time.Time)`

SetLastMessageReceivedAt sets LastMessageReceivedAt field to given value.

### HasLastMessageReceivedAt

`func (o *ListContacts200ResponseContactsInner) HasLastMessageReceivedAt() bool`

HasLastMessageReceivedAt returns a boolean if a field has been set.

### GetMessagesSentCount

`func (o *ListContacts200ResponseContactsInner) GetMessagesSentCount() int32`

GetMessagesSentCount returns the MessagesSentCount field if non-nil, zero value otherwise.

### GetMessagesSentCountOk

`func (o *ListContacts200ResponseContactsInner) GetMessagesSentCountOk() (*int32, bool)`

GetMessagesSentCountOk returns a tuple with the MessagesSentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesSentCount

`func (o *ListContacts200ResponseContactsInner) SetMessagesSentCount(v int32)`

SetMessagesSentCount sets MessagesSentCount field to given value.

### HasMessagesSentCount

`func (o *ListContacts200ResponseContactsInner) HasMessagesSentCount() bool`

HasMessagesSentCount returns a boolean if a field has been set.

### GetMessagesReceivedCount

`func (o *ListContacts200ResponseContactsInner) GetMessagesReceivedCount() int32`

GetMessagesReceivedCount returns the MessagesReceivedCount field if non-nil, zero value otherwise.

### GetMessagesReceivedCountOk

`func (o *ListContacts200ResponseContactsInner) GetMessagesReceivedCountOk() (*int32, bool)`

GetMessagesReceivedCountOk returns a tuple with the MessagesReceivedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesReceivedCount

`func (o *ListContacts200ResponseContactsInner) SetMessagesReceivedCount(v int32)`

SetMessagesReceivedCount sets MessagesReceivedCount field to given value.

### HasMessagesReceivedCount

`func (o *ListContacts200ResponseContactsInner) HasMessagesReceivedCount() bool`

HasMessagesReceivedCount returns a boolean if a field has been set.

### GetCustomFields

`func (o *ListContacts200ResponseContactsInner) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *ListContacts200ResponseContactsInner) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *ListContacts200ResponseContactsInner) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *ListContacts200ResponseContactsInner) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetNotes

`func (o *ListContacts200ResponseContactsInner) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListContacts200ResponseContactsInner) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListContacts200ResponseContactsInner) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListContacts200ResponseContactsInner) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListContacts200ResponseContactsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListContacts200ResponseContactsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListContacts200ResponseContactsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListContacts200ResponseContactsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPlatform

`func (o *ListContacts200ResponseContactsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListContacts200ResponseContactsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListContacts200ResponseContactsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListContacts200ResponseContactsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformIdentifier

`func (o *ListContacts200ResponseContactsInner) GetPlatformIdentifier() string`

GetPlatformIdentifier returns the PlatformIdentifier field if non-nil, zero value otherwise.

### GetPlatformIdentifierOk

`func (o *ListContacts200ResponseContactsInner) GetPlatformIdentifierOk() (*string, bool)`

GetPlatformIdentifierOk returns a tuple with the PlatformIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentifier

`func (o *ListContacts200ResponseContactsInner) SetPlatformIdentifier(v string)`

SetPlatformIdentifier sets PlatformIdentifier field to given value.

### HasPlatformIdentifier

`func (o *ListContacts200ResponseContactsInner) HasPlatformIdentifier() bool`

HasPlatformIdentifier returns a boolean if a field has been set.

### GetDisplayIdentifier

`func (o *ListContacts200ResponseContactsInner) GetDisplayIdentifier() string`

GetDisplayIdentifier returns the DisplayIdentifier field if non-nil, zero value otherwise.

### GetDisplayIdentifierOk

`func (o *ListContacts200ResponseContactsInner) GetDisplayIdentifierOk() (*string, bool)`

GetDisplayIdentifierOk returns a tuple with the DisplayIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIdentifier

`func (o *ListContacts200ResponseContactsInner) SetDisplayIdentifier(v string)`

SetDisplayIdentifier sets DisplayIdentifier field to given value.

### HasDisplayIdentifier

`func (o *ListContacts200ResponseContactsInner) HasDisplayIdentifier() bool`

HasDisplayIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


