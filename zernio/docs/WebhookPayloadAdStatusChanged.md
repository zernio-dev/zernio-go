# WebhookPayloadAdStatusChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Account** | [**WebhookPayloadAccountConnectedAccount**](WebhookPayloadAccountConnectedAccount.md) |  | 
**AdObject** | [**WebhookPayloadAdStatusChangedAdObject**](WebhookPayloadAdStatusChangedAdObject.md) |  | 
**Status** | [**WebhookPayloadAdStatusChangedStatus**](WebhookPayloadAdStatusChangedStatus.md) |  | 
**Error** | Pointer to [**WebhookPayloadAdStatusChangedError**](WebhookPayloadAdStatusChangedError.md) |  | [optional] 
**Timestamp** | **time.Time** | ISO-8601 timestamp the webhook was produced. | 

## Methods

### NewWebhookPayloadAdStatusChanged

`func NewWebhookPayloadAdStatusChanged(id string, event string, account WebhookPayloadAccountConnectedAccount, adObject WebhookPayloadAdStatusChangedAdObject, status WebhookPayloadAdStatusChangedStatus, timestamp time.Time, ) *WebhookPayloadAdStatusChanged`

NewWebhookPayloadAdStatusChanged instantiates a new WebhookPayloadAdStatusChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAdStatusChangedWithDefaults

`func NewWebhookPayloadAdStatusChangedWithDefaults() *WebhookPayloadAdStatusChanged`

NewWebhookPayloadAdStatusChangedWithDefaults instantiates a new WebhookPayloadAdStatusChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadAdStatusChanged) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadAdStatusChanged) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadAdStatusChanged) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadAdStatusChanged) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadAdStatusChanged) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadAdStatusChanged) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAccount

`func (o *WebhookPayloadAdStatusChanged) GetAccount() WebhookPayloadAccountConnectedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadAdStatusChanged) GetAccountOk() (*WebhookPayloadAccountConnectedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadAdStatusChanged) SetAccount(v WebhookPayloadAccountConnectedAccount)`

SetAccount sets Account field to given value.


### GetAdObject

`func (o *WebhookPayloadAdStatusChanged) GetAdObject() WebhookPayloadAdStatusChangedAdObject`

GetAdObject returns the AdObject field if non-nil, zero value otherwise.

### GetAdObjectOk

`func (o *WebhookPayloadAdStatusChanged) GetAdObjectOk() (*WebhookPayloadAdStatusChangedAdObject, bool)`

GetAdObjectOk returns a tuple with the AdObject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdObject

`func (o *WebhookPayloadAdStatusChanged) SetAdObject(v WebhookPayloadAdStatusChangedAdObject)`

SetAdObject sets AdObject field to given value.


### GetStatus

`func (o *WebhookPayloadAdStatusChanged) GetStatus() WebhookPayloadAdStatusChangedStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPayloadAdStatusChanged) GetStatusOk() (*WebhookPayloadAdStatusChangedStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPayloadAdStatusChanged) SetStatus(v WebhookPayloadAdStatusChangedStatus)`

SetStatus sets Status field to given value.


### GetError

`func (o *WebhookPayloadAdStatusChanged) GetError() WebhookPayloadAdStatusChangedError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookPayloadAdStatusChanged) GetErrorOk() (*WebhookPayloadAdStatusChangedError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookPayloadAdStatusChanged) SetError(v WebhookPayloadAdStatusChangedError)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookPayloadAdStatusChanged) HasError() bool`

HasError returns a boolean if a field has been set.

### GetTimestamp

`func (o *WebhookPayloadAdStatusChanged) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadAdStatusChanged) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadAdStatusChanged) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


