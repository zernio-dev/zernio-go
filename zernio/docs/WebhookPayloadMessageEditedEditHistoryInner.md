# WebhookPayloadMessageEditedEditHistoryInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **NullableString** |  | 
**Attachments** | [**[]GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInnerAttachmentsInner**](GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInnerAttachmentsInner.md) |  | 
**EditedAt** | **time.Time** |  | 

## Methods

### NewWebhookPayloadMessageEditedEditHistoryInner

`func NewWebhookPayloadMessageEditedEditHistoryInner(text NullableString, attachments []GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInnerAttachmentsInner, editedAt time.Time, ) *WebhookPayloadMessageEditedEditHistoryInner`

NewWebhookPayloadMessageEditedEditHistoryInner instantiates a new WebhookPayloadMessageEditedEditHistoryInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageEditedEditHistoryInnerWithDefaults

`func NewWebhookPayloadMessageEditedEditHistoryInnerWithDefaults() *WebhookPayloadMessageEditedEditHistoryInner`

NewWebhookPayloadMessageEditedEditHistoryInnerWithDefaults instantiates a new WebhookPayloadMessageEditedEditHistoryInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *WebhookPayloadMessageEditedEditHistoryInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WebhookPayloadMessageEditedEditHistoryInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WebhookPayloadMessageEditedEditHistoryInner) SetText(v string)`

SetText sets Text field to given value.


### SetTextNil

`func (o *WebhookPayloadMessageEditedEditHistoryInner) SetTextNil(b bool)`

 SetTextNil sets the value for Text to be an explicit nil

### UnsetText
`func (o *WebhookPayloadMessageEditedEditHistoryInner) UnsetText()`

UnsetText ensures that no value is present for Text, not even an explicit nil
### GetAttachments

`func (o *WebhookPayloadMessageEditedEditHistoryInner) GetAttachments() []GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInnerAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *WebhookPayloadMessageEditedEditHistoryInner) GetAttachmentsOk() (*[]GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInnerAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *WebhookPayloadMessageEditedEditHistoryInner) SetAttachments(v []GetInboxConversationMessages200ResponseMessagesInnerEditHistoryInnerAttachmentsInner)`

SetAttachments sets Attachments field to given value.


### GetEditedAt

`func (o *WebhookPayloadMessageEditedEditHistoryInner) GetEditedAt() time.Time`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *WebhookPayloadMessageEditedEditHistoryInner) GetEditedAtOk() (*time.Time, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *WebhookPayloadMessageEditedEditHistoryInner) SetEditedAt(v time.Time)`

SetEditedAt sets EditedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


