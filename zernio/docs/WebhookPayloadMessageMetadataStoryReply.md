# WebhookPayloadMessageMetadataStoryReply

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoryId** | **string** | The Instagram story ID the user replied to. | 
**StoryUrl** | Pointer to **string** | Meta CDN URL for the story media. Expires approximately 24 hours after the story posted; consumers must fetch promptly or treat 404s as expected.  | [optional] 

## Methods

### NewWebhookPayloadMessageMetadataStoryReply

`func NewWebhookPayloadMessageMetadataStoryReply(storyId string, ) *WebhookPayloadMessageMetadataStoryReply`

NewWebhookPayloadMessageMetadataStoryReply instantiates a new WebhookPayloadMessageMetadataStoryReply object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageMetadataStoryReplyWithDefaults

`func NewWebhookPayloadMessageMetadataStoryReplyWithDefaults() *WebhookPayloadMessageMetadataStoryReply`

NewWebhookPayloadMessageMetadataStoryReplyWithDefaults instantiates a new WebhookPayloadMessageMetadataStoryReply object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoryId

`func (o *WebhookPayloadMessageMetadataStoryReply) GetStoryId() string`

GetStoryId returns the StoryId field if non-nil, zero value otherwise.

### GetStoryIdOk

`func (o *WebhookPayloadMessageMetadataStoryReply) GetStoryIdOk() (*string, bool)`

GetStoryIdOk returns a tuple with the StoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryId

`func (o *WebhookPayloadMessageMetadataStoryReply) SetStoryId(v string)`

SetStoryId sets StoryId field to given value.


### GetStoryUrl

`func (o *WebhookPayloadMessageMetadataStoryReply) GetStoryUrl() string`

GetStoryUrl returns the StoryUrl field if non-nil, zero value otherwise.

### GetStoryUrlOk

`func (o *WebhookPayloadMessageMetadataStoryReply) GetStoryUrlOk() (*string, bool)`

GetStoryUrlOk returns a tuple with the StoryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoryUrl

`func (o *WebhookPayloadMessageMetadataStoryReply) SetStoryUrl(v string)`

SetStoryUrl sets StoryUrl field to given value.

### HasStoryUrl

`func (o *WebhookPayloadMessageMetadataStoryReply) HasStoryUrl() bool`

HasStoryUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


