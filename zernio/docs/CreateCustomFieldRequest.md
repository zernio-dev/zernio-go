# CreateCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**Name** | **string** |  | 
**Slug** | Pointer to **string** | Auto-generated from name if not provided | [optional] 
**Type** | **string** |  | 
**Options** | Pointer to **[]string** | Required for select type | [optional] 

## Methods

### NewCreateCustomFieldRequest

`func NewCreateCustomFieldRequest(profileId string, name string, type_ string, ) *CreateCustomFieldRequest`

NewCreateCustomFieldRequest instantiates a new CreateCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomFieldRequestWithDefaults

`func NewCreateCustomFieldRequestWithDefaults() *CreateCustomFieldRequest`

NewCreateCustomFieldRequestWithDefaults instantiates a new CreateCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateCustomFieldRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateCustomFieldRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateCustomFieldRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetName

`func (o *CreateCustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.


### GetSlug

`func (o *CreateCustomFieldRequest) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateCustomFieldRequest) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateCustomFieldRequest) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateCustomFieldRequest) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetType

`func (o *CreateCustomFieldRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateCustomFieldRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateCustomFieldRequest) SetType(v string)`

SetType sets Type field to given value.


### GetOptions

`func (o *CreateCustomFieldRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CreateCustomFieldRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CreateCustomFieldRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CreateCustomFieldRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


