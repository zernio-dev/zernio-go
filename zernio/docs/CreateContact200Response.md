# CreateContact200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Contact** | Pointer to [**CreateContact200ResponseContact**](CreateContact200ResponseContact.md) |  | [optional] 
**Channel** | Pointer to [**CreateContact200ResponseChannel**](CreateContact200ResponseChannel.md) |  | [optional] 
**Warning** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateContact200Response

`func NewCreateContact200Response() *CreateContact200Response`

NewCreateContact200Response instantiates a new CreateContact200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContact200ResponseWithDefaults

`func NewCreateContact200ResponseWithDefaults() *CreateContact200Response`

NewCreateContact200ResponseWithDefaults instantiates a new CreateContact200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateContact200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateContact200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateContact200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateContact200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetContact

`func (o *CreateContact200Response) GetContact() CreateContact200ResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *CreateContact200Response) GetContactOk() (*CreateContact200ResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *CreateContact200Response) SetContact(v CreateContact200ResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *CreateContact200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetChannel

`func (o *CreateContact200Response) GetChannel() CreateContact200ResponseChannel`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *CreateContact200Response) GetChannelOk() (*CreateContact200ResponseChannel, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *CreateContact200Response) SetChannel(v CreateContact200ResponseChannel)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *CreateContact200Response) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetWarning

`func (o *CreateContact200Response) GetWarning() string`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *CreateContact200Response) GetWarningOk() (*string, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *CreateContact200Response) SetWarning(v string)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *CreateContact200Response) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


