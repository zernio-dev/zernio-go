# ListWorkflows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Workflows** | Pointer to [**[]ListWorkflows200ResponseWorkflowsInner**](ListWorkflows200ResponseWorkflowsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListWorkflows200Response

`func NewListWorkflows200Response() *ListWorkflows200Response`

NewListWorkflows200Response instantiates a new ListWorkflows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflows200ResponseWithDefaults

`func NewListWorkflows200ResponseWithDefaults() *ListWorkflows200Response`

NewListWorkflows200ResponseWithDefaults instantiates a new ListWorkflows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListWorkflows200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListWorkflows200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListWorkflows200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListWorkflows200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetWorkflows

`func (o *ListWorkflows200Response) GetWorkflows() []ListWorkflows200ResponseWorkflowsInner`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *ListWorkflows200Response) GetWorkflowsOk() (*[]ListWorkflows200ResponseWorkflowsInner, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *ListWorkflows200Response) SetWorkflows(v []ListWorkflows200ResponseWorkflowsInner)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *ListWorkflows200Response) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.

### GetPagination

`func (o *ListWorkflows200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListWorkflows200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListWorkflows200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListWorkflows200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


