# ListInboxConversations200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountsQueried** | Pointer to **int32** |  | [optional] 
**AccountsFailed** | Pointer to **int32** |  | [optional] 
**FailedAccounts** | Pointer to [**[]ListInboxConversations200ResponseMetaFailedAccountsInner**](ListInboxConversations200ResponseMetaFailedAccountsInner.md) |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListInboxConversations200ResponseMeta

`func NewListInboxConversations200ResponseMeta() *ListInboxConversations200ResponseMeta`

NewListInboxConversations200ResponseMeta instantiates a new ListInboxConversations200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversations200ResponseMetaWithDefaults

`func NewListInboxConversations200ResponseMetaWithDefaults() *ListInboxConversations200ResponseMeta`

NewListInboxConversations200ResponseMetaWithDefaults instantiates a new ListInboxConversations200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountsQueried

`func (o *ListInboxConversations200ResponseMeta) GetAccountsQueried() int32`

GetAccountsQueried returns the AccountsQueried field if non-nil, zero value otherwise.

### GetAccountsQueriedOk

`func (o *ListInboxConversations200ResponseMeta) GetAccountsQueriedOk() (*int32, bool)`

GetAccountsQueriedOk returns a tuple with the AccountsQueried field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsQueried

`func (o *ListInboxConversations200ResponseMeta) SetAccountsQueried(v int32)`

SetAccountsQueried sets AccountsQueried field to given value.

### HasAccountsQueried

`func (o *ListInboxConversations200ResponseMeta) HasAccountsQueried() bool`

HasAccountsQueried returns a boolean if a field has been set.

### GetAccountsFailed

`func (o *ListInboxConversations200ResponseMeta) GetAccountsFailed() int32`

GetAccountsFailed returns the AccountsFailed field if non-nil, zero value otherwise.

### GetAccountsFailedOk

`func (o *ListInboxConversations200ResponseMeta) GetAccountsFailedOk() (*int32, bool)`

GetAccountsFailedOk returns a tuple with the AccountsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsFailed

`func (o *ListInboxConversations200ResponseMeta) SetAccountsFailed(v int32)`

SetAccountsFailed sets AccountsFailed field to given value.

### HasAccountsFailed

`func (o *ListInboxConversations200ResponseMeta) HasAccountsFailed() bool`

HasAccountsFailed returns a boolean if a field has been set.

### GetFailedAccounts

`func (o *ListInboxConversations200ResponseMeta) GetFailedAccounts() []ListInboxConversations200ResponseMetaFailedAccountsInner`

GetFailedAccounts returns the FailedAccounts field if non-nil, zero value otherwise.

### GetFailedAccountsOk

`func (o *ListInboxConversations200ResponseMeta) GetFailedAccountsOk() (*[]ListInboxConversations200ResponseMetaFailedAccountsInner, bool)`

GetFailedAccountsOk returns a tuple with the FailedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAccounts

`func (o *ListInboxConversations200ResponseMeta) SetFailedAccounts(v []ListInboxConversations200ResponseMetaFailedAccountsInner)`

SetFailedAccounts sets FailedAccounts field to given value.

### HasFailedAccounts

`func (o *ListInboxConversations200ResponseMeta) HasFailedAccounts() bool`

HasFailedAccounts returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListInboxConversations200ResponseMeta) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListInboxConversations200ResponseMeta) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListInboxConversations200ResponseMeta) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListInboxConversations200ResponseMeta) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


