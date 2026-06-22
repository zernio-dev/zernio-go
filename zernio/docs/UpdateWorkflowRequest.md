# UpdateWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]WorkflowNode**](WorkflowNode.md) |  | [optional] 
**Edges** | Pointer to [**[]WorkflowEdge**](WorkflowEdge.md) |  | [optional] 
**EntryNodeId** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **string** | Reassign the workflow to a different &#x60;SocialAccount&#x60;. &#x60;platform&#x60; and &#x60;profileId&#x60; are derived server-side from the new account (the client never sends them directly). The account must belong to the caller&#39;s workspace and be on a workflow-supported platform (whatsapp, instagram, facebook, telegram, twitter, bluesky, reddit). Changing this triggers a graph revalidation against the new platform.  | [optional] 

## Methods

### NewUpdateWorkflowRequest

`func NewUpdateWorkflowRequest() *UpdateWorkflowRequest`

NewUpdateWorkflowRequest instantiates a new UpdateWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflowRequestWithDefaults

`func NewUpdateWorkflowRequestWithDefaults() *UpdateWorkflowRequest`

NewUpdateWorkflowRequestWithDefaults instantiates a new UpdateWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWorkflowRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateWorkflowRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWorkflowRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWorkflowRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWorkflowRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodes

`func (o *UpdateWorkflowRequest) GetNodes() []WorkflowNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *UpdateWorkflowRequest) GetNodesOk() (*[]WorkflowNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *UpdateWorkflowRequest) SetNodes(v []WorkflowNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *UpdateWorkflowRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *UpdateWorkflowRequest) GetEdges() []WorkflowEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *UpdateWorkflowRequest) GetEdgesOk() (*[]WorkflowEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *UpdateWorkflowRequest) SetEdges(v []WorkflowEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *UpdateWorkflowRequest) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetEntryNodeId

`func (o *UpdateWorkflowRequest) GetEntryNodeId() string`

GetEntryNodeId returns the EntryNodeId field if non-nil, zero value otherwise.

### GetEntryNodeIdOk

`func (o *UpdateWorkflowRequest) GetEntryNodeIdOk() (*string, bool)`

GetEntryNodeIdOk returns a tuple with the EntryNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryNodeId

`func (o *UpdateWorkflowRequest) SetEntryNodeId(v string)`

SetEntryNodeId sets EntryNodeId field to given value.

### HasEntryNodeId

`func (o *UpdateWorkflowRequest) HasEntryNodeId() bool`

HasEntryNodeId returns a boolean if a field has been set.

### SetEntryNodeIdNil

`func (o *UpdateWorkflowRequest) SetEntryNodeIdNil(b bool)`

 SetEntryNodeIdNil sets the value for EntryNodeId to be an explicit nil

### UnsetEntryNodeId
`func (o *UpdateWorkflowRequest) UnsetEntryNodeId()`

UnsetEntryNodeId ensures that no value is present for EntryNodeId, not even an explicit nil
### GetAccountId

`func (o *UpdateWorkflowRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWorkflowRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWorkflowRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateWorkflowRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


