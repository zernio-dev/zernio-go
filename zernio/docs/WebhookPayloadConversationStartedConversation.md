# WebhookPayloadConversationStartedConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Internal conversation ID | 
**Platform** | **string** |  | 
**PlatformConversationId** | **string** |  | 
**ParticipantId** | Pointer to **string** | Contact&#39;s platform identifier (IGSID, PSID, wa_id, etc.) | [optional] 
**ParticipantName** | **string** |  | 
**ParticipantUsername** | Pointer to **string** | Contact&#39;s handle when the platform exposes one | [optional] 
**ParticipantPicture** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**ContactId** | Pointer to **string** | Zernio CRM Contact ID for the participant, when one exists. Resolved by joining &#x60;participantId&#x60; to the ContactChannel collection (same join used by message.*, reaction.received, and call.* webhooks). Best-effort: omitted when no channel matches or &#x60;participantId&#x60; is absent. Lets integrators seed the CRM straight from &#x60;conversation.started&#x60; without waiting for the first &#x60;message.*&#x60; event.  | [optional] 

## Methods

### NewWebhookPayloadConversationStartedConversation

`func NewWebhookPayloadConversationStartedConversation(id string, platform string, platformConversationId string, participantName string, status string, ) *WebhookPayloadConversationStartedConversation`

NewWebhookPayloadConversationStartedConversation instantiates a new WebhookPayloadConversationStartedConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadConversationStartedConversationWithDefaults

`func NewWebhookPayloadConversationStartedConversationWithDefaults() *WebhookPayloadConversationStartedConversation`

NewWebhookPayloadConversationStartedConversationWithDefaults instantiates a new WebhookPayloadConversationStartedConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadConversationStartedConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadConversationStartedConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadConversationStartedConversation) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *WebhookPayloadConversationStartedConversation) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebhookPayloadConversationStartedConversation) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebhookPayloadConversationStartedConversation) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformConversationId

`func (o *WebhookPayloadConversationStartedConversation) GetPlatformConversationId() string`

GetPlatformConversationId returns the PlatformConversationId field if non-nil, zero value otherwise.

### GetPlatformConversationIdOk

`func (o *WebhookPayloadConversationStartedConversation) GetPlatformConversationIdOk() (*string, bool)`

GetPlatformConversationIdOk returns a tuple with the PlatformConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationId

`func (o *WebhookPayloadConversationStartedConversation) SetPlatformConversationId(v string)`

SetPlatformConversationId sets PlatformConversationId field to given value.


### GetParticipantId

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *WebhookPayloadConversationStartedConversation) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *WebhookPayloadConversationStartedConversation) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetParticipantName

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *WebhookPayloadConversationStartedConversation) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.


### GetParticipantUsername

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantUsername() string`

GetParticipantUsername returns the ParticipantUsername field if non-nil, zero value otherwise.

### GetParticipantUsernameOk

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantUsernameOk() (*string, bool)`

GetParticipantUsernameOk returns a tuple with the ParticipantUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantUsername

`func (o *WebhookPayloadConversationStartedConversation) SetParticipantUsername(v string)`

SetParticipantUsername sets ParticipantUsername field to given value.

### HasParticipantUsername

`func (o *WebhookPayloadConversationStartedConversation) HasParticipantUsername() bool`

HasParticipantUsername returns a boolean if a field has been set.

### GetParticipantPicture

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantPicture() string`

GetParticipantPicture returns the ParticipantPicture field if non-nil, zero value otherwise.

### GetParticipantPictureOk

`func (o *WebhookPayloadConversationStartedConversation) GetParticipantPictureOk() (*string, bool)`

GetParticipantPictureOk returns a tuple with the ParticipantPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantPicture

`func (o *WebhookPayloadConversationStartedConversation) SetParticipantPicture(v string)`

SetParticipantPicture sets ParticipantPicture field to given value.

### HasParticipantPicture

`func (o *WebhookPayloadConversationStartedConversation) HasParticipantPicture() bool`

HasParticipantPicture returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookPayloadConversationStartedConversation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPayloadConversationStartedConversation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPayloadConversationStartedConversation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetContactId

`func (o *WebhookPayloadConversationStartedConversation) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *WebhookPayloadConversationStartedConversation) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *WebhookPayloadConversationStartedConversation) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *WebhookPayloadConversationStartedConversation) HasContactId() bool`

HasContactId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


