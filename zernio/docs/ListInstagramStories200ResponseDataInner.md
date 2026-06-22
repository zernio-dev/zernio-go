# ListInstagramStories200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Instagram media ID of the story. | 
**MediaType** | Pointer to **NullableString** | IMAGE / VIDEO / CAROUSEL_ALBUM | [optional] 
**MediaProductType** | Pointer to **NullableString** | Always &#39;STORY&#39; for this endpoint. | [optional] 
**MediaUrl** | Pointer to **NullableString** | Direct media URL. Null if Meta flagged the story for copyright. URL expires when the story expires. | [optional] 
**Permalink** | Pointer to **NullableString** | Public Instagram permalink to the story (only viewable while live). | [optional] 
**ThumbnailUrl** | Pointer to **NullableString** | Thumbnail URL for video stories. | [optional] 
**Timestamp** | Pointer to **NullableTime** | When the story was posted. | [optional] 

## Methods

### NewListInstagramStories200ResponseDataInner

`func NewListInstagramStories200ResponseDataInner(id string, ) *ListInstagramStories200ResponseDataInner`

NewListInstagramStories200ResponseDataInner instantiates a new ListInstagramStories200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstagramStories200ResponseDataInnerWithDefaults

`func NewListInstagramStories200ResponseDataInnerWithDefaults() *ListInstagramStories200ResponseDataInner`

NewListInstagramStories200ResponseDataInnerWithDefaults instantiates a new ListInstagramStories200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstagramStories200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstagramStories200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstagramStories200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.


### GetMediaType

`func (o *ListInstagramStories200ResponseDataInner) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *ListInstagramStories200ResponseDataInner) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *ListInstagramStories200ResponseDataInner) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *ListInstagramStories200ResponseDataInner) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### SetMediaTypeNil

`func (o *ListInstagramStories200ResponseDataInner) SetMediaTypeNil(b bool)`

 SetMediaTypeNil sets the value for MediaType to be an explicit nil

### UnsetMediaType
`func (o *ListInstagramStories200ResponseDataInner) UnsetMediaType()`

UnsetMediaType ensures that no value is present for MediaType, not even an explicit nil
### GetMediaProductType

`func (o *ListInstagramStories200ResponseDataInner) GetMediaProductType() string`

GetMediaProductType returns the MediaProductType field if non-nil, zero value otherwise.

### GetMediaProductTypeOk

`func (o *ListInstagramStories200ResponseDataInner) GetMediaProductTypeOk() (*string, bool)`

GetMediaProductTypeOk returns a tuple with the MediaProductType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaProductType

`func (o *ListInstagramStories200ResponseDataInner) SetMediaProductType(v string)`

SetMediaProductType sets MediaProductType field to given value.

### HasMediaProductType

`func (o *ListInstagramStories200ResponseDataInner) HasMediaProductType() bool`

HasMediaProductType returns a boolean if a field has been set.

### SetMediaProductTypeNil

`func (o *ListInstagramStories200ResponseDataInner) SetMediaProductTypeNil(b bool)`

 SetMediaProductTypeNil sets the value for MediaProductType to be an explicit nil

### UnsetMediaProductType
`func (o *ListInstagramStories200ResponseDataInner) UnsetMediaProductType()`

UnsetMediaProductType ensures that no value is present for MediaProductType, not even an explicit nil
### GetMediaUrl

`func (o *ListInstagramStories200ResponseDataInner) GetMediaUrl() string`

GetMediaUrl returns the MediaUrl field if non-nil, zero value otherwise.

### GetMediaUrlOk

`func (o *ListInstagramStories200ResponseDataInner) GetMediaUrlOk() (*string, bool)`

GetMediaUrlOk returns a tuple with the MediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrl

`func (o *ListInstagramStories200ResponseDataInner) SetMediaUrl(v string)`

SetMediaUrl sets MediaUrl field to given value.

### HasMediaUrl

`func (o *ListInstagramStories200ResponseDataInner) HasMediaUrl() bool`

HasMediaUrl returns a boolean if a field has been set.

### SetMediaUrlNil

`func (o *ListInstagramStories200ResponseDataInner) SetMediaUrlNil(b bool)`

 SetMediaUrlNil sets the value for MediaUrl to be an explicit nil

### UnsetMediaUrl
`func (o *ListInstagramStories200ResponseDataInner) UnsetMediaUrl()`

UnsetMediaUrl ensures that no value is present for MediaUrl, not even an explicit nil
### GetPermalink

`func (o *ListInstagramStories200ResponseDataInner) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *ListInstagramStories200ResponseDataInner) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *ListInstagramStories200ResponseDataInner) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *ListInstagramStories200ResponseDataInner) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### SetPermalinkNil

`func (o *ListInstagramStories200ResponseDataInner) SetPermalinkNil(b bool)`

 SetPermalinkNil sets the value for Permalink to be an explicit nil

### UnsetPermalink
`func (o *ListInstagramStories200ResponseDataInner) UnsetPermalink()`

UnsetPermalink ensures that no value is present for Permalink, not even an explicit nil
### GetThumbnailUrl

`func (o *ListInstagramStories200ResponseDataInner) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *ListInstagramStories200ResponseDataInner) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *ListInstagramStories200ResponseDataInner) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *ListInstagramStories200ResponseDataInner) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *ListInstagramStories200ResponseDataInner) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *ListInstagramStories200ResponseDataInner) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetTimestamp

`func (o *ListInstagramStories200ResponseDataInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListInstagramStories200ResponseDataInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListInstagramStories200ResponseDataInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ListInstagramStories200ResponseDataInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### SetTimestampNil

`func (o *ListInstagramStories200ResponseDataInner) SetTimestampNil(b bool)`

 SetTimestampNil sets the value for Timestamp to be an explicit nil

### UnsetTimestamp
`func (o *ListInstagramStories200ResponseDataInner) UnsetTimestamp()`

UnsetTimestamp ensures that no value is present for Timestamp, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


