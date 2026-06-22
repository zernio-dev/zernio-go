# UpdateContactRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**IsBlocked** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateContactRequest

`func NewUpdateContactRequest() *UpdateContactRequest`

NewUpdateContactRequest instantiates a new UpdateContactRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContactRequestWithDefaults

`func NewUpdateContactRequestWithDefaults() *UpdateContactRequest`

NewUpdateContactRequestWithDefaults instantiates a new UpdateContactRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateContactRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateContactRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateContactRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateContactRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateContactRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateContactRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateContactRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateContactRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *UpdateContactRequest) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateContactRequest) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateContactRequest) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdateContactRequest) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UpdateContactRequest) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UpdateContactRequest) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UpdateContactRequest) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UpdateContactRequest) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetTags

`func (o *UpdateContactRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateContactRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateContactRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateContactRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *UpdateContactRequest) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *UpdateContactRequest) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *UpdateContactRequest) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *UpdateContactRequest) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsBlocked

`func (o *UpdateContactRequest) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *UpdateContactRequest) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *UpdateContactRequest) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *UpdateContactRequest) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateContactRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateContactRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateContactRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateContactRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


