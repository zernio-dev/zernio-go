# AdTreeAdSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformAdSetId** | Pointer to **string** |  | [optional] 
**AdSetName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**AdStatus**](AdStatus.md) | Derived from child ad statuses | [optional] 
**AdCount** | Pointer to **int32** |  | [optional] 
**Budget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**AdSetBudget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**Metrics** | Pointer to [**AdMetrics**](AdMetrics.md) |  | [optional] 
**OptimizationGoal** | Pointer to **NullableString** | Meta ad set optimization goal (e.g. OFFSITE_CONVERSIONS, VALUE, LEAD_GENERATION) | [optional] 
**BidStrategy** | Pointer to [**NullableBidStrategy**](BidStrategy.md) | Bid strategy for this ad set (overrides campaign level when set) | [optional] 
**BidAmount** | Pointer to **NullableFloat32** | Bid cap in whole currency units. Populated when bidStrategy is LOWEST_COST_WITH_BID_CAP or COST_CAP. | [optional] 
**RoasAverageFloor** | Pointer to **NullableFloat32** | Minimum ROAS as a decimal multiplier (2.0 &#x3D; 2.0x). Populated when bidStrategy is LOWEST_COST_WITH_MIN_ROAS. | [optional] 
**PromotedObject** | Pointer to [**AdTreeAdSetPromotedObject**](AdTreeAdSetPromotedObject.md) |  | [optional] 
**Ads** | Pointer to [**[]Ad**](Ad.md) | Individual ads within this ad set (capped at 100). Returns a subset of Ad fields from the aggregation (core fields like _id, name, platform, status, budget, metrics, creative, goal are included; targeting and schedule may be absent). | [optional] 

## Methods

### NewAdTreeAdSet

`func NewAdTreeAdSet() *AdTreeAdSet`

NewAdTreeAdSet instantiates a new AdTreeAdSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdTreeAdSetWithDefaults

`func NewAdTreeAdSetWithDefaults() *AdTreeAdSet`

NewAdTreeAdSetWithDefaults instantiates a new AdTreeAdSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformAdSetId

`func (o *AdTreeAdSet) GetPlatformAdSetId() string`

GetPlatformAdSetId returns the PlatformAdSetId field if non-nil, zero value otherwise.

### GetPlatformAdSetIdOk

`func (o *AdTreeAdSet) GetPlatformAdSetIdOk() (*string, bool)`

GetPlatformAdSetIdOk returns a tuple with the PlatformAdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdSetId

`func (o *AdTreeAdSet) SetPlatformAdSetId(v string)`

SetPlatformAdSetId sets PlatformAdSetId field to given value.

### HasPlatformAdSetId

`func (o *AdTreeAdSet) HasPlatformAdSetId() bool`

HasPlatformAdSetId returns a boolean if a field has been set.

### GetAdSetName

`func (o *AdTreeAdSet) GetAdSetName() string`

GetAdSetName returns the AdSetName field if non-nil, zero value otherwise.

### GetAdSetNameOk

`func (o *AdTreeAdSet) GetAdSetNameOk() (*string, bool)`

GetAdSetNameOk returns a tuple with the AdSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetName

`func (o *AdTreeAdSet) SetAdSetName(v string)`

SetAdSetName sets AdSetName field to given value.

### HasAdSetName

`func (o *AdTreeAdSet) HasAdSetName() bool`

HasAdSetName returns a boolean if a field has been set.

### GetStatus

`func (o *AdTreeAdSet) GetStatus() AdStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdTreeAdSet) GetStatusOk() (*AdStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdTreeAdSet) SetStatus(v AdStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdTreeAdSet) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAdCount

`func (o *AdTreeAdSet) GetAdCount() int32`

GetAdCount returns the AdCount field if non-nil, zero value otherwise.

### GetAdCountOk

`func (o *AdTreeAdSet) GetAdCountOk() (*int32, bool)`

GetAdCountOk returns a tuple with the AdCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdCount

`func (o *AdTreeAdSet) SetAdCount(v int32)`

SetAdCount sets AdCount field to given value.

### HasAdCount

`func (o *AdTreeAdSet) HasAdCount() bool`

HasAdCount returns a boolean if a field has been set.

### GetBudget

`func (o *AdTreeAdSet) GetBudget() AdBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *AdTreeAdSet) GetBudgetOk() (*AdBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *AdTreeAdSet) SetBudget(v AdBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *AdTreeAdSet) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetAdSetBudget

`func (o *AdTreeAdSet) GetAdSetBudget() AdBudget`

GetAdSetBudget returns the AdSetBudget field if non-nil, zero value otherwise.

### GetAdSetBudgetOk

`func (o *AdTreeAdSet) GetAdSetBudgetOk() (*AdBudget, bool)`

GetAdSetBudgetOk returns a tuple with the AdSetBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSetBudget

`func (o *AdTreeAdSet) SetAdSetBudget(v AdBudget)`

SetAdSetBudget sets AdSetBudget field to given value.

### HasAdSetBudget

`func (o *AdTreeAdSet) HasAdSetBudget() bool`

HasAdSetBudget returns a boolean if a field has been set.

### GetMetrics

`func (o *AdTreeAdSet) GetMetrics() AdMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *AdTreeAdSet) GetMetricsOk() (*AdMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *AdTreeAdSet) SetMetrics(v AdMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *AdTreeAdSet) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetOptimizationGoal

`func (o *AdTreeAdSet) GetOptimizationGoal() string`

GetOptimizationGoal returns the OptimizationGoal field if non-nil, zero value otherwise.

### GetOptimizationGoalOk

`func (o *AdTreeAdSet) GetOptimizationGoalOk() (*string, bool)`

GetOptimizationGoalOk returns a tuple with the OptimizationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationGoal

`func (o *AdTreeAdSet) SetOptimizationGoal(v string)`

SetOptimizationGoal sets OptimizationGoal field to given value.

### HasOptimizationGoal

`func (o *AdTreeAdSet) HasOptimizationGoal() bool`

HasOptimizationGoal returns a boolean if a field has been set.

### SetOptimizationGoalNil

`func (o *AdTreeAdSet) SetOptimizationGoalNil(b bool)`

 SetOptimizationGoalNil sets the value for OptimizationGoal to be an explicit nil

### UnsetOptimizationGoal
`func (o *AdTreeAdSet) UnsetOptimizationGoal()`

UnsetOptimizationGoal ensures that no value is present for OptimizationGoal, not even an explicit nil
### GetBidStrategy

`func (o *AdTreeAdSet) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *AdTreeAdSet) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *AdTreeAdSet) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *AdTreeAdSet) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### SetBidStrategyNil

`func (o *AdTreeAdSet) SetBidStrategyNil(b bool)`

 SetBidStrategyNil sets the value for BidStrategy to be an explicit nil

### UnsetBidStrategy
`func (o *AdTreeAdSet) UnsetBidStrategy()`

UnsetBidStrategy ensures that no value is present for BidStrategy, not even an explicit nil
### GetBidAmount

`func (o *AdTreeAdSet) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *AdTreeAdSet) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *AdTreeAdSet) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *AdTreeAdSet) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### SetBidAmountNil

`func (o *AdTreeAdSet) SetBidAmountNil(b bool)`

 SetBidAmountNil sets the value for BidAmount to be an explicit nil

### UnsetBidAmount
`func (o *AdTreeAdSet) UnsetBidAmount()`

UnsetBidAmount ensures that no value is present for BidAmount, not even an explicit nil
### GetRoasAverageFloor

`func (o *AdTreeAdSet) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *AdTreeAdSet) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *AdTreeAdSet) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *AdTreeAdSet) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### SetRoasAverageFloorNil

`func (o *AdTreeAdSet) SetRoasAverageFloorNil(b bool)`

 SetRoasAverageFloorNil sets the value for RoasAverageFloor to be an explicit nil

### UnsetRoasAverageFloor
`func (o *AdTreeAdSet) UnsetRoasAverageFloor()`

UnsetRoasAverageFloor ensures that no value is present for RoasAverageFloor, not even an explicit nil
### GetPromotedObject

`func (o *AdTreeAdSet) GetPromotedObject() AdTreeAdSetPromotedObject`

GetPromotedObject returns the PromotedObject field if non-nil, zero value otherwise.

### GetPromotedObjectOk

`func (o *AdTreeAdSet) GetPromotedObjectOk() (*AdTreeAdSetPromotedObject, bool)`

GetPromotedObjectOk returns a tuple with the PromotedObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotedObject

`func (o *AdTreeAdSet) SetPromotedObject(v AdTreeAdSetPromotedObject)`

SetPromotedObject sets PromotedObject field to given value.

### HasPromotedObject

`func (o *AdTreeAdSet) HasPromotedObject() bool`

HasPromotedObject returns a boolean if a field has been set.

### GetAds

`func (o *AdTreeAdSet) GetAds() []Ad`

GetAds returns the Ads field if non-nil, zero value otherwise.

### GetAdsOk

`func (o *AdTreeAdSet) GetAdsOk() (*[]Ad, bool)`

GetAdsOk returns a tuple with the Ads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAds

`func (o *AdTreeAdSet) SetAds(v []Ad)`

SetAds sets Ads field to given value.

### HasAds

`func (o *AdTreeAdSet) HasAds() bool`

HasAds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


