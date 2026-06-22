# CompleteWhatsAppPhoneSelectionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | The Zernio profile ID | 
**PhoneNumberId** | **string** | The selected phone number ID (from listWhatsAppPhoneNumbers) | 
**WabaId** | **string** | The WABA ID containing the selected phone | 
**TempToken** | **string** | The temporary access token from the headless redirect | 
**UserProfile** | Pointer to **map[string]interface{}** | Optional user profile data (passthrough) | [optional] 
**RedirectUrl** | Pointer to **string** | Optional URL to receive the post-connection redirect target | [optional] 

## Methods

### NewCompleteWhatsAppPhoneSelectionRequest

`func NewCompleteWhatsAppPhoneSelectionRequest(profileId string, phoneNumberId string, wabaId string, tempToken string, ) *CompleteWhatsAppPhoneSelectionRequest`

NewCompleteWhatsAppPhoneSelectionRequest instantiates a new CompleteWhatsAppPhoneSelectionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteWhatsAppPhoneSelectionRequestWithDefaults

`func NewCompleteWhatsAppPhoneSelectionRequestWithDefaults() *CompleteWhatsAppPhoneSelectionRequest`

NewCompleteWhatsAppPhoneSelectionRequestWithDefaults instantiates a new CompleteWhatsAppPhoneSelectionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CompleteWhatsAppPhoneSelectionRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetPhoneNumberId

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *CompleteWhatsAppPhoneSelectionRequest) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.


### GetWabaId

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *CompleteWhatsAppPhoneSelectionRequest) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.


### GetTempToken

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *CompleteWhatsAppPhoneSelectionRequest) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.


### GetUserProfile

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetUserProfile() map[string]interface{}`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetUserProfileOk() (*map[string]interface{}, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *CompleteWhatsAppPhoneSelectionRequest) SetUserProfile(v map[string]interface{})`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *CompleteWhatsAppPhoneSelectionRequest) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CompleteWhatsAppPhoneSelectionRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CompleteWhatsAppPhoneSelectionRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CompleteWhatsAppPhoneSelectionRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


