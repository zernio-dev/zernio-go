# ListWorkflowExecutionEvents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Execution** | Pointer to [**ListWorkflowExecutionEvents200ResponseExecution**](ListWorkflowExecutionEvents200ResponseExecution.md) |  | [optional] 
**Events** | Pointer to [**[]WorkflowExecutionEvent**](WorkflowExecutionEvent.md) | Events in chronological order (oldest first). | [optional] 

## Methods

### NewListWorkflowExecutionEvents200Response

`func NewListWorkflowExecutionEvents200Response() *ListWorkflowExecutionEvents200Response`

NewListWorkflowExecutionEvents200Response instantiates a new ListWorkflowExecutionEvents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowExecutionEvents200ResponseWithDefaults

`func NewListWorkflowExecutionEvents200ResponseWithDefaults() *ListWorkflowExecutionEvents200Response`

NewListWorkflowExecutionEvents200ResponseWithDefaults instantiates a new ListWorkflowExecutionEvents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListWorkflowExecutionEvents200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListWorkflowExecutionEvents200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListWorkflowExecutionEvents200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListWorkflowExecutionEvents200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetExecution

`func (o *ListWorkflowExecutionEvents200Response) GetExecution() ListWorkflowExecutionEvents200ResponseExecution`

GetExecution returns the Execution field if non-nil, zero value otherwise.

### GetExecutionOk

`func (o *ListWorkflowExecutionEvents200Response) GetExecutionOk() (*ListWorkflowExecutionEvents200ResponseExecution, bool)`

GetExecutionOk returns a tuple with the Execution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecution

`func (o *ListWorkflowExecutionEvents200Response) SetExecution(v ListWorkflowExecutionEvents200ResponseExecution)`

SetExecution sets Execution field to given value.

### HasExecution

`func (o *ListWorkflowExecutionEvents200Response) HasExecution() bool`

HasExecution returns a boolean if a field has been set.

### GetEvents

`func (o *ListWorkflowExecutionEvents200Response) GetEvents() []WorkflowExecutionEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListWorkflowExecutionEvents200Response) GetEventsOk() (*[]WorkflowExecutionEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListWorkflowExecutionEvents200Response) SetEvents(v []WorkflowExecutionEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListWorkflowExecutionEvents200Response) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


