# TelegramPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParseMode** | Pointer to **string** | Text formatting mode for the message (default is HTML) | [optional] 
**DisableWebPagePreview** | Pointer to **bool** | Disable link preview generation for URLs in the message | [optional] 
**DisableNotification** | Pointer to **bool** | Send the message silently (users will receive notification without sound) | [optional] 
**ProtectContent** | Pointer to **bool** | Protect message content from forwarding and saving | [optional] 

## Methods

### NewTelegramPlatformData

`func NewTelegramPlatformData() *TelegramPlatformData`

NewTelegramPlatformData instantiates a new TelegramPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelegramPlatformDataWithDefaults

`func NewTelegramPlatformDataWithDefaults() *TelegramPlatformData`

NewTelegramPlatformDataWithDefaults instantiates a new TelegramPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParseMode

`func (o *TelegramPlatformData) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *TelegramPlatformData) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *TelegramPlatformData) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *TelegramPlatformData) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetDisableWebPagePreview

`func (o *TelegramPlatformData) GetDisableWebPagePreview() bool`

GetDisableWebPagePreview returns the DisableWebPagePreview field if non-nil, zero value otherwise.

### GetDisableWebPagePreviewOk

`func (o *TelegramPlatformData) GetDisableWebPagePreviewOk() (*bool, bool)`

GetDisableWebPagePreviewOk returns a tuple with the DisableWebPagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableWebPagePreview

`func (o *TelegramPlatformData) SetDisableWebPagePreview(v bool)`

SetDisableWebPagePreview sets DisableWebPagePreview field to given value.

### HasDisableWebPagePreview

`func (o *TelegramPlatformData) HasDisableWebPagePreview() bool`

HasDisableWebPagePreview returns a boolean if a field has been set.

### GetDisableNotification

`func (o *TelegramPlatformData) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *TelegramPlatformData) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *TelegramPlatformData) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *TelegramPlatformData) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *TelegramPlatformData) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *TelegramPlatformData) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *TelegramPlatformData) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *TelegramPlatformData) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


