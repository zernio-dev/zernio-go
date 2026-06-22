# UpdatePostMetadataRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | The platform to update metadata on | 
**VideoId** | Pointer to **string** | YouTube video ID (required for direct mode, ignored for post-based mode) | [optional] 
**AccountId** | Pointer to **string** | Zernio social account ID (required for direct mode, ignored for post-based mode) | [optional] 
**Title** | Pointer to **string** | New video title (max 100 characters for YouTube) | [optional] 
**Description** | Pointer to **string** | New video description | [optional] 
**Tags** | Pointer to **[]string** | Array of keyword tags (max 500 characters combined for YouTube) | [optional] 
**CategoryId** | Pointer to **string** | YouTube video category ID | [optional] 
**PrivacyStatus** | Pointer to **string** | Video privacy setting | [optional] 
**ThumbnailUrl** | Pointer to **string** | Public URL of a custom thumbnail image (JPEG, PNG, or GIF, max 2 MB, recommended 1280x720). Works on any video you own, including existing videos not published through Zernio. The channel must be verified (phone verification) to set custom thumbnails. | [optional] 
**MadeForKids** | Pointer to **bool** | COPPA compliance flag. Set true for child-directed content (restricts comments, notifications, ad targeting). | [optional] 
**ContainsSyntheticMedia** | Pointer to **bool** | AI-generated content disclosure. Set true if the video contains synthetic content that could be mistaken for real. YouTube may add a label. | [optional] 
**PlaylistId** | Pointer to **string** | YouTube playlist ID to add the video to (e.g. &#39;PLxxxxxxxxxxxxx&#39;). Use GET /v1/accounts/{id}/youtube-playlists to list available playlists. Only playlists owned by the channel are supported. | [optional] 

## Methods

### NewUpdatePostMetadataRequest

`func NewUpdatePostMetadataRequest(platform string, ) *UpdatePostMetadataRequest`

NewUpdatePostMetadataRequest instantiates a new UpdatePostMetadataRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostMetadataRequestWithDefaults

`func NewUpdatePostMetadataRequestWithDefaults() *UpdatePostMetadataRequest`

NewUpdatePostMetadataRequestWithDefaults instantiates a new UpdatePostMetadataRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *UpdatePostMetadataRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdatePostMetadataRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdatePostMetadataRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetVideoId

`func (o *UpdatePostMetadataRequest) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *UpdatePostMetadataRequest) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *UpdatePostMetadataRequest) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *UpdatePostMetadataRequest) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdatePostMetadataRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdatePostMetadataRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdatePostMetadataRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdatePostMetadataRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetTitle

`func (o *UpdatePostMetadataRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdatePostMetadataRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdatePostMetadataRequest) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdatePostMetadataRequest) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePostMetadataRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePostMetadataRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePostMetadataRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePostMetadataRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *UpdatePostMetadataRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdatePostMetadataRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdatePostMetadataRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdatePostMetadataRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetCategoryId

`func (o *UpdatePostMetadataRequest) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *UpdatePostMetadataRequest) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *UpdatePostMetadataRequest) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *UpdatePostMetadataRequest) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPrivacyStatus

`func (o *UpdatePostMetadataRequest) GetPrivacyStatus() string`

GetPrivacyStatus returns the PrivacyStatus field if non-nil, zero value otherwise.

### GetPrivacyStatusOk

`func (o *UpdatePostMetadataRequest) GetPrivacyStatusOk() (*string, bool)`

GetPrivacyStatusOk returns a tuple with the PrivacyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyStatus

`func (o *UpdatePostMetadataRequest) SetPrivacyStatus(v string)`

SetPrivacyStatus sets PrivacyStatus field to given value.

### HasPrivacyStatus

`func (o *UpdatePostMetadataRequest) HasPrivacyStatus() bool`

HasPrivacyStatus returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *UpdatePostMetadataRequest) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *UpdatePostMetadataRequest) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *UpdatePostMetadataRequest) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *UpdatePostMetadataRequest) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetMadeForKids

`func (o *UpdatePostMetadataRequest) GetMadeForKids() bool`

GetMadeForKids returns the MadeForKids field if non-nil, zero value otherwise.

### GetMadeForKidsOk

`func (o *UpdatePostMetadataRequest) GetMadeForKidsOk() (*bool, bool)`

GetMadeForKidsOk returns a tuple with the MadeForKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadeForKids

`func (o *UpdatePostMetadataRequest) SetMadeForKids(v bool)`

SetMadeForKids sets MadeForKids field to given value.

### HasMadeForKids

`func (o *UpdatePostMetadataRequest) HasMadeForKids() bool`

HasMadeForKids returns a boolean if a field has been set.

### GetContainsSyntheticMedia

`func (o *UpdatePostMetadataRequest) GetContainsSyntheticMedia() bool`

GetContainsSyntheticMedia returns the ContainsSyntheticMedia field if non-nil, zero value otherwise.

### GetContainsSyntheticMediaOk

`func (o *UpdatePostMetadataRequest) GetContainsSyntheticMediaOk() (*bool, bool)`

GetContainsSyntheticMediaOk returns a tuple with the ContainsSyntheticMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsSyntheticMedia

`func (o *UpdatePostMetadataRequest) SetContainsSyntheticMedia(v bool)`

SetContainsSyntheticMedia sets ContainsSyntheticMedia field to given value.

### HasContainsSyntheticMedia

`func (o *UpdatePostMetadataRequest) HasContainsSyntheticMedia() bool`

HasContainsSyntheticMedia returns a boolean if a field has been set.

### GetPlaylistId

`func (o *UpdatePostMetadataRequest) GetPlaylistId() string`

GetPlaylistId returns the PlaylistId field if non-nil, zero value otherwise.

### GetPlaylistIdOk

`func (o *UpdatePostMetadataRequest) GetPlaylistIdOk() (*string, bool)`

GetPlaylistIdOk returns a tuple with the PlaylistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistId

`func (o *UpdatePostMetadataRequest) SetPlaylistId(v string)`

SetPlaylistId sets PlaylistId field to given value.

### HasPlaylistId

`func (o *UpdatePostMetadataRequest) HasPlaylistId() bool`

HasPlaylistId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


