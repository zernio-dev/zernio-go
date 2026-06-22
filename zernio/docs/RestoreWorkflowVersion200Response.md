# RestoreWorkflowVersion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Workflow** | Pointer to [**RestoreWorkflowVersion200ResponseWorkflow**](RestoreWorkflowVersion200ResponseWorkflow.md) |  | [optional] 
**RestoredFromVersion** | Pointer to **int32** |  | [optional] 

## Methods

### NewRestoreWorkflowVersion200Response

`func NewRestoreWorkflowVersion200Response() *RestoreWorkflowVersion200Response`

NewRestoreWorkflowVersion200Response instantiates a new RestoreWorkflowVersion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreWorkflowVersion200ResponseWithDefaults

`func NewRestoreWorkflowVersion200ResponseWithDefaults() *RestoreWorkflowVersion200Response`

NewRestoreWorkflowVersion200ResponseWithDefaults instantiates a new RestoreWorkflowVersion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RestoreWorkflowVersion200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RestoreWorkflowVersion200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RestoreWorkflowVersion200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RestoreWorkflowVersion200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetWorkflow

`func (o *RestoreWorkflowVersion200Response) GetWorkflow() RestoreWorkflowVersion200ResponseWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *RestoreWorkflowVersion200Response) GetWorkflowOk() (*RestoreWorkflowVersion200ResponseWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *RestoreWorkflowVersion200Response) SetWorkflow(v RestoreWorkflowVersion200ResponseWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *RestoreWorkflowVersion200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.

### GetRestoredFromVersion

`func (o *RestoreWorkflowVersion200Response) GetRestoredFromVersion() int32`

GetRestoredFromVersion returns the RestoredFromVersion field if non-nil, zero value otherwise.

### GetRestoredFromVersionOk

`func (o *RestoreWorkflowVersion200Response) GetRestoredFromVersionOk() (*int32, bool)`

GetRestoredFromVersionOk returns a tuple with the RestoredFromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredFromVersion

`func (o *RestoreWorkflowVersion200Response) SetRestoredFromVersion(v int32)`

SetRestoredFromVersion sets RestoredFromVersion field to given value.

### HasRestoredFromVersion

`func (o *RestoreWorkflowVersion200Response) HasRestoredFromVersion() bool`

HasRestoredFromVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


