# CreateContact200ResponseContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**IsSubscribed** | Pointer to **bool** |  | [optional] 
**IsBlocked** | Pointer to **bool** |  | [optional] 
**CustomFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateContact200ResponseContact

`func NewCreateContact200ResponseContact() *CreateContact200ResponseContact`

NewCreateContact200ResponseContact instantiates a new CreateContact200ResponseContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContact200ResponseContactWithDefaults

`func NewCreateContact200ResponseContactWithDefaults() *CreateContact200ResponseContact`

NewCreateContact200ResponseContactWithDefaults instantiates a new CreateContact200ResponseContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateContact200ResponseContact) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateContact200ResponseContact) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateContact200ResponseContact) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateContact200ResponseContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateContact200ResponseContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateContact200ResponseContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateContact200ResponseContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateContact200ResponseContact) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *CreateContact200ResponseContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateContact200ResponseContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateContact200ResponseContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateContact200ResponseContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetCompany

`func (o *CreateContact200ResponseContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *CreateContact200ResponseContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *CreateContact200ResponseContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *CreateContact200ResponseContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetTags

`func (o *CreateContact200ResponseContact) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateContact200ResponseContact) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateContact200ResponseContact) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateContact200ResponseContact) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetIsSubscribed

`func (o *CreateContact200ResponseContact) GetIsSubscribed() bool`

GetIsSubscribed returns the IsSubscribed field if non-nil, zero value otherwise.

### GetIsSubscribedOk

`func (o *CreateContact200ResponseContact) GetIsSubscribedOk() (*bool, bool)`

GetIsSubscribedOk returns a tuple with the IsSubscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSubscribed

`func (o *CreateContact200ResponseContact) SetIsSubscribed(v bool)`

SetIsSubscribed sets IsSubscribed field to given value.

### HasIsSubscribed

`func (o *CreateContact200ResponseContact) HasIsSubscribed() bool`

HasIsSubscribed returns a boolean if a field has been set.

### GetIsBlocked

`func (o *CreateContact200ResponseContact) GetIsBlocked() bool`

GetIsBlocked returns the IsBlocked field if non-nil, zero value otherwise.

### GetIsBlockedOk

`func (o *CreateContact200ResponseContact) GetIsBlockedOk() (*bool, bool)`

GetIsBlockedOk returns a tuple with the IsBlocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBlocked

`func (o *CreateContact200ResponseContact) SetIsBlocked(v bool)`

SetIsBlocked sets IsBlocked field to given value.

### HasIsBlocked

`func (o *CreateContact200ResponseContact) HasIsBlocked() bool`

HasIsBlocked returns a boolean if a field has been set.

### GetCustomFields

`func (o *CreateContact200ResponseContact) GetCustomFields() map[string]interface{}`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *CreateContact200ResponseContact) GetCustomFieldsOk() (*map[string]interface{}, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *CreateContact200ResponseContact) SetCustomFields(v map[string]interface{})`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *CreateContact200ResponseContact) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.

### GetNotes

`func (o *CreateContact200ResponseContact) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateContact200ResponseContact) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateContact200ResponseContact) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateContact200ResponseContact) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateContact200ResponseContact) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateContact200ResponseContact) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateContact200ResponseContact) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateContact200ResponseContact) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


