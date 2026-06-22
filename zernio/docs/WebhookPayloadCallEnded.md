# WebhookPayloadCallEnded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Call** | [**WebhookPayloadCallEndedCall**](WebhookPayloadCallEndedCall.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadCallEnded

`func NewWebhookPayloadCallEnded(id string, event string, call WebhookPayloadCallEndedCall, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadCallEnded`

NewWebhookPayloadCallEnded instantiates a new WebhookPayloadCallEnded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallEndedWithDefaults

`func NewWebhookPayloadCallEndedWithDefaults() *WebhookPayloadCallEnded`

NewWebhookPayloadCallEndedWithDefaults instantiates a new WebhookPayloadCallEnded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallEnded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallEnded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallEnded) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadCallEnded) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadCallEnded) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadCallEnded) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetCall

`func (o *WebhookPayloadCallEnded) GetCall() WebhookPayloadCallEndedCall`

GetCall returns the Call field if non-nil, zero value otherwise.

### GetCallOk

`func (o *WebhookPayloadCallEnded) GetCallOk() (*WebhookPayloadCallEndedCall, bool)`

GetCallOk returns a tuple with the Call field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCall

`func (o *WebhookPayloadCallEnded) SetCall(v WebhookPayloadCallEndedCall)`

SetCall sets Call field to given value.


### GetAccount

`func (o *WebhookPayloadCallEnded) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadCallEnded) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadCallEnded) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadCallEnded) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadCallEnded) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadCallEnded) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


