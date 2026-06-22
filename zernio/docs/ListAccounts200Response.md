# ListAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | [**[]SocialAccount**](SocialAccount.md) |  | 
**HasAnalyticsAccess** | **bool** | Whether user has analytics add-on access | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) | Only present when page/limit params are provided | [optional] 

## Methods

### NewListAccounts200Response

`func NewListAccounts200Response(accounts []SocialAccount, hasAnalyticsAccess bool, ) *ListAccounts200Response`

NewListAccounts200Response instantiates a new ListAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccounts200ResponseWithDefaults

`func NewListAccounts200ResponseWithDefaults() *ListAccounts200Response`

NewListAccounts200ResponseWithDefaults instantiates a new ListAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ListAccounts200Response) GetAccounts() []SocialAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListAccounts200Response) GetAccountsOk() (*[]SocialAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListAccounts200Response) SetAccounts(v []SocialAccount)`

SetAccounts sets Accounts field to given value.


### GetHasAnalyticsAccess

`func (o *ListAccounts200Response) GetHasAnalyticsAccess() bool`

GetHasAnalyticsAccess returns the HasAnalyticsAccess field if non-nil, zero value otherwise.

### GetHasAnalyticsAccessOk

`func (o *ListAccounts200Response) GetHasAnalyticsAccessOk() (*bool, bool)`

GetHasAnalyticsAccessOk returns a tuple with the HasAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticsAccess

`func (o *ListAccounts200Response) SetHasAnalyticsAccess(v bool)`

SetHasAnalyticsAccess sets HasAnalyticsAccess field to given value.


### GetPagination

`func (o *ListAccounts200Response) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *ListAccounts200Response) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *ListAccounts200Response) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *ListAccounts200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


