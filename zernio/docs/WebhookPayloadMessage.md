# WebhookPayloadMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Message** | [**WebhookPayloadMessageMessage**](WebhookPayloadMessageMessage.md) |  | 
**Conversation** | [**InboxWebhookConversation**](InboxWebhookConversation.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Metadata** | Pointer to [**WebhookPayloadMessageMetadata**](WebhookPayloadMessageMetadata.md) |  | [optional] 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadMessage

`func NewWebhookPayloadMessage(id string, event string, message WebhookPayloadMessageMessage, conversation InboxWebhookConversation, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadMessage`

NewWebhookPayloadMessage instantiates a new WebhookPayloadMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageWithDefaults

`func NewWebhookPayloadMessageWithDefaults() *WebhookPayloadMessage`

NewWebhookPayloadMessageWithDefaults instantiates a new WebhookPayloadMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadMessage) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadMessage) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadMessage) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadMessage) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMessage

`func (o *WebhookPayloadMessage) GetMessage() WebhookPayloadMessageMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookPayloadMessage) GetMessageOk() (*WebhookPayloadMessageMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookPayloadMessage) SetMessage(v WebhookPayloadMessageMessage)`

SetMessage sets Message field to given value.


### GetConversation

`func (o *WebhookPayloadMessage) GetConversation() InboxWebhookConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WebhookPayloadMessage) GetConversationOk() (*InboxWebhookConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WebhookPayloadMessage) SetConversation(v InboxWebhookConversation)`

SetConversation sets Conversation field to given value.


### GetAccount

`func (o *WebhookPayloadMessage) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadMessage) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadMessage) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetMetadata

`func (o *WebhookPayloadMessage) GetMetadata() WebhookPayloadMessageMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *WebhookPayloadMessage) GetMetadataOk() (*WebhookPayloadMessageMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *WebhookPayloadMessage) SetMetadata(v WebhookPayloadMessageMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *WebhookPayloadMessage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookPayloadMessage) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadMessage) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadMessage) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


