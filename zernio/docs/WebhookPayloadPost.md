# WebhookPayloadPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Post** | [**WebhookPayloadPostPost**](WebhookPayloadPostPost.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadPost

`func NewWebhookPayloadPost(id string, event string, post WebhookPayloadPostPost, timestamp time.Time, ) *WebhookPayloadPost`

NewWebhookPayloadPost instantiates a new WebhookPayloadPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadPostWithDefaults

`func NewWebhookPayloadPostWithDefaults() *WebhookPayloadPost`

NewWebhookPayloadPostWithDefaults instantiates a new WebhookPayloadPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadPost) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadPost) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadPost) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadPost) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetPost

`func (o *WebhookPayloadPost) GetPost() WebhookPayloadPostPost`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *WebhookPayloadPost) GetPostOk() (*WebhookPayloadPostPost, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *WebhookPayloadPost) SetPost(v WebhookPayloadPostPost)`

SetPost sets Post field to given value.


### GetTimestamp

`func (o *WebhookPayloadPost) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadPost) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadPost) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


