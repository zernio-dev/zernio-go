# Profile

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

### NewProfile

`func NewProfile() *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Profile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Profile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Profile) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Profile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *Profile) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *Profile) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *Profile) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *Profile) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetName

`func (o *Profile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Profile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Profile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Profile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Profile) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Profile) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Profile) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Profile) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetColor

`func (o *Profile) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *Profile) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *Profile) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *Profile) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetIsDefault

`func (o *Profile) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Profile) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Profile) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Profile) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsOverLimit

`func (o *Profile) GetIsOverLimit() bool`

GetIsOverLimit returns the IsOverLimit field if non-nil, zero value otherwise.

### GetIsOverLimitOk

`func (o *Profile) GetIsOverLimitOk() (*bool, bool)`

GetIsOverLimitOk returns a tuple with the IsOverLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverLimit

`func (o *Profile) SetIsOverLimit(v bool)`

SetIsOverLimit sets IsOverLimit field to given value.

### HasIsOverLimit

`func (o *Profile) HasIsOverLimit() bool`

HasIsOverLimit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Profile) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Profile) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Profile) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Profile) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


