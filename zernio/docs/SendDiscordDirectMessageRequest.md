# SendDiscordDirectMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | SocialAccount _id of the connected Discord account the bot speaks as. Caller must own the account (directly or via team membership). | 
**UserId** | **string** | Discord snowflake ID of the recipient (15-21 digits). | 
**Content** | Pointer to **string** | Message text, up to 2,000 characters. | [optional] 
**Embeds** | Pointer to **[]map[string]interface{}** | Up to 10 Discord embeds. Same shape as channel-post embeds (title, description, color, fields, etc.). See DiscordPlatformData.embeds for the embed object schema. | [optional] 
**Attachments** | Pointer to [**[]SendDiscordDirectMessageRequestAttachmentsInner**](SendDiscordDirectMessageRequestAttachmentsInner.md) | Up to 10 media attachments. Each is &#x60;{ type: image|video|gif|document, url, filename?, mimeType?, size? }&#x60;. | [optional] 
**Tts** | Pointer to **bool** | Send as text-to-speech message. | [optional] 

## Methods

### NewSendDiscordDirectMessageRequest

`func NewSendDiscordDirectMessageRequest(accountId string, userId string, ) *SendDiscordDirectMessageRequest`

NewSendDiscordDirectMessageRequest instantiates a new SendDiscordDirectMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDiscordDirectMessageRequestWithDefaults

`func NewSendDiscordDirectMessageRequestWithDefaults() *SendDiscordDirectMessageRequest`

NewSendDiscordDirectMessageRequestWithDefaults instantiates a new SendDiscordDirectMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SendDiscordDirectMessageRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SendDiscordDirectMessageRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SendDiscordDirectMessageRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetUserId

`func (o *SendDiscordDirectMessageRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SendDiscordDirectMessageRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SendDiscordDirectMessageRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetContent

`func (o *SendDiscordDirectMessageRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendDiscordDirectMessageRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendDiscordDirectMessageRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *SendDiscordDirectMessageRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEmbeds

`func (o *SendDiscordDirectMessageRequest) GetEmbeds() []map[string]interface{}`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *SendDiscordDirectMessageRequest) GetEmbedsOk() (*[]map[string]interface{}, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *SendDiscordDirectMessageRequest) SetEmbeds(v []map[string]interface{})`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *SendDiscordDirectMessageRequest) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetAttachments

`func (o *SendDiscordDirectMessageRequest) GetAttachments() []SendDiscordDirectMessageRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *SendDiscordDirectMessageRequest) GetAttachmentsOk() (*[]SendDiscordDirectMessageRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *SendDiscordDirectMessageRequest) SetAttachments(v []SendDiscordDirectMessageRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *SendDiscordDirectMessageRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetTts

`func (o *SendDiscordDirectMessageRequest) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *SendDiscordDirectMessageRequest) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *SendDiscordDirectMessageRequest) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *SendDiscordDirectMessageRequest) HasTts() bool`

HasTts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


