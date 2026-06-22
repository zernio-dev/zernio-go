# CreateStandaloneAdRequestPlacementAssetsRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrl** | Pointer to **string** | Image mode. The image to deliver for this rule&#39;s placements. | [optional] 
**VideoUrl** | Pointer to **string** | Video mode. The video to deliver for this rule&#39;s placements. | [optional] 
**ThumbnailUrl** | Pointer to **string** | Video mode (optional). Poster image for this rule&#39;s video; auto-generated when omitted. | [optional] 
**Placements** | [**CreateStandaloneAdRequestPlacements**](CreateStandaloneAdRequestPlacements.md) |  | 

## Methods

### NewCreateStandaloneAdRequestPlacementAssetsRulesInner

`func NewCreateStandaloneAdRequestPlacementAssetsRulesInner(placements CreateStandaloneAdRequestPlacements, ) *CreateStandaloneAdRequestPlacementAssetsRulesInner`

NewCreateStandaloneAdRequestPlacementAssetsRulesInner instantiates a new CreateStandaloneAdRequestPlacementAssetsRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestPlacementAssetsRulesInnerWithDefaults

`func NewCreateStandaloneAdRequestPlacementAssetsRulesInnerWithDefaults() *CreateStandaloneAdRequestPlacementAssetsRulesInner`

NewCreateStandaloneAdRequestPlacementAssetsRulesInnerWithDefaults instantiates a new CreateStandaloneAdRequestPlacementAssetsRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideoUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetPlacements

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetPlacements() CreateStandaloneAdRequestPlacements`

GetPlacements returns the Placements field if non-nil, zero value otherwise.

### GetPlacementsOk

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) GetPlacementsOk() (*CreateStandaloneAdRequestPlacements, bool)`

GetPlacementsOk returns a tuple with the Placements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacements

`func (o *CreateStandaloneAdRequestPlacementAssetsRulesInner) SetPlacements(v CreateStandaloneAdRequestPlacements)`

SetPlacements sets Placements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


