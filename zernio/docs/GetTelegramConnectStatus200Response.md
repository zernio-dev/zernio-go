# GetTelegramConnectStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The access code to send to the Telegram bot | [optional] 
**ExpiresAt** | Pointer to **time.Time** | When the code expires | [optional] 
**ExpiresIn** | Pointer to **int32** | Seconds until expiration | [optional] 
**BotUsername** | Pointer to **string** | The Telegram bot username to message | [optional] 
**Instructions** | Pointer to **[]string** | Step-by-step connection instructions | [optional] 

## Methods

### NewGetTelegramConnectStatus200Response

`func NewGetTelegramConnectStatus200Response() *GetTelegramConnectStatus200Response`

NewGetTelegramConnectStatus200Response instantiates a new GetTelegramConnectStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTelegramConnectStatus200ResponseWithDefaults

`func NewGetTelegramConnectStatus200ResponseWithDefaults() *GetTelegramConnectStatus200Response`

NewGetTelegramConnectStatus200ResponseWithDefaults instantiates a new GetTelegramConnectStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GetTelegramConnectStatus200Response) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetTelegramConnectStatus200Response) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetTelegramConnectStatus200Response) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetTelegramConnectStatus200Response) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetTelegramConnectStatus200Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetTelegramConnectStatus200Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetTelegramConnectStatus200Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetTelegramConnectStatus200Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetTelegramConnectStatus200Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetTelegramConnectStatus200Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetTelegramConnectStatus200Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetTelegramConnectStatus200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetBotUsername

`func (o *GetTelegramConnectStatus200Response) GetBotUsername() string`

GetBotUsername returns the BotUsername field if non-nil, zero value otherwise.

### GetBotUsernameOk

`func (o *GetTelegramConnectStatus200Response) GetBotUsernameOk() (*string, bool)`

GetBotUsernameOk returns a tuple with the BotUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotUsername

`func (o *GetTelegramConnectStatus200Response) SetBotUsername(v string)`

SetBotUsername sets BotUsername field to given value.

### HasBotUsername

`func (o *GetTelegramConnectStatus200Response) HasBotUsername() bool`

HasBotUsername returns a boolean if a field has been set.

### GetInstructions

`func (o *GetTelegramConnectStatus200Response) GetInstructions() []string`

GetInstructions returns the Instructions field if non-nil, zero value otherwise.

### GetInstructionsOk

`func (o *GetTelegramConnectStatus200Response) GetInstructionsOk() (*[]string, bool)`

GetInstructionsOk returns a tuple with the Instructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstructions

`func (o *GetTelegramConnectStatus200Response) SetInstructions(v []string)`

SetInstructions sets Instructions field to given value.

### HasInstructions

`func (o *GetTelegramConnectStatus200Response) HasInstructions() bool`

HasInstructions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


