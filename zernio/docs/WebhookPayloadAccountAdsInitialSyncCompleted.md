# WebhookPayloadAccountAdsInitialSyncCompleted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Account** | [**WebhookPayloadAccountAdsInitialSyncCompletedAccount**](WebhookPayloadAccountAdsInitialSyncCompletedAccount.md) |  | 
**Sync** | [**WebhookPayloadAccountAdsInitialSyncCompletedSync**](WebhookPayloadAccountAdsInitialSyncCompletedSync.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadAccountAdsInitialSyncCompleted

`func NewWebhookPayloadAccountAdsInitialSyncCompleted(id string, event string, account WebhookPayloadAccountAdsInitialSyncCompletedAccount, sync WebhookPayloadAccountAdsInitialSyncCompletedSync, timestamp time.Time, ) *WebhookPayloadAccountAdsInitialSyncCompleted`

NewWebhookPayloadAccountAdsInitialSyncCompleted instantiates a new WebhookPayloadAccountAdsInitialSyncCompleted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountAdsInitialSyncCompletedWithDefaults

`func NewWebhookPayloadAccountAdsInitialSyncCompletedWithDefaults() *WebhookPayloadAccountAdsInitialSyncCompleted`

NewWebhookPayloadAccountAdsInitialSyncCompletedWithDefaults instantiates a new WebhookPayloadAccountAdsInitialSyncCompleted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAccount

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetAccount() WebhookPayloadAccountAdsInitialSyncCompletedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetAccountOk() (*WebhookPayloadAccountAdsInitialSyncCompletedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) SetAccount(v WebhookPayloadAccountAdsInitialSyncCompletedAccount)`

SetAccount sets Account field to given value.


### GetSync

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetSync() WebhookPayloadAccountAdsInitialSyncCompletedSync`

GetSync returns the Sync field if non-nil, zero value otherwise.

### GetSyncOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetSyncOk() (*WebhookPayloadAccountAdsInitialSyncCompletedSync, bool)`

GetSyncOk returns a tuple with the Sync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSync

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) SetSync(v WebhookPayloadAccountAdsInitialSyncCompletedSync)`

SetSync sets Sync field to given value.


### GetTimestamp

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadAccountAdsInitialSyncCompleted) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


