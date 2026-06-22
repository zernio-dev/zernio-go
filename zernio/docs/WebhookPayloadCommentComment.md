# WebhookPayloadCommentComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Platform comment ID | 
**PostId** | **NullableString** | Internal post ID (null for posts not published through Zernio) | 
**PlatformPostId** | **string** | Platform&#39;s post ID | 
**Platform** | **string** |  | 
**Text** | **string** | Comment text content | 
**Author** | [**WebhookPayloadCommentCommentAuthor**](WebhookPayloadCommentCommentAuthor.md) |  | 
**CreatedAt** | **time.Time** |  | 
**IsReply** | **bool** | Whether this is a reply to another comment | 
**ParentCommentId** | **NullableString** | Parent comment ID if this is a reply | 
**Ad** | Pointer to [**WebhookPayloadCommentCommentAd**](WebhookPayloadCommentCommentAd.md) |  | [optional] 

## Methods

### NewWebhookPayloadCommentComment

`func NewWebhookPayloadCommentComment(id string, postId NullableString, platformPostId string, platform string, text string, author WebhookPayloadCommentCommentAuthor, createdAt time.Time, isReply bool, parentCommentId NullableString, ) *WebhookPayloadCommentComment`

NewWebhookPayloadCommentComment instantiates a new WebhookPayloadCommentComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCommentCommentWithDefaults

`func NewWebhookPayloadCommentCommentWithDefaults() *WebhookPayloadCommentComment`

NewWebhookPayloadCommentCommentWithDefaults instantiates a new WebhookPayloadCommentComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCommentComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCommentComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCommentComment) SetId(v string)`

SetId sets Id field to given value.


### GetPostId

`func (o *WebhookPayloadCommentComment) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *WebhookPayloadCommentComment) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *WebhookPayloadCommentComment) SetPostId(v string)`

SetPostId sets PostId field to given value.


### SetPostIdNil

`func (o *WebhookPayloadCommentComment) SetPostIdNil(b bool)`

 SetPostIdNil sets the value for PostId to be an explicit nil

### UnsetPostId
`func (o *WebhookPayloadCommentComment) UnsetPostId()`

UnsetPostId ensures that no value is present for PostId, not even an explicit nil
### GetPlatformPostId

`func (o *WebhookPayloadCommentComment) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *WebhookPayloadCommentComment) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *WebhookPayloadCommentComment) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.


### GetPlatform

`func (o *WebhookPayloadCommentComment) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebhookPayloadCommentComment) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebhookPayloadCommentComment) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetText

`func (o *WebhookPayloadCommentComment) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *WebhookPayloadCommentComment) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *WebhookPayloadCommentComment) SetText(v string)`

SetText sets Text field to given value.


### GetAuthor

`func (o *WebhookPayloadCommentComment) GetAuthor() WebhookPayloadCommentCommentAuthor`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *WebhookPayloadCommentComment) GetAuthorOk() (*WebhookPayloadCommentCommentAuthor, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *WebhookPayloadCommentComment) SetAuthor(v WebhookPayloadCommentCommentAuthor)`

SetAuthor sets Author field to given value.


### GetCreatedAt

`func (o *WebhookPayloadCommentComment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookPayloadCommentComment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookPayloadCommentComment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIsReply

`func (o *WebhookPayloadCommentComment) GetIsReply() bool`

GetIsReply returns the IsReply field if non-nil, zero value otherwise.

### GetIsReplyOk

`func (o *WebhookPayloadCommentComment) GetIsReplyOk() (*bool, bool)`

GetIsReplyOk returns a tuple with the IsReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReply

`func (o *WebhookPayloadCommentComment) SetIsReply(v bool)`

SetIsReply sets IsReply field to given value.


### GetParentCommentId

`func (o *WebhookPayloadCommentComment) GetParentCommentId() string`

GetParentCommentId returns the ParentCommentId field if non-nil, zero value otherwise.

### GetParentCommentIdOk

`func (o *WebhookPayloadCommentComment) GetParentCommentIdOk() (*string, bool)`

GetParentCommentIdOk returns a tuple with the ParentCommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCommentId

`func (o *WebhookPayloadCommentComment) SetParentCommentId(v string)`

SetParentCommentId sets ParentCommentId field to given value.


### SetParentCommentIdNil

`func (o *WebhookPayloadCommentComment) SetParentCommentIdNil(b bool)`

 SetParentCommentIdNil sets the value for ParentCommentId to be an explicit nil

### UnsetParentCommentId
`func (o *WebhookPayloadCommentComment) UnsetParentCommentId()`

UnsetParentCommentId ensures that no value is present for ParentCommentId, not even an explicit nil
### GetAd

`func (o *WebhookPayloadCommentComment) GetAd() WebhookPayloadCommentCommentAd`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *WebhookPayloadCommentComment) GetAdOk() (*WebhookPayloadCommentCommentAd, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *WebhookPayloadCommentComment) SetAd(v WebhookPayloadCommentCommentAd)`

SetAd sets Ad field to given value.

### HasAd

`func (o *WebhookPayloadCommentComment) HasAd() bool`

HasAd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


