# WebhookPayloadCallReceived

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Call** | [**WebhookPayloadCallReceivedCall**](WebhookPayloadCallReceivedCall.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadCallReceived

`func NewWebhookPayloadCallReceived(id string, event string, call WebhookPayloadCallReceivedCall, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadCallReceived`

NewWebhookPayloadCallReceived instantiates a new WebhookPayloadCallReceived object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallReceivedWithDefaults

`func NewWebhookPayloadCallReceivedWithDefaults() *WebhookPayloadCallReceived`

NewWebhookPayloadCallReceivedWithDefaults instantiates a new WebhookPayloadCallReceived object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallReceived) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallReceived) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallReceived) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadCallReceived) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadCallReceived) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadCallReceived) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCall

`func (o *WebhookPayloadCallReceived) GetCall() WebhookPayloadCallReceivedCall`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *WebhookPayloadCallReceived) GetCallOk() (*WebhookPayloadCallReceivedCall, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *WebhookPayloadCallReceived) SetCall(v WebhookPayloadCallReceivedCall)`

SetCall sets Call field to given value.


### GetAccount

`func (o *WebhookPayloadCallReceived) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadCallReceived) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadCallReceived) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadCallReceived) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadCallReceived) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadCallReceived) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


