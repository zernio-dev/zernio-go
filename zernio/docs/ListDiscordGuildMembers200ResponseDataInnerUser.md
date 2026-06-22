# ListDiscordGuildMembers200ResponseDataInnerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | User snowflake | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Discriminator** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **NullableString** |  | [optional] 
**GlobalName** | Pointer to **NullableString** | User&#39;s display name (post-2023 Discord rebrand) | [optional] 

## Methods

### NewListDiscordGuildMembers200ResponseDataInnerUser

`func NewListDiscordGuildMembers200ResponseDataInnerUser() *ListDiscordGuildMembers200ResponseDataInnerUser`

NewListDiscordGuildMembers200ResponseDataInnerUser instantiates a new ListDiscordGuildMembers200ResponseDataInnerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscordGuildMembers200ResponseDataInnerUserWithDefaults

`func NewListDiscordGuildMembers200ResponseDataInnerUserWithDefaults() *ListDiscordGuildMembers200ResponseDataInnerUser`

NewListDiscordGuildMembers200ResponseDataInnerUserWithDefaults instantiates a new ListDiscordGuildMembers200ResponseDataInnerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDiscriminator

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetDiscriminator() string`

GetDiscriminator returns the Discriminator field if non-nil, zero value otherwise.

### GetDiscriminatorOk

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetDiscriminatorOk() (*string, bool)`

GetDiscriminatorOk returns a tuple with the Discriminator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscriminator

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetDiscriminator(v string)`

SetDiscriminator sets Discriminator field to given value.

### HasDiscriminator

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) HasDiscriminator() bool`

HasDiscriminator returns a boolean if a field has been set.

### GetAvatar

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### SetAvatarNil

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetAvatarNil(b bool)`

 SetAvatarNil sets the value for Avatar to be an explicit nil

### UnsetAvatar
`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) UnsetAvatar()`

UnsetAvatar ensures that no value is present for Avatar, not even an explicit nil
### GetGlobalName

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetGlobalName() string`

GetGlobalName returns the GlobalName field if non-nil, zero value otherwise.

### GetGlobalNameOk

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) GetGlobalNameOk() (*string, bool)`

GetGlobalNameOk returns a tuple with the GlobalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalName

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetGlobalName(v string)`

SetGlobalName sets GlobalName field to given value.

### HasGlobalName

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) HasGlobalName() bool`

HasGlobalName returns a boolean if a field has been set.

### SetGlobalNameNil

`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) SetGlobalNameNil(b bool)`

 SetGlobalNameNil sets the value for GlobalName to be an explicit nil

### UnsetGlobalName
`func (o *ListDiscordGuildMembers200ResponseDataInnerUser) UnsetGlobalName()`

UnsetGlobalName ensures that no value is present for GlobalName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


