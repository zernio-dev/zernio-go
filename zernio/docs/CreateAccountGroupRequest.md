# CreateAccountGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**AccountIds** | **[]string** |  | 
**ProfileId** | Pointer to **string** | Deprecated. Accepted for backward compatibility but ignored. Groups are no longer scoped to a single profile.  | [optional] 

## Methods

### NewCreateAccountGroupRequest

`func NewCreateAccountGroupRequest(name string, accountIds []string, ) *CreateAccountGroupRequest`

NewCreateAccountGroupRequest instantiates a new CreateAccountGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountGroupRequestWithDefaults

`func NewCreateAccountGroupRequestWithDefaults() *CreateAccountGroupRequest`

NewCreateAccountGroupRequestWithDefaults instantiates a new CreateAccountGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateAccountGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAccountGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAccountGroupRequest) SetName(v string)`

SetName sets Name field to given value.


### GetAccountIds

`func (o *CreateAccountGroupRequest) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *CreateAccountGroupRequest) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *CreateAccountGroupRequest) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.


### GetProfileId

`func (o *CreateAccountGroupRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateAccountGroupRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateAccountGroupRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *CreateAccountGroupRequest) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


