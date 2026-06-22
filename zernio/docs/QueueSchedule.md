# QueueSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique queue identifier | [optional] 
**ProfileId** | Pointer to **string** | Profile ID this queue belongs to | [optional] 
**Name** | Pointer to **string** | Queue name (e.g., \&quot;Morning Posts\&quot;, \&quot;Evening Content\&quot;) | [optional] 
**Timezone** | Pointer to **string** | IANA timezone (e.g., America/New_York) | [optional] 
**Slots** | Pointer to [**[]QueueSlot**](QueueSlot.md) |  | [optional] 
**Active** | Pointer to **bool** | Whether the queue is active | [optional] 
**IsDefault** | Pointer to **bool** | Whether this is the default queue for the profile (used when no queueId specified) | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewQueueSchedule

`func NewQueueSchedule() *QueueSchedule`

NewQueueSchedule instantiates a new QueueSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueScheduleWithDefaults

`func NewQueueScheduleWithDefaults() *QueueSchedule`

NewQueueScheduleWithDefaults instantiates a new QueueSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueueSchedule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueSchedule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueSchedule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QueueSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileId

`func (o *QueueSchedule) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *QueueSchedule) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *QueueSchedule) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *QueueSchedule) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetName

`func (o *QueueSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueueSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueueSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueueSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimezone

`func (o *QueueSchedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *QueueSchedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *QueueSchedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *QueueSchedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetSlots

`func (o *QueueSchedule) GetSlots() []QueueSlot`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *QueueSchedule) GetSlotsOk() (*[]QueueSlot, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *QueueSchedule) SetSlots(v []QueueSlot)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *QueueSchedule) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetActive

`func (o *QueueSchedule) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *QueueSchedule) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *QueueSchedule) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *QueueSchedule) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetIsDefault

`func (o *QueueSchedule) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *QueueSchedule) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *QueueSchedule) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *QueueSchedule) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetCreatedAt

`func (o *QueueSchedule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *QueueSchedule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *QueueSchedule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *QueueSchedule) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *QueueSchedule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *QueueSchedule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *QueueSchedule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *QueueSchedule) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


