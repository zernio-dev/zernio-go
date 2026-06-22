# ListDiscordGuildMembers200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**ListDiscordGuildMembers200ResponseDataInnerUser**](ListDiscordGuildMembers200ResponseDataInnerUser.md) |  | [optional] 
**Nick** | Pointer to **NullableString** | Guild-specific nickname | [optional] 
**Roles** | Pointer to **[]string** | Snowflake IDs of roles assigned to this member | [optional] 
**JoinedAt** | Pointer to **time.Time** |  | [optional] 
**PremiumSince** | Pointer to **NullableTime** | When the user started boosting the server | [optional] 

## Methods

### NewListDiscordGuildMembers200ResponseDataInner

`func NewListDiscordGuildMembers200ResponseDataInner() *ListDiscordGuildMembers200ResponseDataInner`

NewListDiscordGuildMembers200ResponseDataInner instantiates a new ListDiscordGuildMembers200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscordGuildMembers200ResponseDataInnerWithDefaults

`func NewListDiscordGuildMembers200ResponseDataInnerWithDefaults() *ListDiscordGuildMembers200ResponseDataInner`

NewListDiscordGuildMembers200ResponseDataInnerWithDefaults instantiates a new ListDiscordGuildMembers200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetUser() ListDiscordGuildMembers200ResponseDataInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetUserOk() (*ListDiscordGuildMembers200ResponseDataInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetUser(v ListDiscordGuildMembers200ResponseDataInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ListDiscordGuildMembers200ResponseDataInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetNick

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetNick() string`

GetNick returns the Nick field if non-nil, zero value otherwise.

### GetNickOk

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetNickOk() (*string, bool)`

GetNickOk returns a tuple with the Nick field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNick

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetNick(v string)`

SetNick sets Nick field to given value.

### HasNick

`func (o *ListDiscordGuildMembers200ResponseDataInner) HasNick() bool`

HasNick returns a boolean if a field has been set.

### SetNickNil

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetNickNil(b bool)`

 SetNickNil sets the value for Nick to be an explicit nil

### UnsetNick
`func (o *ListDiscordGuildMembers200ResponseDataInner) UnsetNick()`

UnsetNick ensures that no value is present for Nick, not even an explicit nil
### GetRoles

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ListDiscordGuildMembers200ResponseDataInner) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetJoinedAt

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *ListDiscordGuildMembers200ResponseDataInner) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetPremiumSince

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetPremiumSince() time.Time`

GetPremiumSince returns the PremiumSince field if non-nil, zero value otherwise.

### GetPremiumSinceOk

`func (o *ListDiscordGuildMembers200ResponseDataInner) GetPremiumSinceOk() (*time.Time, bool)`

GetPremiumSinceOk returns a tuple with the PremiumSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremiumSince

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetPremiumSince(v time.Time)`

SetPremiumSince sets PremiumSince field to given value.

### HasPremiumSince

`func (o *ListDiscordGuildMembers200ResponseDataInner) HasPremiumSince() bool`

HasPremiumSince returns a boolean if a field has been set.

### SetPremiumSinceNil

`func (o *ListDiscordGuildMembers200ResponseDataInner) SetPremiumSinceNil(b bool)`

 SetPremiumSinceNil sets the value for PremiumSince to be an explicit nil

### UnsetPremiumSince
`func (o *ListDiscordGuildMembers200ResponseDataInner) UnsetPremiumSince()`

UnsetPremiumSince ensures that no value is present for PremiumSince, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


