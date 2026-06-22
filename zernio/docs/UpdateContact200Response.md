# UpdateContact200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Contact** | Pointer to [**UpdateContact200ResponseContact**](UpdateContact200ResponseContact.md) |  | [optional] 

## Methods

### NewUpdateContact200Response

`func NewUpdateContact200Response() *UpdateContact200Response`

NewUpdateContact200Response instantiates a new UpdateContact200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateContact200ResponseWithDefaults

`func NewUpdateContact200ResponseWithDefaults() *UpdateContact200Response`

NewUpdateContact200ResponseWithDefaults instantiates a new UpdateContact200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateContact200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateContact200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateContact200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateContact200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetContact

`func (o *UpdateContact200Response) GetContact() UpdateContact200ResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *UpdateContact200Response) GetContactOk() (*UpdateContact200ResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *UpdateContact200Response) SetContact(v UpdateContact200ResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *UpdateContact200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


