# WebhookPayloadCallEndedCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MetaCallId** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**PhoneNumberId** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**DurationSeconds** | Pointer to **int32** |  | [optional] 
**EndReason** | Pointer to **string** |  | [optional] 
**RecordingUrl** | Pointer to **string** |  | [optional] 
**RecordingExpiresAt** | Pointer to **time.Time** |  | [optional] 
**Billing** | Pointer to [**WebhookPayloadCallEndedCallBilling**](WebhookPayloadCallEndedCallBilling.md) |  | [optional] 

## Methods

### NewWebhookPayloadCallEndedCall

`func NewWebhookPayloadCallEndedCall() *WebhookPayloadCallEndedCall`

NewWebhookPayloadCallEndedCall instantiates a new WebhookPayloadCallEndedCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallEndedCallWithDefaults

`func NewWebhookPayloadCallEndedCallWithDefaults() *WebhookPayloadCallEndedCall`

NewWebhookPayloadCallEndedCallWithDefaults instantiates a new WebhookPayloadCallEndedCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallEndedCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallEndedCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallEndedCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPayloadCallEndedCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetaCallId

`func (o *WebhookPayloadCallEndedCall) GetMetaCallId() string`

GetMetaCallId returns the MetaCallId field if non-nil, zero value otherwise.

### GetMetaCallIdOk

`func (o *WebhookPayloadCallEndedCall) GetMetaCallIdOk() (*string, bool)`

GetMetaCallIdOk returns a tuple with the MetaCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCallId

`func (o *WebhookPayloadCallEndedCall) SetMetaCallId(v string)`

SetMetaCallId sets MetaCallId field to given value.

### HasMetaCallId

`func (o *WebhookPayloadCallEndedCall) HasMetaCallId() bool`

HasMetaCallId returns a boolean if a field has been set.

### SetMetaCallIdNil

`func (o *WebhookPayloadCallEndedCall) SetMetaCallIdNil(b bool)`

 SetMetaCallIdNil sets the value for MetaCallId to be an explicit nil

### UnsetMetaCallId
`func (o *WebhookPayloadCallEndedCall) UnsetMetaCallId()`

UnsetMetaCallId ensures that no value is present for MetaCallId, not even an explicit nil
### GetAccountId

`func (o *WebhookPayloadCallEndedCall) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookPayloadCallEndedCall) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookPayloadCallEndedCall) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *WebhookPayloadCallEndedCall) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPhoneNumberId

`func (o *WebhookPayloadCallEndedCall) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *WebhookPayloadCallEndedCall) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *WebhookPayloadCallEndedCall) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.

### HasPhoneNumberId

`func (o *WebhookPayloadCallEndedCall) HasPhoneNumberId() bool`

HasPhoneNumberId returns a boolean if a field has been set.

### GetDirection

`func (o *WebhookPayloadCallEndedCall) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *WebhookPayloadCallEndedCall) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *WebhookPayloadCallEndedCall) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *WebhookPayloadCallEndedCall) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFrom

`func (o *WebhookPayloadCallEndedCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WebhookPayloadCallEndedCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WebhookPayloadCallEndedCall) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WebhookPayloadCallEndedCall) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *WebhookPayloadCallEndedCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WebhookPayloadCallEndedCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WebhookPayloadCallEndedCall) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WebhookPayloadCallEndedCall) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetStartedAt

`func (o *WebhookPayloadCallEndedCall) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *WebhookPayloadCallEndedCall) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *WebhookPayloadCallEndedCall) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *WebhookPayloadCallEndedCall) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *WebhookPayloadCallEndedCall) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *WebhookPayloadCallEndedCall) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *WebhookPayloadCallEndedCall) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *WebhookPayloadCallEndedCall) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *WebhookPayloadCallEndedCall) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *WebhookPayloadCallEndedCall) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *WebhookPayloadCallEndedCall) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *WebhookPayloadCallEndedCall) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetEndReason

`func (o *WebhookPayloadCallEndedCall) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *WebhookPayloadCallEndedCall) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *WebhookPayloadCallEndedCall) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *WebhookPayloadCallEndedCall) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### GetRecordingUrl

`func (o *WebhookPayloadCallEndedCall) GetRecordingUrl() string`

GetRecordingUrl returns the RecordingUrl field if non-nil, zero value otherwise.

### GetRecordingUrlOk

`func (o *WebhookPayloadCallEndedCall) GetRecordingUrlOk() (*string, bool)`

GetRecordingUrlOk returns a tuple with the RecordingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingUrl

`func (o *WebhookPayloadCallEndedCall) SetRecordingUrl(v string)`

SetRecordingUrl sets RecordingUrl field to given value.

### HasRecordingUrl

`func (o *WebhookPayloadCallEndedCall) HasRecordingUrl() bool`

HasRecordingUrl returns a boolean if a field has been set.

### GetRecordingExpiresAt

`func (o *WebhookPayloadCallEndedCall) GetRecordingExpiresAt() time.Time`

GetRecordingExpiresAt returns the RecordingExpiresAt field if non-nil, zero value otherwise.

### GetRecordingExpiresAtOk

`func (o *WebhookPayloadCallEndedCall) GetRecordingExpiresAtOk() (*time.Time, bool)`

GetRecordingExpiresAtOk returns a tuple with the RecordingExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingExpiresAt

`func (o *WebhookPayloadCallEndedCall) SetRecordingExpiresAt(v time.Time)`

SetRecordingExpiresAt sets RecordingExpiresAt field to given value.

### HasRecordingExpiresAt

`func (o *WebhookPayloadCallEndedCall) HasRecordingExpiresAt() bool`

HasRecordingExpiresAt returns a boolean if a field has been set.

### GetBilling

`func (o *WebhookPayloadCallEndedCall) GetBilling() WebhookPayloadCallEndedCallBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *WebhookPayloadCallEndedCall) GetBillingOk() (*WebhookPayloadCallEndedCallBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *WebhookPayloadCallEndedCall) SetBilling(v WebhookPayloadCallEndedCallBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *WebhookPayloadCallEndedCall) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


