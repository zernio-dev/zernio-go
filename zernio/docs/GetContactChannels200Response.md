# GetContactChannels200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Channels** | Pointer to [**[]GetContactChannels200ResponseChannelsInner**](GetContactChannels200ResponseChannelsInner.md) |  | [optional] 

## Methods

### NewGetContactChannels200Response

`func NewGetContactChannels200Response() *GetContactChannels200Response`

NewGetContactChannels200Response instantiates a new GetContactChannels200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContactChannels200ResponseWithDefaults

`func NewGetContactChannels200ResponseWithDefaults() *GetContactChannels200Response`

NewGetContactChannels200ResponseWithDefaults instantiates a new GetContactChannels200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetContactChannels200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetContactChannels200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetContactChannels200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetContactChannels200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetChannels

`func (o *GetContactChannels200Response) GetChannels() []GetContactChannels200ResponseChannelsInner`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *GetContactChannels200Response) GetChannelsOk() (*[]GetContactChannels200ResponseChannelsInner, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *GetContactChannels200Response) SetChannels(v []GetContactChannels200ResponseChannelsInner)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *GetContactChannels200Response) HasChannels() bool`

HasChannels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


