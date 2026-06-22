# GetWhatsAppCallingConfig200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumberDocId** | Pointer to **string** | WhatsAppPhoneNumber Mongo ID (use on /v1/whatsapp/phone-numbers/{id}/calling) | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**CallingEnabled** | Pointer to **bool** |  | [optional] 
**CallDeepLink** | Pointer to **NullableString** | Public calling deep link (https://wa.me/call/&lt;number&gt;). Tapping it on a phone starts a WhatsApp voice call to this number. Embed it on websites, emails, or QR codes. Null while calling is disabled; not supported by WhatsApp desktop clients. | [optional] 
**ForwardTo** | Pointer to **NullableString** | tel:+E164 / sip:... / wss://... destination | [optional] 
**RecordingEnabled** | Pointer to **bool** |  | [optional] 
**SipAuthUsername** | Pointer to **NullableString** |  | [optional] 
**SipAuthPasswordConfigured** | Pointer to **bool** | True when a SIP digest password is stored. The plaintext is never returned. | [optional] 
**CallIconCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetWhatsAppCallingConfig200Response

`func NewGetWhatsAppCallingConfig200Response() *GetWhatsAppCallingConfig200Response`

NewGetWhatsAppCallingConfig200Response instantiates a new GetWhatsAppCallingConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppCallingConfig200ResponseWithDefaults

`func NewGetWhatsAppCallingConfig200ResponseWithDefaults() *GetWhatsAppCallingConfig200Response`

NewGetWhatsAppCallingConfig200ResponseWithDefaults instantiates a new GetWhatsAppCallingConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumberDocId

`func (o *GetWhatsAppCallingConfig200Response) GetPhoneNumberDocId() string`

GetPhoneNumberDocId returns the PhoneNumberDocId field if non-nil, zero value otherwise.

### GetPhoneNumberDocIdOk

`func (o *GetWhatsAppCallingConfig200Response) GetPhoneNumberDocIdOk() (*string, bool)`

GetPhoneNumberDocIdOk returns a tuple with the PhoneNumberDocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberDocId

`func (o *GetWhatsAppCallingConfig200Response) SetPhoneNumberDocId(v string)`

SetPhoneNumberDocId sets PhoneNumberDocId field to given value.

### HasPhoneNumberDocId

`func (o *GetWhatsAppCallingConfig200Response) HasPhoneNumberDocId() bool`

HasPhoneNumberDocId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GetWhatsAppCallingConfig200Response) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GetWhatsAppCallingConfig200Response) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GetWhatsAppCallingConfig200Response) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GetWhatsAppCallingConfig200Response) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCallingEnabled

`func (o *GetWhatsAppCallingConfig200Response) GetCallingEnabled() bool`

GetCallingEnabled returns the CallingEnabled field if non-nil, zero value otherwise.

### GetCallingEnabledOk

`func (o *GetWhatsAppCallingConfig200Response) GetCallingEnabledOk() (*bool, bool)`

GetCallingEnabledOk returns a tuple with the CallingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallingEnabled

`func (o *GetWhatsAppCallingConfig200Response) SetCallingEnabled(v bool)`

SetCallingEnabled sets CallingEnabled field to given value.

### HasCallingEnabled

`func (o *GetWhatsAppCallingConfig200Response) HasCallingEnabled() bool`

HasCallingEnabled returns a boolean if a field has been set.

### GetCallDeepLink

`func (o *GetWhatsAppCallingConfig200Response) GetCallDeepLink() string`

GetCallDeepLink returns the CallDeepLink field if non-nil, zero value otherwise.

### GetCallDeepLinkOk

`func (o *GetWhatsAppCallingConfig200Response) GetCallDeepLinkOk() (*string, bool)`

GetCallDeepLinkOk returns a tuple with the CallDeepLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDeepLink

`func (o *GetWhatsAppCallingConfig200Response) SetCallDeepLink(v string)`

SetCallDeepLink sets CallDeepLink field to given value.

### HasCallDeepLink

`func (o *GetWhatsAppCallingConfig200Response) HasCallDeepLink() bool`

HasCallDeepLink returns a boolean if a field has been set.

### SetCallDeepLinkNil

`func (o *GetWhatsAppCallingConfig200Response) SetCallDeepLinkNil(b bool)`

 SetCallDeepLinkNil sets the value for CallDeepLink to be an explicit nil

### UnsetCallDeepLink
`func (o *GetWhatsAppCallingConfig200Response) UnsetCallDeepLink()`

UnsetCallDeepLink ensures that no value is present for CallDeepLink, not even an explicit nil
### GetForwardTo

`func (o *GetWhatsAppCallingConfig200Response) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *GetWhatsAppCallingConfig200Response) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *GetWhatsAppCallingConfig200Response) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *GetWhatsAppCallingConfig200Response) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### SetForwardToNil

`func (o *GetWhatsAppCallingConfig200Response) SetForwardToNil(b bool)`

 SetForwardToNil sets the value for ForwardTo to be an explicit nil

### UnsetForwardTo
`func (o *GetWhatsAppCallingConfig200Response) UnsetForwardTo()`

UnsetForwardTo ensures that no value is present for ForwardTo, not even an explicit nil
### GetRecordingEnabled

`func (o *GetWhatsAppCallingConfig200Response) GetRecordingEnabled() bool`

GetRecordingEnabled returns the RecordingEnabled field if non-nil, zero value otherwise.

### GetRecordingEnabledOk

`func (o *GetWhatsAppCallingConfig200Response) GetRecordingEnabledOk() (*bool, bool)`

GetRecordingEnabledOk returns a tuple with the RecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEnabled

`func (o *GetWhatsAppCallingConfig200Response) SetRecordingEnabled(v bool)`

SetRecordingEnabled sets RecordingEnabled field to given value.

### HasRecordingEnabled

`func (o *GetWhatsAppCallingConfig200Response) HasRecordingEnabled() bool`

HasRecordingEnabled returns a boolean if a field has been set.

### GetSipAuthUsername

`func (o *GetWhatsAppCallingConfig200Response) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *GetWhatsAppCallingConfig200Response) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *GetWhatsAppCallingConfig200Response) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *GetWhatsAppCallingConfig200Response) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### SetSipAuthUsernameNil

`func (o *GetWhatsAppCallingConfig200Response) SetSipAuthUsernameNil(b bool)`

 SetSipAuthUsernameNil sets the value for SipAuthUsername to be an explicit nil

### UnsetSipAuthUsername
`func (o *GetWhatsAppCallingConfig200Response) UnsetSipAuthUsername()`

UnsetSipAuthUsername ensures that no value is present for SipAuthUsername, not even an explicit nil
### GetSipAuthPasswordConfigured

`func (o *GetWhatsAppCallingConfig200Response) GetSipAuthPasswordConfigured() bool`

GetSipAuthPasswordConfigured returns the SipAuthPasswordConfigured field if non-nil, zero value otherwise.

### GetSipAuthPasswordConfiguredOk

`func (o *GetWhatsAppCallingConfig200Response) GetSipAuthPasswordConfiguredOk() (*bool, bool)`

GetSipAuthPasswordConfiguredOk returns a tuple with the SipAuthPasswordConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPasswordConfigured

`func (o *GetWhatsAppCallingConfig200Response) SetSipAuthPasswordConfigured(v bool)`

SetSipAuthPasswordConfigured sets SipAuthPasswordConfigured field to given value.

### HasSipAuthPasswordConfigured

`func (o *GetWhatsAppCallingConfig200Response) HasSipAuthPasswordConfigured() bool`

HasSipAuthPasswordConfigured returns a boolean if a field has been set.

### GetCallIconCountries

`func (o *GetWhatsAppCallingConfig200Response) GetCallIconCountries() []string`

GetCallIconCountries returns the CallIconCountries field if non-nil, zero value otherwise.

### GetCallIconCountriesOk

`func (o *GetWhatsAppCallingConfig200Response) GetCallIconCountriesOk() (*[]string, bool)`

GetCallIconCountriesOk returns a tuple with the CallIconCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIconCountries

`func (o *GetWhatsAppCallingConfig200Response) SetCallIconCountries(v []string)`

SetCallIconCountries sets CallIconCountries field to given value.

### HasCallIconCountries

`func (o *GetWhatsAppCallingConfig200Response) HasCallIconCountries() bool`

HasCallIconCountries returns a boolean if a field has been set.

### SetCallIconCountriesNil

`func (o *GetWhatsAppCallingConfig200Response) SetCallIconCountriesNil(b bool)`

 SetCallIconCountriesNil sets the value for CallIconCountries to be an explicit nil

### UnsetCallIconCountries
`func (o *GetWhatsAppCallingConfig200Response) UnsetCallIconCountries()`

UnsetCallIconCountries ensures that no value is present for CallIconCountries, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


