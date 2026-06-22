# GetPostTimeline200ResponseTimelineInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Date in YYYY-MM-DD format | [optional] 
**Platform** | Pointer to **string** | Platform name (e.g. instagram, tiktok) | [optional] 
**PlatformPostId** | Pointer to **string** | Platform-specific post ID | [optional] 
**Impressions** | Pointer to **int32** | Total impressions on this date | [optional] 
**Reach** | Pointer to **int32** | Total reach on this date | [optional] 
**Likes** | Pointer to **int32** | Total likes on this date | [optional] 
**Comments** | Pointer to **int32** | Total comments on this date | [optional] 
**Shares** | Pointer to **int32** | Total shares on this date | [optional] 
**Saves** | Pointer to **int32** | Total saves on this date | [optional] 
**Clicks** | Pointer to **int32** | Total clicks on this date | [optional] 
**Views** | Pointer to **int32** | Total views on this date | [optional] 

## Methods

### NewGetPostTimeline200ResponseTimelineInner

`func NewGetPostTimeline200ResponseTimelineInner() *GetPostTimeline200ResponseTimelineInner`

NewGetPostTimeline200ResponseTimelineInner instantiates a new GetPostTimeline200ResponseTimelineInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPostTimeline200ResponseTimelineInnerWithDefaults

`func NewGetPostTimeline200ResponseTimelineInnerWithDefaults() *GetPostTimeline200ResponseTimelineInner`

NewGetPostTimeline200ResponseTimelineInnerWithDefaults instantiates a new GetPostTimeline200ResponseTimelineInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetPostTimeline200ResponseTimelineInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetPostTimeline200ResponseTimelineInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetPostTimeline200ResponseTimelineInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPlatform

`func (o *GetPostTimeline200ResponseTimelineInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetPostTimeline200ResponseTimelineInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetPostTimeline200ResponseTimelineInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *GetPostTimeline200ResponseTimelineInner) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *GetPostTimeline200ResponseTimelineInner) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *GetPostTimeline200ResponseTimelineInner) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetImpressions

`func (o *GetPostTimeline200ResponseTimelineInner) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *GetPostTimeline200ResponseTimelineInner) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *GetPostTimeline200ResponseTimelineInner) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetReach

`func (o *GetPostTimeline200ResponseTimelineInner) GetReach() int32`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetReachOk() (*int32, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *GetPostTimeline200ResponseTimelineInner) SetReach(v int32)`

SetReach sets Reach field to given value.

### HasReach

`func (o *GetPostTimeline200ResponseTimelineInner) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetLikes

`func (o *GetPostTimeline200ResponseTimelineInner) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *GetPostTimeline200ResponseTimelineInner) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *GetPostTimeline200ResponseTimelineInner) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetComments

`func (o *GetPostTimeline200ResponseTimelineInner) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetPostTimeline200ResponseTimelineInner) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GetPostTimeline200ResponseTimelineInner) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetShares

`func (o *GetPostTimeline200ResponseTimelineInner) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetPostTimeline200ResponseTimelineInner) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *GetPostTimeline200ResponseTimelineInner) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSaves

`func (o *GetPostTimeline200ResponseTimelineInner) GetSaves() int32`

GetSaves returns the Saves field if non-nil, zero value otherwise.

### GetSavesOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetSavesOk() (*int32, bool)`

GetSavesOk returns a tuple with the Saves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaves

`func (o *GetPostTimeline200ResponseTimelineInner) SetSaves(v int32)`

SetSaves sets Saves field to given value.

### HasSaves

`func (o *GetPostTimeline200ResponseTimelineInner) HasSaves() bool`

HasSaves returns a boolean if a field has been set.

### GetClicks

`func (o *GetPostTimeline200ResponseTimelineInner) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetPostTimeline200ResponseTimelineInner) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *GetPostTimeline200ResponseTimelineInner) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetViews

`func (o *GetPostTimeline200ResponseTimelineInner) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetPostTimeline200ResponseTimelineInner) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetPostTimeline200ResponseTimelineInner) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *GetPostTimeline200ResponseTimelineInner) HasViews() bool`

HasViews returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


