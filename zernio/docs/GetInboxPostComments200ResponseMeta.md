# GetInboxPostComments200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**PostId** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Subreddit** | Pointer to **NullableString** | (Reddit only) Subreddit name | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**AdComments** | Pointer to [**GetInboxPostComments200ResponseMetaAdComments**](GetInboxPostComments200ResponseMetaAdComments.md) |  | [optional] 

## Methods

### NewGetInboxPostComments200ResponseMeta

`func NewGetInboxPostComments200ResponseMeta() *GetInboxPostComments200ResponseMeta`

NewGetInboxPostComments200ResponseMeta instantiates a new GetInboxPostComments200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxPostComments200ResponseMetaWithDefaults

`func NewGetInboxPostComments200ResponseMetaWithDefaults() *GetInboxPostComments200ResponseMeta`

NewGetInboxPostComments200ResponseMetaWithDefaults instantiates a new GetInboxPostComments200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetInboxPostComments200ResponseMeta) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxPostComments200ResponseMeta) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxPostComments200ResponseMeta) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetInboxPostComments200ResponseMeta) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPostId

`func (o *GetInboxPostComments200ResponseMeta) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *GetInboxPostComments200ResponseMeta) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *GetInboxPostComments200ResponseMeta) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *GetInboxPostComments200ResponseMeta) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetInboxPostComments200ResponseMeta) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInboxPostComments200ResponseMeta) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInboxPostComments200ResponseMeta) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInboxPostComments200ResponseMeta) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetSubreddit

`func (o *GetInboxPostComments200ResponseMeta) GetSubreddit() string`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *GetInboxPostComments200ResponseMeta) GetSubredditOk() (*string, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *GetInboxPostComments200ResponseMeta) SetSubreddit(v string)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *GetInboxPostComments200ResponseMeta) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### SetSubredditNil

`func (o *GetInboxPostComments200ResponseMeta) SetSubredditNil(b bool)`

 SetSubredditNil sets the value for Subreddit to be an explicit nil

### UnsetSubreddit
`func (o *GetInboxPostComments200ResponseMeta) UnsetSubreddit()`

UnsetSubreddit ensures that no value is present for Subreddit, not even an explicit nil
### GetLastUpdated

`func (o *GetInboxPostComments200ResponseMeta) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInboxPostComments200ResponseMeta) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInboxPostComments200ResponseMeta) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInboxPostComments200ResponseMeta) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAdComments

`func (o *GetInboxPostComments200ResponseMeta) GetAdComments() GetInboxPostComments200ResponseMetaAdComments`

GetAdComments returns the AdComments field if non-nil, zero value otherwise.

### GetAdCommentsOk

`func (o *GetInboxPostComments200ResponseMeta) GetAdCommentsOk() (*GetInboxPostComments200ResponseMetaAdComments, bool)`

GetAdCommentsOk returns a tuple with the AdComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComments

`func (o *GetInboxPostComments200ResponseMeta) SetAdComments(v GetInboxPostComments200ResponseMetaAdComments)`

SetAdComments sets AdComments field to given value.

### HasAdComments

`func (o *GetInboxPostComments200ResponseMeta) HasAdComments() bool`

HasAdComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


