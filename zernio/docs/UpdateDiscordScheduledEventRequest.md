# UpdateDiscordScheduledEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to **string** | For external events. | [optional] 
**Status** | Pointer to **string** | Status transition. Most common: &#39;cancelled&#39; to cancel an event. | [optional] 
**ImageDataUri** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDiscordScheduledEventRequest

`func NewUpdateDiscordScheduledEventRequest(accountId string, ) *UpdateDiscordScheduledEventRequest`

NewUpdateDiscordScheduledEventRequest instantiates a new UpdateDiscordScheduledEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDiscordScheduledEventRequestWithDefaults

`func NewUpdateDiscordScheduledEventRequestWithDefaults() *UpdateDiscordScheduledEventRequest`

NewUpdateDiscordScheduledEventRequestWithDefaults instantiates a new UpdateDiscordScheduledEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateDiscordScheduledEventRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateDiscordScheduledEventRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateDiscordScheduledEventRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *UpdateDiscordScheduledEventRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDiscordScheduledEventRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDiscordScheduledEventRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDiscordScheduledEventRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDiscordScheduledEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDiscordScheduledEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDiscordScheduledEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDiscordScheduledEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartsAt

`func (o *UpdateDiscordScheduledEventRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *UpdateDiscordScheduledEventRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *UpdateDiscordScheduledEventRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *UpdateDiscordScheduledEventRequest) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *UpdateDiscordScheduledEventRequest) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *UpdateDiscordScheduledEventRequest) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *UpdateDiscordScheduledEventRequest) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *UpdateDiscordScheduledEventRequest) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateDiscordScheduledEventRequest) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateDiscordScheduledEventRequest) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateDiscordScheduledEventRequest) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateDiscordScheduledEventRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateDiscordScheduledEventRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDiscordScheduledEventRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDiscordScheduledEventRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDiscordScheduledEventRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetImageDataUri

`func (o *UpdateDiscordScheduledEventRequest) GetImageDataUri() string`

GetImageDataUri returns the ImageDataUri field if non-nil, zero value otherwise.

### GetImageDataUriOk

`func (o *UpdateDiscordScheduledEventRequest) GetImageDataUriOk() (*string, bool)`

GetImageDataUriOk returns a tuple with the ImageDataUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDataUri

`func (o *UpdateDiscordScheduledEventRequest) SetImageDataUri(v string)`

SetImageDataUri sets ImageDataUri field to given value.

### HasImageDataUri

`func (o *UpdateDiscordScheduledEventRequest) HasImageDataUri() bool`

HasImageDataUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


