# ListSequenceEnrollments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Enrollments** | Pointer to [**[]ListSequenceEnrollments200ResponseEnrollmentsInner**](ListSequenceEnrollments200ResponseEnrollmentsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListSequenceEnrollments200Response

`func NewListSequenceEnrollments200Response() *ListSequenceEnrollments200Response`

NewListSequenceEnrollments200Response instantiates a new ListSequenceEnrollments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSequenceEnrollments200ResponseWithDefaults

`func NewListSequenceEnrollments200ResponseWithDefaults() *ListSequenceEnrollments200Response`

NewListSequenceEnrollments200ResponseWithDefaults instantiates a new ListSequenceEnrollments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListSequenceEnrollments200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListSequenceEnrollments200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListSequenceEnrollments200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListSequenceEnrollments200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetEnrollments

`func (o *ListSequenceEnrollments200Response) GetEnrollments() []ListSequenceEnrollments200ResponseEnrollmentsInner`

GetEnrollments returns the Enrollments field if non-nil, zero value otherwise.

### GetEnrollmentsOk

`func (o *ListSequenceEnrollments200Response) GetEnrollmentsOk() (*[]ListSequenceEnrollments200ResponseEnrollmentsInner, bool)`

GetEnrollmentsOk returns a tuple with the Enrollments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollments

`func (o *ListSequenceEnrollments200Response) SetEnrollments(v []ListSequenceEnrollments200ResponseEnrollmentsInner)`

SetEnrollments sets Enrollments field to given value.

### HasEnrollments

`func (o *ListSequenceEnrollments200Response) HasEnrollments() bool`

HasEnrollments returns a boolean if a field has been set.

### GetPagination

`func (o *ListSequenceEnrollments200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListSequenceEnrollments200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListSequenceEnrollments200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListSequenceEnrollments200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


