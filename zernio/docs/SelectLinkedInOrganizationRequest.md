# SelectLinkedInOrganizationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**TempToken** | **string** |  | 
**UserProfile** | **map[string]interface{}** |  | 
**AccountType** | **string** |  | 
**SelectedOrganization** | Pointer to **map[string]interface{}** |  | [optional] 
**RedirectUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewSelectLinkedInOrganizationRequest

`func NewSelectLinkedInOrganizationRequest(profileId string, tempToken string, userProfile map[string]interface{}, accountType string, ) *SelectLinkedInOrganizationRequest`

NewSelectLinkedInOrganizationRequest instantiates a new SelectLinkedInOrganizationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectLinkedInOrganizationRequestWithDefaults

`func NewSelectLinkedInOrganizationRequestWithDefaults() *SelectLinkedInOrganizationRequest`

NewSelectLinkedInOrganizationRequestWithDefaults instantiates a new SelectLinkedInOrganizationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SelectLinkedInOrganizationRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SelectLinkedInOrganizationRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SelectLinkedInOrganizationRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetTempToken

`func (o *SelectLinkedInOrganizationRequest) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *SelectLinkedInOrganizationRequest) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *SelectLinkedInOrganizationRequest) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.


### GetUserProfile

`func (o *SelectLinkedInOrganizationRequest) GetUserProfile() map[string]interface{}`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *SelectLinkedInOrganizationRequest) GetUserProfileOk() (*map[string]interface{}, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *SelectLinkedInOrganizationRequest) SetUserProfile(v map[string]interface{})`

SetUserProfile sets UserProfile field to given value.


### GetAccountType

`func (o *SelectLinkedInOrganizationRequest) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *SelectLinkedInOrganizationRequest) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *SelectLinkedInOrganizationRequest) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.


### GetSelectedOrganization

`func (o *SelectLinkedInOrganizationRequest) GetSelectedOrganization() map[string]interface{}`

GetSelectedOrganization returns the SelectedOrganization field if non-nil, zero value otherwise.

### GetSelectedOrganizationOk

`func (o *SelectLinkedInOrganizationRequest) GetSelectedOrganizationOk() (*map[string]interface{}, bool)`

GetSelectedOrganizationOk returns a tuple with the SelectedOrganization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedOrganization

`func (o *SelectLinkedInOrganizationRequest) SetSelectedOrganization(v map[string]interface{})`

SetSelectedOrganization sets SelectedOrganization field to given value.

### HasSelectedOrganization

`func (o *SelectLinkedInOrganizationRequest) HasSelectedOrganization() bool`

HasSelectedOrganization returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectLinkedInOrganizationRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectLinkedInOrganizationRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectLinkedInOrganizationRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectLinkedInOrganizationRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


