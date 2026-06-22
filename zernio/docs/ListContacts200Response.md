# ListContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Contacts** | Pointer to [**[]ListContacts200ResponseContactsInner**](ListContacts200ResponseContactsInner.md) |  | [optional] 
**Filters** | Pointer to [**ListContacts200ResponseFilters**](ListContacts200ResponseFilters.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListContacts200Response

`func NewListContacts200Response() *ListContacts200Response`

NewListContacts200Response instantiates a new ListContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListContacts200ResponseWithDefaults

`func NewListContacts200ResponseWithDefaults() *ListContacts200Response`

NewListContacts200ResponseWithDefaults instantiates a new ListContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListContacts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListContacts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListContacts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListContacts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetContacts

`func (o *ListContacts200Response) GetContacts() []ListContacts200ResponseContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *ListContacts200Response) GetContactsOk() (*[]ListContacts200ResponseContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *ListContacts200Response) SetContacts(v []ListContacts200ResponseContactsInner)`

SetContacts sets Contacts field to given value.

### HasContacts

`func (o *ListContacts200Response) HasContacts() bool`

HasContacts returns a boolean if a field has been set.

### GetFilters

`func (o *ListContacts200Response) GetFilters() ListContacts200ResponseFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListContacts200Response) GetFiltersOk() (*ListContacts200ResponseFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListContacts200Response) SetFilters(v ListContacts200ResponseFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListContacts200Response) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetPagination

`func (o *ListContacts200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListContacts200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListContacts200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListContacts200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


