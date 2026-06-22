# GetInstagramStoryInsights200ResponseDataMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Views** | **int32** | Total story plays. Replaces deprecated &#39;impressions&#39; for media created after 2024-07-02. | 
**Reach** | **int32** | Unique accounts that saw the story. | 
**Replies** | **int32** | DMs sent in reply to the story. | 
**Shares** | **int32** |  | 
**Navigation** | **int32** | Total nav actions (tapsForward + tapsBack + exits + swipesForward). | 
**TapsForward** | **int32** | Tapped right to next slide of SAME story. | 
**TapsBack** | **int32** | Tapped left to previous slide. | 
**Exits** | **int32** | Closed Stories interface entirely. | 
**SwipesForward** | **int32** | Swiped left to next account&#39;s story. | 
**ProfileVisits** | **int32** |  | 
**Follows** | **int32** |  | 
**Reposts** | **int32** |  | 
**TotalInteractions** | **int32** |  | 

## Methods

### NewGetInstagramStoryInsights200ResponseDataMetrics

`func NewGetInstagramStoryInsights200ResponseDataMetrics(views int32, reach int32, replies int32, shares int32, navigation int32, tapsForward int32, tapsBack int32, exits int32, swipesForward int32, profileVisits int32, follows int32, reposts int32, totalInteractions int32, ) *GetInstagramStoryInsights200ResponseDataMetrics`

NewGetInstagramStoryInsights200ResponseDataMetrics instantiates a new GetInstagramStoryInsights200ResponseDataMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstagramStoryInsights200ResponseDataMetricsWithDefaults

`func NewGetInstagramStoryInsights200ResponseDataMetricsWithDefaults() *GetInstagramStoryInsights200ResponseDataMetrics`

NewGetInstagramStoryInsights200ResponseDataMetricsWithDefaults instantiates a new GetInstagramStoryInsights200ResponseDataMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViews

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetViews() int32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetViewsOk() (*int32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetViews(v int32)`

SetViews sets Views field to given value.


### GetReach

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetReach() int32`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetReachOk() (*int32, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetReach(v int32)`

SetReach sets Reach field to given value.


### GetReplies

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetReplies() int32`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetRepliesOk() (*int32, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetReplies(v int32)`

SetReplies sets Replies field to given value.


### GetShares

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetShares() int32`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetSharesOk() (*int32, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetShares(v int32)`

SetShares sets Shares field to given value.


### GetNavigation

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetNavigation() int32`

GetNavigation returns the Navigation field if non-nil, zero value otherwise.

### GetNavigationOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetNavigationOk() (*int32, bool)`

GetNavigationOk returns a tuple with the Navigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigation

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetNavigation(v int32)`

SetNavigation sets Navigation field to given value.


### GetTapsForward

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetTapsForward() int32`

GetTapsForward returns the TapsForward field if non-nil, zero value otherwise.

### GetTapsForwardOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetTapsForwardOk() (*int32, bool)`

GetTapsForwardOk returns a tuple with the TapsForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTapsForward

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetTapsForward(v int32)`

SetTapsForward sets TapsForward field to given value.


### GetTapsBack

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetTapsBack() int32`

GetTapsBack returns the TapsBack field if non-nil, zero value otherwise.

### GetTapsBackOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetTapsBackOk() (*int32, bool)`

GetTapsBackOk returns a tuple with the TapsBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTapsBack

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetTapsBack(v int32)`

SetTapsBack sets TapsBack field to given value.


### GetExits

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetExits() int32`

GetExits returns the Exits field if non-nil, zero value otherwise.

### GetExitsOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetExitsOk() (*int32, bool)`

GetExitsOk returns a tuple with the Exits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExits

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetExits(v int32)`

SetExits sets Exits field to given value.


### GetSwipesForward

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetSwipesForward() int32`

GetSwipesForward returns the SwipesForward field if non-nil, zero value otherwise.

### GetSwipesForwardOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetSwipesForwardOk() (*int32, bool)`

GetSwipesForwardOk returns a tuple with the SwipesForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwipesForward

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetSwipesForward(v int32)`

SetSwipesForward sets SwipesForward field to given value.


### GetProfileVisits

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetProfileVisits() int32`

GetProfileVisits returns the ProfileVisits field if non-nil, zero value otherwise.

### GetProfileVisitsOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetProfileVisitsOk() (*int32, bool)`

GetProfileVisitsOk returns a tuple with the ProfileVisits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileVisits

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetProfileVisits(v int32)`

SetProfileVisits sets ProfileVisits field to given value.


### GetFollows

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetFollows() int32`

GetFollows returns the Follows field if non-nil, zero value otherwise.

### GetFollowsOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetFollowsOk() (*int32, bool)`

GetFollowsOk returns a tuple with the Follows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollows

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetFollows(v int32)`

SetFollows sets Follows field to given value.


### GetReposts

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetReposts() int32`

GetReposts returns the Reposts field if non-nil, zero value otherwise.

### GetRepostsOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetRepostsOk() (*int32, bool)`

GetRepostsOk returns a tuple with the Reposts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReposts

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetReposts(v int32)`

SetReposts sets Reposts field to given value.


### GetTotalInteractions

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetTotalInteractions() int32`

GetTotalInteractions returns the TotalInteractions field if non-nil, zero value otherwise.

### GetTotalInteractionsOk

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) GetTotalInteractionsOk() (*int32, bool)`

GetTotalInteractionsOk returns a tuple with the TotalInteractions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalInteractions

`func (o *GetInstagramStoryInsights200ResponseDataMetrics) SetTotalInteractions(v int32)`

SetTotalInteractions sets TotalInteractions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


