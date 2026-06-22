# EnrollContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Enrolled** | Pointer to **int32** | Number of contacts successfully enrolled | [optional] 
**Skipped** | Pointer to **int32** | Number skipped (already enrolled or missing channel) | [optional] 

## Methods

### NewEnrollContacts200Response

`func NewEnrollContacts200Response() *EnrollContacts200Response`

NewEnrollContacts200Response instantiates a new EnrollContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollContacts200ResponseWithDefaults

`func NewEnrollContacts200ResponseWithDefaults() *EnrollContacts200Response`

NewEnrollContacts200ResponseWithDefaults instantiates a new EnrollContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *EnrollContacts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *EnrollContacts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *EnrollContacts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *EnrollContacts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetEnrolled

`func (o *EnrollContacts200Response) GetEnrolled() int32`

GetEnrolled returns the Enrolled field if non-nil, zero value otherwise.

### GetEnrolledOk

`func (o *EnrollContacts200Response) GetEnrolledOk() (*int32, bool)`

GetEnrolledOk returns a tuple with the Enrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrolled

`func (o *EnrollContacts200Response) SetEnrolled(v int32)`

SetEnrolled sets Enrolled field to given value.

### HasEnrolled

`func (o *EnrollContacts200Response) HasEnrolled() bool`

HasEnrolled returns a boolean if a field has been set.

### GetSkipped

`func (o *EnrollContacts200Response) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *EnrollContacts200Response) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *EnrollContacts200Response) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *EnrollContacts200Response) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


