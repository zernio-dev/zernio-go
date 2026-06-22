# PinterestPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Pin title. Defaults to first line of content or \&quot;Pin\&quot;. Must be ≤ 100 characters. | [optional] 
**BoardId** | Pointer to **string** | Target Pinterest board ID. If omitted, the first available board is used. | [optional] 
**Link** | Pointer to **string** | Destination link (pin URL) | [optional] 
**CoverImageUrl** | Pointer to **string** | Optional cover image for video pins | [optional] 
**CoverImageKeyFrameTime** | Pointer to **int32** | Optional key frame time in seconds for derived video cover | [optional] 

## Methods

### NewPinterestPlatformData

`func NewPinterestPlatformData() *PinterestPlatformData`

NewPinterestPlatformData instantiates a new PinterestPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinterestPlatformDataWithDefaults

`func NewPinterestPlatformDataWithDefaults() *PinterestPlatformData`

NewPinterestPlatformDataWithDefaults instantiates a new PinterestPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *PinterestPlatformData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PinterestPlatformData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PinterestPlatformData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PinterestPlatformData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetBoardId

`func (o *PinterestPlatformData) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *PinterestPlatformData) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *PinterestPlatformData) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *PinterestPlatformData) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetLink

`func (o *PinterestPlatformData) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PinterestPlatformData) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PinterestPlatformData) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PinterestPlatformData) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *PinterestPlatformData) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *PinterestPlatformData) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *PinterestPlatformData) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *PinterestPlatformData) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetCoverImageKeyFrameTime

`func (o *PinterestPlatformData) GetCoverImageKeyFrameTime() int32`

GetCoverImageKeyFrameTime returns the CoverImageKeyFrameTime field if non-nil, zero value otherwise.

### GetCoverImageKeyFrameTimeOk

`func (o *PinterestPlatformData) GetCoverImageKeyFrameTimeOk() (*int32, bool)`

GetCoverImageKeyFrameTimeOk returns a tuple with the CoverImageKeyFrameTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageKeyFrameTime

`func (o *PinterestPlatformData) SetCoverImageKeyFrameTime(v int32)`

SetCoverImageKeyFrameTime sets CoverImageKeyFrameTime field to given value.

### HasCoverImageKeyFrameTime

`func (o *PinterestPlatformData) HasCoverImageKeyFrameTime() bool`

HasCoverImageKeyFrameTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


