# LikeInboxCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The social account ID | 
**Cid** | Pointer to **string** | (Bluesky only) Content identifier for the comment | [optional] 

## Methods

### NewLikeInboxCommentRequest

`func NewLikeInboxCommentRequest(accountId string, ) *LikeInboxCommentRequest`

NewLikeInboxCommentRequest instantiates a new LikeInboxCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLikeInboxCommentRequestWithDefaults

`func NewLikeInboxCommentRequestWithDefaults() *LikeInboxCommentRequest`

NewLikeInboxCommentRequestWithDefaults instantiates a new LikeInboxCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LikeInboxCommentRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LikeInboxCommentRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LikeInboxCommentRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCid

`func (o *LikeInboxCommentRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *LikeInboxCommentRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *LikeInboxCommentRequest) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *LikeInboxCommentRequest) HasCid() bool`

HasCid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


