# ConversionEventUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | Plaintext email. Hashed server-side. | [optional] 
**Phone** | Pointer to **string** | Phone number, ideally E.164. Hashed server-side. | [optional] 
**FirstName** | Pointer to **string** | Plaintext first name. Hashed server-side. | [optional] 
**LastName** | Pointer to **string** | Plaintext last name. Hashed server-side. | [optional] 
**ExternalId** | Pointer to **string** | Stable customer identifier (e.g. CRM user ID). Hashed server-side for Meta and Google. Sent as plaintext to LinkedIn (LinkedIn&#39;s Conversions API spec requires the raw value). Maximum effective list size on LinkedIn is 1.  | [optional] 
**IpAddress** | Pointer to **string** | Client IP address. Sent plaintext. | [optional] 
**UserAgent** | Pointer to **string** | Client user-agent string. Sent plaintext. | [optional] 
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country code, e.g. &#39;us&#39;. | [optional] 
**City** | Pointer to **string** | Meta advanced matching (ct). Plaintext city; normalized + SHA-256 hashed server-side. Meta only. | [optional] 
**State** | Pointer to **string** | Meta advanced matching (st). 2-letter ANSI for US; hashed server-side. Meta only. | [optional] 
**Zip** | Pointer to **string** | Meta advanced matching (zp). US uses first 5 digits; hashed server-side. Meta only. | [optional] 
**Dob** | Pointer to **string** | Meta advanced matching (db). YYYYMMDD; hashed server-side. Meta only. | [optional] 
**Gender** | Pointer to **string** | Meta advanced matching (ge). &#39;f&#39; or &#39;m&#39;; hashed server-side. Meta only. | [optional] 
**ClickIds** | Pointer to [**ConversionEventUserClickIds**](ConversionEventUserClickIds.md) |  | [optional] 

## Methods

### NewConversionEventUser

`func NewConversionEventUser() *ConversionEventUser`

NewConversionEventUser instantiates a new ConversionEventUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionEventUserWithDefaults

`func NewConversionEventUserWithDefaults() *ConversionEventUser`

NewConversionEventUserWithDefaults instantiates a new ConversionEventUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ConversionEventUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConversionEventUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConversionEventUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ConversionEventUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPhone

`func (o *ConversionEventUser) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *ConversionEventUser) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *ConversionEventUser) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *ConversionEventUser) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetFirstName

`func (o *ConversionEventUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *ConversionEventUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *ConversionEventUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *ConversionEventUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *ConversionEventUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *ConversionEventUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *ConversionEventUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *ConversionEventUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetExternalId

`func (o *ConversionEventUser) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ConversionEventUser) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ConversionEventUser) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ConversionEventUser) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetIpAddress

`func (o *ConversionEventUser) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ConversionEventUser) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ConversionEventUser) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ConversionEventUser) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetUserAgent

`func (o *ConversionEventUser) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ConversionEventUser) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ConversionEventUser) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ConversionEventUser) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetCountry

`func (o *ConversionEventUser) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ConversionEventUser) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ConversionEventUser) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *ConversionEventUser) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCity

`func (o *ConversionEventUser) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *ConversionEventUser) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *ConversionEventUser) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *ConversionEventUser) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *ConversionEventUser) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConversionEventUser) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConversionEventUser) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ConversionEventUser) HasState() bool`

HasState returns a boolean if a field has been set.

### GetZip

`func (o *ConversionEventUser) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *ConversionEventUser) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *ConversionEventUser) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *ConversionEventUser) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetDob

`func (o *ConversionEventUser) GetDob() string`

GetDob returns the Dob field if non-nil, zero value otherwise.

### GetDobOk

`func (o *ConversionEventUser) GetDobOk() (*string, bool)`

GetDobOk returns a tuple with the Dob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDob

`func (o *ConversionEventUser) SetDob(v string)`

SetDob sets Dob field to given value.

### HasDob

`func (o *ConversionEventUser) HasDob() bool`

HasDob returns a boolean if a field has been set.

### GetGender

`func (o *ConversionEventUser) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ConversionEventUser) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ConversionEventUser) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ConversionEventUser) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetClickIds

`func (o *ConversionEventUser) GetClickIds() ConversionEventUserClickIds`

GetClickIds returns the ClickIds field if non-nil, zero value otherwise.

### GetClickIdsOk

`func (o *ConversionEventUser) GetClickIdsOk() (*ConversionEventUserClickIds, bool)`

GetClickIdsOk returns a tuple with the ClickIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickIds

`func (o *ConversionEventUser) SetClickIds(v ConversionEventUserClickIds)`

SetClickIds sets ClickIds field to given value.

### HasClickIds

`func (o *ConversionEventUser) HasClickIds() bool`

HasClickIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


