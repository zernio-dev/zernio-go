# UpdateDiscordSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebhookUsername** | Pointer to **string** | Custom display name for the webhook (1-80 chars). Empty string resets to default (\&quot;Zernio\&quot;). Cannot contain \&quot;clyde\&quot; or \&quot;discord\&quot;. | [optional] 
**WebhookAvatarUrl** | Pointer to **string** | Custom avatar URL. Empty string resets to default bot avatar. | [optional] 
**ChannelId** | Pointer to **string** | Switch to a different channel in the same guild. Must be a text (0), announcement (5), or forum (15) channel. | [optional] 

## Methods

### NewUpdateDiscordSettingsRequest

`func NewUpdateDiscordSettingsRequest() *UpdateDiscordSettingsRequest`

NewUpdateDiscordSettingsRequest instantiates a new UpdateDiscordSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDiscordSettingsRequestWithDefaults

`func NewUpdateDiscordSettingsRequestWithDefaults() *UpdateDiscordSettingsRequest`

NewUpdateDiscordSettingsRequestWithDefaults instantiates a new UpdateDiscordSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhookUsername

`func (o *UpdateDiscordSettingsRequest) GetWebhookUsername() string`

GetWebhookUsername returns the WebhookUsername field if non-nil, zero value otherwise.

### GetWebhookUsernameOk

`func (o *UpdateDiscordSettingsRequest) GetWebhookUsernameOk() (*string, bool)`

GetWebhookUsernameOk returns a tuple with the WebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUsername

`func (o *UpdateDiscordSettingsRequest) SetWebhookUsername(v string)`

SetWebhookUsername sets WebhookUsername field to given value.

### HasWebhookUsername

`func (o *UpdateDiscordSettingsRequest) HasWebhookUsername() bool`

HasWebhookUsername returns a boolean if a field has been set.

### GetWebhookAvatarUrl

`func (o *UpdateDiscordSettingsRequest) GetWebhookAvatarUrl() string`

GetWebhookAvatarUrl returns the WebhookAvatarUrl field if non-nil, zero value otherwise.

### GetWebhookAvatarUrlOk

`func (o *UpdateDiscordSettingsRequest) GetWebhookAvatarUrlOk() (*string, bool)`

GetWebhookAvatarUrlOk returns a tuple with the WebhookAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAvatarUrl

`func (o *UpdateDiscordSettingsRequest) SetWebhookAvatarUrl(v string)`

SetWebhookAvatarUrl sets WebhookAvatarUrl field to given value.

### HasWebhookAvatarUrl

`func (o *UpdateDiscordSettingsRequest) HasWebhookAvatarUrl() bool`

HasWebhookAvatarUrl returns a boolean if a field has been set.

### GetChannelId

`func (o *UpdateDiscordSettingsRequest) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *UpdateDiscordSettingsRequest) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *UpdateDiscordSettingsRequest) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *UpdateDiscordSettingsRequest) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


