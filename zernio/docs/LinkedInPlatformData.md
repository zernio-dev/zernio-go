# LinkedInPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentTitle** | Pointer to **string** | Title displayed on LinkedIn document (PDF/carousel) posts. Required by LinkedIn for document posts. If omitted, falls back to the media item title, then the filename. | [optional] 
**OrganizationUrn** | Pointer to **string** | Target LinkedIn Organization URN (e.g. \&quot;urn:li:organization:123456789\&quot;). If omitted, uses the default org. Use GET /v1/accounts/{id}/linkedin-organizations to list orgs. | [optional] 
**FirstComment** | Pointer to **string** | Optional first comment to add after the post is created | [optional] 
**DisableLinkPreview** | Pointer to **bool** | Set to true to disable automatic link previews for URLs in the post content (default is false) | [optional] 
**ReshareUrl** | Pointer to **string** | LinkedIn post link to repost (use the post&#39;s \&quot;Copy link to post\&quot; action), or a urn:li:share / urn:li:ugcPost / urn:li:groupPost URN. The published post becomes a quote-reshare: your content is shown as the commentary and the original post is embedded underneath (LinkedIn&#39;s \&quot;repost with your thoughts\&quot;). Mutually exclusive with media. Works on personal profiles and organization pages. | [optional] 
**GeoRestriction** | Pointer to [**GeoRestriction**](GeoRestriction.md) |  | [optional] 

## Methods

### NewLinkedInPlatformData

`func NewLinkedInPlatformData() *LinkedInPlatformData`

NewLinkedInPlatformData instantiates a new LinkedInPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedInPlatformDataWithDefaults

`func NewLinkedInPlatformDataWithDefaults() *LinkedInPlatformData`

NewLinkedInPlatformDataWithDefaults instantiates a new LinkedInPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentTitle

`func (o *LinkedInPlatformData) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *LinkedInPlatformData) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *LinkedInPlatformData) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *LinkedInPlatformData) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetOrganizationUrn

`func (o *LinkedInPlatformData) GetOrganizationUrn() string`

GetOrganizationUrn returns the OrganizationUrn field if non-nil, zero value otherwise.

### GetOrganizationUrnOk

`func (o *LinkedInPlatformData) GetOrganizationUrnOk() (*string, bool)`

GetOrganizationUrnOk returns a tuple with the OrganizationUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrn

`func (o *LinkedInPlatformData) SetOrganizationUrn(v string)`

SetOrganizationUrn sets OrganizationUrn field to given value.

### HasOrganizationUrn

`func (o *LinkedInPlatformData) HasOrganizationUrn() bool`

HasOrganizationUrn returns a boolean if a field has been set.

### GetFirstComment

`func (o *LinkedInPlatformData) GetFirstComment() string`

GetFirstComment returns the FirstComment field if non-nil, zero value otherwise.

### GetFirstCommentOk

`func (o *LinkedInPlatformData) GetFirstCommentOk() (*string, bool)`

GetFirstCommentOk returns a tuple with the FirstComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstComment

`func (o *LinkedInPlatformData) SetFirstComment(v string)`

SetFirstComment sets FirstComment field to given value.

### HasFirstComment

`func (o *LinkedInPlatformData) HasFirstComment() bool`

HasFirstComment returns a boolean if a field has been set.

### GetDisableLinkPreview

`func (o *LinkedInPlatformData) GetDisableLinkPreview() bool`

GetDisableLinkPreview returns the DisableLinkPreview field if non-nil, zero value otherwise.

### GetDisableLinkPreviewOk

`func (o *LinkedInPlatformData) GetDisableLinkPreviewOk() (*bool, bool)`

GetDisableLinkPreviewOk returns a tuple with the DisableLinkPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLinkPreview

`func (o *LinkedInPlatformData) SetDisableLinkPreview(v bool)`

SetDisableLinkPreview sets DisableLinkPreview field to given value.

### HasDisableLinkPreview

`func (o *LinkedInPlatformData) HasDisableLinkPreview() bool`

HasDisableLinkPreview returns a boolean if a field has been set.

### GetReshareUrl

`func (o *LinkedInPlatformData) GetReshareUrl() string`

GetReshareUrl returns the ReshareUrl field if non-nil, zero value otherwise.

### GetReshareUrlOk

`func (o *LinkedInPlatformData) GetReshareUrlOk() (*string, bool)`

GetReshareUrlOk returns a tuple with the ReshareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReshareUrl

`func (o *LinkedInPlatformData) SetReshareUrl(v string)`

SetReshareUrl sets ReshareUrl field to given value.

### HasReshareUrl

`func (o *LinkedInPlatformData) HasReshareUrl() bool`

HasReshareUrl returns a boolean if a field has been set.

### GetGeoRestriction

`func (o *LinkedInPlatformData) GetGeoRestriction() GeoRestriction`

GetGeoRestriction returns the GeoRestriction field if non-nil, zero value otherwise.

### GetGeoRestrictionOk

`func (o *LinkedInPlatformData) GetGeoRestrictionOk() (*GeoRestriction, bool)`

GetGeoRestrictionOk returns a tuple with the GeoRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestriction

`func (o *LinkedInPlatformData) SetGeoRestriction(v GeoRestriction)`

SetGeoRestriction sets GeoRestriction field to given value.

### HasGeoRestriction

`func (o *LinkedInPlatformData) HasGeoRestriction() bool`

HasGeoRestriction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


