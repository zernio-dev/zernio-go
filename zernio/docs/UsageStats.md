# UsageStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingSystem** | Pointer to **string** | Which billing system the account is on. Shape of &#x60;usage&#x60;/&#x60;spend&#x60; differs. | [optional] 
**PlanName** | Pointer to **string** |  | [optional] 
**BillingPeriod** | Pointer to **string** |  | [optional] 
**SignupDate** | Pointer to **time.Time** |  | [optional] 
**BillingAnchorDay** | Pointer to **int32** | Day of month (1-31) when the billing cycle resets | [optional] 
**HasAccess** | Pointer to **bool** | True if the account is in good standing. False for past-due/unpaid/paused subscriptions. | [optional] 
**CustomerId** | Pointer to **NullableString** | Stripe customer ID, when present. | [optional] 
**IsInvitedUser** | Pointer to **bool** | True if this is a team member; limits/usage reflect the account owner. | [optional] 
**AutoUpgradeEnabled** | Pointer to **bool** | Stripe-only. Always false for Metronome users. | [optional] 
**Limits** | Pointer to [**UsageStatsLimits**](UsageStatsLimits.md) |  | [optional] 
**Usage** | Pointer to [**UsageStatsUsage**](UsageStatsUsage.md) |  | [optional] 
**Spend** | Pointer to [**UsageStatsSpend**](UsageStatsSpend.md) |  | [optional] 

## Methods

### NewUsageStats

`func NewUsageStats() *UsageStats`

NewUsageStats instantiates a new UsageStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageStatsWithDefaults

`func NewUsageStatsWithDefaults() *UsageStats`

NewUsageStatsWithDefaults instantiates a new UsageStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingSystem

`func (o *UsageStats) GetBillingSystem() string`

GetBillingSystem returns the BillingSystem field if non-nil, zero value otherwise.

### GetBillingSystemOk

`func (o *UsageStats) GetBillingSystemOk() (*string, bool)`

GetBillingSystemOk returns a tuple with the BillingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingSystem

`func (o *UsageStats) SetBillingSystem(v string)`

SetBillingSystem sets BillingSystem field to given value.

### HasBillingSystem

`func (o *UsageStats) HasBillingSystem() bool`

HasBillingSystem returns a boolean if a field has been set.

### GetPlanName

`func (o *UsageStats) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *UsageStats) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *UsageStats) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *UsageStats) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetBillingPeriod

`func (o *UsageStats) GetBillingPeriod() string`

GetBillingPeriod returns the BillingPeriod field if non-nil, zero value otherwise.

### GetBillingPeriodOk

`func (o *UsageStats) GetBillingPeriodOk() (*string, bool)`

GetBillingPeriodOk returns a tuple with the BillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingPeriod

`func (o *UsageStats) SetBillingPeriod(v string)`

SetBillingPeriod sets BillingPeriod field to given value.

### HasBillingPeriod

`func (o *UsageStats) HasBillingPeriod() bool`

HasBillingPeriod returns a boolean if a field has been set.

### GetSignupDate

`func (o *UsageStats) GetSignupDate() time.Time`

GetSignupDate returns the SignupDate field if non-nil, zero value otherwise.

### GetSignupDateOk

`func (o *UsageStats) GetSignupDateOk() (*time.Time, bool)`

GetSignupDateOk returns a tuple with the SignupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignupDate

`func (o *UsageStats) SetSignupDate(v time.Time)`

SetSignupDate sets SignupDate field to given value.

### HasSignupDate

`func (o *UsageStats) HasSignupDate() bool`

HasSignupDate returns a boolean if a field has been set.

### GetBillingAnchorDay

`func (o *UsageStats) GetBillingAnchorDay() int32`

GetBillingAnchorDay returns the BillingAnchorDay field if non-nil, zero value otherwise.

### GetBillingAnchorDayOk

`func (o *UsageStats) GetBillingAnchorDayOk() (*int32, bool)`

GetBillingAnchorDayOk returns a tuple with the BillingAnchorDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAnchorDay

`func (o *UsageStats) SetBillingAnchorDay(v int32)`

SetBillingAnchorDay sets BillingAnchorDay field to given value.

### HasBillingAnchorDay

`func (o *UsageStats) HasBillingAnchorDay() bool`

HasBillingAnchorDay returns a boolean if a field has been set.

### GetHasAccess

`func (o *UsageStats) GetHasAccess() bool`

GetHasAccess returns the HasAccess field if non-nil, zero value otherwise.

### GetHasAccessOk

`func (o *UsageStats) GetHasAccessOk() (*bool, bool)`

GetHasAccessOk returns a tuple with the HasAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAccess

`func (o *UsageStats) SetHasAccess(v bool)`

SetHasAccess sets HasAccess field to given value.

### HasHasAccess

`func (o *UsageStats) HasHasAccess() bool`

HasHasAccess returns a boolean if a field has been set.

### GetCustomerId

`func (o *UsageStats) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *UsageStats) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *UsageStats) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *UsageStats) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### SetCustomerIdNil

`func (o *UsageStats) SetCustomerIdNil(b bool)`

 SetCustomerIdNil sets the value for CustomerId to be an explicit nil

### UnsetCustomerId
`func (o *UsageStats) UnsetCustomerId()`

UnsetCustomerId ensures that no value is present for CustomerId, not even an explicit nil
### GetIsInvitedUser

`func (o *UsageStats) GetIsInvitedUser() bool`

GetIsInvitedUser returns the IsInvitedUser field if non-nil, zero value otherwise.

### GetIsInvitedUserOk

`func (o *UsageStats) GetIsInvitedUserOk() (*bool, bool)`

GetIsInvitedUserOk returns a tuple with the IsInvitedUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInvitedUser

`func (o *UsageStats) SetIsInvitedUser(v bool)`

SetIsInvitedUser sets IsInvitedUser field to given value.

### HasIsInvitedUser

`func (o *UsageStats) HasIsInvitedUser() bool`

HasIsInvitedUser returns a boolean if a field has been set.

### GetAutoUpgradeEnabled

`func (o *UsageStats) GetAutoUpgradeEnabled() bool`

GetAutoUpgradeEnabled returns the AutoUpgradeEnabled field if non-nil, zero value otherwise.

### GetAutoUpgradeEnabledOk

`func (o *UsageStats) GetAutoUpgradeEnabledOk() (*bool, bool)`

GetAutoUpgradeEnabledOk returns a tuple with the AutoUpgradeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoUpgradeEnabled

`func (o *UsageStats) SetAutoUpgradeEnabled(v bool)`

SetAutoUpgradeEnabled sets AutoUpgradeEnabled field to given value.

### HasAutoUpgradeEnabled

`func (o *UsageStats) HasAutoUpgradeEnabled() bool`

HasAutoUpgradeEnabled returns a boolean if a field has been set.

### GetLimits

`func (o *UsageStats) GetLimits() UsageStatsLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *UsageStats) GetLimitsOk() (*UsageStatsLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *UsageStats) SetLimits(v UsageStatsLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *UsageStats) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetUsage

`func (o *UsageStats) GetUsage() UsageStatsUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *UsageStats) GetUsageOk() (*UsageStatsUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *UsageStats) SetUsage(v UsageStatsUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *UsageStats) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetSpend

`func (o *UsageStats) GetSpend() UsageStatsSpend`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *UsageStats) GetSpendOk() (*UsageStatsSpend, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *UsageStats) SetSpend(v UsageStatsSpend)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *UsageStats) HasSpend() bool`

HasSpend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


