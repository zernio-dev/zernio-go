# UpdateCustomFieldRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateCustomFieldRequest

`func NewUpdateCustomFieldRequest() *UpdateCustomFieldRequest`

NewUpdateCustomFieldRequest instantiates a new UpdateCustomFieldRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCustomFieldRequestWithDefaults

`func NewUpdateCustomFieldRequestWithDefaults() *UpdateCustomFieldRequest`

NewUpdateCustomFieldRequestWithDefaults instantiates a new UpdateCustomFieldRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCustomFieldRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCustomFieldRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCustomFieldRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCustomFieldRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOptions

`func (o *UpdateCustomFieldRequest) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateCustomFieldRequest) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateCustomFieldRequest) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateCustomFieldRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


