# CreatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** | Post caption/text. Optional when media is attached or all platforms have customContent. Required for text-only posts. | [optional] 
**MediaItems** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 
**Platforms** | Pointer to [**[]CreatePostRequestPlatformsInner**](CreatePostRequestPlatformsInner.md) | Target platforms and accounts for this post. Required for non-draft posts (returns 400 if empty). Drafts can omit platforms. | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**PublishNow** | Pointer to **bool** |  | [optional] [default to false]
**IsDraft** | Pointer to **bool** | When true, saves the post as a draft. When none of scheduledFor, publishNow, or queuedFromProfile are provided, the post defaults to draft automatically. | [optional] [default to false]
**Timezone** | Pointer to **string** |  | [optional] [default to "UTC"]
**Tags** | Pointer to **[]string** | Tags/keywords. YouTube constraints: each tag max 100 chars, combined max 500 chars, duplicates auto-removed. | [optional] 
**Hashtags** | Pointer to **[]string** |  | [optional] 
**Mentions** | Pointer to **[]string** | Stored for reference only. This field does NOT automatically create @mentions when publishing. For LinkedIn @mentions, use the /v1/accounts/{accountId}/linkedin-mentions endpoint to resolve profile URLs to URNs, then embed the returned mentionFormat directly in the post content field. | [optional] 
**CrosspostingEnabled** | Pointer to **bool** |  | [optional] [default to true]
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**TiktokSettings** | Pointer to [**TikTokPlatformData**](TikTokPlatformData.md) | Root-level TikTok settings applied to all TikTok platforms. Merged into each platform&#39;s platformSpecificData, with platform-specific settings taking precedence. | [optional] 
**FacebookSettings** | Pointer to [**FacebookPlatformData**](FacebookPlatformData.md) | Root-level Facebook settings applied to all Facebook platforms. Merged into each platform&#39;s platformSpecificData, with platform-specific settings taking precedence. | [optional] 
**Recycling** | Pointer to [**RecyclingConfig**](RecyclingConfig.md) |  | [optional] 
**QueuedFromProfile** | Pointer to **string** | Profile ID to schedule via queue. When provided without scheduledFor, the post is auto-assigned to the next available slot. Do not call /v1/queue/next-slot and use that time in scheduledFor, as that bypasses queue locking. | [optional] 
**QueueId** | Pointer to **string** | Specific queue ID to use when scheduling via queue. Only used when queuedFromProfile is also provided. If omitted, uses the profile&#39;s default queue.  | [optional] 

## Methods

### NewCreatePostRequest

`func NewCreatePostRequest() *CreatePostRequest`

NewCreatePostRequest instantiates a new CreatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePostRequestWithDefaults

`func NewCreatePostRequestWithDefaults() *CreatePostRequest`

NewCreatePostRequestWithDefaults instantiates a new CreatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreatePostRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreatePostRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreatePostRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreatePostRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *CreatePostRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CreatePostRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CreatePostRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *CreatePostRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMediaItems

`func (o *CreatePostRequest) GetMediaItems() []MediaItem`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *CreatePostRequest) GetMediaItemsOk() (*[]MediaItem, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *CreatePostRequest) SetMediaItems(v []MediaItem)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *CreatePostRequest) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.

### GetPlatforms

`func (o *CreatePostRequest) GetPlatforms() []CreatePostRequestPlatformsInner`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *CreatePostRequest) GetPlatformsOk() (*[]CreatePostRequestPlatformsInner, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *CreatePostRequest) SetPlatforms(v []CreatePostRequestPlatformsInner)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *CreatePostRequest) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetScheduledFor

`func (o *CreatePostRequest) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *CreatePostRequest) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *CreatePostRequest) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *CreatePostRequest) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPublishNow

`func (o *CreatePostRequest) GetPublishNow() bool`

GetPublishNow returns the PublishNow field if non-nil, zero value otherwise.

### GetPublishNowOk

`func (o *CreatePostRequest) GetPublishNowOk() (*bool, bool)`

GetPublishNowOk returns a tuple with the PublishNow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishNow

`func (o *CreatePostRequest) SetPublishNow(v bool)`

SetPublishNow sets PublishNow field to given value.

### HasPublishNow

`func (o *CreatePostRequest) HasPublishNow() bool`

HasPublishNow returns a boolean if a field has been set.

### GetIsDraft

`func (o *CreatePostRequest) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *CreatePostRequest) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *CreatePostRequest) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.

### HasIsDraft

`func (o *CreatePostRequest) HasIsDraft() bool`

HasIsDraft returns a boolean if a field has been set.

### GetTimezone

`func (o *CreatePostRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreatePostRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreatePostRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *CreatePostRequest) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetTags

`func (o *CreatePostRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreatePostRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreatePostRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreatePostRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHashtags

`func (o *CreatePostRequest) GetHashtags() []string`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *CreatePostRequest) GetHashtagsOk() (*[]string, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *CreatePostRequest) SetHashtags(v []string)`

SetHashtags sets Hashtags field to given value.

### HasHashtags

`func (o *CreatePostRequest) HasHashtags() bool`

HasHashtags returns a boolean if a field has been set.

### GetMentions

`func (o *CreatePostRequest) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *CreatePostRequest) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *CreatePostRequest) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *CreatePostRequest) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetCrosspostingEnabled

`func (o *CreatePostRequest) GetCrosspostingEnabled() bool`

GetCrosspostingEnabled returns the CrosspostingEnabled field if non-nil, zero value otherwise.

### GetCrosspostingEnabledOk

`func (o *CreatePostRequest) GetCrosspostingEnabledOk() (*bool, bool)`

GetCrosspostingEnabledOk returns a tuple with the CrosspostingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosspostingEnabled

`func (o *CreatePostRequest) SetCrosspostingEnabled(v bool)`

SetCrosspostingEnabled sets CrosspostingEnabled field to given value.

### HasCrosspostingEnabled

`func (o *CreatePostRequest) HasCrosspostingEnabled() bool`

HasCrosspostingEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *CreatePostRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreatePostRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreatePostRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreatePostRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetTiktokSettings

`func (o *CreatePostRequest) GetTiktokSettings() TikTokPlatformData`

GetTiktokSettings returns the TiktokSettings field if non-nil, zero value otherwise.

### GetTiktokSettingsOk

`func (o *CreatePostRequest) GetTiktokSettingsOk() (*TikTokPlatformData, bool)`

GetTiktokSettingsOk returns a tuple with the TiktokSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokSettings

`func (o *CreatePostRequest) SetTiktokSettings(v TikTokPlatformData)`

SetTiktokSettings sets TiktokSettings field to given value.

### HasTiktokSettings

`func (o *CreatePostRequest) HasTiktokSettings() bool`

HasTiktokSettings returns a boolean if a field has been set.

### GetFacebookSettings

`func (o *CreatePostRequest) GetFacebookSettings() FacebookPlatformData`

GetFacebookSettings returns the FacebookSettings field if non-nil, zero value otherwise.

### GetFacebookSettingsOk

`func (o *CreatePostRequest) GetFacebookSettingsOk() (*FacebookPlatformData, bool)`

GetFacebookSettingsOk returns a tuple with the FacebookSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookSettings

`func (o *CreatePostRequest) SetFacebookSettings(v FacebookPlatformData)`

SetFacebookSettings sets FacebookSettings field to given value.

### HasFacebookSettings

`func (o *CreatePostRequest) HasFacebookSettings() bool`

HasFacebookSettings returns a boolean if a field has been set.

### GetRecycling

`func (o *CreatePostRequest) GetRecycling() RecyclingConfig`

GetRecycling returns the Recycling field if non-nil, zero value otherwise.

### GetRecyclingOk

`func (o *CreatePostRequest) GetRecyclingOk() (*RecyclingConfig, bool)`

GetRecyclingOk returns a tuple with the Recycling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycling

`func (o *CreatePostRequest) SetRecycling(v RecyclingConfig)`

SetRecycling sets Recycling field to given value.

### HasRecycling

`func (o *CreatePostRequest) HasRecycling() bool`

HasRecycling returns a boolean if a field has been set.

### GetQueuedFromProfile

`func (o *CreatePostRequest) GetQueuedFromProfile() string`

GetQueuedFromProfile returns the QueuedFromProfile field if non-nil, zero value otherwise.

### GetQueuedFromProfileOk

`func (o *CreatePostRequest) GetQueuedFromProfileOk() (*string, bool)`

GetQueuedFromProfileOk returns a tuple with the QueuedFromProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedFromProfile

`func (o *CreatePostRequest) SetQueuedFromProfile(v string)`

SetQueuedFromProfile sets QueuedFromProfile field to given value.

### HasQueuedFromProfile

`func (o *CreatePostRequest) HasQueuedFromProfile() bool`

HasQueuedFromProfile returns a boolean if a field has been set.

### GetQueueId

`func (o *CreatePostRequest) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *CreatePostRequest) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *CreatePostRequest) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *CreatePostRequest) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


