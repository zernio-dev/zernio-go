# WhatsAppHeaderComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Format** | **string** |  | 
**Text** | Pointer to **string** | Header text (may include {{1}} variable). Used when format is TEXT. | [optional] 
**Example** | Pointer to [**WhatsAppHeaderComponentExample**](WhatsAppHeaderComponentExample.md) |  | [optional] 

## Methods

### NewWhatsAppHeaderComponent

`func NewWhatsAppHeaderComponent(type_ string, format string, ) *WhatsAppHeaderComponent`

NewWhatsAppHeaderComponent instantiates a new WhatsAppHeaderComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppHeaderComponentWithDefaults

`func NewWhatsAppHeaderComponentWithDefaults() *WhatsAppHeaderComponent`

NewWhatsAppHeaderComponentWithDefaults instantiates a new WhatsAppHeaderComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsAppHeaderComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsAppHeaderComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsAppHeaderComponent) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *WhatsAppHeaderComponent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WhatsAppHeaderComponent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WhatsAppHeaderComponent) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetText

`func (o *WhatsAppHeaderComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsAppHeaderComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsAppHeaderComponent) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsAppHeaderComponent) HasText() bool`

HasText returns a boolean if a field has been set.

### GetExample

`func (o *WhatsAppHeaderComponent) GetExample() WhatsAppHeaderComponentExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsAppHeaderComponent) GetExampleOk() (*WhatsAppHeaderComponentExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsAppHeaderComponent) SetExample(v WhatsAppHeaderComponentExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsAppHeaderComponent) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


