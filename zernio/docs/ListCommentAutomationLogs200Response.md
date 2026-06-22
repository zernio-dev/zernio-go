# ListCommentAutomationLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Logs** | Pointer to [**[]GetCommentAutomation200ResponseLogsInner**](GetCommentAutomation200ResponseLogsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListCommentAutomationLogs200Response

`func NewListCommentAutomationLogs200Response() *ListCommentAutomationLogs200Response`

NewListCommentAutomationLogs200Response instantiates a new ListCommentAutomationLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentAutomationLogs200ResponseWithDefaults

`func NewListCommentAutomationLogs200ResponseWithDefaults() *ListCommentAutomationLogs200Response`

NewListCommentAutomationLogs200ResponseWithDefaults instantiates a new ListCommentAutomationLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListCommentAutomationLogs200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListCommentAutomationLogs200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListCommentAutomationLogs200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListCommentAutomationLogs200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetLogs

`func (o *ListCommentAutomationLogs200Response) GetLogs() []GetCommentAutomation200ResponseLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ListCommentAutomationLogs200Response) GetLogsOk() (*[]GetCommentAutomation200ResponseLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ListCommentAutomationLogs200Response) SetLogs(v []GetCommentAutomation200ResponseLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ListCommentAutomationLogs200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetPagination

`func (o *ListCommentAutomationLogs200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListCommentAutomationLogs200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListCommentAutomationLogs200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListCommentAutomationLogs200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


