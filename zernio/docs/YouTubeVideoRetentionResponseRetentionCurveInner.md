# YouTubeVideoRetentionResponseRetentionCurveInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElapsedVideoTimeRatio** | Pointer to **float32** | Position in the video as a ratio (0.01-1.0, exclusive end of each interval) | [optional] 
**AudienceWatchRatio** | Pointer to **float32** | Absolute share of viewers watching at this point. Can exceed 1 (rewinds/looping, common on Shorts). | [optional] 
**RelativeRetentionPerformance** | Pointer to **float32** | Retention vs videos of similar length (0 &#x3D; worst, 0.5 &#x3D; median, 1 &#x3D; best) | [optional] 
**StartedWatching** | Pointer to **int32** | Viewers who started watching in this segment | [optional] 
**StoppedWatching** | Pointer to **int32** | Viewers who stopped watching in this segment | [optional] 
**TotalSegmentImpressions** | Pointer to **int32** | Total views of this segment, including rewatches | [optional] 

## Methods

### NewYouTubeVideoRetentionResponseRetentionCurveInner

`func NewYouTubeVideoRetentionResponseRetentionCurveInner() *YouTubeVideoRetentionResponseRetentionCurveInner`

NewYouTubeVideoRetentionResponseRetentionCurveInner instantiates a new YouTubeVideoRetentionResponseRetentionCurveInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubeVideoRetentionResponseRetentionCurveInnerWithDefaults

`func NewYouTubeVideoRetentionResponseRetentionCurveInnerWithDefaults() *YouTubeVideoRetentionResponseRetentionCurveInner`

NewYouTubeVideoRetentionResponseRetentionCurveInnerWithDefaults instantiates a new YouTubeVideoRetentionResponseRetentionCurveInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsedVideoTimeRatio

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetElapsedVideoTimeRatio() float32`

GetElapsedVideoTimeRatio returns the ElapsedVideoTimeRatio field if non-nil, zero value otherwise.

### GetElapsedVideoTimeRatioOk

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetElapsedVideoTimeRatioOk() (*float32, bool)`

GetElapsedVideoTimeRatioOk returns a tuple with the ElapsedVideoTimeRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedVideoTimeRatio

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) SetElapsedVideoTimeRatio(v float32)`

SetElapsedVideoTimeRatio sets ElapsedVideoTimeRatio field to given value.

### HasElapsedVideoTimeRatio

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) HasElapsedVideoTimeRatio() bool`

HasElapsedVideoTimeRatio returns a boolean if a field has been set.

### GetAudienceWatchRatio

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetAudienceWatchRatio() float32`

GetAudienceWatchRatio returns the AudienceWatchRatio field if non-nil, zero value otherwise.

### GetAudienceWatchRatioOk

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetAudienceWatchRatioOk() (*float32, bool)`

GetAudienceWatchRatioOk returns a tuple with the AudienceWatchRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceWatchRatio

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) SetAudienceWatchRatio(v float32)`

SetAudienceWatchRatio sets AudienceWatchRatio field to given value.

### HasAudienceWatchRatio

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) HasAudienceWatchRatio() bool`

HasAudienceWatchRatio returns a boolean if a field has been set.

### GetRelativeRetentionPerformance

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetRelativeRetentionPerformance() float32`

GetRelativeRetentionPerformance returns the RelativeRetentionPerformance field if non-nil, zero value otherwise.

### GetRelativeRetentionPerformanceOk

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetRelativeRetentionPerformanceOk() (*float32, bool)`

GetRelativeRetentionPerformanceOk returns a tuple with the RelativeRetentionPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeRetentionPerformance

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) SetRelativeRetentionPerformance(v float32)`

SetRelativeRetentionPerformance sets RelativeRetentionPerformance field to given value.

### HasRelativeRetentionPerformance

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) HasRelativeRetentionPerformance() bool`

HasRelativeRetentionPerformance returns a boolean if a field has been set.

### GetStartedWatching

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetStartedWatching() int32`

GetStartedWatching returns the StartedWatching field if non-nil, zero value otherwise.

### GetStartedWatchingOk

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetStartedWatchingOk() (*int32, bool)`

GetStartedWatchingOk returns a tuple with the StartedWatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedWatching

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) SetStartedWatching(v int32)`

SetStartedWatching sets StartedWatching field to given value.

### HasStartedWatching

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) HasStartedWatching() bool`

HasStartedWatching returns a boolean if a field has been set.

### GetStoppedWatching

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetStoppedWatching() int32`

GetStoppedWatching returns the StoppedWatching field if non-nil, zero value otherwise.

### GetStoppedWatchingOk

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetStoppedWatchingOk() (*int32, bool)`

GetStoppedWatchingOk returns a tuple with the StoppedWatching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoppedWatching

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) SetStoppedWatching(v int32)`

SetStoppedWatching sets StoppedWatching field to given value.

### HasStoppedWatching

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) HasStoppedWatching() bool`

HasStoppedWatching returns a boolean if a field has been set.

### GetTotalSegmentImpressions

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetTotalSegmentImpressions() int32`

GetTotalSegmentImpressions returns the TotalSegmentImpressions field if non-nil, zero value otherwise.

### GetTotalSegmentImpressionsOk

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) GetTotalSegmentImpressionsOk() (*int32, bool)`

GetTotalSegmentImpressionsOk returns a tuple with the TotalSegmentImpressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSegmentImpressions

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) SetTotalSegmentImpressions(v int32)`

SetTotalSegmentImpressions sets TotalSegmentImpressions field to given value.

### HasTotalSegmentImpressions

`func (o *YouTubeVideoRetentionResponseRetentionCurveInner) HasTotalSegmentImpressions() bool`

HasTotalSegmentImpressions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


