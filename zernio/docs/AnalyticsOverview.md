# AnalyticsOverview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalPosts** | Pointer to **int32** |  | [optional] 
**PublishedPosts** | Pointer to **int32** |  | [optional] 
**ScheduledPosts** | Pointer to **int32** |  | [optional] 
**LastSync** | Pointer to **NullableTime** |  | [optional] 
**DataStaleness** | Pointer to [**AnalyticsOverviewDataStaleness**](AnalyticsOverviewDataStaleness.md) |  | [optional] 

## Methods

### NewAnalyticsOverview

`func NewAnalyticsOverview() *AnalyticsOverview`

NewAnalyticsOverview instantiates a new AnalyticsOverview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsOverviewWithDefaults

`func NewAnalyticsOverviewWithDefaults() *AnalyticsOverview`

NewAnalyticsOverviewWithDefaults instantiates a new AnalyticsOverview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalPosts

`func (o *AnalyticsOverview) GetTotalPosts() int32`

GetTotalPosts returns the TotalPosts field if non-nil, zero value otherwise.

### GetTotalPostsOk

`func (o *AnalyticsOverview) GetTotalPostsOk() (*int32, bool)`

GetTotalPostsOk returns a tuple with the TotalPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalPosts

`func (o *AnalyticsOverview) SetTotalPosts(v int32)`

SetTotalPosts sets TotalPosts field to given value.

### HasTotalPosts

`func (o *AnalyticsOverview) HasTotalPosts() bool`

HasTotalPosts returns a boolean if a field has been set.

### GetPublishedPosts

`func (o *AnalyticsOverview) GetPublishedPosts() int32`

GetPublishedPosts returns the PublishedPosts field if non-nil, zero value otherwise.

### GetPublishedPostsOk

`func (o *AnalyticsOverview) GetPublishedPostsOk() (*int32, bool)`

GetPublishedPostsOk returns a tuple with the PublishedPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedPosts

`func (o *AnalyticsOverview) SetPublishedPosts(v int32)`

SetPublishedPosts sets PublishedPosts field to given value.

### HasPublishedPosts

`func (o *AnalyticsOverview) HasPublishedPosts() bool`

HasPublishedPosts returns a boolean if a field has been set.

### GetScheduledPosts

`func (o *AnalyticsOverview) GetScheduledPosts() int32`

GetScheduledPosts returns the ScheduledPosts field if non-nil, zero value otherwise.

### GetScheduledPostsOk

`func (o *AnalyticsOverview) GetScheduledPostsOk() (*int32, bool)`

GetScheduledPostsOk returns a tuple with the ScheduledPosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPosts

`func (o *AnalyticsOverview) SetScheduledPosts(v int32)`

SetScheduledPosts sets ScheduledPosts field to given value.

### HasScheduledPosts

`func (o *AnalyticsOverview) HasScheduledPosts() bool`

HasScheduledPosts returns a boolean if a field has been set.

### GetLastSync

`func (o *AnalyticsOverview) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AnalyticsOverview) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AnalyticsOverview) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AnalyticsOverview) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *AnalyticsOverview) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *AnalyticsOverview) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetDataStaleness

`func (o *AnalyticsOverview) GetDataStaleness() AnalyticsOverviewDataStaleness`

GetDataStaleness returns the DataStaleness field if non-nil, zero value otherwise.

### GetDataStalenessOk

`func (o *AnalyticsOverview) GetDataStalenessOk() (*AnalyticsOverviewDataStaleness, bool)`

GetDataStalenessOk returns a tuple with the DataStaleness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataStaleness

`func (o *AnalyticsOverview) SetDataStaleness(v AnalyticsOverviewDataStaleness)`

SetDataStaleness sets DataStaleness field to given value.

### HasDataStaleness

`func (o *AnalyticsOverview) HasDataStaleness() bool`

HasDataStaleness returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


