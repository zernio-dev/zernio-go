# DiscordScheduledEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Event snowflake ID | [optional] 
**GuildId** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **NullableString** | Voice/stage channel ID; null for external events. | [optional] 
**CreatorId** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ScheduledStartTime** | Pointer to **time.Time** |  | [optional] 
**ScheduledEndTime** | Pointer to **NullableTime** | Required for external events; optional for voice/stage. | [optional] 
**PrivacyLevel** | Pointer to **int32** | Always 2 (GUILD_ONLY) — Discord deprecated PUBLIC events. | [optional] 
**Status** | Pointer to **int32** | 1&#x3D;SCHEDULED, 2&#x3D;ACTIVE, 3&#x3D;COMPLETED, 4&#x3D;CANCELED | [optional] 
**EntityType** | Pointer to **int32** | 1&#x3D;STAGE_INSTANCE, 2&#x3D;VOICE, 3&#x3D;EXTERNAL | [optional] 
**EntityId** | Pointer to **NullableString** |  | [optional] 
**EntityMetadata** | Pointer to [**DiscordScheduledEventEntityMetadata**](DiscordScheduledEventEntityMetadata.md) |  | [optional] 
**UserCount** | Pointer to **int32** | Number of members who RSVP&#39;d. Only present when withUserCount&#x3D;true on list. | [optional] 
**Image** | Pointer to **NullableString** | Cover image hash; build URL via cdn.discordapp.com. | [optional] 

## Methods

### NewDiscordScheduledEvent

`func NewDiscordScheduledEvent() *DiscordScheduledEvent`

NewDiscordScheduledEvent instantiates a new DiscordScheduledEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordScheduledEventWithDefaults

`func NewDiscordScheduledEventWithDefaults() *DiscordScheduledEvent`

NewDiscordScheduledEventWithDefaults instantiates a new DiscordScheduledEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DiscordScheduledEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DiscordScheduledEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DiscordScheduledEvent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DiscordScheduledEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetGuildId

`func (o *DiscordScheduledEvent) GetGuildId() string`

GetGuildId returns the GuildId field if non-nil, zero value otherwise.

### GetGuildIdOk

`func (o *DiscordScheduledEvent) GetGuildIdOk() (*string, bool)`

GetGuildIdOk returns a tuple with the GuildId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuildId

`func (o *DiscordScheduledEvent) SetGuildId(v string)`

SetGuildId sets GuildId field to given value.

### HasGuildId

`func (o *DiscordScheduledEvent) HasGuildId() bool`

HasGuildId returns a boolean if a field has been set.

### GetChannelId

`func (o *DiscordScheduledEvent) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *DiscordScheduledEvent) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *DiscordScheduledEvent) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *DiscordScheduledEvent) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### SetChannelIdNil

`func (o *DiscordScheduledEvent) SetChannelIdNil(b bool)`

 SetChannelIdNil sets the value for ChannelId to be an explicit nil

### UnsetChannelId
`func (o *DiscordScheduledEvent) UnsetChannelId()`

UnsetChannelId ensures that no value is present for ChannelId, not even an explicit nil
### GetCreatorId

`func (o *DiscordScheduledEvent) GetCreatorId() string`

GetCreatorId returns the CreatorId field if non-nil, zero value otherwise.

### GetCreatorIdOk

`func (o *DiscordScheduledEvent) GetCreatorIdOk() (*string, bool)`

GetCreatorIdOk returns a tuple with the CreatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatorId

`func (o *DiscordScheduledEvent) SetCreatorId(v string)`

SetCreatorId sets CreatorId field to given value.

### HasCreatorId

`func (o *DiscordScheduledEvent) HasCreatorId() bool`

HasCreatorId returns a boolean if a field has been set.

### SetCreatorIdNil

`func (o *DiscordScheduledEvent) SetCreatorIdNil(b bool)`

 SetCreatorIdNil sets the value for CreatorId to be an explicit nil

### UnsetCreatorId
`func (o *DiscordScheduledEvent) UnsetCreatorId()`

UnsetCreatorId ensures that no value is present for CreatorId, not even an explicit nil
### GetName

`func (o *DiscordScheduledEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiscordScheduledEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiscordScheduledEvent) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiscordScheduledEvent) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DiscordScheduledEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DiscordScheduledEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DiscordScheduledEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DiscordScheduledEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *DiscordScheduledEvent) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *DiscordScheduledEvent) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetScheduledStartTime

`func (o *DiscordScheduledEvent) GetScheduledStartTime() time.Time`

GetScheduledStartTime returns the ScheduledStartTime field if non-nil, zero value otherwise.

### GetScheduledStartTimeOk

`func (o *DiscordScheduledEvent) GetScheduledStartTimeOk() (*time.Time, bool)`

GetScheduledStartTimeOk returns a tuple with the ScheduledStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledStartTime

`func (o *DiscordScheduledEvent) SetScheduledStartTime(v time.Time)`

SetScheduledStartTime sets ScheduledStartTime field to given value.

### HasScheduledStartTime

`func (o *DiscordScheduledEvent) HasScheduledStartTime() bool`

HasScheduledStartTime returns a boolean if a field has been set.

### GetScheduledEndTime

`func (o *DiscordScheduledEvent) GetScheduledEndTime() time.Time`

GetScheduledEndTime returns the ScheduledEndTime field if non-nil, zero value otherwise.

### GetScheduledEndTimeOk

`func (o *DiscordScheduledEvent) GetScheduledEndTimeOk() (*time.Time, bool)`

GetScheduledEndTimeOk returns a tuple with the ScheduledEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledEndTime

`func (o *DiscordScheduledEvent) SetScheduledEndTime(v time.Time)`

SetScheduledEndTime sets ScheduledEndTime field to given value.

### HasScheduledEndTime

`func (o *DiscordScheduledEvent) HasScheduledEndTime() bool`

HasScheduledEndTime returns a boolean if a field has been set.

### SetScheduledEndTimeNil

`func (o *DiscordScheduledEvent) SetScheduledEndTimeNil(b bool)`

 SetScheduledEndTimeNil sets the value for ScheduledEndTime to be an explicit nil

### UnsetScheduledEndTime
`func (o *DiscordScheduledEvent) UnsetScheduledEndTime()`

UnsetScheduledEndTime ensures that no value is present for ScheduledEndTime, not even an explicit nil
### GetPrivacyLevel

`func (o *DiscordScheduledEvent) GetPrivacyLevel() int32`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *DiscordScheduledEvent) GetPrivacyLevelOk() (*int32, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *DiscordScheduledEvent) SetPrivacyLevel(v int32)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *DiscordScheduledEvent) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### GetStatus

`func (o *DiscordScheduledEvent) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DiscordScheduledEvent) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DiscordScheduledEvent) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DiscordScheduledEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntityType

`func (o *DiscordScheduledEvent) GetEntityType() int32`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *DiscordScheduledEvent) GetEntityTypeOk() (*int32, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *DiscordScheduledEvent) SetEntityType(v int32)`

SetEntityType sets EntityType field to given value.

### HasEntityType

`func (o *DiscordScheduledEvent) HasEntityType() bool`

HasEntityType returns a boolean if a field has been set.

### GetEntityId

`func (o *DiscordScheduledEvent) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *DiscordScheduledEvent) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *DiscordScheduledEvent) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *DiscordScheduledEvent) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.

### SetEntityIdNil

`func (o *DiscordScheduledEvent) SetEntityIdNil(b bool)`

 SetEntityIdNil sets the value for EntityId to be an explicit nil

### UnsetEntityId
`func (o *DiscordScheduledEvent) UnsetEntityId()`

UnsetEntityId ensures that no value is present for EntityId, not even an explicit nil
### GetEntityMetadata

`func (o *DiscordScheduledEvent) GetEntityMetadata() DiscordScheduledEventEntityMetadata`

GetEntityMetadata returns the EntityMetadata field if non-nil, zero value otherwise.

### GetEntityMetadataOk

`func (o *DiscordScheduledEvent) GetEntityMetadataOk() (*DiscordScheduledEventEntityMetadata, bool)`

GetEntityMetadataOk returns a tuple with the EntityMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityMetadata

`func (o *DiscordScheduledEvent) SetEntityMetadata(v DiscordScheduledEventEntityMetadata)`

SetEntityMetadata sets EntityMetadata field to given value.

### HasEntityMetadata

`func (o *DiscordScheduledEvent) HasEntityMetadata() bool`

HasEntityMetadata returns a boolean if a field has been set.

### GetUserCount

`func (o *DiscordScheduledEvent) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *DiscordScheduledEvent) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *DiscordScheduledEvent) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *DiscordScheduledEvent) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetImage

`func (o *DiscordScheduledEvent) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *DiscordScheduledEvent) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *DiscordScheduledEvent) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *DiscordScheduledEvent) HasImage() bool`

HasImage returns a boolean if a field has been set.

### SetImageNil

`func (o *DiscordScheduledEvent) SetImageNil(b bool)`

 SetImageNil sets the value for Image to be an explicit nil

### UnsetImage
`func (o *DiscordScheduledEvent) UnsetImage()`

UnsetImage ensures that no value is present for Image, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


