# AddWhatsAppGroupParticipantsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | **[]string** | Phone numbers in E.164 format (max 8) | 

## Methods

### NewAddWhatsAppGroupParticipantsRequest

`func NewAddWhatsAppGroupParticipantsRequest(phoneNumbers []string, ) *AddWhatsAppGroupParticipantsRequest`

NewAddWhatsAppGroupParticipantsRequest instantiates a new AddWhatsAppGroupParticipantsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddWhatsAppGroupParticipantsRequestWithDefaults

`func NewAddWhatsAppGroupParticipantsRequestWithDefaults() *AddWhatsAppGroupParticipantsRequest`

NewAddWhatsAppGroupParticipantsRequestWithDefaults instantiates a new AddWhatsAppGroupParticipantsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *AddWhatsAppGroupParticipantsRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *AddWhatsAppGroupParticipantsRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *AddWhatsAppGroupParticipantsRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


