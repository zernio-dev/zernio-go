# InitiateTelegramConnectRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChatId** | **string** | The Telegram chat ID. Numeric ID (e.g. \&quot;-1001234567890\&quot;) or username with @ prefix (e.g. \&quot;@mychannel\&quot;). | 
**ProfileId** | **string** | The profile ID to connect the account to | 

## Methods

### NewInitiateTelegramConnectRequest

`func NewInitiateTelegramConnectRequest(chatId string, profileId string, ) *InitiateTelegramConnectRequest`

NewInitiateTelegramConnectRequest instantiates a new InitiateTelegramConnectRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateTelegramConnectRequestWithDefaults

`func NewInitiateTelegramConnectRequestWithDefaults() *InitiateTelegramConnectRequest`

NewInitiateTelegramConnectRequestWithDefaults instantiates a new InitiateTelegramConnectRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChatId

`func (o *InitiateTelegramConnectRequest) GetChatId() string`

GetChatId returns the ChatId field if non-nil, zero value otherwise.

### GetChatIdOk

`func (o *InitiateTelegramConnectRequest) GetChatIdOk() (*string, bool)`

GetChatIdOk returns a tuple with the ChatId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatId

`func (o *InitiateTelegramConnectRequest) SetChatId(v string)`

SetChatId sets ChatId field to given value.


### GetProfileId

`func (o *InitiateTelegramConnectRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *InitiateTelegramConnectRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *InitiateTelegramConnectRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


