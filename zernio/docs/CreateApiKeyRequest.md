# CreateApiKeyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ExpiresIn** | Pointer to **int32** | Days until expiry | [optional] 
**Scope** | Pointer to **string** | &#39;full&#39; grants access to all profiles (default), &#39;profiles&#39; restricts to specific profiles | [optional] [default to "full"]
**ProfileIds** | Pointer to **[]string** | Profile IDs this key can access. Required when scope is &#39;profiles&#39;. | [optional] 
**Permission** | Pointer to **string** | &#39;read-write&#39; allows all operations (default), &#39;read&#39; restricts to GET requests only | [optional] [default to "read-write"]

## Methods

### NewCreateApiKeyRequest

`func NewCreateApiKeyRequest(name string, ) *CreateApiKeyRequest`

NewCreateApiKeyRequest instantiates a new CreateApiKeyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKeyRequestWithDefaults

`func NewCreateApiKeyRequestWithDefaults() *CreateApiKeyRequest`

NewCreateApiKeyRequestWithDefaults instantiates a new CreateApiKeyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateApiKeyRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateApiKeyRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateApiKeyRequest) SetName(v string)`

SetName sets Name field to given value.


### GetExpiresIn

`func (o *CreateApiKeyRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *CreateApiKeyRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *CreateApiKeyRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *CreateApiKeyRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetScope

`func (o *CreateApiKeyRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CreateApiKeyRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CreateApiKeyRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CreateApiKeyRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetProfileIds

`func (o *CreateApiKeyRequest) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *CreateApiKeyRequest) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *CreateApiKeyRequest) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *CreateApiKeyRequest) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetPermission

`func (o *CreateApiKeyRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *CreateApiKeyRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *CreateApiKeyRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *CreateApiKeyRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


