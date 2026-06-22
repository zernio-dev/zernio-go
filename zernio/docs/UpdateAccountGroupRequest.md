# UpdateAccountGroupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AccountIds** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateAccountGroupRequest

`func NewUpdateAccountGroupRequest() *UpdateAccountGroupRequest`

NewUpdateAccountGroupRequest instantiates a new UpdateAccountGroupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountGroupRequestWithDefaults

`func NewUpdateAccountGroupRequestWithDefaults() *UpdateAccountGroupRequest`

NewUpdateAccountGroupRequestWithDefaults instantiates a new UpdateAccountGroupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateAccountGroupRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAccountGroupRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAccountGroupRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAccountGroupRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountIds

`func (o *UpdateAccountGroupRequest) GetAccountIds() []string`

GetAccountIds returns the AccountIds field if non-nil, zero value otherwise.

### GetAccountIdsOk

`func (o *UpdateAccountGroupRequest) GetAccountIdsOk() (*[]string, bool)`

GetAccountIdsOk returns a tuple with the AccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIds

`func (o *UpdateAccountGroupRequest) SetAccountIds(v []string)`

SetAccountIds sets AccountIds field to given value.

### HasAccountIds

`func (o *UpdateAccountGroupRequest) HasAccountIds() bool`

HasAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


