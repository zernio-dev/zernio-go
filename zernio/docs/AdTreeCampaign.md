# AdTreeCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformCampaignId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**CampaignName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AdStatus**](AdStatus.md) | Delivery status derived from child ad statuses. Distinct from &#x60;reviewStatus&#x60;, which reflects the platform-side review state. | [optional] 
**ReviewStatus** | Pointer to **NullableString** | Platform-side review state of the campaign. Independent of the children-derived delivery &#x60;status&#x60;: a campaign can have ads already active (status&#x3D;active) while the campaign itself is still being reviewed by the platform (reviewStatus&#x3D;in_review). For Meta, derived from &#x60;effective_status&#x60; + &#x60;issues_info&#x60; on the Campaign, plus ad-level PENDING_REVIEW rollup.  | [optional] 
**PlatformCampaignStatus** | Pointer to **NullableString** | Raw platform-level campaign status (Meta &#x60;effective_status&#x60;: ACTIVE, PAUSED, DELETED, ARCHIVED, IN_PROCESS, WITH_ISSUES). Distinct from per-ad &#x60;platformStatus&#x60;. | [optional] 
**CampaignIssuesInfo** | Pointer to **[]map[string]interface{}** | Platform-reported campaign issues (Meta &#x60;issues_info[]&#x60;). Populated only when the platform has delivery issues to report; contains the specific error codes and messages. | [optional] 
**AdCount** | Pointer to **int32** | Total ads across all ad sets | [optional] 
**AdSetCount** | Pointer to **int32** |  | [optional] 
**Budget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**CampaignBudget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**BudgetLevel** | Pointer to **NullableString** | Canonical CBO/ABO indicator. &#x60;campaign&#x60; &#x3D; CBO (Advantage Campaign Budget, budget lives on the campaign). &#x60;adset&#x60; &#x3D; ABO (budget lives on each ad set). Route budget updates to the matching Meta entity. | [optional] 
**IsBudgetScheduleEnabled** | Pointer to **bool** | Meta-only. Mirrors Campaign.is_budget_schedule_enabled — true when the campaign uses budget scheduling (time-based budget changes). Independent of CBO/ABO. | [optional] [default to false]
**Currency** | Pointer to **NullableString** | ISO 4217 currency code (e.g. USD, EUR, CLP, JPY) for all budget amounts in this campaign node. Budgets are NOT normalized to USD. | [optional] 
**Metrics** | Pointer to [**AdMetrics**](AdMetrics.md) |  | [optional] 
**PlatformAdAccountId** | Pointer to **string** |  | [optional] 
**PlatformAdAccountName** | Pointer to **NullableString** | Human-readable advertiser/account name from the platform. Refreshed on every sync. | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**AdvertisingChannelType** | Pointer to **NullableString** | Google-only. Raw campaign.advertising_channel_type (SEARCH, PERFORMANCE_MAX, VIDEO, DEMAND_GEN, DISPLAY, SHOPPING, ...). Serving surface, distinct from platformObjective (advertiser intent). Null/absent for non-Google platforms. | [optional] 
**PlatformObjective** | Pointer to **NullableString** | Raw Meta campaign objective (e.g. OUTCOME_SALES, OUTCOME_LEADS, OUTCOME_TRAFFIC) | [optional] 
**OptimizationGoal** | Pointer to **NullableString** | Meta optimization goal shared across ad sets, or comma-separated values when ad sets differ (e.g. OFFSITE_CONVERSIONS, VALUE, LEAD_GENERATION) | [optional] 
**BidStrategy** | Pointer to [**NullableBidStrategy**](BidStrategy.md) | Campaign-level bid strategy. Ad sets inherit this unless they override. | [optional] 
**BidAmount** | Pointer to **NullableFloat32** | Representative bid cap for the campaign — bubbled up from the top-spending ad set&#39;s &#x60;bid_amount&#x60; (whole currency units). Populated when the ad-set bidStrategy is LOWEST_COST_WITH_BID_CAP or COST_CAP. | [optional] 
**RoasAverageFloor** | Pointer to **NullableFloat32** | Representative ROAS floor for the campaign — bubbled up from the top-spending ad set. Decimal multiplier (2.0 &#x3D; 2.0x). | [optional] 
**PromotedObject** | Pointer to [**AdTreeAdSetPromotedObject**](AdTreeAdSetPromotedObject.md) |  | [optional] 
**AdSets** | Pointer to [**[]AdTreeAdSet**](AdTreeAdSet.md) |  | [optional] 

## Methods

### NewAdTreeCampaign

`func NewAdTreeCampaign() *AdTreeCampaign`

NewAdTreeCampaign instantiates a new AdTreeCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTreeCampaignWithDefaults

`func NewAdTreeCampaignWithDefaults() *AdTreeCampaign`

NewAdTreeCampaignWithDefaults instantiates a new AdTreeCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformCampaignId

`func (o *AdTreeCampaign) GetPlatformCampaignId() string`

GetPlatformCampaignId returns the PlatformCampaignId field if non-nil, zero value otherwise.

### GetPlatformCampaignIdOk

`func (o *AdTreeCampaign) GetPlatformCampaignIdOk() (*string, bool)`

GetPlatformCampaignIdOk returns a tuple with the PlatformCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignId

`func (o *AdTreeCampaign) SetPlatformCampaignId(v string)`

SetPlatformCampaignId sets PlatformCampaignId field to given value.

### HasPlatformCampaignId

`func (o *AdTreeCampaign) HasPlatformCampaignId() bool`

HasPlatformCampaignId returns a boolean if a field has been set.

### GetPlatform

`func (o *AdTreeCampaign) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AdTreeCampaign) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AdTreeCampaign) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AdTreeCampaign) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCampaignName

`func (o *AdTreeCampaign) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *AdTreeCampaign) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *AdTreeCampaign) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.

### HasCampaignName

`func (o *AdTreeCampaign) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### GetStatus

`func (o *AdTreeCampaign) GetStatus() AdStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdTreeCampaign) GetStatusOk() (*AdStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdTreeCampaign) SetStatus(v AdStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdTreeCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewStatus

`func (o *AdTreeCampaign) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *AdTreeCampaign) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *AdTreeCampaign) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.

### HasReviewStatus

`func (o *AdTreeCampaign) HasReviewStatus() bool`

HasReviewStatus returns a boolean if a field has been set.

### SetReviewStatusNil

`func (o *AdTreeCampaign) SetReviewStatusNil(b bool)`

 SetReviewStatusNil sets the value for ReviewStatus to be an explicit nil

### UnsetReviewStatus
`func (o *AdTreeCampaign) UnsetReviewStatus()`

UnsetReviewStatus ensures that no value is present for ReviewStatus, not even an explicit nil
### GetPlatformCampaignStatus

`func (o *AdTreeCampaign) GetPlatformCampaignStatus() string`

GetPlatformCampaignStatus returns the PlatformCampaignStatus field if non-nil, zero value otherwise.

### GetPlatformCampaignStatusOk

`func (o *AdTreeCampaign) GetPlatformCampaignStatusOk() (*string, bool)`

GetPlatformCampaignStatusOk returns a tuple with the PlatformCampaignStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignStatus

`func (o *AdTreeCampaign) SetPlatformCampaignStatus(v string)`

SetPlatformCampaignStatus sets PlatformCampaignStatus field to given value.

### HasPlatformCampaignStatus

`func (o *AdTreeCampaign) HasPlatformCampaignStatus() bool`

HasPlatformCampaignStatus returns a boolean if a field has been set.

### SetPlatformCampaignStatusNil

`func (o *AdTreeCampaign) SetPlatformCampaignStatusNil(b bool)`

 SetPlatformCampaignStatusNil sets the value for PlatformCampaignStatus to be an explicit nil

### UnsetPlatformCampaignStatus
`func (o *AdTreeCampaign) UnsetPlatformCampaignStatus()`

UnsetPlatformCampaignStatus ensures that no value is present for PlatformCampaignStatus, not even an explicit nil
### GetCampaignIssuesInfo

`func (o *AdTreeCampaign) GetCampaignIssuesInfo() []map[string]interface{}`

GetCampaignIssuesInfo returns the CampaignIssuesInfo field if non-nil, zero value otherwise.

### GetCampaignIssuesInfoOk

`func (o *AdTreeCampaign) GetCampaignIssuesInfoOk() (*[]map[string]interface{}, bool)`

GetCampaignIssuesInfoOk returns a tuple with the CampaignIssuesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIssuesInfo

`func (o *AdTreeCampaign) SetCampaignIssuesInfo(v []map[string]interface{})`

SetCampaignIssuesInfo sets CampaignIssuesInfo field to given value.

### HasCampaignIssuesInfo

`func (o *AdTreeCampaign) HasCampaignIssuesInfo() bool`

HasCampaignIssuesInfo returns a boolean if a field has been set.

### SetCampaignIssuesInfoNil

`func (o *AdTreeCampaign) SetCampaignIssuesInfoNil(b bool)`

 SetCampaignIssuesInfoNil sets the value for CampaignIssuesInfo to be an explicit nil

### UnsetCampaignIssuesInfo
`func (o *AdTreeCampaign) UnsetCampaignIssuesInfo()`

UnsetCampaignIssuesInfo ensures that no value is present for CampaignIssuesInfo, not even an explicit nil
### GetAdCount

`func (o *AdTreeCampaign) GetAdCount() int32`

GetAdCount returns the AdCount field if non-nil, zero value otherwise.

### GetAdCountOk

`func (o *AdTreeCampaign) GetAdCountOk() (*int32, bool)`

GetAdCountOk returns a tuple with the AdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdCount

`func (o *AdTreeCampaign) SetAdCount(v int32)`

SetAdCount sets AdCount field to given value.

### HasAdCount

`func (o *AdTreeCampaign) HasAdCount() bool`

HasAdCount returns a boolean if a field has been set.

### GetAdSetCount

`func (o *AdTreeCampaign) GetAdSetCount() int32`

GetAdSetCount returns the AdSetCount field if non-nil, zero value otherwise.

### GetAdSetCountOk

`func (o *AdTreeCampaign) GetAdSetCountOk() (*int32, bool)`

GetAdSetCountOk returns a tuple with the AdSetCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetCount

`func (o *AdTreeCampaign) SetAdSetCount(v int32)`

SetAdSetCount sets AdSetCount field to given value.

### HasAdSetCount

`func (o *AdTreeCampaign) HasAdSetCount() bool`

HasAdSetCount returns a boolean if a field has been set.

### GetBudget

`func (o *AdTreeCampaign) GetBudget() AdBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdTreeCampaign) GetBudgetOk() (*AdBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdTreeCampaign) SetBudget(v AdBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdTreeCampaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaignBudget

`func (o *AdTreeCampaign) GetCampaignBudget() AdBudget`

GetCampaignBudget returns the CampaignBudget field if non-nil, zero value otherwise.

### GetCampaignBudgetOk

`func (o *AdTreeCampaign) GetCampaignBudgetOk() (*AdBudget, bool)`

GetCampaignBudgetOk returns a tuple with the CampaignBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignBudget

`func (o *AdTreeCampaign) SetCampaignBudget(v AdBudget)`

SetCampaignBudget sets CampaignBudget field to given value.

### HasCampaignBudget

`func (o *AdTreeCampaign) HasCampaignBudget() bool`

HasCampaignBudget returns a boolean if a field has been set.

### GetBudgetLevel

`func (o *AdTreeCampaign) GetBudgetLevel() string`

GetBudgetLevel returns the BudgetLevel field if non-nil, zero value otherwise.

### GetBudgetLevelOk

`func (o *AdTreeCampaign) GetBudgetLevelOk() (*string, bool)`

GetBudgetLevelOk returns a tuple with the BudgetLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLevel

`func (o *AdTreeCampaign) SetBudgetLevel(v string)`

SetBudgetLevel sets BudgetLevel field to given value.

### HasBudgetLevel

`func (o *AdTreeCampaign) HasBudgetLevel() bool`

HasBudgetLevel returns a boolean if a field has been set.

### SetBudgetLevelNil

`func (o *AdTreeCampaign) SetBudgetLevelNil(b bool)`

 SetBudgetLevelNil sets the value for BudgetLevel to be an explicit nil

### UnsetBudgetLevel
`func (o *AdTreeCampaign) UnsetBudgetLevel()`

UnsetBudgetLevel ensures that no value is present for BudgetLevel, not even an explicit nil
### GetIsBudgetScheduleEnabled

`func (o *AdTreeCampaign) GetIsBudgetScheduleEnabled() bool`

GetIsBudgetScheduleEnabled returns the IsBudgetScheduleEnabled field if non-nil, zero value otherwise.

### GetIsBudgetScheduleEnabledOk

`func (o *AdTreeCampaign) GetIsBudgetScheduleEnabledOk() (*bool, bool)`

GetIsBudgetScheduleEnabledOk returns a tuple with the IsBudgetScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBudgetScheduleEnabled

`func (o *AdTreeCampaign) SetIsBudgetScheduleEnabled(v bool)`

SetIsBudgetScheduleEnabled sets IsBudgetScheduleEnabled field to given value.

### HasIsBudgetScheduleEnabled

`func (o *AdTreeCampaign) HasIsBudgetScheduleEnabled() bool`

HasIsBudgetScheduleEnabled returns a boolean if a field has been set.

### GetCurrency

`func (o *AdTreeCampaign) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdTreeCampaign) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdTreeCampaign) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AdTreeCampaign) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *AdTreeCampaign) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *AdTreeCampaign) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetMetrics

`func (o *AdTreeCampaign) GetMetrics() AdMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdTreeCampaign) GetMetricsOk() (*AdMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdTreeCampaign) SetMetrics(v AdMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdTreeCampaign) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPlatformAdAccountId

`func (o *AdTreeCampaign) GetPlatformAdAccountId() string`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *AdTreeCampaign) GetPlatformAdAccountIdOk() (*string, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *AdTreeCampaign) SetPlatformAdAccountId(v string)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.

### HasPlatformAdAccountId

`func (o *AdTreeCampaign) HasPlatformAdAccountId() bool`

HasPlatformAdAccountId returns a boolean if a field has been set.

### GetPlatformAdAccountName

`func (o *AdTreeCampaign) GetPlatformAdAccountName() string`

GetPlatformAdAccountName returns the PlatformAdAccountName field if non-nil, zero value otherwise.

### GetPlatformAdAccountNameOk

`func (o *AdTreeCampaign) GetPlatformAdAccountNameOk() (*string, bool)`

GetPlatformAdAccountNameOk returns a tuple with the PlatformAdAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountName

`func (o *AdTreeCampaign) SetPlatformAdAccountName(v string)`

SetPlatformAdAccountName sets PlatformAdAccountName field to given value.

### HasPlatformAdAccountName

`func (o *AdTreeCampaign) HasPlatformAdAccountName() bool`

HasPlatformAdAccountName returns a boolean if a field has been set.

### SetPlatformAdAccountNameNil

`func (o *AdTreeCampaign) SetPlatformAdAccountNameNil(b bool)`

 SetPlatformAdAccountNameNil sets the value for PlatformAdAccountName to be an explicit nil

### UnsetPlatformAdAccountName
`func (o *AdTreeCampaign) UnsetPlatformAdAccountName()`

UnsetPlatformAdAccountName ensures that no value is present for PlatformAdAccountName, not even an explicit nil
### GetAccountId

`func (o *AdTreeCampaign) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AdTreeCampaign) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AdTreeCampaign) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AdTreeCampaign) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProfileId

`func (o *AdTreeCampaign) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AdTreeCampaign) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AdTreeCampaign) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AdTreeCampaign) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetAdvertisingChannelType

`func (o *AdTreeCampaign) GetAdvertisingChannelType() string`

GetAdvertisingChannelType returns the AdvertisingChannelType field if non-nil, zero value otherwise.

### GetAdvertisingChannelTypeOk

`func (o *AdTreeCampaign) GetAdvertisingChannelTypeOk() (*string, bool)`

GetAdvertisingChannelTypeOk returns a tuple with the AdvertisingChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisingChannelType

`func (o *AdTreeCampaign) SetAdvertisingChannelType(v string)`

SetAdvertisingChannelType sets AdvertisingChannelType field to given value.

### HasAdvertisingChannelType

`func (o *AdTreeCampaign) HasAdvertisingChannelType() bool`

HasAdvertisingChannelType returns a boolean if a field has been set.

### SetAdvertisingChannelTypeNil

`func (o *AdTreeCampaign) SetAdvertisingChannelTypeNil(b bool)`

 SetAdvertisingChannelTypeNil sets the value for AdvertisingChannelType to be an explicit nil

### UnsetAdvertisingChannelType
`func (o *AdTreeCampaign) UnsetAdvertisingChannelType()`

UnsetAdvertisingChannelType ensures that no value is present for AdvertisingChannelType, not even an explicit nil
### GetPlatformObjective

`func (o *AdTreeCampaign) GetPlatformObjective() string`

GetPlatformObjective returns the PlatformObjective field if non-nil, zero value otherwise.

### GetPlatformObjectiveOk

`func (o *AdTreeCampaign) GetPlatformObjectiveOk() (*string, bool)`

GetPlatformObjectiveOk returns a tuple with the PlatformObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformObjective

`func (o *AdTreeCampaign) SetPlatformObjective(v string)`

SetPlatformObjective sets PlatformObjective field to given value.

### HasPlatformObjective

`func (o *AdTreeCampaign) HasPlatformObjective() bool`

HasPlatformObjective returns a boolean if a field has been set.

### SetPlatformObjectiveNil

`func (o *AdTreeCampaign) SetPlatformObjectiveNil(b bool)`

 SetPlatformObjectiveNil sets the value for PlatformObjective to be an explicit nil

### UnsetPlatformObjective
`func (o *AdTreeCampaign) UnsetPlatformObjective()`

UnsetPlatformObjective ensures that no value is present for PlatformObjective, not even an explicit nil
### GetOptimizationGoal

`func (o *AdTreeCampaign) GetOptimizationGoal() string`

GetOptimizationGoal returns the OptimizationGoal field if non-nil, zero value otherwise.

### GetOptimizationGoalOk

`func (o *AdTreeCampaign) GetOptimizationGoalOk() (*string, bool)`

GetOptimizationGoalOk returns a tuple with the OptimizationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationGoal

`func (o *AdTreeCampaign) SetOptimizationGoal(v string)`

SetOptimizationGoal sets OptimizationGoal field to given value.

### HasOptimizationGoal

`func (o *AdTreeCampaign) HasOptimizationGoal() bool`

HasOptimizationGoal returns a boolean if a field has been set.

### SetOptimizationGoalNil

`func (o *AdTreeCampaign) SetOptimizationGoalNil(b bool)`

 SetOptimizationGoalNil sets the value for OptimizationGoal to be an explicit nil

### UnsetOptimizationGoal
`func (o *AdTreeCampaign) UnsetOptimizationGoal()`

UnsetOptimizationGoal ensures that no value is present for OptimizationGoal, not even an explicit nil
### GetBidStrategy

`func (o *AdTreeCampaign) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *AdTreeCampaign) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *AdTreeCampaign) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *AdTreeCampaign) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### SetBidStrategyNil

`func (o *AdTreeCampaign) SetBidStrategyNil(b bool)`

 SetBidStrategyNil sets the value for BidStrategy to be an explicit nil

### UnsetBidStrategy
`func (o *AdTreeCampaign) UnsetBidStrategy()`

UnsetBidStrategy ensures that no value is present for BidStrategy, not even an explicit nil
### GetBidAmount

`func (o *AdTreeCampaign) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *AdTreeCampaign) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *AdTreeCampaign) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *AdTreeCampaign) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### SetBidAmountNil

`func (o *AdTreeCampaign) SetBidAmountNil(b bool)`

 SetBidAmountNil sets the value for BidAmount to be an explicit nil

### UnsetBidAmount
`func (o *AdTreeCampaign) UnsetBidAmount()`

UnsetBidAmount ensures that no value is present for BidAmount, not even an explicit nil
### GetRoasAverageFloor

`func (o *AdTreeCampaign) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *AdTreeCampaign) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *AdTreeCampaign) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *AdTreeCampaign) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### SetRoasAverageFloorNil

`func (o *AdTreeCampaign) SetRoasAverageFloorNil(b bool)`

 SetRoasAverageFloorNil sets the value for RoasAverageFloor to be an explicit nil

### UnsetRoasAverageFloor
`func (o *AdTreeCampaign) UnsetRoasAverageFloor()`

UnsetRoasAverageFloor ensures that no value is present for RoasAverageFloor, not even an explicit nil
### GetPromotedObject

`func (o *AdTreeCampaign) GetPromotedObject() AdTreeAdSetPromotedObject`

GetPromotedObject returns the PromotedObject field if non-nil, zero value otherwise.

### GetPromotedObjectOk

`func (o *AdTreeCampaign) GetPromotedObjectOk() (*AdTreeAdSetPromotedObject, bool)`

GetPromotedObjectOk returns a tuple with the PromotedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedObject

`func (o *AdTreeCampaign) SetPromotedObject(v AdTreeAdSetPromotedObject)`

SetPromotedObject sets PromotedObject field to given value.

### HasPromotedObject

`func (o *AdTreeCampaign) HasPromotedObject() bool`

HasPromotedObject returns a boolean if a field has been set.

### GetAdSets

`func (o *AdTreeCampaign) GetAdSets() []AdTreeAdSet`

GetAdSets returns the AdSets field if non-nil, zero value otherwise.

### GetAdSetsOk

`func (o *AdTreeCampaign) GetAdSetsOk() (*[]AdTreeAdSet, bool)`

GetAdSetsOk returns a tuple with the AdSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSets

`func (o *AdTreeCampaign) SetAdSets(v []AdTreeAdSet)`

SetAdSets sets AdSets field to given value.

### HasAdSets

`func (o *AdTreeCampaign) HasAdSets() bool`

HasAdSets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


