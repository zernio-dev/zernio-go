# CreateAccountGroup201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Group** | Pointer to [**CreateAccountGroup201ResponseGroup**](CreateAccountGroup201ResponseGroup.md) |  | [optional] 

## Methods

### NewCreateAccountGroup201Response

`func NewCreateAccountGroup201Response() *CreateAccountGroup201Response`

NewCreateAccountGroup201Response instantiates a new CreateAccountGroup201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountGroup201ResponseWithDefaults

`func NewCreateAccountGroup201ResponseWithDefaults() *CreateAccountGroup201Response`

NewCreateAccountGroup201ResponseWithDefaults instantiates a new CreateAccountGroup201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateAccountGroup201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateAccountGroup201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateAccountGroup201Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateAccountGroup201Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetGroup

`func (o *CreateAccountGroup201Response) GetGroup() CreateAccountGroup201ResponseGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CreateAccountGroup201Response) GetGroupOk() (*CreateAccountGroup201ResponseGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CreateAccountGroup201Response) SetGroup(v CreateAccountGroup201ResponseGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *CreateAccountGroup201Response) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


