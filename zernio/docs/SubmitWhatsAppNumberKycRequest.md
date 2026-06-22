# SubmitWhatsAppNumberKycRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**Country** | **string** |  | 
**SubmissionId** | Pointer to **string** | Idempotency token for this submission attempt. A retry/double-submit with the same token returns the same number; omit and each call creates a new number. | [optional] 
**Reuse** | Pointer to **bool** | Reuse a prior approved verification for this country (skips document/field collection; places the order immediately). | [optional] 
**ReuseFrom** | Pointer to **string** | Which approved verification to reuse when several exist: the phone number it was originally approved for (GET reusable.options[].fromPhoneNumber). Omitted &#x3D; newest. No match &#x3D; 409. | [optional] 
**EndUserFirstName** | Pointer to **string** | End user&#39;s legal first name. Required when the country has an action/ID-verification (Onfido) requirement. | [optional] 
**EndUserLastName** | Pointer to **string** | End user&#39;s legal last name. Same condition as endUserFirstName. | [optional] 
**Values** | Pointer to **map[string]string** | requirementId → textual value | [optional] 
**Documents** | Pointer to [**[]SubmitWhatsAppNumberKycRequestDocumentsInner**](SubmitWhatsAppNumberKycRequestDocumentsInner.md) | One per document requirement. Each is EITHER inline base64 OR a &#x60;documentId&#x60; returned by POST /v1/whatsapp/phone-numbers/kyc/upload-document (use the upload endpoint for large files to stay under the request-size limit). | [optional] 
**Address** | Pointer to [**SubmitWhatsAppNumberKycRequestAddress**](SubmitWhatsAppNumberKycRequestAddress.md) |  | [optional] 

## Methods

### NewSubmitWhatsAppNumberKycRequest

`func NewSubmitWhatsAppNumberKycRequest(profileId string, country string, ) *SubmitWhatsAppNumberKycRequest`

NewSubmitWhatsAppNumberKycRequest instantiates a new SubmitWhatsAppNumberKycRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitWhatsAppNumberKycRequestWithDefaults

`func NewSubmitWhatsAppNumberKycRequestWithDefaults() *SubmitWhatsAppNumberKycRequest`

NewSubmitWhatsAppNumberKycRequestWithDefaults instantiates a new SubmitWhatsAppNumberKycRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SubmitWhatsAppNumberKycRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SubmitWhatsAppNumberKycRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SubmitWhatsAppNumberKycRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetCountry

`func (o *SubmitWhatsAppNumberKycRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SubmitWhatsAppNumberKycRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SubmitWhatsAppNumberKycRequest) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetSubmissionId

`func (o *SubmitWhatsAppNumberKycRequest) GetSubmissionId() string`

GetSubmissionId returns the SubmissionId field if non-nil, zero value otherwise.

### GetSubmissionIdOk

`func (o *SubmitWhatsAppNumberKycRequest) GetSubmissionIdOk() (*string, bool)`

GetSubmissionIdOk returns a tuple with the SubmissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionId

`func (o *SubmitWhatsAppNumberKycRequest) SetSubmissionId(v string)`

SetSubmissionId sets SubmissionId field to given value.

### HasSubmissionId

`func (o *SubmitWhatsAppNumberKycRequest) HasSubmissionId() bool`

HasSubmissionId returns a boolean if a field has been set.

### GetReuse

`func (o *SubmitWhatsAppNumberKycRequest) GetReuse() bool`

GetReuse returns the Reuse field if non-nil, zero value otherwise.

### GetReuseOk

`func (o *SubmitWhatsAppNumberKycRequest) GetReuseOk() (*bool, bool)`

GetReuseOk returns a tuple with the Reuse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuse

`func (o *SubmitWhatsAppNumberKycRequest) SetReuse(v bool)`

SetReuse sets Reuse field to given value.

### HasReuse

`func (o *SubmitWhatsAppNumberKycRequest) HasReuse() bool`

HasReuse returns a boolean if a field has been set.

### GetReuseFrom

`func (o *SubmitWhatsAppNumberKycRequest) GetReuseFrom() string`

GetReuseFrom returns the ReuseFrom field if non-nil, zero value otherwise.

### GetReuseFromOk

`func (o *SubmitWhatsAppNumberKycRequest) GetReuseFromOk() (*string, bool)`

GetReuseFromOk returns a tuple with the ReuseFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseFrom

`func (o *SubmitWhatsAppNumberKycRequest) SetReuseFrom(v string)`

SetReuseFrom sets ReuseFrom field to given value.

### HasReuseFrom

`func (o *SubmitWhatsAppNumberKycRequest) HasReuseFrom() bool`

HasReuseFrom returns a boolean if a field has been set.

### GetEndUserFirstName

`func (o *SubmitWhatsAppNumberKycRequest) GetEndUserFirstName() string`

GetEndUserFirstName returns the EndUserFirstName field if non-nil, zero value otherwise.

### GetEndUserFirstNameOk

`func (o *SubmitWhatsAppNumberKycRequest) GetEndUserFirstNameOk() (*string, bool)`

GetEndUserFirstNameOk returns a tuple with the EndUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserFirstName

`func (o *SubmitWhatsAppNumberKycRequest) SetEndUserFirstName(v string)`

SetEndUserFirstName sets EndUserFirstName field to given value.

### HasEndUserFirstName

`func (o *SubmitWhatsAppNumberKycRequest) HasEndUserFirstName() bool`

HasEndUserFirstName returns a boolean if a field has been set.

### GetEndUserLastName

`func (o *SubmitWhatsAppNumberKycRequest) GetEndUserLastName() string`

GetEndUserLastName returns the EndUserLastName field if non-nil, zero value otherwise.

### GetEndUserLastNameOk

`func (o *SubmitWhatsAppNumberKycRequest) GetEndUserLastNameOk() (*string, bool)`

GetEndUserLastNameOk returns a tuple with the EndUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserLastName

`func (o *SubmitWhatsAppNumberKycRequest) SetEndUserLastName(v string)`

SetEndUserLastName sets EndUserLastName field to given value.

### HasEndUserLastName

`func (o *SubmitWhatsAppNumberKycRequest) HasEndUserLastName() bool`

HasEndUserLastName returns a boolean if a field has been set.

### GetValues

`func (o *SubmitWhatsAppNumberKycRequest) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SubmitWhatsAppNumberKycRequest) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SubmitWhatsAppNumberKycRequest) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *SubmitWhatsAppNumberKycRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetDocuments

`func (o *SubmitWhatsAppNumberKycRequest) GetDocuments() []SubmitWhatsAppNumberKycRequestDocumentsInner`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *SubmitWhatsAppNumberKycRequest) GetDocumentsOk() (*[]SubmitWhatsAppNumberKycRequestDocumentsInner, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *SubmitWhatsAppNumberKycRequest) SetDocuments(v []SubmitWhatsAppNumberKycRequestDocumentsInner)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *SubmitWhatsAppNumberKycRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetAddress

`func (o *SubmitWhatsAppNumberKycRequest) GetAddress() SubmitWhatsAppNumberKycRequestAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SubmitWhatsAppNumberKycRequest) GetAddressOk() (*SubmitWhatsAppNumberKycRequestAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SubmitWhatsAppNumberKycRequest) SetAddress(v SubmitWhatsAppNumberKycRequestAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SubmitWhatsAppNumberKycRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


