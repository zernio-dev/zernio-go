# WebhookPayloadPostPlatformPlatform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Platform name (e.g. &#x60;twitter&#x60;, &#x60;tiktok&#x60;, &#x60;instagram&#x60;). | 
**Status** | **string** | Terminal status this event fires on. Matches the event suffix. | 
**PlatformPostId** | Pointer to **string** | Platform-native post id. Present on &#x60;published&#x60;, absent on &#x60;failed&#x60;. | [optional] 
**PublishedUrl** | Pointer to **string** | Public URL to the platform-side post. Present on &#x60;published&#x60; (when the platform exposes one and it is not a draft). | [optional] 
**Error** | Pointer to **string** | Error message from the platform. Present on &#x60;failed&#x60;, absent on &#x60;published&#x60;. | [optional] 

## Methods

### NewWebhookPayloadPostPlatformPlatform

`func NewWebhookPayloadPostPlatformPlatform(name string, status string, ) *WebhookPayloadPostPlatformPlatform`

NewWebhookPayloadPostPlatformPlatform instantiates a new WebhookPayloadPostPlatformPlatform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadPostPlatformPlatformWithDefaults

`func NewWebhookPayloadPostPlatformPlatformWithDefaults() *WebhookPayloadPostPlatformPlatform`

NewWebhookPayloadPostPlatformPlatformWithDefaults instantiates a new WebhookPayloadPostPlatformPlatform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *WebhookPayloadPostPlatformPlatform) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookPayloadPostPlatformPlatform) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookPayloadPostPlatformPlatform) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *WebhookPayloadPostPlatformPlatform) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPayloadPostPlatformPlatform) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPayloadPostPlatformPlatform) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatformPostId

`func (o *WebhookPayloadPostPlatformPlatform) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *WebhookPayloadPostPlatformPlatform) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *WebhookPayloadPostPlatformPlatform) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *WebhookPayloadPostPlatformPlatform) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetPublishedUrl

`func (o *WebhookPayloadPostPlatformPlatform) GetPublishedUrl() string`

GetPublishedUrl returns the PublishedUrl field if non-nil, zero value otherwise.

### GetPublishedUrlOk

`func (o *WebhookPayloadPostPlatformPlatform) GetPublishedUrlOk() (*string, bool)`

GetPublishedUrlOk returns a tuple with the PublishedUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedUrl

`func (o *WebhookPayloadPostPlatformPlatform) SetPublishedUrl(v string)`

SetPublishedUrl sets PublishedUrl field to given value.

### HasPublishedUrl

`func (o *WebhookPayloadPostPlatformPlatform) HasPublishedUrl() bool`

HasPublishedUrl returns a boolean if a field has been set.

### GetError

`func (o *WebhookPayloadPostPlatformPlatform) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookPayloadPostPlatformPlatform) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookPayloadPostPlatformPlatform) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookPayloadPostPlatformPlatform) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


