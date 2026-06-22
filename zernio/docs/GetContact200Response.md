# GetContact200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Contact** | Pointer to [**GetContact200ResponseContact**](GetContact200ResponseContact.md) |  | [optional] 
**Channels** | Pointer to [**[]GetContact200ResponseChannelsInner**](GetContact200ResponseChannelsInner.md) |  | [optional] 

## Methods

### NewGetContact200Response

`func NewGetContact200Response() *GetContact200Response`

NewGetContact200Response instantiates a new GetContact200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContact200ResponseWithDefaults

`func NewGetContact200ResponseWithDefaults() *GetContact200Response`

NewGetContact200ResponseWithDefaults instantiates a new GetContact200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetContact200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetContact200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetContact200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetContact200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetContact

`func (o *GetContact200Response) GetContact() GetContact200ResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GetContact200Response) GetContactOk() (*GetContact200ResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GetContact200Response) SetContact(v GetContact200ResponseContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *GetContact200Response) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetChannels

`func (o *GetContact200Response) GetChannels() []GetContact200ResponseChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *GetContact200Response) GetChannelsOk() (*[]GetContact200ResponseChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *GetContact200Response) SetChannels(v []GetContact200ResponseChannelsInner)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *GetContact200Response) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


