# GetDiscordSettings200ResponseAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** | Channel name | [optional] 
**DisplayName** | Pointer to **string** | Guild - #channel display name | [optional] 
**ProfilePicture** | Pointer to **string** | Guild icon URL | [optional] 
**ChannelId** | Pointer to **string** | Connected channel snowflake ID | [optional] 
**ChannelName** | Pointer to **string** | Channel name | [optional] 
**ChannelType** | Pointer to **string** | Channel type (0 &#x3D; text, 5 &#x3D; announcement, 15 &#x3D; forum) | [optional] 
**GuildId** | Pointer to **string** | Guild (server) snowflake ID | [optional] 
**WebhookUsername** | Pointer to **NullableString** | Custom webhook display name (null &#x3D; default \&quot;Zernio\&quot;) | [optional] 
**WebhookAvatarUrl** | Pointer to **NullableString** | Custom webhook avatar URL (null &#x3D; default bot avatar) | [optional] 

## Methods

### NewGetDiscordSettings200ResponseAccount

`func NewGetDiscordSettings200ResponseAccount() *GetDiscordSettings200ResponseAccount`

NewGetDiscordSettings200ResponseAccount instantiates a new GetDiscordSettings200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscordSettings200ResponseAccountWithDefaults

`func NewGetDiscordSettings200ResponseAccountWithDefaults() *GetDiscordSettings200ResponseAccount`

NewGetDiscordSettings200ResponseAccountWithDefaults instantiates a new GetDiscordSettings200ResponseAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDiscordSettings200ResponseAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDiscordSettings200ResponseAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDiscordSettings200ResponseAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDiscordSettings200ResponseAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetDiscordSettings200ResponseAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetDiscordSettings200ResponseAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetDiscordSettings200ResponseAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetDiscordSettings200ResponseAccount) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUsername

`func (o *GetDiscordSettings200ResponseAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetDiscordSettings200ResponseAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetDiscordSettings200ResponseAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetDiscordSettings200ResponseAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetDiscordSettings200ResponseAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetDiscordSettings200ResponseAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetDiscordSettings200ResponseAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetDiscordSettings200ResponseAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProfilePicture

`func (o *GetDiscordSettings200ResponseAccount) GetProfilePicture() string`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *GetDiscordSettings200ResponseAccount) GetProfilePictureOk() (*string, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *GetDiscordSettings200ResponseAccount) SetProfilePicture(v string)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *GetDiscordSettings200ResponseAccount) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### GetChannelId

`func (o *GetDiscordSettings200ResponseAccount) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GetDiscordSettings200ResponseAccount) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GetDiscordSettings200ResponseAccount) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *GetDiscordSettings200ResponseAccount) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannelName

`func (o *GetDiscordSettings200ResponseAccount) GetChannelName() string`

GetChannelName returns the ChannelName field if non-nil, zero value otherwise.

### GetChannelNameOk

`func (o *GetDiscordSettings200ResponseAccount) GetChannelNameOk() (*string, bool)`

GetChannelNameOk returns a tuple with the ChannelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelName

`func (o *GetDiscordSettings200ResponseAccount) SetChannelName(v string)`

SetChannelName sets ChannelName field to given value.

### HasChannelName

`func (o *GetDiscordSettings200ResponseAccount) HasChannelName() bool`

HasChannelName returns a boolean if a field has been set.

### GetChannelType

`func (o *GetDiscordSettings200ResponseAccount) GetChannelType() string`

GetChannelType returns the ChannelType field if non-nil, zero value otherwise.

### GetChannelTypeOk

`func (o *GetDiscordSettings200ResponseAccount) GetChannelTypeOk() (*string, bool)`

GetChannelTypeOk returns a tuple with the ChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelType

`func (o *GetDiscordSettings200ResponseAccount) SetChannelType(v string)`

SetChannelType sets ChannelType field to given value.

### HasChannelType

`func (o *GetDiscordSettings200ResponseAccount) HasChannelType() bool`

HasChannelType returns a boolean if a field has been set.

### GetGuildId

`func (o *GetDiscordSettings200ResponseAccount) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *GetDiscordSettings200ResponseAccount) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *GetDiscordSettings200ResponseAccount) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *GetDiscordSettings200ResponseAccount) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetWebhookUsername

`func (o *GetDiscordSettings200ResponseAccount) GetWebhookUsername() string`

GetWebhookUsername returns the WebhookUsername field if non-nil, zero value otherwise.

### GetWebhookUsernameOk

`func (o *GetDiscordSettings200ResponseAccount) GetWebhookUsernameOk() (*string, bool)`

GetWebhookUsernameOk returns a tuple with the WebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUsername

`func (o *GetDiscordSettings200ResponseAccount) SetWebhookUsername(v string)`

SetWebhookUsername sets WebhookUsername field to given value.

### HasWebhookUsername

`func (o *GetDiscordSettings200ResponseAccount) HasWebhookUsername() bool`

HasWebhookUsername returns a boolean if a field has been set.

### SetWebhookUsernameNil

`func (o *GetDiscordSettings200ResponseAccount) SetWebhookUsernameNil(b bool)`

 SetWebhookUsernameNil sets the value for WebhookUsername to be an explicit nil

### UnsetWebhookUsername
`func (o *GetDiscordSettings200ResponseAccount) UnsetWebhookUsername()`

UnsetWebhookUsername ensures that no value is present for WebhookUsername, not even an explicit nil
### GetWebhookAvatarUrl

`func (o *GetDiscordSettings200ResponseAccount) GetWebhookAvatarUrl() string`

GetWebhookAvatarUrl returns the WebhookAvatarUrl field if non-nil, zero value otherwise.

### GetWebhookAvatarUrlOk

`func (o *GetDiscordSettings200ResponseAccount) GetWebhookAvatarUrlOk() (*string, bool)`

GetWebhookAvatarUrlOk returns a tuple with the WebhookAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAvatarUrl

`func (o *GetDiscordSettings200ResponseAccount) SetWebhookAvatarUrl(v string)`

SetWebhookAvatarUrl sets WebhookAvatarUrl field to given value.

### HasWebhookAvatarUrl

`func (o *GetDiscordSettings200ResponseAccount) HasWebhookAvatarUrl() bool`

HasWebhookAvatarUrl returns a boolean if a field has been set.

### SetWebhookAvatarUrlNil

`func (o *GetDiscordSettings200ResponseAccount) SetWebhookAvatarUrlNil(b bool)`

 SetWebhookAvatarUrlNil sets the value for WebhookAvatarUrl to be an explicit nil

### UnsetWebhookAvatarUrl
`func (o *GetDiscordSettings200ResponseAccount) UnsetWebhookAvatarUrl()`

UnsetWebhookAvatarUrl ensures that no value is present for WebhookAvatarUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


