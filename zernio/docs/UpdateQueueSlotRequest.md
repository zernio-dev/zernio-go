# UpdateQueueSlotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**QueueId** | Pointer to **string** | Queue ID to update (optional) | [optional] 
**Name** | Pointer to **string** | Queue name | [optional] 
**Timezone** | **string** |  | 
**Slots** | [**[]QueueSlot**](QueueSlot.md) |  | 
**Active** | Pointer to **bool** |  | [optional] [default to true]
**SetAsDefault** | Pointer to **bool** | Make this queue the default | [optional] 
**ReshuffleExisting** | Pointer to **bool** | Whether to reschedule existing queued posts to match new slots | [optional] [default to false]

## Methods

### NewUpdateQueueSlotRequest

`func NewUpdateQueueSlotRequest(profileId string, timezone string, slots []QueueSlot, ) *UpdateQueueSlotRequest`

NewUpdateQueueSlotRequest instantiates a new UpdateQueueSlotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateQueueSlotRequestWithDefaults

`func NewUpdateQueueSlotRequestWithDefaults() *UpdateQueueSlotRequest`

NewUpdateQueueSlotRequestWithDefaults instantiates a new UpdateQueueSlotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *UpdateQueueSlotRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *UpdateQueueSlotRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *UpdateQueueSlotRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetQueueId

`func (o *UpdateQueueSlotRequest) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *UpdateQueueSlotRequest) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *UpdateQueueSlotRequest) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *UpdateQueueSlotRequest) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetName

`func (o *UpdateQueueSlotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateQueueSlotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateQueueSlotRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateQueueSlotRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimezone

`func (o *UpdateQueueSlotRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UpdateQueueSlotRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UpdateQueueSlotRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetSlots

`func (o *UpdateQueueSlotRequest) GetSlots() []QueueSlot`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *UpdateQueueSlotRequest) GetSlotsOk() (*[]QueueSlot, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *UpdateQueueSlotRequest) SetSlots(v []QueueSlot)`

SetSlots sets Slots field to given value.


### GetActive

`func (o *UpdateQueueSlotRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateQueueSlotRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateQueueSlotRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateQueueSlotRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSetAsDefault

`func (o *UpdateQueueSlotRequest) GetSetAsDefault() bool`

GetSetAsDefault returns the SetAsDefault field if non-nil, zero value otherwise.

### GetSetAsDefaultOk

`func (o *UpdateQueueSlotRequest) GetSetAsDefaultOk() (*bool, bool)`

GetSetAsDefaultOk returns a tuple with the SetAsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetAsDefault

`func (o *UpdateQueueSlotRequest) SetSetAsDefault(v bool)`

SetSetAsDefault sets SetAsDefault field to given value.

### HasSetAsDefault

`func (o *UpdateQueueSlotRequest) HasSetAsDefault() bool`

HasSetAsDefault returns a boolean if a field has been set.

### GetReshuffleExisting

`func (o *UpdateQueueSlotRequest) GetReshuffleExisting() bool`

GetReshuffleExisting returns the ReshuffleExisting field if non-nil, zero value otherwise.

### GetReshuffleExistingOk

`func (o *UpdateQueueSlotRequest) GetReshuffleExistingOk() (*bool, bool)`

GetReshuffleExistingOk returns a tuple with the ReshuffleExisting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReshuffleExisting

`func (o *UpdateQueueSlotRequest) SetReshuffleExisting(v bool)`

SetReshuffleExisting sets ReshuffleExisting field to given value.

### HasReshuffleExisting

`func (o *UpdateQueueSlotRequest) HasReshuffleExisting() bool`

HasReshuffleExisting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


