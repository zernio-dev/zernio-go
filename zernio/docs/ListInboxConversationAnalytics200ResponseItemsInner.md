# ListInboxConversationAnalytics200ResponseItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConversationId** | Pointer to **string** | The platformConversationId (the same identity used by metadata.conversationId) | [optional] 
**MongoId** | Pointer to **NullableString** | The Conversation document _id, when a matching doc exists | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**ParticipantName** | Pointer to **NullableString** |  | [optional] 
**ParticipantUsername** | Pointer to **NullableString** |  | [optional] 
**ParticipantPicture** | Pointer to **NullableString** |  | [optional] 
**LastMessage** | Pointer to **NullableString** | Cached preview from the Conversation doc | [optional] 
**TotalMessages** | Pointer to **int32** |  | [optional] 
**Received** | Pointer to **int32** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Read** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**FirstMessageAt** | Pointer to **time.Time** |  | [optional] 
**LastMessageAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListInboxConversationAnalytics200ResponseItemsInner

`func NewListInboxConversationAnalytics200ResponseItemsInner() *ListInboxConversationAnalytics200ResponseItemsInner`

NewListInboxConversationAnalytics200ResponseItemsInner instantiates a new ListInboxConversationAnalytics200ResponseItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversationAnalytics200ResponseItemsInnerWithDefaults

`func NewListInboxConversationAnalytics200ResponseItemsInnerWithDefaults() *ListInboxConversationAnalytics200ResponseItemsInner`

NewListInboxConversationAnalytics200ResponseItemsInnerWithDefaults instantiates a new ListInboxConversationAnalytics200ResponseItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConversationId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetMongoId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetMongoId() string`

GetMongoId returns the MongoId field if non-nil, zero value otherwise.

### GetMongoIdOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetMongoIdOk() (*string, bool)`

GetMongoIdOk returns a tuple with the MongoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongoId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetMongoId(v string)`

SetMongoId sets MongoId field to given value.

### HasMongoId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasMongoId() bool`

HasMongoId returns a boolean if a field has been set.

### SetMongoIdNil

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetMongoIdNil(b bool)`

 SetMongoIdNil sets the value for MongoId to be an explicit nil

### UnsetMongoId
`func (o *ListInboxConversationAnalytics200ResponseItemsInner) UnsetMongoId()`

UnsetMongoId ensures that no value is present for MongoId, not even an explicit nil
### GetAccountId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetParticipantName

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### SetParticipantNameNil

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetParticipantNameNil(b bool)`

 SetParticipantNameNil sets the value for ParticipantName to be an explicit nil

### UnsetParticipantName
`func (o *ListInboxConversationAnalytics200ResponseItemsInner) UnsetParticipantName()`

UnsetParticipantName ensures that no value is present for ParticipantName, not even an explicit nil
### GetParticipantUsername

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetParticipantUsername() string`

GetParticipantUsername returns the ParticipantUsername field if non-nil, zero value otherwise.

### GetParticipantUsernameOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetParticipantUsernameOk() (*string, bool)`

GetParticipantUsernameOk returns a tuple with the ParticipantUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantUsername

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetParticipantUsername(v string)`

SetParticipantUsername sets ParticipantUsername field to given value.

### HasParticipantUsername

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasParticipantUsername() bool`

HasParticipantUsername returns a boolean if a field has been set.

### SetParticipantUsernameNil

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetParticipantUsernameNil(b bool)`

 SetParticipantUsernameNil sets the value for ParticipantUsername to be an explicit nil

### UnsetParticipantUsername
`func (o *ListInboxConversationAnalytics200ResponseItemsInner) UnsetParticipantUsername()`

UnsetParticipantUsername ensures that no value is present for ParticipantUsername, not even an explicit nil
### GetParticipantPicture

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetParticipantPicture() string`

GetParticipantPicture returns the ParticipantPicture field if non-nil, zero value otherwise.

### GetParticipantPictureOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetParticipantPictureOk() (*string, bool)`

GetParticipantPictureOk returns a tuple with the ParticipantPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantPicture

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetParticipantPicture(v string)`

SetParticipantPicture sets ParticipantPicture field to given value.

### HasParticipantPicture

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasParticipantPicture() bool`

HasParticipantPicture returns a boolean if a field has been set.

### SetParticipantPictureNil

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetParticipantPictureNil(b bool)`

 SetParticipantPictureNil sets the value for ParticipantPicture to be an explicit nil

### UnsetParticipantPicture
`func (o *ListInboxConversationAnalytics200ResponseItemsInner) UnsetParticipantPicture()`

UnsetParticipantPicture ensures that no value is present for ParticipantPicture, not even an explicit nil
### GetLastMessage

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### SetLastMessageNil

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetLastMessageNil(b bool)`

 SetLastMessageNil sets the value for LastMessage to be an explicit nil

### UnsetLastMessage
`func (o *ListInboxConversationAnalytics200ResponseItemsInner) UnsetLastMessage()`

UnsetLastMessage ensures that no value is present for LastMessage, not even an explicit nil
### GetTotalMessages

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetTotalMessages() int32`

GetTotalMessages returns the TotalMessages field if non-nil, zero value otherwise.

### GetTotalMessagesOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetTotalMessagesOk() (*int32, bool)`

GetTotalMessagesOk returns a tuple with the TotalMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessages

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetTotalMessages(v int32)`

SetTotalMessages sets TotalMessages field to given value.

### HasTotalMessages

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasTotalMessages() bool`

HasTotalMessages returns a boolean if a field has been set.

### GetReceived

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSent

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRead

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetRead() int32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetReadOk() (*int32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetRead(v int32)`

SetRead sets Read field to given value.

### HasRead

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetFailed

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetFirstMessageAt

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetFirstMessageAt() time.Time`

GetFirstMessageAt returns the FirstMessageAt field if non-nil, zero value otherwise.

### GetFirstMessageAtOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetFirstMessageAtOk() (*time.Time, bool)`

GetFirstMessageAtOk returns a tuple with the FirstMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageAt

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetFirstMessageAt(v time.Time)`

SetFirstMessageAt sets FirstMessageAt field to given value.

### HasFirstMessageAt

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasFirstMessageAt() bool`

HasFirstMessageAt returns a boolean if a field has been set.

### GetLastMessageAt

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *ListInboxConversationAnalytics200ResponseItemsInner) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


