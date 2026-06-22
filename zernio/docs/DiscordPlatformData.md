# DiscordPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChannelId** | **string** | Target channel snowflake ID. Determines which channel in the connected server receives the message. | 
**Embeds** | Pointer to [**[]DiscordPlatformDataEmbedsInner**](DiscordPlatformDataEmbedsInner.md) | Up to 10 Discord embed objects (combined max 6,000 characters across all embeds). Sent alongside or instead of plain-text content. | [optional] 
**Poll** | Pointer to [**DiscordPlatformDataPoll**](DiscordPlatformDataPoll.md) |  | [optional] 
**Crosspost** | Pointer to **bool** | Auto-crosspost to every server following this announcement channel (type 5). No-op for regular text channels. | [optional] 
**ForumThreadName** | Pointer to **string** | Thread title for forum channel posts (type 15). Required when posting to a forum channel. | [optional] 
**ForumAppliedTags** | Pointer to **[]string** | Tag snowflake IDs to apply to forum posts. Max 5 tags. | [optional] 
**ThreadFromMessage** | Pointer to [**DiscordPlatformDataThreadFromMessage**](DiscordPlatformDataThreadFromMessage.md) |  | [optional] 
**Tts** | Pointer to **bool** | Send as text-to-speech message. Discord reads the message aloud in the channel. | [optional] 
**WebhookUsername** | Pointer to **string** | Override the webhook display name for this post only (1-80 chars). Falls back to the account-level default set via PATCH /v1/connect/discord. | [optional] 
**WebhookAvatarUrl** | Pointer to **string** | Override the webhook avatar URL for this post only. Falls back to the account-level default. | [optional] 

## Methods

### NewDiscordPlatformData

`func NewDiscordPlatformData(channelId string, ) *DiscordPlatformData`

NewDiscordPlatformData instantiates a new DiscordPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordPlatformDataWithDefaults

`func NewDiscordPlatformDataWithDefaults() *DiscordPlatformData`

NewDiscordPlatformDataWithDefaults instantiates a new DiscordPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannelId

`func (o *DiscordPlatformData) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *DiscordPlatformData) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *DiscordPlatformData) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetEmbeds

`func (o *DiscordPlatformData) GetEmbeds() []DiscordPlatformDataEmbedsInner`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *DiscordPlatformData) GetEmbedsOk() (*[]DiscordPlatformDataEmbedsInner, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *DiscordPlatformData) SetEmbeds(v []DiscordPlatformDataEmbedsInner)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *DiscordPlatformData) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetPoll

`func (o *DiscordPlatformData) GetPoll() DiscordPlatformDataPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *DiscordPlatformData) GetPollOk() (*DiscordPlatformDataPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *DiscordPlatformData) SetPoll(v DiscordPlatformDataPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *DiscordPlatformData) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetCrosspost

`func (o *DiscordPlatformData) GetCrosspost() bool`

GetCrosspost returns the Crosspost field if non-nil, zero value otherwise.

### GetCrosspostOk

`func (o *DiscordPlatformData) GetCrosspostOk() (*bool, bool)`

GetCrosspostOk returns a tuple with the Crosspost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosspost

`func (o *DiscordPlatformData) SetCrosspost(v bool)`

SetCrosspost sets Crosspost field to given value.

### HasCrosspost

`func (o *DiscordPlatformData) HasCrosspost() bool`

HasCrosspost returns a boolean if a field has been set.

### GetForumThreadName

`func (o *DiscordPlatformData) GetForumThreadName() string`

GetForumThreadName returns the ForumThreadName field if non-nil, zero value otherwise.

### GetForumThreadNameOk

`func (o *DiscordPlatformData) GetForumThreadNameOk() (*string, bool)`

GetForumThreadNameOk returns a tuple with the ForumThreadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumThreadName

`func (o *DiscordPlatformData) SetForumThreadName(v string)`

SetForumThreadName sets ForumThreadName field to given value.

### HasForumThreadName

`func (o *DiscordPlatformData) HasForumThreadName() bool`

HasForumThreadName returns a boolean if a field has been set.

### GetForumAppliedTags

`func (o *DiscordPlatformData) GetForumAppliedTags() []string`

GetForumAppliedTags returns the ForumAppliedTags field if non-nil, zero value otherwise.

### GetForumAppliedTagsOk

`func (o *DiscordPlatformData) GetForumAppliedTagsOk() (*[]string, bool)`

GetForumAppliedTagsOk returns a tuple with the ForumAppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumAppliedTags

`func (o *DiscordPlatformData) SetForumAppliedTags(v []string)`

SetForumAppliedTags sets ForumAppliedTags field to given value.

### HasForumAppliedTags

`func (o *DiscordPlatformData) HasForumAppliedTags() bool`

HasForumAppliedTags returns a boolean if a field has been set.

### GetThreadFromMessage

`func (o *DiscordPlatformData) GetThreadFromMessage() DiscordPlatformDataThreadFromMessage`

GetThreadFromMessage returns the ThreadFromMessage field if non-nil, zero value otherwise.

### GetThreadFromMessageOk

`func (o *DiscordPlatformData) GetThreadFromMessageOk() (*DiscordPlatformDataThreadFromMessage, bool)`

GetThreadFromMessageOk returns a tuple with the ThreadFromMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadFromMessage

`func (o *DiscordPlatformData) SetThreadFromMessage(v DiscordPlatformDataThreadFromMessage)`

SetThreadFromMessage sets ThreadFromMessage field to given value.

### HasThreadFromMessage

`func (o *DiscordPlatformData) HasThreadFromMessage() bool`

HasThreadFromMessage returns a boolean if a field has been set.

### GetTts

`func (o *DiscordPlatformData) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *DiscordPlatformData) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *DiscordPlatformData) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *DiscordPlatformData) HasTts() bool`

HasTts returns a boolean if a field has been set.

### GetWebhookUsername

`func (o *DiscordPlatformData) GetWebhookUsername() string`

GetWebhookUsername returns the WebhookUsername field if non-nil, zero value otherwise.

### GetWebhookUsernameOk

`func (o *DiscordPlatformData) GetWebhookUsernameOk() (*string, bool)`

GetWebhookUsernameOk returns a tuple with the WebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUsername

`func (o *DiscordPlatformData) SetWebhookUsername(v string)`

SetWebhookUsername sets WebhookUsername field to given value.

### HasWebhookUsername

`func (o *DiscordPlatformData) HasWebhookUsername() bool`

HasWebhookUsername returns a boolean if a field has been set.

### GetWebhookAvatarUrl

`func (o *DiscordPlatformData) GetWebhookAvatarUrl() string`

GetWebhookAvatarUrl returns the WebhookAvatarUrl field if non-nil, zero value otherwise.

### GetWebhookAvatarUrlOk

`func (o *DiscordPlatformData) GetWebhookAvatarUrlOk() (*string, bool)`

GetWebhookAvatarUrlOk returns a tuple with the WebhookAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAvatarUrl

`func (o *DiscordPlatformData) SetWebhookAvatarUrl(v string)`

SetWebhookAvatarUrl sets WebhookAvatarUrl field to given value.

### HasWebhookAvatarUrl

`func (o *DiscordPlatformData) HasWebhookAvatarUrl() bool`

HasWebhookAvatarUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


