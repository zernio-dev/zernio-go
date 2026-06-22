# WhatsAppTemplateButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Text** | Pointer to **string** | Visible button label. Required for all types except copy_code (whose label is fixed by WhatsApp). | [optional] 
**Url** | Pointer to **string** | Required when type is URL | [optional] 
**Example** | Pointer to **interface{}** |  | [optional] 
**PhoneNumber** | Pointer to **string** | Required when type is phone_number | [optional] 
**OtpType** | Pointer to **string** | Required when type is otp | [optional] 
**AutofillText** | Pointer to **string** |  | [optional] 
**PackageName** | Pointer to **string** |  | [optional] 
**SignatureHash** | Pointer to **string** |  | [optional] 
**FlowId** | Pointer to **string** |  | [optional] 
**FlowName** | Pointer to **string** |  | [optional] 
**FlowJson** | Pointer to **string** |  | [optional] 
**FlowAction** | Pointer to **string** |  | [optional] 
**NavigateScreen** | Pointer to **string** |  | [optional] 

## Methods

### NewWhatsAppTemplateButton

`func NewWhatsAppTemplateButton(type_ string, ) *WhatsAppTemplateButton`

NewWhatsAppTemplateButton instantiates a new WhatsAppTemplateButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppTemplateButtonWithDefaults

`func NewWhatsAppTemplateButtonWithDefaults() *WhatsAppTemplateButton`

NewWhatsAppTemplateButtonWithDefaults instantiates a new WhatsAppTemplateButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *WhatsAppTemplateButton) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WhatsAppTemplateButton) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WhatsAppTemplateButton) SetType(v string)`

SetType sets Type field to given value.


### GetText

`func (o *WhatsAppTemplateButton) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WhatsAppTemplateButton) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WhatsAppTemplateButton) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *WhatsAppTemplateButton) HasText() bool`

HasText returns a boolean if a field has been set.

### GetUrl

`func (o *WhatsAppTemplateButton) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WhatsAppTemplateButton) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WhatsAppTemplateButton) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *WhatsAppTemplateButton) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetExample

`func (o *WhatsAppTemplateButton) GetExample() interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *WhatsAppTemplateButton) GetExampleOk() (*interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *WhatsAppTemplateButton) SetExample(v interface{})`

SetExample sets Example field to given value.

### HasExample

`func (o *WhatsAppTemplateButton) HasExample() bool`

HasExample returns a boolean if a field has been set.

### SetExampleNil

`func (o *WhatsAppTemplateButton) SetExampleNil(b bool)`

 SetExampleNil sets the value for Example to be an explicit nil

### UnsetExample
`func (o *WhatsAppTemplateButton) UnsetExample()`

UnsetExample ensures that no value is present for Example, not even an explicit nil
### GetPhoneNumber

`func (o *WhatsAppTemplateButton) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WhatsAppTemplateButton) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WhatsAppTemplateButton) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WhatsAppTemplateButton) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetOtpType

`func (o *WhatsAppTemplateButton) GetOtpType() string`

GetOtpType returns the OtpType field if non-nil, zero value otherwise.

### GetOtpTypeOk

`func (o *WhatsAppTemplateButton) GetOtpTypeOk() (*string, bool)`

GetOtpTypeOk returns a tuple with the OtpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpType

`func (o *WhatsAppTemplateButton) SetOtpType(v string)`

SetOtpType sets OtpType field to given value.

### HasOtpType

`func (o *WhatsAppTemplateButton) HasOtpType() bool`

HasOtpType returns a boolean if a field has been set.

### GetAutofillText

`func (o *WhatsAppTemplateButton) GetAutofillText() string`

GetAutofillText returns the AutofillText field if non-nil, zero value otherwise.

### GetAutofillTextOk

`func (o *WhatsAppTemplateButton) GetAutofillTextOk() (*string, bool)`

GetAutofillTextOk returns a tuple with the AutofillText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutofillText

`func (o *WhatsAppTemplateButton) SetAutofillText(v string)`

SetAutofillText sets AutofillText field to given value.

### HasAutofillText

`func (o *WhatsAppTemplateButton) HasAutofillText() bool`

HasAutofillText returns a boolean if a field has been set.

### GetPackageName

`func (o *WhatsAppTemplateButton) GetPackageName() string`

GetPackageName returns the PackageName field if non-nil, zero value otherwise.

### GetPackageNameOk

`func (o *WhatsAppTemplateButton) GetPackageNameOk() (*string, bool)`

GetPackageNameOk returns a tuple with the PackageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageName

`func (o *WhatsAppTemplateButton) SetPackageName(v string)`

SetPackageName sets PackageName field to given value.

### HasPackageName

`func (o *WhatsAppTemplateButton) HasPackageName() bool`

HasPackageName returns a boolean if a field has been set.

### GetSignatureHash

`func (o *WhatsAppTemplateButton) GetSignatureHash() string`

GetSignatureHash returns the SignatureHash field if non-nil, zero value otherwise.

### GetSignatureHashOk

`func (o *WhatsAppTemplateButton) GetSignatureHashOk() (*string, bool)`

GetSignatureHashOk returns a tuple with the SignatureHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureHash

`func (o *WhatsAppTemplateButton) SetSignatureHash(v string)`

SetSignatureHash sets SignatureHash field to given value.

### HasSignatureHash

`func (o *WhatsAppTemplateButton) HasSignatureHash() bool`

HasSignatureHash returns a boolean if a field has been set.

### GetFlowId

`func (o *WhatsAppTemplateButton) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *WhatsAppTemplateButton) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *WhatsAppTemplateButton) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.

### HasFlowId

`func (o *WhatsAppTemplateButton) HasFlowId() bool`

HasFlowId returns a boolean if a field has been set.

### GetFlowName

`func (o *WhatsAppTemplateButton) GetFlowName() string`

GetFlowName returns the FlowName field if non-nil, zero value otherwise.

### GetFlowNameOk

`func (o *WhatsAppTemplateButton) GetFlowNameOk() (*string, bool)`

GetFlowNameOk returns a tuple with the FlowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowName

`func (o *WhatsAppTemplateButton) SetFlowName(v string)`

SetFlowName sets FlowName field to given value.

### HasFlowName

`func (o *WhatsAppTemplateButton) HasFlowName() bool`

HasFlowName returns a boolean if a field has been set.

### GetFlowJson

`func (o *WhatsAppTemplateButton) GetFlowJson() string`

GetFlowJson returns the FlowJson field if non-nil, zero value otherwise.

### GetFlowJsonOk

`func (o *WhatsAppTemplateButton) GetFlowJsonOk() (*string, bool)`

GetFlowJsonOk returns a tuple with the FlowJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowJson

`func (o *WhatsAppTemplateButton) SetFlowJson(v string)`

SetFlowJson sets FlowJson field to given value.

### HasFlowJson

`func (o *WhatsAppTemplateButton) HasFlowJson() bool`

HasFlowJson returns a boolean if a field has been set.

### GetFlowAction

`func (o *WhatsAppTemplateButton) GetFlowAction() string`

GetFlowAction returns the FlowAction field if non-nil, zero value otherwise.

### GetFlowActionOk

`func (o *WhatsAppTemplateButton) GetFlowActionOk() (*string, bool)`

GetFlowActionOk returns a tuple with the FlowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAction

`func (o *WhatsAppTemplateButton) SetFlowAction(v string)`

SetFlowAction sets FlowAction field to given value.

### HasFlowAction

`func (o *WhatsAppTemplateButton) HasFlowAction() bool`

HasFlowAction returns a boolean if a field has been set.

### GetNavigateScreen

`func (o *WhatsAppTemplateButton) GetNavigateScreen() string`

GetNavigateScreen returns the NavigateScreen field if non-nil, zero value otherwise.

### GetNavigateScreenOk

`func (o *WhatsAppTemplateButton) GetNavigateScreenOk() (*string, bool)`

GetNavigateScreenOk returns a tuple with the NavigateScreen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigateScreen

`func (o *WhatsAppTemplateButton) SetNavigateScreen(v string)`

SetNavigateScreen sets NavigateScreen field to given value.

### HasNavigateScreen

`func (o *WhatsAppTemplateButton) HasNavigateScreen() bool`

HasNavigateScreen returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


