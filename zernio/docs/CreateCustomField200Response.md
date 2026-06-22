# CreateCustomField200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Field** | Pointer to [**ListCustomFields200ResponseFieldsInner**](ListCustomFields200ResponseFieldsInner.md) |  | [optional] 

## Methods

### NewCreateCustomField200Response

`func NewCreateCustomField200Response() *CreateCustomField200Response`

NewCreateCustomField200Response instantiates a new CreateCustomField200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomField200ResponseWithDefaults

`func NewCreateCustomField200ResponseWithDefaults() *CreateCustomField200Response`

NewCreateCustomField200ResponseWithDefaults instantiates a new CreateCustomField200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateCustomField200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateCustomField200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateCustomField200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateCustomField200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetField

`func (o *CreateCustomField200Response) GetField() ListCustomFields200ResponseFieldsInner`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *CreateCustomField200Response) GetFieldOk() (*ListCustomFields200ResponseFieldsInner, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *CreateCustomField200Response) SetField(v ListCustomFields200ResponseFieldsInner)`

SetField sets Field field to given value.

### HasField

`func (o *CreateCustomField200Response) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


