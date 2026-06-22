# PlatformTargetAccountId

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

## Methods

### NewPlatformTargetAccountId

`func NewPlatformTargetAccountId(id string, platform string, profileId SocialAccountProfileId, isActive bool, ) *PlatformTargetAccountId`

NewPlatformTargetAccountId instantiates a new PlatformTargetAccountId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTargetAccountIdWithDefaults

`func NewPlatformTargetAccountIdWithDefaults() *PlatformTargetAccountId`

NewPlatformTargetAccountIdWithDefaults instantiates a new PlatformTargetAccountId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PlatformTargetAccountId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PlatformTargetAccountId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PlatformTargetAccountId) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *PlatformTargetAccountId) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PlatformTargetAccountId) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PlatformTargetAccountId) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetProfileId

`func (o *PlatformTargetAccountId) GetProfileId() SocialAccountProfileId`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PlatformTargetAccountId) GetProfileIdOk() (*SocialAccountProfileId, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PlatformTargetAccountId) SetProfileId(v SocialAccountProfileId)`

SetProfileId sets ProfileId field to given value.


### GetUsername

`func (o *PlatformTargetAccountId) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PlatformTargetAccountId) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PlatformTargetAccountId) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PlatformTargetAccountId) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *PlatformTargetAccountId) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *PlatformTargetAccountId) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *PlatformTargetAccountId) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *PlatformTargetAccountId) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProfilePicture

`func (o *PlatformTargetAccountId) GetProfilePicture() string`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *PlatformTargetAccountId) GetProfilePictureOk() (*string, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *PlatformTargetAccountId) SetProfilePicture(v string)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *PlatformTargetAccountId) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### GetProfileUrl

`func (o *PlatformTargetAccountId) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *PlatformTargetAccountId) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *PlatformTargetAccountId) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *PlatformTargetAccountId) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetIsActive

`func (o *PlatformTargetAccountId) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PlatformTargetAccountId) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PlatformTargetAccountId) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetFollowersCount

`func (o *PlatformTargetAccountId) GetFollowersCount() float32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *PlatformTargetAccountId) GetFollowersCountOk() (*float32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *PlatformTargetAccountId) SetFollowersCount(v float32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *PlatformTargetAccountId) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetFollowersLastUpdated

`func (o *PlatformTargetAccountId) GetFollowersLastUpdated() time.Time`

GetFollowersLastUpdated returns the FollowersLastUpdated field if non-nil, zero value otherwise.

### GetFollowersLastUpdatedOk

`func (o *PlatformTargetAccountId) GetFollowersLastUpdatedOk() (*time.Time, bool)`

GetFollowersLastUpdatedOk returns a tuple with the FollowersLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersLastUpdated

`func (o *PlatformTargetAccountId) SetFollowersLastUpdated(v time.Time)`

SetFollowersLastUpdated sets FollowersLastUpdated field to given value.

### HasFollowersLastUpdated

`func (o *PlatformTargetAccountId) HasFollowersLastUpdated() bool`

HasFollowersLastUpdated returns a boolean if a field has been set.

### GetParentAccountId

`func (o *PlatformTargetAccountId) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *PlatformTargetAccountId) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *PlatformTargetAccountId) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *PlatformTargetAccountId) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *PlatformTargetAccountId) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PlatformTargetAccountId) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PlatformTargetAccountId) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PlatformTargetAccountId) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *PlatformTargetAccountId) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PlatformTargetAccountId) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PlatformTargetAccountId) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PlatformTargetAccountId) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


