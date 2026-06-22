# WebhookPayloadMessageDeleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Message** | [**InboxWebhookMessage**](InboxWebhookMessage.md) |  | 
**DeletedAt** | **time.Time** |  | 
**Conversation** | [**InboxWebhookConversation**](InboxWebhookConversation.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadMessageDeleted

`func NewWebhookPayloadMessageDeleted(id string, event string, message InboxWebhookMessage, deletedAt time.Time, conversation InboxWebhookConversation, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadMessageDeleted`

NewWebhookPayloadMessageDeleted instantiates a new WebhookPayloadMessageDeleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageDeletedWithDefaults

`func NewWebhookPayloadMessageDeletedWithDefaults() *WebhookPayloadMessageDeleted`

NewWebhookPayloadMessageDeletedWithDefaults instantiates a new WebhookPayloadMessageDeleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadMessageDeleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadMessageDeleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadMessageDeleted) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadMessageDeleted) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadMessageDeleted) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadMessageDeleted) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMessage

`func (o *WebhookPayloadMessageDeleted) GetMessage() InboxWebhookMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookPayloadMessageDeleted) GetMessageOk() (*InboxWebhookMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookPayloadMessageDeleted) SetMessage(v InboxWebhookMessage)`

SetMessage sets Message field to given value.


### GetDeletedAt

`func (o *WebhookPayloadMessageDeleted) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *WebhookPayloadMessageDeleted) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *WebhookPayloadMessageDeleted) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.


### GetConversation

`func (o *WebhookPayloadMessageDeleted) GetConversation() InboxWebhookConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WebhookPayloadMessageDeleted) GetConversationOk() (*InboxWebhookConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WebhookPayloadMessageDeleted) SetConversation(v InboxWebhookConversation)`

SetConversation sets Conversation field to given value.


### GetAccount

`func (o *WebhookPayloadMessageDeleted) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadMessageDeleted) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadMessageDeleted) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadMessageDeleted) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadMessageDeleted) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadMessageDeleted) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


