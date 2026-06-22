# GetRedditSubreddits200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subreddits** | Pointer to [**[]GetRedditSubreddits200ResponseSubredditsInner**](GetRedditSubreddits200ResponseSubredditsInner.md) |  | [optional] 
**DefaultSubreddit** | Pointer to **string** | Currently set default subreddit for posting | [optional] 

## Methods

### NewGetRedditSubreddits200Response

`func NewGetRedditSubreddits200Response() *GetRedditSubreddits200Response`

NewGetRedditSubreddits200Response instantiates a new GetRedditSubreddits200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRedditSubreddits200ResponseWithDefaults

`func NewGetRedditSubreddits200ResponseWithDefaults() *GetRedditSubreddits200Response`

NewGetRedditSubreddits200ResponseWithDefaults instantiates a new GetRedditSubreddits200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubreddits

`func (o *GetRedditSubreddits200Response) GetSubreddits() []GetRedditSubreddits200ResponseSubredditsInner`

GetSubreddits returns the Subreddits field if non-nil, zero value otherwise.

### GetSubredditsOk

`func (o *GetRedditSubreddits200Response) GetSubredditsOk() (*[]GetRedditSubreddits200ResponseSubredditsInner, bool)`

GetSubredditsOk returns a tuple with the Subreddits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddits

`func (o *GetRedditSubreddits200Response) SetSubreddits(v []GetRedditSubreddits200ResponseSubredditsInner)`

SetSubreddits sets Subreddits field to given value.

### HasSubreddits

`func (o *GetRedditSubreddits200Response) HasSubreddits() bool`

HasSubreddits returns a boolean if a field has been set.

### GetDefaultSubreddit

`func (o *GetRedditSubreddits200Response) GetDefaultSubreddit() string`

GetDefaultSubreddit returns the DefaultSubreddit field if non-nil, zero value otherwise.

### GetDefaultSubredditOk

`func (o *GetRedditSubreddits200Response) GetDefaultSubredditOk() (*string, bool)`

GetDefaultSubredditOk returns a tuple with the DefaultSubreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSubreddit

`func (o *GetRedditSubreddits200Response) SetDefaultSubreddit(v string)`

SetDefaultSubreddit sets DefaultSubreddit field to given value.

### HasDefaultSubreddit

`func (o *GetRedditSubreddits200Response) HasDefaultSubreddit() bool`

HasDefaultSubreddit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


