# UsageStatsUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uploads** | Pointer to **int32** | Stripe users only. Uploads consumed in the current period. | [optional] 
**Profiles** | Pointer to **int32** | Stripe users only. Profiles currently owned. | [optional] 
**LastReset** | Pointer to **time.Time** | Stripe users only. | [optional] 
**ConnectedAccounts** | Pointer to **int32** | Metronome users only. Accounts currently connected across the team. | [optional] 
**XApiCalls** | Pointer to [**UsageStatsUsageXApiCalls**](UsageStatsUsageXApiCalls.md) |  | [optional] 
**XApiCallsByOperation** | Pointer to **map[string]int32** | Metronome users only. Per-operation X API call counts keyed by operation (e.g. &#x60;posts_read&#x60;, &#x60;content_create&#x60;, &#x60;content_create_with_url&#x60;). Resolve each key to price and metadata via &#x60;GET /v1/billing/x-pricing&#x60;. This is the canonical source — covers every price tier including the $0.200 URL tier that &#x60;xApiCalls&#x60; excludes.  | [optional] 

## Methods

### NewUsageStatsUsage

`func NewUsageStatsUsage() *UsageStatsUsage`

NewUsageStatsUsage instantiates a new UsageStatsUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageStatsUsageWithDefaults

`func NewUsageStatsUsageWithDefaults() *UsageStatsUsage`

NewUsageStatsUsageWithDefaults instantiates a new UsageStatsUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploads

`func (o *UsageStatsUsage) GetUploads() int32`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *UsageStatsUsage) GetUploadsOk() (*int32, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *UsageStatsUsage) SetUploads(v int32)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *UsageStatsUsage) HasUploads() bool`

HasUploads returns a boolean if a field has been set.

### GetProfiles

`func (o *UsageStatsUsage) GetProfiles() int32`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *UsageStatsUsage) GetProfilesOk() (*int32, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *UsageStatsUsage) SetProfiles(v int32)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *UsageStatsUsage) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetLastReset

`func (o *UsageStatsUsage) GetLastReset() time.Time`

GetLastReset returns the LastReset field if non-nil, zero value otherwise.

### GetLastResetOk

`func (o *UsageStatsUsage) GetLastResetOk() (*time.Time, bool)`

GetLastResetOk returns a tuple with the LastReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReset

`func (o *UsageStatsUsage) SetLastReset(v time.Time)`

SetLastReset sets LastReset field to given value.

### HasLastReset

`func (o *UsageStatsUsage) HasLastReset() bool`

HasLastReset returns a boolean if a field has been set.

### GetConnectedAccounts

`func (o *UsageStatsUsage) GetConnectedAccounts() int32`

GetConnectedAccounts returns the ConnectedAccounts field if non-nil, zero value otherwise.

### GetConnectedAccountsOk

`func (o *UsageStatsUsage) GetConnectedAccountsOk() (*int32, bool)`

GetConnectedAccountsOk returns a tuple with the ConnectedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAccounts

`func (o *UsageStatsUsage) SetConnectedAccounts(v int32)`

SetConnectedAccounts sets ConnectedAccounts field to given value.

### HasConnectedAccounts

`func (o *UsageStatsUsage) HasConnectedAccounts() bool`

HasConnectedAccounts returns a boolean if a field has been set.

### GetXApiCalls

`func (o *UsageStatsUsage) GetXApiCalls() UsageStatsUsageXApiCalls`

GetXApiCalls returns the XApiCalls field if non-nil, zero value otherwise.

### GetXApiCallsOk

`func (o *UsageStatsUsage) GetXApiCallsOk() (*UsageStatsUsageXApiCalls, bool)`

GetXApiCallsOk returns a tuple with the XApiCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXApiCalls

`func (o *UsageStatsUsage) SetXApiCalls(v UsageStatsUsageXApiCalls)`

SetXApiCalls sets XApiCalls field to given value.

### HasXApiCalls

`func (o *UsageStatsUsage) HasXApiCalls() bool`

HasXApiCalls returns a boolean if a field has been set.

### GetXApiCallsByOperation

`func (o *UsageStatsUsage) GetXApiCallsByOperation() map[string]int32`

GetXApiCallsByOperation returns the XApiCallsByOperation field if non-nil, zero value otherwise.

### GetXApiCallsByOperationOk

`func (o *UsageStatsUsage) GetXApiCallsByOperationOk() (*map[string]int32, bool)`

GetXApiCallsByOperationOk returns a tuple with the XApiCallsByOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXApiCallsByOperation

`func (o *UsageStatsUsage) SetXApiCallsByOperation(v map[string]int32)`

SetXApiCallsByOperation sets XApiCallsByOperation field to given value.

### HasXApiCallsByOperation

`func (o *UsageStatsUsage) HasXApiCallsByOperation() bool`

HasXApiCallsByOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


