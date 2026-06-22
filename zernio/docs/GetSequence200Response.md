# GetSequence200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Sequence** | Pointer to [**GetSequence200ResponseSequence**](GetSequence200ResponseSequence.md) |  | [optional] 

## Methods

### NewGetSequence200Response

`func NewGetSequence200Response() *GetSequence200Response`

NewGetSequence200Response instantiates a new GetSequence200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSequence200ResponseWithDefaults

`func NewGetSequence200ResponseWithDefaults() *GetSequence200Response`

NewGetSequence200ResponseWithDefaults instantiates a new GetSequence200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetSequence200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetSequence200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetSequence200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetSequence200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSequence

`func (o *GetSequence200Response) GetSequence() GetSequence200ResponseSequence`

GetSequence returns the Sequence field if non-nil, zero value otherwise.

### GetSequenceOk

`func (o *GetSequence200Response) GetSequenceOk() (*GetSequence200ResponseSequence, bool)`

GetSequenceOk returns a tuple with the Sequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSequence

`func (o *GetSequence200Response) SetSequence(v GetSequence200ResponseSequence)`

SetSequence sets Sequence field to given value.

### HasSequence

`func (o *GetSequence200Response) HasSequence() bool`

HasSequence returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


