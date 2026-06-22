# AccountWithFollowerStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**ProfileId** | [**SocialAccountProfileId**](SocialAccountProfileId.md) |  | 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ProfilePicture** | Pointer to **string** | URL to the account&#39;s profile picture on the platform. May be null if the platform does not provide one. | [optional] 
**ProfileUrl** | Pointer to **string** | Full profile URL for the connected account on its platform. | [optional] 
**IsActive** | **bool** |  | 
**FollowersCount** | Pointer to **float32** | Follower count (only included if user has analytics add-on) | [optional] 
**FollowersLastUpdated** | Pointer to **time.Time** | Last time follower count was updated (only included if user has analytics add-on) | [optional] 
**ParentAccountId** | Pointer to **string** | Reference to the parent posting SocialAccount. Set for ads accounts that share or derive from a posting account&#39;s OAuth token. null for standalone ads (Google Ads) and all posting accounts.  | [optional] 
**Enabled** | Pointer to **bool** | Whether the user explicitly activated this account. false means the account was created as a side effect (e.g., posting account auto-created when user connected ads first). Posting UI and scheduler ignore accounts with enabled: false.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Platform-specific metadata. Fields vary by platform. For WhatsApp accounts, includes: - qualityRating: Phone number quality rating from Meta (GREEN, YELLOW, RED, or UNKNOWN) - nameStatus: Display name review status (APPROVED, PENDING_REVIEW, DECLINED, or NONE). Messages cannot be sent until the display name is approved by Meta. - messagingLimitTier: Maximum unique business-initiated conversations per 24h rolling window (TIER_250, TIER_1K, TIER_10K, TIER_100K, or TIER_UNLIMITED). Scales automatically as quality rating improves. - verifiedName: Meta-verified business display name - displayPhoneNumber: Formatted phone number (e.g., \&quot;+1 555-123-4567\&quot;) - wabaId: WhatsApp Business Account ID - phoneNumberId: Meta phone number ID  | [optional] 
**CurrentFollowers** | Pointer to **float32** | Current follower count | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Growth** | Pointer to **float32** | Follower change over period | [optional] 
**GrowthPercentage** | Pointer to **float32** | Percentage growth | [optional] 
**DataPoints** | Pointer to **float32** | Number of historical snapshots | [optional] 
**AccountStats** | Pointer to [**AccountWithFollowerStatsAllOfAccountStats**](AccountWithFollowerStatsAllOfAccountStats.md) |  | [optional] 

## Methods

### NewAccountWithFollowerStats

`func NewAccountWithFollowerStats(id string, platform string, profileId SocialAccountProfileId, isActive bool, ) *AccountWithFollowerStats`

NewAccountWithFollowerStats instantiates a new AccountWithFollowerStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountWithFollowerStatsWithDefaults

`func NewAccountWithFollowerStatsWithDefaults() *AccountWithFollowerStats`

NewAccountWithFollowerStatsWithDefaults instantiates a new AccountWithFollowerStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccountWithFollowerStats) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccountWithFollowerStats) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccountWithFollowerStats) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *AccountWithFollowerStats) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AccountWithFollowerStats) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AccountWithFollowerStats) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetProfileId

`func (o *AccountWithFollowerStats) GetProfileId() SocialAccountProfileId`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *AccountWithFollowerStats) GetProfileIdOk() (*SocialAccountProfileId, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *AccountWithFollowerStats) SetProfileId(v SocialAccountProfileId)`

SetProfileId sets ProfileId field to given value.


### GetUsername

`func (o *AccountWithFollowerStats) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AccountWithFollowerStats) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AccountWithFollowerStats) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AccountWithFollowerStats) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *AccountWithFollowerStats) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AccountWithFollowerStats) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AccountWithFollowerStats) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AccountWithFollowerStats) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProfilePicture

`func (o *AccountWithFollowerStats) GetProfilePicture() string`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *AccountWithFollowerStats) GetProfilePictureOk() (*string, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *AccountWithFollowerStats) SetProfilePicture(v string)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *AccountWithFollowerStats) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### GetProfileUrl

`func (o *AccountWithFollowerStats) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *AccountWithFollowerStats) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *AccountWithFollowerStats) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *AccountWithFollowerStats) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetIsActive

`func (o *AccountWithFollowerStats) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AccountWithFollowerStats) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AccountWithFollowerStats) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetFollowersCount

`func (o *AccountWithFollowerStats) GetFollowersCount() float32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *AccountWithFollowerStats) GetFollowersCountOk() (*float32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *AccountWithFollowerStats) SetFollowersCount(v float32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *AccountWithFollowerStats) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetFollowersLastUpdated

`func (o *AccountWithFollowerStats) GetFollowersLastUpdated() time.Time`

GetFollowersLastUpdated returns the FollowersLastUpdated field if non-nil, zero value otherwise.

### GetFollowersLastUpdatedOk

`func (o *AccountWithFollowerStats) GetFollowersLastUpdatedOk() (*time.Time, bool)`

GetFollowersLastUpdatedOk returns a tuple with the FollowersLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersLastUpdated

`func (o *AccountWithFollowerStats) SetFollowersLastUpdated(v time.Time)`

SetFollowersLastUpdated sets FollowersLastUpdated field to given value.

### HasFollowersLastUpdated

`func (o *AccountWithFollowerStats) HasFollowersLastUpdated() bool`

HasFollowersLastUpdated returns a boolean if a field has been set.

### GetParentAccountId

`func (o *AccountWithFollowerStats) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *AccountWithFollowerStats) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *AccountWithFollowerStats) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *AccountWithFollowerStats) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *AccountWithFollowerStats) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AccountWithFollowerStats) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AccountWithFollowerStats) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AccountWithFollowerStats) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *AccountWithFollowerStats) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AccountWithFollowerStats) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AccountWithFollowerStats) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AccountWithFollowerStats) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetCurrentFollowers

`func (o *AccountWithFollowerStats) GetCurrentFollowers() float32`

GetCurrentFollowers returns the CurrentFollowers field if non-nil, zero value otherwise.

### GetCurrentFollowersOk

`func (o *AccountWithFollowerStats) GetCurrentFollowersOk() (*float32, bool)`

GetCurrentFollowersOk returns a tuple with the CurrentFollowers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFollowers

`func (o *AccountWithFollowerStats) SetCurrentFollowers(v float32)`

SetCurrentFollowers sets CurrentFollowers field to given value.

### HasCurrentFollowers

`func (o *AccountWithFollowerStats) HasCurrentFollowers() bool`

HasCurrentFollowers returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AccountWithFollowerStats) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AccountWithFollowerStats) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AccountWithFollowerStats) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AccountWithFollowerStats) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetGrowth

`func (o *AccountWithFollowerStats) GetGrowth() float32`

GetGrowth returns the Growth field if non-nil, zero value otherwise.

### GetGrowthOk

`func (o *AccountWithFollowerStats) GetGrowthOk() (*float32, bool)`

GetGrowthOk returns a tuple with the Growth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowth

`func (o *AccountWithFollowerStats) SetGrowth(v float32)`

SetGrowth sets Growth field to given value.

### HasGrowth

`func (o *AccountWithFollowerStats) HasGrowth() bool`

HasGrowth returns a boolean if a field has been set.

### GetGrowthPercentage

`func (o *AccountWithFollowerStats) GetGrowthPercentage() float32`

GetGrowthPercentage returns the GrowthPercentage field if non-nil, zero value otherwise.

### GetGrowthPercentageOk

`func (o *AccountWithFollowerStats) GetGrowthPercentageOk() (*float32, bool)`

GetGrowthPercentageOk returns a tuple with the GrowthPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrowthPercentage

`func (o *AccountWithFollowerStats) SetGrowthPercentage(v float32)`

SetGrowthPercentage sets GrowthPercentage field to given value.

### HasGrowthPercentage

`func (o *AccountWithFollowerStats) HasGrowthPercentage() bool`

HasGrowthPercentage returns a boolean if a field has been set.

### GetDataPoints

`func (o *AccountWithFollowerStats) GetDataPoints() float32`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *AccountWithFollowerStats) GetDataPointsOk() (*float32, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *AccountWithFollowerStats) SetDataPoints(v float32)`

SetDataPoints sets DataPoints field to given value.

### HasDataPoints

`func (o *AccountWithFollowerStats) HasDataPoints() bool`

HasDataPoints returns a boolean if a field has been set.

### GetAccountStats

`func (o *AccountWithFollowerStats) GetAccountStats() AccountWithFollowerStatsAllOfAccountStats`

GetAccountStats returns the AccountStats field if non-nil, zero value otherwise.

### GetAccountStatsOk

`func (o *AccountWithFollowerStats) GetAccountStatsOk() (*AccountWithFollowerStatsAllOfAccountStats, bool)`

GetAccountStatsOk returns a tuple with the AccountStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountStats

`func (o *AccountWithFollowerStats) SetAccountStats(v AccountWithFollowerStatsAllOfAccountStats)`

SetAccountStats sets AccountStats field to given value.

### HasAccountStats

`func (o *AccountWithFollowerStats) HasAccountStats() bool`

HasAccountStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


