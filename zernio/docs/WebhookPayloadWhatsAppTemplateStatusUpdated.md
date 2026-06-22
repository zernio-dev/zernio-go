# WebhookPayloadWhatsAppTemplateStatusUpdated

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Stable webhook event ID | 
**Event** | **string** |  | 
**Account** | [**WebhookPayloadWhatsAppTemplateStatusUpdatedAccount**](WebhookPayloadWhatsAppTemplateStatusUpdatedAccount.md) |  | 
**Template** | [**WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate**](WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate.md) |  | 
**Timestamp** | **time.Time** | ISO-8601 timestamp the webhook was produced. | 

## Methods

### NewWebhookPayloadWhatsAppTemplateStatusUpdated

`func NewWebhookPayloadWhatsAppTemplateStatusUpdated(id string, event string, account WebhookPayloadWhatsAppTemplateStatusUpdatedAccount, template WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate, timestamp time.Time, ) *WebhookPayloadWhatsAppTemplateStatusUpdated`

NewWebhookPayloadWhatsAppTemplateStatusUpdated instantiates a new WebhookPayloadWhatsAppTemplateStatusUpdated object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadWhatsAppTemplateStatusUpdatedWithDefaults

`func NewWebhookPayloadWhatsAppTemplateStatusUpdatedWithDefaults() *WebhookPayloadWhatsAppTemplateStatusUpdated`

NewWebhookPayloadWhatsAppTemplateStatusUpdatedWithDefaults instantiates a new WebhookPayloadWhatsAppTemplateStatusUpdated object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetAccount

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetAccount() WebhookPayloadWhatsAppTemplateStatusUpdatedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetAccountOk() (*WebhookPayloadWhatsAppTemplateStatusUpdatedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) SetAccount(v WebhookPayloadWhatsAppTemplateStatusUpdatedAccount)`

SetAccount sets Account field to given value.


### GetTemplate

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetTemplate() WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetTemplateOk() (*WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) SetTemplate(v WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate)`

SetTemplate sets Template field to given value.


### GetTimestamp

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdated) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


