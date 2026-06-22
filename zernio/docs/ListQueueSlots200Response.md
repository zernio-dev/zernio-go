# ListQueueSlots200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**QueueSchedule**](QueueSchedule.md) |  | [optional] 
**NextSlots** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 
**Queues** | Pointer to [**[]QueueSchedule**](QueueSchedule.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewListQueueSlots200Response

`func NewListQueueSlots200Response() *ListQueueSlots200Response`

NewListQueueSlots200Response instantiates a new ListQueueSlots200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListQueueSlots200ResponseWithDefaults

`func NewListQueueSlots200ResponseWithDefaults() *ListQueueSlots200Response`

NewListQueueSlots200ResponseWithDefaults instantiates a new ListQueueSlots200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *ListQueueSlots200Response) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *ListQueueSlots200Response) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *ListQueueSlots200Response) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *ListQueueSlots200Response) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetSchedule

`func (o *ListQueueSlots200Response) GetSchedule() QueueSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ListQueueSlots200Response) GetScheduleOk() (*QueueSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ListQueueSlots200Response) SetSchedule(v QueueSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ListQueueSlots200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextSlots

`func (o *ListQueueSlots200Response) GetNextSlots() []time.Time`

GetNextSlots returns the NextSlots field if non-nil, zero value otherwise.

### GetNextSlotsOk

`func (o *ListQueueSlots200Response) GetNextSlotsOk() (*[]time.Time, bool)`

GetNextSlotsOk returns a tuple with the NextSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSlots

`func (o *ListQueueSlots200Response) SetNextSlots(v []time.Time)`

SetNextSlots sets NextSlots field to given value.

### HasNextSlots

`func (o *ListQueueSlots200Response) HasNextSlots() bool`

HasNextSlots returns a boolean if a field has been set.

### GetQueues

`func (o *ListQueueSlots200Response) GetQueues() []QueueSchedule`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *ListQueueSlots200Response) GetQueuesOk() (*[]QueueSchedule, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *ListQueueSlots200Response) SetQueues(v []QueueSchedule)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *ListQueueSlots200Response) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetCount

`func (o *ListQueueSlots200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListQueueSlots200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListQueueSlots200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListQueueSlots200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


