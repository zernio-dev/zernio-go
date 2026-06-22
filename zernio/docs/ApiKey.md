# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**KeyPreview** | Pointer to **string** |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Key** | Pointer to **string** | Returned only once, on creation | [optional] 
**Scope** | Pointer to **string** | &#39;full&#39; grants access to all profiles, &#39;profiles&#39; restricts to specific profiles | [optional] [default to "full"]
**ProfileIds** | Pointer to [**[]ApiKeyProfileIdsInner**](ApiKeyProfileIdsInner.md) | Profiles this key can access (populated with name and color). Only present when scope is &#39;profiles&#39;. | [optional] 
**Permission** | Pointer to **string** | &#39;read-write&#39; allows all operations, &#39;read&#39; restricts to GET requests only | [optional] [default to "read-write"]

## Methods

### NewApiKey

`func NewApiKey() *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApiKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiKey) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiKey) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeyPreview

`func (o *ApiKey) GetKeyPreview() string`

GetKeyPreview returns the KeyPreview field if non-nil, zero value otherwise.

### GetKeyPreviewOk

`func (o *ApiKey) GetKeyPreviewOk() (*string, bool)`

GetKeyPreviewOk returns a tuple with the KeyPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPreview

`func (o *ApiKey) SetKeyPreview(v string)`

SetKeyPreview sets KeyPreview field to given value.

### HasKeyPreview

`func (o *ApiKey) HasKeyPreview() bool`

HasKeyPreview returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ApiKey) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ApiKey) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ApiKey) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ApiKey) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiKey) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiKey) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiKey) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiKey) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetKey

`func (o *ApiKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApiKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApiKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ApiKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetScope

`func (o *ApiKey) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiKey) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiKey) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiKey) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetProfileIds

`func (o *ApiKey) GetProfileIds() []ApiKeyProfileIdsInner`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *ApiKey) GetProfileIdsOk() (*[]ApiKeyProfileIdsInner, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *ApiKey) SetProfileIds(v []ApiKeyProfileIdsInner)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *ApiKey) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetPermission

`func (o *ApiKey) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *ApiKey) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *ApiKey) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *ApiKey) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


