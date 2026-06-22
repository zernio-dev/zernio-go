# OnWhatsAppNumberDeclinedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Number** | Pointer to [**OnWhatsAppNumberActivatedRequestNumber**](OnWhatsAppNumberActivatedRequestNumber.md) |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewOnWhatsAppNumberDeclinedRequest

`func NewOnWhatsAppNumberDeclinedRequest() *OnWhatsAppNumberDeclinedRequest`

NewOnWhatsAppNumberDeclinedRequest instantiates a new OnWhatsAppNumberDeclinedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnWhatsAppNumberDeclinedRequestWithDefaults

`func NewOnWhatsAppNumberDeclinedRequestWithDefaults() *OnWhatsAppNumberDeclinedRequest`

NewOnWhatsAppNumberDeclinedRequestWithDefaults instantiates a new OnWhatsAppNumberDeclinedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnWhatsAppNumberDeclinedRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnWhatsAppNumberDeclinedRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnWhatsAppNumberDeclinedRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnWhatsAppNumberDeclinedRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *OnWhatsAppNumberDeclinedRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OnWhatsAppNumberDeclinedRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OnWhatsAppNumberDeclinedRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OnWhatsAppNumberDeclinedRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *OnWhatsAppNumberDeclinedRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OnWhatsAppNumberDeclinedRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OnWhatsAppNumberDeclinedRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OnWhatsAppNumberDeclinedRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNumber

`func (o *OnWhatsAppNumberDeclinedRequest) GetNumber() OnWhatsAppNumberActivatedRequestNumber`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnWhatsAppNumberDeclinedRequest) GetNumberOk() (*OnWhatsAppNumberActivatedRequestNumber, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnWhatsAppNumberDeclinedRequest) SetNumber(v OnWhatsAppNumberActivatedRequestNumber)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnWhatsAppNumberDeclinedRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetReason

`func (o *OnWhatsAppNumberDeclinedRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *OnWhatsAppNumberDeclinedRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *OnWhatsAppNumberDeclinedRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *OnWhatsAppNumberDeclinedRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *OnWhatsAppNumberDeclinedRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *OnWhatsAppNumberDeclinedRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


