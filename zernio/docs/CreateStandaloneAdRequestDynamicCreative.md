# CreateStandaloneAdRequestDynamicCreative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageUrls** | **[]string** | Pool of image URLs (1-10). Uploaded to the ad account and referenced by hash in the asset feed. | 
**Bodies** | Pointer to **[]string** | Primary-text variations (the body copy). | [optional] 
**Titles** | Pointer to **[]string** | Headline variations. | [optional] 
**Descriptions** | Pointer to **[]string** | Description (link caption) variations. | [optional] 
**LinkUrls** | Pointer to **[]string** | Destination URL variations. At least one is required unless &#x60;goal&#x60; is &#x60;lead_generation&#x60;. | [optional] 
**CallToActionTypes** | Pointer to **[]string** | CTA-button variations. Required. | [optional] 
**AdFormat** | Pointer to **string** | Asset-feed ad format. Defaults to SINGLE_IMAGE. | [optional] [default to "SINGLE_IMAGE"]

## Methods

### NewCreateStandaloneAdRequestDynamicCreative

`func NewCreateStandaloneAdRequestDynamicCreative(imageUrls []string, ) *CreateStandaloneAdRequestDynamicCreative`

NewCreateStandaloneAdRequestDynamicCreative instantiates a new CreateStandaloneAdRequestDynamicCreative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestDynamicCreativeWithDefaults

`func NewCreateStandaloneAdRequestDynamicCreativeWithDefaults() *CreateStandaloneAdRequestDynamicCreative`

NewCreateStandaloneAdRequestDynamicCreativeWithDefaults instantiates a new CreateStandaloneAdRequestDynamicCreative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageUrls

`func (o *CreateStandaloneAdRequestDynamicCreative) GetImageUrls() []string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetImageUrlsOk() (*[]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *CreateStandaloneAdRequestDynamicCreative) SetImageUrls(v []string)`

SetImageUrls sets ImageUrls field to given value.


### GetBodies

`func (o *CreateStandaloneAdRequestDynamicCreative) GetBodies() []string`

GetBodies returns the Bodies field if non-nil, zero value otherwise.

### GetBodiesOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetBodiesOk() (*[]string, bool)`

GetBodiesOk returns a tuple with the Bodies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodies

`func (o *CreateStandaloneAdRequestDynamicCreative) SetBodies(v []string)`

SetBodies sets Bodies field to given value.

### HasBodies

`func (o *CreateStandaloneAdRequestDynamicCreative) HasBodies() bool`

HasBodies returns a boolean if a field has been set.

### GetTitles

`func (o *CreateStandaloneAdRequestDynamicCreative) GetTitles() []string`

GetTitles returns the Titles field if non-nil, zero value otherwise.

### GetTitlesOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetTitlesOk() (*[]string, bool)`

GetTitlesOk returns a tuple with the Titles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitles

`func (o *CreateStandaloneAdRequestDynamicCreative) SetTitles(v []string)`

SetTitles sets Titles field to given value.

### HasTitles

`func (o *CreateStandaloneAdRequestDynamicCreative) HasTitles() bool`

HasTitles returns a boolean if a field has been set.

### GetDescriptions

`func (o *CreateStandaloneAdRequestDynamicCreative) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *CreateStandaloneAdRequestDynamicCreative) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.

### HasDescriptions

`func (o *CreateStandaloneAdRequestDynamicCreative) HasDescriptions() bool`

HasDescriptions returns a boolean if a field has been set.

### GetLinkUrls

`func (o *CreateStandaloneAdRequestDynamicCreative) GetLinkUrls() []string`

GetLinkUrls returns the LinkUrls field if non-nil, zero value otherwise.

### GetLinkUrlsOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetLinkUrlsOk() (*[]string, bool)`

GetLinkUrlsOk returns a tuple with the LinkUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrls

`func (o *CreateStandaloneAdRequestDynamicCreative) SetLinkUrls(v []string)`

SetLinkUrls sets LinkUrls field to given value.

### HasLinkUrls

`func (o *CreateStandaloneAdRequestDynamicCreative) HasLinkUrls() bool`

HasLinkUrls returns a boolean if a field has been set.

### GetCallToActionTypes

`func (o *CreateStandaloneAdRequestDynamicCreative) GetCallToActionTypes() []string`

GetCallToActionTypes returns the CallToActionTypes field if non-nil, zero value otherwise.

### GetCallToActionTypesOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetCallToActionTypesOk() (*[]string, bool)`

GetCallToActionTypesOk returns a tuple with the CallToActionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToActionTypes

`func (o *CreateStandaloneAdRequestDynamicCreative) SetCallToActionTypes(v []string)`

SetCallToActionTypes sets CallToActionTypes field to given value.

### HasCallToActionTypes

`func (o *CreateStandaloneAdRequestDynamicCreative) HasCallToActionTypes() bool`

HasCallToActionTypes returns a boolean if a field has been set.

### GetAdFormat

`func (o *CreateStandaloneAdRequestDynamicCreative) GetAdFormat() string`

GetAdFormat returns the AdFormat field if non-nil, zero value otherwise.

### GetAdFormatOk

`func (o *CreateStandaloneAdRequestDynamicCreative) GetAdFormatOk() (*string, bool)`

GetAdFormatOk returns a tuple with the AdFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdFormat

`func (o *CreateStandaloneAdRequestDynamicCreative) SetAdFormat(v string)`

SetAdFormat sets AdFormat field to given value.

### HasAdFormat

`func (o *CreateStandaloneAdRequestDynamicCreative) HasAdFormat() bool`

HasAdFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


