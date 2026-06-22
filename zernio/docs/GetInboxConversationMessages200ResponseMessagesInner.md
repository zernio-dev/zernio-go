# GetInboxConversationMessages200ResponseMessagesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**SenderId** | Pointer to **string** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**SenderVerifiedType** | Pointer to **NullableString** | X/Twitter verified badge type. Only present for Twitter/X messages. | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Attachments** | Pointer to [**[]GetInboxConversationMessages200ResponseMessagesInnerAttachmentsInner**](GetInboxConversationMessages200ResponseMessagesInnerAttachmentsInner.md) |  | [optional] 
**Subject** | Pointer to **NullableString** | Reddit message subject | [optional] 
**StoryReply** | Pointer to **NullableBool** | Instagram story reply | [optional] 
**IsStoryMention** | Pointer to **NullableBool** | Instagram story mention | [optional] 
**IsEdited** | Pointer to **bool** | True if the sender has edited this message at least once. | [optional] 
**EditedAt** | Pointer to **NullableTime** | When the most recent edit happened. | [optional] 
**EditCount** | Pointer to **int32** | Total number of edits applied. | [optional] 
**EditHistory** | Pointer to [**[]GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInner**](GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInner.md) | Every prior version of the message, oldest first. | [optional] 
**IsDeleted** | Pointer to **bool** | True if the sender has deleted (unsent) this message. The original message and attachments fields remain populated. | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**DeliveryStatus** | Pointer to **NullableString** | Lifecycle status for outgoing messages. Not all platforms emit every state (see webhook support matrix). | [optional] 
**DeliveredAt** | Pointer to **NullableTime** |  | [optional] 
**ReadAt** | Pointer to **NullableTime** |  | [optional] 
**SentAt** | Pointer to **NullableTime** | Original send time for outgoing messages (used for Messenger watermark queries). | [optional] 
**DeliveryError** | Pointer to [**GetInboxConversationMessages200ResponseMessagesInnerDeliveryError**](GetInboxConversationMessages200ResponseMessagesInnerDeliveryError.md) |  | [optional] 
**Reactions** | Pointer to [**[]GetInboxConversationMessages200ResponseMessagesInnerReactionsInner**](GetInboxConversationMessages200ResponseMessagesInnerReactionsInner.md) | Emoji reactions on this message (WhatsApp / Telegram). At most one per party in a 1:1 thread. | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Platform-specific extras. Free-form, but commonly includes: &#x60;quotedMessageId&#x60; (platformMessageId this message replies to), &#x60;waInteractive&#x60; (a compact descriptor of WhatsApp interactive content sent: buttons / list / cta_url / flow / location_request), and for inbound interactive taps &#x60;interactiveType&#x60; / &#x60;interactiveId&#x60;.  | [optional] 

## Methods

### NewGetInboxConversationMessages200ResponseMessagesInner

`func NewGetInboxConversationMessages200ResponseMessagesInner() *GetInboxConversationMessages200ResponseMessagesInner`

NewGetInboxConversationMessages200ResponseMessagesInner instantiates a new GetInboxConversationMessages200ResponseMessagesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversationMessages200ResponseMessagesInnerWithDefaults

`func NewGetInboxConversationMessages200ResponseMessagesInnerWithDefaults() *GetInboxConversationMessages200ResponseMessagesInner`

NewGetInboxConversationMessages200ResponseMessagesInnerWithDefaults instantiates a new GetInboxConversationMessages200ResponseMessagesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetConversationId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetMessage

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSenderId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSenderId() string`

GetSenderId returns the SenderId field if non-nil, zero value otherwise.

### GetSenderIdOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSenderIdOk() (*string, bool)`

GetSenderIdOk returns a tuple with the SenderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSenderId(v string)`

SetSenderId sets SenderId field to given value.

### HasSenderId

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasSenderId() bool`

HasSenderId returns a boolean if a field has been set.

### GetSenderName

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetSenderVerifiedType

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSenderVerifiedType() string`

GetSenderVerifiedType returns the SenderVerifiedType field if non-nil, zero value otherwise.

### GetSenderVerifiedTypeOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSenderVerifiedTypeOk() (*string, bool)`

GetSenderVerifiedTypeOk returns a tuple with the SenderVerifiedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderVerifiedType

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSenderVerifiedType(v string)`

SetSenderVerifiedType sets SenderVerifiedType field to given value.

### HasSenderVerifiedType

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasSenderVerifiedType() bool`

HasSenderVerifiedType returns a boolean if a field has been set.

### SetSenderVerifiedTypeNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSenderVerifiedTypeNil(b bool)`

 SetSenderVerifiedTypeNil sets the value for SenderVerifiedType to be an explicit nil

### UnsetSenderVerifiedType
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetSenderVerifiedType()`

UnsetSenderVerifiedType ensures that no value is present for SenderVerifiedType, not even an explicit nil
### GetDirection

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetAttachments

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetAttachments() []GetInboxConversationMessages200ResponseMessagesInnerAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetAttachmentsOk() (*[]GetInboxConversationMessages200ResponseMessagesInnerAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetAttachments(v []GetInboxConversationMessages200ResponseMessagesInnerAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetSubject

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetStoryReply

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetStoryReply() bool`

GetStoryReply returns the StoryReply field if non-nil, zero value otherwise.

### GetStoryReplyOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetStoryReplyOk() (*bool, bool)`

GetStoryReplyOk returns a tuple with the StoryReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryReply

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetStoryReply(v bool)`

SetStoryReply sets StoryReply field to given value.

### HasStoryReply

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasStoryReply() bool`

HasStoryReply returns a boolean if a field has been set.

### SetStoryReplyNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetStoryReplyNil(b bool)`

 SetStoryReplyNil sets the value for StoryReply to be an explicit nil

### UnsetStoryReply
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetStoryReply()`

UnsetStoryReply ensures that no value is present for StoryReply, not even an explicit nil
### GetIsStoryMention

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIsStoryMention() bool`

GetIsStoryMention returns the IsStoryMention field if non-nil, zero value otherwise.

### GetIsStoryMentionOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIsStoryMentionOk() (*bool, bool)`

GetIsStoryMentionOk returns a tuple with the IsStoryMention field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStoryMention

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetIsStoryMention(v bool)`

SetIsStoryMention sets IsStoryMention field to given value.

### HasIsStoryMention

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasIsStoryMention() bool`

HasIsStoryMention returns a boolean if a field has been set.

### SetIsStoryMentionNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetIsStoryMentionNil(b bool)`

 SetIsStoryMentionNil sets the value for IsStoryMention to be an explicit nil

### UnsetIsStoryMention
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetIsStoryMention()`

UnsetIsStoryMention ensures that no value is present for IsStoryMention, not even an explicit nil
### GetIsEdited

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIsEdited() bool`

GetIsEdited returns the IsEdited field if non-nil, zero value otherwise.

### GetIsEditedOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIsEditedOk() (*bool, bool)`

GetIsEditedOk returns a tuple with the IsEdited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEdited

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetIsEdited(v bool)`

SetIsEdited sets IsEdited field to given value.

### HasIsEdited

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasIsEdited() bool`

HasIsEdited returns a boolean if a field has been set.

### GetEditedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetEditedAt() time.Time`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetEditedAtOk() (*time.Time, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetEditedAt(v time.Time)`

SetEditedAt sets EditedAt field to given value.

### HasEditedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasEditedAt() bool`

HasEditedAt returns a boolean if a field has been set.

### SetEditedAtNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetEditedAtNil(b bool)`

 SetEditedAtNil sets the value for EditedAt to be an explicit nil

### UnsetEditedAt
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetEditedAt()`

UnsetEditedAt ensures that no value is present for EditedAt, not even an explicit nil
### GetEditCount

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetEditCount() int32`

GetEditCount returns the EditCount field if non-nil, zero value otherwise.

### GetEditCountOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetEditCountOk() (*int32, bool)`

GetEditCountOk returns a tuple with the EditCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditCount

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetEditCount(v int32)`

SetEditCount sets EditCount field to given value.

### HasEditCount

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasEditCount() bool`

HasEditCount returns a boolean if a field has been set.

### GetEditHistory

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetEditHistory() []GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInner`

GetEditHistory returns the EditHistory field if non-nil, zero value otherwise.

### GetEditHistoryOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetEditHistoryOk() (*[]GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInner, bool)`

GetEditHistoryOk returns a tuple with the EditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditHistory

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetEditHistory(v []GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInner)`

SetEditHistory sets EditHistory field to given value.

### HasEditHistory

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasEditHistory() bool`

HasEditHistory returns a boolean if a field has been set.

### GetIsDeleted

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetDeletedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetDeliveryStatus

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### SetDeliveryStatusNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeliveryStatusNil(b bool)`

 SetDeliveryStatusNil sets the value for DeliveryStatus to be an explicit nil

### UnsetDeliveryStatus
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetDeliveryStatus()`

UnsetDeliveryStatus ensures that no value is present for DeliveryStatus, not even an explicit nil
### GetDeliveredAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetReadAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.

### SetReadAtNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetReadAtNil(b bool)`

 SetReadAtNil sets the value for ReadAt to be an explicit nil

### UnsetReadAt
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetReadAt()`

UnsetReadAt ensures that no value is present for ReadAt, not even an explicit nil
### GetSentAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *GetInboxConversationMessages200ResponseMessagesInner) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetDeliveryError

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeliveryError() GetInboxConversationMessages200ResponseMessagesInnerDeliveryError`

GetDeliveryError returns the DeliveryError field if non-nil, zero value otherwise.

### GetDeliveryErrorOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetDeliveryErrorOk() (*GetInboxConversationMessages200ResponseMessagesInnerDeliveryError, bool)`

GetDeliveryErrorOk returns a tuple with the DeliveryError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryError

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetDeliveryError(v GetInboxConversationMessages200ResponseMessagesInnerDeliveryError)`

SetDeliveryError sets DeliveryError field to given value.

### HasDeliveryError

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasDeliveryError() bool`

HasDeliveryError returns a boolean if a field has been set.

### GetReactions

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetReactions() []GetInboxConversationMessages200ResponseMessagesInnerReactionsInner`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetReactionsOk() (*[]GetInboxConversationMessages200ResponseMessagesInnerReactionsInner, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetReactions(v []GetInboxConversationMessages200ResponseMessagesInnerReactionsInner)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetMetadata

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetInboxConversationMessages200ResponseMessagesInner) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetInboxConversationMessages200ResponseMessagesInner) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetInboxConversationMessages200ResponseMessagesInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


