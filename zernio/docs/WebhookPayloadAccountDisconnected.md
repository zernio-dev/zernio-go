# WebhookPayloadAccountDisconnected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Account** | [**WebhookPayloadAccountDisconnectedAccount**](WebhookPayloadAccountDisconnectedAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadAccountDisconnected

`func NewWebhookPayloadAccountDisconnected(id string, event string, account WebhookPayloadAccountDisconnectedAccount, timestamp time.Time, ) *WebhookPayloadAccountDisconnected`

NewWebhookPayloadAccountDisconnected instantiates a new WebhookPayloadAccountDisconnected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountDisconnectedWithDefaults

`func NewWebhookPayloadAccountDisconnectedWithDefaults() *WebhookPayloadAccountDisconnected`

NewWebhookPayloadAccountDisconnectedWithDefaults instantiates a new WebhookPayloadAccountDisconnected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadAccountDisconnected) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadAccountDisconnected) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadAccountDisconnected) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadAccountDisconnected) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadAccountDisconnected) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadAccountDisconnected) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAccount

`func (o *WebhookPayloadAccountDisconnected) GetAccount() WebhookPayloadAccountDisconnectedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadAccountDisconnected) GetAccountOk() (*WebhookPayloadAccountDisconnectedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadAccountDisconnected) SetAccount(v WebhookPayloadAccountDisconnectedAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadAccountDisconnected) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadAccountDisconnected) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadAccountDisconnected) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


