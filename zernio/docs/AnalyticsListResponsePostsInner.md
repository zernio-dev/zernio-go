# AnalyticsListResponsePostsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**LatePostId** | Pointer to **NullableString** | Original Zernio post ID if scheduled via Zernio | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Analytics** | Pointer to [**PostAnalytics**](PostAnalytics.md) |  | [optional] 
**Platforms** | Pointer to [**[]PlatformAnalytics**](PlatformAnalytics.md) |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformPostUrl** | Pointer to **string** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**ProfileId** | Pointer to **NullableString** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**MediaItems** | Pointer to [**[]AnalyticsSinglePostResponseMediaItemsInner**](AnalyticsSinglePostResponseMediaItemsInner.md) | All media items for this post. Carousel posts contain one entry per slide. | [optional] 

## Methods

### NewAnalyticsListResponsePostsInner

`func NewAnalyticsListResponsePostsInner() *AnalyticsListResponsePostsInner`

NewAnalyticsListResponsePostsInner instantiates a new AnalyticsListResponsePostsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsListResponsePostsInnerWithDefaults

`func NewAnalyticsListResponsePostsInnerWithDefaults() *AnalyticsListResponsePostsInner`

NewAnalyticsListResponsePostsInnerWithDefaults instantiates a new AnalyticsListResponsePostsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AnalyticsListResponsePostsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AnalyticsListResponsePostsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AnalyticsListResponsePostsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AnalyticsListResponsePostsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatePostId

`func (o *AnalyticsListResponsePostsInner) GetLatePostId() string`

GetLatePostId returns the LatePostId field if non-nil, zero value otherwise.

### GetLatePostIdOk

`func (o *AnalyticsListResponsePostsInner) GetLatePostIdOk() (*string, bool)`

GetLatePostIdOk returns a tuple with the LatePostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatePostId

`func (o *AnalyticsListResponsePostsInner) SetLatePostId(v string)`

SetLatePostId sets LatePostId field to given value.

### HasLatePostId

`func (o *AnalyticsListResponsePostsInner) HasLatePostId() bool`

HasLatePostId returns a boolean if a field has been set.

### SetLatePostIdNil

`func (o *AnalyticsListResponsePostsInner) SetLatePostIdNil(b bool)`

 SetLatePostIdNil sets the value for LatePostId to be an explicit nil

### UnsetLatePostId
`func (o *AnalyticsListResponsePostsInner) UnsetLatePostId()`

UnsetLatePostId ensures that no value is present for LatePostId, not even an explicit nil
### GetContent

`func (o *AnalyticsListResponsePostsInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AnalyticsListResponsePostsInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AnalyticsListResponsePostsInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AnalyticsListResponsePostsInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetScheduledFor

`func (o *AnalyticsListResponsePostsInner) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *AnalyticsListResponsePostsInner) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *AnalyticsListResponsePostsInner) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *AnalyticsListResponsePostsInner) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPublishedAt

`func (o *AnalyticsListResponsePostsInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *AnalyticsListResponsePostsInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *AnalyticsListResponsePostsInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *AnalyticsListResponsePostsInner) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetStatus

`func (o *AnalyticsListResponsePostsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalyticsListResponsePostsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalyticsListResponsePostsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnalyticsListResponsePostsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAnalytics

`func (o *AnalyticsListResponsePostsInner) GetAnalytics() PostAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *AnalyticsListResponsePostsInner) GetAnalyticsOk() (*PostAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *AnalyticsListResponsePostsInner) SetAnalytics(v PostAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *AnalyticsListResponsePostsInner) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetPlatforms

`func (o *AnalyticsListResponsePostsInner) GetPlatforms() []PlatformAnalytics`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *AnalyticsListResponsePostsInner) GetPlatformsOk() (*[]PlatformAnalytics, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *AnalyticsListResponsePostsInner) SetPlatforms(v []PlatformAnalytics)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *AnalyticsListResponsePostsInner) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetPlatform

`func (o *AnalyticsListResponsePostsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AnalyticsListResponsePostsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AnalyticsListResponsePostsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AnalyticsListResponsePostsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformPostUrl

`func (o *AnalyticsListResponsePostsInner) GetPlatformPostUrl() string`

GetPlatformPostUrl returns the PlatformPostUrl field if non-nil, zero value otherwise.

### GetPlatformPostUrlOk

`func (o *AnalyticsListResponsePostsInner) GetPlatformPostUrlOk() (*string, bool)`

GetPlatformPostUrlOk returns a tuple with the PlatformPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostUrl

`func (o *AnalyticsListResponsePostsInner) SetPlatformPostUrl(v string)`

SetPlatformPostUrl sets PlatformPostUrl field to given value.

### HasPlatformPostUrl

`func (o *AnalyticsListResponsePostsInner) HasPlatformPostUrl() bool`

HasPlatformPostUrl returns a boolean if a field has been set.

### GetIsExternal

`func (o *AnalyticsListResponsePostsInner) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *AnalyticsListResponsePostsInner) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *AnalyticsListResponsePostsInner) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *AnalyticsListResponsePostsInner) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetProfileId

`func (o *AnalyticsListResponsePostsInner) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AnalyticsListResponsePostsInner) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AnalyticsListResponsePostsInner) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AnalyticsListResponsePostsInner) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### SetProfileIdNil

`func (o *AnalyticsListResponsePostsInner) SetProfileIdNil(b bool)`

 SetProfileIdNil sets the value for ProfileId to be an explicit nil

### UnsetProfileId
`func (o *AnalyticsListResponsePostsInner) UnsetProfileId()`

UnsetProfileId ensures that no value is present for ProfileId, not even an explicit nil
### GetThumbnailUrl

`func (o *AnalyticsListResponsePostsInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *AnalyticsListResponsePostsInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *AnalyticsListResponsePostsInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *AnalyticsListResponsePostsInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetMediaType

`func (o *AnalyticsListResponsePostsInner) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AnalyticsListResponsePostsInner) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AnalyticsListResponsePostsInner) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AnalyticsListResponsePostsInner) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaItems

`func (o *AnalyticsListResponsePostsInner) GetMediaItems() []AnalyticsSinglePostResponseMediaItemsInner`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *AnalyticsListResponsePostsInner) GetMediaItemsOk() (*[]AnalyticsSinglePostResponseMediaItemsInner, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *AnalyticsListResponsePostsInner) SetMediaItems(v []AnalyticsSinglePostResponseMediaItemsInner)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *AnalyticsListResponsePostsInner) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


