# ActivateWorkflow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Workflow** | Pointer to [**ActivateWorkflow200ResponseWorkflow**](ActivateWorkflow200ResponseWorkflow.md) |  | [optional] 

## Methods

### NewActivateWorkflow200Response

`func NewActivateWorkflow200Response() *ActivateWorkflow200Response`

NewActivateWorkflow200Response instantiates a new ActivateWorkflow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivateWorkflow200ResponseWithDefaults

`func NewActivateWorkflow200ResponseWithDefaults() *ActivateWorkflow200Response`

NewActivateWorkflow200ResponseWithDefaults instantiates a new ActivateWorkflow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ActivateWorkflow200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ActivateWorkflow200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ActivateWorkflow200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ActivateWorkflow200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetWorkflow

`func (o *ActivateWorkflow200Response) GetWorkflow() ActivateWorkflow200ResponseWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *ActivateWorkflow200Response) GetWorkflowOk() (*ActivateWorkflow200ResponseWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *ActivateWorkflow200Response) SetWorkflow(v ActivateWorkflow200ResponseWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *ActivateWorkflow200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


