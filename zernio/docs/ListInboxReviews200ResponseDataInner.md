# ListInboxReviews200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**Reviewer** | Pointer to [**ListInboxReviews200ResponseDataInnerReviewer**](ListInboxReviews200ResponseDataInnerReviewer.md) |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Created** | Pointer to **time.Time** |  | [optional] 
**Replied** | Pointer to **bool** |  | [optional] 
**Reply** | Pointer to [**ListInboxReviews200ResponseDataInnerReply**](ListInboxReviews200ResponseDataInnerReply.md) |  | [optional] 
**ReviewUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListInboxReviews200ResponseDataInner

`func NewListInboxReviews200ResponseDataInner() *ListInboxReviews200ResponseDataInner`

NewListInboxReviews200ResponseDataInner instantiates a new ListInboxReviews200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxReviews200ResponseDataInnerWithDefaults

`func NewListInboxReviews200ResponseDataInnerWithDefaults() *ListInboxReviews200ResponseDataInner`

NewListInboxReviews200ResponseDataInnerWithDefaults instantiates a new ListInboxReviews200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInboxReviews200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInboxReviews200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInboxReviews200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListInboxReviews200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListInboxReviews200ResponseDataInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInboxReviews200ResponseDataInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInboxReviews200ResponseDataInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInboxReviews200ResponseDataInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *ListInboxReviews200ResponseDataInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInboxReviews200ResponseDataInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInboxReviews200ResponseDataInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInboxReviews200ResponseDataInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUsername

`func (o *ListInboxReviews200ResponseDataInner) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *ListInboxReviews200ResponseDataInner) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *ListInboxReviews200ResponseDataInner) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *ListInboxReviews200ResponseDataInner) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetReviewer

`func (o *ListInboxReviews200ResponseDataInner) GetReviewer() ListInboxReviews200ResponseDataInnerReviewer`

GetReviewer returns the Reviewer field if non-nil, zero value otherwise.

### GetReviewerOk

`func (o *ListInboxReviews200ResponseDataInner) GetReviewerOk() (*ListInboxReviews200ResponseDataInnerReviewer, bool)`

GetReviewerOk returns a tuple with the Reviewer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewer

`func (o *ListInboxReviews200ResponseDataInner) SetReviewer(v ListInboxReviews200ResponseDataInnerReviewer)`

SetReviewer sets Reviewer field to given value.

### HasReviewer

`func (o *ListInboxReviews200ResponseDataInner) HasReviewer() bool`

HasReviewer returns a boolean if a field has been set.

### GetRating

`func (o *ListInboxReviews200ResponseDataInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ListInboxReviews200ResponseDataInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ListInboxReviews200ResponseDataInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ListInboxReviews200ResponseDataInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetText

`func (o *ListInboxReviews200ResponseDataInner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ListInboxReviews200ResponseDataInner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ListInboxReviews200ResponseDataInner) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *ListInboxReviews200ResponseDataInner) HasText() bool`

HasText returns a boolean if a field has been set.

### GetCreated

`func (o *ListInboxReviews200ResponseDataInner) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ListInboxReviews200ResponseDataInner) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ListInboxReviews200ResponseDataInner) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ListInboxReviews200ResponseDataInner) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetReplied

`func (o *ListInboxReviews200ResponseDataInner) GetReplied() bool`

GetReplied returns the Replied field if non-nil, zero value otherwise.

### GetRepliedOk

`func (o *ListInboxReviews200ResponseDataInner) GetRepliedOk() (*bool, bool)`

GetRepliedOk returns a tuple with the Replied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplied

`func (o *ListInboxReviews200ResponseDataInner) SetReplied(v bool)`

SetReplied sets Replied field to given value.

### HasReplied

`func (o *ListInboxReviews200ResponseDataInner) HasReplied() bool`

HasReplied returns a boolean if a field has been set.

### GetReply

`func (o *ListInboxReviews200ResponseDataInner) GetReply() ListInboxReviews200ResponseDataInnerReply`

GetReply returns the Reply field if non-nil, zero value otherwise.

### GetReplyOk

`func (o *ListInboxReviews200ResponseDataInner) GetReplyOk() (*ListInboxReviews200ResponseDataInnerReply, bool)`

GetReplyOk returns a tuple with the Reply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReply

`func (o *ListInboxReviews200ResponseDataInner) SetReply(v ListInboxReviews200ResponseDataInnerReply)`

SetReply sets Reply field to given value.

### HasReply

`func (o *ListInboxReviews200ResponseDataInner) HasReply() bool`

HasReply returns a boolean if a field has been set.

### GetReviewUrl

`func (o *ListInboxReviews200ResponseDataInner) GetReviewUrl() string`

GetReviewUrl returns the ReviewUrl field if non-nil, zero value otherwise.

### GetReviewUrlOk

`func (o *ListInboxReviews200ResponseDataInner) GetReviewUrlOk() (*string, bool)`

GetReviewUrlOk returns a tuple with the ReviewUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewUrl

`func (o *ListInboxReviews200ResponseDataInner) SetReviewUrl(v string)`

SetReviewUrl sets ReviewUrl field to given value.

### HasReviewUrl

`func (o *ListInboxReviews200ResponseDataInner) HasReviewUrl() bool`

HasReviewUrl returns a boolean if a field has been set.

### SetReviewUrlNil

`func (o *ListInboxReviews200ResponseDataInner) SetReviewUrlNil(b bool)`

 SetReviewUrlNil sets the value for ReviewUrl to be an explicit nil

### UnsetReviewUrl
`func (o *ListInboxReviews200ResponseDataInner) UnsetReviewUrl()`

UnsetReviewUrl ensures that no value is present for ReviewUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


