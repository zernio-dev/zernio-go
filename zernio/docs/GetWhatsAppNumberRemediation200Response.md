# GetWhatsAppNumberRemediation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **string** |  | [optional] 
**NumberType** | Pointer to **string** |  | [optional] 
**DeclineReason** | Pointer to **NullableString** |  | [optional] 
**Fields** | Pointer to **[]map[string]interface{}** | Same field shape as GET /v1/whatsapp/phone-numbers/kyc. | [optional] 

## Methods

### NewGetWhatsAppNumberRemediation200Response

`func NewGetWhatsAppNumberRemediation200Response() *GetWhatsAppNumberRemediation200Response`

NewGetWhatsAppNumberRemediation200Response instantiates a new GetWhatsAppNumberRemediation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppNumberRemediation200ResponseWithDefaults

`func NewGetWhatsAppNumberRemediation200ResponseWithDefaults() *GetWhatsAppNumberRemediation200Response`

NewGetWhatsAppNumberRemediation200ResponseWithDefaults instantiates a new GetWhatsAppNumberRemediation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *GetWhatsAppNumberRemediation200Response) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GetWhatsAppNumberRemediation200Response) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GetWhatsAppNumberRemediation200Response) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GetWhatsAppNumberRemediation200Response) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetNumberType

`func (o *GetWhatsAppNumberRemediation200Response) GetNumberType() string`

GetNumberType returns the NumberType field if non-nil, zero value otherwise.

### GetNumberTypeOk

`func (o *GetWhatsAppNumberRemediation200Response) GetNumberTypeOk() (*string, bool)`

GetNumberTypeOk returns a tuple with the NumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberType

`func (o *GetWhatsAppNumberRemediation200Response) SetNumberType(v string)`

SetNumberType sets NumberType field to given value.

### HasNumberType

`func (o *GetWhatsAppNumberRemediation200Response) HasNumberType() bool`

HasNumberType returns a boolean if a field has been set.

### GetDeclineReason

`func (o *GetWhatsAppNumberRemediation200Response) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *GetWhatsAppNumberRemediation200Response) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *GetWhatsAppNumberRemediation200Response) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *GetWhatsAppNumberRemediation200Response) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### SetDeclineReasonNil

`func (o *GetWhatsAppNumberRemediation200Response) SetDeclineReasonNil(b bool)`

 SetDeclineReasonNil sets the value for DeclineReason to be an explicit nil

### UnsetDeclineReason
`func (o *GetWhatsAppNumberRemediation200Response) UnsetDeclineReason()`

UnsetDeclineReason ensures that no value is present for DeclineReason, not even an explicit nil
### GetFields

`func (o *GetWhatsAppNumberRemediation200Response) GetFields() []map[string]interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetWhatsAppNumberRemediation200Response) GetFieldsOk() (*[]map[string]interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetWhatsAppNumberRemediation200Response) SetFields(v []map[string]interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *GetWhatsAppNumberRemediation200Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


