# ListInboxConversationAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Items** | Pointer to [**[]ListInboxConversationAnalytics200ResponseItemsInner**](ListInboxConversationAnalytics200ResponseItemsInner.md) |  | [optional] 
**Pagination** | Pointer to [**ListInboxConversationAnalytics200ResponsePagination**](ListInboxConversationAnalytics200ResponsePagination.md) |  | [optional] 

## Methods

### NewListInboxConversationAnalytics200Response

`func NewListInboxConversationAnalytics200Response() *ListInboxConversationAnalytics200Response`

NewListInboxConversationAnalytics200Response instantiates a new ListInboxConversationAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversationAnalytics200ResponseWithDefaults

`func NewListInboxConversationAnalytics200ResponseWithDefaults() *ListInboxConversationAnalytics200Response`

NewListInboxConversationAnalytics200ResponseWithDefaults instantiates a new ListInboxConversationAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListInboxConversationAnalytics200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListInboxConversationAnalytics200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListInboxConversationAnalytics200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListInboxConversationAnalytics200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFrom

`func (o *ListInboxConversationAnalytics200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListInboxConversationAnalytics200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListInboxConversationAnalytics200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListInboxConversationAnalytics200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ListInboxConversationAnalytics200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListInboxConversationAnalytics200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListInboxConversationAnalytics200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ListInboxConversationAnalytics200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *ListInboxConversationAnalytics200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *ListInboxConversationAnalytics200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetItems

`func (o *ListInboxConversationAnalytics200Response) GetItems() []ListInboxConversationAnalytics200ResponseItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListInboxConversationAnalytics200Response) GetItemsOk() (*[]ListInboxConversationAnalytics200ResponseItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListInboxConversationAnalytics200Response) SetItems(v []ListInboxConversationAnalytics200ResponseItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListInboxConversationAnalytics200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetPagination

`func (o *ListInboxConversationAnalytics200Response) GetPagination() ListInboxConversationAnalytics200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListInboxConversationAnalytics200Response) GetPaginationOk() (*ListInboxConversationAnalytics200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListInboxConversationAnalytics200Response) SetPagination(v ListInboxConversationAnalytics200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListInboxConversationAnalytics200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


