# SelectSnapchatProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Your Zernio profile ID | 
**SelectedPublicProfile** | [**SelectSnapchatProfileRequestSelectedPublicProfile**](SelectSnapchatProfileRequestSelectedPublicProfile.md) |  | 
**TempToken** | **string** | Temporary Snapchat access token from OAuth | 
**UserProfile** | **map[string]interface{}** | User profile data from OAuth redirect | 
**RefreshToken** | Pointer to **string** | Snapchat refresh token (if available) | [optional] 
**ExpiresIn** | Pointer to **int32** | Token expiration time in seconds | [optional] 
**RedirectUrl** | Pointer to **string** | Custom redirect URL after connection completes | [optional] 

## Methods

### NewSelectSnapchatProfileRequest

`func NewSelectSnapchatProfileRequest(profileId string, selectedPublicProfile SelectSnapchatProfileRequestSelectedPublicProfile, tempToken string, userProfile map[string]interface{}, ) *SelectSnapchatProfileRequest`

NewSelectSnapchatProfileRequest instantiates a new SelectSnapchatProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectSnapchatProfileRequestWithDefaults

`func NewSelectSnapchatProfileRequestWithDefaults() *SelectSnapchatProfileRequest`

NewSelectSnapchatProfileRequestWithDefaults instantiates a new SelectSnapchatProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SelectSnapchatProfileRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SelectSnapchatProfileRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SelectSnapchatProfileRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetSelectedPublicProfile

`func (o *SelectSnapchatProfileRequest) GetSelectedPublicProfile() SelectSnapchatProfileRequestSelectedPublicProfile`

GetSelectedPublicProfile returns the SelectedPublicProfile field if non-nil, zero value otherwise.

### GetSelectedPublicProfileOk

`func (o *SelectSnapchatProfileRequest) GetSelectedPublicProfileOk() (*SelectSnapchatProfileRequestSelectedPublicProfile, bool)`

GetSelectedPublicProfileOk returns a tuple with the SelectedPublicProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedPublicProfile

`func (o *SelectSnapchatProfileRequest) SetSelectedPublicProfile(v SelectSnapchatProfileRequestSelectedPublicProfile)`

SetSelectedPublicProfile sets SelectedPublicProfile field to given value.


### GetTempToken

`func (o *SelectSnapchatProfileRequest) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *SelectSnapchatProfileRequest) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *SelectSnapchatProfileRequest) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.


### GetUserProfile

`func (o *SelectSnapchatProfileRequest) GetUserProfile() map[string]interface{}`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *SelectSnapchatProfileRequest) GetUserProfileOk() (*map[string]interface{}, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *SelectSnapchatProfileRequest) SetUserProfile(v map[string]interface{})`

SetUserProfile sets UserProfile field to given value.


### GetRefreshToken

`func (o *SelectSnapchatProfileRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *SelectSnapchatProfileRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *SelectSnapchatProfileRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *SelectSnapchatProfileRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *SelectSnapchatProfileRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *SelectSnapchatProfileRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *SelectSnapchatProfileRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *SelectSnapchatProfileRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectSnapchatProfileRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectSnapchatProfileRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectSnapchatProfileRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectSnapchatProfileRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


