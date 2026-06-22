# WebhookPayloadReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Reaction** | [**WebhookPayloadReactionReaction**](WebhookPayloadReactionReaction.md) |  | 
**Conversation** | [**InboxWebhookConversation**](InboxWebhookConversation.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadReaction

`func NewWebhookPayloadReaction(id string, event string, reaction WebhookPayloadReactionReaction, conversation InboxWebhookConversation, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadReaction`

NewWebhookPayloadReaction instantiates a new WebhookPayloadReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadReactionWithDefaults

`func NewWebhookPayloadReactionWithDefaults() *WebhookPayloadReaction`

NewWebhookPayloadReactionWithDefaults instantiates a new WebhookPayloadReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadReaction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadReaction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadReaction) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadReaction) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadReaction) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadReaction) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetReaction

`func (o *WebhookPayloadReaction) GetReaction() WebhookPayloadReactionReaction`

GetReaction returns the Reaction field if non-nil, zero value otherwise.

### GetReactionOk

`func (o *WebhookPayloadReaction) GetReactionOk() (*WebhookPayloadReactionReaction, bool)`

GetReactionOk returns a tuple with the Reaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaction

`func (o *WebhookPayloadReaction) SetReaction(v WebhookPayloadReactionReaction)`

SetReaction sets Reaction field to given value.


### GetConversation

`func (o *WebhookPayloadReaction) GetConversation() InboxWebhookConversation`

GetConversation returns the Conversation field if non-nil, zero value otherwise.

### GetConversationOk

`func (o *WebhookPayloadReaction) GetConversationOk() (*InboxWebhookConversation, bool)`

GetConversationOk returns a tuple with the Conversation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversation

`func (o *WebhookPayloadReaction) SetConversation(v InboxWebhookConversation)`

SetConversation sets Conversation field to given value.


### GetAccount

`func (o *WebhookPayloadReaction) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadReaction) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadReaction) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadReaction) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadReaction) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadReaction) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


