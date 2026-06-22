# CreateDiscordScheduledEventRequestEntityOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Location** | **string** | Where the event takes place (e.g. \&quot;Zoom link\&quot;, \&quot;123 Main St\&quot;) | 
**EndsAt** | **time.Time** |  | 

## Methods

### NewCreateDiscordScheduledEventRequestEntityOneOf

`func NewCreateDiscordScheduledEventRequestEntityOneOf(type_ string, location string, endsAt time.Time, ) *CreateDiscordScheduledEventRequestEntityOneOf`

NewCreateDiscordScheduledEventRequestEntityOneOf instantiates a new CreateDiscordScheduledEventRequestEntityOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscordScheduledEventRequestEntityOneOfWithDefaults

`func NewCreateDiscordScheduledEventRequestEntityOneOfWithDefaults() *CreateDiscordScheduledEventRequestEntityOneOf`

NewCreateDiscordScheduledEventRequestEntityOneOfWithDefaults instantiates a new CreateDiscordScheduledEventRequestEntityOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetLocation

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *CreateDiscordScheduledEventRequestEntityOneOf) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


