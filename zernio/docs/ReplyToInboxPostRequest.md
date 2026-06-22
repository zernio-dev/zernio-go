# ReplyToInboxPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Message** | **string** |  | 
**CommentId** | Pointer to **string** | Reply to specific comment (optional) | [optional] 
**ParentCid** | Pointer to **string** | (Bluesky only) Parent content identifier | [optional] 
**RootUri** | Pointer to **string** | (Bluesky only) Root post URI | [optional] 
**RootCid** | Pointer to **string** | (Bluesky only) Root post CID | [optional] 

## Methods

### NewReplyToInboxPostRequest

`func NewReplyToInboxPostRequest(accountId string, message string, ) *ReplyToInboxPostRequest`

NewReplyToInboxPostRequest instantiates a new ReplyToInboxPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyToInboxPostRequestWithDefaults

`func NewReplyToInboxPostRequestWithDefaults() *ReplyToInboxPostRequest`

NewReplyToInboxPostRequestWithDefaults instantiates a new ReplyToInboxPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *ReplyToInboxPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ReplyToInboxPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ReplyToInboxPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetMessage

`func (o *ReplyToInboxPostRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ReplyToInboxPostRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ReplyToInboxPostRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCommentId

`func (o *ReplyToInboxPostRequest) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *ReplyToInboxPostRequest) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *ReplyToInboxPostRequest) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *ReplyToInboxPostRequest) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetParentCid

`func (o *ReplyToInboxPostRequest) GetParentCid() string`

GetParentCid returns the ParentCid field if non-nil, zero value otherwise.

### GetParentCidOk

`func (o *ReplyToInboxPostRequest) GetParentCidOk() (*string, bool)`

GetParentCidOk returns a tuple with the ParentCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCid

`func (o *ReplyToInboxPostRequest) SetParentCid(v string)`

SetParentCid sets ParentCid field to given value.

### HasParentCid

`func (o *ReplyToInboxPostRequest) HasParentCid() bool`

HasParentCid returns a boolean if a field has been set.

### GetRootUri

`func (o *ReplyToInboxPostRequest) GetRootUri() string`

GetRootUri returns the RootUri field if non-nil, zero value otherwise.

### GetRootUriOk

`func (o *ReplyToInboxPostRequest) GetRootUriOk() (*string, bool)`

GetRootUriOk returns a tuple with the RootUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUri

`func (o *ReplyToInboxPostRequest) SetRootUri(v string)`

SetRootUri sets RootUri field to given value.

### HasRootUri

`func (o *ReplyToInboxPostRequest) HasRootUri() bool`

HasRootUri returns a boolean if a field has been set.

### GetRootCid

`func (o *ReplyToInboxPostRequest) GetRootCid() string`

GetRootCid returns the RootCid field if non-nil, zero value otherwise.

### GetRootCidOk

`func (o *ReplyToInboxPostRequest) GetRootCidOk() (*string, bool)`

GetRootCidOk returns a tuple with the RootCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCid

`func (o *ReplyToInboxPostRequest) SetRootCid(v string)`

SetRootCid sets RootCid field to given value.

### HasRootCid

`func (o *ReplyToInboxPostRequest) HasRootCid() bool`

HasRootCid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


