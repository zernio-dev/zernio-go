# ListTrackingTags200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to [**[]TrackingTag**](TrackingTag.md) |  | [optional] 

## Methods

### NewListTrackingTags200Response

`func NewListTrackingTags200Response() *ListTrackingTags200Response`

NewListTrackingTags200Response instantiates a new ListTrackingTags200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTrackingTags200ResponseWithDefaults

`func NewListTrackingTags200ResponseWithDefaults() *ListTrackingTags200Response`

NewListTrackingTags200ResponseWithDefaults instantiates a new ListTrackingTags200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *ListTrackingTags200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListTrackingTags200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListTrackingTags200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListTrackingTags200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetTags

`func (o *ListTrackingTags200Response) GetTags() []TrackingTag`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListTrackingTags200Response) GetTagsOk() (*[]TrackingTag, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListTrackingTags200Response) SetTags(v []TrackingTag)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListTrackingTags200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


