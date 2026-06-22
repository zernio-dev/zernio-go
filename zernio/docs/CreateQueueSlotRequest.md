# CreateQueueSlotRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Profile ID | 
**Name** | **string** | Queue name (e.g., Evening Posts) | 
**Timezone** | **string** | IANA timezone | 
**Slots** | [**[]QueueSlot**](QueueSlot.md) |  | 
**Active** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateQueueSlotRequest

`func NewCreateQueueSlotRequest(profileId string, name string, timezone string, slots []QueueSlot, ) *CreateQueueSlotRequest`

NewCreateQueueSlotRequest instantiates a new CreateQueueSlotRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateQueueSlotRequestWithDefaults

`func NewCreateQueueSlotRequestWithDefaults() *CreateQueueSlotRequest`

NewCreateQueueSlotRequestWithDefaults instantiates a new CreateQueueSlotRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateQueueSlotRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateQueueSlotRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateQueueSlotRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetName

`func (o *CreateQueueSlotRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateQueueSlotRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateQueueSlotRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTimezone

`func (o *CreateQueueSlotRequest) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *CreateQueueSlotRequest) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *CreateQueueSlotRequest) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetSlots

`func (o *CreateQueueSlotRequest) GetSlots() []QueueSlot`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *CreateQueueSlotRequest) GetSlotsOk() (*[]QueueSlot, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *CreateQueueSlotRequest) SetSlots(v []QueueSlot)`

SetSlots sets Slots field to given value.


### GetActive

`func (o *CreateQueueSlotRequest) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateQueueSlotRequest) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateQueueSlotRequest) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateQueueSlotRequest) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


