# WorkflowExecutionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**NodeId** | Pointer to **NullableString** | Present on &#x60;node_*&#x60; events | [optional] 
**NodeType** | Pointer to **NullableString** | Present on &#x60;node_*&#x60; events | [optional] 
**SourceHandle** | Pointer to **NullableString** | The edge handle the executor followed out of this node (see &#x60;WorkflowEdge.sourceHandle&#x60;) | [optional] 
**DurationMs** | Pointer to **NullableInt32** | Node run time; present on &#x60;node_completed&#x60; and &#x60;node_failed&#x60; | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Failure detail; present on &#x60;node_failed&#x60; and &#x60;execution_exited&#x60; | [optional] 
**Meta** | Pointer to **map[string]interface{}** | Per-node-type payload. Shape varies — see WorkflowNode &#x60;type&#x60;. Examples:   &#x60;send_message&#x60; → &#x60;{ messageType, text, recipient }&#x60;,   &#x60;webhook&#x60; → &#x60;{ url, method, statusCode, responseTimeMs, responsePreview }&#x60;,   &#x60;ai&#x60; → &#x60;{ model, provider, inputTokens, outputTokens, responsePreview }&#x60;,   &#x60;condition&#x60; → &#x60;{ matchedHandle, rulesEvaluated }&#x60;,   &#x60;a_b_split&#x60; → &#x60;{ percentage, chosen }&#x60;.  | [optional] 
**At** | Pointer to **time.Time** | Event timestamp (UTC) | [optional] 

## Methods

### NewWorkflowExecutionEvent

`func NewWorkflowExecutionEvent() *WorkflowExecutionEvent`

NewWorkflowExecutionEvent instantiates a new WorkflowExecutionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowExecutionEventWithDefaults

`func NewWorkflowExecutionEventWithDefaults() *WorkflowExecutionEvent`

NewWorkflowExecutionEventWithDefaults instantiates a new WorkflowExecutionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *WorkflowExecutionEvent) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *WorkflowExecutionEvent) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *WorkflowExecutionEvent) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *WorkflowExecutionEvent) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetStatus

`func (o *WorkflowExecutionEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WorkflowExecutionEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WorkflowExecutionEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WorkflowExecutionEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *WorkflowExecutionEvent) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *WorkflowExecutionEvent) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetNodeId

`func (o *WorkflowExecutionEvent) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *WorkflowExecutionEvent) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *WorkflowExecutionEvent) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *WorkflowExecutionEvent) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### SetNodeIdNil

`func (o *WorkflowExecutionEvent) SetNodeIdNil(b bool)`

 SetNodeIdNil sets the value for NodeId to be an explicit nil

### UnsetNodeId
`func (o *WorkflowExecutionEvent) UnsetNodeId()`

UnsetNodeId ensures that no value is present for NodeId, not even an explicit nil
### GetNodeType

`func (o *WorkflowExecutionEvent) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *WorkflowExecutionEvent) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *WorkflowExecutionEvent) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *WorkflowExecutionEvent) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### SetNodeTypeNil

`func (o *WorkflowExecutionEvent) SetNodeTypeNil(b bool)`

 SetNodeTypeNil sets the value for NodeType to be an explicit nil

### UnsetNodeType
`func (o *WorkflowExecutionEvent) UnsetNodeType()`

UnsetNodeType ensures that no value is present for NodeType, not even an explicit nil
### GetSourceHandle

`func (o *WorkflowExecutionEvent) GetSourceHandle() string`

GetSourceHandle returns the SourceHandle field if non-nil, zero value otherwise.

### GetSourceHandleOk

`func (o *WorkflowExecutionEvent) GetSourceHandleOk() (*string, bool)`

GetSourceHandleOk returns a tuple with the SourceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHandle

`func (o *WorkflowExecutionEvent) SetSourceHandle(v string)`

SetSourceHandle sets SourceHandle field to given value.

### HasSourceHandle

`func (o *WorkflowExecutionEvent) HasSourceHandle() bool`

HasSourceHandle returns a boolean if a field has been set.

### SetSourceHandleNil

`func (o *WorkflowExecutionEvent) SetSourceHandleNil(b bool)`

 SetSourceHandleNil sets the value for SourceHandle to be an explicit nil

### UnsetSourceHandle
`func (o *WorkflowExecutionEvent) UnsetSourceHandle()`

UnsetSourceHandle ensures that no value is present for SourceHandle, not even an explicit nil
### GetDurationMs

`func (o *WorkflowExecutionEvent) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *WorkflowExecutionEvent) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *WorkflowExecutionEvent) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *WorkflowExecutionEvent) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *WorkflowExecutionEvent) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *WorkflowExecutionEvent) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetErrorMessage

`func (o *WorkflowExecutionEvent) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *WorkflowExecutionEvent) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *WorkflowExecutionEvent) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *WorkflowExecutionEvent) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *WorkflowExecutionEvent) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *WorkflowExecutionEvent) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetMeta

`func (o *WorkflowExecutionEvent) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *WorkflowExecutionEvent) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *WorkflowExecutionEvent) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *WorkflowExecutionEvent) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *WorkflowExecutionEvent) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *WorkflowExecutionEvent) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetAt

`func (o *WorkflowExecutionEvent) GetAt() time.Time`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *WorkflowExecutionEvent) GetAtOk() (*time.Time, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *WorkflowExecutionEvent) SetAt(v time.Time)`

SetAt sets At field to given value.

### HasAt

`func (o *WorkflowExecutionEvent) HasAt() bool`

HasAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


