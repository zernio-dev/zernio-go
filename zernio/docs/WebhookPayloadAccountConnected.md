# WebhookPayloadAccountConnected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Account** | [**WebhookPayloadAccountConnectedAccount**](WebhookPayloadAccountConnectedAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadAccountConnected

`func NewWebhookPayloadAccountConnected(id string, event string, account WebhookPayloadAccountConnectedAccount, timestamp time.Time, ) *WebhookPayloadAccountConnected`

NewWebhookPayloadAccountConnected instantiates a new WebhookPayloadAccountConnected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountConnectedWithDefaults

`func NewWebhookPayloadAccountConnectedWithDefaults() *WebhookPayloadAccountConnected`

NewWebhookPayloadAccountConnectedWithDefaults instantiates a new WebhookPayloadAccountConnected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadAccountConnected) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadAccountConnected) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadAccountConnected) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadAccountConnected) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadAccountConnected) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadAccountConnected) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAccount

`func (o *WebhookPayloadAccountConnected) GetAccount() WebhookPayloadAccountConnectedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadAccountConnected) GetAccountOk() (*WebhookPayloadAccountConnectedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadAccountConnected) SetAccount(v WebhookPayloadAccountConnectedAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadAccountConnected) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadAccountConnected) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadAccountConnected) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


