# CompleteTelegramConnect200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds until expiration | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**ChatTitle** | Pointer to **string** |  | [optional] 
**ChatType** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ConnectedAccount**](ConnectedAccount.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewCompleteTelegramConnect200Response

`func NewCompleteTelegramConnect200Response() *CompleteTelegramConnect200Response`

NewCompleteTelegramConnect200Response instantiates a new CompleteTelegramConnect200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteTelegramConnect200ResponseWithDefaults

`func NewCompleteTelegramConnect200ResponseWithDefaults() *CompleteTelegramConnect200Response`

NewCompleteTelegramConnect200ResponseWithDefaults instantiates a new CompleteTelegramConnect200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CompleteTelegramConnect200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CompleteTelegramConnect200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CompleteTelegramConnect200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CompleteTelegramConnect200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CompleteTelegramConnect200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CompleteTelegramConnect200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CompleteTelegramConnect200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CompleteTelegramConnect200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresIn

`func (o *CompleteTelegramConnect200Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CompleteTelegramConnect200Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CompleteTelegramConnect200Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CompleteTelegramConnect200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetChatId

`func (o *CompleteTelegramConnect200Response) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *CompleteTelegramConnect200Response) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *CompleteTelegramConnect200Response) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *CompleteTelegramConnect200Response) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetChatTitle

`func (o *CompleteTelegramConnect200Response) GetChatTitle() string`

GetChatTitle returns the ChatTitle field if non-nil, zero value otherwise.

### GetChatTitleOk

`func (o *CompleteTelegramConnect200Response) GetChatTitleOk() (*string, bool)`

GetChatTitleOk returns a tuple with the ChatTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatTitle

`func (o *CompleteTelegramConnect200Response) SetChatTitle(v string)`

SetChatTitle sets ChatTitle field to given value.

### HasChatTitle

`func (o *CompleteTelegramConnect200Response) HasChatTitle() bool`

HasChatTitle returns a boolean if a field has been set.

### GetChatType

`func (o *CompleteTelegramConnect200Response) GetChatType() string`

GetChatType returns the ChatType field if non-nil, zero value otherwise.

### GetChatTypeOk

`func (o *CompleteTelegramConnect200Response) GetChatTypeOk() (*string, bool)`

GetChatTypeOk returns a tuple with the ChatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatType

`func (o *CompleteTelegramConnect200Response) SetChatType(v string)`

SetChatType sets ChatType field to given value.

### HasChatType

`func (o *CompleteTelegramConnect200Response) HasChatType() bool`

HasChatType returns a boolean if a field has been set.

### GetAccount

`func (o *CompleteTelegramConnect200Response) GetAccount() ConnectedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CompleteTelegramConnect200Response) GetAccountOk() (*ConnectedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CompleteTelegramConnect200Response) SetAccount(v ConnectedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CompleteTelegramConnect200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetMessage

`func (o *CompleteTelegramConnect200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CompleteTelegramConnect200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CompleteTelegramConnect200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CompleteTelegramConnect200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


