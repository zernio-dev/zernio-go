# ReviewWebhookReviewReviewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Platform reviewer ID. Null when the platform does not expose it (common on Google Business anonymous reviews). | 
**Name** | **string** |  | 
**ProfileImage** | **NullableString** |  | 

## Methods

### NewReviewWebhookReviewReviewer

`func NewReviewWebhookReviewReviewer(id NullableString, name string, profileImage NullableString, ) *ReviewWebhookReviewReviewer`

NewReviewWebhookReviewReviewer instantiates a new ReviewWebhookReviewReviewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewWebhookReviewReviewerWithDefaults

`func NewReviewWebhookReviewReviewerWithDefaults() *ReviewWebhookReviewReviewer`

NewReviewWebhookReviewReviewerWithDefaults instantiates a new ReviewWebhookReviewReviewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReviewWebhookReviewReviewer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReviewWebhookReviewReviewer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReviewWebhookReviewReviewer) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *ReviewWebhookReviewReviewer) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *ReviewWebhookReviewReviewer) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetName

`func (o *ReviewWebhookReviewReviewer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReviewWebhookReviewReviewer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReviewWebhookReviewReviewer) SetName(v string)`

SetName sets Name field to given value.


### GetProfileImage

`func (o *ReviewWebhookReviewReviewer) GetProfileImage() string`

GetProfileImage returns the ProfileImage field if non-nil, zero value otherwise.

### GetProfileImageOk

`func (o *ReviewWebhookReviewReviewer) GetProfileImageOk() (*string, bool)`

GetProfileImageOk returns a tuple with the ProfileImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileImage

`func (o *ReviewWebhookReviewReviewer) SetProfileImage(v string)`

SetProfileImage sets ProfileImage field to given value.


### SetProfileImageNil

`func (o *ReviewWebhookReviewReviewer) SetProfileImageNil(b bool)`

 SetProfileImageNil sets the value for ProfileImage to be an explicit nil

### UnsetProfileImage
`func (o *ReviewWebhookReviewReviewer) UnsetProfileImage()`

UnsetProfileImage ensures that no value is present for ProfileImage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


