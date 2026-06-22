# UpdateAdSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**Budget** | Pointer to [**UpdateAdSetRequestBudget**](UpdateAdSetRequestBudget.md) |  | [optional] 
**Status** | Pointer to **string** | Omit if not toggling delivery state | [optional] 
**Name** | Pointer to **string** | Rename the ad set (Meta only; other platforms return 501). At least one of budget/status/bidStrategy/name is required. | [optional] 
**BidStrategy** | Pointer to [**BidStrategy**](BidStrategy.md) | Ad-set-level bid strategy. Overrides the campaign-level default. Supported on Meta (facebook, instagram) and TikTok. On TikTok the Meta-style enum is mapped to bid_type / bid_price / deep_bid_type automatically. Other platforms (linkedin, pinterest, google, twitter) return 501 Not Implemented when bidStrategy is set.  | [optional] 
**BidAmount** | Pointer to **float32** | Bid cap in WHOLE currency units (USD: 5 &#x3D; $5.00; JPY: 100 &#x3D; ¥100). Required when bidStrategy is LOWEST_COST_WITH_BID_CAP or COST_CAP. Internally converted to Meta&#39;s smallest-denomination integer.  | [optional] 
**RoasAverageFloor** | Pointer to **float32** | Minimum ROAS as a decimal multiplier (2.0 &#x3D; 2.0x). Required when bidStrategy is LOWEST_COST_WITH_MIN_ROAS. Sent to Meta as &#x60;bid_constraints.roas_average_floor&#x60; × 10000.  | [optional] 

## Methods

### NewUpdateAdSetRequest

`func NewUpdateAdSetRequest(platform string, ) *UpdateAdSetRequest`

NewUpdateAdSetRequest instantiates a new UpdateAdSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdSetRequestWithDefaults

`func NewUpdateAdSetRequestWithDefaults() *UpdateAdSetRequest`

NewUpdateAdSetRequestWithDefaults instantiates a new UpdateAdSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *UpdateAdSetRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateAdSetRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateAdSetRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetBudget

`func (o *UpdateAdSetRequest) GetBudget() UpdateAdSetRequestBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdateAdSetRequest) GetBudgetOk() (*UpdateAdSetRequestBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdateAdSetRequest) SetBudget(v UpdateAdSetRequestBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdateAdSetRequest) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAdSetRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAdSetRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAdSetRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAdSetRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetName

`func (o *UpdateAdSetRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAdSetRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAdSetRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAdSetRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBidStrategy

`func (o *UpdateAdSetRequest) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *UpdateAdSetRequest) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *UpdateAdSetRequest) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *UpdateAdSetRequest) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### GetBidAmount

`func (o *UpdateAdSetRequest) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *UpdateAdSetRequest) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *UpdateAdSetRequest) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *UpdateAdSetRequest) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### GetRoasAverageFloor

`func (o *UpdateAdSetRequest) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *UpdateAdSetRequest) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *UpdateAdSetRequest) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *UpdateAdSetRequest) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


