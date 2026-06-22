# WebhookPayloadReactionReactionSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ContactId** | Pointer to **string** | Zernio CRM Contact id for this sender, when one exists. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** | WhatsApp only. Sender&#39;s phone number in E.164 format (with leading &#x60;+&#x60;), when available. | [optional] 

## Methods

### NewWebhookPayloadReactionReactionSender

`func NewWebhookPayloadReactionReactionSender(id string, ) *WebhookPayloadReactionReactionSender`

NewWebhookPayloadReactionReactionSender instantiates a new WebhookPayloadReactionReactionSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadReactionReactionSenderWithDefaults

`func NewWebhookPayloadReactionReactionSenderWithDefaults() *WebhookPayloadReactionReactionSender`

NewWebhookPayloadReactionReactionSenderWithDefaults instantiates a new WebhookPayloadReactionReactionSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadReactionReactionSender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadReactionReactionSender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadReactionReactionSender) SetId(v string)`

SetId sets Id field to given value.


### GetContactId

`func (o *WebhookPayloadReactionReactionSender) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *WebhookPayloadReactionReactionSender) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *WebhookPayloadReactionReactionSender) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *WebhookPayloadReactionReactionSender) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetName

`func (o *WebhookPayloadReactionReactionSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WebhookPayloadReactionReactionSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WebhookPayloadReactionReactionSender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WebhookPayloadReactionReactionSender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *WebhookPayloadReactionReactionSender) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookPayloadReactionReactionSender) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookPayloadReactionReactionSender) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *WebhookPayloadReactionReactionSender) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPicture

`func (o *WebhookPayloadReactionReactionSender) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *WebhookPayloadReactionReactionSender) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *WebhookPayloadReactionReactionSender) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *WebhookPayloadReactionReactionSender) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *WebhookPayloadReactionReactionSender) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *WebhookPayloadReactionReactionSender) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *WebhookPayloadReactionReactionSender) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *WebhookPayloadReactionReactionSender) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *WebhookPayloadReactionReactionSender) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *WebhookPayloadReactionReactionSender) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


