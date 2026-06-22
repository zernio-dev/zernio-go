# GetPendingOAuthData200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** | The platform (e.g., \&quot;linkedin\&quot;) | [optional] 
**ProfileId** | Pointer to **string** | The Zernio profile ID | [optional] 
**TempToken** | Pointer to **string** | Temporary access token for the platform | [optional] 
**RefreshToken** | Pointer to **string** | Refresh token (if available) | [optional] 
**ExpiresIn** | Pointer to **float32** | Token expiry in seconds | [optional] 
**UserProfile** | Pointer to **map[string]interface{}** | User profile data (id, username, displayName, profilePicture) | [optional] 
**SelectionType** | Pointer to **string** | Type of selection data | [optional] 
**Organizations** | Pointer to [**[]GetPendingOAuthData200ResponseOrganizationsInner**](GetPendingOAuthData200ResponseOrganizationsInner.md) | LinkedIn organizations (when selectionType is \&quot;organizations\&quot;) | [optional] 

## Methods

### NewGetPendingOAuthData200Response

`func NewGetPendingOAuthData200Response() *GetPendingOAuthData200Response`

NewGetPendingOAuthData200Response instantiates a new GetPendingOAuthData200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPendingOAuthData200ResponseWithDefaults

`func NewGetPendingOAuthData200ResponseWithDefaults() *GetPendingOAuthData200Response`

NewGetPendingOAuthData200ResponseWithDefaults instantiates a new GetPendingOAuthData200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetPendingOAuthData200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetPendingOAuthData200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetPendingOAuthData200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetPendingOAuthData200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetProfileId

`func (o *GetPendingOAuthData200Response) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetPendingOAuthData200Response) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetPendingOAuthData200Response) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetPendingOAuthData200Response) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetTempToken

`func (o *GetPendingOAuthData200Response) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *GetPendingOAuthData200Response) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *GetPendingOAuthData200Response) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.

### HasTempToken

`func (o *GetPendingOAuthData200Response) HasTempToken() bool`

HasTempToken returns a boolean if a field has been set.

### GetRefreshToken

`func (o *GetPendingOAuthData200Response) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *GetPendingOAuthData200Response) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *GetPendingOAuthData200Response) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *GetPendingOAuthData200Response) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetPendingOAuthData200Response) GetExpiresIn() float32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetPendingOAuthData200Response) GetExpiresInOk() (*float32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetPendingOAuthData200Response) SetExpiresIn(v float32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetPendingOAuthData200Response) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetUserProfile

`func (o *GetPendingOAuthData200Response) GetUserProfile() map[string]interface{}`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *GetPendingOAuthData200Response) GetUserProfileOk() (*map[string]interface{}, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *GetPendingOAuthData200Response) SetUserProfile(v map[string]interface{})`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *GetPendingOAuthData200Response) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.

### GetSelectionType

`func (o *GetPendingOAuthData200Response) GetSelectionType() string`

GetSelectionType returns the SelectionType field if non-nil, zero value otherwise.

### GetSelectionTypeOk

`func (o *GetPendingOAuthData200Response) GetSelectionTypeOk() (*string, bool)`

GetSelectionTypeOk returns a tuple with the SelectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionType

`func (o *GetPendingOAuthData200Response) SetSelectionType(v string)`

SetSelectionType sets SelectionType field to given value.

### HasSelectionType

`func (o *GetPendingOAuthData200Response) HasSelectionType() bool`

HasSelectionType returns a boolean if a field has been set.

### GetOrganizations

`func (o *GetPendingOAuthData200Response) GetOrganizations() []GetPendingOAuthData200ResponseOrganizationsInner`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetPendingOAuthData200Response) GetOrganizationsOk() (*[]GetPendingOAuthData200ResponseOrganizationsInner, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetPendingOAuthData200Response) SetOrganizations(v []GetPendingOAuthData200ResponseOrganizationsInner)`

SetOrganizations sets Organizations field to given value.

### HasOrganizations

`func (o *GetPendingOAuthData200Response) HasOrganizations() bool`

HasOrganizations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


