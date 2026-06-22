# AnalyticsOverviewDataStaleness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaleAccountCount** | Pointer to **int32** | Number of accounts with stale analytics data | [optional] 
**SyncTriggered** | Pointer to **bool** | Whether a background sync was triggered for stale accounts | [optional] 

## Methods

### NewAnalyticsOverviewDataStaleness

`func NewAnalyticsOverviewDataStaleness() *AnalyticsOverviewDataStaleness`

NewAnalyticsOverviewDataStaleness instantiates a new AnalyticsOverviewDataStaleness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsOverviewDataStalenessWithDefaults

`func NewAnalyticsOverviewDataStalenessWithDefaults() *AnalyticsOverviewDataStaleness`

NewAnalyticsOverviewDataStalenessWithDefaults instantiates a new AnalyticsOverviewDataStaleness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaleAccountCount

`func (o *AnalyticsOverviewDataStaleness) GetStaleAccountCount() int32`

GetStaleAccountCount returns the StaleAccountCount field if non-nil, zero value otherwise.

### GetStaleAccountCountOk

`func (o *AnalyticsOverviewDataStaleness) GetStaleAccountCountOk() (*int32, bool)`

GetStaleAccountCountOk returns a tuple with the StaleAccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleAccountCount

`func (o *AnalyticsOverviewDataStaleness) SetStaleAccountCount(v int32)`

SetStaleAccountCount sets StaleAccountCount field to given value.

### HasStaleAccountCount

`func (o *AnalyticsOverviewDataStaleness) HasStaleAccountCount() bool`

HasStaleAccountCount returns a boolean if a field has been set.

### GetSyncTriggered

`func (o *AnalyticsOverviewDataStaleness) GetSyncTriggered() bool`

GetSyncTriggered returns the SyncTriggered field if non-nil, zero value otherwise.

### GetSyncTriggeredOk

`func (o *AnalyticsOverviewDataStaleness) GetSyncTriggeredOk() (*bool, bool)`

GetSyncTriggeredOk returns a tuple with the SyncTriggered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncTriggered

`func (o *AnalyticsOverviewDataStaleness) SetSyncTriggered(v bool)`

SetSyncTriggered sets SyncTriggered field to given value.

### HasSyncTriggered

`func (o *AnalyticsOverviewDataStaleness) HasSyncTriggered() bool`

HasSyncTriggered returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


