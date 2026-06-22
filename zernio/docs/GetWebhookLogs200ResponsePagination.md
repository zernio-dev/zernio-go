# GetWebhookLogs200ResponsePagination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of matching logs | [optional] 
**Limit** | Pointer to **int32** | Maximum number of logs returned per page | [optional] 
**Skip** | Pointer to **int32** | Number of logs skipped | [optional] 
**Pages** | Pointer to **int32** | Total number of pages | [optional] 
**HasMore** | Pointer to **bool** | Whether more logs are available beyond this page | [optional] 

## Methods

### NewGetWebhookLogs200ResponsePagination

`func NewGetWebhookLogs200ResponsePagination() *GetWebhookLogs200ResponsePagination`

NewGetWebhookLogs200ResponsePagination instantiates a new GetWebhookLogs200ResponsePagination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhookLogs200ResponsePaginationWithDefaults

`func NewGetWebhookLogs200ResponsePaginationWithDefaults() *GetWebhookLogs200ResponsePagination`

NewGetWebhookLogs200ResponsePaginationWithDefaults instantiates a new GetWebhookLogs200ResponsePagination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetWebhookLogs200ResponsePagination) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetWebhookLogs200ResponsePagination) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetWebhookLogs200ResponsePagination) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetWebhookLogs200ResponsePagination) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLimit

`func (o *GetWebhookLogs200ResponsePagination) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetWebhookLogs200ResponsePagination) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetWebhookLogs200ResponsePagination) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GetWebhookLogs200ResponsePagination) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetSkip

`func (o *GetWebhookLogs200ResponsePagination) GetSkip() int32`

GetSkip returns the Skip field if non-nil, zero value otherwise.

### GetSkipOk

`func (o *GetWebhookLogs200ResponsePagination) GetSkipOk() (*int32, bool)`

GetSkipOk returns a tuple with the Skip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkip

`func (o *GetWebhookLogs200ResponsePagination) SetSkip(v int32)`

SetSkip sets Skip field to given value.

### HasSkip

`func (o *GetWebhookLogs200ResponsePagination) HasSkip() bool`

HasSkip returns a boolean if a field has been set.

### GetPages

`func (o *GetWebhookLogs200ResponsePagination) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *GetWebhookLogs200ResponsePagination) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *GetWebhookLogs200ResponsePagination) SetPages(v int32)`

SetPages sets Pages field to given value.

### HasPages

`func (o *GetWebhookLogs200ResponsePagination) HasPages() bool`

HasPages returns a boolean if a field has been set.

### GetHasMore

`func (o *GetWebhookLogs200ResponsePagination) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetWebhookLogs200ResponsePagination) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetWebhookLogs200ResponsePagination) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetWebhookLogs200ResponsePagination) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


