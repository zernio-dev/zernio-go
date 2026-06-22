# ListCustomFields200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Fields** | Pointer to [**[]ListCustomFields200ResponseFieldsInner**](ListCustomFields200ResponseFieldsInner.md) |  | [optional] 

## Methods

### NewListCustomFields200Response

`func NewListCustomFields200Response() *ListCustomFields200Response`

NewListCustomFields200Response instantiates a new ListCustomFields200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomFields200ResponseWithDefaults

`func NewListCustomFields200ResponseWithDefaults() *ListCustomFields200Response`

NewListCustomFields200ResponseWithDefaults instantiates a new ListCustomFields200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListCustomFields200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListCustomFields200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListCustomFields200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListCustomFields200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFields

`func (o *ListCustomFields200Response) GetFields() []ListCustomFields200ResponseFieldsInner`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ListCustomFields200Response) GetFieldsOk() (*[]ListCustomFields200ResponseFieldsInner, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ListCustomFields200Response) SetFields(v []ListCustomFields200ResponseFieldsInner)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ListCustomFields200Response) HasFields() bool`

HasFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


