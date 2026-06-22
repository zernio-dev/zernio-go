# SelectFacebookPageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Profile ID from your connection flow | 
**PageId** | **string** | The Facebook Page ID selected by the user | 
**TempToken** | **string** | Temporary Facebook access token from OAuth | 
**UserProfile** | [**SelectFacebookPageRequestUserProfile**](SelectFacebookPageRequestUserProfile.md) |  | 
**RedirectUrl** | Pointer to **string** | Optional custom redirect URL to return to after selection | [optional] 

## Methods

### NewSelectFacebookPageRequest

`func NewSelectFacebookPageRequest(profileId string, pageId string, tempToken string, userProfile SelectFacebookPageRequestUserProfile, ) *SelectFacebookPageRequest`

NewSelectFacebookPageRequest instantiates a new SelectFacebookPageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectFacebookPageRequestWithDefaults

`func NewSelectFacebookPageRequestWithDefaults() *SelectFacebookPageRequest`

NewSelectFacebookPageRequestWithDefaults instantiates a new SelectFacebookPageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SelectFacebookPageRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SelectFacebookPageRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SelectFacebookPageRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetPageId

`func (o *SelectFacebookPageRequest) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *SelectFacebookPageRequest) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *SelectFacebookPageRequest) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetTempToken

`func (o *SelectFacebookPageRequest) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *SelectFacebookPageRequest) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *SelectFacebookPageRequest) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.


### GetUserProfile

`func (o *SelectFacebookPageRequest) GetUserProfile() SelectFacebookPageRequestUserProfile`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *SelectFacebookPageRequest) GetUserProfileOk() (*SelectFacebookPageRequestUserProfile, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *SelectFacebookPageRequest) SetUserProfile(v SelectFacebookPageRequestUserProfile)`

SetUserProfile sets UserProfile field to given value.


### GetRedirectUrl

`func (o *SelectFacebookPageRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectFacebookPageRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectFacebookPageRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectFacebookPageRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


