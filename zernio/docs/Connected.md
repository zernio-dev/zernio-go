# Connected

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**ChatId** | Pointer to **string** |  | [optional] 
**ChatTitle** | Pointer to **string** |  | [optional] 
**ChatType** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ConnectedAccount**](ConnectedAccount.md) |  | [optional] 

## Methods

### NewConnected

`func NewConnected() *Connected`

NewConnected instantiates a new Connected object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectedWithDefaults

`func NewConnectedWithDefaults() *Connected`

NewConnectedWithDefaults instantiates a new Connected object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Connected) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Connected) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Connected) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Connected) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetChatId

`func (o *Connected) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *Connected) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *Connected) SetChatId(v string)`

SetChatId sets ChatId field to given value.

### HasChatId

`func (o *Connected) HasChatId() bool`

HasChatId returns a boolean if a field has been set.

### GetChatTitle

`func (o *Connected) GetChatTitle() string`

GetChatTitle returns the ChatTitle field if non-nil, zero value otherwise.

### GetChatTitleOk

`func (o *Connected) GetChatTitleOk() (*string, bool)`

GetChatTitleOk returns a tuple with the ChatTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatTitle

`func (o *Connected) SetChatTitle(v string)`

SetChatTitle sets ChatTitle field to given value.

### HasChatTitle

`func (o *Connected) HasChatTitle() bool`

HasChatTitle returns a boolean if a field has been set.

### GetChatType

`func (o *Connected) GetChatType() string`

GetChatType returns the ChatType field if non-nil, zero value otherwise.

### GetChatTypeOk

`func (o *Connected) GetChatTypeOk() (*string, bool)`

GetChatTypeOk returns a tuple with the ChatType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatType

`func (o *Connected) SetChatType(v string)`

SetChatType sets ChatType field to given value.

### HasChatType

`func (o *Connected) HasChatType() bool`

HasChatType returns a boolean if a field has been set.

### GetAccount

`func (o *Connected) GetAccount() ConnectedAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Connected) GetAccountOk() (*ConnectedAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Connected) SetAccount(v ConnectedAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Connected) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


