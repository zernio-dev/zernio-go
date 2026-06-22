# WebhookPayloadLead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Lead** | [**WebhookPayloadLeadLead**](WebhookPayloadLeadLead.md) |  | 
**Account** | [**WebhookPayloadLeadAccount**](WebhookPayloadLeadAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadLead

`func NewWebhookPayloadLead(id string, event string, lead WebhookPayloadLeadLead, account WebhookPayloadLeadAccount, timestamp time.Time, ) *WebhookPayloadLead`

NewWebhookPayloadLead instantiates a new WebhookPayloadLead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadLeadWithDefaults

`func NewWebhookPayloadLeadWithDefaults() *WebhookPayloadLead`

NewWebhookPayloadLeadWithDefaults instantiates a new WebhookPayloadLead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadLead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadLead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadLead) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadLead) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadLead) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadLead) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetLead

`func (o *WebhookPayloadLead) GetLead() WebhookPayloadLeadLead`

GetLead returns the Lead field if non-nil, zero value otherwise.

### GetLeadOk

`func (o *WebhookPayloadLead) GetLeadOk() (*WebhookPayloadLeadLead, bool)`

GetLeadOk returns a tuple with the Lead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLead

`func (o *WebhookPayloadLead) SetLead(v WebhookPayloadLeadLead)`

SetLead sets Lead field to given value.


### GetAccount

`func (o *WebhookPayloadLead) GetAccount() WebhookPayloadLeadAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadLead) GetAccountOk() (*WebhookPayloadLeadAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadLead) SetAccount(v WebhookPayloadLeadAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadLead) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadLead) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadLead) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


