# WebhookPayloadCallReceivedCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Internal Zernio Call doc id | [optional] 
**MetaCallId** | Pointer to **NullableString** | Meta wacid.* call id when known | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**PhoneNumberId** | Pointer to **string** | Meta phone_number_id | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** | Consumer wa_id / E.164 | [optional] 
**To** | Pointer to **string** | Business number (E.164) | [optional] 
**ForwardTo** | Pointer to **string** | Destination snapshot at routing time | [optional] 
**ContactId** | Pointer to **string** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewWebhookPayloadCallReceivedCall

`func NewWebhookPayloadCallReceivedCall() *WebhookPayloadCallReceivedCall`

NewWebhookPayloadCallReceivedCall instantiates a new WebhookPayloadCallReceivedCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallReceivedCallWithDefaults

`func NewWebhookPayloadCallReceivedCallWithDefaults() *WebhookPayloadCallReceivedCall`

NewWebhookPayloadCallReceivedCallWithDefaults instantiates a new WebhookPayloadCallReceivedCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallReceivedCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallReceivedCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallReceivedCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPayloadCallReceivedCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetaCallId

`func (o *WebhookPayloadCallReceivedCall) GetMetaCallId() string`

GetMetaCallId returns the MetaCallId field if non-nil, zero value otherwise.

### GetMetaCallIdOk

`func (o *WebhookPayloadCallReceivedCall) GetMetaCallIdOk() (*string, bool)`

GetMetaCallIdOk returns a tuple with the MetaCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCallId

`func (o *WebhookPayloadCallReceivedCall) SetMetaCallId(v string)`

SetMetaCallId sets MetaCallId field to given value.

### HasMetaCallId

`func (o *WebhookPayloadCallReceivedCall) HasMetaCallId() bool`

HasMetaCallId returns a boolean if a field has been set.

### SetMetaCallIdNil

`func (o *WebhookPayloadCallReceivedCall) SetMetaCallIdNil(b bool)`

 SetMetaCallIdNil sets the value for MetaCallId to be an explicit nil

### UnsetMetaCallId
`func (o *WebhookPayloadCallReceivedCall) UnsetMetaCallId()`

UnsetMetaCallId ensures that no value is present for MetaCallId, not even an explicit nil
### GetAccountId

`func (o *WebhookPayloadCallReceivedCall) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookPayloadCallReceivedCall) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookPayloadCallReceivedCall) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *WebhookPayloadCallReceivedCall) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPhoneNumberId

`func (o *WebhookPayloadCallReceivedCall) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *WebhookPayloadCallReceivedCall) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *WebhookPayloadCallReceivedCall) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.

### HasPhoneNumberId

`func (o *WebhookPayloadCallReceivedCall) HasPhoneNumberId() bool`

HasPhoneNumberId returns a boolean if a field has been set.

### GetDirection

`func (o *WebhookPayloadCallReceivedCall) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *WebhookPayloadCallReceivedCall) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *WebhookPayloadCallReceivedCall) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *WebhookPayloadCallReceivedCall) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFrom

`func (o *WebhookPayloadCallReceivedCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WebhookPayloadCallReceivedCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WebhookPayloadCallReceivedCall) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WebhookPayloadCallReceivedCall) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *WebhookPayloadCallReceivedCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WebhookPayloadCallReceivedCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WebhookPayloadCallReceivedCall) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WebhookPayloadCallReceivedCall) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetForwardTo

`func (o *WebhookPayloadCallReceivedCall) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *WebhookPayloadCallReceivedCall) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *WebhookPayloadCallReceivedCall) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *WebhookPayloadCallReceivedCall) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetContactId

`func (o *WebhookPayloadCallReceivedCall) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *WebhookPayloadCallReceivedCall) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *WebhookPayloadCallReceivedCall) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *WebhookPayloadCallReceivedCall) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetConversationId

`func (o *WebhookPayloadCallReceivedCall) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *WebhookPayloadCallReceivedCall) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *WebhookPayloadCallReceivedCall) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *WebhookPayloadCallReceivedCall) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetStartedAt

`func (o *WebhookPayloadCallReceivedCall) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WebhookPayloadCallReceivedCall) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WebhookPayloadCallReceivedCall) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WebhookPayloadCallReceivedCall) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


