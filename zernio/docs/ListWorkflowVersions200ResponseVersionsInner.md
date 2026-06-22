# ListWorkflowVersions200ResponseVersionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **int32** | Monotonically increasing version number | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**CreatedBy** | Pointer to **NullableString** | User id that authored this version | [optional] 
**CreatedByEmail** | Pointer to **NullableString** | Denormalized email so the history UI can render without a join | [optional] 
**RestoredFromVersion** | Pointer to **NullableInt32** | When non-null, this snapshot was created by restoring that version | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListWorkflowVersions200ResponseVersionsInner

`func NewListWorkflowVersions200ResponseVersionsInner() *ListWorkflowVersions200ResponseVersionsInner`

NewListWorkflowVersions200ResponseVersionsInner instantiates a new ListWorkflowVersions200ResponseVersionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowVersions200ResponseVersionsInnerWithDefaults

`func NewListWorkflowVersions200ResponseVersionsInnerWithDefaults() *ListWorkflowVersions200ResponseVersionsInner`

NewListWorkflowVersions200ResponseVersionsInnerWithDefaults instantiates a new ListWorkflowVersions200ResponseVersionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListWorkflowVersions200ResponseVersionsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatedBy

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### SetCreatedByNil

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetCreatedByNil(b bool)`

 SetCreatedByNil sets the value for CreatedBy to be an explicit nil

### UnsetCreatedBy
`func (o *ListWorkflowVersions200ResponseVersionsInner) UnsetCreatedBy()`

UnsetCreatedBy ensures that no value is present for CreatedBy, not even an explicit nil
### GetCreatedByEmail

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetCreatedByEmail() string`

GetCreatedByEmail returns the CreatedByEmail field if non-nil, zero value otherwise.

### GetCreatedByEmailOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetCreatedByEmailOk() (*string, bool)`

GetCreatedByEmailOk returns a tuple with the CreatedByEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByEmail

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetCreatedByEmail(v string)`

SetCreatedByEmail sets CreatedByEmail field to given value.

### HasCreatedByEmail

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasCreatedByEmail() bool`

HasCreatedByEmail returns a boolean if a field has been set.

### SetCreatedByEmailNil

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetCreatedByEmailNil(b bool)`

 SetCreatedByEmailNil sets the value for CreatedByEmail to be an explicit nil

### UnsetCreatedByEmail
`func (o *ListWorkflowVersions200ResponseVersionsInner) UnsetCreatedByEmail()`

UnsetCreatedByEmail ensures that no value is present for CreatedByEmail, not even an explicit nil
### GetRestoredFromVersion

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetRestoredFromVersion() int32`

GetRestoredFromVersion returns the RestoredFromVersion field if non-nil, zero value otherwise.

### GetRestoredFromVersionOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetRestoredFromVersionOk() (*int32, bool)`

GetRestoredFromVersionOk returns a tuple with the RestoredFromVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoredFromVersion

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetRestoredFromVersion(v int32)`

SetRestoredFromVersion sets RestoredFromVersion field to given value.

### HasRestoredFromVersion

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasRestoredFromVersion() bool`

HasRestoredFromVersion returns a boolean if a field has been set.

### SetRestoredFromVersionNil

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetRestoredFromVersionNil(b bool)`

 SetRestoredFromVersionNil sets the value for RestoredFromVersion to be an explicit nil

### UnsetRestoredFromVersion
`func (o *ListWorkflowVersions200ResponseVersionsInner) UnsetRestoredFromVersion()`

UnsetRestoredFromVersion ensures that no value is present for RestoredFromVersion, not even an explicit nil
### GetCreatedAt

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListWorkflowVersions200ResponseVersionsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListWorkflowVersions200ResponseVersionsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListWorkflowVersions200ResponseVersionsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


