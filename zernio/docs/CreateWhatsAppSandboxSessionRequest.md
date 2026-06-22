# CreateWhatsAppSandboxSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** | Recipient phone in international format. Digits, spaces, dashes and a leading &#x60;+&#x60; are all accepted; the server normalizes to E.164 digits-only. | 

## Methods

### NewCreateWhatsAppSandboxSessionRequest

`func NewCreateWhatsAppSandboxSessionRequest(phone string, ) *CreateWhatsAppSandboxSessionRequest`

NewCreateWhatsAppSandboxSessionRequest instantiates a new CreateWhatsAppSandboxSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhatsAppSandboxSessionRequestWithDefaults

`func NewCreateWhatsAppSandboxSessionRequestWithDefaults() *CreateWhatsAppSandboxSessionRequest`

NewCreateWhatsAppSandboxSessionRequestWithDefaults instantiates a new CreateWhatsAppSandboxSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhone

`func (o *CreateWhatsAppSandboxSessionRequest) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateWhatsAppSandboxSessionRequest) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateWhatsAppSandboxSessionRequest) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


