# CreateStandaloneAdRequestCreativesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Exact name for this ad. Falls back to &#x60;&lt;name&gt; #N&#x60; (N &#x3D; 1-based position). | [optional] 
**Headline** | **string** |  | 
**Body** | **string** |  | 
**Description** | Pointer to **string** | Link description for this ad (link_data.description; video creatives: video_data.link_description). Falls back to the top-level &#x60;description&#x60;; when both are omitted Meta scrapes the destination URL&#39;s OG description. | [optional] 
**ImageUrl** | Pointer to **string** | Image creative. Mutually exclusive with &#x60;video&#x60;. | [optional] 
**Video** | Pointer to [**CreateStandaloneAdRequestVideo**](CreateStandaloneAdRequestVideo.md) |  | [optional] 
**LinkUrl** | **string** |  | 
**CallToAction** | **string** |  | 

## Methods

### NewCreateStandaloneAdRequestCreativesInner

`func NewCreateStandaloneAdRequestCreativesInner(headline string, body string, linkUrl string, callToAction string, ) *CreateStandaloneAdRequestCreativesInner`

NewCreateStandaloneAdRequestCreativesInner instantiates a new CreateStandaloneAdRequestCreativesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestCreativesInnerWithDefaults

`func NewCreateStandaloneAdRequestCreativesInnerWithDefaults() *CreateStandaloneAdRequestCreativesInner`

NewCreateStandaloneAdRequestCreativesInnerWithDefaults instantiates a new CreateStandaloneAdRequestCreativesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateStandaloneAdRequestCreativesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateStandaloneAdRequestCreativesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateStandaloneAdRequestCreativesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHeadline

`func (o *CreateStandaloneAdRequestCreativesInner) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *CreateStandaloneAdRequestCreativesInner) SetHeadline(v string)`

SetHeadline sets Headline field to given value.


### GetBody

`func (o *CreateStandaloneAdRequestCreativesInner) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateStandaloneAdRequestCreativesInner) SetBody(v string)`

SetBody sets Body field to given value.


### GetDescription

`func (o *CreateStandaloneAdRequestCreativesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateStandaloneAdRequestCreativesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateStandaloneAdRequestCreativesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateStandaloneAdRequestCreativesInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateStandaloneAdRequestCreativesInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateStandaloneAdRequestCreativesInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideo

`func (o *CreateStandaloneAdRequestCreativesInner) GetVideo() CreateStandaloneAdRequestVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetVideoOk() (*CreateStandaloneAdRequestVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreateStandaloneAdRequestCreativesInner) SetVideo(v CreateStandaloneAdRequestVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreateStandaloneAdRequestCreativesInner) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetLinkUrl

`func (o *CreateStandaloneAdRequestCreativesInner) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *CreateStandaloneAdRequestCreativesInner) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.


### GetCallToAction

`func (o *CreateStandaloneAdRequestCreativesInner) GetCallToAction() string`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *CreateStandaloneAdRequestCreativesInner) GetCallToActionOk() (*string, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *CreateStandaloneAdRequestCreativesInner) SetCallToAction(v string)`

SetCallToAction sets CallToAction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


