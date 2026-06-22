# WebhookPayloadPostPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Content** | **string** |  | 
**Status** | **string** |  | 
**ScheduledFor** | **time.Time** |  | 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Platforms** | [**[]WebhookPayloadPostPostPlatformsInner**](WebhookPayloadPostPostPlatformsInner.md) |  | 

## Methods

### NewWebhookPayloadPostPost

`func NewWebhookPayloadPostPost(id string, content string, status string, scheduledFor time.Time, platforms []WebhookPayloadPostPostPlatformsInner, ) *WebhookPayloadPostPost`

NewWebhookPayloadPostPost instantiates a new WebhookPayloadPostPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadPostPostWithDefaults

`func NewWebhookPayloadPostPostWithDefaults() *WebhookPayloadPostPost`

NewWebhookPayloadPostPostWithDefaults instantiates a new WebhookPayloadPostPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadPostPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadPostPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadPostPost) SetId(v string)`

SetId sets Id field to given value.


### GetContent

`func (o *WebhookPayloadPostPost) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebhookPayloadPostPost) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebhookPayloadPostPost) SetContent(v string)`

SetContent sets Content field to given value.


### GetStatus

`func (o *WebhookPayloadPostPost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPayloadPostPost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPayloadPostPost) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetScheduledFor

`func (o *WebhookPayloadPostPost) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *WebhookPayloadPostPost) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *WebhookPayloadPostPost) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.


### GetPublishedAt

`func (o *WebhookPayloadPostPost) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *WebhookPayloadPostPost) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *WebhookPayloadPostPost) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *WebhookPayloadPostPost) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetPlatforms

`func (o *WebhookPayloadPostPost) GetPlatforms() []WebhookPayloadPostPostPlatformsInner`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *WebhookPayloadPostPost) GetPlatformsOk() (*[]WebhookPayloadPostPostPlatformsInner, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *WebhookPayloadPostPost) SetPlatforms(v []WebhookPayloadPostPostPlatformsInner)`

SetPlatforms sets Platforms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


