# DuplicateAdCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**DeepCopy** | Pointer to **bool** | Copy child ad sets + ads + creatives + targeting | [optional] [default to true]
**StatusOption** | Pointer to **string** |  | [optional] [default to "PAUSED"]
**StartTime** | Pointer to **time.Time** | Reschedule the copied hierarchy&#39;s start time | [optional] 
**EndTime** | Pointer to **time.Time** |  | [optional] 
**RenameStrategy** | Pointer to **string** |  | [optional] 
**RenamePrefix** | Pointer to **string** |  | [optional] 
**RenameSuffix** | Pointer to **string** |  | [optional] 
**SyncAfter** | Pointer to **bool** | Trigger ads discovery on the owning account after the copy succeeds | [optional] [default to true]

## Methods

### NewDuplicateAdCampaignRequest

`func NewDuplicateAdCampaignRequest(platform string, ) *DuplicateAdCampaignRequest`

NewDuplicateAdCampaignRequest instantiates a new DuplicateAdCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateAdCampaignRequestWithDefaults

`func NewDuplicateAdCampaignRequestWithDefaults() *DuplicateAdCampaignRequest`

NewDuplicateAdCampaignRequestWithDefaults instantiates a new DuplicateAdCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *DuplicateAdCampaignRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *DuplicateAdCampaignRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *DuplicateAdCampaignRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetDeepCopy

`func (o *DuplicateAdCampaignRequest) GetDeepCopy() bool`

GetDeepCopy returns the DeepCopy field if non-nil, zero value otherwise.

### GetDeepCopyOk

`func (o *DuplicateAdCampaignRequest) GetDeepCopyOk() (*bool, bool)`

GetDeepCopyOk returns a tuple with the DeepCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeepCopy

`func (o *DuplicateAdCampaignRequest) SetDeepCopy(v bool)`

SetDeepCopy sets DeepCopy field to given value.

### HasDeepCopy

`func (o *DuplicateAdCampaignRequest) HasDeepCopy() bool`

HasDeepCopy returns a boolean if a field has been set.

### GetStatusOption

`func (o *DuplicateAdCampaignRequest) GetStatusOption() string`

GetStatusOption returns the StatusOption field if non-nil, zero value otherwise.

### GetStatusOptionOk

`func (o *DuplicateAdCampaignRequest) GetStatusOptionOk() (*string, bool)`

GetStatusOptionOk returns a tuple with the StatusOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusOption

`func (o *DuplicateAdCampaignRequest) SetStatusOption(v string)`

SetStatusOption sets StatusOption field to given value.

### HasStatusOption

`func (o *DuplicateAdCampaignRequest) HasStatusOption() bool`

HasStatusOption returns a boolean if a field has been set.

### GetStartTime

`func (o *DuplicateAdCampaignRequest) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DuplicateAdCampaignRequest) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DuplicateAdCampaignRequest) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DuplicateAdCampaignRequest) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *DuplicateAdCampaignRequest) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *DuplicateAdCampaignRequest) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *DuplicateAdCampaignRequest) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *DuplicateAdCampaignRequest) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetRenameStrategy

`func (o *DuplicateAdCampaignRequest) GetRenameStrategy() string`

GetRenameStrategy returns the RenameStrategy field if non-nil, zero value otherwise.

### GetRenameStrategyOk

`func (o *DuplicateAdCampaignRequest) GetRenameStrategyOk() (*string, bool)`

GetRenameStrategyOk returns a tuple with the RenameStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameStrategy

`func (o *DuplicateAdCampaignRequest) SetRenameStrategy(v string)`

SetRenameStrategy sets RenameStrategy field to given value.

### HasRenameStrategy

`func (o *DuplicateAdCampaignRequest) HasRenameStrategy() bool`

HasRenameStrategy returns a boolean if a field has been set.

### GetRenamePrefix

`func (o *DuplicateAdCampaignRequest) GetRenamePrefix() string`

GetRenamePrefix returns the RenamePrefix field if non-nil, zero value otherwise.

### GetRenamePrefixOk

`func (o *DuplicateAdCampaignRequest) GetRenamePrefixOk() (*string, bool)`

GetRenamePrefixOk returns a tuple with the RenamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenamePrefix

`func (o *DuplicateAdCampaignRequest) SetRenamePrefix(v string)`

SetRenamePrefix sets RenamePrefix field to given value.

### HasRenamePrefix

`func (o *DuplicateAdCampaignRequest) HasRenamePrefix() bool`

HasRenamePrefix returns a boolean if a field has been set.

### GetRenameSuffix

`func (o *DuplicateAdCampaignRequest) GetRenameSuffix() string`

GetRenameSuffix returns the RenameSuffix field if non-nil, zero value otherwise.

### GetRenameSuffixOk

`func (o *DuplicateAdCampaignRequest) GetRenameSuffixOk() (*string, bool)`

GetRenameSuffixOk returns a tuple with the RenameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenameSuffix

`func (o *DuplicateAdCampaignRequest) SetRenameSuffix(v string)`

SetRenameSuffix sets RenameSuffix field to given value.

### HasRenameSuffix

`func (o *DuplicateAdCampaignRequest) HasRenameSuffix() bool`

HasRenameSuffix returns a boolean if a field has been set.

### GetSyncAfter

`func (o *DuplicateAdCampaignRequest) GetSyncAfter() bool`

GetSyncAfter returns the SyncAfter field if non-nil, zero value otherwise.

### GetSyncAfterOk

`func (o *DuplicateAdCampaignRequest) GetSyncAfterOk() (*bool, bool)`

GetSyncAfterOk returns a tuple with the SyncAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncAfter

`func (o *DuplicateAdCampaignRequest) SetSyncAfter(v bool)`

SetSyncAfter sets SyncAfter field to given value.

### HasSyncAfter

`func (o *DuplicateAdCampaignRequest) HasSyncAfter() bool`

HasSyncAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


