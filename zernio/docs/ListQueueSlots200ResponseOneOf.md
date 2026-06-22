# ListQueueSlots200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**QueueSchedule**](QueueSchedule.md) |  | [optional] 
**NextSlots** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewListQueueSlots200ResponseOneOf

`func NewListQueueSlots200ResponseOneOf() *ListQueueSlots200ResponseOneOf`

NewListQueueSlots200ResponseOneOf instantiates a new ListQueueSlots200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListQueueSlots200ResponseOneOfWithDefaults

`func NewListQueueSlots200ResponseOneOfWithDefaults() *ListQueueSlots200ResponseOneOf`

NewListQueueSlots200ResponseOneOfWithDefaults instantiates a new ListQueueSlots200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *ListQueueSlots200ResponseOneOf) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *ListQueueSlots200ResponseOneOf) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *ListQueueSlots200ResponseOneOf) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *ListQueueSlots200ResponseOneOf) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetSchedule

`func (o *ListQueueSlots200ResponseOneOf) GetSchedule() QueueSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ListQueueSlots200ResponseOneOf) GetScheduleOk() (*QueueSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ListQueueSlots200ResponseOneOf) SetSchedule(v QueueSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ListQueueSlots200ResponseOneOf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetNextSlots

`func (o *ListQueueSlots200ResponseOneOf) GetNextSlots() []time.Time`

GetNextSlots returns the NextSlots field if non-nil, zero value otherwise.

### GetNextSlotsOk

`func (o *ListQueueSlots200ResponseOneOf) GetNextSlotsOk() (*[]time.Time, bool)`

GetNextSlotsOk returns a tuple with the NextSlots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSlots

`func (o *ListQueueSlots200ResponseOneOf) SetNextSlots(v []time.Time)`

SetNextSlots sets NextSlots field to given value.

### HasNextSlots

`func (o *ListQueueSlots200ResponseOneOf) HasNextSlots() bool`

HasNextSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


