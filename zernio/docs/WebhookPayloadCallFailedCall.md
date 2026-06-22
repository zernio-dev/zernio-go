# WebhookPayloadCallFailedCall

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
**FailedAt** | Pointer to **time.Time** |  | [optional] 
**Error** | Pointer to [**WebhookPayloadCallFailedCallError**](WebhookPayloadCallFailedCallError.md) |  | [optional] 

## Methods

### NewWebhookPayloadCallFailedCall

`func NewWebhookPayloadCallFailedCall() *WebhookPayloadCallFailedCall`

NewWebhookPayloadCallFailedCall instantiates a new WebhookPayloadCallFailedCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallFailedCallWithDefaults

`func NewWebhookPayloadCallFailedCallWithDefaults() *WebhookPayloadCallFailedCall`

NewWebhookPayloadCallFailedCallWithDefaults instantiates a new WebhookPayloadCallFailedCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallFailedCall) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallFailedCall) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallFailedCall) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPayloadCallFailedCall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMetaCallId

`func (o *WebhookPayloadCallFailedCall) GetMetaCallId() string`

GetMetaCallId returns the MetaCallId field if non-nil, zero value otherwise.

### GetMetaCallIdOk

`func (o *WebhookPayloadCallFailedCall) GetMetaCallIdOk() (*string, bool)`

GetMetaCallIdOk returns a tuple with the MetaCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCallId

`func (o *WebhookPayloadCallFailedCall) SetMetaCallId(v string)`

SetMetaCallId sets MetaCallId field to given value.

### HasMetaCallId

`func (o *WebhookPayloadCallFailedCall) HasMetaCallId() bool`

HasMetaCallId returns a boolean if a field has been set.

### SetMetaCallIdNil

`func (o *WebhookPayloadCallFailedCall) SetMetaCallIdNil(b bool)`

 SetMetaCallIdNil sets the value for MetaCallId to be an explicit nil

### UnsetMetaCallId
`func (o *WebhookPayloadCallFailedCall) UnsetMetaCallId()`

UnsetMetaCallId ensures that no value is present for MetaCallId, not even an explicit nil
### GetAccountId

`func (o *WebhookPayloadCallFailedCall) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookPayloadCallFailedCall) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookPayloadCallFailedCall) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *WebhookPayloadCallFailedCall) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPhoneNumberId

`func (o *WebhookPayloadCallFailedCall) GetPhoneNumberId() string`

GetPhoneNumberId returns the PhoneNumberId field if non-nil, zero value otherwise.

### GetPhoneNumberIdOk

`func (o *WebhookPayloadCallFailedCall) GetPhoneNumberIdOk() (*string, bool)`

GetPhoneNumberIdOk returns a tuple with the PhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberId

`func (o *WebhookPayloadCallFailedCall) SetPhoneNumberId(v string)`

SetPhoneNumberId sets PhoneNumberId field to given value.

### HasPhoneNumberId

`func (o *WebhookPayloadCallFailedCall) HasPhoneNumberId() bool`

HasPhoneNumberId returns a boolean if a field has been set.

### GetDirection

`func (o *WebhookPayloadCallFailedCall) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *WebhookPayloadCallFailedCall) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *WebhookPayloadCallFailedCall) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *WebhookPayloadCallFailedCall) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFrom

`func (o *WebhookPayloadCallFailedCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WebhookPayloadCallFailedCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WebhookPayloadCallFailedCall) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WebhookPayloadCallFailedCall) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *WebhookPayloadCallFailedCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WebhookPayloadCallFailedCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WebhookPayloadCallFailedCall) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *WebhookPayloadCallFailedCall) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFailedAt

`func (o *WebhookPayloadCallFailedCall) GetFailedAt() time.Time`

GetFailedAt returns the FailedAt field if non-nil, zero value otherwise.

### GetFailedAtOk

`func (o *WebhookPayloadCallFailedCall) GetFailedAtOk() (*time.Time, bool)`

GetFailedAtOk returns a tuple with the FailedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedAt

`func (o *WebhookPayloadCallFailedCall) SetFailedAt(v time.Time)`

SetFailedAt sets FailedAt field to given value.

### HasFailedAt

`func (o *WebhookPayloadCallFailedCall) HasFailedAt() bool`

HasFailedAt returns a boolean if a field has been set.

### GetError

`func (o *WebhookPayloadCallFailedCall) GetError() WebhookPayloadCallFailedCallError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *WebhookPayloadCallFailedCall) GetErrorOk() (*WebhookPayloadCallFailedCallError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *WebhookPayloadCallFailedCall) SetError(v WebhookPayloadCallFailedCallError)`

SetError sets Error field to given value.

### HasError

`func (o *WebhookPayloadCallFailedCall) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


