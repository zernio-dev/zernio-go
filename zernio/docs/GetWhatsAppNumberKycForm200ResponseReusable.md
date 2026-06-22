# GetWhatsAppNumberKycForm200ResponseReusable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** |  | [optional] 
**FromPhoneNumber** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]GetWhatsAppNumberKycForm200ResponseReusableDetailsInner**](GetWhatsAppNumberKycForm200ResponseReusableDetailsInner.md) | Human-readable summary of the verification on file (field labels + values, plus the address as one line). Best-effort — may be empty if the provider lookup fails. | [optional] 
**Options** | Pointer to [**[]GetWhatsAppNumberKycForm200ResponseReusableOptionsInner**](GetWhatsAppNumberKycForm200ResponseReusableOptionsInner.md) | One entry per distinct approved verification, newest first. | [optional] 

## Methods

### NewGetWhatsAppNumberKycForm200ResponseReusable

`func NewGetWhatsAppNumberKycForm200ResponseReusable() *GetWhatsAppNumberKycForm200ResponseReusable`

NewGetWhatsAppNumberKycForm200ResponseReusable instantiates a new GetWhatsAppNumberKycForm200ResponseReusable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppNumberKycForm200ResponseReusableWithDefaults

`func NewGetWhatsAppNumberKycForm200ResponseReusableWithDefaults() *GetWhatsAppNumberKycForm200ResponseReusable`

NewGetWhatsAppNumberKycForm200ResponseReusableWithDefaults instantiates a new GetWhatsAppNumberKycForm200ResponseReusable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetFromPhoneNumber

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetFromPhoneNumber() string`

GetFromPhoneNumber returns the FromPhoneNumber field if non-nil, zero value otherwise.

### GetFromPhoneNumberOk

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetFromPhoneNumberOk() (*string, bool)`

GetFromPhoneNumberOk returns a tuple with the FromPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPhoneNumber

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) SetFromPhoneNumber(v string)`

SetFromPhoneNumber sets FromPhoneNumber field to given value.

### HasFromPhoneNumber

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) HasFromPhoneNumber() bool`

HasFromPhoneNumber returns a boolean if a field has been set.

### GetDetails

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetDetails() []GetWhatsAppNumberKycForm200ResponseReusableDetailsInner`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetDetailsOk() (*[]GetWhatsAppNumberKycForm200ResponseReusableDetailsInner, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) SetDetails(v []GetWhatsAppNumberKycForm200ResponseReusableDetailsInner)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetOptions

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetOptions() []GetWhatsAppNumberKycForm200ResponseReusableOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) GetOptionsOk() (*[]GetWhatsAppNumberKycForm200ResponseReusableOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) SetOptions(v []GetWhatsAppNumberKycForm200ResponseReusableOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetWhatsAppNumberKycForm200ResponseReusable) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


