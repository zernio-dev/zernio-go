# WhatsAppTemplateComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Format** | **string** |  | 
**Text** | **string** | Static footer text | 
**Example** | Pointer to [**WhatsAppBodyComponentExample**](WhatsAppBodyComponentExample.md) |  | [optional] 
**AddSecurityRecommendation** | Pointer to **bool** | Add security recommendation text (authentication templates only) | [optional] 
**CodeExpirationMinutes** | Pointer to **int32** | OTP code expiry in minutes (authentication templates only) | [optional] 
**Buttons** | [**[]WhatsAppTemplateButton**](WhatsAppTemplateButton.md) |  | 

## Methods

### NewWhatsAppTemplateComponent

`func NewWhatsAppTemplateComponent(type_ string, format string, text string, buttons []WhatsAppTemplateButton, ) *WhatsAppTemplateComponent`

NewWhatsAppTemplateComponent instantiates a new WhatsAppTemplateComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppTemplateComponentWithDefaults

`func NewWhatsAppTemplateComponentWithDefaults() *WhatsAppTemplateComponent`

NewWhatsAppTemplateComponentWithDefaults instantiates a new WhatsAppTemplateComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsAppTemplateComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsAppTemplateComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsAppTemplateComponent) SetType(v string)`

SetType sets Type field to given value.


### GetFormat

`func (o *WhatsAppTemplateComponent) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *WhatsAppTemplateComponent) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *WhatsAppTemplateComponent) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetText

`func (o *WhatsAppTemplateComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsAppTemplateComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsAppTemplateComponent) SetText(v string)`

SetText sets Text field to given value.


### GetExample

`func (o *WhatsAppTemplateComponent) GetExample() WhatsAppBodyComponentExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsAppTemplateComponent) GetExampleOk() (*WhatsAppBodyComponentExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsAppTemplateComponent) SetExample(v WhatsAppBodyComponentExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsAppTemplateComponent) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetAddSecurityRecommendation

`func (o *WhatsAppTemplateComponent) GetAddSecurityRecommendation() bool`

GetAddSecurityRecommendation returns the AddSecurityRecommendation field if non-nil, zero value otherwise.

### GetAddSecurityRecommendationOk

`func (o *WhatsAppTemplateComponent) GetAddSecurityRecommendationOk() (*bool, bool)`

GetAddSecurityRecommendationOk returns a tuple with the AddSecurityRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSecurityRecommendation

`func (o *WhatsAppTemplateComponent) SetAddSecurityRecommendation(v bool)`

SetAddSecurityRecommendation sets AddSecurityRecommendation field to given value.

### HasAddSecurityRecommendation

`func (o *WhatsAppTemplateComponent) HasAddSecurityRecommendation() bool`

HasAddSecurityRecommendation returns a boolean if a field has been set.

### GetCodeExpirationMinutes

`func (o *WhatsAppTemplateComponent) GetCodeExpirationMinutes() int32`

GetCodeExpirationMinutes returns the CodeExpirationMinutes field if non-nil, zero value otherwise.

### GetCodeExpirationMinutesOk

`func (o *WhatsAppTemplateComponent) GetCodeExpirationMinutesOk() (*int32, bool)`

GetCodeExpirationMinutesOk returns a tuple with the CodeExpirationMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeExpirationMinutes

`func (o *WhatsAppTemplateComponent) SetCodeExpirationMinutes(v int32)`

SetCodeExpirationMinutes sets CodeExpirationMinutes field to given value.

### HasCodeExpirationMinutes

`func (o *WhatsAppTemplateComponent) HasCodeExpirationMinutes() bool`

HasCodeExpirationMinutes returns a boolean if a field has been set.

### GetButtons

`func (o *WhatsAppTemplateComponent) GetButtons() []WhatsAppTemplateButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *WhatsAppTemplateComponent) GetButtonsOk() (*[]WhatsAppTemplateButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *WhatsAppTemplateComponent) SetButtons(v []WhatsAppTemplateButton)`

SetButtons sets Buttons field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


