# Post

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to [**PostUserId**](PostUserId.md) |  | [optional] 
**Title** | Pointer to **string** | YouTube: title must be ≤ 100 characters.  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**MediaItems** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 
**Platforms** | Pointer to [**[]PlatformTarget**](PlatformTarget.md) |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** | YouTube constraints: each tag max 100 chars, combined max 500 chars, duplicates removed. | [optional] 
**Hashtags** | Pointer to **[]string** |  | [optional] 
**Mentions** | Pointer to **[]string** | Stored for reference only. This field does NOT automatically create @mentions when publishing. For LinkedIn @mentions, use the /v1/accounts/{accountId}/linkedin-mentions endpoint to resolve profile URLs to URNs, then embed the returned mentionFormat directly in the post content field. | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Recycling** | Pointer to [**RecyclingState**](RecyclingState.md) |  | [optional] 
**RecycledFromPostId** | Pointer to **string** | ID of the original post if this post was created via recycling | [optional] 
**QueuedFromProfile** | Pointer to **string** | Profile ID if the post was scheduled via the queue | [optional] 
**QueueId** | Pointer to **string** | Queue ID if the post was scheduled via a specific queue | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPost

`func NewPost() *Post`

NewPost instantiates a new Post object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostWithDefaults

`func NewPostWithDefaults() *Post`

NewPostWithDefaults instantiates a new Post object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Post) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Post) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Post) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Post) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Post) GetUserId() PostUserId`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Post) GetUserIdOk() (*PostUserId, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Post) SetUserId(v PostUserId)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Post) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetTitle

`func (o *Post) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Post) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Post) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Post) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetContent

`func (o *Post) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Post) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Post) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Post) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetMediaItems

`func (o *Post) GetMediaItems() []MediaItem`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *Post) GetMediaItemsOk() (*[]MediaItem, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *Post) SetMediaItems(v []MediaItem)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *Post) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.

### GetPlatforms

`func (o *Post) GetPlatforms() []PlatformTarget`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *Post) GetPlatformsOk() (*[]PlatformTarget, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *Post) SetPlatforms(v []PlatformTarget)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *Post) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetScheduledFor

`func (o *Post) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *Post) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *Post) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *Post) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetTimezone

`func (o *Post) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Post) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Post) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Post) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetStatus

`func (o *Post) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Post) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Post) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Post) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *Post) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Post) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Post) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Post) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetHashtags

`func (o *Post) GetHashtags() []string`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *Post) GetHashtagsOk() (*[]string, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *Post) SetHashtags(v []string)`

SetHashtags sets Hashtags field to given value.

### HasHashtags

`func (o *Post) HasHashtags() bool`

HasHashtags returns a boolean if a field has been set.

### GetMentions

`func (o *Post) GetMentions() []string`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *Post) GetMentionsOk() (*[]string, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *Post) SetMentions(v []string)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *Post) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### GetVisibility

`func (o *Post) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *Post) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *Post) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *Post) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMetadata

`func (o *Post) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Post) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Post) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Post) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetRecycling

`func (o *Post) GetRecycling() RecyclingState`

GetRecycling returns the Recycling field if non-nil, zero value otherwise.

### GetRecyclingOk

`func (o *Post) GetRecyclingOk() (*RecyclingState, bool)`

GetRecyclingOk returns a tuple with the Recycling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycling

`func (o *Post) SetRecycling(v RecyclingState)`

SetRecycling sets Recycling field to given value.

### HasRecycling

`func (o *Post) HasRecycling() bool`

HasRecycling returns a boolean if a field has been set.

### GetRecycledFromPostId

`func (o *Post) GetRecycledFromPostId() string`

GetRecycledFromPostId returns the RecycledFromPostId field if non-nil, zero value otherwise.

### GetRecycledFromPostIdOk

`func (o *Post) GetRecycledFromPostIdOk() (*string, bool)`

GetRecycledFromPostIdOk returns a tuple with the RecycledFromPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycledFromPostId

`func (o *Post) SetRecycledFromPostId(v string)`

SetRecycledFromPostId sets RecycledFromPostId field to given value.

### HasRecycledFromPostId

`func (o *Post) HasRecycledFromPostId() bool`

HasRecycledFromPostId returns a boolean if a field has been set.

### GetQueuedFromProfile

`func (o *Post) GetQueuedFromProfile() string`

GetQueuedFromProfile returns the QueuedFromProfile field if non-nil, zero value otherwise.

### GetQueuedFromProfileOk

`func (o *Post) GetQueuedFromProfileOk() (*string, bool)`

GetQueuedFromProfileOk returns a tuple with the QueuedFromProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuedFromProfile

`func (o *Post) SetQueuedFromProfile(v string)`

SetQueuedFromProfile sets QueuedFromProfile field to given value.

### HasQueuedFromProfile

`func (o *Post) HasQueuedFromProfile() bool`

HasQueuedFromProfile returns a boolean if a field has been set.

### GetQueueId

`func (o *Post) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *Post) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *Post) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *Post) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Post) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Post) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Post) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Post) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Post) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Post) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Post) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Post) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


