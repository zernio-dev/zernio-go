# GetWorkflow200ResponseWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**EntryNodeId** | Pointer to **string** |  | [optional] 
**Nodes** | Pointer to [**[]WorkflowNode**](WorkflowNode.md) |  | [optional] 
**Edges** | Pointer to [**[]WorkflowEdge**](WorkflowEdge.md) |  | [optional] 
**TotalStarted** | Pointer to **int32** |  | [optional] 
**TotalCompleted** | Pointer to **int32** |  | [optional] 
**TotalExited** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetWorkflow200ResponseWorkflow

`func NewGetWorkflow200ResponseWorkflow() *GetWorkflow200ResponseWorkflow`

NewGetWorkflow200ResponseWorkflow instantiates a new GetWorkflow200ResponseWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflow200ResponseWorkflowWithDefaults

`func NewGetWorkflow200ResponseWorkflowWithDefaults() *GetWorkflow200ResponseWorkflow`

NewGetWorkflow200ResponseWorkflowWithDefaults instantiates a new GetWorkflow200ResponseWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetWorkflow200ResponseWorkflow) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetWorkflow200ResponseWorkflow) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetWorkflow200ResponseWorkflow) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetWorkflow200ResponseWorkflow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetWorkflow200ResponseWorkflow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetWorkflow200ResponseWorkflow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetWorkflow200ResponseWorkflow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetWorkflow200ResponseWorkflow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetWorkflow200ResponseWorkflow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetWorkflow200ResponseWorkflow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetWorkflow200ResponseWorkflow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetWorkflow200ResponseWorkflow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *GetWorkflow200ResponseWorkflow) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetWorkflow200ResponseWorkflow) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetWorkflow200ResponseWorkflow) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetWorkflow200ResponseWorkflow) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *GetWorkflow200ResponseWorkflow) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetWorkflow200ResponseWorkflow) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetWorkflow200ResponseWorkflow) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetWorkflow200ResponseWorkflow) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProfileId

`func (o *GetWorkflow200ResponseWorkflow) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetWorkflow200ResponseWorkflow) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetWorkflow200ResponseWorkflow) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetWorkflow200ResponseWorkflow) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetStatus

`func (o *GetWorkflow200ResponseWorkflow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWorkflow200ResponseWorkflow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWorkflow200ResponseWorkflow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetWorkflow200ResponseWorkflow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntryNodeId

`func (o *GetWorkflow200ResponseWorkflow) GetEntryNodeId() string`

GetEntryNodeId returns the EntryNodeId field if non-nil, zero value otherwise.

### GetEntryNodeIdOk

`func (o *GetWorkflow200ResponseWorkflow) GetEntryNodeIdOk() (*string, bool)`

GetEntryNodeIdOk returns a tuple with the EntryNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryNodeId

`func (o *GetWorkflow200ResponseWorkflow) SetEntryNodeId(v string)`

SetEntryNodeId sets EntryNodeId field to given value.

### HasEntryNodeId

`func (o *GetWorkflow200ResponseWorkflow) HasEntryNodeId() bool`

HasEntryNodeId returns a boolean if a field has been set.

### GetNodes

`func (o *GetWorkflow200ResponseWorkflow) GetNodes() []WorkflowNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetWorkflow200ResponseWorkflow) GetNodesOk() (*[]WorkflowNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetWorkflow200ResponseWorkflow) SetNodes(v []WorkflowNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GetWorkflow200ResponseWorkflow) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *GetWorkflow200ResponseWorkflow) GetEdges() []WorkflowEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GetWorkflow200ResponseWorkflow) GetEdgesOk() (*[]WorkflowEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GetWorkflow200ResponseWorkflow) SetEdges(v []WorkflowEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *GetWorkflow200ResponseWorkflow) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetTotalStarted

`func (o *GetWorkflow200ResponseWorkflow) GetTotalStarted() int32`

GetTotalStarted returns the TotalStarted field if non-nil, zero value otherwise.

### GetTotalStartedOk

`func (o *GetWorkflow200ResponseWorkflow) GetTotalStartedOk() (*int32, bool)`

GetTotalStartedOk returns a tuple with the TotalStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStarted

`func (o *GetWorkflow200ResponseWorkflow) SetTotalStarted(v int32)`

SetTotalStarted sets TotalStarted field to given value.

### HasTotalStarted

`func (o *GetWorkflow200ResponseWorkflow) HasTotalStarted() bool`

HasTotalStarted returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *GetWorkflow200ResponseWorkflow) GetTotalCompleted() int32`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *GetWorkflow200ResponseWorkflow) GetTotalCompletedOk() (*int32, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *GetWorkflow200ResponseWorkflow) SetTotalCompleted(v int32)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *GetWorkflow200ResponseWorkflow) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetTotalExited

`func (o *GetWorkflow200ResponseWorkflow) GetTotalExited() int32`

GetTotalExited returns the TotalExited field if non-nil, zero value otherwise.

### GetTotalExitedOk

`func (o *GetWorkflow200ResponseWorkflow) GetTotalExitedOk() (*int32, bool)`

GetTotalExitedOk returns a tuple with the TotalExited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExited

`func (o *GetWorkflow200ResponseWorkflow) SetTotalExited(v int32)`

SetTotalExited sets TotalExited field to given value.

### HasTotalExited

`func (o *GetWorkflow200ResponseWorkflow) HasTotalExited() bool`

HasTotalExited returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetWorkflow200ResponseWorkflow) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetWorkflow200ResponseWorkflow) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetWorkflow200ResponseWorkflow) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetWorkflow200ResponseWorkflow) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetWorkflow200ResponseWorkflow) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetWorkflow200ResponseWorkflow) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetWorkflow200ResponseWorkflow) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetWorkflow200ResponseWorkflow) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


