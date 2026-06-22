# PostAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impressions** | Pointer to **int32** |  | [optional] 
**Reach** | Pointer to **int32** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**Comments** | Pointer to **int32** |  | [optional] 
**Shares** | Pointer to **int32** |  | [optional] 
**Saves** | Pointer to **int32** | Number of saves/bookmarks (Instagram, Pinterest) | [optional] 
**Clicks** | Pointer to **int32** |  | [optional] 
**Views** | Pointer to **int32** |  | [optional] 
**IgReelsAvgWatchTime** | Pointer to **int32** | Instagram Reels only: average watch time per play, in milliseconds. 0 for non-Reels media and other platforms. | [optional] 
**IgReelsVideoViewTotalTime** | Pointer to **int32** | Instagram Reels only: total watch time including replays, in milliseconds. 0 for non-Reels media and other platforms. | [optional] 
**EngagementRate** | Pointer to **float32** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPostAnalytics

`func NewPostAnalytics() *PostAnalytics`

NewPostAnalytics instantiates a new PostAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostAnalyticsWithDefaults

`func NewPostAnalyticsWithDefaults() *PostAnalytics`

NewPostAnalyticsWithDefaults instantiates a new PostAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpressions

`func (o *PostAnalytics) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *PostAnalytics) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *PostAnalytics) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *PostAnalytics) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetReach

`func (o *PostAnalytics) GetReach() int32`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *PostAnalytics) GetReachOk() (*int32, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *PostAnalytics) SetReach(v int32)`

SetReach sets Reach field to given value.

### HasReach

`func (o *PostAnalytics) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetLikes

`func (o *PostAnalytics) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *PostAnalytics) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *PostAnalytics) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *PostAnalytics) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetComments

`func (o *PostAnalytics) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PostAnalytics) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PostAnalytics) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PostAnalytics) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetShares

`func (o *PostAnalytics) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *PostAnalytics) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *PostAnalytics) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *PostAnalytics) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSaves

`func (o *PostAnalytics) GetSaves() int32`

GetSaves returns the Saves field if non-nil, zero value otherwise.

### GetSavesOk

`func (o *PostAnalytics) GetSavesOk() (*int32, bool)`

GetSavesOk returns a tuple with the Saves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaves

`func (o *PostAnalytics) SetSaves(v int32)`

SetSaves sets Saves field to given value.

### HasSaves

`func (o *PostAnalytics) HasSaves() bool`

HasSaves returns a boolean if a field has been set.

### GetClicks

`func (o *PostAnalytics) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *PostAnalytics) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *PostAnalytics) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *PostAnalytics) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetViews

`func (o *PostAnalytics) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *PostAnalytics) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *PostAnalytics) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *PostAnalytics) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetIgReelsAvgWatchTime

`func (o *PostAnalytics) GetIgReelsAvgWatchTime() int32`

GetIgReelsAvgWatchTime returns the IgReelsAvgWatchTime field if non-nil, zero value otherwise.

### GetIgReelsAvgWatchTimeOk

`func (o *PostAnalytics) GetIgReelsAvgWatchTimeOk() (*int32, bool)`

GetIgReelsAvgWatchTimeOk returns a tuple with the IgReelsAvgWatchTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgReelsAvgWatchTime

`func (o *PostAnalytics) SetIgReelsAvgWatchTime(v int32)`

SetIgReelsAvgWatchTime sets IgReelsAvgWatchTime field to given value.

### HasIgReelsAvgWatchTime

`func (o *PostAnalytics) HasIgReelsAvgWatchTime() bool`

HasIgReelsAvgWatchTime returns a boolean if a field has been set.

### GetIgReelsVideoViewTotalTime

`func (o *PostAnalytics) GetIgReelsVideoViewTotalTime() int32`

GetIgReelsVideoViewTotalTime returns the IgReelsVideoViewTotalTime field if non-nil, zero value otherwise.

### GetIgReelsVideoViewTotalTimeOk

`func (o *PostAnalytics) GetIgReelsVideoViewTotalTimeOk() (*int32, bool)`

GetIgReelsVideoViewTotalTimeOk returns a tuple with the IgReelsVideoViewTotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgReelsVideoViewTotalTime

`func (o *PostAnalytics) SetIgReelsVideoViewTotalTime(v int32)`

SetIgReelsVideoViewTotalTime sets IgReelsVideoViewTotalTime field to given value.

### HasIgReelsVideoViewTotalTime

`func (o *PostAnalytics) HasIgReelsVideoViewTotalTime() bool`

HasIgReelsVideoViewTotalTime returns a boolean if a field has been set.

### GetEngagementRate

`func (o *PostAnalytics) GetEngagementRate() float32`

GetEngagementRate returns the EngagementRate field if non-nil, zero value otherwise.

### GetEngagementRateOk

`func (o *PostAnalytics) GetEngagementRateOk() (*float32, bool)`

GetEngagementRateOk returns a tuple with the EngagementRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementRate

`func (o *PostAnalytics) SetEngagementRate(v float32)`

SetEngagementRate sets EngagementRate field to given value.

### HasEngagementRate

`func (o *PostAnalytics) HasEngagementRate() bool`

HasEngagementRate returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PostAnalytics) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PostAnalytics) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PostAnalytics) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PostAnalytics) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


