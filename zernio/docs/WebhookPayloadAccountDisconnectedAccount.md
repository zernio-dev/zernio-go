# WebhookPayloadAccountDisconnectedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account&#39;s unique identifier (same as used in /v1/accounts/{accountId}) | 
**ProfileId** | **string** | The profile&#39;s unique identifier this account belongs to | 
**Platform** | **string** |  | 
**Username** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**DisconnectionType** | **string** | Whether the disconnection was intentional (user action) or unintentional (token expired/revoked) | 
**Reason** | **string** | Human-readable reason for the disconnection | 

## Methods

### NewWebhookPayloadAccountDisconnectedAccount

`func NewWebhookPayloadAccountDisconnectedAccount(accountId string, profileId string, platform string, username string, disconnectionType string, reason string, ) *WebhookPayloadAccountDisconnectedAccount`

NewWebhookPayloadAccountDisconnectedAccount instantiates a new WebhookPayloadAccountDisconnectedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountDisconnectedAccountWithDefaults

`func NewWebhookPayloadAccountDisconnectedAccountWithDefaults() *WebhookPayloadAccountDisconnectedAccount`

NewWebhookPayloadAccountDisconnectedAccountWithDefaults instantiates a new WebhookPayloadAccountDisconnectedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *WebhookPayloadAccountDisconnectedAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookPayloadAccountDisconnectedAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProfileId

`func (o *WebhookPayloadAccountDisconnectedAccount) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *WebhookPayloadAccountDisconnectedAccount) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetPlatform

`func (o *WebhookPayloadAccountDisconnectedAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebhookPayloadAccountDisconnectedAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUsername

`func (o *WebhookPayloadAccountDisconnectedAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookPayloadAccountDisconnectedAccount) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *WebhookPayloadAccountDisconnectedAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WebhookPayloadAccountDisconnectedAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WebhookPayloadAccountDisconnectedAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisconnectionType

`func (o *WebhookPayloadAccountDisconnectedAccount) GetDisconnectionType() string`

GetDisconnectionType returns the DisconnectionType field if non-nil, zero value otherwise.

### GetDisconnectionTypeOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetDisconnectionTypeOk() (*string, bool)`

GetDisconnectionTypeOk returns a tuple with the DisconnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectionType

`func (o *WebhookPayloadAccountDisconnectedAccount) SetDisconnectionType(v string)`

SetDisconnectionType sets DisconnectionType field to given value.


### GetReason

`func (o *WebhookPayloadAccountDisconnectedAccount) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *WebhookPayloadAccountDisconnectedAccount) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *WebhookPayloadAccountDisconnectedAccount) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


