# ListUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentUserId** | Pointer to **string** |  | [optional] 
**Users** | Pointer to [**[]ListUsers200ResponseUsersInner**](ListUsers200ResponseUsersInner.md) |  | [optional] 

## Methods

### NewListUsers200Response

`func NewListUsers200Response() *ListUsers200Response`

NewListUsers200Response instantiates a new ListUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUsers200ResponseWithDefaults

`func NewListUsers200ResponseWithDefaults() *ListUsers200Response`

NewListUsers200ResponseWithDefaults instantiates a new ListUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentUserId

`func (o *ListUsers200Response) GetCurrentUserId() string`

GetCurrentUserId returns the CurrentUserId field if non-nil, zero value otherwise.

### GetCurrentUserIdOk

`func (o *ListUsers200Response) GetCurrentUserIdOk() (*string, bool)`

GetCurrentUserIdOk returns a tuple with the CurrentUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentUserId

`func (o *ListUsers200Response) SetCurrentUserId(v string)`

SetCurrentUserId sets CurrentUserId field to given value.

### HasCurrentUserId

`func (o *ListUsers200Response) HasCurrentUserId() bool`

HasCurrentUserId returns a boolean if a field has been set.

### GetUsers

`func (o *ListUsers200Response) GetUsers() []ListUsers200ResponseUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ListUsers200Response) GetUsersOk() (*[]ListUsers200ResponseUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ListUsers200Response) SetUsers(v []ListUsers200ResponseUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ListUsers200Response) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


