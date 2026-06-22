# AnalyticsListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Overview** | Pointer to [**AnalyticsOverview**](AnalyticsOverview.md) |  | [optional] 
**Posts** | Pointer to [**[]AnalyticsListResponsePostsInner**](AnalyticsListResponsePostsInner.md) |  | [optional] 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Accounts** | Pointer to [**[]SocialAccount**](SocialAccount.md) | Connected social accounts (followerCount and followersLastUpdated only included if user has analytics add-on) | [optional] 
**HasAnalyticsAccess** | Pointer to **bool** | Whether user has analytics add-on access | [optional] 

## Methods

### NewAnalyticsListResponse

`func NewAnalyticsListResponse() *AnalyticsListResponse`

NewAnalyticsListResponse instantiates a new AnalyticsListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsListResponseWithDefaults

`func NewAnalyticsListResponseWithDefaults() *AnalyticsListResponse`

NewAnalyticsListResponseWithDefaults instantiates a new AnalyticsListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOverview

`func (o *AnalyticsListResponse) GetOverview() AnalyticsOverview`

GetOverview returns the Overview field if non-nil, zero value otherwise.

### GetOverviewOk

`func (o *AnalyticsListResponse) GetOverviewOk() (*AnalyticsOverview, bool)`

GetOverviewOk returns a tuple with the Overview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverview

`func (o *AnalyticsListResponse) SetOverview(v AnalyticsOverview)`

SetOverview sets Overview field to given value.

### HasOverview

`func (o *AnalyticsListResponse) HasOverview() bool`

HasOverview returns a boolean if a field has been set.

### GetPosts

`func (o *AnalyticsListResponse) GetPosts() []AnalyticsListResponsePostsInner`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *AnalyticsListResponse) GetPostsOk() (*[]AnalyticsListResponsePostsInner, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *AnalyticsListResponse) SetPosts(v []AnalyticsListResponsePostsInner)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *AnalyticsListResponse) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPagination

`func (o *AnalyticsListResponse) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *AnalyticsListResponse) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *AnalyticsListResponse) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *AnalyticsListResponse) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetAccounts

`func (o *AnalyticsListResponse) GetAccounts() []SocialAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AnalyticsListResponse) GetAccountsOk() (*[]SocialAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AnalyticsListResponse) SetAccounts(v []SocialAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AnalyticsListResponse) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetHasAnalyticsAccess

`func (o *AnalyticsListResponse) GetHasAnalyticsAccess() bool`

GetHasAnalyticsAccess returns the HasAnalyticsAccess field if non-nil, zero value otherwise.

### GetHasAnalyticsAccessOk

`func (o *AnalyticsListResponse) GetHasAnalyticsAccessOk() (*bool, bool)`

GetHasAnalyticsAccessOk returns a tuple with the HasAnalyticsAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAnalyticsAccess

`func (o *AnalyticsListResponse) SetHasAnalyticsAccess(v bool)`

SetHasAnalyticsAccess sets HasAnalyticsAccess field to given value.

### HasHasAnalyticsAccess

`func (o *AnalyticsListResponse) HasHasAnalyticsAccess() bool`

HasHasAnalyticsAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


