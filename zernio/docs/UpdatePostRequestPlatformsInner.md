# UpdatePostRequestPlatformsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**AccountId** | **string** |  | 
**CustomContent** | Pointer to **string** | Platform-specific text override. | [optional] 
**CustomMedia** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** | Optional per-platform scheduled time override. | [optional] 
**PlatformSpecificData** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdatePostRequestPlatformsInner

`func NewUpdatePostRequestPlatformsInner(platform string, accountId string, ) *UpdatePostRequestPlatformsInner`

NewUpdatePostRequestPlatformsInner instantiates a new UpdatePostRequestPlatformsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostRequestPlatformsInnerWithDefaults

`func NewUpdatePostRequestPlatformsInnerWithDefaults() *UpdatePostRequestPlatformsInner`

NewUpdatePostRequestPlatformsInnerWithDefaults instantiates a new UpdatePostRequestPlatformsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *UpdatePostRequestPlatformsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdatePostRequestPlatformsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdatePostRequestPlatformsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAccountId

`func (o *UpdatePostRequestPlatformsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdatePostRequestPlatformsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdatePostRequestPlatformsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomContent

`func (o *UpdatePostRequestPlatformsInner) GetCustomContent() string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *UpdatePostRequestPlatformsInner) GetCustomContentOk() (*string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *UpdatePostRequestPlatformsInner) SetCustomContent(v string)`

SetCustomContent sets CustomContent field to given value.

### HasCustomContent

`func (o *UpdatePostRequestPlatformsInner) HasCustomContent() bool`

HasCustomContent returns a boolean if a field has been set.

### GetCustomMedia

`func (o *UpdatePostRequestPlatformsInner) GetCustomMedia() []MediaItem`

GetCustomMedia returns the CustomMedia field if non-nil, zero value otherwise.

### GetCustomMediaOk

`func (o *UpdatePostRequestPlatformsInner) GetCustomMediaOk() (*[]MediaItem, bool)`

GetCustomMediaOk returns a tuple with the CustomMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMedia

`func (o *UpdatePostRequestPlatformsInner) SetCustomMedia(v []MediaItem)`

SetCustomMedia sets CustomMedia field to given value.

### HasCustomMedia

`func (o *UpdatePostRequestPlatformsInner) HasCustomMedia() bool`

HasCustomMedia returns a boolean if a field has been set.

### GetScheduledFor

`func (o *UpdatePostRequestPlatformsInner) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *UpdatePostRequestPlatformsInner) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *UpdatePostRequestPlatformsInner) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *UpdatePostRequestPlatformsInner) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPlatformSpecificData

`func (o *UpdatePostRequestPlatformsInner) GetPlatformSpecificData() map[string]interface{}`

GetPlatformSpecificData returns the PlatformSpecificData field if non-nil, zero value otherwise.

### GetPlatformSpecificDataOk

`func (o *UpdatePostRequestPlatformsInner) GetPlatformSpecificDataOk() (*map[string]interface{}, bool)`

GetPlatformSpecificDataOk returns a tuple with the PlatformSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformSpecificData

`func (o *UpdatePostRequestPlatformsInner) SetPlatformSpecificData(v map[string]interface{})`

SetPlatformSpecificData sets PlatformSpecificData field to given value.

### HasPlatformSpecificData

`func (o *UpdatePostRequestPlatformsInner) HasPlatformSpecificData() bool`

HasPlatformSpecificData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


