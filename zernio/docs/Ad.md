# Ad

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AdStatus**](AdStatus.md) |  | [optional] 
**AdType** | Pointer to **string** |  | [optional] 
**Goal** | Pointer to **string** | Available goals vary by platform. Meta (Facebook/Instagram) supports all 9 (incl. &#x60;lead_conversion&#x60; &#x3D; website pixel lead optimization and &#x60;catalog_sales&#x60; &#x3D; Advantage+ catalog ads). TikTok supports the 7 non-&#x60;lead_conversion&#x60; goals. LinkedIn supports all except app_promotion / lead_conversion. Twitter/X supports engagement, traffic, awareness, video_views, app_promotion. Pinterest and Google Ads support only engagement, traffic, awareness, video_views. | [optional] 
**IsExternal** | Pointer to **bool** | True for ads synced from platform ad managers | [optional] 
**Budget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**Metrics** | Pointer to [**NullableAdMetrics**](AdMetrics.md) |  | [optional] 
**PlatformAdId** | Pointer to **string** |  | [optional] 
**PlatformAdAccountId** | Pointer to **string** |  | [optional] 
**PlatformCampaignId** | Pointer to **string** |  | [optional] 
**PlatformAdSetId** | Pointer to **string** |  | [optional] 
**CampaignName** | Pointer to **string** |  | [optional] 
**AdSetName** | Pointer to **string** |  | [optional] 
**PlatformObjective** | Pointer to **NullableString** | Raw Meta campaign objective (e.g. OUTCOME_SALES, OUTCOME_LEADS, OUTCOME_TRAFFIC). Only present for Meta ads. | [optional] 
**OptimizationGoal** | Pointer to **NullableString** | Meta ad set optimization goal (e.g. OFFSITE_CONVERSIONS, VALUE, LEAD_GENERATION, LINK_CLICKS). Only present for Meta ads. | [optional] 
**PlatformAdAccountName** | Pointer to **NullableString** | Human-readable advertiser/account name (Meta &#x60;AdAccount.name&#x60;, TikTok &#x60;advertiser_name&#x60;, LinkedIn / X / Pinterest equivalents). Refreshed every sync so platform-side renames propagate within one cycle. &#x60;null&#x60; when the platform doesn&#39;t return a name or the sync hasn&#39;t run yet.  | [optional] 
**PlatformCreatedAt** | Pointer to **NullableTime** | Platform-reported creation timestamp (Meta &#x60;created_time&#x60;, TikTok &#x60;create_time&#x60;). Distinct from &#x60;createdAt&#x60; which reflects when Zernio first synced the doc — for sort/filter by \&quot;when the ad was actually created on the platform\&quot;, read this field. &#x60;null&#x60; for legacy ads synced before this field was added; aggregations fall back to &#x60;createdAt&#x60; in that case.  | [optional] 
**BidStrategy** | Pointer to [**NullableBidStrategy**](BidStrategy.md) | Ad-set bid strategy (overrides campaign level on Meta). Populated for Meta and TikTok. TikTok&#39;s native &#x60;bid_type&#x60; is normalized to the cross-platform Meta enum: &#x60;BID_TYPE_NO_BID&#x60; -&gt; &#x60;LOWEST_COST_WITHOUT_CAP&#x60;, &#x60;BID_TYPE_CUSTOM&#x60; -&gt; &#x60;LOWEST_COST_WITH_BID_CAP&#x60;, deep_bid_type&#x3D;MIN_ROAS or roas_bid&gt;0 -&gt; &#x60;LOWEST_COST_WITH_MIN_ROAS&#x60;, &#x60;BID_TYPE_MAX_CONVERSION&#x60; -&gt; &#x60;LOWEST_COST_WITHOUT_CAP&#x60;.  | [optional] 
**BidAmount** | Pointer to **NullableFloat32** | Bid cap in WHOLE currency units of the ad account (USD: 5 &#x3D; $5.00; JPY: 100 &#x3D; ¥100). Populated when bidStrategy is &#x60;LOWEST_COST_WITH_BID_CAP&#x60; or &#x60;COST_CAP&#x60;. &#x60;null&#x60; for auto-bid (&#x60;LOWEST_COST_WITHOUT_CAP&#x60;).  - Meta source: &#x60;bid_amount&#x60; on the ad set (smallest-denomination int, decoded here). - TikTok source: priority order &#x60;bid_price&#x60; -&gt; &#x60;conversion_bid_price&#x60; -&gt; &#x60;deep_cpa_bid&#x60;   (whichever is set on the ad group). TikTok stores all three in whole currency units.  Source: facebook-business-sdk-codegen api_specs/specs/AdSet.json (&#x60;bid_amount&#x60;).  | [optional] 
**RoasAverageFloor** | Pointer to **NullableFloat32** | Minimum ROAS as a decimal multiplier (2.0 &#x3D; 2.0x ROAS). Populated when bidStrategy is &#x60;LOWEST_COST_WITH_MIN_ROAS&#x60;.  - Meta source: decoded from &#x60;bid_constraints.roas_average_floor&#x60; (Meta stores as   fixed-point int × 10000; we return the decimal). - TikTok source: &#x60;roas_bid&#x60; on the ad group (already a decimal).  Source: facebook-business-sdk-codegen api_specs/specs/AdCampaignBidConstraint.json.  | [optional] 
**PromotedObject** | Pointer to [**AdPromotedObject**](AdPromotedObject.md) |  | [optional] 
**Creative** | Pointer to [**AdCreative**](AdCreative.md) |  | [optional] 
**Targeting** | Pointer to **map[string]interface{}** | The ad set&#39;s targeting (age, gender, geo, interests, placements, audience inclusions/exclusions). For ads created through Zernio this is the spec you supplied. For external ads (synced from Meta Ads Manager, &#x60;isExternal: true&#x60;) targeting lives at the ad set and isn&#39;t stored at ingest, so on the first &#x60;GET /v1/ads/{adId}&#x60; Zernio resolves it live from Meta and caches it on the ad; the value is then Meta&#39;s raw &#x60;targeting&#x60; shape (snake_case, e.g. &#x60;geo_locations&#x60;, &#x60;age_min&#x60;), the same object Ads Manager shows. May be absent if the ad set exposes no targeting or the lookup fails.  | [optional] 
**Schedule** | Pointer to [**AdSchedule**](AdSchedule.md) |  | [optional] 
**RejectionReason** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAd

`func NewAd() *Ad`

NewAd instantiates a new Ad object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdWithDefaults

`func NewAdWithDefaults() *Ad`

NewAdWithDefaults instantiates a new Ad object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Ad) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Ad) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Ad) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Ad) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Ad) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Ad) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Ad) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Ad) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *Ad) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Ad) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Ad) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Ad) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *Ad) GetStatus() AdStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Ad) GetStatusOk() (*AdStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Ad) SetStatus(v AdStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Ad) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdType

`func (o *Ad) GetAdType() string`

GetAdType returns the AdType field if non-nil, zero value otherwise.

### GetAdTypeOk

`func (o *Ad) GetAdTypeOk() (*string, bool)`

GetAdTypeOk returns a tuple with the AdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdType

`func (o *Ad) SetAdType(v string)`

SetAdType sets AdType field to given value.

### HasAdType

`func (o *Ad) HasAdType() bool`

HasAdType returns a boolean if a field has been set.

### GetGoal

`func (o *Ad) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *Ad) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *Ad) SetGoal(v string)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *Ad) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetIsExternal

`func (o *Ad) GetIsExternal() bool`

GetIsExternal returns the IsExternal field if non-nil, zero value otherwise.

### GetIsExternalOk

`func (o *Ad) GetIsExternalOk() (*bool, bool)`

GetIsExternalOk returns a tuple with the IsExternal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExternal

`func (o *Ad) SetIsExternal(v bool)`

SetIsExternal sets IsExternal field to given value.

### HasIsExternal

`func (o *Ad) HasIsExternal() bool`

HasIsExternal returns a boolean if a field has been set.

### GetBudget

`func (o *Ad) GetBudget() AdBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *Ad) GetBudgetOk() (*AdBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *Ad) SetBudget(v AdBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *Ad) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetMetrics

`func (o *Ad) GetMetrics() AdMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Ad) GetMetricsOk() (*AdMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Ad) SetMetrics(v AdMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Ad) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### SetMetricsNil

`func (o *Ad) SetMetricsNil(b bool)`

 SetMetricsNil sets the value for Metrics to be an explicit nil

### UnsetMetrics
`func (o *Ad) UnsetMetrics()`

UnsetMetrics ensures that no value is present for Metrics, not even an explicit nil
### GetPlatformAdId

`func (o *Ad) GetPlatformAdId() string`

GetPlatformAdId returns the PlatformAdId field if non-nil, zero value otherwise.

### GetPlatformAdIdOk

`func (o *Ad) GetPlatformAdIdOk() (*string, bool)`

GetPlatformAdIdOk returns a tuple with the PlatformAdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdId

`func (o *Ad) SetPlatformAdId(v string)`

SetPlatformAdId sets PlatformAdId field to given value.

### HasPlatformAdId

`func (o *Ad) HasPlatformAdId() bool`

HasPlatformAdId returns a boolean if a field has been set.

### GetPlatformAdAccountId

`func (o *Ad) GetPlatformAdAccountId() string`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *Ad) GetPlatformAdAccountIdOk() (*string, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *Ad) SetPlatformAdAccountId(v string)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.

### HasPlatformAdAccountId

`func (o *Ad) HasPlatformAdAccountId() bool`

HasPlatformAdAccountId returns a boolean if a field has been set.

### GetPlatformCampaignId

`func (o *Ad) GetPlatformCampaignId() string`

GetPlatformCampaignId returns the PlatformCampaignId field if non-nil, zero value otherwise.

### GetPlatformCampaignIdOk

`func (o *Ad) GetPlatformCampaignIdOk() (*string, bool)`

GetPlatformCampaignIdOk returns a tuple with the PlatformCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignId

`func (o *Ad) SetPlatformCampaignId(v string)`

SetPlatformCampaignId sets PlatformCampaignId field to given value.

### HasPlatformCampaignId

`func (o *Ad) HasPlatformCampaignId() bool`

HasPlatformCampaignId returns a boolean if a field has been set.

### GetPlatformAdSetId

`func (o *Ad) GetPlatformAdSetId() string`

GetPlatformAdSetId returns the PlatformAdSetId field if non-nil, zero value otherwise.

### GetPlatformAdSetIdOk

`func (o *Ad) GetPlatformAdSetIdOk() (*string, bool)`

GetPlatformAdSetIdOk returns a tuple with the PlatformAdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdSetId

`func (o *Ad) SetPlatformAdSetId(v string)`

SetPlatformAdSetId sets PlatformAdSetId field to given value.

### HasPlatformAdSetId

`func (o *Ad) HasPlatformAdSetId() bool`

HasPlatformAdSetId returns a boolean if a field has been set.

### GetCampaignName

`func (o *Ad) GetCampaignName() string`

GetCampaignName returns the CampaignName field if non-nil, zero value otherwise.

### GetCampaignNameOk

`func (o *Ad) GetCampaignNameOk() (*string, bool)`

GetCampaignNameOk returns a tuple with the CampaignName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignName

`func (o *Ad) SetCampaignName(v string)`

SetCampaignName sets CampaignName field to given value.

### HasCampaignName

`func (o *Ad) HasCampaignName() bool`

HasCampaignName returns a boolean if a field has been set.

### GetAdSetName

`func (o *Ad) GetAdSetName() string`

GetAdSetName returns the AdSetName field if non-nil, zero value otherwise.

### GetAdSetNameOk

`func (o *Ad) GetAdSetNameOk() (*string, bool)`

GetAdSetNameOk returns a tuple with the AdSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetName

`func (o *Ad) SetAdSetName(v string)`

SetAdSetName sets AdSetName field to given value.

### HasAdSetName

`func (o *Ad) HasAdSetName() bool`

HasAdSetName returns a boolean if a field has been set.

### GetPlatformObjective

`func (o *Ad) GetPlatformObjective() string`

GetPlatformObjective returns the PlatformObjective field if non-nil, zero value otherwise.

### GetPlatformObjectiveOk

`func (o *Ad) GetPlatformObjectiveOk() (*string, bool)`

GetPlatformObjectiveOk returns a tuple with the PlatformObjective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformObjective

`func (o *Ad) SetPlatformObjective(v string)`

SetPlatformObjective sets PlatformObjective field to given value.

### HasPlatformObjective

`func (o *Ad) HasPlatformObjective() bool`

HasPlatformObjective returns a boolean if a field has been set.

### SetPlatformObjectiveNil

`func (o *Ad) SetPlatformObjectiveNil(b bool)`

 SetPlatformObjectiveNil sets the value for PlatformObjective to be an explicit nil

### UnsetPlatformObjective
`func (o *Ad) UnsetPlatformObjective()`

UnsetPlatformObjective ensures that no value is present for PlatformObjective, not even an explicit nil
### GetOptimizationGoal

`func (o *Ad) GetOptimizationGoal() string`

GetOptimizationGoal returns the OptimizationGoal field if non-nil, zero value otherwise.

### GetOptimizationGoalOk

`func (o *Ad) GetOptimizationGoalOk() (*string, bool)`

GetOptimizationGoalOk returns a tuple with the OptimizationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationGoal

`func (o *Ad) SetOptimizationGoal(v string)`

SetOptimizationGoal sets OptimizationGoal field to given value.

### HasOptimizationGoal

`func (o *Ad) HasOptimizationGoal() bool`

HasOptimizationGoal returns a boolean if a field has been set.

### SetOptimizationGoalNil

`func (o *Ad) SetOptimizationGoalNil(b bool)`

 SetOptimizationGoalNil sets the value for OptimizationGoal to be an explicit nil

### UnsetOptimizationGoal
`func (o *Ad) UnsetOptimizationGoal()`

UnsetOptimizationGoal ensures that no value is present for OptimizationGoal, not even an explicit nil
### GetPlatformAdAccountName

`func (o *Ad) GetPlatformAdAccountName() string`

GetPlatformAdAccountName returns the PlatformAdAccountName field if non-nil, zero value otherwise.

### GetPlatformAdAccountNameOk

`func (o *Ad) GetPlatformAdAccountNameOk() (*string, bool)`

GetPlatformAdAccountNameOk returns a tuple with the PlatformAdAccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountName

`func (o *Ad) SetPlatformAdAccountName(v string)`

SetPlatformAdAccountName sets PlatformAdAccountName field to given value.

### HasPlatformAdAccountName

`func (o *Ad) HasPlatformAdAccountName() bool`

HasPlatformAdAccountName returns a boolean if a field has been set.

### SetPlatformAdAccountNameNil

`func (o *Ad) SetPlatformAdAccountNameNil(b bool)`

 SetPlatformAdAccountNameNil sets the value for PlatformAdAccountName to be an explicit nil

### UnsetPlatformAdAccountName
`func (o *Ad) UnsetPlatformAdAccountName()`

UnsetPlatformAdAccountName ensures that no value is present for PlatformAdAccountName, not even an explicit nil
### GetPlatformCreatedAt

`func (o *Ad) GetPlatformCreatedAt() time.Time`

GetPlatformCreatedAt returns the PlatformCreatedAt field if non-nil, zero value otherwise.

### GetPlatformCreatedAtOk

`func (o *Ad) GetPlatformCreatedAtOk() (*time.Time, bool)`

GetPlatformCreatedAtOk returns a tuple with the PlatformCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCreatedAt

`func (o *Ad) SetPlatformCreatedAt(v time.Time)`

SetPlatformCreatedAt sets PlatformCreatedAt field to given value.

### HasPlatformCreatedAt

`func (o *Ad) HasPlatformCreatedAt() bool`

HasPlatformCreatedAt returns a boolean if a field has been set.

### SetPlatformCreatedAtNil

`func (o *Ad) SetPlatformCreatedAtNil(b bool)`

 SetPlatformCreatedAtNil sets the value for PlatformCreatedAt to be an explicit nil

### UnsetPlatformCreatedAt
`func (o *Ad) UnsetPlatformCreatedAt()`

UnsetPlatformCreatedAt ensures that no value is present for PlatformCreatedAt, not even an explicit nil
### GetBidStrategy

`func (o *Ad) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *Ad) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *Ad) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *Ad) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### SetBidStrategyNil

`func (o *Ad) SetBidStrategyNil(b bool)`

 SetBidStrategyNil sets the value for BidStrategy to be an explicit nil

### UnsetBidStrategy
`func (o *Ad) UnsetBidStrategy()`

UnsetBidStrategy ensures that no value is present for BidStrategy, not even an explicit nil
### GetBidAmount

`func (o *Ad) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *Ad) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *Ad) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *Ad) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### SetBidAmountNil

`func (o *Ad) SetBidAmountNil(b bool)`

 SetBidAmountNil sets the value for BidAmount to be an explicit nil

### UnsetBidAmount
`func (o *Ad) UnsetBidAmount()`

UnsetBidAmount ensures that no value is present for BidAmount, not even an explicit nil
### GetRoasAverageFloor

`func (o *Ad) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *Ad) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *Ad) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *Ad) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### SetRoasAverageFloorNil

`func (o *Ad) SetRoasAverageFloorNil(b bool)`

 SetRoasAverageFloorNil sets the value for RoasAverageFloor to be an explicit nil

### UnsetRoasAverageFloor
`func (o *Ad) UnsetRoasAverageFloor()`

UnsetRoasAverageFloor ensures that no value is present for RoasAverageFloor, not even an explicit nil
### GetPromotedObject

`func (o *Ad) GetPromotedObject() AdPromotedObject`

GetPromotedObject returns the PromotedObject field if non-nil, zero value otherwise.

### GetPromotedObjectOk

`func (o *Ad) GetPromotedObjectOk() (*AdPromotedObject, bool)`

GetPromotedObjectOk returns a tuple with the PromotedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedObject

`func (o *Ad) SetPromotedObject(v AdPromotedObject)`

SetPromotedObject sets PromotedObject field to given value.

### HasPromotedObject

`func (o *Ad) HasPromotedObject() bool`

HasPromotedObject returns a boolean if a field has been set.

### GetCreative

`func (o *Ad) GetCreative() AdCreative`

GetCreative returns the Creative field if non-nil, zero value otherwise.

### GetCreativeOk

`func (o *Ad) GetCreativeOk() (*AdCreative, bool)`

GetCreativeOk returns a tuple with the Creative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreative

`func (o *Ad) SetCreative(v AdCreative)`

SetCreative sets Creative field to given value.

### HasCreative

`func (o *Ad) HasCreative() bool`

HasCreative returns a boolean if a field has been set.

### GetTargeting

`func (o *Ad) GetTargeting() map[string]interface{}`

GetTargeting returns the Targeting field if non-nil, zero value otherwise.

### GetTargetingOk

`func (o *Ad) GetTargetingOk() (*map[string]interface{}, bool)`

GetTargetingOk returns a tuple with the Targeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargeting

`func (o *Ad) SetTargeting(v map[string]interface{})`

SetTargeting sets Targeting field to given value.

### HasTargeting

`func (o *Ad) HasTargeting() bool`

HasTargeting returns a boolean if a field has been set.

### GetSchedule

`func (o *Ad) GetSchedule() AdSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *Ad) GetScheduleOk() (*AdSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *Ad) SetSchedule(v AdSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *Ad) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetRejectionReason

`func (o *Ad) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *Ad) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *Ad) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *Ad) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Ad) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Ad) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Ad) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Ad) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Ad) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Ad) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Ad) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Ad) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


