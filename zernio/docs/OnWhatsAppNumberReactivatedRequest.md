# OnWhatsAppNumberReactivatedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Number** | Pointer to [**OnWhatsAppNumberActivatedRequestNumber**](OnWhatsAppNumberActivatedRequestNumber.md) |  | [optional] 

## Methods

### NewOnWhatsAppNumberReactivatedRequest

`func NewOnWhatsAppNumberReactivatedRequest() *OnWhatsAppNumberReactivatedRequest`

NewOnWhatsAppNumberReactivatedRequest instantiates a new OnWhatsAppNumberReactivatedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnWhatsAppNumberReactivatedRequestWithDefaults

`func NewOnWhatsAppNumberReactivatedRequestWithDefaults() *OnWhatsAppNumberReactivatedRequest`

NewOnWhatsAppNumberReactivatedRequestWithDefaults instantiates a new OnWhatsAppNumberReactivatedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnWhatsAppNumberReactivatedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnWhatsAppNumberReactivatedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnWhatsAppNumberReactivatedRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnWhatsAppNumberReactivatedRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *OnWhatsAppNumberReactivatedRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OnWhatsAppNumberReactivatedRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OnWhatsAppNumberReactivatedRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OnWhatsAppNumberReactivatedRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *OnWhatsAppNumberReactivatedRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OnWhatsAppNumberReactivatedRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OnWhatsAppNumberReactivatedRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OnWhatsAppNumberReactivatedRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNumber

`func (o *OnWhatsAppNumberReactivatedRequest) GetNumber() OnWhatsAppNumberActivatedRequestNumber`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnWhatsAppNumberReactivatedRequest) GetNumberOk() (*OnWhatsAppNumberActivatedRequestNumber, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnWhatsAppNumberReactivatedRequest) SetNumber(v OnWhatsAppNumberActivatedRequestNumber)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnWhatsAppNumberReactivatedRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


