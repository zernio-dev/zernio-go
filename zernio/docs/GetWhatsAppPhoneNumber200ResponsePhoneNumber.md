# GetWhatsAppPhoneNumber200ResponsePhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**MetaPreverifiedId** | Pointer to **string** |  | [optional] 
**MetaVerificationStatus** | Pointer to **string** |  | [optional] 
**OnfidoVerificationUrl** | Pointer to **NullableString** | For a regulated number with an Onfido ID step — the link to forward to the end user. Appears once the order is placed; null otherwise. | [optional] 
**EndUserFirstName** | Pointer to **NullableString** |  | [optional] 
**EndUserLastName** | Pointer to **NullableString** |  | [optional] 
**RegulatoryDeclineReason** | Pointer to **NullableString** | Reviewer rejection reason when status is regulatory_declined. | [optional] 
**ProvisionedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetWhatsAppPhoneNumber200ResponsePhoneNumber

`func NewGetWhatsAppPhoneNumber200ResponsePhoneNumber() *GetWhatsAppPhoneNumber200ResponsePhoneNumber`

NewGetWhatsAppPhoneNumber200ResponsePhoneNumber instantiates a new GetWhatsAppPhoneNumber200ResponsePhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppPhoneNumber200ResponsePhoneNumberWithDefaults

`func NewGetWhatsAppPhoneNumber200ResponsePhoneNumberWithDefaults() *GetWhatsAppPhoneNumber200ResponsePhoneNumber`

NewGetWhatsAppPhoneNumber200ResponsePhoneNumberWithDefaults instantiates a new GetWhatsAppPhoneNumber200ResponsePhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCountry

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetMetaPreverifiedId

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetMetaPreverifiedId() string`

GetMetaPreverifiedId returns the MetaPreverifiedId field if non-nil, zero value otherwise.

### GetMetaPreverifiedIdOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetMetaPreverifiedIdOk() (*string, bool)`

GetMetaPreverifiedIdOk returns a tuple with the MetaPreverifiedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPreverifiedId

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetMetaPreverifiedId(v string)`

SetMetaPreverifiedId sets MetaPreverifiedId field to given value.

### HasMetaPreverifiedId

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasMetaPreverifiedId() bool`

HasMetaPreverifiedId returns a boolean if a field has been set.

### GetMetaVerificationStatus

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetMetaVerificationStatus() string`

GetMetaVerificationStatus returns the MetaVerificationStatus field if non-nil, zero value otherwise.

### GetMetaVerificationStatusOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetMetaVerificationStatusOk() (*string, bool)`

GetMetaVerificationStatusOk returns a tuple with the MetaVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaVerificationStatus

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetMetaVerificationStatus(v string)`

SetMetaVerificationStatus sets MetaVerificationStatus field to given value.

### HasMetaVerificationStatus

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasMetaVerificationStatus() bool`

HasMetaVerificationStatus returns a boolean if a field has been set.

### GetOnfidoVerificationUrl

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetOnfidoVerificationUrl() string`

GetOnfidoVerificationUrl returns the OnfidoVerificationUrl field if non-nil, zero value otherwise.

### GetOnfidoVerificationUrlOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetOnfidoVerificationUrlOk() (*string, bool)`

GetOnfidoVerificationUrlOk returns a tuple with the OnfidoVerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnfidoVerificationUrl

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetOnfidoVerificationUrl(v string)`

SetOnfidoVerificationUrl sets OnfidoVerificationUrl field to given value.

### HasOnfidoVerificationUrl

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasOnfidoVerificationUrl() bool`

HasOnfidoVerificationUrl returns a boolean if a field has been set.

### SetOnfidoVerificationUrlNil

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetOnfidoVerificationUrlNil(b bool)`

 SetOnfidoVerificationUrlNil sets the value for OnfidoVerificationUrl to be an explicit nil

### UnsetOnfidoVerificationUrl
`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) UnsetOnfidoVerificationUrl()`

UnsetOnfidoVerificationUrl ensures that no value is present for OnfidoVerificationUrl, not even an explicit nil
### GetEndUserFirstName

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetEndUserFirstName() string`

GetEndUserFirstName returns the EndUserFirstName field if non-nil, zero value otherwise.

### GetEndUserFirstNameOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetEndUserFirstNameOk() (*string, bool)`

GetEndUserFirstNameOk returns a tuple with the EndUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserFirstName

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetEndUserFirstName(v string)`

SetEndUserFirstName sets EndUserFirstName field to given value.

### HasEndUserFirstName

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasEndUserFirstName() bool`

HasEndUserFirstName returns a boolean if a field has been set.

### SetEndUserFirstNameNil

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetEndUserFirstNameNil(b bool)`

 SetEndUserFirstNameNil sets the value for EndUserFirstName to be an explicit nil

### UnsetEndUserFirstName
`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) UnsetEndUserFirstName()`

UnsetEndUserFirstName ensures that no value is present for EndUserFirstName, not even an explicit nil
### GetEndUserLastName

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetEndUserLastName() string`

GetEndUserLastName returns the EndUserLastName field if non-nil, zero value otherwise.

### GetEndUserLastNameOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetEndUserLastNameOk() (*string, bool)`

GetEndUserLastNameOk returns a tuple with the EndUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserLastName

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetEndUserLastName(v string)`

SetEndUserLastName sets EndUserLastName field to given value.

### HasEndUserLastName

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasEndUserLastName() bool`

HasEndUserLastName returns a boolean if a field has been set.

### SetEndUserLastNameNil

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetEndUserLastNameNil(b bool)`

 SetEndUserLastNameNil sets the value for EndUserLastName to be an explicit nil

### UnsetEndUserLastName
`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) UnsetEndUserLastName()`

UnsetEndUserLastName ensures that no value is present for EndUserLastName, not even an explicit nil
### GetRegulatoryDeclineReason

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetRegulatoryDeclineReason() string`

GetRegulatoryDeclineReason returns the RegulatoryDeclineReason field if non-nil, zero value otherwise.

### GetRegulatoryDeclineReasonOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetRegulatoryDeclineReasonOk() (*string, bool)`

GetRegulatoryDeclineReasonOk returns a tuple with the RegulatoryDeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDeclineReason

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetRegulatoryDeclineReason(v string)`

SetRegulatoryDeclineReason sets RegulatoryDeclineReason field to given value.

### HasRegulatoryDeclineReason

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasRegulatoryDeclineReason() bool`

HasRegulatoryDeclineReason returns a boolean if a field has been set.

### SetRegulatoryDeclineReasonNil

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetRegulatoryDeclineReasonNil(b bool)`

 SetRegulatoryDeclineReasonNil sets the value for RegulatoryDeclineReason to be an explicit nil

### UnsetRegulatoryDeclineReason
`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) UnsetRegulatoryDeclineReason()`

UnsetRegulatoryDeclineReason ensures that no value is present for RegulatoryDeclineReason, not even an explicit nil
### GetProvisionedAt

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetProvisionedAt() time.Time`

GetProvisionedAt returns the ProvisionedAt field if non-nil, zero value otherwise.

### GetProvisionedAtOk

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) GetProvisionedAtOk() (*time.Time, bool)`

GetProvisionedAtOk returns a tuple with the ProvisionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedAt

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) SetProvisionedAt(v time.Time)`

SetProvisionedAt sets ProvisionedAt field to given value.

### HasProvisionedAt

`func (o *GetWhatsAppPhoneNumber200ResponsePhoneNumber) HasProvisionedAt() bool`

HasProvisionedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


