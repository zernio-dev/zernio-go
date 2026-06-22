# ReviewWebhookReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Platform review ID (e.g. \&quot;accounts/123/locations/456/reviews/789\&quot; for Google Business). | 
**Platform** | **string** | Platform the review originated on. Currently Google Business Profile only. | 
**Rating** | **int32** | Star rating the reviewer gave. | 
**Text** | **string** | Review text content. May be empty if the reviewer left only a rating. | 
**Reviewer** | [**ReviewWebhookReviewReviewer**](ReviewWebhookReviewReviewer.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Replied** | **bool** | Whether the connected account has replied to this review. | 
**Reply** | Pointer to [**ReviewWebhookReviewReply**](ReviewWebhookReviewReply.md) |  | [optional] 

## Methods

### NewReviewWebhookReview

`func NewReviewWebhookReview(id string, platform string, rating int32, text string, reviewer ReviewWebhookReviewReviewer, createdAt time.Time, replied bool, ) *ReviewWebhookReview`

NewReviewWebhookReview instantiates a new ReviewWebhookReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWebhookReviewWithDefaults

`func NewReviewWebhookReviewWithDefaults() *ReviewWebhookReview`

NewReviewWebhookReviewWithDefaults instantiates a new ReviewWebhookReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewWebhookReview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewWebhookReview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewWebhookReview) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ReviewWebhookReview) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ReviewWebhookReview) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ReviewWebhookReview) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetRating

`func (o *ReviewWebhookReview) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ReviewWebhookReview) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ReviewWebhookReview) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetText

`func (o *ReviewWebhookReview) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ReviewWebhookReview) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ReviewWebhookReview) SetText(v string)`

SetText sets Text field to given value.


### GetReviewer

`func (o *ReviewWebhookReview) GetReviewer() ReviewWebhookReviewReviewer`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *ReviewWebhookReview) GetReviewerOk() (*ReviewWebhookReviewReviewer, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *ReviewWebhookReview) SetReviewer(v ReviewWebhookReviewReviewer)`

SetReviewer sets Reviewer field to given value.


### GetCreatedAt

`func (o *ReviewWebhookReview) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReviewWebhookReview) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReviewWebhookReview) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetReplied

`func (o *ReviewWebhookReview) GetReplied() bool`

GetReplied returns the Replied field if non-nil, zero value otherwise.

### GetRepliedOk

`func (o *ReviewWebhookReview) GetRepliedOk() (*bool, bool)`

GetRepliedOk returns a tuple with the Replied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplied

`func (o *ReviewWebhookReview) SetReplied(v bool)`

SetReplied sets Replied field to given value.


### GetReply

`func (o *ReviewWebhookReview) GetReply() ReviewWebhookReviewReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *ReviewWebhookReview) GetReplyOk() (*ReviewWebhookReviewReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *ReviewWebhookReview) SetReply(v ReviewWebhookReviewReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *ReviewWebhookReview) HasReply() bool`

HasReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


