# GetInboxConversation200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ParticipantName** | Pointer to **string** |  | [optional] 
**ParticipantId** | Pointer to **string** |  | [optional] 
**ParticipantVerifiedType** | Pointer to **NullableString** | X/Twitter verified badge type. Only present for Twitter/X conversations. | [optional] 
**LastMessage** | Pointer to **string** |  | [optional] 
**LastMessageAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedTime** | Pointer to **time.Time** |  | [optional] 
**Participants** | Pointer to [**[]UpdateFacebookPage200ResponseSelectedPage**](UpdateFacebookPage200ResponseSelectedPage.md) |  | [optional] 
**InstagramProfile** | Pointer to [**ListInboxConversations200ResponseDataInnerInstagramProfile**](ListInboxConversations200ResponseDataInnerInstagramProfile.md) |  | [optional] 

## Methods

### NewGetInboxConversation200ResponseData

`func NewGetInboxConversation200ResponseData() *GetInboxConversation200ResponseData`

NewGetInboxConversation200ResponseData instantiates a new GetInboxConversation200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversation200ResponseDataWithDefaults

`func NewGetInboxConversation200ResponseDataWithDefaults() *GetInboxConversation200ResponseData`

NewGetInboxConversation200ResponseDataWithDefaults instantiates a new GetInboxConversation200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInboxConversation200ResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxConversation200ResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxConversation200ResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInboxConversation200ResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetInboxConversation200ResponseData) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInboxConversation200ResponseData) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInboxConversation200ResponseData) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInboxConversation200ResponseData) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUsername

`func (o *GetInboxConversation200ResponseData) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *GetInboxConversation200ResponseData) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *GetInboxConversation200ResponseData) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *GetInboxConversation200ResponseData) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetPlatform

`func (o *GetInboxConversation200ResponseData) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxConversation200ResponseData) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxConversation200ResponseData) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetInboxConversation200ResponseData) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *GetInboxConversation200ResponseData) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInboxConversation200ResponseData) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInboxConversation200ResponseData) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInboxConversation200ResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetParticipantName

`func (o *GetInboxConversation200ResponseData) GetParticipantName() string`

GetParticipantName returns the ParticipantName field if non-nil, zero value otherwise.

### GetParticipantNameOk

`func (o *GetInboxConversation200ResponseData) GetParticipantNameOk() (*string, bool)`

GetParticipantNameOk returns a tuple with the ParticipantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantName

`func (o *GetInboxConversation200ResponseData) SetParticipantName(v string)`

SetParticipantName sets ParticipantName field to given value.

### HasParticipantName

`func (o *GetInboxConversation200ResponseData) HasParticipantName() bool`

HasParticipantName returns a boolean if a field has been set.

### GetParticipantId

`func (o *GetInboxConversation200ResponseData) GetParticipantId() string`

GetParticipantId returns the ParticipantId field if non-nil, zero value otherwise.

### GetParticipantIdOk

`func (o *GetInboxConversation200ResponseData) GetParticipantIdOk() (*string, bool)`

GetParticipantIdOk returns a tuple with the ParticipantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantId

`func (o *GetInboxConversation200ResponseData) SetParticipantId(v string)`

SetParticipantId sets ParticipantId field to given value.

### HasParticipantId

`func (o *GetInboxConversation200ResponseData) HasParticipantId() bool`

HasParticipantId returns a boolean if a field has been set.

### GetParticipantVerifiedType

`func (o *GetInboxConversation200ResponseData) GetParticipantVerifiedType() string`

GetParticipantVerifiedType returns the ParticipantVerifiedType field if non-nil, zero value otherwise.

### GetParticipantVerifiedTypeOk

`func (o *GetInboxConversation200ResponseData) GetParticipantVerifiedTypeOk() (*string, bool)`

GetParticipantVerifiedTypeOk returns a tuple with the ParticipantVerifiedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantVerifiedType

`func (o *GetInboxConversation200ResponseData) SetParticipantVerifiedType(v string)`

SetParticipantVerifiedType sets ParticipantVerifiedType field to given value.

### HasParticipantVerifiedType

`func (o *GetInboxConversation200ResponseData) HasParticipantVerifiedType() bool`

HasParticipantVerifiedType returns a boolean if a field has been set.

### SetParticipantVerifiedTypeNil

`func (o *GetInboxConversation200ResponseData) SetParticipantVerifiedTypeNil(b bool)`

 SetParticipantVerifiedTypeNil sets the value for ParticipantVerifiedType to be an explicit nil

### UnsetParticipantVerifiedType
`func (o *GetInboxConversation200ResponseData) UnsetParticipantVerifiedType()`

UnsetParticipantVerifiedType ensures that no value is present for ParticipantVerifiedType, not even an explicit nil
### GetLastMessage

`func (o *GetInboxConversation200ResponseData) GetLastMessage() string`

GetLastMessage returns the LastMessage field if non-nil, zero value otherwise.

### GetLastMessageOk

`func (o *GetInboxConversation200ResponseData) GetLastMessageOk() (*string, bool)`

GetLastMessageOk returns a tuple with the LastMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessage

`func (o *GetInboxConversation200ResponseData) SetLastMessage(v string)`

SetLastMessage sets LastMessage field to given value.

### HasLastMessage

`func (o *GetInboxConversation200ResponseData) HasLastMessage() bool`

HasLastMessage returns a boolean if a field has been set.

### GetLastMessageAt

`func (o *GetInboxConversation200ResponseData) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *GetInboxConversation200ResponseData) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *GetInboxConversation200ResponseData) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *GetInboxConversation200ResponseData) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *GetInboxConversation200ResponseData) GetUpdatedTime() time.Time`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *GetInboxConversation200ResponseData) GetUpdatedTimeOk() (*time.Time, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *GetInboxConversation200ResponseData) SetUpdatedTime(v time.Time)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *GetInboxConversation200ResponseData) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetParticipants

`func (o *GetInboxConversation200ResponseData) GetParticipants() []UpdateFacebookPage200ResponseSelectedPage`

GetParticipants returns the Participants field if non-nil, zero value otherwise.

### GetParticipantsOk

`func (o *GetInboxConversation200ResponseData) GetParticipantsOk() (*[]UpdateFacebookPage200ResponseSelectedPage, bool)`

GetParticipantsOk returns a tuple with the Participants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipants

`func (o *GetInboxConversation200ResponseData) SetParticipants(v []UpdateFacebookPage200ResponseSelectedPage)`

SetParticipants sets Participants field to given value.

### HasParticipants

`func (o *GetInboxConversation200ResponseData) HasParticipants() bool`

HasParticipants returns a boolean if a field has been set.

### GetInstagramProfile

`func (o *GetInboxConversation200ResponseData) GetInstagramProfile() ListInboxConversations200ResponseDataInnerInstagramProfile`

GetInstagramProfile returns the InstagramProfile field if non-nil, zero value otherwise.

### GetInstagramProfileOk

`func (o *GetInboxConversation200ResponseData) GetInstagramProfileOk() (*ListInboxConversations200ResponseDataInnerInstagramProfile, bool)`

GetInstagramProfileOk returns a tuple with the InstagramProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramProfile

`func (o *GetInboxConversation200ResponseData) SetInstagramProfile(v ListInboxConversations200ResponseDataInnerInstagramProfile)`

SetInstagramProfile sets InstagramProfile field to given value.

### HasInstagramProfile

`func (o *GetInboxConversation200ResponseData) HasInstagramProfile() bool`

HasInstagramProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


