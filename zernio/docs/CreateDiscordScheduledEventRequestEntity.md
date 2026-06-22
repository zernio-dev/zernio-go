# CreateDiscordScheduledEventRequestEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Location** | **string** | Where the event takes place (e.g. \&quot;Zoom link\&quot;, \&quot;123 Main St\&quot;) | 
**EndsAt** | **time.Time** |  | 
**ChannelId** | **string** | Stage channel snowflake. | 

## Methods

### NewCreateDiscordScheduledEventRequestEntity

`func NewCreateDiscordScheduledEventRequestEntity(type_ string, location string, endsAt time.Time, channelId string, ) *CreateDiscordScheduledEventRequestEntity`

NewCreateDiscordScheduledEventRequestEntity instantiates a new CreateDiscordScheduledEventRequestEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscordScheduledEventRequestEntityWithDefaults

`func NewCreateDiscordScheduledEventRequestEntityWithDefaults() *CreateDiscordScheduledEventRequestEntity`

NewCreateDiscordScheduledEventRequestEntityWithDefaults instantiates a new CreateDiscordScheduledEventRequestEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateDiscordScheduledEventRequestEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDiscordScheduledEventRequestEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDiscordScheduledEventRequestEntity) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *CreateDiscordScheduledEventRequestEntity) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateDiscordScheduledEventRequestEntity) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateDiscordScheduledEventRequestEntity) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntity) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateDiscordScheduledEventRequestEntity) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntity) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetChannelId

`func (o *CreateDiscordScheduledEventRequestEntity) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *CreateDiscordScheduledEventRequestEntity) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *CreateDiscordScheduledEventRequestEntity) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


