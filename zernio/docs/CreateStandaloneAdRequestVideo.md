# CreateStandaloneAdRequestVideo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | Public URL of the video. Meta: uploaded via chunked transfer on /act_X/advideos, then the request blocks on Meta&#39;s transcoding until status.video_status &#x3D;&#x3D;&#x3D; &#39;ready&#39;. LinkedIn: uploaded via the Videos API (multipart), then the request blocks until LinkedIn finishes transcoding (status AVAILABLE) — short clips take ~10-30s. | 
**ThumbnailUrl** | Pointer to **string** | Public URL of a still-image thumbnail for the video. OPTIONAL: when omitted on Meta, the poster is auto-generated from Meta&#39;s own preferred video thumbnail (the same candidates Ads Manager shows), so video ads publish without supplying one. Provide it to control the poster frame exactly (uploaded as an ad image and referenced in object_story_spec.video_data). Ignored by LinkedIn (auto-generated poster frame). | [optional] 

## Methods

### NewCreateStandaloneAdRequestVideo

`func NewCreateStandaloneAdRequestVideo(url string, ) *CreateStandaloneAdRequestVideo`

NewCreateStandaloneAdRequestVideo instantiates a new CreateStandaloneAdRequestVideo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestVideoWithDefaults

`func NewCreateStandaloneAdRequestVideoWithDefaults() *CreateStandaloneAdRequestVideo`

NewCreateStandaloneAdRequestVideoWithDefaults instantiates a new CreateStandaloneAdRequestVideo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateStandaloneAdRequestVideo) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateStandaloneAdRequestVideo) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateStandaloneAdRequestVideo) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetThumbnailUrl

`func (o *CreateStandaloneAdRequestVideo) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *CreateStandaloneAdRequestVideo) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *CreateStandaloneAdRequestVideo) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *CreateStandaloneAdRequestVideo) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


