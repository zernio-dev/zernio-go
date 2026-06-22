# WebhookPayloadAccountAdsInitialSyncCompletedAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account&#39;s unique identifier (same as used in /v1/accounts/{accountId}) | 
**ProfileId** | **string** | The profile&#39;s unique identifier this account belongs to | 
**Platform** | **string** |  | 
**Username** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**PlatformUserId** | Pointer to **string** | The platform-side account/ad-account ID (e.g. Meta ad account ID). | [optional] 
**ProfilePicture** | Pointer to **string** | URL of the account&#39;s profile picture, when available. | [optional] 
**PlatformAdAccountId** | Pointer to **string** | When the consumer scoped the connect call to a single ad account, this echoes that ID back so the webhook can be correlated to the originating connect request without consulting the consumer&#39;s DB. Meta uses the &#x60;act_*&#x60; shape.  | [optional] 
**PlatformAdAccountIds** | Pointer to **[]string** | Every ad-account ID that the connected token could see at discovery time. Useful for \&quot;we synced ads from these accounts\&quot; UX without a follow-up API call. Empty array when the token had no ad-account visibility.  | [optional] 

## Methods

### NewWebhookPayloadAccountAdsInitialSyncCompletedAccount

`func NewWebhookPayloadAccountAdsInitialSyncCompletedAccount(accountId string, profileId string, platform string, username string, ) *WebhookPayloadAccountAdsInitialSyncCompletedAccount`

NewWebhookPayloadAccountAdsInitialSyncCompletedAccount instantiates a new WebhookPayloadAccountAdsInitialSyncCompletedAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAccountAdsInitialSyncCompletedAccountWithDefaults

`func NewWebhookPayloadAccountAdsInitialSyncCompletedAccountWithDefaults() *WebhookPayloadAccountAdsInitialSyncCompletedAccount`

NewWebhookPayloadAccountAdsInitialSyncCompletedAccountWithDefaults instantiates a new WebhookPayloadAccountAdsInitialSyncCompletedAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetProfileId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetPlatform

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUsername

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPlatformUserId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformUserId() string`

GetPlatformUserId returns the PlatformUserId field if non-nil, zero value otherwise.

### GetPlatformUserIdOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformUserIdOk() (*string, bool)`

GetPlatformUserIdOk returns a tuple with the PlatformUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUserId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetPlatformUserId(v string)`

SetPlatformUserId sets PlatformUserId field to given value.

### HasPlatformUserId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) HasPlatformUserId() bool`

HasPlatformUserId returns a boolean if a field has been set.

### GetProfilePicture

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetProfilePicture() string`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetProfilePictureOk() (*string, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetProfilePicture(v string)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### GetPlatformAdAccountId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformAdAccountId() string`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformAdAccountIdOk() (*string, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetPlatformAdAccountId(v string)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.

### HasPlatformAdAccountId

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) HasPlatformAdAccountId() bool`

HasPlatformAdAccountId returns a boolean if a field has been set.

### GetPlatformAdAccountIds

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformAdAccountIds() []string`

GetPlatformAdAccountIds returns the PlatformAdAccountIds field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdsOk

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) GetPlatformAdAccountIdsOk() (*[]string, bool)`

GetPlatformAdAccountIdsOk returns a tuple with the PlatformAdAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountIds

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) SetPlatformAdAccountIds(v []string)`

SetPlatformAdAccountIds sets PlatformAdAccountIds field to given value.

### HasPlatformAdAccountIds

`func (o *WebhookPayloadAccountAdsInitialSyncCompletedAccount) HasPlatformAdAccountIds() bool`

HasPlatformAdAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


