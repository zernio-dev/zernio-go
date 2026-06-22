# InitiateTelegramConnect200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**InitiateTelegramConnect200ResponseAccount**](InitiateTelegramConnect200ResponseAccount.md) |  | [optional] 

## Methods

### NewInitiateTelegramConnect200Response

`func NewInitiateTelegramConnect200Response() *InitiateTelegramConnect200Response`

NewInitiateTelegramConnect200Response instantiates a new InitiateTelegramConnect200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateTelegramConnect200ResponseWithDefaults

`func NewInitiateTelegramConnect200ResponseWithDefaults() *InitiateTelegramConnect200Response`

NewInitiateTelegramConnect200ResponseWithDefaults instantiates a new InitiateTelegramConnect200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *InitiateTelegramConnect200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InitiateTelegramConnect200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InitiateTelegramConnect200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *InitiateTelegramConnect200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetAccount

`func (o *InitiateTelegramConnect200Response) GetAccount() InitiateTelegramConnect200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InitiateTelegramConnect200Response) GetAccountOk() (*InitiateTelegramConnect200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InitiateTelegramConnect200Response) SetAccount(v InitiateTelegramConnect200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InitiateTelegramConnect200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


