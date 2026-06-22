# WebhookPayloadLeadAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Social account ID (the facebook account owning the Page) | 
**Platform** | **string** |  | 

## Methods

### NewWebhookPayloadLeadAccount

`func NewWebhookPayloadLeadAccount(id string, platform string, ) *WebhookPayloadLeadAccount`

NewWebhookPayloadLeadAccount instantiates a new WebhookPayloadLeadAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadLeadAccountWithDefaults

`func NewWebhookPayloadLeadAccountWithDefaults() *WebhookPayloadLeadAccount`

NewWebhookPayloadLeadAccountWithDefaults instantiates a new WebhookPayloadLeadAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadLeadAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadLeadAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadLeadAccount) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *WebhookPayloadLeadAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebhookPayloadLeadAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebhookPayloadLeadAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


