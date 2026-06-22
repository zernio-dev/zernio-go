# InboxWebhookMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal message ID | 
**ConversationId** | **string** | Internal conversation ID | 
**Platform** | **string** |  | 
**PlatformMessageId** | **string** | Platform&#39;s message ID | 
**Direction** | **string** |  | 
**Text** | **NullableString** | Message text content (retained on deleted messages for API consumers; Zernio dashboard UI hides this) | 
**Attachments** | [**[]InboxWebhookMessageAttachmentsInner**](InboxWebhookMessageAttachmentsInner.md) |  | 
**Sender** | [**InboxWebhookMessageSender**](InboxWebhookMessageSender.md) |  | 
**SentAt** | **time.Time** |  | 
**IsRead** | **bool** |  | 

## Methods

### NewInboxWebhookMessage

`func NewInboxWebhookMessage(id string, conversationId string, platform string, platformMessageId string, direction string, text NullableString, attachments []InboxWebhookMessageAttachmentsInner, sender InboxWebhookMessageSender, sentAt time.Time, isRead bool, ) *InboxWebhookMessage`

NewInboxWebhookMessage instantiates a new InboxWebhookMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxWebhookMessageWithDefaults

`func NewInboxWebhookMessageWithDefaults() *InboxWebhookMessage`

NewInboxWebhookMessageWithDefaults instantiates a new InboxWebhookMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboxWebhookMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxWebhookMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxWebhookMessage) SetId(v string)`

SetId sets Id field to given value.


### GetConversationId

`func (o *InboxWebhookMessage) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *InboxWebhookMessage) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *InboxWebhookMessage) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetPlatform

`func (o *InboxWebhookMessage) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InboxWebhookMessage) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InboxWebhookMessage) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformMessageId

`func (o *InboxWebhookMessage) GetPlatformMessageId() string`

GetPlatformMessageId returns the PlatformMessageId field if non-nil, zero value otherwise.

### GetPlatformMessageIdOk

`func (o *InboxWebhookMessage) GetPlatformMessageIdOk() (*string, bool)`

GetPlatformMessageIdOk returns a tuple with the PlatformMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformMessageId

`func (o *InboxWebhookMessage) SetPlatformMessageId(v string)`

SetPlatformMessageId sets PlatformMessageId field to given value.


### GetDirection

`func (o *InboxWebhookMessage) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InboxWebhookMessage) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InboxWebhookMessage) SetDirection(v string)`

SetDirection sets Direction field to given value.


### GetText

`func (o *InboxWebhookMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *InboxWebhookMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *InboxWebhookMessage) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *InboxWebhookMessage) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *InboxWebhookMessage) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetAttachments

`func (o *InboxWebhookMessage) GetAttachments() []InboxWebhookMessageAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InboxWebhookMessage) GetAttachmentsOk() (*[]InboxWebhookMessageAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InboxWebhookMessage) SetAttachments(v []InboxWebhookMessageAttachmentsInner)`

SetAttachments sets Attachments field to given value.


### GetSender

`func (o *InboxWebhookMessage) GetSender() InboxWebhookMessageSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *InboxWebhookMessage) GetSenderOk() (*InboxWebhookMessageSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *InboxWebhookMessage) SetSender(v InboxWebhookMessageSender)`

SetSender sets Sender field to given value.


### GetSentAt

`func (o *InboxWebhookMessage) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *InboxWebhookMessage) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *InboxWebhookMessage) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.


### GetIsRead

`func (o *InboxWebhookMessage) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *InboxWebhookMessage) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *InboxWebhookMessage) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


