# UpdateWhatsAppTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**Components** | [**[]WhatsAppTemplateComponent**](WhatsAppTemplateComponent.md) | Updated template components | 

## Methods

### NewUpdateWhatsAppTemplateRequest

`func NewUpdateWhatsAppTemplateRequest(accountId string, components []WhatsAppTemplateComponent, ) *UpdateWhatsAppTemplateRequest`

NewUpdateWhatsAppTemplateRequest instantiates a new UpdateWhatsAppTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWhatsAppTemplateRequestWithDefaults

`func NewUpdateWhatsAppTemplateRequestWithDefaults() *UpdateWhatsAppTemplateRequest`

NewUpdateWhatsAppTemplateRequestWithDefaults instantiates a new UpdateWhatsAppTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateWhatsAppTemplateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWhatsAppTemplateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWhatsAppTemplateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetComponents

`func (o *UpdateWhatsAppTemplateRequest) GetComponents() []WhatsAppTemplateComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *UpdateWhatsAppTemplateRequest) GetComponentsOk() (*[]WhatsAppTemplateComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *UpdateWhatsAppTemplateRequest) SetComponents(v []WhatsAppTemplateComponent)`

SetComponents sets Components field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


