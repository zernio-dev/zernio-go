# GetLinkedInPostAnalytics200ResponseAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Impressions** | Pointer to **int32** | Times the post was shown | [optional] 
**Reach** | Pointer to **int32** | Unique members who saw the post | [optional] 
**Likes** | Pointer to **int32** | Reactions on the post | [optional] 
**Comments** | Pointer to **int32** | Comments on the post | [optional] 
**Shares** | Pointer to **int32** | Reshares of the post | [optional] 
**Saves** | Pointer to **int32** | Times the post was saved (personal accounts only; 0 for organization accounts) | [optional] 
**Sends** | Pointer to **int32** | Times the post was sent via LinkedIn messaging (personal accounts only; 0 for organization accounts) | [optional] 
**Clicks** | Pointer to **int32** | Clicks on the post (organization accounts only) | [optional] 
**Views** | Pointer to **int32** | Video views (video posts only) | [optional] 
**EngagementRate** | Pointer to **float32** | Engagement rate as percentage | [optional] 

## Methods

### NewGetLinkedInPostAnalytics200ResponseAnalytics

`func NewGetLinkedInPostAnalytics200ResponseAnalytics() *GetLinkedInPostAnalytics200ResponseAnalytics`

NewGetLinkedInPostAnalytics200ResponseAnalytics instantiates a new GetLinkedInPostAnalytics200ResponseAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkedInPostAnalytics200ResponseAnalyticsWithDefaults

`func NewGetLinkedInPostAnalytics200ResponseAnalyticsWithDefaults() *GetLinkedInPostAnalytics200ResponseAnalytics`

NewGetLinkedInPostAnalytics200ResponseAnalyticsWithDefaults instantiates a new GetLinkedInPostAnalytics200ResponseAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImpressions

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetReach

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetReach() int32`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetReachOk() (*int32, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetReach(v int32)`

SetReach sets Reach field to given value.

### HasReach

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetLikes

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetLikes() int32`

GetLikes returns the Likes field if non-nil, zero value otherwise.

### GetLikesOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetLikesOk() (*int32, bool)`

GetLikesOk returns a tuple with the Likes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikes

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetLikes(v int32)`

SetLikes sets Likes field to given value.

### HasLikes

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasLikes() bool`

HasLikes returns a boolean if a field has been set.

### GetComments

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetComments() int32`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetCommentsOk() (*int32, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetComments(v int32)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetShares

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetShares(v int32)`

SetShares sets Shares field to given value.

### HasShares

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasShares() bool`

HasShares returns a boolean if a field has been set.

### GetSaves

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetSaves() int32`

GetSaves returns the Saves field if non-nil, zero value otherwise.

### GetSavesOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetSavesOk() (*int32, bool)`

GetSavesOk returns a tuple with the Saves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaves

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetSaves(v int32)`

SetSaves sets Saves field to given value.

### HasSaves

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasSaves() bool`

HasSaves returns a boolean if a field has been set.

### GetSends

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetSends() int32`

GetSends returns the Sends field if non-nil, zero value otherwise.

### GetSendsOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetSendsOk() (*int32, bool)`

GetSendsOk returns a tuple with the Sends field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSends

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetSends(v int32)`

SetSends sets Sends field to given value.

### HasSends

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasSends() bool`

HasSends returns a boolean if a field has been set.

### GetClicks

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetViews

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetViews(v int32)`

SetViews sets Views field to given value.

### HasViews

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasViews() bool`

HasViews returns a boolean if a field has been set.

### GetEngagementRate

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetEngagementRate() float32`

GetEngagementRate returns the EngagementRate field if non-nil, zero value otherwise.

### GetEngagementRateOk

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) GetEngagementRateOk() (*float32, bool)`

GetEngagementRateOk returns a tuple with the EngagementRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementRate

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) SetEngagementRate(v float32)`

SetEngagementRate sets EngagementRate field to given value.

### HasEngagementRate

`func (o *GetLinkedInPostAnalytics200ResponseAnalytics) HasEngagementRate() bool`

HasEngagementRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


