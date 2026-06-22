# ListDiscordGuildRoles200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Role snowflake ID | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **int32** | Decimal color (0 &#x3D; no color). Convert to hex via .toString(16). | [optional] 
**Position** | Pointer to **int32** | Position in role hierarchy (higher &#x3D; more authority) | [optional] 
**Permissions** | Pointer to **string** | Permissions bitfield as a stringified integer | [optional] 
**Managed** | Pointer to **bool** | True for integration-managed roles (bot roles) | [optional] 
**Mentionable** | Pointer to **bool** |  | [optional] 
**Hoist** | Pointer to **bool** | True if role is displayed separately in member list | [optional] 

## Methods

### NewListDiscordGuildRoles200ResponseDataInner

`func NewListDiscordGuildRoles200ResponseDataInner() *ListDiscordGuildRoles200ResponseDataInner`

NewListDiscordGuildRoles200ResponseDataInner instantiates a new ListDiscordGuildRoles200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDiscordGuildRoles200ResponseDataInnerWithDefaults

`func NewListDiscordGuildRoles200ResponseDataInnerWithDefaults() *ListDiscordGuildRoles200ResponseDataInner`

NewListDiscordGuildRoles200ResponseDataInnerWithDefaults instantiates a new ListDiscordGuildRoles200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetColor

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetColor() int32`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetColorOk() (*int32, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetColor(v int32)`

SetColor sets Color field to given value.

### HasColor

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetPosition

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPermissions

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetManaged

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasManaged() bool`

HasManaged returns a boolean if a field has been set.

### GetMentionable

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetMentionable() bool`

GetMentionable returns the Mentionable field if non-nil, zero value otherwise.

### GetMentionableOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetMentionableOk() (*bool, bool)`

GetMentionableOk returns a tuple with the Mentionable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentionable

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetMentionable(v bool)`

SetMentionable sets Mentionable field to given value.

### HasMentionable

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasMentionable() bool`

HasMentionable returns a boolean if a field has been set.

### GetHoist

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetHoist() bool`

GetHoist returns the Hoist field if non-nil, zero value otherwise.

### GetHoistOk

`func (o *ListDiscordGuildRoles200ResponseDataInner) GetHoistOk() (*bool, bool)`

GetHoistOk returns a tuple with the Hoist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoist

`func (o *ListDiscordGuildRoles200ResponseDataInner) SetHoist(v bool)`

SetHoist sets Hoist field to given value.

### HasHoist

`func (o *ListDiscordGuildRoles200ResponseDataInner) HasHoist() bool`

HasHoist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


