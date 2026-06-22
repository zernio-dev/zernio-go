# ListWorkflowExecutions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Executions** | Pointer to [**[]ListWorkflowExecutions200ResponseExecutionsInner**](ListWorkflowExecutions200ResponseExecutionsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListWorkflowExecutions200Response

`func NewListWorkflowExecutions200Response() *ListWorkflowExecutions200Response`

NewListWorkflowExecutions200Response instantiates a new ListWorkflowExecutions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowExecutions200ResponseWithDefaults

`func NewListWorkflowExecutions200ResponseWithDefaults() *ListWorkflowExecutions200Response`

NewListWorkflowExecutions200ResponseWithDefaults instantiates a new ListWorkflowExecutions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListWorkflowExecutions200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListWorkflowExecutions200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListWorkflowExecutions200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListWorkflowExecutions200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetExecutions

`func (o *ListWorkflowExecutions200Response) GetExecutions() []ListWorkflowExecutions200ResponseExecutionsInner`

GetExecutions returns the Executions field if non-nil, zero value otherwise.

### GetExecutionsOk

`func (o *ListWorkflowExecutions200Response) GetExecutionsOk() (*[]ListWorkflowExecutions200ResponseExecutionsInner, bool)`

GetExecutionsOk returns a tuple with the Executions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutions

`func (o *ListWorkflowExecutions200Response) SetExecutions(v []ListWorkflowExecutions200ResponseExecutionsInner)`

SetExecutions sets Executions field to given value.

### HasExecutions

`func (o *ListWorkflowExecutions200Response) HasExecutions() bool`

HasExecutions returns a boolean if a field has been set.

### GetPagination

`func (o *ListWorkflowExecutions200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWorkflowExecutions200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWorkflowExecutions200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListWorkflowExecutions200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


