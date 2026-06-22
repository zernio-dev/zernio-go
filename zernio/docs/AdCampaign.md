# AdCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformCampaignId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**CampaignName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AdStatus**](AdStatus.md) | Delivery status derived from child ad statuses. Distinct from &#x60;reviewStatus&#x60;. | [optional] 
**ReviewStatus** | Pointer to **NullableString** | Platform-side review state of the campaign. See AdTreeCampaign.reviewStatus for the full description. | [optional] 
**PlatformCampaignStatus** | Pointer to **NullableString** | Raw platform-level campaign status (Meta &#x60;effective_status&#x60;). | [optional] 
**CampaignIssuesInfo** | Pointer to **[]map[string]interface{}** | Platform-reported campaign issues (Meta &#x60;issues_info[]&#x60;). | [optional] 
**AdCount** | Pointer to **int32** |  | [optional] 
**Budget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**CampaignBudget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**BudgetLevel** | Pointer to **NullableString** | Canonical CBO/ABO indicator. See AdTreeCampaign.budgetLevel. | [optional] 
**IsBudgetScheduleEnabled** | Pointer to **bool** | Meta-only. Mirrors Campaign.is_budget_schedule_enabled. | [optional] [default to false]
**Currency** | Pointer to **NullableString** | ISO 4217 currency code for all budget amounts. Budgets are NOT normalized to USD. | [optional] 
**Metrics** | Pointer to [**AdMetrics**](AdMetrics.md) |  | [optional] 
**PlatformAdAccountId** | Pointer to **string** |  | [optional] 
**PlatformAdAccountName** | Pointer to **NullableString** | Human-readable advertiser/account name from the platform. Refreshed on every sync. | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**AdvertisingChannelType** | Pointer to **NullableString** | Google-only. Raw campaign.advertising_channel_type. See AdTreeCampaign.advertisingChannelType. | [optional] 
**PlatformObjective** | Pointer to **NullableString** | Raw Meta campaign objective (e.g. OUTCOME_SALES, OUTCOME_LEADS, OUTCOME_TRAFFIC) | [optional] 
**OptimizationGoal** | Pointer to **NullableString** | Meta optimization goal shared across ad sets, or comma-separated values when ad sets differ (e.g. OFFSITE_CONVERSIONS, VALUE, LEAD_GENERATION) | [optional] 
**BidStrategy** | Pointer to [**NullableBidStrategy**](BidStrategy.md) | Campaign-level bid strategy. Ad sets inherit this unless they override. | [optional] 
**BidAmount** | Pointer to **NullableFloat32** | Representative bid cap from the top-spending ad set (whole currency units). Populated when bidStrategy is LOWEST_COST_WITH_BID_CAP or COST_CAP. | [optional] 
**RoasAverageFloor** | Pointer to **NullableFloat32** | Representative ROAS floor from the top-spending ad set. Decimal multiplier (2.0 &#x3D; 2.0x). | [optional] 
**PromotedObject** | Pointer to [**AdTreeAdSetPromotedObject**](AdTreeAdSetPromotedObject.md) |  | [optional] 
**EarliestAd** | Pointer to **time.Time** |  | [optional] 
**LatestAd** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAdCampaign

`func NewAdCampaign() *AdCampaign`

NewAdCampaign instantiates a new AdCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCampaignWithDefaults

`func NewAdCampaignWithDefaults() *AdCampaign`

NewAdCampaignWithDefaults instantiates a new AdCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformCampaignId

`func (o *AdCampaign) GetPlatformCampaignId() string`

GetPlatformCampaignId returns the PlatformCampaignId field if non-nil, zero value otherwise.

### GetPlatformCampaignIdOk

`func (o *AdCampaign) GetPlatformCampaignIdOk() (*string, bool)`

GetPlatformCampaignIdOk returns a tuple with the PlatformCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignId

`func (o *AdCampaign) SetPlatformCampaignId(v string)`

SetPlatformCampaignId sets PlatformCampaignId field to given value.

### HasPlatformCampaignId

`func (o *AdCampaign) HasPlatformCampaignId() bool`

HasPlatformCampaignId returns a boolean if a field has been set.

### GetPlatform

`func (o *AdCampaign) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AdCampaign) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AdCampaign) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AdCampaign) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetCampaignName

`func (o *AdCampaign) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *AdCampaign) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *AdCampaign) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.

### HasCampaignName

`func (o *AdCampaign) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### GetStatus

`func (o *AdCampaign) GetStatus() AdStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdCampaign) GetStatusOk() (*AdStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdCampaign) SetStatus(v AdStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReviewStatus

`func (o *AdCampaign) GetReviewStatus() string`

GetReviewStatus returns the ReviewStatus field if non-nil, zero value otherwise.

### GetReviewStatusOk

`func (o *AdCampaign) GetReviewStatusOk() (*string, bool)`

GetReviewStatusOk returns a tuple with the ReviewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewStatus

`func (o *AdCampaign) SetReviewStatus(v string)`

SetReviewStatus sets ReviewStatus field to given value.

### HasReviewStatus

`func (o *AdCampaign) HasReviewStatus() bool`

HasReviewStatus returns a boolean if a field has been set.

### SetReviewStatusNil

`func (o *AdCampaign) SetReviewStatusNil(b bool)`

 SetReviewStatusNil sets the value for ReviewStatus to be an explicit nil

### UnsetReviewStatus
`func (o *AdCampaign) UnsetReviewStatus()`

UnsetReviewStatus ensures that no value is present for ReviewStatus, not even an explicit nil
### GetPlatformCampaignStatus

`func (o *AdCampaign) GetPlatformCampaignStatus() string`

GetPlatformCampaignStatus returns the PlatformCampaignStatus field if non-nil, zero value otherwise.

### GetPlatformCampaignStatusOk

`func (o *AdCampaign) GetPlatformCampaignStatusOk() (*string, bool)`

GetPlatformCampaignStatusOk returns a tuple with the PlatformCampaignStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignStatus

`func (o *AdCampaign) SetPlatformCampaignStatus(v string)`

SetPlatformCampaignStatus sets PlatformCampaignStatus field to given value.

### HasPlatformCampaignStatus

`func (o *AdCampaign) HasPlatformCampaignStatus() bool`

HasPlatformCampaignStatus returns a boolean if a field has been set.

### SetPlatformCampaignStatusNil

`func (o *AdCampaign) SetPlatformCampaignStatusNil(b bool)`

 SetPlatformCampaignStatusNil sets the value for PlatformCampaignStatus to be an explicit nil

### UnsetPlatformCampaignStatus
`func (o *AdCampaign) UnsetPlatformCampaignStatus()`

UnsetPlatformCampaignStatus ensures that no value is present for PlatformCampaignStatus, not even an explicit nil
### GetCampaignIssuesInfo

`func (o *AdCampaign) GetCampaignIssuesInfo() []map[string]interface{}`

GetCampaignIssuesInfo returns the CampaignIssuesInfo field if non-nil, zero value otherwise.

### GetCampaignIssuesInfoOk

`func (o *AdCampaign) GetCampaignIssuesInfoOk() (*[]map[string]interface{}, bool)`

GetCampaignIssuesInfoOk returns a tuple with the CampaignIssuesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignIssuesInfo

`func (o *AdCampaign) SetCampaignIssuesInfo(v []map[string]interface{})`

SetCampaignIssuesInfo sets CampaignIssuesInfo field to given value.

### HasCampaignIssuesInfo

`func (o *AdCampaign) HasCampaignIssuesInfo() bool`

HasCampaignIssuesInfo returns a boolean if a field has been set.

### SetCampaignIssuesInfoNil

`func (o *AdCampaign) SetCampaignIssuesInfoNil(b bool)`

 SetCampaignIssuesInfoNil sets the value for CampaignIssuesInfo to be an explicit nil

### UnsetCampaignIssuesInfo
`func (o *AdCampaign) UnsetCampaignIssuesInfo()`

UnsetCampaignIssuesInfo ensures that no value is present for CampaignIssuesInfo, not even an explicit nil
### GetAdCount

`func (o *AdCampaign) GetAdCount() int32`

GetAdCount returns the AdCount field if non-nil, zero value otherwise.

### GetAdCountOk

`func (o *AdCampaign) GetAdCountOk() (*int32, bool)`

GetAdCountOk returns a tuple with the AdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdCount

`func (o *AdCampaign) SetAdCount(v int32)`

SetAdCount sets AdCount field to given value.

### HasAdCount

`func (o *AdCampaign) HasAdCount() bool`

HasAdCount returns a boolean if a field has been set.

### GetBudget

`func (o *AdCampaign) GetBudget() AdBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdCampaign) GetBudgetOk() (*AdBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdCampaign) SetBudget(v AdBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdCampaign) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetCampaignBudget

`func (o *AdCampaign) GetCampaignBudget() AdBudget`

GetCampaignBudget returns the CampaignBudget field if non-nil, zero value otherwise.

### GetCampaignBudgetOk

`func (o *AdCampaign) GetCampaignBudgetOk() (*AdBudget, bool)`

GetCampaignBudgetOk returns a tuple with the CampaignBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignBudget

`func (o *AdCampaign) SetCampaignBudget(v AdBudget)`

SetCampaignBudget sets CampaignBudget field to given value.

### HasCampaignBudget

`func (o *AdCampaign) HasCampaignBudget() bool`

HasCampaignBudget returns a boolean if a field has been set.

### GetBudgetLevel

`func (o *AdCampaign) GetBudgetLevel() string`

GetBudgetLevel returns the BudgetLevel field if non-nil, zero value otherwise.

### GetBudgetLevelOk

`func (o *AdCampaign) GetBudgetLevelOk() (*string, bool)`

GetBudgetLevelOk returns a tuple with the BudgetLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLevel

`func (o *AdCampaign) SetBudgetLevel(v string)`

SetBudgetLevel sets BudgetLevel field to given value.

### HasBudgetLevel

`func (o *AdCampaign) HasBudgetLevel() bool`

HasBudgetLevel returns a boolean if a field has been set.

### SetBudgetLevelNil

`func (o *AdCampaign) SetBudgetLevelNil(b bool)`

 SetBudgetLevelNil sets the value for BudgetLevel to be an explicit nil

### UnsetBudgetLevel
`func (o *AdCampaign) UnsetBudgetLevel()`

UnsetBudgetLevel ensures that no value is present for BudgetLevel, not even an explicit nil
### GetIsBudgetScheduleEnabled

`func (o *AdCampaign) GetIsBudgetScheduleEnabled() bool`

GetIsBudgetScheduleEnabled returns the IsBudgetScheduleEnabled field if non-nil, zero value otherwise.

### GetIsBudgetScheduleEnabledOk

`func (o *AdCampaign) GetIsBudgetScheduleEnabledOk() (*bool, bool)`

GetIsBudgetScheduleEnabledOk returns a tuple with the IsBudgetScheduleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBudgetScheduleEnabled

`func (o *AdCampaign) SetIsBudgetScheduleEnabled(v bool)`

SetIsBudgetScheduleEnabled sets IsBudgetScheduleEnabled field to given value.

### HasIsBudgetScheduleEnabled

`func (o *AdCampaign) HasIsBudgetScheduleEnabled() bool`

HasIsBudgetScheduleEnabled returns a boolean if a field has been set.

### GetCurrency

`func (o *AdCampaign) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdCampaign) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdCampaign) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AdCampaign) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *AdCampaign) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *AdCampaign) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetMetrics

`func (o *AdCampaign) GetMetrics() AdMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdCampaign) GetMetricsOk() (*AdMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdCampaign) SetMetrics(v AdMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdCampaign) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetPlatformAdAccountId

`func (o *AdCampaign) GetPlatformAdAccountId() string`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *AdCampaign) GetPlatformAdAccountIdOk() (*string, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *AdCampaign) SetPlatformAdAccountId(v string)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.

### HasPlatformAdAccountId

`func (o *AdCampaign) HasPlatformAdAccountId() bool`

HasPlatformAdAccountId returns a boolean if a field has been set.

### GetPlatformAdAccountName

`func (o *AdCampaign) GetPlatformAdAccountName() string`

GetPlatformAdAccountName returns the PlatformAdAccountName field if non-nil, zero value otherwise.

### GetPlatformAdAccountNameOk

`func (o *AdCampaign) GetPlatformAdAccountNameOk() (*string, bool)`

GetPlatformAdAccountNameOk returns a tuple with the PlatformAdAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountName

`func (o *AdCampaign) SetPlatformAdAccountName(v string)`

SetPlatformAdAccountName sets PlatformAdAccountName field to given value.

### HasPlatformAdAccountName

`func (o *AdCampaign) HasPlatformAdAccountName() bool`

HasPlatformAdAccountName returns a boolean if a field has been set.

### SetPlatformAdAccountNameNil

`func (o *AdCampaign) SetPlatformAdAccountNameNil(b bool)`

 SetPlatformAdAccountNameNil sets the value for PlatformAdAccountName to be an explicit nil

### UnsetPlatformAdAccountName
`func (o *AdCampaign) UnsetPlatformAdAccountName()`

UnsetPlatformAdAccountName ensures that no value is present for PlatformAdAccountName, not even an explicit nil
### GetAccountId

`func (o *AdCampaign) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AdCampaign) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AdCampaign) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AdCampaign) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProfileId

`func (o *AdCampaign) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AdCampaign) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AdCampaign) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *AdCampaign) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetAdvertisingChannelType

`func (o *AdCampaign) GetAdvertisingChannelType() string`

GetAdvertisingChannelType returns the AdvertisingChannelType field if non-nil, zero value otherwise.

### GetAdvertisingChannelTypeOk

`func (o *AdCampaign) GetAdvertisingChannelTypeOk() (*string, bool)`

GetAdvertisingChannelTypeOk returns a tuple with the AdvertisingChannelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertisingChannelType

`func (o *AdCampaign) SetAdvertisingChannelType(v string)`

SetAdvertisingChannelType sets AdvertisingChannelType field to given value.

### HasAdvertisingChannelType

`func (o *AdCampaign) HasAdvertisingChannelType() bool`

HasAdvertisingChannelType returns a boolean if a field has been set.

### SetAdvertisingChannelTypeNil

`func (o *AdCampaign) SetAdvertisingChannelTypeNil(b bool)`

 SetAdvertisingChannelTypeNil sets the value for AdvertisingChannelType to be an explicit nil

### UnsetAdvertisingChannelType
`func (o *AdCampaign) UnsetAdvertisingChannelType()`

UnsetAdvertisingChannelType ensures that no value is present for AdvertisingChannelType, not even an explicit nil
### GetPlatformObjective

`func (o *AdCampaign) GetPlatformObjective() string`

GetPlatformObjective returns the PlatformObjective field if non-nil, zero value otherwise.

### GetPlatformObjectiveOk

`func (o *AdCampaign) GetPlatformObjectiveOk() (*string, bool)`

GetPlatformObjectiveOk returns a tuple with the PlatformObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformObjective

`func (o *AdCampaign) SetPlatformObjective(v string)`

SetPlatformObjective sets PlatformObjective field to given value.

### HasPlatformObjective

`func (o *AdCampaign) HasPlatformObjective() bool`

HasPlatformObjective returns a boolean if a field has been set.

### SetPlatformObjectiveNil

`func (o *AdCampaign) SetPlatformObjectiveNil(b bool)`

 SetPlatformObjectiveNil sets the value for PlatformObjective to be an explicit nil

### UnsetPlatformObjective
`func (o *AdCampaign) UnsetPlatformObjective()`

UnsetPlatformObjective ensures that no value is present for PlatformObjective, not even an explicit nil
### GetOptimizationGoal

`func (o *AdCampaign) GetOptimizationGoal() string`

GetOptimizationGoal returns the OptimizationGoal field if non-nil, zero value otherwise.

### GetOptimizationGoalOk

`func (o *AdCampaign) GetOptimizationGoalOk() (*string, bool)`

GetOptimizationGoalOk returns a tuple with the OptimizationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationGoal

`func (o *AdCampaign) SetOptimizationGoal(v string)`

SetOptimizationGoal sets OptimizationGoal field to given value.

### HasOptimizationGoal

`func (o *AdCampaign) HasOptimizationGoal() bool`

HasOptimizationGoal returns a boolean if a field has been set.

### SetOptimizationGoalNil

`func (o *AdCampaign) SetOptimizationGoalNil(b bool)`

 SetOptimizationGoalNil sets the value for OptimizationGoal to be an explicit nil

### UnsetOptimizationGoal
`func (o *AdCampaign) UnsetOptimizationGoal()`

UnsetOptimizationGoal ensures that no value is present for OptimizationGoal, not even an explicit nil
### GetBidStrategy

`func (o *AdCampaign) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *AdCampaign) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *AdCampaign) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *AdCampaign) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### SetBidStrategyNil

`func (o *AdCampaign) SetBidStrategyNil(b bool)`

 SetBidStrategyNil sets the value for BidStrategy to be an explicit nil

### UnsetBidStrategy
`func (o *AdCampaign) UnsetBidStrategy()`

UnsetBidStrategy ensures that no value is present for BidStrategy, not even an explicit nil
### GetBidAmount

`func (o *AdCampaign) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *AdCampaign) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *AdCampaign) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *AdCampaign) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### SetBidAmountNil

`func (o *AdCampaign) SetBidAmountNil(b bool)`

 SetBidAmountNil sets the value for BidAmount to be an explicit nil

### UnsetBidAmount
`func (o *AdCampaign) UnsetBidAmount()`

UnsetBidAmount ensures that no value is present for BidAmount, not even an explicit nil
### GetRoasAverageFloor

`func (o *AdCampaign) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *AdCampaign) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *AdCampaign) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *AdCampaign) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### SetRoasAverageFloorNil

`func (o *AdCampaign) SetRoasAverageFloorNil(b bool)`

 SetRoasAverageFloorNil sets the value for RoasAverageFloor to be an explicit nil

### UnsetRoasAverageFloor
`func (o *AdCampaign) UnsetRoasAverageFloor()`

UnsetRoasAverageFloor ensures that no value is present for RoasAverageFloor, not even an explicit nil
### GetPromotedObject

`func (o *AdCampaign) GetPromotedObject() AdTreeAdSetPromotedObject`

GetPromotedObject returns the PromotedObject field if non-nil, zero value otherwise.

### GetPromotedObjectOk

`func (o *AdCampaign) GetPromotedObjectOk() (*AdTreeAdSetPromotedObject, bool)`

GetPromotedObjectOk returns a tuple with the PromotedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedObject

`func (o *AdCampaign) SetPromotedObject(v AdTreeAdSetPromotedObject)`

SetPromotedObject sets PromotedObject field to given value.

### HasPromotedObject

`func (o *AdCampaign) HasPromotedObject() bool`

HasPromotedObject returns a boolean if a field has been set.

### GetEarliestAd

`func (o *AdCampaign) GetEarliestAd() time.Time`

GetEarliestAd returns the EarliestAd field if non-nil, zero value otherwise.

### GetEarliestAdOk

`func (o *AdCampaign) GetEarliestAdOk() (*time.Time, bool)`

GetEarliestAdOk returns a tuple with the EarliestAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarliestAd

`func (o *AdCampaign) SetEarliestAd(v time.Time)`

SetEarliestAd sets EarliestAd field to given value.

### HasEarliestAd

`func (o *AdCampaign) HasEarliestAd() bool`

HasEarliestAd returns a boolean if a field has been set.

### GetLatestAd

`func (o *AdCampaign) GetLatestAd() time.Time`

GetLatestAd returns the LatestAd field if non-nil, zero value otherwise.

### GetLatestAdOk

`func (o *AdCampaign) GetLatestAdOk() (*time.Time, bool)`

GetLatestAdOk returns a tuple with the LatestAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestAd

`func (o *AdCampaign) SetLatestAd(v time.Time)`

SetLatestAd sets LatestAd field to given value.

### HasLatestAd

`func (o *AdCampaign) HasLatestAd() bool`

HasLatestAd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


