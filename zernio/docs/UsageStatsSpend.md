# UsageStatsSpend

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPeriodCents** | Pointer to **int32** | Total current-period spend in cents (all products combined). | [optional] 
**CreditsRemainingCents** | Pointer to **int32** | Free-tier credit remaining in cents. Applied before any charge. | [optional] 
**XSpendCents** | Pointer to **int32** | Current-period X/Twitter API spend in cents, summed from &#x60;xApiCallsByOperation&#x60; × per-operation prices. Tier-agnostic (covers every price including the $0.200 URL tier). Rounded up for conservative enforcement against &#x60;xSpendLimitCents&#x60;.  | [optional] 
**XSpendLimitCents** | Pointer to **NullableInt32** | Monthly X spend cap set by the account owner, or null if no cap. When current X spend hits this cap, analytics and inbox sync are auto-paused for X accounts. Publishing is never blocked by this cap.  | [optional] 

## Methods

### NewUsageStatsSpend

`func NewUsageStatsSpend() *UsageStatsSpend`

NewUsageStatsSpend instantiates a new UsageStatsSpend object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageStatsSpendWithDefaults

`func NewUsageStatsSpendWithDefaults() *UsageStatsSpend`

NewUsageStatsSpendWithDefaults instantiates a new UsageStatsSpend object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPeriodCents

`func (o *UsageStatsSpend) GetCurrentPeriodCents() int32`

GetCurrentPeriodCents returns the CurrentPeriodCents field if non-nil, zero value otherwise.

### GetCurrentPeriodCentsOk

`func (o *UsageStatsSpend) GetCurrentPeriodCentsOk() (*int32, bool)`

GetCurrentPeriodCentsOk returns a tuple with the CurrentPeriodCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPeriodCents

`func (o *UsageStatsSpend) SetCurrentPeriodCents(v int32)`

SetCurrentPeriodCents sets CurrentPeriodCents field to given value.

### HasCurrentPeriodCents

`func (o *UsageStatsSpend) HasCurrentPeriodCents() bool`

HasCurrentPeriodCents returns a boolean if a field has been set.

### GetCreditsRemainingCents

`func (o *UsageStatsSpend) GetCreditsRemainingCents() int32`

GetCreditsRemainingCents returns the CreditsRemainingCents field if non-nil, zero value otherwise.

### GetCreditsRemainingCentsOk

`func (o *UsageStatsSpend) GetCreditsRemainingCentsOk() (*int32, bool)`

GetCreditsRemainingCentsOk returns a tuple with the CreditsRemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsRemainingCents

`func (o *UsageStatsSpend) SetCreditsRemainingCents(v int32)`

SetCreditsRemainingCents sets CreditsRemainingCents field to given value.

### HasCreditsRemainingCents

`func (o *UsageStatsSpend) HasCreditsRemainingCents() bool`

HasCreditsRemainingCents returns a boolean if a field has been set.

### GetXSpendCents

`func (o *UsageStatsSpend) GetXSpendCents() int32`

GetXSpendCents returns the XSpendCents field if non-nil, zero value otherwise.

### GetXSpendCentsOk

`func (o *UsageStatsSpend) GetXSpendCentsOk() (*int32, bool)`

GetXSpendCentsOk returns a tuple with the XSpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXSpendCents

`func (o *UsageStatsSpend) SetXSpendCents(v int32)`

SetXSpendCents sets XSpendCents field to given value.

### HasXSpendCents

`func (o *UsageStatsSpend) HasXSpendCents() bool`

HasXSpendCents returns a boolean if a field has been set.

### GetXSpendLimitCents

`func (o *UsageStatsSpend) GetXSpendLimitCents() int32`

GetXSpendLimitCents returns the XSpendLimitCents field if non-nil, zero value otherwise.

### GetXSpendLimitCentsOk

`func (o *UsageStatsSpend) GetXSpendLimitCentsOk() (*int32, bool)`

GetXSpendLimitCentsOk returns a tuple with the XSpendLimitCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXSpendLimitCents

`func (o *UsageStatsSpend) SetXSpendLimitCents(v int32)`

SetXSpendLimitCents sets XSpendLimitCents field to given value.

### HasXSpendLimitCents

`func (o *UsageStatsSpend) HasXSpendLimitCents() bool`

HasXSpendLimitCents returns a boolean if a field has been set.

### SetXSpendLimitCentsNil

`func (o *UsageStatsSpend) SetXSpendLimitCentsNil(b bool)`

 SetXSpendLimitCentsNil sets the value for XSpendLimitCents to be an explicit nil

### UnsetXSpendLimitCents
`func (o *UsageStatsSpend) UnsetXSpendLimitCents()`

UnsetXSpendLimitCents ensures that no value is present for XSpendLimitCents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


