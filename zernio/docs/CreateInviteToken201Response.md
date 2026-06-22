# CreateInviteToken201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**InvitedProfileIds** | Pointer to **[]string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**InviteUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateInviteToken201Response

`func NewCreateInviteToken201Response() *CreateInviteToken201Response`

NewCreateInviteToken201Response instantiates a new CreateInviteToken201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInviteToken201ResponseWithDefaults

`func NewCreateInviteToken201ResponseWithDefaults() *CreateInviteToken201Response`

NewCreateInviteToken201ResponseWithDefaults instantiates a new CreateInviteToken201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *CreateInviteToken201Response) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateInviteToken201Response) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateInviteToken201Response) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateInviteToken201Response) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetScope

`func (o *CreateInviteToken201Response) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateInviteToken201Response) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateInviteToken201Response) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateInviteToken201Response) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetInvitedProfileIds

`func (o *CreateInviteToken201Response) GetInvitedProfileIds() []string`

GetInvitedProfileIds returns the InvitedProfileIds field if non-nil, zero value otherwise.

### GetInvitedProfileIdsOk

`func (o *CreateInviteToken201Response) GetInvitedProfileIdsOk() (*[]string, bool)`

GetInvitedProfileIdsOk returns a tuple with the InvitedProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitedProfileIds

`func (o *CreateInviteToken201Response) SetInvitedProfileIds(v []string)`

SetInvitedProfileIds sets InvitedProfileIds field to given value.

### HasInvitedProfileIds

`func (o *CreateInviteToken201Response) HasInvitedProfileIds() bool`

HasInvitedProfileIds returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CreateInviteToken201Response) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CreateInviteToken201Response) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CreateInviteToken201Response) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CreateInviteToken201Response) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetInviteUrl

`func (o *CreateInviteToken201Response) GetInviteUrl() string`

GetInviteUrl returns the InviteUrl field if non-nil, zero value otherwise.

### GetInviteUrlOk

`func (o *CreateInviteToken201Response) GetInviteUrlOk() (*string, bool)`

GetInviteUrlOk returns a tuple with the InviteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteUrl

`func (o *CreateInviteToken201Response) SetInviteUrl(v string)`

SetInviteUrl sets InviteUrl field to given value.

### HasInviteUrl

`func (o *CreateInviteToken201Response) HasInviteUrl() bool`

HasInviteUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


