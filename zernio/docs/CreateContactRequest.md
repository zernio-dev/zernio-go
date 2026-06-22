# CreateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**Name** | **string** |  | 
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] [default to true]
**Notes** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** | Optional. Creates a channel if provided with platform + platformIdentifier | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**PlatformIdentifier** | Pointer to **string** |  | [optional] 
**DisplayIdentifier** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateContactRequest

`func NewCreateContactRequest(profileId string, name string, ) *CreateContactRequest`

NewCreateContactRequest instantiates a new CreateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContactRequestWithDefaults

`func NewCreateContactRequestWithDefaults() *CreateContactRequest`

NewCreateContactRequestWithDefaults instantiates a new CreateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateContactRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateContactRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateContactRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetName

`func (o *CreateContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContactRequest) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *CreateContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *CreateContactRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreateContactRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreateContactRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreateContactRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTags

`func (o *CreateContactRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContactRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContactRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *CreateContactRequest) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *CreateContactRequest) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *CreateContactRequest) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *CreateContactRequest) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetNotes

`func (o *CreateContactRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateContactRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateContactRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateContactRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetAccountId

`func (o *CreateContactRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateContactRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateContactRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *CreateContactRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *CreateContactRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateContactRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateContactRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateContactRequest) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformIdentifier

`func (o *CreateContactRequest) GetPlatformIdentifier() string`

GetPlatformIdentifier returns the PlatformIdentifier field if non-nil, zero value otherwise.

### GetPlatformIdentifierOk

`func (o *CreateContactRequest) GetPlatformIdentifierOk() (*string, bool)`

GetPlatformIdentifierOk returns a tuple with the PlatformIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentifier

`func (o *CreateContactRequest) SetPlatformIdentifier(v string)`

SetPlatformIdentifier sets PlatformIdentifier field to given value.

### HasPlatformIdentifier

`func (o *CreateContactRequest) HasPlatformIdentifier() bool`

HasPlatformIdentifier returns a boolean if a field has been set.

### GetDisplayIdentifier

`func (o *CreateContactRequest) GetDisplayIdentifier() string`

GetDisplayIdentifier returns the DisplayIdentifier field if non-nil, zero value otherwise.

### GetDisplayIdentifierOk

`func (o *CreateContactRequest) GetDisplayIdentifierOk() (*string, bool)`

GetDisplayIdentifierOk returns a tuple with the DisplayIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayIdentifier

`func (o *CreateContactRequest) SetDisplayIdentifier(v string)`

SetDisplayIdentifier sets DisplayIdentifier field to given value.

### HasDisplayIdentifier

`func (o *CreateContactRequest) HasDisplayIdentifier() bool`

HasDisplayIdentifier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


