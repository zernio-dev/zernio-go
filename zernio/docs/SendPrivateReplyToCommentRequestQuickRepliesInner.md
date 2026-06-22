# SendPrivateReplyToCommentRequestQuickRepliesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Label shown on the chip. Truncated by Meta beyond 20 characters. | 
**Payload** | **string** | Opaque value returned in the inbound webhook when the user taps the chip. | 
**ImageUrl** | Pointer to **string** | Optional thumbnail shown next to the chip title. | [optional] 

## Methods

### NewSendPrivateReplyToCommentRequestQuickRepliesInner

`func NewSendPrivateReplyToCommentRequestQuickRepliesInner(title string, payload string, ) *SendPrivateReplyToCommentRequestQuickRepliesInner`

NewSendPrivateReplyToCommentRequestQuickRepliesInner instantiates a new SendPrivateReplyToCommentRequestQuickRepliesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPrivateReplyToCommentRequestQuickRepliesInnerWithDefaults

`func NewSendPrivateReplyToCommentRequestQuickRepliesInnerWithDefaults() *SendPrivateReplyToCommentRequestQuickRepliesInner`

NewSendPrivateReplyToCommentRequestQuickRepliesInnerWithDefaults instantiates a new SendPrivateReplyToCommentRequestQuickRepliesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPayload

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetImageUrl

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *SendPrivateReplyToCommentRequestQuickRepliesInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


