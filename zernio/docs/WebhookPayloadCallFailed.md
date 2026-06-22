# WebhookPayloadCallFailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Call** | [**WebhookPayloadCallFailedCall**](WebhookPayloadCallFailedCall.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadCallFailed

`func NewWebhookPayloadCallFailed(id string, event string, call WebhookPayloadCallFailedCall, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadCallFailed`

NewWebhookPayloadCallFailed instantiates a new WebhookPayloadCallFailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallFailedWithDefaults

`func NewWebhookPayloadCallFailedWithDefaults() *WebhookPayloadCallFailed`

NewWebhookPayloadCallFailedWithDefaults instantiates a new WebhookPayloadCallFailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallFailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallFailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallFailed) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadCallFailed) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadCallFailed) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadCallFailed) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCall

`func (o *WebhookPayloadCallFailed) GetCall() WebhookPayloadCallFailedCall`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *WebhookPayloadCallFailed) GetCallOk() (*WebhookPayloadCallFailedCall, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *WebhookPayloadCallFailed) SetCall(v WebhookPayloadCallFailedCall)`

SetCall sets Call field to given value.


### GetAccount

`func (o *WebhookPayloadCallFailed) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadCallFailed) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadCallFailed) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadCallFailed) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadCallFailed) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadCallFailed) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


