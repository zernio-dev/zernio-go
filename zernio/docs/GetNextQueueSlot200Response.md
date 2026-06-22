# GetNextQueueSlot200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** |  | [optional] 
**NextSlot** | Pointer to **time.Time** |  | [optional] 
**Timezone** | Pointer to **string** |  | [optional] 
**QueueId** | Pointer to **string** | Queue ID this slot belongs to | [optional] 
**QueueName** | Pointer to **string** | Queue name | [optional] 

## Methods

### NewGetNextQueueSlot200Response

`func NewGetNextQueueSlot200Response() *GetNextQueueSlot200Response`

NewGetNextQueueSlot200Response instantiates a new GetNextQueueSlot200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNextQueueSlot200ResponseWithDefaults

`func NewGetNextQueueSlot200ResponseWithDefaults() *GetNextQueueSlot200Response`

NewGetNextQueueSlot200ResponseWithDefaults instantiates a new GetNextQueueSlot200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *GetNextQueueSlot200Response) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetNextQueueSlot200Response) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetNextQueueSlot200Response) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetNextQueueSlot200Response) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetNextSlot

`func (o *GetNextQueueSlot200Response) GetNextSlot() time.Time`

GetNextSlot returns the NextSlot field if non-nil, zero value otherwise.

### GetNextSlotOk

`func (o *GetNextQueueSlot200Response) GetNextSlotOk() (*time.Time, bool)`

GetNextSlotOk returns a tuple with the NextSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSlot

`func (o *GetNextQueueSlot200Response) SetNextSlot(v time.Time)`

SetNextSlot sets NextSlot field to given value.

### HasNextSlot

`func (o *GetNextQueueSlot200Response) HasNextSlot() bool`

HasNextSlot returns a boolean if a field has been set.

### GetTimezone

`func (o *GetNextQueueSlot200Response) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *GetNextQueueSlot200Response) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *GetNextQueueSlot200Response) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *GetNextQueueSlot200Response) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetQueueId

`func (o *GetNextQueueSlot200Response) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *GetNextQueueSlot200Response) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *GetNextQueueSlot200Response) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *GetNextQueueSlot200Response) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetQueueName

`func (o *GetNextQueueSlot200Response) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *GetNextQueueSlot200Response) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *GetNextQueueSlot200Response) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *GetNextQueueSlot200Response) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


