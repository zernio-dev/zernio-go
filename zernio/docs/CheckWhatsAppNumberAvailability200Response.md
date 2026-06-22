# CheckWhatsAppNumberAvailability200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** |  | [optional] 
**NumberType** | Pointer to **string** |  | [optional] 
**Available** | Pointer to **bool** | Whether deliverable voice inventory exists right now. | [optional] 
**AddressConstraint** | Pointer to **string** |  | [optional] 
**Areas** | Pointer to **[]string** | For &#x60;geo&#x60; only — the area(s) the registered address must be in. | [optional] 

## Methods

### NewCheckWhatsAppNumberAvailability200Response

`func NewCheckWhatsAppNumberAvailability200Response() *CheckWhatsAppNumberAvailability200Response`

NewCheckWhatsAppNumberAvailability200Response instantiates a new CheckWhatsAppNumberAvailability200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckWhatsAppNumberAvailability200ResponseWithDefaults

`func NewCheckWhatsAppNumberAvailability200ResponseWithDefaults() *CheckWhatsAppNumberAvailability200Response`

NewCheckWhatsAppNumberAvailability200ResponseWithDefaults instantiates a new CheckWhatsAppNumberAvailability200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *CheckWhatsAppNumberAvailability200Response) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CheckWhatsAppNumberAvailability200Response) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CheckWhatsAppNumberAvailability200Response) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CheckWhatsAppNumberAvailability200Response) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNumberType

`func (o *CheckWhatsAppNumberAvailability200Response) GetNumberType() string`

GetNumberType returns the NumberType field if non-nil, zero value otherwise.

### GetNumberTypeOk

`func (o *CheckWhatsAppNumberAvailability200Response) GetNumberTypeOk() (*string, bool)`

GetNumberTypeOk returns a tuple with the NumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberType

`func (o *CheckWhatsAppNumberAvailability200Response) SetNumberType(v string)`

SetNumberType sets NumberType field to given value.

### HasNumberType

`func (o *CheckWhatsAppNumberAvailability200Response) HasNumberType() bool`

HasNumberType returns a boolean if a field has been set.

### GetAvailable

`func (o *CheckWhatsAppNumberAvailability200Response) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CheckWhatsAppNumberAvailability200Response) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CheckWhatsAppNumberAvailability200Response) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CheckWhatsAppNumberAvailability200Response) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetAddressConstraint

`func (o *CheckWhatsAppNumberAvailability200Response) GetAddressConstraint() string`

GetAddressConstraint returns the AddressConstraint field if non-nil, zero value otherwise.

### GetAddressConstraintOk

`func (o *CheckWhatsAppNumberAvailability200Response) GetAddressConstraintOk() (*string, bool)`

GetAddressConstraintOk returns a tuple with the AddressConstraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressConstraint

`func (o *CheckWhatsAppNumberAvailability200Response) SetAddressConstraint(v string)`

SetAddressConstraint sets AddressConstraint field to given value.

### HasAddressConstraint

`func (o *CheckWhatsAppNumberAvailability200Response) HasAddressConstraint() bool`

HasAddressConstraint returns a boolean if a field has been set.

### GetAreas

`func (o *CheckWhatsAppNumberAvailability200Response) GetAreas() []string`

GetAreas returns the Areas field if non-nil, zero value otherwise.

### GetAreasOk

`func (o *CheckWhatsAppNumberAvailability200Response) GetAreasOk() (*[]string, bool)`

GetAreasOk returns a tuple with the Areas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreas

`func (o *CheckWhatsAppNumberAvailability200Response) SetAreas(v []string)`

SetAreas sets Areas field to given value.

### HasAreas

`func (o *CheckWhatsAppNumberAvailability200Response) HasAreas() bool`

HasAreas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


