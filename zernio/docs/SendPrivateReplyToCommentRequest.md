# SendPrivateReplyToCommentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The social account ID (Instagram or Facebook) | 
**Message** | **string** | The message text to send as a private DM | 
**QuickReplies** | Pointer to [**[]SendPrivateReplyToCommentRequestQuickRepliesInner**](SendPrivateReplyToCommentRequestQuickRepliesInner.md) | Optional quick-reply chips appended to the message. Visible only in the Instagram and Messenger apps (not on web). Maximum 13 entries. Mutually exclusive with &#x60;buttons&#x60;. Note: chips do NOT render in the Instagram Message Requests folder where DMs from non-followers land — use &#x60;buttons&#x60; instead for cold reach.  | [optional] 
**Buttons** | Pointer to [**[]SendPrivateReplyToCommentRequestButtonsInner**](SendPrivateReplyToCommentRequestButtonsInner.md) | Optional 1-3 inline buttons rendered as part of the same message bubble via Meta&#39;s button_template. Visible in the Instagram Message Requests folder (unlike quick replies). Mutually exclusive with &#x60;quickReplies&#x60;.  | [optional] 

## Methods

### NewSendPrivateReplyToCommentRequest

`func NewSendPrivateReplyToCommentRequest(accountId string, message string, ) *SendPrivateReplyToCommentRequest`

NewSendPrivateReplyToCommentRequest instantiates a new SendPrivateReplyToCommentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPrivateReplyToCommentRequestWithDefaults

`func NewSendPrivateReplyToCommentRequestWithDefaults() *SendPrivateReplyToCommentRequest`

NewSendPrivateReplyToCommentRequestWithDefaults instantiates a new SendPrivateReplyToCommentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SendPrivateReplyToCommentRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SendPrivateReplyToCommentRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SendPrivateReplyToCommentRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetMessage

`func (o *SendPrivateReplyToCommentRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendPrivateReplyToCommentRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendPrivateReplyToCommentRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetQuickReplies

`func (o *SendPrivateReplyToCommentRequest) GetQuickReplies() []SendPrivateReplyToCommentRequestQuickRepliesInner`

GetQuickReplies returns the QuickReplies field if non-nil, zero value otherwise.

### GetQuickRepliesOk

`func (o *SendPrivateReplyToCommentRequest) GetQuickRepliesOk() (*[]SendPrivateReplyToCommentRequestQuickRepliesInner, bool)`

GetQuickRepliesOk returns a tuple with the QuickReplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickReplies

`func (o *SendPrivateReplyToCommentRequest) SetQuickReplies(v []SendPrivateReplyToCommentRequestQuickRepliesInner)`

SetQuickReplies sets QuickReplies field to given value.

### HasQuickReplies

`func (o *SendPrivateReplyToCommentRequest) HasQuickReplies() bool`

HasQuickReplies returns a boolean if a field has been set.

### GetButtons

`func (o *SendPrivateReplyToCommentRequest) GetButtons() []SendPrivateReplyToCommentRequestButtonsInner`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *SendPrivateReplyToCommentRequest) GetButtonsOk() (*[]SendPrivateReplyToCommentRequestButtonsInner, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *SendPrivateReplyToCommentRequest) SetButtons(v []SendPrivateReplyToCommentRequestButtonsInner)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *SendPrivateReplyToCommentRequest) HasButtons() bool`

HasButtons returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


