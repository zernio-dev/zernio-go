# ListBroadcasts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Broadcasts** | Pointer to [**[]ListBroadcasts200ResponseBroadcastsInner**](ListBroadcasts200ResponseBroadcastsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListContacts200ResponsePagination**](ListContacts200ResponsePagination.md) |  | [optional] 

## Methods

### NewListBroadcasts200Response

`func NewListBroadcasts200Response() *ListBroadcasts200Response`

NewListBroadcasts200Response instantiates a new ListBroadcasts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBroadcasts200ResponseWithDefaults

`func NewListBroadcasts200ResponseWithDefaults() *ListBroadcasts200Response`

NewListBroadcasts200ResponseWithDefaults instantiates a new ListBroadcasts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListBroadcasts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListBroadcasts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListBroadcasts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListBroadcasts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBroadcasts

`func (o *ListBroadcasts200Response) GetBroadcasts() []ListBroadcasts200ResponseBroadcastsInner`

GetBroadcasts returns the Broadcasts field if non-nil, zero value otherwise.

### GetBroadcastsOk

`func (o *ListBroadcasts200Response) GetBroadcastsOk() (*[]ListBroadcasts200ResponseBroadcastsInner, bool)`

GetBroadcastsOk returns a tuple with the Broadcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcasts

`func (o *ListBroadcasts200Response) SetBroadcasts(v []ListBroadcasts200ResponseBroadcastsInner)`

SetBroadcasts sets Broadcasts field to given value.

### HasBroadcasts

`func (o *ListBroadcasts200Response) HasBroadcasts() bool`

HasBroadcasts returns a boolean if a field has been set.

### GetPagination

`func (o *ListBroadcasts200Response) GetPagination() ListContacts200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListBroadcasts200Response) GetPaginationOk() (*ListContacts200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListBroadcasts200Response) SetPagination(v ListContacts200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListBroadcasts200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


