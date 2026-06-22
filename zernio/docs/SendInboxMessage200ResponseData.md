# SendInboxMessage200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | ID of the sent message (not returned for Reddit) | [optional] 
**ConversationId** | Pointer to **NullableString** | Twitter conversation ID | [optional] 
**SentAt** | Pointer to **NullableTime** | Bluesky sent timestamp | [optional] 
**Message** | Pointer to **NullableString** | Success message (Reddit only) | [optional] 

## Methods

### NewSendInboxMessage200ResponseData

`func NewSendInboxMessage200ResponseData() *SendInboxMessage200ResponseData`

NewSendInboxMessage200ResponseData instantiates a new SendInboxMessage200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessage200ResponseDataWithDefaults

`func NewSendInboxMessage200ResponseDataWithDefaults() *SendInboxMessage200ResponseData`

NewSendInboxMessage200ResponseDataWithDefaults instantiates a new SendInboxMessage200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *SendInboxMessage200ResponseData) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SendInboxMessage200ResponseData) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SendInboxMessage200ResponseData) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SendInboxMessage200ResponseData) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetConversationId

`func (o *SendInboxMessage200ResponseData) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *SendInboxMessage200ResponseData) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *SendInboxMessage200ResponseData) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *SendInboxMessage200ResponseData) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *SendInboxMessage200ResponseData) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *SendInboxMessage200ResponseData) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetSentAt

`func (o *SendInboxMessage200ResponseData) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *SendInboxMessage200ResponseData) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *SendInboxMessage200ResponseData) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.

### HasSentAt

`func (o *SendInboxMessage200ResponseData) HasSentAt() bool`

HasSentAt returns a boolean if a field has been set.

### SetSentAtNil

`func (o *SendInboxMessage200ResponseData) SetSentAtNil(b bool)`

 SetSentAtNil sets the value for SentAt to be an explicit nil

### UnsetSentAt
`func (o *SendInboxMessage200ResponseData) UnsetSentAt()`

UnsetSentAt ensures that no value is present for SentAt, not even an explicit nil
### GetMessage

`func (o *SendInboxMessage200ResponseData) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendInboxMessage200ResponseData) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendInboxMessage200ResponseData) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendInboxMessage200ResponseData) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *SendInboxMessage200ResponseData) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *SendInboxMessage200ResponseData) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


