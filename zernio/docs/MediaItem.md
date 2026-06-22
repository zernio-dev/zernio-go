# MediaItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | Optional title for the media item. Used as the document title for LinkedIn PDF/carousel posts. If omitted, falls back to the post title, then the filename. | [optional] 
**AltText** | Pointer to **string** | Accessibility alternative text for an image, applied on every platform that supports it: Instagram (feed images only, not Reels/Stories), Facebook, Threads, X/Twitter (max 1000 chars), LinkedIn, Bluesky, and Pinterest (max 500 chars). Ignored on platforms without alt-text support (TikTok, YouTube, Snapchat, Telegram, Reddit, Google Business, WhatsApp) and on video items where the platform does not accept it. Set once per image; the same value is sent to each selected platform. | [optional] 
**Filename** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **int32** | Optional file size in bytes | [optional] 
**MimeType** | Pointer to **string** | Optional MIME type (e.g. image/jpeg, video/mp4) | [optional] 
**Thumbnail** | Pointer to **string** | Optional custom thumbnail/cover image URL for videos. Supported for Facebook video posts, Facebook Reels, and regular video uploads. Max 10MB, JPG/PNG recommended. | [optional] 
**InstagramThumbnail** | Pointer to **string** | Custom cover image URL for Instagram Reels. Can also be set via platformSpecificData.instagramThumbnail or platformSpecificData.reelCover. Resolution order: this field &gt; platformSpecificData.instagramThumbnail &gt; platformSpecificData.reelCover &gt; platformSpecificData.thumbnailUrl (legacy). | [optional] 
**TiktokProcessed** | Pointer to **bool** | Internal flag indicating the image was resized for TikTok | [optional] 

## Methods

### NewMediaItem

`func NewMediaItem() *MediaItem`

NewMediaItem instantiates a new MediaItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMediaItemWithDefaults

`func NewMediaItemWithDefaults() *MediaItem`

NewMediaItemWithDefaults instantiates a new MediaItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MediaItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MediaItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MediaItem) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MediaItem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUrl

`func (o *MediaItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MediaItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MediaItem) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *MediaItem) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTitle

`func (o *MediaItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MediaItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MediaItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MediaItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAltText

`func (o *MediaItem) GetAltText() string`

GetAltText returns the AltText field if non-nil, zero value otherwise.

### GetAltTextOk

`func (o *MediaItem) GetAltTextOk() (*string, bool)`

GetAltTextOk returns a tuple with the AltText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltText

`func (o *MediaItem) SetAltText(v string)`

SetAltText sets AltText field to given value.

### HasAltText

`func (o *MediaItem) HasAltText() bool`

HasAltText returns a boolean if a field has been set.

### GetFilename

`func (o *MediaItem) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *MediaItem) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *MediaItem) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *MediaItem) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetSize

`func (o *MediaItem) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *MediaItem) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *MediaItem) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *MediaItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMimeType

`func (o *MediaItem) GetMimeType() string`

GetMimeType returns the MimeType field if non-nil, zero value otherwise.

### GetMimeTypeOk

`func (o *MediaItem) GetMimeTypeOk() (*string, bool)`

GetMimeTypeOk returns a tuple with the MimeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMimeType

`func (o *MediaItem) SetMimeType(v string)`

SetMimeType sets MimeType field to given value.

### HasMimeType

`func (o *MediaItem) HasMimeType() bool`

HasMimeType returns a boolean if a field has been set.

### GetThumbnail

`func (o *MediaItem) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *MediaItem) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *MediaItem) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *MediaItem) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### GetInstagramThumbnail

`func (o *MediaItem) GetInstagramThumbnail() string`

GetInstagramThumbnail returns the InstagramThumbnail field if non-nil, zero value otherwise.

### GetInstagramThumbnailOk

`func (o *MediaItem) GetInstagramThumbnailOk() (*string, bool)`

GetInstagramThumbnailOk returns a tuple with the InstagramThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramThumbnail

`func (o *MediaItem) SetInstagramThumbnail(v string)`

SetInstagramThumbnail sets InstagramThumbnail field to given value.

### HasInstagramThumbnail

`func (o *MediaItem) HasInstagramThumbnail() bool`

HasInstagramThumbnail returns a boolean if a field has been set.

### GetTiktokProcessed

`func (o *MediaItem) GetTiktokProcessed() bool`

GetTiktokProcessed returns the TiktokProcessed field if non-nil, zero value otherwise.

### GetTiktokProcessedOk

`func (o *MediaItem) GetTiktokProcessedOk() (*bool, bool)`

GetTiktokProcessedOk returns a tuple with the TiktokProcessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiktokProcessed

`func (o *MediaItem) SetTiktokProcessed(v bool)`

SetTiktokProcessed sets TiktokProcessed field to given value.

### HasTiktokProcessed

`func (o *MediaItem) HasTiktokProcessed() bool`

HasTiktokProcessed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


