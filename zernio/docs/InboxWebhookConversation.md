# InboxWebhookConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PlatformConversationId** | **string** |  | 
**ParticipantId** | Pointer to **string** |  | [optional] 
**ParticipantName** | Pointer to **string** |  | [optional] 
**ParticipantUsername** | Pointer to **string** |  | [optional] 
**ParticipantPicture** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**ContactId** | Pointer to **string** | Zernio CRM Contact ID for the participant, when one exists. Resolved by joining &#x60;participantId&#x60; to the ContactChannel collection. Best-effort: omitted when no channel matches or &#x60;participantId&#x60; is absent. Lets integrators join any inbox webhook back to the CRM Contact without needing to look at the sender — which matters for outgoing and delivery-status events whose sender is the business.  | [optional] 

## Methods

### NewInboxWebhookConversation

`func NewInboxWebhookConversation(id string, platformConversationId string, status string, ) *InboxWebhookConversation`

NewInboxWebhookConversation instantiates a new InboxWebhookConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxWebhookConversationWithDefaults

`func NewInboxWebhookConversationWithDefaults() *InboxWebhookConversation`

NewInboxWebhookConversationWithDefaults instantiates a new InboxWebhookConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboxWebhookConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxWebhookConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxWebhookConversation) SetId(v string)`

SetId sets Id field to given value.


### GetPlatformConversationId

`func (o *InboxWebhookConversation) GetPlatformConversationId() string`

GetPlatformConversationId returns the PlatformConversationId field if non-nil, zero value otherwise.

### GetPlatformConversationIdOk

`func (o *InboxWebhookConversation) GetPlatformConversationIdOk() (*string, bool)`

GetPlatformConversationIdOk returns a tuple with the PlatformConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationId

`func (o *InboxWebhookConversation) SetPlatformConversationId(v string)`

SetPlatformConversationId sets PlatformConversationId field to given value.


### GetParticipantId

`func (o *InboxWebhookConversation) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *InboxWebhookConversation) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *InboxWebhookConversation) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *InboxWebhookConversation) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetParticipantName

`func (o *InboxWebhookConversation) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *InboxWebhookConversation) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *InboxWebhookConversation) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *InboxWebhookConversation) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### GetParticipantUsername

`func (o *InboxWebhookConversation) GetParticipantUsername() string`

GetParticipantUsername returns the ParticipantUsername field if non-nil, zero value otherwise.

### GetParticipantUsernameOk

`func (o *InboxWebhookConversation) GetParticipantUsernameOk() (*string, bool)`

GetParticipantUsernameOk returns a tuple with the ParticipantUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantUsername

`func (o *InboxWebhookConversation) SetParticipantUsername(v string)`

SetParticipantUsername sets ParticipantUsername field to given value.

### HasParticipantUsername

`func (o *InboxWebhookConversation) HasParticipantUsername() bool`

HasParticipantUsername returns a boolean if a field has been set.

### GetParticipantPicture

`func (o *InboxWebhookConversation) GetParticipantPicture() string`

GetParticipantPicture returns the ParticipantPicture field if non-nil, zero value otherwise.

### GetParticipantPictureOk

`func (o *InboxWebhookConversation) GetParticipantPictureOk() (*string, bool)`

GetParticipantPictureOk returns a tuple with the ParticipantPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantPicture

`func (o *InboxWebhookConversation) SetParticipantPicture(v string)`

SetParticipantPicture sets ParticipantPicture field to given value.

### HasParticipantPicture

`func (o *InboxWebhookConversation) HasParticipantPicture() bool`

HasParticipantPicture returns a boolean if a field has been set.

### GetStatus

`func (o *InboxWebhookConversation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InboxWebhookConversation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InboxWebhookConversation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetContactId

`func (o *InboxWebhookConversation) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *InboxWebhookConversation) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *InboxWebhookConversation) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *InboxWebhookConversation) HasContactId() bool`

HasContactId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


