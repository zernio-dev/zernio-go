# ListInboxReviews200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]ListInboxReviews200ResponseDataInner**](ListInboxReviews200ResponseDataInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListInboxConversations200ResponsePagination**](ListInboxConversations200ResponsePagination.md) |  | [optional] 
**Meta** | Pointer to [**ListInboxConversations200ResponseMeta**](ListInboxConversations200ResponseMeta.md) |  | [optional] 
**Summary** | Pointer to [**ListInboxReviews200ResponseSummary**](ListInboxReviews200ResponseSummary.md) |  | [optional] 

## Methods

### NewListInboxReviews200Response

`func NewListInboxReviews200Response() *ListInboxReviews200Response`

NewListInboxReviews200Response instantiates a new ListInboxReviews200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxReviews200ResponseWithDefaults

`func NewListInboxReviews200ResponseWithDefaults() *ListInboxReviews200Response`

NewListInboxReviews200ResponseWithDefaults instantiates a new ListInboxReviews200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListInboxReviews200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInboxReviews200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInboxReviews200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListInboxReviews200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *ListInboxReviews200Response) GetData() []ListInboxReviews200ResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListInboxReviews200Response) GetDataOk() (*[]ListInboxReviews200ResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListInboxReviews200Response) SetData(v []ListInboxReviews200ResponseDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListInboxReviews200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetPagination

`func (o *ListInboxReviews200Response) GetPagination() ListInboxConversations200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInboxReviews200Response) GetPaginationOk() (*ListInboxConversations200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInboxReviews200Response) SetPagination(v ListInboxConversations200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListInboxReviews200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetMeta

`func (o *ListInboxReviews200Response) GetMeta() ListInboxConversations200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInboxReviews200Response) GetMetaOk() (*ListInboxConversations200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInboxReviews200Response) SetMeta(v ListInboxConversations200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInboxReviews200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetSummary

`func (o *ListInboxReviews200Response) GetSummary() ListInboxReviews200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *ListInboxReviews200Response) GetSummaryOk() (*ListInboxReviews200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *ListInboxReviews200Response) SetSummary(v ListInboxReviews200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *ListInboxReviews200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


