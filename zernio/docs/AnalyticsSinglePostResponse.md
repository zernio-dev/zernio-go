# AnalyticsSinglePostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **string** |  | [optional] 
**LatePostId** | Pointer to **NullableString** | Original Zernio post ID if scheduled via Zernio | [optional] 
**Status** | Pointer to **string** | Overall post status. \&quot;partial\&quot; when some platforms published and others failed. | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | Pointer to **NullableTime** |  | [optional] 
**Analytics** | Pointer to [**PostAnalytics**](PostAnalytics.md) |  | [optional] 
**PlatformAnalytics** | Pointer to [**[]PlatformAnalytics**](PlatformAnalytics.md) |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformPostUrl** | Pointer to **NullableString** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**SyncStatus** | Pointer to **string** | Overall sync state across all platforms | [optional] 
**Message** | Pointer to **NullableString** | Human-readable status message for pending, partial, or failed states | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** |  | [optional] 
**MediaType** | Pointer to **NullableString** |  | [optional] 
**MediaItems** | Pointer to [**[]AnalyticsSinglePostResponseMediaItemsInner**](AnalyticsSinglePostResponseMediaItemsInner.md) | All media items for this post. Carousel posts contain one entry per slide. | [optional] 

## Methods

### NewAnalyticsSinglePostResponse

`func NewAnalyticsSinglePostResponse() *AnalyticsSinglePostResponse`

NewAnalyticsSinglePostResponse instantiates a new AnalyticsSinglePostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsSinglePostResponseWithDefaults

`func NewAnalyticsSinglePostResponseWithDefaults() *AnalyticsSinglePostResponse`

NewAnalyticsSinglePostResponseWithDefaults instantiates a new AnalyticsSinglePostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *AnalyticsSinglePostResponse) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *AnalyticsSinglePostResponse) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *AnalyticsSinglePostResponse) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *AnalyticsSinglePostResponse) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetLatePostId

`func (o *AnalyticsSinglePostResponse) GetLatePostId() string`

GetLatePostId returns the LatePostId field if non-nil, zero value otherwise.

### GetLatePostIdOk

`func (o *AnalyticsSinglePostResponse) GetLatePostIdOk() (*string, bool)`

GetLatePostIdOk returns a tuple with the LatePostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatePostId

`func (o *AnalyticsSinglePostResponse) SetLatePostId(v string)`

SetLatePostId sets LatePostId field to given value.

### HasLatePostId

`func (o *AnalyticsSinglePostResponse) HasLatePostId() bool`

HasLatePostId returns a boolean if a field has been set.

### SetLatePostIdNil

`func (o *AnalyticsSinglePostResponse) SetLatePostIdNil(b bool)`

 SetLatePostIdNil sets the value for LatePostId to be an explicit nil

### UnsetLatePostId
`func (o *AnalyticsSinglePostResponse) UnsetLatePostId()`

UnsetLatePostId ensures that no value is present for LatePostId, not even an explicit nil
### GetStatus

`func (o *AnalyticsSinglePostResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalyticsSinglePostResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalyticsSinglePostResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AnalyticsSinglePostResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContent

`func (o *AnalyticsSinglePostResponse) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AnalyticsSinglePostResponse) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AnalyticsSinglePostResponse) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AnalyticsSinglePostResponse) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetScheduledFor

`func (o *AnalyticsSinglePostResponse) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *AnalyticsSinglePostResponse) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *AnalyticsSinglePostResponse) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *AnalyticsSinglePostResponse) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPublishedAt

`func (o *AnalyticsSinglePostResponse) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *AnalyticsSinglePostResponse) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *AnalyticsSinglePostResponse) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *AnalyticsSinglePostResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *AnalyticsSinglePostResponse) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *AnalyticsSinglePostResponse) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetAnalytics

`func (o *AnalyticsSinglePostResponse) GetAnalytics() PostAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *AnalyticsSinglePostResponse) GetAnalyticsOk() (*PostAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *AnalyticsSinglePostResponse) SetAnalytics(v PostAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *AnalyticsSinglePostResponse) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetPlatformAnalytics

`func (o *AnalyticsSinglePostResponse) GetPlatformAnalytics() []PlatformAnalytics`

GetPlatformAnalytics returns the PlatformAnalytics field if non-nil, zero value otherwise.

### GetPlatformAnalyticsOk

`func (o *AnalyticsSinglePostResponse) GetPlatformAnalyticsOk() (*[]PlatformAnalytics, bool)`

GetPlatformAnalyticsOk returns a tuple with the PlatformAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAnalytics

`func (o *AnalyticsSinglePostResponse) SetPlatformAnalytics(v []PlatformAnalytics)`

SetPlatformAnalytics sets PlatformAnalytics field to given value.

### HasPlatformAnalytics

`func (o *AnalyticsSinglePostResponse) HasPlatformAnalytics() bool`

HasPlatformAnalytics returns a boolean if a field has been set.

### GetPlatform

`func (o *AnalyticsSinglePostResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AnalyticsSinglePostResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AnalyticsSinglePostResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AnalyticsSinglePostResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformPostUrl

`func (o *AnalyticsSinglePostResponse) GetPlatformPostUrl() string`

GetPlatformPostUrl returns the PlatformPostUrl field if non-nil, zero value otherwise.

### GetPlatformPostUrlOk

`func (o *AnalyticsSinglePostResponse) GetPlatformPostUrlOk() (*string, bool)`

GetPlatformPostUrlOk returns a tuple with the PlatformPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostUrl

`func (o *AnalyticsSinglePostResponse) SetPlatformPostUrl(v string)`

SetPlatformPostUrl sets PlatformPostUrl field to given value.

### HasPlatformPostUrl

`func (o *AnalyticsSinglePostResponse) HasPlatformPostUrl() bool`

HasPlatformPostUrl returns a boolean if a field has been set.

### SetPlatformPostUrlNil

`func (o *AnalyticsSinglePostResponse) SetPlatformPostUrlNil(b bool)`

 SetPlatformPostUrlNil sets the value for PlatformPostUrl to be an explicit nil

### UnsetPlatformPostUrl
`func (o *AnalyticsSinglePostResponse) UnsetPlatformPostUrl()`

UnsetPlatformPostUrl ensures that no value is present for PlatformPostUrl, not even an explicit nil
### GetIsExternal

`func (o *AnalyticsSinglePostResponse) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *AnalyticsSinglePostResponse) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *AnalyticsSinglePostResponse) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *AnalyticsSinglePostResponse) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetSyncStatus

`func (o *AnalyticsSinglePostResponse) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *AnalyticsSinglePostResponse) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *AnalyticsSinglePostResponse) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *AnalyticsSinglePostResponse) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetMessage

`func (o *AnalyticsSinglePostResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnalyticsSinglePostResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnalyticsSinglePostResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AnalyticsSinglePostResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AnalyticsSinglePostResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AnalyticsSinglePostResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetThumbnailUrl

`func (o *AnalyticsSinglePostResponse) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *AnalyticsSinglePostResponse) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *AnalyticsSinglePostResponse) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *AnalyticsSinglePostResponse) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *AnalyticsSinglePostResponse) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *AnalyticsSinglePostResponse) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetMediaType

`func (o *AnalyticsSinglePostResponse) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *AnalyticsSinglePostResponse) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *AnalyticsSinglePostResponse) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *AnalyticsSinglePostResponse) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *AnalyticsSinglePostResponse) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *AnalyticsSinglePostResponse) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetMediaItems

`func (o *AnalyticsSinglePostResponse) GetMediaItems() []AnalyticsSinglePostResponseMediaItemsInner`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *AnalyticsSinglePostResponse) GetMediaItemsOk() (*[]AnalyticsSinglePostResponseMediaItemsInner, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *AnalyticsSinglePostResponse) SetMediaItems(v []AnalyticsSinglePostResponseMediaItemsInner)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *AnalyticsSinglePostResponse) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


