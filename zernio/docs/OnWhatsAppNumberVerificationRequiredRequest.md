# OnWhatsAppNumberVerificationRequiredRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Event** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Number** | Pointer to [**OnWhatsAppNumberActivatedRequestNumber**](OnWhatsAppNumberActivatedRequestNumber.md) |  | [optional] 
**VerificationUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewOnWhatsAppNumberVerificationRequiredRequest

`func NewOnWhatsAppNumberVerificationRequiredRequest() *OnWhatsAppNumberVerificationRequiredRequest`

NewOnWhatsAppNumberVerificationRequiredRequest instantiates a new OnWhatsAppNumberVerificationRequiredRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnWhatsAppNumberVerificationRequiredRequestWithDefaults

`func NewOnWhatsAppNumberVerificationRequiredRequestWithDefaults() *OnWhatsAppNumberVerificationRequiredRequest`

NewOnWhatsAppNumberVerificationRequiredRequestWithDefaults instantiates a new OnWhatsAppNumberVerificationRequiredRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OnWhatsAppNumberVerificationRequiredRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OnWhatsAppNumberVerificationRequiredRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEvent

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *OnWhatsAppNumberVerificationRequiredRequest) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *OnWhatsAppNumberVerificationRequiredRequest) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetTimestamp

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *OnWhatsAppNumberVerificationRequiredRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *OnWhatsAppNumberVerificationRequiredRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetNumber

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetNumber() OnWhatsAppNumberActivatedRequestNumber`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetNumberOk() (*OnWhatsAppNumberActivatedRequestNumber, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *OnWhatsAppNumberVerificationRequiredRequest) SetNumber(v OnWhatsAppNumberActivatedRequestNumber)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *OnWhatsAppNumberVerificationRequiredRequest) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetVerificationUrl

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetVerificationUrl() string`

GetVerificationUrl returns the VerificationUrl field if non-nil, zero value otherwise.

### GetVerificationUrlOk

`func (o *OnWhatsAppNumberVerificationRequiredRequest) GetVerificationUrlOk() (*string, bool)`

GetVerificationUrlOk returns a tuple with the VerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationUrl

`func (o *OnWhatsAppNumberVerificationRequiredRequest) SetVerificationUrl(v string)`

SetVerificationUrl sets VerificationUrl field to given value.

### HasVerificationUrl

`func (o *OnWhatsAppNumberVerificationRequiredRequest) HasVerificationUrl() bool`

HasVerificationUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


