# RetweetPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The social account ID | 
**TweetId** | **string** | The ID of the tweet to retweet | 

## Methods

### NewRetweetPostRequest

`func NewRetweetPostRequest(accountId string, tweetId string, ) *RetweetPostRequest`

NewRetweetPostRequest instantiates a new RetweetPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRetweetPostRequestWithDefaults

`func NewRetweetPostRequestWithDefaults() *RetweetPostRequest`

NewRetweetPostRequestWithDefaults instantiates a new RetweetPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *RetweetPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *RetweetPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *RetweetPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTweetId

`func (o *RetweetPostRequest) GetTweetId() string`

GetTweetId returns the TweetId field if non-nil, zero value otherwise.

### GetTweetIdOk

`func (o *RetweetPostRequest) GetTweetIdOk() (*string, bool)`

GetTweetIdOk returns a tuple with the TweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTweetId

`func (o *RetweetPostRequest) SetTweetId(v string)`

SetTweetId sets TweetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


