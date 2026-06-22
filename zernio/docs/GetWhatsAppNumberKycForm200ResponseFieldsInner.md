# GetWhatsAppNumberKycForm200ResponseFieldsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequirementId** | Pointer to **string** |  | [optional] 
**Label** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** | \&quot;action\&quot; &#x3D; an out-of-band verification (e.g. Onfido); not filled here, fulfilled after the order via a link. | [optional] 
**Description** | Pointer to **NullableString** | Plain-English explanation of what to provide. | [optional] 
**Example** | Pointer to **NullableString** | Concrete example value. | [optional] 
**LocalTo** | Pointer to **NullableString** | ISO country the value must be local to | [optional] 

## Methods

### NewGetWhatsAppNumberKycForm200ResponseFieldsInner

`func NewGetWhatsAppNumberKycForm200ResponseFieldsInner() *GetWhatsAppNumberKycForm200ResponseFieldsInner`

NewGetWhatsAppNumberKycForm200ResponseFieldsInner instantiates a new GetWhatsAppNumberKycForm200ResponseFieldsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppNumberKycForm200ResponseFieldsInnerWithDefaults

`func NewGetWhatsAppNumberKycForm200ResponseFieldsInnerWithDefaults() *GetWhatsAppNumberKycForm200ResponseFieldsInner`

NewGetWhatsAppNumberKycForm200ResponseFieldsInnerWithDefaults instantiates a new GetWhatsAppNumberKycForm200ResponseFieldsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequirementId

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetRequirementId() string`

GetRequirementId returns the RequirementId field if non-nil, zero value otherwise.

### GetRequirementIdOk

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetRequirementIdOk() (*string, bool)`

GetRequirementIdOk returns a tuple with the RequirementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementId

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetRequirementId(v string)`

SetRequirementId sets RequirementId field to given value.

### HasRequirementId

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) HasRequirementId() bool`

HasRequirementId returns a boolean if a field has been set.

### GetLabel

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetKind

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetDescription

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetExample

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) HasExample() bool`

HasExample returns a boolean if a field has been set.

### SetExampleNil

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetExampleNil(b bool)`

 SetExampleNil sets the value for Example to be an explicit nil

### UnsetExample
`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) UnsetExample()`

UnsetExample ensures that no value is present for Example, not even an explicit nil
### GetLocalTo

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetLocalTo() string`

GetLocalTo returns the LocalTo field if non-nil, zero value otherwise.

### GetLocalToOk

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) GetLocalToOk() (*string, bool)`

GetLocalToOk returns a tuple with the LocalTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalTo

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetLocalTo(v string)`

SetLocalTo sets LocalTo field to given value.

### HasLocalTo

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) HasLocalTo() bool`

HasLocalTo returns a boolean if a field has been set.

### SetLocalToNil

`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) SetLocalToNil(b bool)`

 SetLocalToNil sets the value for LocalTo to be an explicit nil

### UnsetLocalTo
`func (o *GetWhatsAppNumberKycForm200ResponseFieldsInner) UnsetLocalTo()`

UnsetLocalTo ensures that no value is present for LocalTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


