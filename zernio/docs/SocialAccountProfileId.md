# SocialAccountProfileId

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsOverLimit** | Pointer to **bool** | Only present when includeOverLimit&#x3D;true. Indicates if this profile exceeds the plan limit. | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSocialAccountProfileId

`func NewSocialAccountProfileId() *SocialAccountProfileId`

NewSocialAccountProfileId instantiates a new SocialAccountProfileId object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialAccountProfileIdWithDefaults

`func NewSocialAccountProfileIdWithDefaults() *SocialAccountProfileId`

NewSocialAccountProfileIdWithDefaults instantiates a new SocialAccountProfileId object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SocialAccountProfileId) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialAccountProfileId) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialAccountProfileId) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialAccountProfileId) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *SocialAccountProfileId) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SocialAccountProfileId) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SocialAccountProfileId) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SocialAccountProfileId) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *SocialAccountProfileId) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SocialAccountProfileId) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SocialAccountProfileId) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SocialAccountProfileId) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SocialAccountProfileId) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SocialAccountProfileId) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SocialAccountProfileId) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SocialAccountProfileId) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *SocialAccountProfileId) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *SocialAccountProfileId) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *SocialAccountProfileId) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *SocialAccountProfileId) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetIsDefault

`func (o *SocialAccountProfileId) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SocialAccountProfileId) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SocialAccountProfileId) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SocialAccountProfileId) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsOverLimit

`func (o *SocialAccountProfileId) GetIsOverLimit() bool`

GetIsOverLimit returns the IsOverLimit field if non-nil, zero value otherwise.

### GetIsOverLimitOk

`func (o *SocialAccountProfileId) GetIsOverLimitOk() (*bool, bool)`

GetIsOverLimitOk returns a tuple with the IsOverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverLimit

`func (o *SocialAccountProfileId) SetIsOverLimit(v bool)`

SetIsOverLimit sets IsOverLimit field to given value.

### HasIsOverLimit

`func (o *SocialAccountProfileId) HasIsOverLimit() bool`

HasIsOverLimit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SocialAccountProfileId) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SocialAccountProfileId) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SocialAccountProfileId) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SocialAccountProfileId) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


