# CreatePostRequestPlatformsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**AccountId** | **string** |  | 
**CustomContent** | Pointer to **string** | Platform-specific text override. When set, this content is used instead of the top-level post content for this platform. Useful for tailoring captions per platform (e.g. keeping tweets under 280 characters). | [optional] 
**CustomMedia** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 
**ScheduledFor** | Pointer to **time.Time** | Optional per-platform scheduled time override. When omitted, the top-level scheduledFor is used. | [optional] 
**PlatformSpecificData** | Pointer to [**CreatePostRequestPlatformsInnerPlatformSpecificData**](CreatePostRequestPlatformsInnerPlatformSpecificData.md) |  | [optional] 

## Methods

### NewCreatePostRequestPlatformsInner

`func NewCreatePostRequestPlatformsInner(platform string, accountId string, ) *CreatePostRequestPlatformsInner`

NewCreatePostRequestPlatformsInner instantiates a new CreatePostRequestPlatformsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePostRequestPlatformsInnerWithDefaults

`func NewCreatePostRequestPlatformsInnerWithDefaults() *CreatePostRequestPlatformsInner`

NewCreatePostRequestPlatformsInnerWithDefaults instantiates a new CreatePostRequestPlatformsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *CreatePostRequestPlatformsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreatePostRequestPlatformsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreatePostRequestPlatformsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAccountId

`func (o *CreatePostRequestPlatformsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreatePostRequestPlatformsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreatePostRequestPlatformsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetCustomContent

`func (o *CreatePostRequestPlatformsInner) GetCustomContent() string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *CreatePostRequestPlatformsInner) GetCustomContentOk() (*string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *CreatePostRequestPlatformsInner) SetCustomContent(v string)`

SetCustomContent sets CustomContent field to given value.

### HasCustomContent

`func (o *CreatePostRequestPlatformsInner) HasCustomContent() bool`

HasCustomContent returns a boolean if a field has been set.

### GetCustomMedia

`func (o *CreatePostRequestPlatformsInner) GetCustomMedia() []MediaItem`

GetCustomMedia returns the CustomMedia field if non-nil, zero value otherwise.

### GetCustomMediaOk

`func (o *CreatePostRequestPlatformsInner) GetCustomMediaOk() (*[]MediaItem, bool)`

GetCustomMediaOk returns a tuple with the CustomMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMedia

`func (o *CreatePostRequestPlatformsInner) SetCustomMedia(v []MediaItem)`

SetCustomMedia sets CustomMedia field to given value.

### HasCustomMedia

`func (o *CreatePostRequestPlatformsInner) HasCustomMedia() bool`

HasCustomMedia returns a boolean if a field has been set.

### GetScheduledFor

`func (o *CreatePostRequestPlatformsInner) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *CreatePostRequestPlatformsInner) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *CreatePostRequestPlatformsInner) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.

### HasScheduledFor

`func (o *CreatePostRequestPlatformsInner) HasScheduledFor() bool`

HasScheduledFor returns a boolean if a field has been set.

### GetPlatformSpecificData

`func (o *CreatePostRequestPlatformsInner) GetPlatformSpecificData() CreatePostRequestPlatformsInnerPlatformSpecificData`

GetPlatformSpecificData returns the PlatformSpecificData field if non-nil, zero value otherwise.

### GetPlatformSpecificDataOk

`func (o *CreatePostRequestPlatformsInner) GetPlatformSpecificDataOk() (*CreatePostRequestPlatformsInnerPlatformSpecificData, bool)`

GetPlatformSpecificDataOk returns a tuple with the PlatformSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformSpecificData

`func (o *CreatePostRequestPlatformsInner) SetPlatformSpecificData(v CreatePostRequestPlatformsInnerPlatformSpecificData)`

SetPlatformSpecificData sets PlatformSpecificData field to given value.

### HasPlatformSpecificData

`func (o *CreatePostRequestPlatformsInner) HasPlatformSpecificData() bool`

HasPlatformSpecificData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


