# ListBroadcastRecipients200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Recipients** | Pointer to [**[]ListBroadcastRecipients200ResponseRecipientsInner**](ListBroadcastRecipients200ResponseRecipientsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListBroadcastRecipients200Response

`func NewListBroadcastRecipients200Response() *ListBroadcastRecipients200Response`

NewListBroadcastRecipients200Response instantiates a new ListBroadcastRecipients200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBroadcastRecipients200ResponseWithDefaults

`func NewListBroadcastRecipients200ResponseWithDefaults() *ListBroadcastRecipients200Response`

NewListBroadcastRecipients200ResponseWithDefaults instantiates a new ListBroadcastRecipients200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListBroadcastRecipients200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListBroadcastRecipients200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListBroadcastRecipients200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListBroadcastRecipients200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetRecipients

`func (o *ListBroadcastRecipients200Response) GetRecipients() []ListBroadcastRecipients200ResponseRecipientsInner`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListBroadcastRecipients200Response) GetRecipientsOk() (*[]ListBroadcastRecipients200ResponseRecipientsInner, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListBroadcastRecipients200Response) SetRecipients(v []ListBroadcastRecipients200ResponseRecipientsInner)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *ListBroadcastRecipients200Response) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetPagination

`func (o *ListBroadcastRecipients200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListBroadcastRecipients200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListBroadcastRecipients200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListBroadcastRecipients200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


