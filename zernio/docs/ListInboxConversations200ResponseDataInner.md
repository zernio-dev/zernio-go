# ListInboxConversations200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**ParticipantId** | Pointer to **string** |  | [optional] 
**ParticipantName** | Pointer to **string** |  | [optional] 
**ParticipantPicture** | Pointer to **NullableString** |  | [optional] 
**ParticipantVerifiedType** | Pointer to **NullableString** | X/Twitter verified badge type. Only present for Twitter/X conversations. | [optional] 
**LastMessage** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UnreadCount** | Pointer to **NullableInt32** | Number of unread messages | [optional] 
**Url** | Pointer to **NullableString** | Direct link to open the conversation on the platform (if available) | [optional] 
**InstagramProfile** | Pointer to [**ListInboxConversations200ResponseDataInnerInstagramProfile**](ListInboxConversations200ResponseDataInnerInstagramProfile.md) |  | [optional] 

## Methods

### NewListInboxConversations200ResponseDataInner

`func NewListInboxConversations200ResponseDataInner() *ListInboxConversations200ResponseDataInner`

NewListInboxConversations200ResponseDataInner instantiates a new ListInboxConversations200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversations200ResponseDataInnerWithDefaults

`func NewListInboxConversations200ResponseDataInnerWithDefaults() *ListInboxConversations200ResponseDataInner`

NewListInboxConversations200ResponseDataInnerWithDefaults instantiates a new ListInboxConversations200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInboxConversations200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInboxConversations200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInboxConversations200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListInboxConversations200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListInboxConversations200ResponseDataInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInboxConversations200ResponseDataInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInboxConversations200ResponseDataInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInboxConversations200ResponseDataInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *ListInboxConversations200ResponseDataInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInboxConversations200ResponseDataInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInboxConversations200ResponseDataInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInboxConversations200ResponseDataInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUsername

`func (o *ListInboxConversations200ResponseDataInner) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *ListInboxConversations200ResponseDataInner) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *ListInboxConversations200ResponseDataInner) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *ListInboxConversations200ResponseDataInner) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetParticipantId

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *ListInboxConversations200ResponseDataInner) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *ListInboxConversations200ResponseDataInner) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetParticipantName

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *ListInboxConversations200ResponseDataInner) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *ListInboxConversations200ResponseDataInner) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### GetParticipantPicture

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantPicture() string`

GetParticipantPicture returns the ParticipantPicture field if non-nil, zero value otherwise.

### GetParticipantPictureOk

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantPictureOk() (*string, bool)`

GetParticipantPictureOk returns a tuple with the ParticipantPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantPicture

`func (o *ListInboxConversations200ResponseDataInner) SetParticipantPicture(v string)`

SetParticipantPicture sets ParticipantPicture field to given value.

### HasParticipantPicture

`func (o *ListInboxConversations200ResponseDataInner) HasParticipantPicture() bool`

HasParticipantPicture returns a boolean if a field has been set.

### SetParticipantPictureNil

`func (o *ListInboxConversations200ResponseDataInner) SetParticipantPictureNil(b bool)`

 SetParticipantPictureNil sets the value for ParticipantPicture to be an explicit nil

### UnsetParticipantPicture
`func (o *ListInboxConversations200ResponseDataInner) UnsetParticipantPicture()`

UnsetParticipantPicture ensures that no value is present for ParticipantPicture, not even an explicit nil
### GetParticipantVerifiedType

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantVerifiedType() string`

GetParticipantVerifiedType returns the ParticipantVerifiedType field if non-nil, zero value otherwise.

### GetParticipantVerifiedTypeOk

`func (o *ListInboxConversations200ResponseDataInner) GetParticipantVerifiedTypeOk() (*string, bool)`

GetParticipantVerifiedTypeOk returns a tuple with the ParticipantVerifiedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantVerifiedType

`func (o *ListInboxConversations200ResponseDataInner) SetParticipantVerifiedType(v string)`

SetParticipantVerifiedType sets ParticipantVerifiedType field to given value.

### HasParticipantVerifiedType

`func (o *ListInboxConversations200ResponseDataInner) HasParticipantVerifiedType() bool`

HasParticipantVerifiedType returns a boolean if a field has been set.

### SetParticipantVerifiedTypeNil

`func (o *ListInboxConversations200ResponseDataInner) SetParticipantVerifiedTypeNil(b bool)`

 SetParticipantVerifiedTypeNil sets the value for ParticipantVerifiedType to be an explicit nil

### UnsetParticipantVerifiedType
`func (o *ListInboxConversations200ResponseDataInner) UnsetParticipantVerifiedType()`

UnsetParticipantVerifiedType ensures that no value is present for ParticipantVerifiedType, not even an explicit nil
### GetLastMessage

`func (o *ListInboxConversations200ResponseDataInner) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *ListInboxConversations200ResponseDataInner) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *ListInboxConversations200ResponseDataInner) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *ListInboxConversations200ResponseDataInner) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *ListInboxConversations200ResponseDataInner) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *ListInboxConversations200ResponseDataInner) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *ListInboxConversations200ResponseDataInner) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *ListInboxConversations200ResponseDataInner) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetStatus

`func (o *ListInboxConversations200ResponseDataInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInboxConversations200ResponseDataInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInboxConversations200ResponseDataInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListInboxConversations200ResponseDataInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUnreadCount

`func (o *ListInboxConversations200ResponseDataInner) GetUnreadCount() int32`

GetUnreadCount returns the UnreadCount field if non-nil, zero value otherwise.

### GetUnreadCountOk

`func (o *ListInboxConversations200ResponseDataInner) GetUnreadCountOk() (*int32, bool)`

GetUnreadCountOk returns a tuple with the UnreadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadCount

`func (o *ListInboxConversations200ResponseDataInner) SetUnreadCount(v int32)`

SetUnreadCount sets UnreadCount field to given value.

### HasUnreadCount

`func (o *ListInboxConversations200ResponseDataInner) HasUnreadCount() bool`

HasUnreadCount returns a boolean if a field has been set.

### SetUnreadCountNil

`func (o *ListInboxConversations200ResponseDataInner) SetUnreadCountNil(b bool)`

 SetUnreadCountNil sets the value for UnreadCount to be an explicit nil

### UnsetUnreadCount
`func (o *ListInboxConversations200ResponseDataInner) UnsetUnreadCount()`

UnsetUnreadCount ensures that no value is present for UnreadCount, not even an explicit nil
### GetUrl

`func (o *ListInboxConversations200ResponseDataInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListInboxConversations200ResponseDataInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListInboxConversations200ResponseDataInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListInboxConversations200ResponseDataInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *ListInboxConversations200ResponseDataInner) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ListInboxConversations200ResponseDataInner) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetInstagramProfile

`func (o *ListInboxConversations200ResponseDataInner) GetInstagramProfile() ListInboxConversations200ResponseDataInnerInstagramProfile`

GetInstagramProfile returns the InstagramProfile field if non-nil, zero value otherwise.

### GetInstagramProfileOk

`func (o *ListInboxConversations200ResponseDataInner) GetInstagramProfileOk() (*ListInboxConversations200ResponseDataInnerInstagramProfile, bool)`

GetInstagramProfileOk returns a tuple with the InstagramProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramProfile

`func (o *ListInboxConversations200ResponseDataInner) SetInstagramProfile(v ListInboxConversations200ResponseDataInnerInstagramProfile)`

SetInstagramProfile sets InstagramProfile field to given value.

### HasInstagramProfile

`func (o *ListInboxConversations200ResponseDataInner) HasInstagramProfile() bool`

HasInstagramProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


