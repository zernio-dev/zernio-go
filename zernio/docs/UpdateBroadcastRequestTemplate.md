# UpdateBroadcastRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**VariableMapping** | Pointer to [**map[string]CreateBroadcastRequestTemplateVariableMappingValue**](CreateBroadcastRequestTemplateVariableMappingValue.md) | Maps template variable positions to contact fields. Keys are position strings (\&quot;1\&quot;, \&quot;2\&quot;); values are { field, customValue }. | [optional] 

## Methods

### NewUpdateBroadcastRequestTemplate

`func NewUpdateBroadcastRequestTemplate() *UpdateBroadcastRequestTemplate`

NewUpdateBroadcastRequestTemplate instantiates a new UpdateBroadcastRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBroadcastRequestTemplateWithDefaults

`func NewUpdateBroadcastRequestTemplateWithDefaults() *UpdateBroadcastRequestTemplate`

NewUpdateBroadcastRequestTemplateWithDefaults instantiates a new UpdateBroadcastRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBroadcastRequestTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBroadcastRequestTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBroadcastRequestTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateBroadcastRequestTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *UpdateBroadcastRequestTemplate) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UpdateBroadcastRequestTemplate) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UpdateBroadcastRequestTemplate) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UpdateBroadcastRequestTemplate) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetVariableMapping

`func (o *UpdateBroadcastRequestTemplate) GetVariableMapping() map[string]CreateBroadcastRequestTemplateVariableMappingValue`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *UpdateBroadcastRequestTemplate) GetVariableMappingOk() (*map[string]CreateBroadcastRequestTemplateVariableMappingValue, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *UpdateBroadcastRequestTemplate) SetVariableMapping(v map[string]CreateBroadcastRequestTemplateVariableMappingValue)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *UpdateBroadcastRequestTemplate) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


