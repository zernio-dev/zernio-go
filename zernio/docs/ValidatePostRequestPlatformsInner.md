# ValidatePostRequestPlatformsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**CustomContent** | Pointer to **string** |  | [optional] 
**PlatformSpecificData** | Pointer to **map[string]interface{}** |  | [optional] 
**CustomMedia** | Pointer to [**[]MediaItem**](MediaItem.md) |  | [optional] 

## Methods

### NewValidatePostRequestPlatformsInner

`func NewValidatePostRequestPlatformsInner(platform string, ) *ValidatePostRequestPlatformsInner`

NewValidatePostRequestPlatformsInner instantiates a new ValidatePostRequestPlatformsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePostRequestPlatformsInnerWithDefaults

`func NewValidatePostRequestPlatformsInnerWithDefaults() *ValidatePostRequestPlatformsInner`

NewValidatePostRequestPlatformsInnerWithDefaults instantiates a new ValidatePostRequestPlatformsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *ValidatePostRequestPlatformsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ValidatePostRequestPlatformsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ValidatePostRequestPlatformsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCustomContent

`func (o *ValidatePostRequestPlatformsInner) GetCustomContent() string`

GetCustomContent returns the CustomContent field if non-nil, zero value otherwise.

### GetCustomContentOk

`func (o *ValidatePostRequestPlatformsInner) GetCustomContentOk() (*string, bool)`

GetCustomContentOk returns a tuple with the CustomContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomContent

`func (o *ValidatePostRequestPlatformsInner) SetCustomContent(v string)`

SetCustomContent sets CustomContent field to given value.

### HasCustomContent

`func (o *ValidatePostRequestPlatformsInner) HasCustomContent() bool`

HasCustomContent returns a boolean if a field has been set.

### GetPlatformSpecificData

`func (o *ValidatePostRequestPlatformsInner) GetPlatformSpecificData() map[string]interface{}`

GetPlatformSpecificData returns the PlatformSpecificData field if non-nil, zero value otherwise.

### GetPlatformSpecificDataOk

`func (o *ValidatePostRequestPlatformsInner) GetPlatformSpecificDataOk() (*map[string]interface{}, bool)`

GetPlatformSpecificDataOk returns a tuple with the PlatformSpecificData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformSpecificData

`func (o *ValidatePostRequestPlatformsInner) SetPlatformSpecificData(v map[string]interface{})`

SetPlatformSpecificData sets PlatformSpecificData field to given value.

### HasPlatformSpecificData

`func (o *ValidatePostRequestPlatformsInner) HasPlatformSpecificData() bool`

HasPlatformSpecificData returns a boolean if a field has been set.

### GetCustomMedia

`func (o *ValidatePostRequestPlatformsInner) GetCustomMedia() []MediaItem`

GetCustomMedia returns the CustomMedia field if non-nil, zero value otherwise.

### GetCustomMediaOk

`func (o *ValidatePostRequestPlatformsInner) GetCustomMediaOk() (*[]MediaItem, bool)`

GetCustomMediaOk returns a tuple with the CustomMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMedia

`func (o *ValidatePostRequestPlatformsInner) SetCustomMedia(v []MediaItem)`

SetCustomMedia sets CustomMedia field to given value.

### HasCustomMedia

`func (o *ValidatePostRequestPlatformsInner) HasCustomMedia() bool`

HasCustomMedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


