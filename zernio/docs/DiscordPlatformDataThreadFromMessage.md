# DiscordPlatformDataThreadFromMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Thread name (1-100 chars) | [optional] 
**AutoArchiveDuration** | Pointer to **int32** | Auto-archive after inactivity (minutes) | [optional] 
**RateLimitPerUser** | Pointer to **int32** | Slow-mode duration in seconds (0-21600) | [optional] 

## Methods

### NewDiscordPlatformDataThreadFromMessage

`func NewDiscordPlatformDataThreadFromMessage() *DiscordPlatformDataThreadFromMessage`

NewDiscordPlatformDataThreadFromMessage instantiates a new DiscordPlatformDataThreadFromMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordPlatformDataThreadFromMessageWithDefaults

`func NewDiscordPlatformDataThreadFromMessageWithDefaults() *DiscordPlatformDataThreadFromMessage`

NewDiscordPlatformDataThreadFromMessageWithDefaults instantiates a new DiscordPlatformDataThreadFromMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DiscordPlatformDataThreadFromMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscordPlatformDataThreadFromMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscordPlatformDataThreadFromMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscordPlatformDataThreadFromMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAutoArchiveDuration

`func (o *DiscordPlatformDataThreadFromMessage) GetAutoArchiveDuration() int32`

GetAutoArchiveDuration returns the AutoArchiveDuration field if non-nil, zero value otherwise.

### GetAutoArchiveDurationOk

`func (o *DiscordPlatformDataThreadFromMessage) GetAutoArchiveDurationOk() (*int32, bool)`

GetAutoArchiveDurationOk returns a tuple with the AutoArchiveDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoArchiveDuration

`func (o *DiscordPlatformDataThreadFromMessage) SetAutoArchiveDuration(v int32)`

SetAutoArchiveDuration sets AutoArchiveDuration field to given value.

### HasAutoArchiveDuration

`func (o *DiscordPlatformDataThreadFromMessage) HasAutoArchiveDuration() bool`

HasAutoArchiveDuration returns a boolean if a field has been set.

### GetRateLimitPerUser

`func (o *DiscordPlatformDataThreadFromMessage) GetRateLimitPerUser() int32`

GetRateLimitPerUser returns the RateLimitPerUser field if non-nil, zero value otherwise.

### GetRateLimitPerUserOk

`func (o *DiscordPlatformDataThreadFromMessage) GetRateLimitPerUserOk() (*int32, bool)`

GetRateLimitPerUserOk returns a tuple with the RateLimitPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitPerUser

`func (o *DiscordPlatformDataThreadFromMessage) SetRateLimitPerUser(v int32)`

SetRateLimitPerUser sets RateLimitPerUser field to given value.

### HasRateLimitPerUser

`func (o *DiscordPlatformDataThreadFromMessage) HasRateLimitPerUser() bool`

HasRateLimitPerUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


