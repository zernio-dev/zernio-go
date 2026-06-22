# InboxWebhookMessageSender

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Sender&#39;s platform identifier. For WhatsApp this is the phone number (without leading &#x60;+&#x60;) when available, otherwise the &#x60;businessScopedUserId&#x60;. For other platforms, the platform&#39;s own user ID.  | 
**ContactId** | Pointer to **string** | Zernio CRM Contact id for this sender, when one exists (joined via the ContactChannel mapping). Lets integrators link a message straight to a Contact without a follow-up Contacts API call. Omitted when the sender isn&#39;t a tracked contact (e.g. outgoing messages where the sender is the business, or first-touch messages before the contact is created).  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **NullableString** | WhatsApp only. Sender&#39;s phone number in E.164 format (with leading &#x60;+&#x60;).  **Nullable during the BSUID rollout (April 2026+).** WhatsApp users who adopt a username can message businesses without exposing a phone number — this field is omitted for them. Match by &#x60;businessScopedUserId&#x60; instead. See &#x60;docs/whatsapp-bsuid-migration.md&#x60;.  | [optional] 
**BusinessScopedUserId** | Pointer to **string** | WhatsApp only. Business-scoped user ID (BSUID) — Meta&#39;s canonical identifier for a WhatsApp user within your business. Present when Meta includes it in the inbound payload (rollout in progress since early April 2026). **Recommended primary identity anchor** going forward; fall back to &#x60;phoneNumber&#x60; only when this field is absent.  | [optional] 
**ParentBusinessScopedUserId** | Pointer to **string** | WhatsApp only. Parent BSUID for businesses with linked business portfolios. Omitted for standalone portfolios.  | [optional] 
**WhatsappUsername** | Pointer to **string** | WhatsApp only. User&#39;s WhatsApp username (e.g. &#x60;@jane&#x60;). Not a stable identifier — users can change it. Useful for display, not recommended as an identity anchor.  | [optional] 
**InstagramProfile** | Pointer to [**InboxWebhookMessageSenderInstagramProfile**](InboxWebhookMessageSenderInstagramProfile.md) |  | [optional] 

## Methods

### NewInboxWebhookMessageSender

`func NewInboxWebhookMessageSender(id string, ) *InboxWebhookMessageSender`

NewInboxWebhookMessageSender instantiates a new InboxWebhookMessageSender object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxWebhookMessageSenderWithDefaults

`func NewInboxWebhookMessageSenderWithDefaults() *InboxWebhookMessageSender`

NewInboxWebhookMessageSenderWithDefaults instantiates a new InboxWebhookMessageSender object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboxWebhookMessageSender) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxWebhookMessageSender) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxWebhookMessageSender) SetId(v string)`

SetId sets Id field to given value.


### GetContactId

`func (o *InboxWebhookMessageSender) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *InboxWebhookMessageSender) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *InboxWebhookMessageSender) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *InboxWebhookMessageSender) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetName

`func (o *InboxWebhookMessageSender) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InboxWebhookMessageSender) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InboxWebhookMessageSender) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InboxWebhookMessageSender) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsername

`func (o *InboxWebhookMessageSender) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InboxWebhookMessageSender) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InboxWebhookMessageSender) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InboxWebhookMessageSender) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPicture

`func (o *InboxWebhookMessageSender) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *InboxWebhookMessageSender) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *InboxWebhookMessageSender) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *InboxWebhookMessageSender) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *InboxWebhookMessageSender) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *InboxWebhookMessageSender) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *InboxWebhookMessageSender) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *InboxWebhookMessageSender) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### SetPhoneNumberNil

`func (o *InboxWebhookMessageSender) SetPhoneNumberNil(b bool)`

 SetPhoneNumberNil sets the value for PhoneNumber to be an explicit nil

### UnsetPhoneNumber
`func (o *InboxWebhookMessageSender) UnsetPhoneNumber()`

UnsetPhoneNumber ensures that no value is present for PhoneNumber, not even an explicit nil
### GetBusinessScopedUserId

`func (o *InboxWebhookMessageSender) GetBusinessScopedUserId() string`

GetBusinessScopedUserId returns the BusinessScopedUserId field if non-nil, zero value otherwise.

### GetBusinessScopedUserIdOk

`func (o *InboxWebhookMessageSender) GetBusinessScopedUserIdOk() (*string, bool)`

GetBusinessScopedUserIdOk returns a tuple with the BusinessScopedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessScopedUserId

`func (o *InboxWebhookMessageSender) SetBusinessScopedUserId(v string)`

SetBusinessScopedUserId sets BusinessScopedUserId field to given value.

### HasBusinessScopedUserId

`func (o *InboxWebhookMessageSender) HasBusinessScopedUserId() bool`

HasBusinessScopedUserId returns a boolean if a field has been set.

### GetParentBusinessScopedUserId

`func (o *InboxWebhookMessageSender) GetParentBusinessScopedUserId() string`

GetParentBusinessScopedUserId returns the ParentBusinessScopedUserId field if non-nil, zero value otherwise.

### GetParentBusinessScopedUserIdOk

`func (o *InboxWebhookMessageSender) GetParentBusinessScopedUserIdOk() (*string, bool)`

GetParentBusinessScopedUserIdOk returns a tuple with the ParentBusinessScopedUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentBusinessScopedUserId

`func (o *InboxWebhookMessageSender) SetParentBusinessScopedUserId(v string)`

SetParentBusinessScopedUserId sets ParentBusinessScopedUserId field to given value.

### HasParentBusinessScopedUserId

`func (o *InboxWebhookMessageSender) HasParentBusinessScopedUserId() bool`

HasParentBusinessScopedUserId returns a boolean if a field has been set.

### GetWhatsappUsername

`func (o *InboxWebhookMessageSender) GetWhatsappUsername() string`

GetWhatsappUsername returns the WhatsappUsername field if non-nil, zero value otherwise.

### GetWhatsappUsernameOk

`func (o *InboxWebhookMessageSender) GetWhatsappUsernameOk() (*string, bool)`

GetWhatsappUsernameOk returns a tuple with the WhatsappUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhatsappUsername

`func (o *InboxWebhookMessageSender) SetWhatsappUsername(v string)`

SetWhatsappUsername sets WhatsappUsername field to given value.

### HasWhatsappUsername

`func (o *InboxWebhookMessageSender) HasWhatsappUsername() bool`

HasWhatsappUsername returns a boolean if a field has been set.

### GetInstagramProfile

`func (o *InboxWebhookMessageSender) GetInstagramProfile() InboxWebhookMessageSenderInstagramProfile`

GetInstagramProfile returns the InstagramProfile field if non-nil, zero value otherwise.

### GetInstagramProfileOk

`func (o *InboxWebhookMessageSender) GetInstagramProfileOk() (*InboxWebhookMessageSenderInstagramProfile, bool)`

GetInstagramProfileOk returns a tuple with the InstagramProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramProfile

`func (o *InboxWebhookMessageSender) SetInstagramProfile(v InboxWebhookMessageSenderInstagramProfile)`

SetInstagramProfile sets InstagramProfile field to given value.

### HasInstagramProfile

`func (o *InboxWebhookMessageSender) HasInstagramProfile() bool`

HasInstagramProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


