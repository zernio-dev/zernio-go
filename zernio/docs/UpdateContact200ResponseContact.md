# UpdateContact200ResponseContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**AvatarUrl** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**IsBlocked** | Pointer to **bool** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateContact200ResponseContact

`func NewUpdateContact200ResponseContact() *UpdateContact200ResponseContact`

NewUpdateContact200ResponseContact instantiates a new UpdateContact200ResponseContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContact200ResponseContactWithDefaults

`func NewUpdateContact200ResponseContactWithDefaults() *UpdateContact200ResponseContact`

NewUpdateContact200ResponseContactWithDefaults instantiates a new UpdateContact200ResponseContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateContact200ResponseContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateContact200ResponseContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateContact200ResponseContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateContact200ResponseContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateContact200ResponseContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateContact200ResponseContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateContact200ResponseContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateContact200ResponseContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateContact200ResponseContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateContact200ResponseContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateContact200ResponseContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateContact200ResponseContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *UpdateContact200ResponseContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdateContact200ResponseContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdateContact200ResponseContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *UpdateContact200ResponseContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UpdateContact200ResponseContact) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UpdateContact200ResponseContact) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UpdateContact200ResponseContact) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UpdateContact200ResponseContact) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetTags

`func (o *UpdateContact200ResponseContact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateContact200ResponseContact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateContact200ResponseContact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateContact200ResponseContact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *UpdateContact200ResponseContact) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *UpdateContact200ResponseContact) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *UpdateContact200ResponseContact) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *UpdateContact200ResponseContact) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsBlocked

`func (o *UpdateContact200ResponseContact) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *UpdateContact200ResponseContact) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *UpdateContact200ResponseContact) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *UpdateContact200ResponseContact) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetNotes

`func (o *UpdateContact200ResponseContact) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateContact200ResponseContact) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateContact200ResponseContact) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateContact200ResponseContact) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UpdateContact200ResponseContact) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateContact200ResponseContact) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateContact200ResponseContact) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateContact200ResponseContact) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


