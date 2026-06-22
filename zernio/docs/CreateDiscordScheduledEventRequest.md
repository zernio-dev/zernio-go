# CreateDiscordScheduledEventRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**StartsAt** | **time.Time** | ISO 8601 start time. Must be in the future. | 
**Entity** | [**CreateDiscordScheduledEventRequestEntity**](CreateDiscordScheduledEventRequestEntity.md) |  | 
**ImageDataUri** | Pointer to **string** | Optional cover image as a base64 data URI. | [optional] 

## Methods

### NewCreateDiscordScheduledEventRequest

`func NewCreateDiscordScheduledEventRequest(accountId string, name string, startsAt time.Time, entity CreateDiscordScheduledEventRequestEntity, ) *CreateDiscordScheduledEventRequest`

NewCreateDiscordScheduledEventRequest instantiates a new CreateDiscordScheduledEventRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDiscordScheduledEventRequestWithDefaults

`func NewCreateDiscordScheduledEventRequestWithDefaults() *CreateDiscordScheduledEventRequest`

NewCreateDiscordScheduledEventRequestWithDefaults instantiates a new CreateDiscordScheduledEventRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateDiscordScheduledEventRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateDiscordScheduledEventRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateDiscordScheduledEventRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CreateDiscordScheduledEventRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDiscordScheduledEventRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDiscordScheduledEventRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateDiscordScheduledEventRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDiscordScheduledEventRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDiscordScheduledEventRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDiscordScheduledEventRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStartsAt

`func (o *CreateDiscordScheduledEventRequest) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *CreateDiscordScheduledEventRequest) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *CreateDiscordScheduledEventRequest) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEntity

`func (o *CreateDiscordScheduledEventRequest) GetEntity() CreateDiscordScheduledEventRequestEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *CreateDiscordScheduledEventRequest) GetEntityOk() (*CreateDiscordScheduledEventRequestEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *CreateDiscordScheduledEventRequest) SetEntity(v CreateDiscordScheduledEventRequestEntity)`

SetEntity sets Entity field to given value.


### GetImageDataUri

`func (o *CreateDiscordScheduledEventRequest) GetImageDataUri() string`

GetImageDataUri returns the ImageDataUri field if non-nil, zero value otherwise.

### GetImageDataUriOk

`func (o *CreateDiscordScheduledEventRequest) GetImageDataUriOk() (*string, bool)`

GetImageDataUriOk returns a tuple with the ImageDataUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageDataUri

`func (o *CreateDiscordScheduledEventRequest) SetImageDataUri(v string)`

SetImageDataUri sets ImageDataUri field to given value.

### HasImageDataUri

`func (o *CreateDiscordScheduledEventRequest) HasImageDataUri() bool`

HasImageDataUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


