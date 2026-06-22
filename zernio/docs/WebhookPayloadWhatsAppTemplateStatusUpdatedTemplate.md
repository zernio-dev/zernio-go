# WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateId** | **string** | Meta&#39;s &#x60;message_template_id&#x60;, returned as a string. | 
**Name** | **string** | Meta&#39;s &#x60;message_template_name&#x60;. | 
**Language** | **string** | Meta&#39;s &#x60;message_template_language&#x60; (e.g. &#x60;en_US&#x60;). | 
**Status** | **string** | New status. Forwarded verbatim from Meta&#39;s &#x60;event&#x60; field. &#x60;PENDING_DELETION&#x60; is the 24h-grace state after a delete request before the template is actually removed.  | 
**Reason** | **string** | Meta&#39;s free-form reason for the transition. &#x60;\&quot;NONE\&quot;&#x60; on approval; an explanation string on rejection.  | 

## Methods

### NewWebhookPayloadWhatsAppTemplateStatusUpdatedTemplate

`func NewWebhookPayloadWhatsAppTemplateStatusUpdatedTemplate(templateId string, name string, language string, status string, reason string, ) *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate`

NewWebhookPayloadWhatsAppTemplateStatusUpdatedTemplate instantiates a new WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadWhatsAppTemplateStatusUpdatedTemplateWithDefaults

`func NewWebhookPayloadWhatsAppTemplateStatusUpdatedTemplateWithDefaults() *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate`

NewWebhookPayloadWhatsAppTemplateStatusUpdatedTemplateWithDefaults instantiates a new WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateId

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.


### GetName

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetLanguage

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetStatus

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WebhookPayloadWhatsAppTemplateStatusUpdatedTemplate) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


