# GetAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **string** |  | [optional] 
**LatePostId** | Pointer to **string** | Original Zernio post ID if scheduled via Zernio | [optional] 
**Status** | Pointer to **string** | Overall post status. \&quot;partial\&quot; when some platforms published and others failed. | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Analytics** | Pointer to [**PostAnalytics**](PostAnalytics.md) |  | [optional] 
**PlatformAnalytics** | Pointer to [**[]PlatformAnalytics**](PlatformAnalytics.md) |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformPostUrl** | Pointer to **string** |  | [optional] 
**IsExternal** | Pointer to **bool** |  | [optional] 
**SyncStatus** | Pointer to **string** | Overall sync state across all platforms | [optional] 
**Message** | Pointer to **string** | Human-readable status message for pending, partial, or failed states | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**MediaItems** | Pointer to [**[]AnalyticsSinglePostResponseMediaItemsInner**](AnalyticsSinglePostResponseMediaItemsInner.md) | All media items for this post. Carousel posts contain one entry per slide. | [optional] 
**Overview** | Pointer to [**AnalyticsOverview**](AnalyticsOverview.md) |  | [optional] 
**Posts** | Pointer to [**[]AnalyticsListResponsePostsInner**](AnalyticsListResponsePostsInner.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Accounts** | Pointer to [**[]SocialAccount**](SocialAccount.md) | Connected social accounts (followerCount and followersLastUpdated only included if user has analytics add-on) | [optional] 
**HasAnalyticsAccess** | Pointer to **bool** | Whether user has analytics add-on access | [optional] 

## Methods

### NewGetAnalytics200Response

`func NewGetAnalytics200Response() *GetAnalytics200Response`

NewGetAnalytics200Response instantiates a new GetAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalytics200ResponseWithDefaults

`func NewGetAnalytics200ResponseWithDefaults() *GetAnalytics200Response`

NewGetAnalytics200ResponseWithDefaults instantiates a new GetAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *GetAnalytics200Response) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *GetAnalytics200Response) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *GetAnalytics200Response) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *GetAnalytics200Response) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetLatePostId

`func (o *GetAnalytics200Response) GetLatePostId() string`

GetLatePostId returns the LatePostId field if non-nil, zero value otherwise.

### GetLatePostIdOk

`func (o *GetAnalytics200Response) GetLatePostIdOk() (*string, bool)`

GetLatePostIdOk returns a tuple with the LatePostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatePostId

`func (o *GetAnalytics200Response) SetLatePostId(v string)`

SetLatePostId sets LatePostId field to given value.

### HasLatePostId

`func (o *GetAnalytics200Response) HasLatePostId() bool`

HasLatePostId returns a boolean if a field has been set.

### GetStatus

`func (o *GetAnalytics200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAnalytics200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAnalytics200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAnalytics200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetContent

`func (o *GetAnalytics200Response) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *GetAnalytics200Response) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *GetAnalytics200Response) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *GetAnalytics200Response) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetScheduledFor

`func (o *GetAnalytics200Response) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *GetAnalytics200Response) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *GetAnalytics200Response) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *GetAnalytics200Response) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPublishedAt

`func (o *GetAnalytics200Response) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetAnalytics200Response) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetAnalytics200Response) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *GetAnalytics200Response) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetAnalytics

`func (o *GetAnalytics200Response) GetAnalytics() PostAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *GetAnalytics200Response) GetAnalyticsOk() (*PostAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *GetAnalytics200Response) SetAnalytics(v PostAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *GetAnalytics200Response) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetPlatformAnalytics

`func (o *GetAnalytics200Response) GetPlatformAnalytics() []PlatformAnalytics`

GetPlatformAnalytics returns the PlatformAnalytics field if non-nil, zero value otherwise.

### GetPlatformAnalyticsOk

`func (o *GetAnalytics200Response) GetPlatformAnalyticsOk() (*[]PlatformAnalytics, bool)`

GetPlatformAnalyticsOk returns a tuple with the PlatformAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAnalytics

`func (o *GetAnalytics200Response) SetPlatformAnalytics(v []PlatformAnalytics)`

SetPlatformAnalytics sets PlatformAnalytics field to given value.

### HasPlatformAnalytics

`func (o *GetAnalytics200Response) HasPlatformAnalytics() bool`

HasPlatformAnalytics returns a boolean if a field has been set.

### GetPlatform

`func (o *GetAnalytics200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAnalytics200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAnalytics200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetAnalytics200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformPostUrl

`func (o *GetAnalytics200Response) GetPlatformPostUrl() string`

GetPlatformPostUrl returns the PlatformPostUrl field if non-nil, zero value otherwise.

### GetPlatformPostUrlOk

`func (o *GetAnalytics200Response) GetPlatformPostUrlOk() (*string, bool)`

GetPlatformPostUrlOk returns a tuple with the PlatformPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostUrl

`func (o *GetAnalytics200Response) SetPlatformPostUrl(v string)`

SetPlatformPostUrl sets PlatformPostUrl field to given value.

### HasPlatformPostUrl

`func (o *GetAnalytics200Response) HasPlatformPostUrl() bool`

HasPlatformPostUrl returns a boolean if a field has been set.

### GetIsExternal

`func (o *GetAnalytics200Response) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *GetAnalytics200Response) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *GetAnalytics200Response) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *GetAnalytics200Response) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetSyncStatus

`func (o *GetAnalytics200Response) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *GetAnalytics200Response) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *GetAnalytics200Response) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *GetAnalytics200Response) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetMessage

`func (o *GetAnalytics200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetAnalytics200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetAnalytics200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetAnalytics200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *GetAnalytics200Response) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *GetAnalytics200Response) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *GetAnalytics200Response) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *GetAnalytics200Response) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetMediaType

`func (o *GetAnalytics200Response) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *GetAnalytics200Response) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *GetAnalytics200Response) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *GetAnalytics200Response) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetMediaItems

`func (o *GetAnalytics200Response) GetMediaItems() []AnalyticsSinglePostResponseMediaItemsInner`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *GetAnalytics200Response) GetMediaItemsOk() (*[]AnalyticsSinglePostResponseMediaItemsInner, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *GetAnalytics200Response) SetMediaItems(v []AnalyticsSinglePostResponseMediaItemsInner)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *GetAnalytics200Response) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.

### GetOverview

`func (o *GetAnalytics200Response) GetOverview() AnalyticsOverview`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *GetAnalytics200Response) GetOverviewOk() (*AnalyticsOverview, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *GetAnalytics200Response) SetOverview(v AnalyticsOverview)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *GetAnalytics200Response) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetPosts

`func (o *GetAnalytics200Response) GetPosts() []AnalyticsListResponsePostsInner`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *GetAnalytics200Response) GetPostsOk() (*[]AnalyticsListResponsePostsInner, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *GetAnalytics200Response) SetPosts(v []AnalyticsListResponsePostsInner)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *GetAnalytics200Response) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPagination

`func (o *GetAnalytics200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAnalytics200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAnalytics200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetAnalytics200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetAccounts

`func (o *GetAnalytics200Response) GetAccounts() []SocialAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetAnalytics200Response) GetAccountsOk() (*[]SocialAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetAnalytics200Response) SetAccounts(v []SocialAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetAnalytics200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetHasAnalyticsAccess

`func (o *GetAnalytics200Response) GetHasAnalyticsAccess() bool`

GetHasAnalyticsAccess returns the HasAnalyticsAccess field if non-nil, zero value otherwise.

### GetHasAnalyticsAccessOk

`func (o *GetAnalytics200Response) GetHasAnalyticsAccessOk() (*bool, bool)`

GetHasAnalyticsAccessOk returns a tuple with the HasAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticsAccess

`func (o *GetAnalytics200Response) SetHasAnalyticsAccess(v bool)`

SetHasAnalyticsAccess sets HasAnalyticsAccess field to given value.

### HasHasAnalyticsAccess

`func (o *GetAnalytics200Response) HasHasAnalyticsAccess() bool`

HasHasAnalyticsAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


