# WhatsAppFooterComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | Pointer to **string** | Static footer text | [optional] 
**CodeExpirationMinutes** | Pointer to **int32** | OTP code expiry in minutes (authentication templates only) | [optional] 

## Methods

### NewWhatsAppFooterComponent

`func NewWhatsAppFooterComponent(type_ string, ) *WhatsAppFooterComponent`

NewWhatsAppFooterComponent instantiates a new WhatsAppFooterComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppFooterComponentWithDefaults

`func NewWhatsAppFooterComponentWithDefaults() *WhatsAppFooterComponent`

NewWhatsAppFooterComponentWithDefaults instantiates a new WhatsAppFooterComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsAppFooterComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsAppFooterComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsAppFooterComponent) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *WhatsAppFooterComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsAppFooterComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsAppFooterComponent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsAppFooterComponent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetCodeExpirationMinutes

`func (o *WhatsAppFooterComponent) GetCodeExpirationMinutes() int32`

GetCodeExpirationMinutes returns the CodeExpirationMinutes field if non-nil, zero value otherwise.

### GetCodeExpirationMinutesOk

`func (o *WhatsAppFooterComponent) GetCodeExpirationMinutesOk() (*int32, bool)`

GetCodeExpirationMinutesOk returns a tuple with the CodeExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExpirationMinutes

`func (o *WhatsAppFooterComponent) SetCodeExpirationMinutes(v int32)`

SetCodeExpirationMinutes sets CodeExpirationMinutes field to given value.

### HasCodeExpirationMinutes

`func (o *WhatsAppFooterComponent) HasCodeExpirationMinutes() bool`

HasCodeExpirationMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


