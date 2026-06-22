# GetDiscordChannels200ResponseChannelsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Channel snowflake ID | [optional] 
**Name** | Pointer to **string** | Channel name | [optional] 
**Type** | Pointer to **int32** | Channel type: 0 (text), 5 (announcement), 15 (forum) | [optional] 

## Methods

### NewGetDiscordChannels200ResponseChannelsInner

`func NewGetDiscordChannels200ResponseChannelsInner() *GetDiscordChannels200ResponseChannelsInner`

NewGetDiscordChannels200ResponseChannelsInner instantiates a new GetDiscordChannels200ResponseChannelsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDiscordChannels200ResponseChannelsInnerWithDefaults

`func NewGetDiscordChannels200ResponseChannelsInnerWithDefaults() *GetDiscordChannels200ResponseChannelsInner`

NewGetDiscordChannels200ResponseChannelsInnerWithDefaults instantiates a new GetDiscordChannels200ResponseChannelsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDiscordChannels200ResponseChannelsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDiscordChannels200ResponseChannelsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDiscordChannels200ResponseChannelsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetDiscordChannels200ResponseChannelsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetDiscordChannels200ResponseChannelsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDiscordChannels200ResponseChannelsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDiscordChannels200ResponseChannelsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDiscordChannels200ResponseChannelsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetDiscordChannels200ResponseChannelsInner) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetDiscordChannels200ResponseChannelsInner) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetDiscordChannels200ResponseChannelsInner) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GetDiscordChannels200ResponseChannelsInner) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


