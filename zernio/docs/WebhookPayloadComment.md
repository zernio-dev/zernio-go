# WebhookPayloadComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Comment** | [**WebhookPayloadCommentComment**](WebhookPayloadCommentComment.md) |  | 
**Post** | [**WebhookPayloadCommentPost**](WebhookPayloadCommentPost.md) |  | 
**Account** | [**WebhookPayloadCommentAccount**](WebhookPayloadCommentAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadComment

`func NewWebhookPayloadComment(id string, event string, comment WebhookPayloadCommentComment, post WebhookPayloadCommentPost, account WebhookPayloadCommentAccount, timestamp time.Time, ) *WebhookPayloadComment`

NewWebhookPayloadComment instantiates a new WebhookPayloadComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCommentWithDefaults

`func NewWebhookPayloadCommentWithDefaults() *WebhookPayloadComment`

NewWebhookPayloadCommentWithDefaults instantiates a new WebhookPayloadComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadComment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadComment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadComment) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadComment) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadComment) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadComment) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetComment

`func (o *WebhookPayloadComment) GetComment() WebhookPayloadCommentComment`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WebhookPayloadComment) GetCommentOk() (*WebhookPayloadCommentComment, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WebhookPayloadComment) SetComment(v WebhookPayloadCommentComment)`

SetComment sets Comment field to given value.


### GetPost

`func (o *WebhookPayloadComment) GetPost() WebhookPayloadCommentPost`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *WebhookPayloadComment) GetPostOk() (*WebhookPayloadCommentPost, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *WebhookPayloadComment) SetPost(v WebhookPayloadCommentPost)`

SetPost sets Post field to given value.


### GetAccount

`func (o *WebhookPayloadComment) GetAccount() WebhookPayloadCommentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadComment) GetAccountOk() (*WebhookPayloadCommentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadComment) SetAccount(v WebhookPayloadCommentAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadComment) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadComment) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadComment) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


