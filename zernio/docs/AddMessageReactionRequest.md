# AddMessageReactionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Social account ID | 
**Emoji** | **string** | Emoji character (e.g. \&quot;👍\&quot;, \&quot;❤️\&quot;) | 

## Methods

### NewAddMessageReactionRequest

`func NewAddMessageReactionRequest(accountId string, emoji string, ) *AddMessageReactionRequest`

NewAddMessageReactionRequest instantiates a new AddMessageReactionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddMessageReactionRequestWithDefaults

`func NewAddMessageReactionRequestWithDefaults() *AddMessageReactionRequest`

NewAddMessageReactionRequestWithDefaults instantiates a new AddMessageReactionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *AddMessageReactionRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddMessageReactionRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddMessageReactionRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEmoji

`func (o *AddMessageReactionRequest) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *AddMessageReactionRequest) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *AddMessageReactionRequest) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


