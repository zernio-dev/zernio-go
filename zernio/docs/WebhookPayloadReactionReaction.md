# WebhookPayloadReactionReaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | **string** | The emoji reacted with. May be an empty string when &#x60;action&#x60; is &#x60;removed&#x60; on WhatsApp (Meta does not report which emoji was removed).  | 
**Action** | **string** |  | 
**MessageId** | Pointer to **string** | Internal Zernio message ID of the reacted-to message, when resolvable from the platform ID. | [optional] 
**PlatformMessageId** | **string** | Platform-native ID of the reacted-to message (e.g. WhatsApp wamid). | 
**Sender** | [**WebhookPayloadReactionReactionSender**](WebhookPayloadReactionReactionSender.md) |  | 
**ReactedAt** | **time.Time** |  | 

## Methods

### NewWebhookPayloadReactionReaction

`func NewWebhookPayloadReactionReaction(emoji string, action string, platformMessageId string, sender WebhookPayloadReactionReactionSender, reactedAt time.Time, ) *WebhookPayloadReactionReaction`

NewWebhookPayloadReactionReaction instantiates a new WebhookPayloadReactionReaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadReactionReactionWithDefaults

`func NewWebhookPayloadReactionReactionWithDefaults() *WebhookPayloadReactionReaction`

NewWebhookPayloadReactionReactionWithDefaults instantiates a new WebhookPayloadReactionReaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *WebhookPayloadReactionReaction) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *WebhookPayloadReactionReaction) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *WebhookPayloadReactionReaction) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.


### GetAction

`func (o *WebhookPayloadReactionReaction) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WebhookPayloadReactionReaction) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WebhookPayloadReactionReaction) SetAction(v string)`

SetAction sets Action field to given value.


### GetMessageId

`func (o *WebhookPayloadReactionReaction) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *WebhookPayloadReactionReaction) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *WebhookPayloadReactionReaction) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *WebhookPayloadReactionReaction) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetPlatformMessageId

`func (o *WebhookPayloadReactionReaction) GetPlatformMessageId() string`

GetPlatformMessageId returns the PlatformMessageId field if non-nil, zero value otherwise.

### GetPlatformMessageIdOk

`func (o *WebhookPayloadReactionReaction) GetPlatformMessageIdOk() (*string, bool)`

GetPlatformMessageIdOk returns a tuple with the PlatformMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformMessageId

`func (o *WebhookPayloadReactionReaction) SetPlatformMessageId(v string)`

SetPlatformMessageId sets PlatformMessageId field to given value.


### GetSender

`func (o *WebhookPayloadReactionReaction) GetSender() WebhookPayloadReactionReactionSender`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *WebhookPayloadReactionReaction) GetSenderOk() (*WebhookPayloadReactionReactionSender, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *WebhookPayloadReactionReaction) SetSender(v WebhookPayloadReactionReactionSender)`

SetSender sets Sender field to given value.


### GetReactedAt

`func (o *WebhookPayloadReactionReaction) GetReactedAt() time.Time`

GetReactedAt returns the ReactedAt field if non-nil, zero value otherwise.

### GetReactedAtOk

`func (o *WebhookPayloadReactionReaction) GetReactedAtOk() (*time.Time, bool)`

GetReactedAtOk returns a tuple with the ReactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactedAt

`func (o *WebhookPayloadReactionReaction) SetReactedAt(v time.Time)`

SetReactedAt sets ReactedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


