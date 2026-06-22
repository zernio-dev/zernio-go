# GetGoogleBusinessReviews200ResponseReviewsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Review ID | [optional] 
**Name** | Pointer to **string** | Full resource name | [optional] 
**Reviewer** | Pointer to [**GetGoogleBusinessReviews200ResponseReviewsInnerReviewer**](GetGoogleBusinessReviews200ResponseReviewsInnerReviewer.md) |  | [optional] 
**Rating** | Pointer to **int32** | Numeric star rating | [optional] 
**StarRating** | Pointer to **string** | Google&#39;s string rating | [optional] 
**Comment** | Pointer to **string** | Review text | [optional] 
**CreateTime** | Pointer to **time.Time** |  | [optional] 
**UpdateTime** | Pointer to **time.Time** |  | [optional] 
**ReviewReply** | Pointer to [**GetGoogleBusinessReviews200ResponseReviewsInnerReviewReply**](GetGoogleBusinessReviews200ResponseReviewsInnerReviewReply.md) |  | [optional] 

## Methods

### NewGetGoogleBusinessReviews200ResponseReviewsInner

`func NewGetGoogleBusinessReviews200ResponseReviewsInner() *GetGoogleBusinessReviews200ResponseReviewsInner`

NewGetGoogleBusinessReviews200ResponseReviewsInner instantiates a new GetGoogleBusinessReviews200ResponseReviewsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessReviews200ResponseReviewsInnerWithDefaults

`func NewGetGoogleBusinessReviews200ResponseReviewsInnerWithDefaults() *GetGoogleBusinessReviews200ResponseReviewsInner`

NewGetGoogleBusinessReviews200ResponseReviewsInnerWithDefaults instantiates a new GetGoogleBusinessReviews200ResponseReviewsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReviewer

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetReviewer() GetGoogleBusinessReviews200ResponseReviewsInnerReviewer`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetReviewerOk() (*GetGoogleBusinessReviews200ResponseReviewsInnerReviewer, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetReviewer(v GetGoogleBusinessReviews200ResponseReviewsInnerReviewer)`

SetReviewer sets Reviewer field to given value.

### HasReviewer

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasReviewer() bool`

HasReviewer returns a boolean if a field has been set.

### GetRating

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetStarRating

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetStarRating() string`

GetStarRating returns the StarRating field if non-nil, zero value otherwise.

### GetStarRatingOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetStarRatingOk() (*string, bool)`

GetStarRatingOk returns a tuple with the StarRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarRating

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetStarRating(v string)`

SetStarRating sets StarRating field to given value.

### HasStarRating

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasStarRating() bool`

HasStarRating returns a boolean if a field has been set.

### GetComment

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreateTime

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetCreateTime() time.Time`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetCreateTimeOk() (*time.Time, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetCreateTime(v time.Time)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetUpdateTime

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetReviewReply

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetReviewReply() GetGoogleBusinessReviews200ResponseReviewsInnerReviewReply`

GetReviewReply returns the ReviewReply field if non-nil, zero value otherwise.

### GetReviewReplyOk

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) GetReviewReplyOk() (*GetGoogleBusinessReviews200ResponseReviewsInnerReviewReply, bool)`

GetReviewReplyOk returns a tuple with the ReviewReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewReply

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) SetReviewReply(v GetGoogleBusinessReviews200ResponseReviewsInnerReviewReply)`

SetReviewReply sets ReviewReply field to given value.

### HasReviewReply

`func (o *GetGoogleBusinessReviews200ResponseReviewsInner) HasReviewReply() bool`

HasReviewReply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


