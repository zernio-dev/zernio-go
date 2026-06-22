# WebhookPayloadCommentPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **NullableString** | Internal post ID (null for posts not published through Zernio) | 
**PlatformPostId** | **string** | Platform&#39;s post ID | 

## Methods

### NewWebhookPayloadCommentPost

`func NewWebhookPayloadCommentPost(id NullableString, platformPostId string, ) *WebhookPayloadCommentPost`

NewWebhookPayloadCommentPost instantiates a new WebhookPayloadCommentPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCommentPostWithDefaults

`func NewWebhookPayloadCommentPostWithDefaults() *WebhookPayloadCommentPost`

NewWebhookPayloadCommentPostWithDefaults instantiates a new WebhookPayloadCommentPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCommentPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCommentPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCommentPost) SetId(v string)`

SetId sets Id field to given value.


### SetIdNil

`func (o *WebhookPayloadCommentPost) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WebhookPayloadCommentPost) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetPlatformPostId

`func (o *WebhookPayloadCommentPost) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *WebhookPayloadCommentPost) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *WebhookPayloadCommentPost) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


