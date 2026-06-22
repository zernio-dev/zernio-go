# WebhookPayloadMessageEdited

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Message** | [**InboxWebhookMessage**](InboxWebhookMessage.md) |  | 
**EditHistory** | [**[]WebhookPayloadMessageEditedEditHistoryInner**](WebhookPayloadMessageEditedEditHistoryInner.md) | Prior versions of the message, oldest first. | 
**EditCount** | **int32** | Total number of edits applied to this message. | 
**EditedAt** | **time.Time** | When the most recent edit happened. | 
**Conversation** | [**InboxWebhookConversation**](InboxWebhookConversation.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadMessageEdited

`func NewWebhookPayloadMessageEdited(id string, event string, message InboxWebhookMessage, editHistory []WebhookPayloadMessageEditedEditHistoryInner, editCount int32, editedAt time.Time, conversation InboxWebhookConversation, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadMessageEdited`

NewWebhookPayloadMessageEdited instantiates a new WebhookPayloadMessageEdited object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageEditedWithDefaults

`func NewWebhookPayloadMessageEditedWithDefaults() *WebhookPayloadMessageEdited`

NewWebhookPayloadMessageEditedWithDefaults instantiates a new WebhookPayloadMessageEdited object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadMessageEdited) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadMessageEdited) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadMessageEdited) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadMessageEdited) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadMessageEdited) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadMessageEdited) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMessage

`func (o *WebhookPayloadMessageEdited) GetMessage() InboxWebhookMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookPayloadMessageEdited) GetMessageOk() (*InboxWebhookMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookPayloadMessageEdited) SetMessage(v InboxWebhookMessage)`

SetMessage sets Message field to given value.


### GetEditHistory

`func (o *WebhookPayloadMessageEdited) GetEditHistory() []WebhookPayloadMessageEditedEditHistoryInner`

GetEditHistory returns the EditHistory field if non-nil, zero value otherwise.

### GetEditHistoryOk

`func (o *WebhookPayloadMessageEdited) GetEditHistoryOk() (*[]WebhookPayloadMessageEditedEditHistoryInner, bool)`

GetEditHistoryOk returns a tuple with the EditHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditHistory

`func (o *WebhookPayloadMessageEdited) SetEditHistory(v []WebhookPayloadMessageEditedEditHistoryInner)`

SetEditHistory sets EditHistory field to given value.


### GetEditCount

`func (o *WebhookPayloadMessageEdited) GetEditCount() int32`

GetEditCount returns the EditCount field if non-nil, zero value otherwise.

### GetEditCountOk

`func (o *WebhookPayloadMessageEdited) GetEditCountOk() (*int32, bool)`

GetEditCountOk returns a tuple with the EditCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditCount

`func (o *WebhookPayloadMessageEdited) SetEditCount(v int32)`

SetEditCount sets EditCount field to given value.


### GetEditedAt

`func (o *WebhookPayloadMessageEdited) GetEditedAt() time.Time`

GetEditedAt returns the EditedAt field if non-nil, zero value otherwise.

### GetEditedAtOk

`func (o *WebhookPayloadMessageEdited) GetEditedAtOk() (*time.Time, bool)`

GetEditedAtOk returns a tuple with the EditedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedAt

`func (o *WebhookPayloadMessageEdited) SetEditedAt(v time.Time)`

SetEditedAt sets EditedAt field to given value.


### GetConversation

`func (o *WebhookPayloadMessageEdited) GetConversation() InboxWebhookConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WebhookPayloadMessageEdited) GetConversationOk() (*InboxWebhookConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WebhookPayloadMessageEdited) SetConversation(v InboxWebhookConversation)`

SetConversation sets Conversation field to given value.


### GetAccount

`func (o *WebhookPayloadMessageEdited) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadMessageEdited) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadMessageEdited) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadMessageEdited) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadMessageEdited) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadMessageEdited) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


