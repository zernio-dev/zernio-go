# ListLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]ListLogs200ResponseLogsInner**](ListLogs200ResponseLogsInner.md) |  | [optional] 
**Pagination** | Pointer to [**GetWebhookLogs200ResponsePagination**](GetWebhookLogs200ResponsePagination.md) |  | [optional] 

## Methods

### NewListLogs200Response

`func NewListLogs200Response() *ListLogs200Response`

NewListLogs200Response instantiates a new ListLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseWithDefaults

`func NewListLogs200ResponseWithDefaults() *ListLogs200Response`

NewListLogs200ResponseWithDefaults instantiates a new ListLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *ListLogs200Response) GetLogs() []ListLogs200ResponseLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *ListLogs200Response) GetLogsOk() (*[]ListLogs200ResponseLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *ListLogs200Response) SetLogs(v []ListLogs200ResponseLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *ListLogs200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetPagination

`func (o *ListLogs200Response) GetPagination() GetWebhookLogs200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListLogs200Response) GetPaginationOk() (*GetWebhookLogs200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListLogs200Response) SetPagination(v GetWebhookLogs200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListLogs200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


