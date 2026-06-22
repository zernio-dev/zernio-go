# WebhookPayloadAccountConnectedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account&#39;s unique identifier (same as used in /v1/accounts/{accountId}) | 
**ProfileId** | **string** | The profile&#39;s unique identifier this account belongs to | 
**Platform** | **string** |  | 
**Username** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewWebhookPayloadAccountConnectedAccount

`func NewWebhookPayloadAccountConnectedAccount(accountId string, profileId string, platform string, username string, ) *WebhookPayloadAccountConnectedAccount`

NewWebhookPayloadAccountConnectedAccount instantiates a new WebhookPayloadAccountConnectedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountConnectedAccountWithDefaults

`func NewWebhookPayloadAccountConnectedAccountWithDefaults() *WebhookPayloadAccountConnectedAccount`

NewWebhookPayloadAccountConnectedAccountWithDefaults instantiates a new WebhookPayloadAccountConnectedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *WebhookPayloadAccountConnectedAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookPayloadAccountConnectedAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookPayloadAccountConnectedAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProfileId

`func (o *WebhookPayloadAccountConnectedAccount) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *WebhookPayloadAccountConnectedAccount) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *WebhookPayloadAccountConnectedAccount) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetPlatform

`func (o *WebhookPayloadAccountConnectedAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebhookPayloadAccountConnectedAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebhookPayloadAccountConnectedAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUsername

`func (o *WebhookPayloadAccountConnectedAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookPayloadAccountConnectedAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookPayloadAccountConnectedAccount) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *WebhookPayloadAccountConnectedAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WebhookPayloadAccountConnectedAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WebhookPayloadAccountConnectedAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WebhookPayloadAccountConnectedAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


