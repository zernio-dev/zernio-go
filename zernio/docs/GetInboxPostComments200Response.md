# GetInboxPostComments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to [**[]GetInboxPostComments200ResponseCommentsInner**](GetInboxPostComments200ResponseCommentsInner.md) |  | [optional] 
**Post** | Pointer to [**GetInboxPostComments200ResponsePost**](GetInboxPostComments200ResponsePost.md) |  | [optional] 
**Pagination** | Pointer to [**GetInboxPostComments200ResponsePagination**](GetInboxPostComments200ResponsePagination.md) |  | [optional] 
**Meta** | Pointer to [**GetInboxPostComments200ResponseMeta**](GetInboxPostComments200ResponseMeta.md) |  | [optional] 

## Methods

### NewGetInboxPostComments200Response

`func NewGetInboxPostComments200Response() *GetInboxPostComments200Response`

NewGetInboxPostComments200Response instantiates a new GetInboxPostComments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxPostComments200ResponseWithDefaults

`func NewGetInboxPostComments200ResponseWithDefaults() *GetInboxPostComments200Response`

NewGetInboxPostComments200ResponseWithDefaults instantiates a new GetInboxPostComments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetInboxPostComments200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInboxPostComments200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInboxPostComments200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInboxPostComments200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetComments

`func (o *GetInboxPostComments200Response) GetComments() []GetInboxPostComments200ResponseCommentsInner`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetInboxPostComments200Response) GetCommentsOk() (*[]GetInboxPostComments200ResponseCommentsInner, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetInboxPostComments200Response) SetComments(v []GetInboxPostComments200ResponseCommentsInner)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GetInboxPostComments200Response) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetPost

`func (o *GetInboxPostComments200Response) GetPost() GetInboxPostComments200ResponsePost`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *GetInboxPostComments200Response) GetPostOk() (*GetInboxPostComments200ResponsePost, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *GetInboxPostComments200Response) SetPost(v GetInboxPostComments200ResponsePost)`

SetPost sets Post field to given value.

### HasPost

`func (o *GetInboxPostComments200Response) HasPost() bool`

HasPost returns a boolean if a field has been set.

### GetPagination

`func (o *GetInboxPostComments200Response) GetPagination() GetInboxPostComments200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetInboxPostComments200Response) GetPaginationOk() (*GetInboxPostComments200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetInboxPostComments200Response) SetPagination(v GetInboxPostComments200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetInboxPostComments200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetMeta

`func (o *GetInboxPostComments200Response) GetMeta() GetInboxPostComments200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInboxPostComments200Response) GetMetaOk() (*GetInboxPostComments200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInboxPostComments200Response) SetMeta(v GetInboxPostComments200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInboxPostComments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


