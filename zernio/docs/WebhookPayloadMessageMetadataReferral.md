# WebhookPayloadMessageMetadataReferral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CtwaClid** | Pointer to **string** | Meta&#39;s GCLID-equivalent click identifier. | [optional] 
**SourceId** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**Headline** | Pointer to **string** |  | [optional] 
**Body** | Pointer to **string** |  | [optional] 
**MediaType** | Pointer to **string** |  | [optional] 
**ImageUrl** | Pointer to **string** |  | [optional] 
**VideoUrl** | Pointer to **string** |  | [optional] 
**ThumbnailUrl** | Pointer to **string** |  | [optional] 
**AdId** | Pointer to **string** | Facebook Messenger CTM / Instagram CTD only. The Meta ad ID the user clicked to start the conversation.  | [optional] 
**Ref** | Pointer to **string** | Optional &#x60;ref&#x60; parameter passed through from the Meta ad creative. Facebook Messenger CTM / Instagram CTD only.  | [optional] 
**Source** | Pointer to **string** | Meta-supplied source identifier (e.g. &#x60;ADS&#x60;). Facebook Messenger CTM / Instagram CTD only.  | [optional] 
**Type** | Pointer to **string** | Meta-supplied referral type (e.g. &#x60;OPEN_THREAD&#x60;). Facebook Messenger CTM / Instagram CTD only.  | [optional] 
**AdsContextData** | Pointer to [**WebhookPayloadMessageMetadataReferralAdsContextData**](WebhookPayloadMessageMetadataReferralAdsContextData.md) |  | [optional] 

## Methods

### NewWebhookPayloadMessageMetadataReferral

`func NewWebhookPayloadMessageMetadataReferral() *WebhookPayloadMessageMetadataReferral`

NewWebhookPayloadMessageMetadataReferral instantiates a new WebhookPayloadMessageMetadataReferral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageMetadataReferralWithDefaults

`func NewWebhookPayloadMessageMetadataReferralWithDefaults() *WebhookPayloadMessageMetadataReferral`

NewWebhookPayloadMessageMetadataReferralWithDefaults instantiates a new WebhookPayloadMessageMetadataReferral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCtwaClid

`func (o *WebhookPayloadMessageMetadataReferral) GetCtwaClid() string`

GetCtwaClid returns the CtwaClid field if non-nil, zero value otherwise.

### GetCtwaClidOk

`func (o *WebhookPayloadMessageMetadataReferral) GetCtwaClidOk() (*string, bool)`

GetCtwaClidOk returns a tuple with the CtwaClid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtwaClid

`func (o *WebhookPayloadMessageMetadataReferral) SetCtwaClid(v string)`

SetCtwaClid sets CtwaClid field to given value.

### HasCtwaClid

`func (o *WebhookPayloadMessageMetadataReferral) HasCtwaClid() bool`

HasCtwaClid returns a boolean if a field has been set.

### GetSourceId

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *WebhookPayloadMessageMetadataReferral) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.

### HasSourceId

`func (o *WebhookPayloadMessageMetadataReferral) HasSourceId() bool`

HasSourceId returns a boolean if a field has been set.

### GetSourceType

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *WebhookPayloadMessageMetadataReferral) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *WebhookPayloadMessageMetadataReferral) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *WebhookPayloadMessageMetadataReferral) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *WebhookPayloadMessageMetadataReferral) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetHeadline

`func (o *WebhookPayloadMessageMetadataReferral) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *WebhookPayloadMessageMetadataReferral) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *WebhookPayloadMessageMetadataReferral) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *WebhookPayloadMessageMetadataReferral) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetBody

`func (o *WebhookPayloadMessageMetadataReferral) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *WebhookPayloadMessageMetadataReferral) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *WebhookPayloadMessageMetadataReferral) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *WebhookPayloadMessageMetadataReferral) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetMediaType

`func (o *WebhookPayloadMessageMetadataReferral) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *WebhookPayloadMessageMetadataReferral) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *WebhookPayloadMessageMetadataReferral) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *WebhookPayloadMessageMetadataReferral) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetImageUrl

`func (o *WebhookPayloadMessageMetadataReferral) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *WebhookPayloadMessageMetadataReferral) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *WebhookPayloadMessageMetadataReferral) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *WebhookPayloadMessageMetadataReferral) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideoUrl

`func (o *WebhookPayloadMessageMetadataReferral) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *WebhookPayloadMessageMetadataReferral) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *WebhookPayloadMessageMetadataReferral) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *WebhookPayloadMessageMetadataReferral) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### GetThumbnailUrl

`func (o *WebhookPayloadMessageMetadataReferral) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *WebhookPayloadMessageMetadataReferral) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *WebhookPayloadMessageMetadataReferral) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *WebhookPayloadMessageMetadataReferral) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetAdId

`func (o *WebhookPayloadMessageMetadataReferral) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *WebhookPayloadMessageMetadataReferral) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *WebhookPayloadMessageMetadataReferral) SetAdId(v string)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *WebhookPayloadMessageMetadataReferral) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### GetRef

`func (o *WebhookPayloadMessageMetadataReferral) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *WebhookPayloadMessageMetadataReferral) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *WebhookPayloadMessageMetadataReferral) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *WebhookPayloadMessageMetadataReferral) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSource

`func (o *WebhookPayloadMessageMetadataReferral) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *WebhookPayloadMessageMetadataReferral) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *WebhookPayloadMessageMetadataReferral) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *WebhookPayloadMessageMetadataReferral) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *WebhookPayloadMessageMetadataReferral) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WebhookPayloadMessageMetadataReferral) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WebhookPayloadMessageMetadataReferral) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WebhookPayloadMessageMetadataReferral) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdsContextData

`func (o *WebhookPayloadMessageMetadataReferral) GetAdsContextData() WebhookPayloadMessageMetadataReferralAdsContextData`

GetAdsContextData returns the AdsContextData field if non-nil, zero value otherwise.

### GetAdsContextDataOk

`func (o *WebhookPayloadMessageMetadataReferral) GetAdsContextDataOk() (*WebhookPayloadMessageMetadataReferralAdsContextData, bool)`

GetAdsContextDataOk returns a tuple with the AdsContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsContextData

`func (o *WebhookPayloadMessageMetadataReferral) SetAdsContextData(v WebhookPayloadMessageMetadataReferralAdsContextData)`

SetAdsContextData sets AdsContextData field to given value.

### HasAdsContextData

`func (o *WebhookPayloadMessageMetadataReferral) HasAdsContextData() bool`

HasAdsContextData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


