# UpdateAdSet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Budget** | Pointer to [**AdBudget**](AdBudget.md) |  | [optional] 
**BudgetLevel** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusUpdated** | Pointer to **int32** |  | [optional] 
**StatusSkipped** | Pointer to **int32** |  | [optional] 
**BidStrategy** | Pointer to [**BidStrategy**](BidStrategy.md) |  | [optional] 
**BidAmount** | Pointer to **NullableFloat32** |  | [optional] 
**RoasAverageFloor** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewUpdateAdSet200Response

`func NewUpdateAdSet200Response() *UpdateAdSet200Response`

NewUpdateAdSet200Response instantiates a new UpdateAdSet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdSet200ResponseWithDefaults

`func NewUpdateAdSet200ResponseWithDefaults() *UpdateAdSet200Response`

NewUpdateAdSet200ResponseWithDefaults instantiates a new UpdateAdSet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudget

`func (o *UpdateAdSet200Response) GetBudget() AdBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdateAdSet200Response) GetBudgetOk() (*AdBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdateAdSet200Response) SetBudget(v AdBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdateAdSet200Response) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetBudgetLevel

`func (o *UpdateAdSet200Response) GetBudgetLevel() string`

GetBudgetLevel returns the BudgetLevel field if non-nil, zero value otherwise.

### GetBudgetLevelOk

`func (o *UpdateAdSet200Response) GetBudgetLevelOk() (*string, bool)`

GetBudgetLevelOk returns a tuple with the BudgetLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetLevel

`func (o *UpdateAdSet200Response) SetBudgetLevel(v string)`

SetBudgetLevel sets BudgetLevel field to given value.

### HasBudgetLevel

`func (o *UpdateAdSet200Response) HasBudgetLevel() bool`

HasBudgetLevel returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateAdSet200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAdSet200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAdSet200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAdSet200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusUpdated

`func (o *UpdateAdSet200Response) GetStatusUpdated() int32`

GetStatusUpdated returns the StatusUpdated field if non-nil, zero value otherwise.

### GetStatusUpdatedOk

`func (o *UpdateAdSet200Response) GetStatusUpdatedOk() (*int32, bool)`

GetStatusUpdatedOk returns a tuple with the StatusUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusUpdated

`func (o *UpdateAdSet200Response) SetStatusUpdated(v int32)`

SetStatusUpdated sets StatusUpdated field to given value.

### HasStatusUpdated

`func (o *UpdateAdSet200Response) HasStatusUpdated() bool`

HasStatusUpdated returns a boolean if a field has been set.

### GetStatusSkipped

`func (o *UpdateAdSet200Response) GetStatusSkipped() int32`

GetStatusSkipped returns the StatusSkipped field if non-nil, zero value otherwise.

### GetStatusSkippedOk

`func (o *UpdateAdSet200Response) GetStatusSkippedOk() (*int32, bool)`

GetStatusSkippedOk returns a tuple with the StatusSkipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusSkipped

`func (o *UpdateAdSet200Response) SetStatusSkipped(v int32)`

SetStatusSkipped sets StatusSkipped field to given value.

### HasStatusSkipped

`func (o *UpdateAdSet200Response) HasStatusSkipped() bool`

HasStatusSkipped returns a boolean if a field has been set.

### GetBidStrategy

`func (o *UpdateAdSet200Response) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *UpdateAdSet200Response) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *UpdateAdSet200Response) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *UpdateAdSet200Response) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### GetBidAmount

`func (o *UpdateAdSet200Response) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *UpdateAdSet200Response) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *UpdateAdSet200Response) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *UpdateAdSet200Response) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### SetBidAmountNil

`func (o *UpdateAdSet200Response) SetBidAmountNil(b bool)`

 SetBidAmountNil sets the value for BidAmount to be an explicit nil

### UnsetBidAmount
`func (o *UpdateAdSet200Response) UnsetBidAmount()`

UnsetBidAmount ensures that no value is present for BidAmount, not even an explicit nil
### GetRoasAverageFloor

`func (o *UpdateAdSet200Response) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *UpdateAdSet200Response) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *UpdateAdSet200Response) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *UpdateAdSet200Response) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### SetRoasAverageFloorNil

`func (o *UpdateAdSet200Response) SetRoasAverageFloorNil(b bool)`

 SetRoasAverageFloorNil sets the value for RoasAverageFloor to be an explicit nil

### UnsetRoasAverageFloor
`func (o *UpdateAdSet200Response) UnsetRoasAverageFloor()`

UnsetRoasAverageFloor ensures that no value is present for RoasAverageFloor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


