# ListInboxConversations200ResponseMetaFailedAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **NullableString** | Error code if available | [optional] 
**RetryAfter** | Pointer to **NullableInt32** | Seconds to wait before retry (rate limits) | [optional] 

## Methods

### NewListInboxConversations200ResponseMetaFailedAccountsInner

`func NewListInboxConversations200ResponseMetaFailedAccountsInner() *ListInboxConversations200ResponseMetaFailedAccountsInner`

NewListInboxConversations200ResponseMetaFailedAccountsInner instantiates a new ListInboxConversations200ResponseMetaFailedAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversations200ResponseMetaFailedAccountsInnerWithDefaults

`func NewListInboxConversations200ResponseMetaFailedAccountsInnerWithDefaults() *ListInboxConversations200ResponseMetaFailedAccountsInner`

NewListInboxConversations200ResponseMetaFailedAccountsInnerWithDefaults instantiates a new ListInboxConversations200ResponseMetaFailedAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUsername

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### SetAccountUsernameNil

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetAccountUsernameNil(b bool)`

 SetAccountUsernameNil sets the value for AccountUsername to be an explicit nil

### UnsetAccountUsername
`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) UnsetAccountUsername()`

UnsetAccountUsername ensures that no value is present for AccountUsername, not even an explicit nil
### GetPlatform

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetError

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCode

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### SetCodeNil

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetCodeNil(b bool)`

 SetCodeNil sets the value for Code to be an explicit nil

### UnsetCode
`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) UnsetCode()`

UnsetCode ensures that no value is present for Code, not even an explicit nil
### GetRetryAfter

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetRetryAfter() int32`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) GetRetryAfterOk() (*int32, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetRetryAfter(v int32)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.

### SetRetryAfterNil

`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) SetRetryAfterNil(b bool)`

 SetRetryAfterNil sets the value for RetryAfter to be an explicit nil

### UnsetRetryAfter
`func (o *ListInboxConversations200ResponseMetaFailedAccountsInner) UnsetRetryAfter()`

UnsetRetryAfter ensures that no value is present for RetryAfter, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


