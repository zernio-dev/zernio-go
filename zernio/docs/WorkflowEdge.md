# WorkflowEdge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Source** | **string** | Source node id | 
**Target** | **string** | Target node id | 
**SourceHandle** | Pointer to **NullableString** | Selects a branch output of a multi-output node. Null (or omitted) &#x3D; the node&#39;s single/default output. Known handles per node type:    - **condition** — a rule&#39;s &#x60;id&#x60;, or &#x60;&#39;default&#39;&#x60; (no rule matched)   - **wait_for_reply** — &#x60;&#39;reply&#39;&#x60; (contact replied) | &#x60;&#39;timeout&#39;&#x60; (no reply in window)   - **webhook** — &#x60;&#39;success&#39;&#x60; (2xx) | &#x60;&#39;error&#39;&#x60; (non-2xx / fetch failed)   - **ai** — &#x60;&#39;success&#39;&#x60; (text/JSON response) | &#x60;&#39;tool:&lt;toolName&gt;&#39;&#x60; (model invoked     that tool) | &#x60;&#39;error&#39;&#x60; (upstream failure / non-JSON in JSON mode)   - **start_call** — &#x60;&#39;success&#39;&#x60; | &#x60;&#39;permission_required&#39;&#x60; | &#x60;&#39;failed&#39;&#x60;   - **a_b_split** — &#x60;&#39;a&#39;&#x60; | &#x60;&#39;b&#39;&#x60;   - **enroll_sequence** — &#x60;&#39;success&#39;&#x60; | &#x60;&#39;error&#39;&#x60;  | [optional] 

## Methods

### NewWorkflowEdge

`func NewWorkflowEdge(id string, source string, target string, ) *WorkflowEdge`

NewWorkflowEdge instantiates a new WorkflowEdge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowEdgeWithDefaults

`func NewWorkflowEdgeWithDefaults() *WorkflowEdge`

NewWorkflowEdgeWithDefaults instantiates a new WorkflowEdge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowEdge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowEdge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowEdge) SetId(v string)`

SetId sets Id field to given value.


### GetSource

`func (o *WorkflowEdge) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WorkflowEdge) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WorkflowEdge) SetSource(v string)`

SetSource sets Source field to given value.


### GetTarget

`func (o *WorkflowEdge) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *WorkflowEdge) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *WorkflowEdge) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetSourceHandle

`func (o *WorkflowEdge) GetSourceHandle() string`

GetSourceHandle returns the SourceHandle field if non-nil, zero value otherwise.

### GetSourceHandleOk

`func (o *WorkflowEdge) GetSourceHandleOk() (*string, bool)`

GetSourceHandleOk returns a tuple with the SourceHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHandle

`func (o *WorkflowEdge) SetSourceHandle(v string)`

SetSourceHandle sets SourceHandle field to given value.

### HasSourceHandle

`func (o *WorkflowEdge) HasSourceHandle() bool`

HasSourceHandle returns a boolean if a field has been set.

### SetSourceHandleNil

`func (o *WorkflowEdge) SetSourceHandleNil(b bool)`

 SetSourceHandleNil sets the value for SourceHandle to be an explicit nil

### UnsetSourceHandle
`func (o *WorkflowEdge) UnsetSourceHandle()`

UnsetSourceHandle ensures that no value is present for SourceHandle, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


