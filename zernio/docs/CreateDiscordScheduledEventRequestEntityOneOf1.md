# CreateDiscordScheduledEventRequestEntityOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**ChannelId** | **string** | Voice channel snowflake. | 
**EndsAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateDiscordScheduledEventRequestEntityOneOf1

`func NewCreateDiscordScheduledEventRequestEntityOneOf1(type_ string, channelId string, ) *CreateDiscordScheduledEventRequestEntityOneOf1`

NewCreateDiscordScheduledEventRequestEntityOneOf1 instantiates a new CreateDiscordScheduledEventRequestEntityOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscordScheduledEventRequestEntityOneOf1WithDefaults

`func NewCreateDiscordScheduledEventRequestEntityOneOf1WithDefaults() *CreateDiscordScheduledEventRequestEntityOneOf1`

NewCreateDiscordScheduledEventRequestEntityOneOf1WithDefaults instantiates a new CreateDiscordScheduledEventRequestEntityOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetChannelId

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf1) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


