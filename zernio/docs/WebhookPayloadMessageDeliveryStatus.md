# WebhookPayloadMessageDeliveryStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Message** | [**InboxWebhookMessage**](InboxWebhookMessage.md) |  | 
**StatusAt** | **time.Time** | When the platform reported this status. | 
**Error** | Pointer to [**WebhookPayloadMessageDeliveryStatusError**](WebhookPayloadMessageDeliveryStatusError.md) |  | [optional] 
**Conversation** | [**InboxWebhookConversation**](InboxWebhookConversation.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadMessageDeliveryStatus

`func NewWebhookPayloadMessageDeliveryStatus(id string, event string, message InboxWebhookMessage, statusAt time.Time, conversation InboxWebhookConversation, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadMessageDeliveryStatus`

NewWebhookPayloadMessageDeliveryStatus instantiates a new WebhookPayloadMessageDeliveryStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageDeliveryStatusWithDefaults

`func NewWebhookPayloadMessageDeliveryStatusWithDefaults() *WebhookPayloadMessageDeliveryStatus`

NewWebhookPayloadMessageDeliveryStatusWithDefaults instantiates a new WebhookPayloadMessageDeliveryStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadMessageDeliveryStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadMessageDeliveryStatus) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadMessageDeliveryStatus) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadMessageDeliveryStatus) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetMessage

`func (o *WebhookPayloadMessageDeliveryStatus) GetMessage() InboxWebhookMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetMessageOk() (*InboxWebhookMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookPayloadMessageDeliveryStatus) SetMessage(v InboxWebhookMessage)`

SetMessage sets Message field to given value.


### GetStatusAt

`func (o *WebhookPayloadMessageDeliveryStatus) GetStatusAt() time.Time`

GetStatusAt returns the StatusAt field if non-nil, zero value otherwise.

### GetStatusAtOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetStatusAtOk() (*time.Time, bool)`

GetStatusAtOk returns a tuple with the StatusAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusAt

`func (o *WebhookPayloadMessageDeliveryStatus) SetStatusAt(v time.Time)`

SetStatusAt sets StatusAt field to given value.


### GetError

`func (o *WebhookPayloadMessageDeliveryStatus) GetError() WebhookPayloadMessageDeliveryStatusError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetErrorOk() (*WebhookPayloadMessageDeliveryStatusError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookPayloadMessageDeliveryStatus) SetError(v WebhookPayloadMessageDeliveryStatusError)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookPayloadMessageDeliveryStatus) HasError() bool`

HasError returns a boolean if a field has been set.

### GetConversation

`func (o *WebhookPayloadMessageDeliveryStatus) GetConversation() InboxWebhookConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetConversationOk() (*InboxWebhookConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WebhookPayloadMessageDeliveryStatus) SetConversation(v InboxWebhookConversation)`

SetConversation sets Conversation field to given value.


### GetAccount

`func (o *WebhookPayloadMessageDeliveryStatus) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadMessageDeliveryStatus) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadMessageDeliveryStatus) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadMessageDeliveryStatus) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadMessageDeliveryStatus) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


