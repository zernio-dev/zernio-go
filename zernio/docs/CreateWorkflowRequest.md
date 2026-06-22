# CreateWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**AccountId** | **string** |  | 
**Platform** | Pointer to **string** |  | [optional] [default to "whatsapp"]
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]WorkflowNode**](WorkflowNode.md) |  | [optional] 
**Edges** | Pointer to [**[]WorkflowEdge**](WorkflowEdge.md) |  | [optional] 
**EntryNodeId** | Pointer to **string** | The trigger node id; derived from the single trigger node if omitted | [optional] 

## Methods

### NewCreateWorkflowRequest

`func NewCreateWorkflowRequest(profileId string, accountId string, name string, ) *CreateWorkflowRequest`

NewCreateWorkflowRequest instantiates a new CreateWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowRequestWithDefaults

`func NewCreateWorkflowRequestWithDefaults() *CreateWorkflowRequest`

NewCreateWorkflowRequestWithDefaults instantiates a new CreateWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateWorkflowRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateWorkflowRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateWorkflowRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *CreateWorkflowRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateWorkflowRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateWorkflowRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPlatform

`func (o *CreateWorkflowRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateWorkflowRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateWorkflowRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateWorkflowRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetName

`func (o *CreateWorkflowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWorkflowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWorkflowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateWorkflowRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWorkflowRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWorkflowRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWorkflowRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNodes

`func (o *CreateWorkflowRequest) GetNodes() []WorkflowNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *CreateWorkflowRequest) GetNodesOk() (*[]WorkflowNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *CreateWorkflowRequest) SetNodes(v []WorkflowNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *CreateWorkflowRequest) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *CreateWorkflowRequest) GetEdges() []WorkflowEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *CreateWorkflowRequest) GetEdgesOk() (*[]WorkflowEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *CreateWorkflowRequest) SetEdges(v []WorkflowEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *CreateWorkflowRequest) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetEntryNodeId

`func (o *CreateWorkflowRequest) GetEntryNodeId() string`

GetEntryNodeId returns the EntryNodeId field if non-nil, zero value otherwise.

### GetEntryNodeIdOk

`func (o *CreateWorkflowRequest) GetEntryNodeIdOk() (*string, bool)`

GetEntryNodeIdOk returns a tuple with the EntryNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryNodeId

`func (o *CreateWorkflowRequest) SetEntryNodeId(v string)`

SetEntryNodeId sets EntryNodeId field to given value.

### HasEntryNodeId

`func (o *CreateWorkflowRequest) HasEntryNodeId() bool`

HasEntryNodeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


