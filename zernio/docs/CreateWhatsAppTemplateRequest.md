# CreateWhatsAppTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**Name** | **string** | Template name (lowercase, letters/numbers/underscores, must start with a letter) | 
**Category** | **string** | Template category | 
**Language** | **string** | Template language code (e.g., en_US) | 
**Components** | Pointer to [**[]WhatsAppTemplateComponent**](WhatsAppTemplateComponent.md) | Template components (header, body, footer, buttons). Required for custom templates, omit when using library_template_name. | [optional] 
**LibraryTemplateName** | Pointer to **string** | Name of a pre-built template from Meta&#39;s template library (e.g., \&quot;appointment_reminder\&quot;, \&quot;auto_pay_reminder_1\&quot;, \&quot;address_update\&quot;). When provided, the template is pre-approved by Meta with no review wait. Omit components when using this field.  | [optional] 
**LibraryTemplateBodyInputs** | Pointer to **map[string]interface{}** | Optional body customizations for library templates. Available options depend on the template (e.g., add_contact_number, add_learn_more_link, add_security_recommendation, add_track_package_link, code_expiration_minutes).  | [optional] 
**LibraryTemplateButtonInputs** | Pointer to [**[]CreateWhatsAppTemplateRequestLibraryTemplateButtonInputsInner**](CreateWhatsAppTemplateRequestLibraryTemplateButtonInputsInner.md) | Optional button customizations for library templates. Each item specifies button type and configuration (e.g., URL, phone number, quick reply).  | [optional] 

## Methods

### NewCreateWhatsAppTemplateRequest

`func NewCreateWhatsAppTemplateRequest(accountId string, name string, category string, language string, ) *CreateWhatsAppTemplateRequest`

NewCreateWhatsAppTemplateRequest instantiates a new CreateWhatsAppTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhatsAppTemplateRequestWithDefaults

`func NewCreateWhatsAppTemplateRequestWithDefaults() *CreateWhatsAppTemplateRequest`

NewCreateWhatsAppTemplateRequestWithDefaults instantiates a new CreateWhatsAppTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateWhatsAppTemplateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateWhatsAppTemplateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateWhatsAppTemplateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CreateWhatsAppTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWhatsAppTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWhatsAppTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategory

`func (o *CreateWhatsAppTemplateRequest) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateWhatsAppTemplateRequest) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateWhatsAppTemplateRequest) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetLanguage

`func (o *CreateWhatsAppTemplateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *CreateWhatsAppTemplateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *CreateWhatsAppTemplateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetComponents

`func (o *CreateWhatsAppTemplateRequest) GetComponents() []WhatsAppTemplateComponent`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *CreateWhatsAppTemplateRequest) GetComponentsOk() (*[]WhatsAppTemplateComponent, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *CreateWhatsAppTemplateRequest) SetComponents(v []WhatsAppTemplateComponent)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *CreateWhatsAppTemplateRequest) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetLibraryTemplateName

`func (o *CreateWhatsAppTemplateRequest) GetLibraryTemplateName() string`

GetLibraryTemplateName returns the LibraryTemplateName field if non-nil, zero value otherwise.

### GetLibraryTemplateNameOk

`func (o *CreateWhatsAppTemplateRequest) GetLibraryTemplateNameOk() (*string, bool)`

GetLibraryTemplateNameOk returns a tuple with the LibraryTemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryTemplateName

`func (o *CreateWhatsAppTemplateRequest) SetLibraryTemplateName(v string)`

SetLibraryTemplateName sets LibraryTemplateName field to given value.

### HasLibraryTemplateName

`func (o *CreateWhatsAppTemplateRequest) HasLibraryTemplateName() bool`

HasLibraryTemplateName returns a boolean if a field has been set.

### GetLibraryTemplateBodyInputs

`func (o *CreateWhatsAppTemplateRequest) GetLibraryTemplateBodyInputs() map[string]interface{}`

GetLibraryTemplateBodyInputs returns the LibraryTemplateBodyInputs field if non-nil, zero value otherwise.

### GetLibraryTemplateBodyInputsOk

`func (o *CreateWhatsAppTemplateRequest) GetLibraryTemplateBodyInputsOk() (*map[string]interface{}, bool)`

GetLibraryTemplateBodyInputsOk returns a tuple with the LibraryTemplateBodyInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryTemplateBodyInputs

`func (o *CreateWhatsAppTemplateRequest) SetLibraryTemplateBodyInputs(v map[string]interface{})`

SetLibraryTemplateBodyInputs sets LibraryTemplateBodyInputs field to given value.

### HasLibraryTemplateBodyInputs

`func (o *CreateWhatsAppTemplateRequest) HasLibraryTemplateBodyInputs() bool`

HasLibraryTemplateBodyInputs returns a boolean if a field has been set.

### GetLibraryTemplateButtonInputs

`func (o *CreateWhatsAppTemplateRequest) GetLibraryTemplateButtonInputs() []CreateWhatsAppTemplateRequestLibraryTemplateButtonInputsInner`

GetLibraryTemplateButtonInputs returns the LibraryTemplateButtonInputs field if non-nil, zero value otherwise.

### GetLibraryTemplateButtonInputsOk

`func (o *CreateWhatsAppTemplateRequest) GetLibraryTemplateButtonInputsOk() (*[]CreateWhatsAppTemplateRequestLibraryTemplateButtonInputsInner, bool)`

GetLibraryTemplateButtonInputsOk returns a tuple with the LibraryTemplateButtonInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryTemplateButtonInputs

`func (o *CreateWhatsAppTemplateRequest) SetLibraryTemplateButtonInputs(v []CreateWhatsAppTemplateRequestLibraryTemplateButtonInputsInner)`

SetLibraryTemplateButtonInputs sets LibraryTemplateButtonInputs field to given value.

### HasLibraryTemplateButtonInputs

`func (o *CreateWhatsAppTemplateRequest) HasLibraryTemplateButtonInputs() bool`

HasLibraryTemplateButtonInputs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


