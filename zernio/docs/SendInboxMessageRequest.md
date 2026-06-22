# SendInboxMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Social account ID | 
**Message** | Pointer to **string** | Message text | [optional] 
**AttachmentUrl** | Pointer to **string** | URL of the attachment to send (image, video, audio, or file). The URL must be publicly accessible. For binary file uploads, use multipart/form-data instead. | [optional] 
**AttachmentType** | Pointer to **string** | Type of attachment. Defaults to file if not specified. | [optional] 
**VoiceNote** | Pointer to **bool** | WhatsApp only. When &#x60;true&#x60; on an audio attachment, the message is sent as a voice message (PTT) — the recipient sees the waveform + voice-note UI instead of a basic audio attachment. The audio file MUST be &#x60;.ogg&#x60; encoded with the OPUS codec (mono) per Meta&#39;s voice-message contract; other formats are rejected by WhatsApp. Ignored for non-audio attachments.  | [optional] 
**QuickReplies** | Pointer to [**[]SendInboxMessageRequestQuickRepliesInner**](SendInboxMessageRequestQuickRepliesInner.md) | Quick reply buttons. Mutually exclusive with buttons. Max 13 items. | [optional] 
**Buttons** | Pointer to [**[]SendInboxMessageRequestButtonsInner**](SendInboxMessageRequestButtonsInner.md) | Action buttons. Mutually exclusive with quickReplies. Max 3 items. | [optional] 
**Template** | Pointer to [**SendInboxMessageRequestTemplate**](SendInboxMessageRequestTemplate.md) |  | [optional] 
**Interactive** | Pointer to [**SendInboxMessageRequestInteractive**](SendInboxMessageRequestInteractive.md) |  | [optional] 
**ReplyMarkup** | Pointer to [**SendInboxMessageRequestReplyMarkup**](SendInboxMessageRequestReplyMarkup.md) |  | [optional] 
**MessagingType** | Pointer to **string** | Facebook messaging type. Required when using messageTag. | [optional] 
**MessageTag** | Pointer to **string** | Facebook message tag for messaging outside 24h window. Requires messagingType MESSAGE_TAG. Instagram only supports HUMAN_AGENT. | [optional] 
**ReplyTo** | Pointer to **string** | Platform message ID to quote-reply to. For WhatsApp, pass the wamid (available in message.platformMessageId from webhooks). For Telegram, pass the Telegram message ID. | [optional] 
**Location** | Pointer to [**SendInboxMessageRequestLocation**](SendInboxMessageRequestLocation.md) |  | [optional] 
**Contacts** | Pointer to [**[]SendInboxMessageRequestContactsInner**](SendInboxMessageRequestContactsInner.md) | WhatsApp-only. Send one or more contact cards. | [optional] 

## Methods

### NewSendInboxMessageRequest

`func NewSendInboxMessageRequest(accountId string, ) *SendInboxMessageRequest`

NewSendInboxMessageRequest instantiates a new SendInboxMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestWithDefaults

`func NewSendInboxMessageRequestWithDefaults() *SendInboxMessageRequest`

NewSendInboxMessageRequestWithDefaults instantiates a new SendInboxMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SendInboxMessageRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SendInboxMessageRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SendInboxMessageRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetMessage

`func (o *SendInboxMessageRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendInboxMessageRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendInboxMessageRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendInboxMessageRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAttachmentUrl

`func (o *SendInboxMessageRequest) GetAttachmentUrl() string`

GetAttachmentUrl returns the AttachmentUrl field if non-nil, zero value otherwise.

### GetAttachmentUrlOk

`func (o *SendInboxMessageRequest) GetAttachmentUrlOk() (*string, bool)`

GetAttachmentUrlOk returns a tuple with the AttachmentUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentUrl

`func (o *SendInboxMessageRequest) SetAttachmentUrl(v string)`

SetAttachmentUrl sets AttachmentUrl field to given value.

### HasAttachmentUrl

`func (o *SendInboxMessageRequest) HasAttachmentUrl() bool`

HasAttachmentUrl returns a boolean if a field has been set.

### GetAttachmentType

`func (o *SendInboxMessageRequest) GetAttachmentType() string`

GetAttachmentType returns the AttachmentType field if non-nil, zero value otherwise.

### GetAttachmentTypeOk

`func (o *SendInboxMessageRequest) GetAttachmentTypeOk() (*string, bool)`

GetAttachmentTypeOk returns a tuple with the AttachmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachmentType

`func (o *SendInboxMessageRequest) SetAttachmentType(v string)`

SetAttachmentType sets AttachmentType field to given value.

### HasAttachmentType

`func (o *SendInboxMessageRequest) HasAttachmentType() bool`

HasAttachmentType returns a boolean if a field has been set.

### GetVoiceNote

`func (o *SendInboxMessageRequest) GetVoiceNote() bool`

GetVoiceNote returns the VoiceNote field if non-nil, zero value otherwise.

### GetVoiceNoteOk

`func (o *SendInboxMessageRequest) GetVoiceNoteOk() (*bool, bool)`

GetVoiceNoteOk returns a tuple with the VoiceNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceNote

`func (o *SendInboxMessageRequest) SetVoiceNote(v bool)`

SetVoiceNote sets VoiceNote field to given value.

### HasVoiceNote

`func (o *SendInboxMessageRequest) HasVoiceNote() bool`

HasVoiceNote returns a boolean if a field has been set.

### GetQuickReplies

`func (o *SendInboxMessageRequest) GetQuickReplies() []SendInboxMessageRequestQuickRepliesInner`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *SendInboxMessageRequest) GetQuickRepliesOk() (*[]SendInboxMessageRequestQuickRepliesInner, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *SendInboxMessageRequest) SetQuickReplies(v []SendInboxMessageRequestQuickRepliesInner)`

SetQuickReplies sets QuickReplies field to given value.

### HasQuickReplies

`func (o *SendInboxMessageRequest) HasQuickReplies() bool`

HasQuickReplies returns a boolean if a field has been set.

### GetButtons

`func (o *SendInboxMessageRequest) GetButtons() []SendInboxMessageRequestButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *SendInboxMessageRequest) GetButtonsOk() (*[]SendInboxMessageRequestButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *SendInboxMessageRequest) SetButtons(v []SendInboxMessageRequestButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *SendInboxMessageRequest) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetTemplate

`func (o *SendInboxMessageRequest) GetTemplate() SendInboxMessageRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *SendInboxMessageRequest) GetTemplateOk() (*SendInboxMessageRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *SendInboxMessageRequest) SetTemplate(v SendInboxMessageRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *SendInboxMessageRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetInteractive

`func (o *SendInboxMessageRequest) GetInteractive() SendInboxMessageRequestInteractive`

GetInteractive returns the Interactive field if non-nil, zero value otherwise.

### GetInteractiveOk

`func (o *SendInboxMessageRequest) GetInteractiveOk() (*SendInboxMessageRequestInteractive, bool)`

GetInteractiveOk returns a tuple with the Interactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractive

`func (o *SendInboxMessageRequest) SetInteractive(v SendInboxMessageRequestInteractive)`

SetInteractive sets Interactive field to given value.

### HasInteractive

`func (o *SendInboxMessageRequest) HasInteractive() bool`

HasInteractive returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *SendInboxMessageRequest) GetReplyMarkup() SendInboxMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *SendInboxMessageRequest) GetReplyMarkupOk() (*SendInboxMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *SendInboxMessageRequest) SetReplyMarkup(v SendInboxMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *SendInboxMessageRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.

### GetMessagingType

`func (o *SendInboxMessageRequest) GetMessagingType() string`

GetMessagingType returns the MessagingType field if non-nil, zero value otherwise.

### GetMessagingTypeOk

`func (o *SendInboxMessageRequest) GetMessagingTypeOk() (*string, bool)`

GetMessagingTypeOk returns a tuple with the MessagingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingType

`func (o *SendInboxMessageRequest) SetMessagingType(v string)`

SetMessagingType sets MessagingType field to given value.

### HasMessagingType

`func (o *SendInboxMessageRequest) HasMessagingType() bool`

HasMessagingType returns a boolean if a field has been set.

### GetMessageTag

`func (o *SendInboxMessageRequest) GetMessageTag() string`

GetMessageTag returns the MessageTag field if non-nil, zero value otherwise.

### GetMessageTagOk

`func (o *SendInboxMessageRequest) GetMessageTagOk() (*string, bool)`

GetMessageTagOk returns a tuple with the MessageTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTag

`func (o *SendInboxMessageRequest) SetMessageTag(v string)`

SetMessageTag sets MessageTag field to given value.

### HasMessageTag

`func (o *SendInboxMessageRequest) HasMessageTag() bool`

HasMessageTag returns a boolean if a field has been set.

### GetReplyTo

`func (o *SendInboxMessageRequest) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *SendInboxMessageRequest) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *SendInboxMessageRequest) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *SendInboxMessageRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetLocation

`func (o *SendInboxMessageRequest) GetLocation() SendInboxMessageRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SendInboxMessageRequest) GetLocationOk() (*SendInboxMessageRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SendInboxMessageRequest) SetLocation(v SendInboxMessageRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SendInboxMessageRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetContacts

`func (o *SendInboxMessageRequest) GetContacts() []SendInboxMessageRequestContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *SendInboxMessageRequest) GetContactsOk() (*[]SendInboxMessageRequestContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *SendInboxMessageRequest) SetContacts(v []SendInboxMessageRequestContactsInner)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *SendInboxMessageRequest) HasContacts() bool`

HasContacts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


