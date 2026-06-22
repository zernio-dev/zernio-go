# CreateBroadcastRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 
**Components** | Pointer to **[]interface{}** |  | [optional] 
**VariableMapping** | Pointer to [**map[string]CreateBroadcastRequestTemplateVariableMappingValue**](CreateBroadcastRequestTemplateVariableMappingValue.md) | Maps template variable positions (\&quot;1\&quot;, \&quot;2\&quot;) to contact fields or static values. Resolved per recipient at send time. | [optional] 

## Methods

### NewCreateBroadcastRequestTemplate

`func NewCreateBroadcastRequestTemplate() *CreateBroadcastRequestTemplate`

NewCreateBroadcastRequestTemplate instantiates a new CreateBroadcastRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBroadcastRequestTemplateWithDefaults

`func NewCreateBroadcastRequestTemplateWithDefaults() *CreateBroadcastRequestTemplate`

NewCreateBroadcastRequestTemplateWithDefaults instantiates a new CreateBroadcastRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateBroadcastRequestTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBroadcastRequestTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBroadcastRequestTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBroadcastRequestTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLanguage

`func (o *CreateBroadcastRequestTemplate) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateBroadcastRequestTemplate) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateBroadcastRequestTemplate) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *CreateBroadcastRequestTemplate) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetComponents

`func (o *CreateBroadcastRequestTemplate) GetComponents() []interface{}`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CreateBroadcastRequestTemplate) GetComponentsOk() (*[]interface{}, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CreateBroadcastRequestTemplate) SetComponents(v []interface{})`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CreateBroadcastRequestTemplate) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetVariableMapping

`func (o *CreateBroadcastRequestTemplate) GetVariableMapping() map[string]CreateBroadcastRequestTemplateVariableMappingValue`

GetVariableMapping returns the VariableMapping field if non-nil, zero value otherwise.

### GetVariableMappingOk

`func (o *CreateBroadcastRequestTemplate) GetVariableMappingOk() (*map[string]CreateBroadcastRequestTemplateVariableMappingValue, bool)`

GetVariableMappingOk returns a tuple with the VariableMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableMapping

`func (o *CreateBroadcastRequestTemplate) SetVariableMapping(v map[string]CreateBroadcastRequestTemplateVariableMappingValue)`

SetVariableMapping sets VariableMapping field to given value.

### HasVariableMapping

`func (o *CreateBroadcastRequestTemplate) HasVariableMapping() bool`

HasVariableMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


