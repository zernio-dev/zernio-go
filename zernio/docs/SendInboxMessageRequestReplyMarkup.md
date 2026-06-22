# SendInboxMessageRequestReplyMarkup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Keyboard type | [optional] 
**Keyboard** | Pointer to [**[][]SendInboxMessageRequestReplyMarkupKeyboardInnerInner**]([]SendInboxMessageRequestReplyMarkupKeyboardInnerInner.md) | Array of rows, each row is an array of buttons | [optional] 
**OneTime** | Pointer to **bool** | Hide keyboard after use (reply_keyboard only) | [optional] [default to true]

## Methods

### NewSendInboxMessageRequestReplyMarkup

`func NewSendInboxMessageRequestReplyMarkup() *SendInboxMessageRequestReplyMarkup`

NewSendInboxMessageRequestReplyMarkup instantiates a new SendInboxMessageRequestReplyMarkup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestReplyMarkupWithDefaults

`func NewSendInboxMessageRequestReplyMarkupWithDefaults() *SendInboxMessageRequestReplyMarkup`

NewSendInboxMessageRequestReplyMarkupWithDefaults instantiates a new SendInboxMessageRequestReplyMarkup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendInboxMessageRequestReplyMarkup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendInboxMessageRequestReplyMarkup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendInboxMessageRequestReplyMarkup) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SendInboxMessageRequestReplyMarkup) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKeyboard

`func (o *SendInboxMessageRequestReplyMarkup) GetKeyboard() [][]SendInboxMessageRequestReplyMarkupKeyboardInnerInner`

GetKeyboard returns the Keyboard field if non-nil, zero value otherwise.

### GetKeyboardOk

`func (o *SendInboxMessageRequestReplyMarkup) GetKeyboardOk() (*[][]SendInboxMessageRequestReplyMarkupKeyboardInnerInner, bool)`

GetKeyboardOk returns a tuple with the Keyboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboard

`func (o *SendInboxMessageRequestReplyMarkup) SetKeyboard(v [][]SendInboxMessageRequestReplyMarkupKeyboardInnerInner)`

SetKeyboard sets Keyboard field to given value.

### HasKeyboard

`func (o *SendInboxMessageRequestReplyMarkup) HasKeyboard() bool`

HasKeyboard returns a boolean if a field has been set.

### GetOneTime

`func (o *SendInboxMessageRequestReplyMarkup) GetOneTime() bool`

GetOneTime returns the OneTime field if non-nil, zero value otherwise.

### GetOneTimeOk

`func (o *SendInboxMessageRequestReplyMarkup) GetOneTimeOk() (*bool, bool)`

GetOneTimeOk returns a tuple with the OneTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOneTime

`func (o *SendInboxMessageRequestReplyMarkup) SetOneTime(v bool)`

SetOneTime sets OneTime field to given value.

### HasOneTime

`func (o *SendInboxMessageRequestReplyMarkup) HasOneTime() bool`

HasOneTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


