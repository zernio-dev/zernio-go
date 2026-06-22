# YouTubeDailyViewsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**VideoId** | Pointer to **string** | The YouTube video ID | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** | Video length in seconds (from YouTube contentDetails.duration) | [optional] 
**DateRange** | Pointer to [**GetGoogleBusinessPerformance200ResponseDateRange**](GetGoogleBusinessPerformance200ResponseDateRange.md) |  | [optional] 
**TotalViews** | Pointer to **int32** | Sum of views across all days in the range | [optional] 
**DailyViews** | Pointer to [**[]YouTubeDailyViewsResponseDailyViewsInner**](YouTubeDailyViewsResponseDailyViewsInner.md) |  | [optional] 
**LastSyncedAt** | Pointer to **NullableTime** | When the data was last synced from YouTube | [optional] 
**ScopeStatus** | Pointer to [**YouTubeDailyViewsResponseScopeStatus**](YouTubeDailyViewsResponseScopeStatus.md) |  | [optional] 

## Methods

### NewYouTubeDailyViewsResponse

`func NewYouTubeDailyViewsResponse() *YouTubeDailyViewsResponse`

NewYouTubeDailyViewsResponse instantiates a new YouTubeDailyViewsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubeDailyViewsResponseWithDefaults

`func NewYouTubeDailyViewsResponseWithDefaults() *YouTubeDailyViewsResponse`

NewYouTubeDailyViewsResponseWithDefaults instantiates a new YouTubeDailyViewsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *YouTubeDailyViewsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *YouTubeDailyViewsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *YouTubeDailyViewsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *YouTubeDailyViewsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetVideoId

`func (o *YouTubeDailyViewsResponse) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *YouTubeDailyViewsResponse) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *YouTubeDailyViewsResponse) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *YouTubeDailyViewsResponse) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *YouTubeDailyViewsResponse) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *YouTubeDailyViewsResponse) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *YouTubeDailyViewsResponse) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *YouTubeDailyViewsResponse) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *YouTubeDailyViewsResponse) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *YouTubeDailyViewsResponse) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetDateRange

`func (o *YouTubeDailyViewsResponse) GetDateRange() GetGoogleBusinessPerformance200ResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *YouTubeDailyViewsResponse) GetDateRangeOk() (*GetGoogleBusinessPerformance200ResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *YouTubeDailyViewsResponse) SetDateRange(v GetGoogleBusinessPerformance200ResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *YouTubeDailyViewsResponse) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetTotalViews

`func (o *YouTubeDailyViewsResponse) GetTotalViews() int32`

GetTotalViews returns the TotalViews field if non-nil, zero value otherwise.

### GetTotalViewsOk

`func (o *YouTubeDailyViewsResponse) GetTotalViewsOk() (*int32, bool)`

GetTotalViewsOk returns a tuple with the TotalViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalViews

`func (o *YouTubeDailyViewsResponse) SetTotalViews(v int32)`

SetTotalViews sets TotalViews field to given value.

### HasTotalViews

`func (o *YouTubeDailyViewsResponse) HasTotalViews() bool`

HasTotalViews returns a boolean if a field has been set.

### GetDailyViews

`func (o *YouTubeDailyViewsResponse) GetDailyViews() []YouTubeDailyViewsResponseDailyViewsInner`

GetDailyViews returns the DailyViews field if non-nil, zero value otherwise.

### GetDailyViewsOk

`func (o *YouTubeDailyViewsResponse) GetDailyViewsOk() (*[]YouTubeDailyViewsResponseDailyViewsInner, bool)`

GetDailyViewsOk returns a tuple with the DailyViews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyViews

`func (o *YouTubeDailyViewsResponse) SetDailyViews(v []YouTubeDailyViewsResponseDailyViewsInner)`

SetDailyViews sets DailyViews field to given value.

### HasDailyViews

`func (o *YouTubeDailyViewsResponse) HasDailyViews() bool`

HasDailyViews returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *YouTubeDailyViewsResponse) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *YouTubeDailyViewsResponse) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *YouTubeDailyViewsResponse) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *YouTubeDailyViewsResponse) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *YouTubeDailyViewsResponse) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *YouTubeDailyViewsResponse) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetScopeStatus

`func (o *YouTubeDailyViewsResponse) GetScopeStatus() YouTubeDailyViewsResponseScopeStatus`

GetScopeStatus returns the ScopeStatus field if non-nil, zero value otherwise.

### GetScopeStatusOk

`func (o *YouTubeDailyViewsResponse) GetScopeStatusOk() (*YouTubeDailyViewsResponseScopeStatus, bool)`

GetScopeStatusOk returns a tuple with the ScopeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStatus

`func (o *YouTubeDailyViewsResponse) SetScopeStatus(v YouTubeDailyViewsResponseScopeStatus)`

SetScopeStatus sets ScopeStatus field to given value.

### HasScopeStatus

`func (o *YouTubeDailyViewsResponse) HasScopeStatus() bool`

HasScopeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


