# ValidatePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | Post text content | [optional] 
**Platforms** | [**[]ValidatePostRequestPlatformsInner**](ValidatePostRequestPlatformsInner.md) | Target platforms (same format as POST /v1/posts) | 
**MediaItems** | Pointer to [**[]MediaItem**](MediaItem.md) | Root media items shared across platforms | [optional] 

## Methods

### NewValidatePostRequest

`func NewValidatePostRequest(platforms []ValidatePostRequestPlatformsInner, ) *ValidatePostRequest`

NewValidatePostRequest instantiates a new ValidatePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePostRequestWithDefaults

`func NewValidatePostRequestWithDefaults() *ValidatePostRequest`

NewValidatePostRequestWithDefaults instantiates a new ValidatePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *ValidatePostRequest) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ValidatePostRequest) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ValidatePostRequest) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ValidatePostRequest) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPlatforms

`func (o *ValidatePostRequest) GetPlatforms() []ValidatePostRequestPlatformsInner`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ValidatePostRequest) GetPlatformsOk() (*[]ValidatePostRequestPlatformsInner, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ValidatePostRequest) SetPlatforms(v []ValidatePostRequestPlatformsInner)`

SetPlatforms sets Platforms field to given value.


### GetMediaItems

`func (o *ValidatePostRequest) GetMediaItems() []MediaItem`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *ValidatePostRequest) GetMediaItemsOk() (*[]MediaItem, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *ValidatePostRequest) SetMediaItems(v []MediaItem)`

SetMediaItems sets MediaItems field to given value.

### HasMediaItems

`func (o *ValidatePostRequest) HasMediaItems() bool`

HasMediaItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


