# OnWhatsAppNumberActionRequiredRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Number** | Pointer to [**OnWhatsAppNumberActivatedRequestNumber**](OnWhatsAppNumberActivatedRequestNumber.md) |  | [optional] 

## Methods

### NewOnWhatsAppNumberActionRequiredRequest

`func NewOnWhatsAppNumberActionRequiredRequest() *OnWhatsAppNumberActionRequiredRequest`

NewOnWhatsAppNumberActionRequiredRequest instantiates a new OnWhatsAppNumberActionRequiredRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnWhatsAppNumberActionRequiredRequestWithDefaults

`func NewOnWhatsAppNumberActionRequiredRequestWithDefaults() *OnWhatsAppNumberActionRequiredRequest`

NewOnWhatsAppNumberActionRequiredRequestWithDefaults instantiates a new OnWhatsAppNumberActionRequiredRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnWhatsAppNumberActionRequiredRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnWhatsAppNumberActionRequiredRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnWhatsAppNumberActionRequiredRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnWhatsAppNumberActionRequiredRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *OnWhatsAppNumberActionRequiredRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OnWhatsAppNumberActionRequiredRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OnWhatsAppNumberActionRequiredRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OnWhatsAppNumberActionRequiredRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *OnWhatsAppNumberActionRequiredRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OnWhatsAppNumberActionRequiredRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OnWhatsAppNumberActionRequiredRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OnWhatsAppNumberActionRequiredRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetReason

`func (o *OnWhatsAppNumberActionRequiredRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OnWhatsAppNumberActionRequiredRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OnWhatsAppNumberActionRequiredRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OnWhatsAppNumberActionRequiredRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetNumber

`func (o *OnWhatsAppNumberActionRequiredRequest) GetNumber() OnWhatsAppNumberActivatedRequestNumber`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnWhatsAppNumberActionRequiredRequest) GetNumberOk() (*OnWhatsAppNumberActivatedRequestNumber, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnWhatsAppNumberActionRequiredRequest) SetNumber(v OnWhatsAppNumberActivatedRequestNumber)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnWhatsAppNumberActionRequiredRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


