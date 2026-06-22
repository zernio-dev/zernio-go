# CreateStandaloneAdRequestTracking

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PixelId** | Pointer to **string** | Meta Pixel ID to attach for offsite-conversion measurement. | [optional] 
**UrlTags** | Pointer to [**[]UpdateAdTrackingTagsRequestUrlTagsInner**](UpdateAdTrackingTagsRequestUrlTagsInner.md) | Click-URL params appended to the ad&#39;s destination as &#x60;url_tags&#x60; (e.g. utm_source). | [optional] 

## Methods

### NewCreateStandaloneAdRequestTracking

`func NewCreateStandaloneAdRequestTracking() *CreateStandaloneAdRequestTracking`

NewCreateStandaloneAdRequestTracking instantiates a new CreateStandaloneAdRequestTracking object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestTrackingWithDefaults

`func NewCreateStandaloneAdRequestTrackingWithDefaults() *CreateStandaloneAdRequestTracking`

NewCreateStandaloneAdRequestTrackingWithDefaults instantiates a new CreateStandaloneAdRequestTracking object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPixelId

`func (o *CreateStandaloneAdRequestTracking) GetPixelId() string`

GetPixelId returns the PixelId field if non-nil, zero value otherwise.

### GetPixelIdOk

`func (o *CreateStandaloneAdRequestTracking) GetPixelIdOk() (*string, bool)`

GetPixelIdOk returns a tuple with the PixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelId

`func (o *CreateStandaloneAdRequestTracking) SetPixelId(v string)`

SetPixelId sets PixelId field to given value.

### HasPixelId

`func (o *CreateStandaloneAdRequestTracking) HasPixelId() bool`

HasPixelId returns a boolean if a field has been set.

### GetUrlTags

`func (o *CreateStandaloneAdRequestTracking) GetUrlTags() []UpdateAdTrackingTagsRequestUrlTagsInner`

GetUrlTags returns the UrlTags field if non-nil, zero value otherwise.

### GetUrlTagsOk

`func (o *CreateStandaloneAdRequestTracking) GetUrlTagsOk() (*[]UpdateAdTrackingTagsRequestUrlTagsInner, bool)`

GetUrlTagsOk returns a tuple with the UrlTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTags

`func (o *CreateStandaloneAdRequestTracking) SetUrlTags(v []UpdateAdTrackingTagsRequestUrlTagsInner)`

SetUrlTags sets UrlTags field to given value.

### HasUrlTags

`func (o *CreateStandaloneAdRequestTracking) HasUrlTags() bool`

HasUrlTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


