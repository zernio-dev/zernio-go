# UpdateQueueSlot200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**QueueSchedule**](QueueSchedule.md) |  | [optional] 
**NextSlots** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**ReshuffledCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewUpdateQueueSlot200Response

`func NewUpdateQueueSlot200Response() *UpdateQueueSlot200Response`

NewUpdateQueueSlot200Response instantiates a new UpdateQueueSlot200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQueueSlot200ResponseWithDefaults

`func NewUpdateQueueSlot200ResponseWithDefaults() *UpdateQueueSlot200Response`

NewUpdateQueueSlot200ResponseWithDefaults instantiates a new UpdateQueueSlot200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateQueueSlot200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateQueueSlot200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateQueueSlot200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateQueueSlot200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSchedule

`func (o *UpdateQueueSlot200Response) GetSchedule() QueueSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateQueueSlot200Response) GetScheduleOk() (*QueueSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateQueueSlot200Response) SetSchedule(v QueueSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateQueueSlot200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextSlots

`func (o *UpdateQueueSlot200Response) GetNextSlots() []time.Time`

GetNextSlots returns the NextSlots field if non-nil, zero value otherwise.

### GetNextSlotsOk

`func (o *UpdateQueueSlot200Response) GetNextSlotsOk() (*[]time.Time, bool)`

GetNextSlotsOk returns a tuple with the NextSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSlots

`func (o *UpdateQueueSlot200Response) SetNextSlots(v []time.Time)`

SetNextSlots sets NextSlots field to given value.

### HasNextSlots

`func (o *UpdateQueueSlot200Response) HasNextSlots() bool`

HasNextSlots returns a boolean if a field has been set.

### GetReshuffledCount

`func (o *UpdateQueueSlot200Response) GetReshuffledCount() int32`

GetReshuffledCount returns the ReshuffledCount field if non-nil, zero value otherwise.

### GetReshuffledCountOk

`func (o *UpdateQueueSlot200Response) GetReshuffledCountOk() (*int32, bool)`

GetReshuffledCountOk returns a tuple with the ReshuffledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReshuffledCount

`func (o *UpdateQueueSlot200Response) SetReshuffledCount(v int32)`

SetReshuffledCount sets ReshuffledCount field to given value.

### HasReshuffledCount

`func (o *UpdateQueueSlot200Response) HasReshuffledCount() bool`

HasReshuffledCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


