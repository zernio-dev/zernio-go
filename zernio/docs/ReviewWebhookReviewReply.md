# ReviewWebhookReviewReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewReviewWebhookReviewReply

`func NewReviewWebhookReviewReply(text string, createdAt time.Time, ) *ReviewWebhookReviewReply`

NewReviewWebhookReviewReply instantiates a new ReviewWebhookReviewReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWebhookReviewReplyWithDefaults

`func NewReviewWebhookReviewReplyWithDefaults() *ReviewWebhookReviewReply`

NewReviewWebhookReviewReplyWithDefaults instantiates a new ReviewWebhookReviewReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ReviewWebhookReviewReply) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReviewWebhookReviewReply) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReviewWebhookReviewReply) SetText(v string)`

SetText sets Text field to given value.


### GetCreatedAt

`func (o *ReviewWebhookReviewReply) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReviewWebhookReviewReply) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReviewWebhookReviewReply) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


