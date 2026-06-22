# ListSequences200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Sequences** | Pointer to [**[]ListSequences200ResponseSequencesInner**](ListSequences200ResponseSequencesInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListSequences200Response

`func NewListSequences200Response() *ListSequences200Response`

NewListSequences200Response instantiates a new ListSequences200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSequences200ResponseWithDefaults

`func NewListSequences200ResponseWithDefaults() *ListSequences200Response`

NewListSequences200ResponseWithDefaults instantiates a new ListSequences200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListSequences200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListSequences200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListSequences200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListSequences200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSequences

`func (o *ListSequences200Response) GetSequences() []ListSequences200ResponseSequencesInner`

GetSequences returns the Sequences field if non-nil, zero value otherwise.

### GetSequencesOk

`func (o *ListSequences200Response) GetSequencesOk() (*[]ListSequences200ResponseSequencesInner, bool)`

GetSequencesOk returns a tuple with the Sequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequences

`func (o *ListSequences200Response) SetSequences(v []ListSequences200ResponseSequencesInner)`

SetSequences sets Sequences field to given value.

### HasSequences

`func (o *ListSequences200Response) HasSequences() bool`

HasSequences returns a boolean if a field has been set.

### GetPagination

`func (o *ListSequences200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSequences200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSequences200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSequences200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


