# GetInboxPostComments200ResponseCommentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**From** | Pointer to [**GetInboxPostComments200ResponseCommentsInnerFrom**](GetInboxPostComments200ResponseCommentsInnerFrom.md) |  | [optional] 
**LikeCount** | Pointer to **int32** |  | [optional] 
**ReplyCount** | Pointer to **int32** |  | [optional] 
**Platform** | Pointer to **string** | The platform this comment is from | [optional] 
**Url** | Pointer to **NullableString** | Direct link to the comment on the platform (if available) | [optional] 
**Replies** | Pointer to **[]map[string]interface{}** |  | [optional] 
**CanReply** | Pointer to **bool** |  | [optional] 
**CanDelete** | Pointer to **bool** |  | [optional] 
**CanHide** | Pointer to **bool** | Whether this comment can be hidden (Facebook, Instagram, Threads) | [optional] 
**CanLike** | Pointer to **bool** | Whether this comment can be liked (Facebook, Twitter/X, Bluesky, Reddit) | [optional] 
**IsHidden** | Pointer to **bool** | Whether the comment is currently hidden | [optional] 
**IsLiked** | Pointer to **bool** | Whether the current user has liked this comment | [optional] 
**LikeUri** | Pointer to **NullableString** | Bluesky like URI for unliking | [optional] 
**Cid** | Pointer to **NullableString** | Bluesky content identifier | [optional] 
**ParentId** | Pointer to **NullableString** | Parent comment ID for nested replies | [optional] 
**RootUri** | Pointer to **NullableString** | Bluesky root post URI | [optional] 
**RootCid** | Pointer to **NullableString** | Bluesky root post CID | [optional] 

## Methods

### NewGetInboxPostComments200ResponseCommentsInner

`func NewGetInboxPostComments200ResponseCommentsInner() *GetInboxPostComments200ResponseCommentsInner`

NewGetInboxPostComments200ResponseCommentsInner instantiates a new GetInboxPostComments200ResponseCommentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxPostComments200ResponseCommentsInnerWithDefaults

`func NewGetInboxPostComments200ResponseCommentsInnerWithDefaults() *GetInboxPostComments200ResponseCommentsInner`

NewGetInboxPostComments200ResponseCommentsInnerWithDefaults instantiates a new GetInboxPostComments200ResponseCommentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInboxPostComments200ResponseCommentsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxPostComments200ResponseCommentsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInboxPostComments200ResponseCommentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *GetInboxPostComments200ResponseCommentsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetInboxPostComments200ResponseCommentsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetInboxPostComments200ResponseCommentsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCreatedTime

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *GetInboxPostComments200ResponseCommentsInner) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetFrom

`func (o *GetInboxPostComments200ResponseCommentsInner) GetFrom() GetInboxPostComments200ResponseCommentsInnerFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetFromOk() (*GetInboxPostComments200ResponseCommentsInnerFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxPostComments200ResponseCommentsInner) SetFrom(v GetInboxPostComments200ResponseCommentsInnerFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxPostComments200ResponseCommentsInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLikeCount

`func (o *GetInboxPostComments200ResponseCommentsInner) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *GetInboxPostComments200ResponseCommentsInner) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *GetInboxPostComments200ResponseCommentsInner) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetReplyCount

`func (o *GetInboxPostComments200ResponseCommentsInner) GetReplyCount() int32`

GetReplyCount returns the ReplyCount field if non-nil, zero value otherwise.

### GetReplyCountOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetReplyCountOk() (*int32, bool)`

GetReplyCountOk returns a tuple with the ReplyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyCount

`func (o *GetInboxPostComments200ResponseCommentsInner) SetReplyCount(v int32)`

SetReplyCount sets ReplyCount field to given value.

### HasReplyCount

`func (o *GetInboxPostComments200ResponseCommentsInner) HasReplyCount() bool`

HasReplyCount returns a boolean if a field has been set.

### GetPlatform

`func (o *GetInboxPostComments200ResponseCommentsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxPostComments200ResponseCommentsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetInboxPostComments200ResponseCommentsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUrl

`func (o *GetInboxPostComments200ResponseCommentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetInboxPostComments200ResponseCommentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetInboxPostComments200ResponseCommentsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *GetInboxPostComments200ResponseCommentsInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *GetInboxPostComments200ResponseCommentsInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetReplies

`func (o *GetInboxPostComments200ResponseCommentsInner) GetReplies() []map[string]interface{}`

GetReplies returns the Replies field if non-nil, zero value otherwise.

### GetRepliesOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetRepliesOk() (*[]map[string]interface{}, bool)`

GetRepliesOk returns a tuple with the Replies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplies

`func (o *GetInboxPostComments200ResponseCommentsInner) SetReplies(v []map[string]interface{})`

SetReplies sets Replies field to given value.

### HasReplies

`func (o *GetInboxPostComments200ResponseCommentsInner) HasReplies() bool`

HasReplies returns a boolean if a field has been set.

### GetCanReply

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanReply() bool`

GetCanReply returns the CanReply field if non-nil, zero value otherwise.

### GetCanReplyOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanReplyOk() (*bool, bool)`

GetCanReplyOk returns a tuple with the CanReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanReply

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCanReply(v bool)`

SetCanReply sets CanReply field to given value.

### HasCanReply

`func (o *GetInboxPostComments200ResponseCommentsInner) HasCanReply() bool`

HasCanReply returns a boolean if a field has been set.

### GetCanDelete

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanDelete() bool`

GetCanDelete returns the CanDelete field if non-nil, zero value otherwise.

### GetCanDeleteOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanDeleteOk() (*bool, bool)`

GetCanDeleteOk returns a tuple with the CanDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanDelete

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCanDelete(v bool)`

SetCanDelete sets CanDelete field to given value.

### HasCanDelete

`func (o *GetInboxPostComments200ResponseCommentsInner) HasCanDelete() bool`

HasCanDelete returns a boolean if a field has been set.

### GetCanHide

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanHide() bool`

GetCanHide returns the CanHide field if non-nil, zero value otherwise.

### GetCanHideOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanHideOk() (*bool, bool)`

GetCanHideOk returns a tuple with the CanHide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanHide

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCanHide(v bool)`

SetCanHide sets CanHide field to given value.

### HasCanHide

`func (o *GetInboxPostComments200ResponseCommentsInner) HasCanHide() bool`

HasCanHide returns a boolean if a field has been set.

### GetCanLike

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanLike() bool`

GetCanLike returns the CanLike field if non-nil, zero value otherwise.

### GetCanLikeOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCanLikeOk() (*bool, bool)`

GetCanLikeOk returns a tuple with the CanLike field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanLike

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCanLike(v bool)`

SetCanLike sets CanLike field to given value.

### HasCanLike

`func (o *GetInboxPostComments200ResponseCommentsInner) HasCanLike() bool`

HasCanLike returns a boolean if a field has been set.

### GetIsHidden

`func (o *GetInboxPostComments200ResponseCommentsInner) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *GetInboxPostComments200ResponseCommentsInner) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *GetInboxPostComments200ResponseCommentsInner) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsLiked

`func (o *GetInboxPostComments200ResponseCommentsInner) GetIsLiked() bool`

GetIsLiked returns the IsLiked field if non-nil, zero value otherwise.

### GetIsLikedOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetIsLikedOk() (*bool, bool)`

GetIsLikedOk returns a tuple with the IsLiked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLiked

`func (o *GetInboxPostComments200ResponseCommentsInner) SetIsLiked(v bool)`

SetIsLiked sets IsLiked field to given value.

### HasIsLiked

`func (o *GetInboxPostComments200ResponseCommentsInner) HasIsLiked() bool`

HasIsLiked returns a boolean if a field has been set.

### GetLikeUri

`func (o *GetInboxPostComments200ResponseCommentsInner) GetLikeUri() string`

GetLikeUri returns the LikeUri field if non-nil, zero value otherwise.

### GetLikeUriOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetLikeUriOk() (*string, bool)`

GetLikeUriOk returns a tuple with the LikeUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeUri

`func (o *GetInboxPostComments200ResponseCommentsInner) SetLikeUri(v string)`

SetLikeUri sets LikeUri field to given value.

### HasLikeUri

`func (o *GetInboxPostComments200ResponseCommentsInner) HasLikeUri() bool`

HasLikeUri returns a boolean if a field has been set.

### SetLikeUriNil

`func (o *GetInboxPostComments200ResponseCommentsInner) SetLikeUriNil(b bool)`

 SetLikeUriNil sets the value for LikeUri to be an explicit nil

### UnsetLikeUri
`func (o *GetInboxPostComments200ResponseCommentsInner) UnsetLikeUri()`

UnsetLikeUri ensures that no value is present for LikeUri, not even an explicit nil
### GetCid

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *GetInboxPostComments200ResponseCommentsInner) HasCid() bool`

HasCid returns a boolean if a field has been set.

### SetCidNil

`func (o *GetInboxPostComments200ResponseCommentsInner) SetCidNil(b bool)`

 SetCidNil sets the value for Cid to be an explicit nil

### UnsetCid
`func (o *GetInboxPostComments200ResponseCommentsInner) UnsetCid()`

UnsetCid ensures that no value is present for Cid, not even an explicit nil
### GetParentId

`func (o *GetInboxPostComments200ResponseCommentsInner) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetInboxPostComments200ResponseCommentsInner) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetInboxPostComments200ResponseCommentsInner) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *GetInboxPostComments200ResponseCommentsInner) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *GetInboxPostComments200ResponseCommentsInner) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetRootUri

`func (o *GetInboxPostComments200ResponseCommentsInner) GetRootUri() string`

GetRootUri returns the RootUri field if non-nil, zero value otherwise.

### GetRootUriOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetRootUriOk() (*string, bool)`

GetRootUriOk returns a tuple with the RootUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootUri

`func (o *GetInboxPostComments200ResponseCommentsInner) SetRootUri(v string)`

SetRootUri sets RootUri field to given value.

### HasRootUri

`func (o *GetInboxPostComments200ResponseCommentsInner) HasRootUri() bool`

HasRootUri returns a boolean if a field has been set.

### SetRootUriNil

`func (o *GetInboxPostComments200ResponseCommentsInner) SetRootUriNil(b bool)`

 SetRootUriNil sets the value for RootUri to be an explicit nil

### UnsetRootUri
`func (o *GetInboxPostComments200ResponseCommentsInner) UnsetRootUri()`

UnsetRootUri ensures that no value is present for RootUri, not even an explicit nil
### GetRootCid

`func (o *GetInboxPostComments200ResponseCommentsInner) GetRootCid() string`

GetRootCid returns the RootCid field if non-nil, zero value otherwise.

### GetRootCidOk

`func (o *GetInboxPostComments200ResponseCommentsInner) GetRootCidOk() (*string, bool)`

GetRootCidOk returns a tuple with the RootCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootCid

`func (o *GetInboxPostComments200ResponseCommentsInner) SetRootCid(v string)`

SetRootCid sets RootCid field to given value.

### HasRootCid

`func (o *GetInboxPostComments200ResponseCommentsInner) HasRootCid() bool`

HasRootCid returns a boolean if a field has been set.

### SetRootCidNil

`func (o *GetInboxPostComments200ResponseCommentsInner) SetRootCidNil(b bool)`

 SetRootCidNil sets the value for RootCid to be an explicit nil

### UnsetRootCid
`func (o *GetInboxPostComments200ResponseCommentsInner) UnsetRootCid()`

UnsetRootCid ensures that no value is present for RootCid, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


