# ListLeadForms200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Forms** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Pagination** | Pointer to [**GetInboxPostComments200ResponsePagination**](GetInboxPostComments200ResponsePagination.md) |  | [optional] 

## Methods

### NewListLeadForms200Response

`func NewListLeadForms200Response() *ListLeadForms200Response`

NewListLeadForms200Response instantiates a new ListLeadForms200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLeadForms200ResponseWithDefaults

`func NewListLeadForms200ResponseWithDefaults() *ListLeadForms200Response`

NewListLeadForms200ResponseWithDefaults instantiates a new ListLeadForms200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ListLeadForms200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLeadForms200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLeadForms200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLeadForms200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetForms

`func (o *ListLeadForms200Response) GetForms() []map[string]interface{}`

GetForms returns the Forms field if non-nil, zero value otherwise.

### GetFormsOk

`func (o *ListLeadForms200Response) GetFormsOk() (*[]map[string]interface{}, bool)`

GetFormsOk returns a tuple with the Forms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForms

`func (o *ListLeadForms200Response) SetForms(v []map[string]interface{})`

SetForms sets Forms field to given value.

### HasForms

`func (o *ListLeadForms200Response) HasForms() bool`

HasForms returns a boolean if a field has been set.

### GetPagination

`func (o *ListLeadForms200Response) GetPagination() GetInboxPostComments200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListLeadForms200Response) GetPaginationOk() (*GetInboxPostComments200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListLeadForms200Response) SetPagination(v GetInboxPostComments200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListLeadForms200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


