# CreateInboxConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The social account ID to send from | 
**ParticipantId** | Pointer to **string** | Recipient identifier. For X this is the numeric user ID; for WhatsApp, the recipient phone number in international format (digits, country code included). Provide either this or participantUsername. | [optional] 
**ParticipantUsername** | Pointer to **string** | Recipient handle/username — an X or Bluesky handle (with or without @) or a Reddit username (with or without u/). Resolved via lookup. Provide either this or participantId. | [optional] 
**Message** | Pointer to **string** | Text content of the message. At least one of message, attachment, or (for WhatsApp) templateName is required. | [optional] 
**SkipDmCheck** | Pointer to **bool** | X/Twitter only. Skip the receives_your_dm eligibility check before sending. Use if you have already verified the recipient accepts DMs. | [optional] [default to false]
**TemplateName** | Pointer to **string** | WhatsApp only. Name of the approved template to start the conversation with (required for WhatsApp). | [optional] 
**TemplateLanguage** | Pointer to **string** | WhatsApp only. Template language code (e.g. en_US). | [optional] 
**TemplateParams** | Pointer to **[]string** | WhatsApp only. Body variable values, in order. Works with positional placeholders ({{1}}, {{2}}, ...) and with named placeholders ({{name}}, {{company}} - how Meta Business Manager creates templates), where values fill the named slots in order of appearance. | [optional] 

## Methods

### NewCreateInboxConversationRequest

`func NewCreateInboxConversationRequest(accountId string, ) *CreateInboxConversationRequest`

NewCreateInboxConversationRequest instantiates a new CreateInboxConversationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInboxConversationRequestWithDefaults

`func NewCreateInboxConversationRequestWithDefaults() *CreateInboxConversationRequest`

NewCreateInboxConversationRequestWithDefaults instantiates a new CreateInboxConversationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateInboxConversationRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateInboxConversationRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateInboxConversationRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetParticipantId

`func (o *CreateInboxConversationRequest) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *CreateInboxConversationRequest) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *CreateInboxConversationRequest) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *CreateInboxConversationRequest) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetParticipantUsername

`func (o *CreateInboxConversationRequest) GetParticipantUsername() string`

GetParticipantUsername returns the ParticipantUsername field if non-nil, zero value otherwise.

### GetParticipantUsernameOk

`func (o *CreateInboxConversationRequest) GetParticipantUsernameOk() (*string, bool)`

GetParticipantUsernameOk returns a tuple with the ParticipantUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantUsername

`func (o *CreateInboxConversationRequest) SetParticipantUsername(v string)`

SetParticipantUsername sets ParticipantUsername field to given value.

### HasParticipantUsername

`func (o *CreateInboxConversationRequest) HasParticipantUsername() bool`

HasParticipantUsername returns a boolean if a field has been set.

### GetMessage

`func (o *CreateInboxConversationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateInboxConversationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateInboxConversationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateInboxConversationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetSkipDmCheck

`func (o *CreateInboxConversationRequest) GetSkipDmCheck() bool`

GetSkipDmCheck returns the SkipDmCheck field if non-nil, zero value otherwise.

### GetSkipDmCheckOk

`func (o *CreateInboxConversationRequest) GetSkipDmCheckOk() (*bool, bool)`

GetSkipDmCheckOk returns a tuple with the SkipDmCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipDmCheck

`func (o *CreateInboxConversationRequest) SetSkipDmCheck(v bool)`

SetSkipDmCheck sets SkipDmCheck field to given value.

### HasSkipDmCheck

`func (o *CreateInboxConversationRequest) HasSkipDmCheck() bool`

HasSkipDmCheck returns a boolean if a field has been set.

### GetTemplateName

`func (o *CreateInboxConversationRequest) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *CreateInboxConversationRequest) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *CreateInboxConversationRequest) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *CreateInboxConversationRequest) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateLanguage

`func (o *CreateInboxConversationRequest) GetTemplateLanguage() string`

GetTemplateLanguage returns the TemplateLanguage field if non-nil, zero value otherwise.

### GetTemplateLanguageOk

`func (o *CreateInboxConversationRequest) GetTemplateLanguageOk() (*string, bool)`

GetTemplateLanguageOk returns a tuple with the TemplateLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateLanguage

`func (o *CreateInboxConversationRequest) SetTemplateLanguage(v string)`

SetTemplateLanguage sets TemplateLanguage field to given value.

### HasTemplateLanguage

`func (o *CreateInboxConversationRequest) HasTemplateLanguage() bool`

HasTemplateLanguage returns a boolean if a field has been set.

### GetTemplateParams

`func (o *CreateInboxConversationRequest) GetTemplateParams() []string`

GetTemplateParams returns the TemplateParams field if non-nil, zero value otherwise.

### GetTemplateParamsOk

`func (o *CreateInboxConversationRequest) GetTemplateParamsOk() (*[]string, bool)`

GetTemplateParamsOk returns a tuple with the TemplateParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParams

`func (o *CreateInboxConversationRequest) SetTemplateParams(v []string)`

SetTemplateParams sets TemplateParams field to given value.

### HasTemplateParams

`func (o *CreateInboxConversationRequest) HasTemplateParams() bool`

HasTemplateParams returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


