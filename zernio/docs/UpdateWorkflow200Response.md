# UpdateWorkflow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Workflow** | Pointer to [**UpdateWorkflow200ResponseWorkflow**](UpdateWorkflow200ResponseWorkflow.md) |  | [optional] 

## Methods

### NewUpdateWorkflow200Response

`func NewUpdateWorkflow200Response() *UpdateWorkflow200Response`

NewUpdateWorkflow200Response instantiates a new UpdateWorkflow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWorkflow200ResponseWithDefaults

`func NewUpdateWorkflow200ResponseWithDefaults() *UpdateWorkflow200Response`

NewUpdateWorkflow200ResponseWithDefaults instantiates a new UpdateWorkflow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateWorkflow200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateWorkflow200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateWorkflow200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateWorkflow200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetWorkflow

`func (o *UpdateWorkflow200Response) GetWorkflow() UpdateWorkflow200ResponseWorkflow`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *UpdateWorkflow200Response) GetWorkflowOk() (*UpdateWorkflow200ResponseWorkflow, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *UpdateWorkflow200Response) SetWorkflow(v UpdateWorkflow200ResponseWorkflow)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *UpdateWorkflow200Response) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


