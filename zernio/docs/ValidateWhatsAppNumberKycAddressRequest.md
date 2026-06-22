# ValidateWhatsAppNumberKycAddressRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | **string** | ISO 3166-1 alpha-2 country code. | 
**StreetAddress** | **string** |  | 
**Locality** | **string** | City / town. | 
**AdministrativeArea** | Pointer to **string** | State / province / region. When omitted, the pre-check is skipped (the final submit still validates). | [optional] 
**PostalCode** | **string** |  | 

## Methods

### NewValidateWhatsAppNumberKycAddressRequest

`func NewValidateWhatsAppNumberKycAddressRequest(country string, streetAddress string, locality string, postalCode string, ) *ValidateWhatsAppNumberKycAddressRequest`

NewValidateWhatsAppNumberKycAddressRequest instantiates a new ValidateWhatsAppNumberKycAddressRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateWhatsAppNumberKycAddressRequestWithDefaults

`func NewValidateWhatsAppNumberKycAddressRequestWithDefaults() *ValidateWhatsAppNumberKycAddressRequest`

NewValidateWhatsAppNumberKycAddressRequestWithDefaults instantiates a new ValidateWhatsAppNumberKycAddressRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *ValidateWhatsAppNumberKycAddressRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetStreetAddress

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetStreetAddress() string`

GetStreetAddress returns the StreetAddress field if non-nil, zero value otherwise.

### GetStreetAddressOk

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetStreetAddressOk() (*string, bool)`

GetStreetAddressOk returns a tuple with the StreetAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress

`func (o *ValidateWhatsAppNumberKycAddressRequest) SetStreetAddress(v string)`

SetStreetAddress sets StreetAddress field to given value.


### GetLocality

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetLocality() string`

GetLocality returns the Locality field if non-nil, zero value otherwise.

### GetLocalityOk

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetLocalityOk() (*string, bool)`

GetLocalityOk returns a tuple with the Locality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocality

`func (o *ValidateWhatsAppNumberKycAddressRequest) SetLocality(v string)`

SetLocality sets Locality field to given value.


### GetAdministrativeArea

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetAdministrativeArea() string`

GetAdministrativeArea returns the AdministrativeArea field if non-nil, zero value otherwise.

### GetAdministrativeAreaOk

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetAdministrativeAreaOk() (*string, bool)`

GetAdministrativeAreaOk returns a tuple with the AdministrativeArea field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministrativeArea

`func (o *ValidateWhatsAppNumberKycAddressRequest) SetAdministrativeArea(v string)`

SetAdministrativeArea sets AdministrativeArea field to given value.

### HasAdministrativeArea

`func (o *ValidateWhatsAppNumberKycAddressRequest) HasAdministrativeArea() bool`

HasAdministrativeArea returns a boolean if a field has been set.

### GetPostalCode

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *ValidateWhatsAppNumberKycAddressRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *ValidateWhatsAppNumberKycAddressRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


