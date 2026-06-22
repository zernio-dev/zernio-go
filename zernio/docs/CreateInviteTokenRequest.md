# CreateInviteTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | &#39;all&#39; grants access to all profiles, &#39;profiles&#39; restricts to specific profiles | 
**ProfileIds** | Pointer to **[]string** | Required if scope is &#39;profiles&#39;. Array of profile IDs to grant access to. | [optional] 
**Role** | Pointer to **string** | Org role granted to the invitee. Defaults to &#39;member&#39;. &#39;viewer&#39; creates a read-only member who can view everything in their profile scope but cannot perform any content mutation (publish, edit, delete, connect accounts). | [optional] [default to "member"]
**ReadOnly** | Pointer to **bool** | Deprecated. Use role &#39;viewer&#39; instead. When true, the invite is created with role &#39;viewer&#39;. Cannot be combined with role &#39;billing_admin&#39;. | [optional] 

## Methods

### NewCreateInviteTokenRequest

`func NewCreateInviteTokenRequest(scope string, ) *CreateInviteTokenRequest`

NewCreateInviteTokenRequest instantiates a new CreateInviteTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInviteTokenRequestWithDefaults

`func NewCreateInviteTokenRequestWithDefaults() *CreateInviteTokenRequest`

NewCreateInviteTokenRequestWithDefaults instantiates a new CreateInviteTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CreateInviteTokenRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateInviteTokenRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateInviteTokenRequest) SetScope(v string)`

SetScope sets Scope field to given value.


### GetProfileIds

`func (o *CreateInviteTokenRequest) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *CreateInviteTokenRequest) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *CreateInviteTokenRequest) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *CreateInviteTokenRequest) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetRole

`func (o *CreateInviteTokenRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *CreateInviteTokenRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *CreateInviteTokenRequest) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *CreateInviteTokenRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetReadOnly

`func (o *CreateInviteTokenRequest) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *CreateInviteTokenRequest) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *CreateInviteTokenRequest) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *CreateInviteTokenRequest) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


