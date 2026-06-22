# GetWorkflowVersion200ResponseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EntryNodeId** | Pointer to **NullableString** |  | [optional] 
**Nodes** | Pointer to [**[]WorkflowNode**](WorkflowNode.md) |  | [optional] 
**Edges** | Pointer to [**[]WorkflowEdge**](WorkflowEdge.md) |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** |  | [optional] 
**CreatedByEmail** | Pointer to **NullableString** |  | [optional] 
**RestoredFromVersion** | Pointer to **NullableInt32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetWorkflowVersion200ResponseVersion

`func NewGetWorkflowVersion200ResponseVersion() *GetWorkflowVersion200ResponseVersion`

NewGetWorkflowVersion200ResponseVersion instantiates a new GetWorkflowVersion200ResponseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowVersion200ResponseVersionWithDefaults

`func NewGetWorkflowVersion200ResponseVersionWithDefaults() *GetWorkflowVersion200ResponseVersion`

NewGetWorkflowVersion200ResponseVersionWithDefaults instantiates a new GetWorkflowVersion200ResponseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *GetWorkflowVersion200ResponseVersion) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetWorkflowVersion200ResponseVersion) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetWorkflowVersion200ResponseVersion) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetWorkflowVersion200ResponseVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *GetWorkflowVersion200ResponseVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetWorkflowVersion200ResponseVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetWorkflowVersion200ResponseVersion) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetWorkflowVersion200ResponseVersion) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetWorkflowVersion200ResponseVersion) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetWorkflowVersion200ResponseVersion) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetWorkflowVersion200ResponseVersion) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetWorkflowVersion200ResponseVersion) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetWorkflowVersion200ResponseVersion) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetWorkflowVersion200ResponseVersion) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEntryNodeId

`func (o *GetWorkflowVersion200ResponseVersion) GetEntryNodeId() string`

GetEntryNodeId returns the EntryNodeId field if non-nil, zero value otherwise.

### GetEntryNodeIdOk

`func (o *GetWorkflowVersion200ResponseVersion) GetEntryNodeIdOk() (*string, bool)`

GetEntryNodeIdOk returns a tuple with the EntryNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryNodeId

`func (o *GetWorkflowVersion200ResponseVersion) SetEntryNodeId(v string)`

SetEntryNodeId sets EntryNodeId field to given value.

### HasEntryNodeId

`func (o *GetWorkflowVersion200ResponseVersion) HasEntryNodeId() bool`

HasEntryNodeId returns a boolean if a field has been set.

### SetEntryNodeIdNil

`func (o *GetWorkflowVersion200ResponseVersion) SetEntryNodeIdNil(b bool)`

 SetEntryNodeIdNil sets the value for EntryNodeId to be an explicit nil

### UnsetEntryNodeId
`func (o *GetWorkflowVersion200ResponseVersion) UnsetEntryNodeId()`

UnsetEntryNodeId ensures that no value is present for EntryNodeId, not even an explicit nil
### GetNodes

`func (o *GetWorkflowVersion200ResponseVersion) GetNodes() []WorkflowNode`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *GetWorkflowVersion200ResponseVersion) GetNodesOk() (*[]WorkflowNode, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *GetWorkflowVersion200ResponseVersion) SetNodes(v []WorkflowNode)`

SetNodes sets Nodes field to given value.

### HasNodes

`func (o *GetWorkflowVersion200ResponseVersion) HasNodes() bool`

HasNodes returns a boolean if a field has been set.

### GetEdges

`func (o *GetWorkflowVersion200ResponseVersion) GetEdges() []WorkflowEdge`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GetWorkflowVersion200ResponseVersion) GetEdgesOk() (*[]WorkflowEdge, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GetWorkflowVersion200ResponseVersion) SetEdges(v []WorkflowEdge)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *GetWorkflowVersion200ResponseVersion) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetPlatform

`func (o *GetWorkflowVersion200ResponseVersion) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetWorkflowVersion200ResponseVersion) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetWorkflowVersion200ResponseVersion) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetWorkflowVersion200ResponseVersion) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *GetWorkflowVersion200ResponseVersion) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetWorkflowVersion200ResponseVersion) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetWorkflowVersion200ResponseVersion) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetWorkflowVersion200ResponseVersion) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProfileId

`func (o *GetWorkflowVersion200ResponseVersion) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetWorkflowVersion200ResponseVersion) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetWorkflowVersion200ResponseVersion) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetWorkflowVersion200ResponseVersion) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetWorkflowVersion200ResponseVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetWorkflowVersion200ResponseVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetWorkflowVersion200ResponseVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetWorkflowVersion200ResponseVersion) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *GetWorkflowVersion200ResponseVersion) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *GetWorkflowVersion200ResponseVersion) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedByEmail

`func (o *GetWorkflowVersion200ResponseVersion) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *GetWorkflowVersion200ResponseVersion) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *GetWorkflowVersion200ResponseVersion) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *GetWorkflowVersion200ResponseVersion) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### SetCreatedByEmailNil

`func (o *GetWorkflowVersion200ResponseVersion) SetCreatedByEmailNil(b bool)`

 SetCreatedByEmailNil sets the value for CreatedByEmail to be an explicit nil

### UnsetCreatedByEmail
`func (o *GetWorkflowVersion200ResponseVersion) UnsetCreatedByEmail()`

UnsetCreatedByEmail ensures that no value is present for CreatedByEmail, not even an explicit nil
### GetRestoredFromVersion

`func (o *GetWorkflowVersion200ResponseVersion) GetRestoredFromVersion() int32`

GetRestoredFromVersion returns the RestoredFromVersion field if non-nil, zero value otherwise.

### GetRestoredFromVersionOk

`func (o *GetWorkflowVersion200ResponseVersion) GetRestoredFromVersionOk() (*int32, bool)`

GetRestoredFromVersionOk returns a tuple with the RestoredFromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredFromVersion

`func (o *GetWorkflowVersion200ResponseVersion) SetRestoredFromVersion(v int32)`

SetRestoredFromVersion sets RestoredFromVersion field to given value.

### HasRestoredFromVersion

`func (o *GetWorkflowVersion200ResponseVersion) HasRestoredFromVersion() bool`

HasRestoredFromVersion returns a boolean if a field has been set.

### SetRestoredFromVersionNil

`func (o *GetWorkflowVersion200ResponseVersion) SetRestoredFromVersionNil(b bool)`

 SetRestoredFromVersionNil sets the value for RestoredFromVersion to be an explicit nil

### UnsetRestoredFromVersion
`func (o *GetWorkflowVersion200ResponseVersion) UnsetRestoredFromVersion()`

UnsetRestoredFromVersion ensures that no value is present for RestoredFromVersion, not even an explicit nil
### GetCreatedAt

`func (o *GetWorkflowVersion200ResponseVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetWorkflowVersion200ResponseVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetWorkflowVersion200ResponseVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetWorkflowVersion200ResponseVersion) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


