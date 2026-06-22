# GetFollowerStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountWithFollowerStats**](AccountWithFollowerStats.md) |  | [optional] 
**Stats** | Pointer to [**map[string][]GetFollowerStats200ResponseStatsValueInner**](array.md) |  | [optional] 
**DateRange** | Pointer to [**GetFollowerStats200ResponseDateRange**](GetFollowerStats200ResponseDateRange.md) |  | [optional] 
**Granularity** | Pointer to **string** |  | [optional] 

## Methods

### NewGetFollowerStats200Response

`func NewGetFollowerStats200Response() *GetFollowerStats200Response`

NewGetFollowerStats200Response instantiates a new GetFollowerStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFollowerStats200ResponseWithDefaults

`func NewGetFollowerStats200ResponseWithDefaults() *GetFollowerStats200Response`

NewGetFollowerStats200ResponseWithDefaults instantiates a new GetFollowerStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *GetFollowerStats200Response) GetAccounts() []AccountWithFollowerStats`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetFollowerStats200Response) GetAccountsOk() (*[]AccountWithFollowerStats, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetFollowerStats200Response) SetAccounts(v []AccountWithFollowerStats)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetFollowerStats200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetStats

`func (o *GetFollowerStats200Response) GetStats() map[string][]GetFollowerStats200ResponseStatsValueInner`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetFollowerStats200Response) GetStatsOk() (*map[string][]GetFollowerStats200ResponseStatsValueInner, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetFollowerStats200Response) SetStats(v map[string][]GetFollowerStats200ResponseStatsValueInner)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetFollowerStats200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetDateRange

`func (o *GetFollowerStats200Response) GetDateRange() GetFollowerStats200ResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *GetFollowerStats200Response) GetDateRangeOk() (*GetFollowerStats200ResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *GetFollowerStats200Response) SetDateRange(v GetFollowerStats200ResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *GetFollowerStats200Response) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetGranularity

`func (o *GetFollowerStats200Response) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GetFollowerStats200Response) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GetFollowerStats200Response) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *GetFollowerStats200Response) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


