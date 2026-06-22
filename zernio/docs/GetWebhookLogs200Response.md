# GetWebhookLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]WebhookLog**](WebhookLog.md) |  | [optional] 
**Pagination** | Pointer to [**GetWebhookLogs200ResponsePagination**](GetWebhookLogs200ResponsePagination.md) |  | [optional] 

## Methods

### NewGetWebhookLogs200Response

`func NewGetWebhookLogs200Response() *GetWebhookLogs200Response`

NewGetWebhookLogs200Response instantiates a new GetWebhookLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhookLogs200ResponseWithDefaults

`func NewGetWebhookLogs200ResponseWithDefaults() *GetWebhookLogs200Response`

NewGetWebhookLogs200ResponseWithDefaults instantiates a new GetWebhookLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *GetWebhookLogs200Response) GetLogs() []WebhookLog`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetWebhookLogs200Response) GetLogsOk() (*[]WebhookLog, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetWebhookLogs200Response) SetLogs(v []WebhookLog)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *GetWebhookLogs200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.

### GetPagination

`func (o *GetWebhookLogs200Response) GetPagination() GetWebhookLogs200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetWebhookLogs200Response) GetPaginationOk() (*GetWebhookLogs200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetWebhookLogs200Response) SetPagination(v GetWebhookLogs200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetWebhookLogs200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


