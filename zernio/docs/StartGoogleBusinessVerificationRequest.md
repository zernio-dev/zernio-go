# StartGoogleBusinessVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | **string** | The verification method. Selects which method-specific field below is required. | 
**LanguageCode** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** | For PHONE_CALL / SMS. | [optional] 
**EmailAddress** | Pointer to **string** | For EMAIL. | [optional] 
**MailerContact** | Pointer to **map[string]interface{}** | For ADDRESS (postcard) verification. | [optional] 
**Context** | Pointer to **map[string]interface{}** | ServiceBusinessContext (e.g. service address). Required for service-area businesses. | [optional] 

## Methods

### NewStartGoogleBusinessVerificationRequest

`func NewStartGoogleBusinessVerificationRequest(method string, ) *StartGoogleBusinessVerificationRequest`

NewStartGoogleBusinessVerificationRequest instantiates a new StartGoogleBusinessVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartGoogleBusinessVerificationRequestWithDefaults

`func NewStartGoogleBusinessVerificationRequestWithDefaults() *StartGoogleBusinessVerificationRequest`

NewStartGoogleBusinessVerificationRequestWithDefaults instantiates a new StartGoogleBusinessVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *StartGoogleBusinessVerificationRequest) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *StartGoogleBusinessVerificationRequest) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *StartGoogleBusinessVerificationRequest) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetLanguageCode

`func (o *StartGoogleBusinessVerificationRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *StartGoogleBusinessVerificationRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *StartGoogleBusinessVerificationRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *StartGoogleBusinessVerificationRequest) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *StartGoogleBusinessVerificationRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *StartGoogleBusinessVerificationRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *StartGoogleBusinessVerificationRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *StartGoogleBusinessVerificationRequest) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetEmailAddress

`func (o *StartGoogleBusinessVerificationRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *StartGoogleBusinessVerificationRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *StartGoogleBusinessVerificationRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *StartGoogleBusinessVerificationRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetMailerContact

`func (o *StartGoogleBusinessVerificationRequest) GetMailerContact() map[string]interface{}`

GetMailerContact returns the MailerContact field if non-nil, zero value otherwise.

### GetMailerContactOk

`func (o *StartGoogleBusinessVerificationRequest) GetMailerContactOk() (*map[string]interface{}, bool)`

GetMailerContactOk returns a tuple with the MailerContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailerContact

`func (o *StartGoogleBusinessVerificationRequest) SetMailerContact(v map[string]interface{})`

SetMailerContact sets MailerContact field to given value.

### HasMailerContact

`func (o *StartGoogleBusinessVerificationRequest) HasMailerContact() bool`

HasMailerContact returns a boolean if a field has been set.

### GetContext

`func (o *StartGoogleBusinessVerificationRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *StartGoogleBusinessVerificationRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *StartGoogleBusinessVerificationRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *StartGoogleBusinessVerificationRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


