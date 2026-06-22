# FacebookPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Draft** | Pointer to **bool** | When true, creates the post as an unpublished draft visible in Facebook Publishing Tools instead of publishing immediately. Supported for feed posts (text, link, image, video) and reels. Not supported for stories. Drafts expire after ~30 days. | [optional] [default to false]
**ContentType** | Pointer to **string** | Set to &#39;story&#39; for Page Stories (24h ephemeral) or &#39;reel&#39; for Reels (short vertical video). Defaults to feed post if omitted. | [optional] 
**Title** | Pointer to **string** | Reel title (only for contentType&#x3D;reel). Separate from the caption/content field. | [optional] 
**FirstComment** | Pointer to **string** | Optional first comment to post immediately after publishing (feed posts and reels, not stories). Skipped when draft is true. | [optional] 
**PageId** | Pointer to **string** | Target Facebook Page ID for multi-page posting. If omitted, uses the default page. Use GET /v1/accounts/{id}/facebook-page to list pages. | [optional] 
**GeoRestriction** | Pointer to [**GeoRestriction**](GeoRestriction.md) |  | [optional] 
**CarouselCards** | Pointer to [**[]FacebookPlatformDataCarouselCardsInner**](FacebookPlatformDataCarouselCardsInner.md) | Renders the post as a multi-link carousel (organic Page post). When set, mediaItems must be provided with the same length and all items must be images (no videos). Each cards[i] adds the click-through link and headline for the image at mediaItems[i]. Mutually exclusive with contentType&#x3D;story|reel. Facebook display truncates name at ~35 chars and description at ~30 chars; longer strings are accepted but get truncated on render.  | [optional] 
**CarouselLink** | Pointer to **string** | Optional top-level \&quot;See more\&quot; destination shown on the carousel end card. Defaults to the first card&#39;s link when omitted. Only used together with carouselCards.  | [optional] 

## Methods

### NewFacebookPlatformData

`func NewFacebookPlatformData() *FacebookPlatformData`

NewFacebookPlatformData instantiates a new FacebookPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFacebookPlatformDataWithDefaults

`func NewFacebookPlatformDataWithDefaults() *FacebookPlatformData`

NewFacebookPlatformDataWithDefaults instantiates a new FacebookPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraft

`func (o *FacebookPlatformData) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *FacebookPlatformData) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *FacebookPlatformData) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *FacebookPlatformData) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetContentType

`func (o *FacebookPlatformData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FacebookPlatformData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FacebookPlatformData) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FacebookPlatformData) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTitle

`func (o *FacebookPlatformData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FacebookPlatformData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FacebookPlatformData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FacebookPlatformData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstComment

`func (o *FacebookPlatformData) GetFirstComment() string`

GetFirstComment returns the FirstComment field if non-nil, zero value otherwise.

### GetFirstCommentOk

`func (o *FacebookPlatformData) GetFirstCommentOk() (*string, bool)`

GetFirstCommentOk returns a tuple with the FirstComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstComment

`func (o *FacebookPlatformData) SetFirstComment(v string)`

SetFirstComment sets FirstComment field to given value.

### HasFirstComment

`func (o *FacebookPlatformData) HasFirstComment() bool`

HasFirstComment returns a boolean if a field has been set.

### GetPageId

`func (o *FacebookPlatformData) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *FacebookPlatformData) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *FacebookPlatformData) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *FacebookPlatformData) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetGeoRestriction

`func (o *FacebookPlatformData) GetGeoRestriction() GeoRestriction`

GetGeoRestriction returns the GeoRestriction field if non-nil, zero value otherwise.

### GetGeoRestrictionOk

`func (o *FacebookPlatformData) GetGeoRestrictionOk() (*GeoRestriction, bool)`

GetGeoRestrictionOk returns a tuple with the GeoRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestriction

`func (o *FacebookPlatformData) SetGeoRestriction(v GeoRestriction)`

SetGeoRestriction sets GeoRestriction field to given value.

### HasGeoRestriction

`func (o *FacebookPlatformData) HasGeoRestriction() bool`

HasGeoRestriction returns a boolean if a field has been set.

### GetCarouselCards

`func (o *FacebookPlatformData) GetCarouselCards() []FacebookPlatformDataCarouselCardsInner`

GetCarouselCards returns the CarouselCards field if non-nil, zero value otherwise.

### GetCarouselCardsOk

`func (o *FacebookPlatformData) GetCarouselCardsOk() (*[]FacebookPlatformDataCarouselCardsInner, bool)`

GetCarouselCardsOk returns a tuple with the CarouselCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarouselCards

`func (o *FacebookPlatformData) SetCarouselCards(v []FacebookPlatformDataCarouselCardsInner)`

SetCarouselCards sets CarouselCards field to given value.

### HasCarouselCards

`func (o *FacebookPlatformData) HasCarouselCards() bool`

HasCarouselCards returns a boolean if a field has been set.

### GetCarouselLink

`func (o *FacebookPlatformData) GetCarouselLink() string`

GetCarouselLink returns the CarouselLink field if non-nil, zero value otherwise.

### GetCarouselLinkOk

`func (o *FacebookPlatformData) GetCarouselLinkOk() (*string, bool)`

GetCarouselLinkOk returns a tuple with the CarouselLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarouselLink

`func (o *FacebookPlatformData) SetCarouselLink(v string)`

SetCarouselLink sets CarouselLink field to given value.

### HasCarouselLink

`func (o *FacebookPlatformData) HasCarouselLink() bool`

HasCarouselLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


