# CreateDiscordScheduledEventRequestEntityOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ChannelId** | **string** | Stage channel snowflake. | 
**EndsAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateDiscordScheduledEventRequestEntityOneOf2

`func NewCreateDiscordScheduledEventRequestEntityOneOf2(type_ string, channelId string, ) *CreateDiscordScheduledEventRequestEntityOneOf2`

NewCreateDiscordScheduledEventRequestEntityOneOf2 instantiates a new CreateDiscordScheduledEventRequestEntityOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscordScheduledEventRequestEntityOneOf2WithDefaults

`func NewCreateDiscordScheduledEventRequestEntityOneOf2WithDefaults() *CreateDiscordScheduledEventRequestEntityOneOf2`

NewCreateDiscordScheduledEventRequestEntityOneOf2WithDefaults instantiates a new CreateDiscordScheduledEventRequestEntityOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetChannelId

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf2) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


