# EditInboxMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Social account ID | 
**Text** | Pointer to **string** | New message text | [optional] 
**ReplyMarkup** | Pointer to [**EditInboxMessageRequestReplyMarkup**](EditInboxMessageRequestReplyMarkup.md) |  | [optional] 

## Methods

### NewEditInboxMessageRequest

`func NewEditInboxMessageRequest(accountId string, ) *EditInboxMessageRequest`

NewEditInboxMessageRequest instantiates a new EditInboxMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditInboxMessageRequestWithDefaults

`func NewEditInboxMessageRequestWithDefaults() *EditInboxMessageRequest`

NewEditInboxMessageRequestWithDefaults instantiates a new EditInboxMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EditInboxMessageRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EditInboxMessageRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EditInboxMessageRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetText

`func (o *EditInboxMessageRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *EditInboxMessageRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *EditInboxMessageRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *EditInboxMessageRequest) HasText() bool`

HasText returns a boolean if a field has been set.

### GetReplyMarkup

`func (o *EditInboxMessageRequest) GetReplyMarkup() EditInboxMessageRequestReplyMarkup`

GetReplyMarkup returns the ReplyMarkup field if non-nil, zero value otherwise.

### GetReplyMarkupOk

`func (o *EditInboxMessageRequest) GetReplyMarkupOk() (*EditInboxMessageRequestReplyMarkup, bool)`

GetReplyMarkupOk returns a tuple with the ReplyMarkup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyMarkup

`func (o *EditInboxMessageRequest) SetReplyMarkup(v EditInboxMessageRequestReplyMarkup)`

SetReplyMarkup sets ReplyMarkup field to given value.

### HasReplyMarkup

`func (o *EditInboxMessageRequest) HasReplyMarkup() bool`

HasReplyMarkup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


