# QueueSlot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **int32** | Day of week (0&#x3D;Sunday, 6&#x3D;Saturday) | [optional] 
**Time** | Pointer to **string** | Time in HH:mm format (24-hour) | [optional] 

## Methods

### NewQueueSlot

`func NewQueueSlot() *QueueSlot`

NewQueueSlot instantiates a new QueueSlot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueSlotWithDefaults

`func NewQueueSlotWithDefaults() *QueueSlot`

NewQueueSlotWithDefaults instantiates a new QueueSlot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *QueueSlot) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *QueueSlot) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *QueueSlot) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *QueueSlot) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetTime

`func (o *QueueSlot) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *QueueSlot) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *QueueSlot) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *QueueSlot) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


