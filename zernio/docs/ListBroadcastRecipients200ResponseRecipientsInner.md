# ListBroadcastRecipients200ResponseRecipientsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContactId** | Pointer to **string** |  | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**PlatformIdentifier** | Pointer to **string** |  | [optional] 
**ContactName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**MessageId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **NullableInt32** | Meta WhatsApp error code (e.g. 131049 for antispam, 131021 for invalid phone, 131026 for re-engagement required). Only populated for status&#x3D;failed. | [optional] 
**ErrorExplanation** | Pointer to **NullableString** | Plain-language translation of errorCode (e.g. for 131026, that the recipient has likely opted out of marketing messages). Null for unmapped codes; fall back to error. | [optional] 
**SentAt** | Pointer to **time.Time** |  | [optional] 
**DeliveredAt** | Pointer to **time.Time** |  | [optional] 
**ReadAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListBroadcastRecipients200ResponseRecipientsInner

`func NewListBroadcastRecipients200ResponseRecipientsInner() *ListBroadcastRecipients200ResponseRecipientsInner`

NewListBroadcastRecipients200ResponseRecipientsInner instantiates a new ListBroadcastRecipients200ResponseRecipientsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBroadcastRecipients200ResponseRecipientsInnerWithDefaults

`func NewListBroadcastRecipients200ResponseRecipientsInnerWithDefaults() *ListBroadcastRecipients200ResponseRecipientsInner`

NewListBroadcastRecipients200ResponseRecipientsInnerWithDefaults instantiates a new ListBroadcastRecipients200ResponseRecipientsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContactId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetContactId() string`

GetContactId returns the ContactId field if non-nil, zero value otherwise.

### GetContactIdOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetContactIdOk() (*string, bool)`

GetContactIdOk returns a tuple with the ContactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetContactId(v string)`

SetContactId sets ContactId field to given value.

### HasContactId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasContactId() bool`

HasContactId returns a boolean if a field has been set.

### GetChannelId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetPlatformIdentifier

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetPlatformIdentifier() string`

GetPlatformIdentifier returns the PlatformIdentifier field if non-nil, zero value otherwise.

### GetPlatformIdentifierOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetPlatformIdentifierOk() (*string, bool)`

GetPlatformIdentifierOk returns a tuple with the PlatformIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentifier

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetPlatformIdentifier(v string)`

SetPlatformIdentifier sets PlatformIdentifier field to given value.

### HasPlatformIdentifier

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasPlatformIdentifier() bool`

HasPlatformIdentifier returns a boolean if a field has been set.

### GetContactName

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetContactName() string`

GetContactName returns the ContactName field if non-nil, zero value otherwise.

### GetContactNameOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetContactNameOk() (*string, bool)`

GetContactNameOk returns a tuple with the ContactName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactName

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetContactName(v string)`

SetContactName sets ContactName field to given value.

### HasContactName

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasContactName() bool`

HasContactName returns a boolean if a field has been set.

### GetStatus

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessageId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetError

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorCode

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *ListBroadcastRecipients200ResponseRecipientsInner) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetErrorExplanation

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetErrorExplanation() string`

GetErrorExplanation returns the ErrorExplanation field if non-nil, zero value otherwise.

### GetErrorExplanationOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetErrorExplanationOk() (*string, bool)`

GetErrorExplanationOk returns a tuple with the ErrorExplanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorExplanation

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetErrorExplanation(v string)`

SetErrorExplanation sets ErrorExplanation field to given value.

### HasErrorExplanation

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasErrorExplanation() bool`

HasErrorExplanation returns a boolean if a field has been set.

### SetErrorExplanationNil

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetErrorExplanationNil(b bool)`

 SetErrorExplanationNil sets the value for ErrorExplanation to be an explicit nil

### UnsetErrorExplanation
`func (o *ListBroadcastRecipients200ResponseRecipientsInner) UnsetErrorExplanation()`

UnsetErrorExplanation ensures that no value is present for ErrorExplanation, not even an explicit nil
### GetSentAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### GetDeliveredAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### GetReadAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.

### HasReadAt

`func (o *ListBroadcastRecipients200ResponseRecipientsInner) HasReadAt() bool`

HasReadAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


