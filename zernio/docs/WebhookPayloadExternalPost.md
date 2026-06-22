# WebhookPayloadExternalPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Post** | [**ExternalPostWebhookPost**](ExternalPostWebhookPost.md) |  | 
**Account** | [**WebhookPayloadCommentAccount**](WebhookPayloadCommentAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadExternalPost

`func NewWebhookPayloadExternalPost(id string, event string, post ExternalPostWebhookPost, account WebhookPayloadCommentAccount, timestamp time.Time, ) *WebhookPayloadExternalPost`

NewWebhookPayloadExternalPost instantiates a new WebhookPayloadExternalPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadExternalPostWithDefaults

`func NewWebhookPayloadExternalPostWithDefaults() *WebhookPayloadExternalPost`

NewWebhookPayloadExternalPostWithDefaults instantiates a new WebhookPayloadExternalPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadExternalPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadExternalPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadExternalPost) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadExternalPost) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadExternalPost) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadExternalPost) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetPost

`func (o *WebhookPayloadExternalPost) GetPost() ExternalPostWebhookPost`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *WebhookPayloadExternalPost) GetPostOk() (*ExternalPostWebhookPost, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *WebhookPayloadExternalPost) SetPost(v ExternalPostWebhookPost)`

SetPost sets Post field to given value.


### GetAccount

`func (o *WebhookPayloadExternalPost) GetAccount() WebhookPayloadCommentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadExternalPost) GetAccountOk() (*WebhookPayloadCommentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadExternalPost) SetAccount(v WebhookPayloadCommentAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadExternalPost) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadExternalPost) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadExternalPost) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


