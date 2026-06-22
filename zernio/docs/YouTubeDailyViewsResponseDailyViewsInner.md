# YouTubeDailyViewsResponseDailyViewsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Views** | Pointer to **int32** |  | [optional] 
**EstimatedMinutesWatched** | Pointer to **float32** |  | [optional] 
**AverageViewDuration** | Pointer to **float32** | Average view duration in seconds | [optional] 
**AverageViewPercentage** | Pointer to **float32** | Average percentage of the video watched per view. Can exceed 100 on Shorts (looping rewatches), so do not clamp it client-side. | [optional] 
**SubscribersGained** | Pointer to **int32** |  | [optional] 
**SubscribersLost** | Pointer to **int32** |  | [optional] 
**Likes** | Pointer to **int32** |  | [optional] 
**Comments** | Pointer to **int32** |  | [optional] 
**Shares** | Pointer to **int32** |  | [optional] 

## Methods

### NewYouTubeDailyViewsResponseDailyViewsInner

`func NewYouTubeDailyViewsResponseDailyViewsInner() *YouTubeDailyViewsResponseDailyViewsInner`

NewYouTubeDailyViewsResponseDailyViewsInner instantiates a new YouTubeDailyViewsResponseDailyViewsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubeDailyViewsResponseDailyViewsInnerWithDefaults

`func NewYouTubeDailyViewsResponseDailyViewsInnerWithDefaults() *YouTubeDailyViewsResponseDailyViewsInner`

NewYouTubeDailyViewsResponseDailyViewsInnerWithDefaults instantiates a new YouTubeDailyViewsResponseDailyViewsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetViews

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetEstimatedMinutesWatched

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetEstimatedMinutesWatched() float32`

GetEstimatedMinutesWatched returns the EstimatedMinutesWatched field if non-nil, zero value otherwise.

### GetEstimatedMinutesWatchedOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetEstimatedMinutesWatchedOk() (*float32, bool)`

GetEstimatedMinutesWatchedOk returns a tuple with the EstimatedMinutesWatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedMinutesWatched

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetEstimatedMinutesWatched(v float32)`

SetEstimatedMinutesWatched sets EstimatedMinutesWatched field to given value.

### HasEstimatedMinutesWatched

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasEstimatedMinutesWatched() bool`

HasEstimatedMinutesWatched returns a boolean if a field has been set.

### GetAverageViewDuration

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetAverageViewDuration() float32`

GetAverageViewDuration returns the AverageViewDuration field if non-nil, zero value otherwise.

### GetAverageViewDurationOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetAverageViewDurationOk() (*float32, bool)`

GetAverageViewDurationOk returns a tuple with the AverageViewDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageViewDuration

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetAverageViewDuration(v float32)`

SetAverageViewDuration sets AverageViewDuration field to given value.

### HasAverageViewDuration

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasAverageViewDuration() bool`

HasAverageViewDuration returns a boolean if a field has been set.

### GetAverageViewPercentage

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetAverageViewPercentage() float32`

GetAverageViewPercentage returns the AverageViewPercentage field if non-nil, zero value otherwise.

### GetAverageViewPercentageOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetAverageViewPercentageOk() (*float32, bool)`

GetAverageViewPercentageOk returns a tuple with the AverageViewPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageViewPercentage

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetAverageViewPercentage(v float32)`

SetAverageViewPercentage sets AverageViewPercentage field to given value.

### HasAverageViewPercentage

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasAverageViewPercentage() bool`

HasAverageViewPercentage returns a boolean if a field has been set.

### GetSubscribersGained

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetSubscribersGained() int32`

GetSubscribersGained returns the SubscribersGained field if non-nil, zero value otherwise.

### GetSubscribersGainedOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetSubscribersGainedOk() (*int32, bool)`

GetSubscribersGainedOk returns a tuple with the SubscribersGained field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersGained

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetSubscribersGained(v int32)`

SetSubscribersGained sets SubscribersGained field to given value.

### HasSubscribersGained

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasSubscribersGained() bool`

HasSubscribersGained returns a boolean if a field has been set.

### GetSubscribersLost

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetSubscribersLost() int32`

GetSubscribersLost returns the SubscribersLost field if non-nil, zero value otherwise.

### GetSubscribersLostOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetSubscribersLostOk() (*int32, bool)`

GetSubscribersLostOk returns a tuple with the SubscribersLost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribersLost

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetSubscribersLost(v int32)`

SetSubscribersLost sets SubscribersLost field to given value.

### HasSubscribersLost

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasSubscribersLost() bool`

HasSubscribersLost returns a boolean if a field has been set.

### GetLikes

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetComments

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetShares

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *YouTubeDailyViewsResponseDailyViewsInner) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *YouTubeDailyViewsResponseDailyViewsInner) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *YouTubeDailyViewsResponseDailyViewsInner) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


