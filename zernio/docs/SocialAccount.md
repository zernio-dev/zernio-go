# SocialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**ProfileId** | [**SocialAccountProfileId**](SocialAccountProfileId.md) |  | 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ProfilePicture** | Pointer to **NullableString** | URL to the account&#39;s profile picture on the platform. May be null if the platform does not provide one. | [optional] 
**ProfileUrl** | Pointer to **string** | Full profile URL for the connected account on its platform. | [optional] 
**IsActive** | **bool** |  | 
**FollowersCount** | Pointer to **float32** | Follower count (only included if user has analytics add-on) | [optional] 
**FollowersLastUpdated** | Pointer to **time.Time** | Last time follower count was updated (only included if user has analytics add-on) | [optional] 
**ParentAccountId** | Pointer to **NullableString** | Reference to the parent posting SocialAccount. Set for ads accounts that share or derive from a posting account&#39;s OAuth token. null for standalone ads (Google Ads) and all posting accounts.  | [optional] 
**Enabled** | Pointer to **bool** | Whether the user explicitly activated this account. false means the account was created as a side effect (e.g., posting account auto-created when user connected ads first). Posting UI and scheduler ignore accounts with enabled: false.  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Platform-specific metadata. Fields vary by platform. For WhatsApp accounts, includes: - qualityRating: Phone number quality rating from Meta (GREEN, YELLOW, RED, or UNKNOWN) - nameStatus: Display name review status (APPROVED, PENDING_REVIEW, DECLINED, or NONE). Messages cannot be sent until the display name is approved by Meta. - messagingLimitTier: Maximum unique business-initiated conversations per 24h rolling window (TIER_250, TIER_1K, TIER_10K, TIER_100K, or TIER_UNLIMITED). Scales automatically as quality rating improves. - verifiedName: Meta-verified business display name - displayPhoneNumber: Formatted phone number (e.g., \&quot;+1 555-123-4567\&quot;) - wabaId: WhatsApp Business Account ID - phoneNumberId: Meta phone number ID  | [optional] 

## Methods

### NewSocialAccount

`func NewSocialAccount(id string, platform string, profileId SocialAccountProfileId, isActive bool, ) *SocialAccount`

NewSocialAccount instantiates a new SocialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialAccountWithDefaults

`func NewSocialAccountWithDefaults() *SocialAccount`

NewSocialAccountWithDefaults instantiates a new SocialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialAccount) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *SocialAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SocialAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SocialAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetProfileId

`func (o *SocialAccount) GetProfileId() SocialAccountProfileId`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SocialAccount) GetProfileIdOk() (*SocialAccountProfileId, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SocialAccount) SetProfileId(v SocialAccountProfileId)`

SetProfileId sets ProfileId field to given value.


### GetUsername

`func (o *SocialAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SocialAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SocialAccount) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SocialAccount) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *SocialAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SocialAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SocialAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SocialAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProfilePicture

`func (o *SocialAccount) GetProfilePicture() string`

GetProfilePicture returns the ProfilePicture field if non-nil, zero value otherwise.

### GetProfilePictureOk

`func (o *SocialAccount) GetProfilePictureOk() (*string, bool)`

GetProfilePictureOk returns a tuple with the ProfilePicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePicture

`func (o *SocialAccount) SetProfilePicture(v string)`

SetProfilePicture sets ProfilePicture field to given value.

### HasProfilePicture

`func (o *SocialAccount) HasProfilePicture() bool`

HasProfilePicture returns a boolean if a field has been set.

### SetProfilePictureNil

`func (o *SocialAccount) SetProfilePictureNil(b bool)`

 SetProfilePictureNil sets the value for ProfilePicture to be an explicit nil

### UnsetProfilePicture
`func (o *SocialAccount) UnsetProfilePicture()`

UnsetProfilePicture ensures that no value is present for ProfilePicture, not even an explicit nil
### GetProfileUrl

`func (o *SocialAccount) GetProfileUrl() string`

GetProfileUrl returns the ProfileUrl field if non-nil, zero value otherwise.

### GetProfileUrlOk

`func (o *SocialAccount) GetProfileUrlOk() (*string, bool)`

GetProfileUrlOk returns a tuple with the ProfileUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileUrl

`func (o *SocialAccount) SetProfileUrl(v string)`

SetProfileUrl sets ProfileUrl field to given value.

### HasProfileUrl

`func (o *SocialAccount) HasProfileUrl() bool`

HasProfileUrl returns a boolean if a field has been set.

### GetIsActive

`func (o *SocialAccount) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *SocialAccount) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *SocialAccount) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetFollowersCount

`func (o *SocialAccount) GetFollowersCount() float32`

GetFollowersCount returns the FollowersCount field if non-nil, zero value otherwise.

### GetFollowersCountOk

`func (o *SocialAccount) GetFollowersCountOk() (*float32, bool)`

GetFollowersCountOk returns a tuple with the FollowersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersCount

`func (o *SocialAccount) SetFollowersCount(v float32)`

SetFollowersCount sets FollowersCount field to given value.

### HasFollowersCount

`func (o *SocialAccount) HasFollowersCount() bool`

HasFollowersCount returns a boolean if a field has been set.

### GetFollowersLastUpdated

`func (o *SocialAccount) GetFollowersLastUpdated() time.Time`

GetFollowersLastUpdated returns the FollowersLastUpdated field if non-nil, zero value otherwise.

### GetFollowersLastUpdatedOk

`func (o *SocialAccount) GetFollowersLastUpdatedOk() (*time.Time, bool)`

GetFollowersLastUpdatedOk returns a tuple with the FollowersLastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowersLastUpdated

`func (o *SocialAccount) SetFollowersLastUpdated(v time.Time)`

SetFollowersLastUpdated sets FollowersLastUpdated field to given value.

### HasFollowersLastUpdated

`func (o *SocialAccount) HasFollowersLastUpdated() bool`

HasFollowersLastUpdated returns a boolean if a field has been set.

### GetParentAccountId

`func (o *SocialAccount) GetParentAccountId() string`

GetParentAccountId returns the ParentAccountId field if non-nil, zero value otherwise.

### GetParentAccountIdOk

`func (o *SocialAccount) GetParentAccountIdOk() (*string, bool)`

GetParentAccountIdOk returns a tuple with the ParentAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentAccountId

`func (o *SocialAccount) SetParentAccountId(v string)`

SetParentAccountId sets ParentAccountId field to given value.

### HasParentAccountId

`func (o *SocialAccount) HasParentAccountId() bool`

HasParentAccountId returns a boolean if a field has been set.

### SetParentAccountIdNil

`func (o *SocialAccount) SetParentAccountIdNil(b bool)`

 SetParentAccountIdNil sets the value for ParentAccountId to be an explicit nil

### UnsetParentAccountId
`func (o *SocialAccount) UnsetParentAccountId()`

UnsetParentAccountId ensures that no value is present for ParentAccountId, not even an explicit nil
### GetEnabled

`func (o *SocialAccount) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SocialAccount) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SocialAccount) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SocialAccount) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *SocialAccount) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SocialAccount) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SocialAccount) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SocialAccount) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


