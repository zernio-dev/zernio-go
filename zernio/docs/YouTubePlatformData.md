# YouTubePlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** | Video title. Defaults to first line of content or \&quot;Untitled Video\&quot;. Must be ≤ 100 characters. | [optional] 
**Visibility** | Pointer to **string** | Video visibility: public (default, anyone can watch), unlisted (link only), private (invite only) | [optional] [default to "public"]
**MadeForKids** | Pointer to **bool** | COPPA compliance flag. Set true for child-directed content (restricts comments, notifications, ad targeting). Defaults to false. YouTube may block views if not explicitly set. | [optional] [default to false]
**FirstComment** | Pointer to **string** | Optional first comment to post immediately after video upload. Up to 10,000 characters (YouTube&#39;s comment limit). | [optional] 
**ContainsSyntheticMedia** | Pointer to **bool** | AI-generated content disclosure. Set true if the video contains synthetic content that could be mistaken for real. YouTube may add a label. | [optional] [default to false]
**CategoryId** | Pointer to **string** | YouTube video category ID. Defaults to 22 (People &amp; Blogs). Common: 1 (Film), 2 (Autos), 10 (Music), 15 (Pets), 17 (Sports), 20 (Gaming), 23 (Comedy), 24 (Entertainment), 25 (News), 26 (Howto), 27 (Education), 28 (Science &amp; Tech). | [optional] [default to "22"]
**PlaylistId** | Pointer to **string** | Optional YouTube playlist ID to add the video to after upload (e.g. &#39;PLxxxxxxxxxxxxx&#39;). Use GET /v1/accounts/{id}/youtube-playlists to list available playlists. Works for both immediate and scheduled uploads. Quota cost: 50 YouTube API units per call. | [optional] 

## Methods

### NewYouTubePlatformData

`func NewYouTubePlatformData() *YouTubePlatformData`

NewYouTubePlatformData instantiates a new YouTubePlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubePlatformDataWithDefaults

`func NewYouTubePlatformDataWithDefaults() *YouTubePlatformData`

NewYouTubePlatformDataWithDefaults instantiates a new YouTubePlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *YouTubePlatformData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *YouTubePlatformData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *YouTubePlatformData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *YouTubePlatformData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisibility

`func (o *YouTubePlatformData) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *YouTubePlatformData) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *YouTubePlatformData) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *YouTubePlatformData) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMadeForKids

`func (o *YouTubePlatformData) GetMadeForKids() bool`

GetMadeForKids returns the MadeForKids field if non-nil, zero value otherwise.

### GetMadeForKidsOk

`func (o *YouTubePlatformData) GetMadeForKidsOk() (*bool, bool)`

GetMadeForKidsOk returns a tuple with the MadeForKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadeForKids

`func (o *YouTubePlatformData) SetMadeForKids(v bool)`

SetMadeForKids sets MadeForKids field to given value.

### HasMadeForKids

`func (o *YouTubePlatformData) HasMadeForKids() bool`

HasMadeForKids returns a boolean if a field has been set.

### GetFirstComment

`func (o *YouTubePlatformData) GetFirstComment() string`

GetFirstComment returns the FirstComment field if non-nil, zero value otherwise.

### GetFirstCommentOk

`func (o *YouTubePlatformData) GetFirstCommentOk() (*string, bool)`

GetFirstCommentOk returns a tuple with the FirstComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstComment

`func (o *YouTubePlatformData) SetFirstComment(v string)`

SetFirstComment sets FirstComment field to given value.

### HasFirstComment

`func (o *YouTubePlatformData) HasFirstComment() bool`

HasFirstComment returns a boolean if a field has been set.

### GetContainsSyntheticMedia

`func (o *YouTubePlatformData) GetContainsSyntheticMedia() bool`

GetContainsSyntheticMedia returns the ContainsSyntheticMedia field if non-nil, zero value otherwise.

### GetContainsSyntheticMediaOk

`func (o *YouTubePlatformData) GetContainsSyntheticMediaOk() (*bool, bool)`

GetContainsSyntheticMediaOk returns a tuple with the ContainsSyntheticMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsSyntheticMedia

`func (o *YouTubePlatformData) SetContainsSyntheticMedia(v bool)`

SetContainsSyntheticMedia sets ContainsSyntheticMedia field to given value.

### HasContainsSyntheticMedia

`func (o *YouTubePlatformData) HasContainsSyntheticMedia() bool`

HasContainsSyntheticMedia returns a boolean if a field has been set.

### GetCategoryId

`func (o *YouTubePlatformData) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *YouTubePlatformData) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *YouTubePlatformData) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *YouTubePlatformData) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPlaylistId

`func (o *YouTubePlatformData) GetPlaylistId() string`

GetPlaylistId returns the PlaylistId field if non-nil, zero value otherwise.

### GetPlaylistIdOk

`func (o *YouTubePlatformData) GetPlaylistIdOk() (*string, bool)`

GetPlaylistIdOk returns a tuple with the PlaylistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistId

`func (o *YouTubePlatformData) SetPlaylistId(v string)`

SetPlaylistId sets PlaylistId field to given value.

### HasPlaylistId

`func (o *YouTubePlatformData) HasPlaylistId() bool`

HasPlaylistId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


