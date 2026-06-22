# CreateSequence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Sequence** | Pointer to [**CreateSequence200ResponseSequence**](CreateSequence200ResponseSequence.md) |  | [optional] 

## Methods

### NewCreateSequence200Response

`func NewCreateSequence200Response() *CreateSequence200Response`

NewCreateSequence200Response instantiates a new CreateSequence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSequence200ResponseWithDefaults

`func NewCreateSequence200ResponseWithDefaults() *CreateSequence200Response`

NewCreateSequence200ResponseWithDefaults instantiates a new CreateSequence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateSequence200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateSequence200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateSequence200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateSequence200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSequence

`func (o *CreateSequence200Response) GetSequence() CreateSequence200ResponseSequence`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *CreateSequence200Response) GetSequenceOk() (*CreateSequence200ResponseSequence, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *CreateSequence200Response) SetSequence(v CreateSequence200ResponseSequence)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *CreateSequence200Response) HasSequence() bool`

HasSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


