# WebhookPayloadReviewUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Review** | [**ReviewWebhookReview**](ReviewWebhookReview.md) |  | 
**Account** | [**WebhookPayloadCommentAccount**](WebhookPayloadCommentAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadReviewUpdated

`func NewWebhookPayloadReviewUpdated(id string, event string, review ReviewWebhookReview, account WebhookPayloadCommentAccount, timestamp time.Time, ) *WebhookPayloadReviewUpdated`

NewWebhookPayloadReviewUpdated instantiates a new WebhookPayloadReviewUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadReviewUpdatedWithDefaults

`func NewWebhookPayloadReviewUpdatedWithDefaults() *WebhookPayloadReviewUpdated`

NewWebhookPayloadReviewUpdatedWithDefaults instantiates a new WebhookPayloadReviewUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadReviewUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadReviewUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadReviewUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadReviewUpdated) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadReviewUpdated) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadReviewUpdated) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetReview

`func (o *WebhookPayloadReviewUpdated) GetReview() ReviewWebhookReview`

GetReview returns the Review field if non-nil, zero value otherwise.

### GetReviewOk

`func (o *WebhookPayloadReviewUpdated) GetReviewOk() (*ReviewWebhookReview, bool)`

GetReviewOk returns a tuple with the Review field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReview

`func (o *WebhookPayloadReviewUpdated) SetReview(v ReviewWebhookReview)`

SetReview sets Review field to given value.


### GetAccount

`func (o *WebhookPayloadReviewUpdated) GetAccount() WebhookPayloadCommentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadReviewUpdated) GetAccountOk() (*WebhookPayloadCommentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadReviewUpdated) SetAccount(v WebhookPayloadCommentAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadReviewUpdated) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadReviewUpdated) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadReviewUpdated) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


