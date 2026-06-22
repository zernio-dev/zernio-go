# ExternalPostWebhookPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Platform-native post ID (NOT a Zernio post ID). | 
**Platform** | **string** | Platform the post lives on (e.g. \&quot;googlebusiness\&quot;). | 
**AccountId** | **string** | Zernio social account ID the post belongs to. | 
**Url** | **NullableString** | Direct URL to the post on the platform, when available. | 
**Content** | **string** | Post text. May be empty. | 
**MediaType** | **string** | One of image, video, gif, document, text, carousel. | 
**MediaItems** | [**[]ExternalPostMediaItem**](ExternalPostMediaItem.md) |  | 
**ThumbnailUrl** | **NullableString** |  | 
**PublishedAt** | **time.Time** |  | 
**Source** | **string** | Always \&quot;external\&quot; — distinguishes these from Zernio-originated post.* events. | 
**DeletedAt** | Pointer to **NullableTime** | Detection time of deletion. Present on post.external.deleted; null/absent otherwise. | [optional] 

## Methods

### NewExternalPostWebhookPost

`func NewExternalPostWebhookPost(id string, platform string, accountId string, url NullableString, content string, mediaType string, mediaItems []ExternalPostMediaItem, thumbnailUrl NullableString, publishedAt time.Time, source string, ) *ExternalPostWebhookPost`

NewExternalPostWebhookPost instantiates a new ExternalPostWebhookPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPostWebhookPostWithDefaults

`func NewExternalPostWebhookPostWithDefaults() *ExternalPostWebhookPost`

NewExternalPostWebhookPostWithDefaults instantiates a new ExternalPostWebhookPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalPostWebhookPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalPostWebhookPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalPostWebhookPost) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ExternalPostWebhookPost) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ExternalPostWebhookPost) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ExternalPostWebhookPost) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAccountId

`func (o *ExternalPostWebhookPost) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ExternalPostWebhookPost) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ExternalPostWebhookPost) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetUrl

`func (o *ExternalPostWebhookPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExternalPostWebhookPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExternalPostWebhookPost) SetUrl(v string)`

SetUrl sets Url field to given value.


### SetUrlNil

`func (o *ExternalPostWebhookPost) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *ExternalPostWebhookPost) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetContent

`func (o *ExternalPostWebhookPost) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ExternalPostWebhookPost) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ExternalPostWebhookPost) SetContent(v string)`

SetContent sets Content field to given value.


### GetMediaType

`func (o *ExternalPostWebhookPost) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ExternalPostWebhookPost) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ExternalPostWebhookPost) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.


### GetMediaItems

`func (o *ExternalPostWebhookPost) GetMediaItems() []ExternalPostMediaItem`

GetMediaItems returns the MediaItems field if non-nil, zero value otherwise.

### GetMediaItemsOk

`func (o *ExternalPostWebhookPost) GetMediaItemsOk() (*[]ExternalPostMediaItem, bool)`

GetMediaItemsOk returns a tuple with the MediaItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaItems

`func (o *ExternalPostWebhookPost) SetMediaItems(v []ExternalPostMediaItem)`

SetMediaItems sets MediaItems field to given value.


### GetThumbnailUrl

`func (o *ExternalPostWebhookPost) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ExternalPostWebhookPost) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ExternalPostWebhookPost) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.


### SetThumbnailUrlNil

`func (o *ExternalPostWebhookPost) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *ExternalPostWebhookPost) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetPublishedAt

`func (o *ExternalPostWebhookPost) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *ExternalPostWebhookPost) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *ExternalPostWebhookPost) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.


### GetSource

`func (o *ExternalPostWebhookPost) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ExternalPostWebhookPost) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ExternalPostWebhookPost) SetSource(v string)`

SetSource sets Source field to given value.


### GetDeletedAt

`func (o *ExternalPostWebhookPost) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ExternalPostWebhookPost) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ExternalPostWebhookPost) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ExternalPostWebhookPost) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ExternalPostWebhookPost) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ExternalPostWebhookPost) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


