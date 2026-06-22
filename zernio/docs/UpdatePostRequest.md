# UpdatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**MediaItems** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 
**Platforms** | Pointer to [**[]UpdatePostRequestPlatformsInner**](UpdatePostRequestPlatformsInner.md) | Target platforms and accounts for this post. Each item must include platform and accountId. | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**PublishNow** | Pointer to **bool** |  | [optional] [default to false]
**IsDraft** | Pointer to **bool** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Hashtags** | Pointer to **[]string** |  | [optional] 
**Mentions** | Pointer to **[]string** |  | [optional] 
**CrosspostingEnabled** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**QueuedFromProfile** | Pointer to **string** | Profile ID to schedule via queue. | [optional] 
**QueueId** | Pointer to **string** | Specific queue ID to use when scheduling via queue. | [optional] 
**TiktokSettings** | Pointer to [**TikTokPlatformData**](TikTokPlatformData.md) | Root-level TikTok settings applied to all TikTok platforms. Merged into each platform&#39;s platformSpecificData, with platform-specific settings taking precedence. | [optional] 
**FacebookSettings** | Pointer to [**FacebookPlatformData**](FacebookPlatformData.md) | Root-level Facebook settings applied to all Facebook platforms. Merged into each platform&#39;s platformSpecificData, with platform-specific settings taking precedence. | [optional] 
**Recycling** | Pointer to [**RecyclingConfig**](RecyclingConfig.md) |  | [optional] 

## Methods

### NewUpdatePostRequest

`func NewUpdatePostRequest() *UpdatePostRequest`

NewUpdatePostRequest instantiates a new UpdatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostRequestWithDefaults

`func NewUpdatePostRequestWithDefaults() *UpdatePostRequest`

NewUpdatePostRequestWithDefaults instantiates a new UpdatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *UpdatePostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdatePostRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *UpdatePostRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *UpdatePostRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *UpdatePostRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *UpdatePostRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMediaItems

`func (o *UpdatePostRequest) GetMediaItems() []MediaItem`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *UpdatePostRequest) GetMediaItemsOk() (*[]MediaItem, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *UpdatePostRequest) SetMediaItems(v []MediaItem)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *UpdatePostRequest) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.

### GetPlatforms

`func (o *UpdatePostRequest) GetPlatforms() []UpdatePostRequestPlatformsInner`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *UpdatePostRequest) GetPlatformsOk() (*[]UpdatePostRequestPlatformsInner, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *UpdatePostRequest) SetPlatforms(v []UpdatePostRequestPlatformsInner)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *UpdatePostRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetScheduledFor

`func (o *UpdatePostRequest) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *UpdatePostRequest) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *UpdatePostRequest) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *UpdatePostRequest) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPublishNow

`func (o *UpdatePostRequest) GetPublishNow() bool`

GetPublishNow returns the PublishNow field if non-nil, zero value otherwise.

### GetPublishNowOk

`func (o *UpdatePostRequest) GetPublishNowOk() (*bool, bool)`

GetPublishNowOk returns a tuple with the PublishNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishNow

`func (o *UpdatePostRequest) SetPublishNow(v bool)`

SetPublishNow sets PublishNow field to given value.

### HasPublishNow

`func (o *UpdatePostRequest) HasPublishNow() bool`

HasPublishNow returns a boolean if a field has been set.

### GetIsDraft

`func (o *UpdatePostRequest) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *UpdatePostRequest) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *UpdatePostRequest) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *UpdatePostRequest) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdatePostRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdatePostRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdatePostRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UpdatePostRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdatePostRequest) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdatePostRequest) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdatePostRequest) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdatePostRequest) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePostRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePostRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePostRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHashtags

`func (o *UpdatePostRequest) GetHashtags() []string`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *UpdatePostRequest) GetHashtagsOk() (*[]string, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *UpdatePostRequest) SetHashtags(v []string)`

SetHashtags sets Hashtags field to given value.

### HasHashtags

`func (o *UpdatePostRequest) HasHashtags() bool`

HasHashtags returns a boolean if a field has been set.

### GetMentions

`func (o *UpdatePostRequest) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *UpdatePostRequest) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *UpdatePostRequest) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *UpdatePostRequest) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetCrosspostingEnabled

`func (o *UpdatePostRequest) GetCrosspostingEnabled() bool`

GetCrosspostingEnabled returns the CrosspostingEnabled field if non-nil, zero value otherwise.

### GetCrosspostingEnabledOk

`func (o *UpdatePostRequest) GetCrosspostingEnabledOk() (*bool, bool)`

GetCrosspostingEnabledOk returns a tuple with the CrosspostingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosspostingEnabled

`func (o *UpdatePostRequest) SetCrosspostingEnabled(v bool)`

SetCrosspostingEnabled sets CrosspostingEnabled field to given value.

### HasCrosspostingEnabled

`func (o *UpdatePostRequest) HasCrosspostingEnabled() bool`

HasCrosspostingEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *UpdatePostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdatePostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdatePostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdatePostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetQueuedFromProfile

`func (o *UpdatePostRequest) GetQueuedFromProfile() string`

GetQueuedFromProfile returns the QueuedFromProfile field if non-nil, zero value otherwise.

### GetQueuedFromProfileOk

`func (o *UpdatePostRequest) GetQueuedFromProfileOk() (*string, bool)`

GetQueuedFromProfileOk returns a tuple with the QueuedFromProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedFromProfile

`func (o *UpdatePostRequest) SetQueuedFromProfile(v string)`

SetQueuedFromProfile sets QueuedFromProfile field to given value.

### HasQueuedFromProfile

`func (o *UpdatePostRequest) HasQueuedFromProfile() bool`

HasQueuedFromProfile returns a boolean if a field has been set.

### GetQueueId

`func (o *UpdatePostRequest) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *UpdatePostRequest) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *UpdatePostRequest) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *UpdatePostRequest) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetTiktokSettings

`func (o *UpdatePostRequest) GetTiktokSettings() TikTokPlatformData`

GetTiktokSettings returns the TiktokSettings field if non-nil, zero value otherwise.

### GetTiktokSettingsOk

`func (o *UpdatePostRequest) GetTiktokSettingsOk() (*TikTokPlatformData, bool)`

GetTiktokSettingsOk returns a tuple with the TiktokSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokSettings

`func (o *UpdatePostRequest) SetTiktokSettings(v TikTokPlatformData)`

SetTiktokSettings sets TiktokSettings field to given value.

### HasTiktokSettings

`func (o *UpdatePostRequest) HasTiktokSettings() bool`

HasTiktokSettings returns a boolean if a field has been set.

### GetFacebookSettings

`func (o *UpdatePostRequest) GetFacebookSettings() FacebookPlatformData`

GetFacebookSettings returns the FacebookSettings field if non-nil, zero value otherwise.

### GetFacebookSettingsOk

`func (o *UpdatePostRequest) GetFacebookSettingsOk() (*FacebookPlatformData, bool)`

GetFacebookSettingsOk returns a tuple with the FacebookSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookSettings

`func (o *UpdatePostRequest) SetFacebookSettings(v FacebookPlatformData)`

SetFacebookSettings sets FacebookSettings field to given value.

### HasFacebookSettings

`func (o *UpdatePostRequest) HasFacebookSettings() bool`

HasFacebookSettings returns a boolean if a field has been set.

### GetRecycling

`func (o *UpdatePostRequest) GetRecycling() RecyclingConfig`

GetRecycling returns the Recycling field if non-nil, zero value otherwise.

### GetRecyclingOk

`func (o *UpdatePostRequest) GetRecyclingOk() (*RecyclingConfig, bool)`

GetRecyclingOk returns a tuple with the Recycling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycling

`func (o *UpdatePostRequest) SetRecycling(v RecyclingConfig)`

SetRecycling sets Recycling field to given value.

### HasRecycling

`func (o *UpdatePostRequest) HasRecycling() bool`

HasRecycling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


