# GetAdComments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Comments** | **[]map[string]interface{}** |  | 
**Pagination** | [**GetAdComments200ResponsePagination**](GetAdComments200ResponsePagination.md) |  | 
**Meta** | [**GetAdComments200ResponseMeta**](GetAdComments200ResponseMeta.md) |  | 

## Methods

### NewGetAdComments200Response

`func NewGetAdComments200Response(status string, comments []map[string]interface{}, pagination GetAdComments200ResponsePagination, meta GetAdComments200ResponseMeta, ) *GetAdComments200Response`

NewGetAdComments200Response instantiates a new GetAdComments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdComments200ResponseWithDefaults

`func NewGetAdComments200ResponseWithDefaults() *GetAdComments200Response`

NewGetAdComments200ResponseWithDefaults instantiates a new GetAdComments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAdComments200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAdComments200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAdComments200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetComments

`func (o *GetAdComments200Response) GetComments() []map[string]interface{}`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetAdComments200Response) GetCommentsOk() (*[]map[string]interface{}, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetAdComments200Response) SetComments(v []map[string]interface{})`

SetComments sets Comments field to given value.


### GetPagination

`func (o *GetAdComments200Response) GetPagination() GetAdComments200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetAdComments200Response) GetPaginationOk() (*GetAdComments200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetAdComments200Response) SetPagination(v GetAdComments200ResponsePagination)`

SetPagination sets Pagination field to given value.


### GetMeta

`func (o *GetAdComments200Response) GetMeta() GetAdComments200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAdComments200Response) GetMetaOk() (*GetAdComments200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAdComments200Response) SetMeta(v GetAdComments200ResponseMeta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


