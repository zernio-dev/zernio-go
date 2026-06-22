# ListWhatsAppFlows200ResponseFlowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 
**ValidationErrors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Version** | Pointer to **int32** | 1-based version within the flow&#39;s clone lineage (Zernio-tracked; Meta has no native versioning). Standalone flows are version 1. | [optional] 
**LineageId** | Pointer to **string** | Stable group key for the flow&#39;s version lineage (the root flow&#39;s ID). | [optional] 

## Methods

### NewListWhatsAppFlows200ResponseFlowsInner

`func NewListWhatsAppFlows200ResponseFlowsInner() *ListWhatsAppFlows200ResponseFlowsInner`

NewListWhatsAppFlows200ResponseFlowsInner instantiates a new ListWhatsAppFlows200ResponseFlowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppFlows200ResponseFlowsInnerWithDefaults

`func NewListWhatsAppFlows200ResponseFlowsInnerWithDefaults() *ListWhatsAppFlows200ResponseFlowsInner`

NewListWhatsAppFlows200ResponseFlowsInnerWithDefaults instantiates a new ListWhatsAppFlows200ResponseFlowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCategories

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetValidationErrors

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetValidationErrors() []map[string]interface{}`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetValidationErrorsOk() (*[]map[string]interface{}, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetValidationErrors(v []map[string]interface{})`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.

### GetVersion

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLineageId

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetLineageId() string`

GetLineageId returns the LineageId field if non-nil, zero value otherwise.

### GetLineageIdOk

`func (o *ListWhatsAppFlows200ResponseFlowsInner) GetLineageIdOk() (*string, bool)`

GetLineageIdOk returns a tuple with the LineageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineageId

`func (o *ListWhatsAppFlows200ResponseFlowsInner) SetLineageId(v string)`

SetLineageId sets LineageId field to given value.

### HasLineageId

`func (o *ListWhatsAppFlows200ResponseFlowsInner) HasLineageId() bool`

HasLineageId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


