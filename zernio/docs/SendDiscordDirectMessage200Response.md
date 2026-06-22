# SendDiscordDirectMessage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageId** | Pointer to **string** | Discord message snowflake ID | [optional] 
**ChannelId** | Pointer to **string** | DM channel snowflake (Discord auto-creates one per recipient pair) | [optional] 
**Url** | Pointer to **string** | Direct link to the message — uses Discord&#39;s @me path for DMs | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Recipient** | Pointer to [**SendDiscordDirectMessage200ResponseRecipient**](SendDiscordDirectMessage200ResponseRecipient.md) |  | [optional] 
**Account** | Pointer to [**SendDiscordDirectMessage200ResponseAccount**](SendDiscordDirectMessage200ResponseAccount.md) |  | [optional] 

## Methods

### NewSendDiscordDirectMessage200Response

`func NewSendDiscordDirectMessage200Response() *SendDiscordDirectMessage200Response`

NewSendDiscordDirectMessage200Response instantiates a new SendDiscordDirectMessage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendDiscordDirectMessage200ResponseWithDefaults

`func NewSendDiscordDirectMessage200ResponseWithDefaults() *SendDiscordDirectMessage200Response`

NewSendDiscordDirectMessage200ResponseWithDefaults instantiates a new SendDiscordDirectMessage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageId

`func (o *SendDiscordDirectMessage200Response) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *SendDiscordDirectMessage200Response) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *SendDiscordDirectMessage200Response) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *SendDiscordDirectMessage200Response) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetChannelId

`func (o *SendDiscordDirectMessage200Response) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *SendDiscordDirectMessage200Response) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *SendDiscordDirectMessage200Response) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *SendDiscordDirectMessage200Response) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetUrl

`func (o *SendDiscordDirectMessage200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SendDiscordDirectMessage200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SendDiscordDirectMessage200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SendDiscordDirectMessage200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTimestamp

`func (o *SendDiscordDirectMessage200Response) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SendDiscordDirectMessage200Response) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SendDiscordDirectMessage200Response) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SendDiscordDirectMessage200Response) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetRecipient

`func (o *SendDiscordDirectMessage200Response) GetRecipient() SendDiscordDirectMessage200ResponseRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *SendDiscordDirectMessage200Response) GetRecipientOk() (*SendDiscordDirectMessage200ResponseRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *SendDiscordDirectMessage200Response) SetRecipient(v SendDiscordDirectMessage200ResponseRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *SendDiscordDirectMessage200Response) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetAccount

`func (o *SendDiscordDirectMessage200Response) GetAccount() SendDiscordDirectMessage200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SendDiscordDirectMessage200Response) GetAccountOk() (*SendDiscordDirectMessage200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SendDiscordDirectMessage200Response) SetAccount(v SendDiscordDirectMessage200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SendDiscordDirectMessage200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


