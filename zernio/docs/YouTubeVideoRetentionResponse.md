# YouTubeVideoRetentionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** | The Zernio account ID for the YouTube account | [optional] 
**VideoId** | Pointer to **string** | The YouTube video ID | [optional] 
**Title** | Pointer to **NullableString** | Video title | [optional] 
**PublishedAt** | Pointer to **NullableTime** | When the video was published on YouTube | [optional] 
**DurationSeconds** | Pointer to **NullableInt32** | Video length in seconds (from YouTube contentDetails.duration) | [optional] 
**DateRange** | Pointer to [**GetGoogleBusinessPerformance200ResponseDateRange**](GetGoogleBusinessPerformance200ResponseDateRange.md) |  | [optional] 
**RetentionCurve** | Pointer to [**[]YouTubeVideoRetentionResponseRetentionCurveInner**](YouTubeVideoRetentionResponseRetentionCurveInner.md) | Up to 100 points covering the video timeline, aggregated over the date range. Empty for videos with very few views. | [optional] 
**Note** | Pointer to **string** | Present only when the curve is empty, explaining why | [optional] 
**ScopeStatus** | Pointer to [**YouTubeDailyViewsResponseScopeStatus**](YouTubeDailyViewsResponseScopeStatus.md) |  | [optional] 

## Methods

### NewYouTubeVideoRetentionResponse

`func NewYouTubeVideoRetentionResponse() *YouTubeVideoRetentionResponse`

NewYouTubeVideoRetentionResponse instantiates a new YouTubeVideoRetentionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubeVideoRetentionResponseWithDefaults

`func NewYouTubeVideoRetentionResponseWithDefaults() *YouTubeVideoRetentionResponse`

NewYouTubeVideoRetentionResponseWithDefaults instantiates a new YouTubeVideoRetentionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *YouTubeVideoRetentionResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *YouTubeVideoRetentionResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *YouTubeVideoRetentionResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *YouTubeVideoRetentionResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *YouTubeVideoRetentionResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *YouTubeVideoRetentionResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *YouTubeVideoRetentionResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *YouTubeVideoRetentionResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetVideoId

`func (o *YouTubeVideoRetentionResponse) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *YouTubeVideoRetentionResponse) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *YouTubeVideoRetentionResponse) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *YouTubeVideoRetentionResponse) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetTitle

`func (o *YouTubeVideoRetentionResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *YouTubeVideoRetentionResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *YouTubeVideoRetentionResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *YouTubeVideoRetentionResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *YouTubeVideoRetentionResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *YouTubeVideoRetentionResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPublishedAt

`func (o *YouTubeVideoRetentionResponse) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *YouTubeVideoRetentionResponse) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *YouTubeVideoRetentionResponse) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *YouTubeVideoRetentionResponse) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *YouTubeVideoRetentionResponse) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *YouTubeVideoRetentionResponse) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetDurationSeconds

`func (o *YouTubeVideoRetentionResponse) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *YouTubeVideoRetentionResponse) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *YouTubeVideoRetentionResponse) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *YouTubeVideoRetentionResponse) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### SetDurationSecondsNil

`func (o *YouTubeVideoRetentionResponse) SetDurationSecondsNil(b bool)`

 SetDurationSecondsNil sets the value for DurationSeconds to be an explicit nil

### UnsetDurationSeconds
`func (o *YouTubeVideoRetentionResponse) UnsetDurationSeconds()`

UnsetDurationSeconds ensures that no value is present for DurationSeconds, not even an explicit nil
### GetDateRange

`func (o *YouTubeVideoRetentionResponse) GetDateRange() GetGoogleBusinessPerformance200ResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *YouTubeVideoRetentionResponse) GetDateRangeOk() (*GetGoogleBusinessPerformance200ResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *YouTubeVideoRetentionResponse) SetDateRange(v GetGoogleBusinessPerformance200ResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *YouTubeVideoRetentionResponse) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetRetentionCurve

`func (o *YouTubeVideoRetentionResponse) GetRetentionCurve() []YouTubeVideoRetentionResponseRetentionCurveInner`

GetRetentionCurve returns the RetentionCurve field if non-nil, zero value otherwise.

### GetRetentionCurveOk

`func (o *YouTubeVideoRetentionResponse) GetRetentionCurveOk() (*[]YouTubeVideoRetentionResponseRetentionCurveInner, bool)`

GetRetentionCurveOk returns a tuple with the RetentionCurve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionCurve

`func (o *YouTubeVideoRetentionResponse) SetRetentionCurve(v []YouTubeVideoRetentionResponseRetentionCurveInner)`

SetRetentionCurve sets RetentionCurve field to given value.

### HasRetentionCurve

`func (o *YouTubeVideoRetentionResponse) HasRetentionCurve() bool`

HasRetentionCurve returns a boolean if a field has been set.

### GetNote

`func (o *YouTubeVideoRetentionResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *YouTubeVideoRetentionResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *YouTubeVideoRetentionResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *YouTubeVideoRetentionResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetScopeStatus

`func (o *YouTubeVideoRetentionResponse) GetScopeStatus() YouTubeDailyViewsResponseScopeStatus`

GetScopeStatus returns the ScopeStatus field if non-nil, zero value otherwise.

### GetScopeStatusOk

`func (o *YouTubeVideoRetentionResponse) GetScopeStatusOk() (*YouTubeDailyViewsResponseScopeStatus, bool)`

GetScopeStatusOk returns a tuple with the ScopeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStatus

`func (o *YouTubeVideoRetentionResponse) SetScopeStatus(v YouTubeDailyViewsResponseScopeStatus)`

SetScopeStatus sets ScopeStatus field to given value.

### HasScopeStatus

`func (o *YouTubeVideoRetentionResponse) HasScopeStatus() bool`

HasScopeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


