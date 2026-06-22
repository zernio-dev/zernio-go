# EnableWhatsAppCallingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**ForwardTo** | **string** | tel:+E164 / sip:... / wss://... destination | 
**SipAuthUsername** | Pointer to **string** |  | [optional] 
**SipAuthPassword** | Pointer to **string** | Stored encrypted, never returned by any endpoint. | [optional] 
**RecordingEnabled** | Pointer to **bool** |  | [optional] [default to false]
**CallIconCountries** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEnableWhatsAppCallingRequest

`func NewEnableWhatsAppCallingRequest(accountId string, forwardTo string, ) *EnableWhatsAppCallingRequest`

NewEnableWhatsAppCallingRequest instantiates a new EnableWhatsAppCallingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableWhatsAppCallingRequestWithDefaults

`func NewEnableWhatsAppCallingRequestWithDefaults() *EnableWhatsAppCallingRequest`

NewEnableWhatsAppCallingRequestWithDefaults instantiates a new EnableWhatsAppCallingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EnableWhatsAppCallingRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EnableWhatsAppCallingRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EnableWhatsAppCallingRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetForwardTo

`func (o *EnableWhatsAppCallingRequest) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *EnableWhatsAppCallingRequest) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *EnableWhatsAppCallingRequest) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.


### GetSipAuthUsername

`func (o *EnableWhatsAppCallingRequest) GetSipAuthUsername() string`

GetSipAuthUsername returns the SipAuthUsername field if non-nil, zero value otherwise.

### GetSipAuthUsernameOk

`func (o *EnableWhatsAppCallingRequest) GetSipAuthUsernameOk() (*string, bool)`

GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthUsername

`func (o *EnableWhatsAppCallingRequest) SetSipAuthUsername(v string)`

SetSipAuthUsername sets SipAuthUsername field to given value.

### HasSipAuthUsername

`func (o *EnableWhatsAppCallingRequest) HasSipAuthUsername() bool`

HasSipAuthUsername returns a boolean if a field has been set.

### GetSipAuthPassword

`func (o *EnableWhatsAppCallingRequest) GetSipAuthPassword() string`

GetSipAuthPassword returns the SipAuthPassword field if non-nil, zero value otherwise.

### GetSipAuthPasswordOk

`func (o *EnableWhatsAppCallingRequest) GetSipAuthPasswordOk() (*string, bool)`

GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthPassword

`func (o *EnableWhatsAppCallingRequest) SetSipAuthPassword(v string)`

SetSipAuthPassword sets SipAuthPassword field to given value.

### HasSipAuthPassword

`func (o *EnableWhatsAppCallingRequest) HasSipAuthPassword() bool`

HasSipAuthPassword returns a boolean if a field has been set.

### GetRecordingEnabled

`func (o *EnableWhatsAppCallingRequest) GetRecordingEnabled() bool`

GetRecordingEnabled returns the RecordingEnabled field if non-nil, zero value otherwise.

### GetRecordingEnabledOk

`func (o *EnableWhatsAppCallingRequest) GetRecordingEnabledOk() (*bool, bool)`

GetRecordingEnabledOk returns a tuple with the RecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEnabled

`func (o *EnableWhatsAppCallingRequest) SetRecordingEnabled(v bool)`

SetRecordingEnabled sets RecordingEnabled field to given value.

### HasRecordingEnabled

`func (o *EnableWhatsAppCallingRequest) HasRecordingEnabled() bool`

HasRecordingEnabled returns a boolean if a field has been set.

### GetCallIconCountries

`func (o *EnableWhatsAppCallingRequest) GetCallIconCountries() []string`

GetCallIconCountries returns the CallIconCountries field if non-nil, zero value otherwise.

### GetCallIconCountriesOk

`func (o *EnableWhatsAppCallingRequest) GetCallIconCountriesOk() (*[]string, bool)`

GetCallIconCountriesOk returns a tuple with the CallIconCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallIconCountries

`func (o *EnableWhatsAppCallingRequest) SetCallIconCountries(v []string)`

SetCallIconCountries sets CallIconCountries field to given value.

### HasCallIconCountries

`func (o *EnableWhatsAppCallingRequest) HasCallIconCountries() bool`

HasCallIconCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


