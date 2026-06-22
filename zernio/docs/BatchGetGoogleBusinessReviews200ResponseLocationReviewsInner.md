# BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Full review resource name (accounts/_*_/locations/_*_/reviews/_*) | [optional] 
**Review** | Pointer to **map[string]interface{}** | The review object (reviewId, starRating, comment, reviewer, createTime, updateTime, reviewReply) | [optional] 

## Methods

### NewBatchGetGoogleBusinessReviews200ResponseLocationReviewsInner

`func NewBatchGetGoogleBusinessReviews200ResponseLocationReviewsInner() *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner`

NewBatchGetGoogleBusinessReviews200ResponseLocationReviewsInner instantiates a new BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchGetGoogleBusinessReviews200ResponseLocationReviewsInnerWithDefaults

`func NewBatchGetGoogleBusinessReviews200ResponseLocationReviewsInnerWithDefaults() *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner`

NewBatchGetGoogleBusinessReviews200ResponseLocationReviewsInnerWithDefaults instantiates a new BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReview

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) GetReview() map[string]interface{}`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) GetReviewOk() (*map[string]interface{}, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) SetReview(v map[string]interface{})`

SetReview sets Review field to given value.

### HasReview

`func (o *BatchGetGoogleBusinessReviews200ResponseLocationReviewsInner) HasReview() bool`

HasReview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


