# GetLinkedInPostReactions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**PostUrn** | Pointer to **string** |  | [optional] 
**Reactions** | Pointer to [**[]GetLinkedInPostReactions200ResponseReactionsInner**](GetLinkedInPostReactions200ResponseReactionsInner.md) |  | [optional] 
**Pagination** | Pointer to [**GetLinkedInPostReactions200ResponsePagination**](GetLinkedInPostReactions200ResponsePagination.md) |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetLinkedInPostReactions200Response

`func NewGetLinkedInPostReactions200Response() *GetLinkedInPostReactions200Response`

NewGetLinkedInPostReactions200Response instantiates a new GetLinkedInPostReactions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkedInPostReactions200ResponseWithDefaults

`func NewGetLinkedInPostReactions200ResponseWithDefaults() *GetLinkedInPostReactions200Response`

NewGetLinkedInPostReactions200ResponseWithDefaults instantiates a new GetLinkedInPostReactions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetLinkedInPostReactions200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetLinkedInPostReactions200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetLinkedInPostReactions200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetLinkedInPostReactions200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetLinkedInPostReactions200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetLinkedInPostReactions200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetLinkedInPostReactions200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetLinkedInPostReactions200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountType

`func (o *GetLinkedInPostReactions200Response) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetLinkedInPostReactions200Response) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetLinkedInPostReactions200Response) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetLinkedInPostReactions200Response) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetUsername

`func (o *GetLinkedInPostReactions200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetLinkedInPostReactions200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetLinkedInPostReactions200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetLinkedInPostReactions200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPostUrn

`func (o *GetLinkedInPostReactions200Response) GetPostUrn() string`

GetPostUrn returns the PostUrn field if non-nil, zero value otherwise.

### GetPostUrnOk

`func (o *GetLinkedInPostReactions200Response) GetPostUrnOk() (*string, bool)`

GetPostUrnOk returns a tuple with the PostUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostUrn

`func (o *GetLinkedInPostReactions200Response) SetPostUrn(v string)`

SetPostUrn sets PostUrn field to given value.

### HasPostUrn

`func (o *GetLinkedInPostReactions200Response) HasPostUrn() bool`

HasPostUrn returns a boolean if a field has been set.

### GetReactions

`func (o *GetLinkedInPostReactions200Response) GetReactions() []GetLinkedInPostReactions200ResponseReactionsInner`

GetReactions returns the Reactions field if non-nil, zero value otherwise.

### GetReactionsOk

`func (o *GetLinkedInPostReactions200Response) GetReactionsOk() (*[]GetLinkedInPostReactions200ResponseReactionsInner, bool)`

GetReactionsOk returns a tuple with the Reactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactions

`func (o *GetLinkedInPostReactions200Response) SetReactions(v []GetLinkedInPostReactions200ResponseReactionsInner)`

SetReactions sets Reactions field to given value.

### HasReactions

`func (o *GetLinkedInPostReactions200Response) HasReactions() bool`

HasReactions returns a boolean if a field has been set.

### GetPagination

`func (o *GetLinkedInPostReactions200Response) GetPagination() GetLinkedInPostReactions200ResponsePagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *GetLinkedInPostReactions200Response) GetPaginationOk() (*GetLinkedInPostReactions200ResponsePagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *GetLinkedInPostReactions200Response) SetPagination(v GetLinkedInPostReactions200ResponsePagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *GetLinkedInPostReactions200Response) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLinkedInPostReactions200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLinkedInPostReactions200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLinkedInPostReactions200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLinkedInPostReactions200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


