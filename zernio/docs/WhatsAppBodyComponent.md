# WhatsAppBodyComponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | **string** | Body text with optional {{n}} variables | 
**AddSecurityRecommendation** | Pointer to **bool** | Add security recommendation text (authentication templates only) | [optional] 
**Example** | Pointer to [**WhatsAppBodyComponentExample**](WhatsAppBodyComponentExample.md) |  | [optional] 

## Methods

### NewWhatsAppBodyComponent

`func NewWhatsAppBodyComponent(type_ string, text string, ) *WhatsAppBodyComponent`

NewWhatsAppBodyComponent instantiates a new WhatsAppBodyComponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppBodyComponentWithDefaults

`func NewWhatsAppBodyComponentWithDefaults() *WhatsAppBodyComponent`

NewWhatsAppBodyComponentWithDefaults instantiates a new WhatsAppBodyComponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsAppBodyComponent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsAppBodyComponent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsAppBodyComponent) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *WhatsAppBodyComponent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsAppBodyComponent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsAppBodyComponent) SetText(v string)`

SetText sets Text field to given value.


### GetAddSecurityRecommendation

`func (o *WhatsAppBodyComponent) GetAddSecurityRecommendation() bool`

GetAddSecurityRecommendation returns the AddSecurityRecommendation field if non-nil, zero value otherwise.

### GetAddSecurityRecommendationOk

`func (o *WhatsAppBodyComponent) GetAddSecurityRecommendationOk() (*bool, bool)`

GetAddSecurityRecommendationOk returns a tuple with the AddSecurityRecommendation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddSecurityRecommendation

`func (o *WhatsAppBodyComponent) SetAddSecurityRecommendation(v bool)`

SetAddSecurityRecommendation sets AddSecurityRecommendation field to given value.

### HasAddSecurityRecommendation

`func (o *WhatsAppBodyComponent) HasAddSecurityRecommendation() bool`

HasAddSecurityRecommendation returns a boolean if a field has been set.

### GetExample

`func (o *WhatsAppBodyComponent) GetExample() WhatsAppBodyComponentExample`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsAppBodyComponent) GetExampleOk() (*WhatsAppBodyComponentExample, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsAppBodyComponent) SetExample(v WhatsAppBodyComponentExample)`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsAppBodyComponent) HasExample() bool`

HasExample returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


