# UpdateWhatsAppCallingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**ForwardTo** | Pointer to **string** |  | [optional] 
**SipAuthUsername** | Pointer to **NullableString** |  | [optional] 
**SipAuthPassword** | Pointer to **NullableString** |  | [optional] 
**RecordingEnabled** | Pointer to **bool** |  | [optional] 
**CallIconCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateWhatsAppCallingRequest

`func NewUpdateWhatsAppCallingRequest(accountId string, ) *UpdateWhatsAppCallingRequest`

NewUpdateWhatsAppCallingRequest instantiates a new UpdateWhatsAppCallingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWhatsAppCallingRequestWithDefaults

`func NewUpdateWhatsAppCallingRequestWithDefaults() *UpdateWhatsAppCallingRequest`

NewUpdateWhatsAppCallingRequestWithDefaults instantiates a new UpdateWhatsAppCallingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateWhatsAppCallingRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWhatsAppCallingRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWhatsAppCallingRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetForwardTo

`func (o *UpdateWhatsAppCallingRequest) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *UpdateWhatsAppCallingRequest) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *UpdateWhatsAppCallingRequest) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *UpdateWhatsAppCallingRequest) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *UpdateWhatsAppCallingRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *UpdateWhatsAppCallingRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *UpdateWhatsAppCallingRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *UpdateWhatsAppCallingRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### SetSipAuthUsernameNil

`func (o *UpdateWhatsAppCallingRequest) SetSipAuthUsernameNil(b bool)`

 SetSipAuthUsernameNil sets the value for SipAuthUsername to be an explicit nil

### UnsetSipAuthUsername
`func (o *UpdateWhatsAppCallingRequest) UnsetSipAuthUsername()`

UnsetSipAuthUsername ensures that no value is present for SipAuthUsername, not even an explicit nil
### GetSipAuthPassword

`func (o *UpdateWhatsAppCallingRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *UpdateWhatsAppCallingRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *UpdateWhatsAppCallingRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *UpdateWhatsAppCallingRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### SetSipAuthPasswordNil

`func (o *UpdateWhatsAppCallingRequest) SetSipAuthPasswordNil(b bool)`

 SetSipAuthPasswordNil sets the value for SipAuthPassword to be an explicit nil

### UnsetSipAuthPassword
`func (o *UpdateWhatsAppCallingRequest) UnsetSipAuthPassword()`

UnsetSipAuthPassword ensures that no value is present for SipAuthPassword, not even an explicit nil
### GetRecordingEnabled

`func (o *UpdateWhatsAppCallingRequest) GetRecordingEnabled() bool`

GetRecordingEnabled returns the RecordingEnabled field if non-nil, zero value otherwise.

### GetRecordingEnabledOk

`func (o *UpdateWhatsAppCallingRequest) GetRecordingEnabledOk() (*bool, bool)`

GetRecordingEnabledOk returns a tuple with the RecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEnabled

`func (o *UpdateWhatsAppCallingRequest) SetRecordingEnabled(v bool)`

SetRecordingEnabled sets RecordingEnabled field to given value.

### HasRecordingEnabled

`func (o *UpdateWhatsAppCallingRequest) HasRecordingEnabled() bool`

HasRecordingEnabled returns a boolean if a field has been set.

### GetCallIconCountries

`func (o *UpdateWhatsAppCallingRequest) GetCallIconCountries() []string`

GetCallIconCountries returns the CallIconCountries field if non-nil, zero value otherwise.

### GetCallIconCountriesOk

`func (o *UpdateWhatsAppCallingRequest) GetCallIconCountriesOk() (*[]string, bool)`

GetCallIconCountriesOk returns a tuple with the CallIconCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIconCountries

`func (o *UpdateWhatsAppCallingRequest) SetCallIconCountries(v []string)`

SetCallIconCountries sets CallIconCountries field to given value.

### HasCallIconCountries

`func (o *UpdateWhatsAppCallingRequest) HasCallIconCountries() bool`

HasCallIconCountries returns a boolean if a field has been set.

### SetCallIconCountriesNil

`func (o *UpdateWhatsAppCallingRequest) SetCallIconCountriesNil(b bool)`

 SetCallIconCountriesNil sets the value for CallIconCountries to be an explicit nil

### UnsetCallIconCountries
`func (o *UpdateWhatsAppCallingRequest) UnsetCallIconCountries()`

UnsetCallIconCountries ensures that no value is present for CallIconCountries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


