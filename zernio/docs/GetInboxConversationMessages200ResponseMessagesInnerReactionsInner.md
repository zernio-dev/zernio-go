# GetInboxConversationMessages200ResponseMessagesInnerReactionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emoji** | Pointer to **string** |  | [optional] 
**FromMe** | Pointer to **bool** | true if the connected account reacted, false if the contact did. | [optional] 
**ReactedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetInboxConversationMessages200ResponseMessagesInnerReactionsInner

`func NewGetInboxConversationMessages200ResponseMessagesInnerReactionsInner() *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner`

NewGetInboxConversationMessages200ResponseMessagesInnerReactionsInner instantiates a new GetInboxConversationMessages200ResponseMessagesInnerReactionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversationMessages200ResponseMessagesInnerReactionsInnerWithDefaults

`func NewGetInboxConversationMessages200ResponseMessagesInnerReactionsInnerWithDefaults() *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner`

NewGetInboxConversationMessages200ResponseMessagesInnerReactionsInnerWithDefaults instantiates a new GetInboxConversationMessages200ResponseMessagesInnerReactionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmoji

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) GetEmoji() string`

GetEmoji returns the Emoji field if non-nil, zero value otherwise.

### GetEmojiOk

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) GetEmojiOk() (*string, bool)`

GetEmojiOk returns a tuple with the Emoji field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmoji

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) SetEmoji(v string)`

SetEmoji sets Emoji field to given value.

### HasEmoji

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) HasEmoji() bool`

HasEmoji returns a boolean if a field has been set.

### GetFromMe

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) GetFromMe() bool`

GetFromMe returns the FromMe field if non-nil, zero value otherwise.

### GetFromMeOk

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) GetFromMeOk() (*bool, bool)`

GetFromMeOk returns a tuple with the FromMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMe

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) SetFromMe(v bool)`

SetFromMe sets FromMe field to given value.

### HasFromMe

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) HasFromMe() bool`

HasFromMe returns a boolean if a field has been set.

### GetReactedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) GetReactedAt() time.Time`

GetReactedAt returns the ReactedAt field if non-nil, zero value otherwise.

### GetReactedAtOk

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) GetReactedAtOk() (*time.Time, bool)`

GetReactedAtOk returns a tuple with the ReactedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReactedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) SetReactedAt(v time.Time)`

SetReactedAt sets ReactedAt field to given value.

### HasReactedAt

`func (o *GetInboxConversationMessages200ResponseMessagesInnerReactionsInner) HasReactedAt() bool`

HasReactedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


