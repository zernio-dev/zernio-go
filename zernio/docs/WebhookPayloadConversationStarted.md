# WebhookPayloadConversationStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Conversation** | [**WebhookPayloadConversationStartedConversation**](WebhookPayloadConversationStartedConversation.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**StartedAt** | **time.Time** | When the conversation document was created. | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadConversationStarted

`func NewWebhookPayloadConversationStarted(id string, event string, conversation WebhookPayloadConversationStartedConversation, account InboxWebhookAccount, startedAt time.Time, timestamp time.Time, ) *WebhookPayloadConversationStarted`

NewWebhookPayloadConversationStarted instantiates a new WebhookPayloadConversationStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadConversationStartedWithDefaults

`func NewWebhookPayloadConversationStartedWithDefaults() *WebhookPayloadConversationStarted`

NewWebhookPayloadConversationStartedWithDefaults instantiates a new WebhookPayloadConversationStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadConversationStarted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadConversationStarted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadConversationStarted) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadConversationStarted) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadConversationStarted) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadConversationStarted) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetConversation

`func (o *WebhookPayloadConversationStarted) GetConversation() WebhookPayloadConversationStartedConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WebhookPayloadConversationStarted) GetConversationOk() (*WebhookPayloadConversationStartedConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WebhookPayloadConversationStarted) SetConversation(v WebhookPayloadConversationStartedConversation)`

SetConversation sets Conversation field to given value.


### GetAccount

`func (o *WebhookPayloadConversationStarted) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadConversationStarted) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadConversationStarted) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetStartedAt

`func (o *WebhookPayloadConversationStarted) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WebhookPayloadConversationStarted) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WebhookPayloadConversationStarted) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetTimestamp

`func (o *WebhookPayloadConversationStarted) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadConversationStarted) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadConversationStarted) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


