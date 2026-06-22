# CreateQueueSlot201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**QueueSchedule**](QueueSchedule.md) |  | [optional] 
**NextSlots** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewCreateQueueSlot201Response

`func NewCreateQueueSlot201Response() *CreateQueueSlot201Response`

NewCreateQueueSlot201Response instantiates a new CreateQueueSlot201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQueueSlot201ResponseWithDefaults

`func NewCreateQueueSlot201ResponseWithDefaults() *CreateQueueSlot201Response`

NewCreateQueueSlot201ResponseWithDefaults instantiates a new CreateQueueSlot201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateQueueSlot201Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateQueueSlot201Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateQueueSlot201Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateQueueSlot201Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSchedule

`func (o *CreateQueueSlot201Response) GetSchedule() QueueSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CreateQueueSlot201Response) GetScheduleOk() (*QueueSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CreateQueueSlot201Response) SetSchedule(v QueueSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CreateQueueSlot201Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextSlots

`func (o *CreateQueueSlot201Response) GetNextSlots() []time.Time`

GetNextSlots returns the NextSlots field if non-nil, zero value otherwise.

### GetNextSlotsOk

`func (o *CreateQueueSlot201Response) GetNextSlotsOk() (*[]time.Time, bool)`

GetNextSlotsOk returns a tuple with the NextSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSlots

`func (o *CreateQueueSlot201Response) SetNextSlots(v []time.Time)`

SetNextSlots sets NextSlots field to given value.

### HasNextSlots

`func (o *CreateQueueSlot201Response) HasNextSlots() bool`

HasNextSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


