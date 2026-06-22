# GetWhatsAppPhoneNumbers200ResponseNumbersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RegistrantName** | Pointer to **NullableString** | For regulated numbers, who it&#39;s registered for (company or person) — set from the submitted KYC. | [optional] 
**TelnyxOrderId** | Pointer to **NullableString** | Present once the number order has been placed (i.e. the requirement group was approved). Absent while still in identity review. | [optional] 
**MonthlyCents** | Pointer to **int32** | Per-country monthly price in cents ($2..$25). | [optional] 
**ProfileId** | Pointer to **map[string]interface{}** |  | [optional] 
**ProvisionedAt** | Pointer to **time.Time** |  | [optional] 
**MetaPreverifiedId** | Pointer to **string** |  | [optional] 
**MetaVerificationStatus** | Pointer to **string** |  | [optional] 
**OnfidoVerificationUrl** | Pointer to **NullableString** | For regulated (Tier 3/4) numbers with an Onfido ID-verification step — the link to forward to the end user. Set once the order is placed; null otherwise. Poll this field after submitting KYC. | [optional] 
**EndUserFirstName** | Pointer to **NullableString** |  | [optional] 
**EndUserLastName** | Pointer to **NullableString** |  | [optional] 
**RegulatoryDeclineReason** | Pointer to **NullableString** | Reviewer rejection reason when status is regulatory_declined. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetWhatsAppPhoneNumbers200ResponseNumbersInner

`func NewGetWhatsAppPhoneNumbers200ResponseNumbersInner() *GetWhatsAppPhoneNumbers200ResponseNumbersInner`

NewGetWhatsAppPhoneNumbers200ResponseNumbersInner instantiates a new GetWhatsAppPhoneNumbers200ResponseNumbersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppPhoneNumbers200ResponseNumbersInnerWithDefaults

`func NewGetWhatsAppPhoneNumbers200ResponseNumbersInnerWithDefaults() *GetWhatsAppPhoneNumbers200ResponseNumbersInner`

NewGetWhatsAppPhoneNumbers200ResponseNumbersInnerWithDefaults instantiates a new GetWhatsAppPhoneNumbers200ResponseNumbersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountry

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetStatus

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRegistrantName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetRegistrantName() string`

GetRegistrantName returns the RegistrantName field if non-nil, zero value otherwise.

### GetRegistrantNameOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetRegistrantNameOk() (*string, bool)`

GetRegistrantNameOk returns a tuple with the RegistrantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrantName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetRegistrantName(v string)`

SetRegistrantName sets RegistrantName field to given value.

### HasRegistrantName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasRegistrantName() bool`

HasRegistrantName returns a boolean if a field has been set.

### SetRegistrantNameNil

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetRegistrantNameNil(b bool)`

 SetRegistrantNameNil sets the value for RegistrantName to be an explicit nil

### UnsetRegistrantName
`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) UnsetRegistrantName()`

UnsetRegistrantName ensures that no value is present for RegistrantName, not even an explicit nil
### GetTelnyxOrderId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetTelnyxOrderId() string`

GetTelnyxOrderId returns the TelnyxOrderId field if non-nil, zero value otherwise.

### GetTelnyxOrderIdOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetTelnyxOrderIdOk() (*string, bool)`

GetTelnyxOrderIdOk returns a tuple with the TelnyxOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxOrderId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetTelnyxOrderId(v string)`

SetTelnyxOrderId sets TelnyxOrderId field to given value.

### HasTelnyxOrderId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasTelnyxOrderId() bool`

HasTelnyxOrderId returns a boolean if a field has been set.

### SetTelnyxOrderIdNil

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetTelnyxOrderIdNil(b bool)`

 SetTelnyxOrderIdNil sets the value for TelnyxOrderId to be an explicit nil

### UnsetTelnyxOrderId
`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) UnsetTelnyxOrderId()`

UnsetTelnyxOrderId ensures that no value is present for TelnyxOrderId, not even an explicit nil
### GetMonthlyCents

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetMonthlyCents() int32`

GetMonthlyCents returns the MonthlyCents field if non-nil, zero value otherwise.

### GetMonthlyCentsOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetMonthlyCentsOk() (*int32, bool)`

GetMonthlyCentsOk returns a tuple with the MonthlyCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyCents

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetMonthlyCents(v int32)`

SetMonthlyCents sets MonthlyCents field to given value.

### HasMonthlyCents

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasMonthlyCents() bool`

HasMonthlyCents returns a boolean if a field has been set.

### GetProfileId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetProfileId() map[string]interface{}`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetProfileIdOk() (*map[string]interface{}, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetProfileId(v map[string]interface{})`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetProvisionedAt

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetProvisionedAt() time.Time`

GetProvisionedAt returns the ProvisionedAt field if non-nil, zero value otherwise.

### GetProvisionedAtOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetProvisionedAtOk() (*time.Time, bool)`

GetProvisionedAtOk returns a tuple with the ProvisionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionedAt

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetProvisionedAt(v time.Time)`

SetProvisionedAt sets ProvisionedAt field to given value.

### HasProvisionedAt

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasProvisionedAt() bool`

HasProvisionedAt returns a boolean if a field has been set.

### GetMetaPreverifiedId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetMetaPreverifiedId() string`

GetMetaPreverifiedId returns the MetaPreverifiedId field if non-nil, zero value otherwise.

### GetMetaPreverifiedIdOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetMetaPreverifiedIdOk() (*string, bool)`

GetMetaPreverifiedIdOk returns a tuple with the MetaPreverifiedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaPreverifiedId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetMetaPreverifiedId(v string)`

SetMetaPreverifiedId sets MetaPreverifiedId field to given value.

### HasMetaPreverifiedId

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasMetaPreverifiedId() bool`

HasMetaPreverifiedId returns a boolean if a field has been set.

### GetMetaVerificationStatus

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetMetaVerificationStatus() string`

GetMetaVerificationStatus returns the MetaVerificationStatus field if non-nil, zero value otherwise.

### GetMetaVerificationStatusOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetMetaVerificationStatusOk() (*string, bool)`

GetMetaVerificationStatusOk returns a tuple with the MetaVerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaVerificationStatus

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetMetaVerificationStatus(v string)`

SetMetaVerificationStatus sets MetaVerificationStatus field to given value.

### HasMetaVerificationStatus

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasMetaVerificationStatus() bool`

HasMetaVerificationStatus returns a boolean if a field has been set.

### GetOnfidoVerificationUrl

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetOnfidoVerificationUrl() string`

GetOnfidoVerificationUrl returns the OnfidoVerificationUrl field if non-nil, zero value otherwise.

### GetOnfidoVerificationUrlOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetOnfidoVerificationUrlOk() (*string, bool)`

GetOnfidoVerificationUrlOk returns a tuple with the OnfidoVerificationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnfidoVerificationUrl

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetOnfidoVerificationUrl(v string)`

SetOnfidoVerificationUrl sets OnfidoVerificationUrl field to given value.

### HasOnfidoVerificationUrl

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasOnfidoVerificationUrl() bool`

HasOnfidoVerificationUrl returns a boolean if a field has been set.

### SetOnfidoVerificationUrlNil

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetOnfidoVerificationUrlNil(b bool)`

 SetOnfidoVerificationUrlNil sets the value for OnfidoVerificationUrl to be an explicit nil

### UnsetOnfidoVerificationUrl
`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) UnsetOnfidoVerificationUrl()`

UnsetOnfidoVerificationUrl ensures that no value is present for OnfidoVerificationUrl, not even an explicit nil
### GetEndUserFirstName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetEndUserFirstName() string`

GetEndUserFirstName returns the EndUserFirstName field if non-nil, zero value otherwise.

### GetEndUserFirstNameOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetEndUserFirstNameOk() (*string, bool)`

GetEndUserFirstNameOk returns a tuple with the EndUserFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserFirstName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetEndUserFirstName(v string)`

SetEndUserFirstName sets EndUserFirstName field to given value.

### HasEndUserFirstName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasEndUserFirstName() bool`

HasEndUserFirstName returns a boolean if a field has been set.

### SetEndUserFirstNameNil

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetEndUserFirstNameNil(b bool)`

 SetEndUserFirstNameNil sets the value for EndUserFirstName to be an explicit nil

### UnsetEndUserFirstName
`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) UnsetEndUserFirstName()`

UnsetEndUserFirstName ensures that no value is present for EndUserFirstName, not even an explicit nil
### GetEndUserLastName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetEndUserLastName() string`

GetEndUserLastName returns the EndUserLastName field if non-nil, zero value otherwise.

### GetEndUserLastNameOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetEndUserLastNameOk() (*string, bool)`

GetEndUserLastNameOk returns a tuple with the EndUserLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserLastName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetEndUserLastName(v string)`

SetEndUserLastName sets EndUserLastName field to given value.

### HasEndUserLastName

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasEndUserLastName() bool`

HasEndUserLastName returns a boolean if a field has been set.

### SetEndUserLastNameNil

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetEndUserLastNameNil(b bool)`

 SetEndUserLastNameNil sets the value for EndUserLastName to be an explicit nil

### UnsetEndUserLastName
`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) UnsetEndUserLastName()`

UnsetEndUserLastName ensures that no value is present for EndUserLastName, not even an explicit nil
### GetRegulatoryDeclineReason

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetRegulatoryDeclineReason() string`

GetRegulatoryDeclineReason returns the RegulatoryDeclineReason field if non-nil, zero value otherwise.

### GetRegulatoryDeclineReasonOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetRegulatoryDeclineReasonOk() (*string, bool)`

GetRegulatoryDeclineReasonOk returns a tuple with the RegulatoryDeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegulatoryDeclineReason

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetRegulatoryDeclineReason(v string)`

SetRegulatoryDeclineReason sets RegulatoryDeclineReason field to given value.

### HasRegulatoryDeclineReason

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasRegulatoryDeclineReason() bool`

HasRegulatoryDeclineReason returns a boolean if a field has been set.

### SetRegulatoryDeclineReasonNil

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetRegulatoryDeclineReasonNil(b bool)`

 SetRegulatoryDeclineReasonNil sets the value for RegulatoryDeclineReason to be an explicit nil

### UnsetRegulatoryDeclineReason
`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) UnsetRegulatoryDeclineReason()`

UnsetRegulatoryDeclineReason ensures that no value is present for RegulatoryDeclineReason, not even an explicit nil
### GetCreatedAt

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetWhatsAppPhoneNumbers200ResponseNumbersInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


