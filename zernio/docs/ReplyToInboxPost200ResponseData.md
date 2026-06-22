# ReplyToInboxPost200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommentId** | Pointer to **string** |  | [optional] 
**IsReply** | Pointer to **bool** |  | [optional] 
**Cid** | Pointer to **NullableString** | Bluesky CID | [optional] 

## Methods

### NewReplyToInboxPost200ResponseData

`func NewReplyToInboxPost200ResponseData() *ReplyToInboxPost200ResponseData`

NewReplyToInboxPost200ResponseData instantiates a new ReplyToInboxPost200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplyToInboxPost200ResponseDataWithDefaults

`func NewReplyToInboxPost200ResponseDataWithDefaults() *ReplyToInboxPost200ResponseData`

NewReplyToInboxPost200ResponseDataWithDefaults instantiates a new ReplyToInboxPost200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommentId

`func (o *ReplyToInboxPost200ResponseData) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *ReplyToInboxPost200ResponseData) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *ReplyToInboxPost200ResponseData) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *ReplyToInboxPost200ResponseData) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetIsReply

`func (o *ReplyToInboxPost200ResponseData) GetIsReply() bool`

GetIsReply returns the IsReply field if non-nil, zero value otherwise.

### GetIsReplyOk

`func (o *ReplyToInboxPost200ResponseData) GetIsReplyOk() (*bool, bool)`

GetIsReplyOk returns a tuple with the IsReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReply

`func (o *ReplyToInboxPost200ResponseData) SetIsReply(v bool)`

SetIsReply sets IsReply field to given value.

### HasIsReply

`func (o *ReplyToInboxPost200ResponseData) HasIsReply() bool`

HasIsReply returns a boolean if a field has been set.

### GetCid

`func (o *ReplyToInboxPost200ResponseData) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *ReplyToInboxPost200ResponseData) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *ReplyToInboxPost200ResponseData) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *ReplyToInboxPost200ResponseData) HasCid() bool`

HasCid returns a boolean if a field has been set.

### SetCidNil

`func (o *ReplyToInboxPost200ResponseData) SetCidNil(b bool)`

 SetCidNil sets the value for Cid to be an explicit nil

### UnsetCid
`func (o *ReplyToInboxPost200ResponseData) UnsetCid()`

UnsetCid ensures that no value is present for Cid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


