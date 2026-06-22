# GetAdTrackingTags200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**UrlTags** | Pointer to **NullableString** | Meta: &amp;-joined click-URL params. | [optional] 
**TemplateUrlSpec** | Pointer to **map[string]interface{}** | Meta: third-party click-tracking template (Dynamic Ads). | [optional] 
**TrackingUrlTemplate** | Pointer to **NullableString** | Google. | [optional] 
**FinalUrlSuffix** | Pointer to **NullableString** | Google. | [optional] 
**DynamicValueParameters** | Pointer to **map[string]interface{}** | LinkedIn. | [optional] 
**CustomValueParameters** | Pointer to **map[string]interface{}** | LinkedIn. | [optional] 

## Methods

### NewGetAdTrackingTags200Response

`func NewGetAdTrackingTags200Response() *GetAdTrackingTags200Response`

NewGetAdTrackingTags200Response instantiates a new GetAdTrackingTags200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdTrackingTags200ResponseWithDefaults

`func NewGetAdTrackingTags200ResponseWithDefaults() *GetAdTrackingTags200Response`

NewGetAdTrackingTags200ResponseWithDefaults instantiates a new GetAdTrackingTags200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetAdTrackingTags200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAdTrackingTags200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAdTrackingTags200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetAdTrackingTags200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetLevel

`func (o *GetAdTrackingTags200Response) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GetAdTrackingTags200Response) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GetAdTrackingTags200Response) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GetAdTrackingTags200Response) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetUrlTags

`func (o *GetAdTrackingTags200Response) GetUrlTags() string`

GetUrlTags returns the UrlTags field if non-nil, zero value otherwise.

### GetUrlTagsOk

`func (o *GetAdTrackingTags200Response) GetUrlTagsOk() (*string, bool)`

GetUrlTagsOk returns a tuple with the UrlTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTags

`func (o *GetAdTrackingTags200Response) SetUrlTags(v string)`

SetUrlTags sets UrlTags field to given value.

### HasUrlTags

`func (o *GetAdTrackingTags200Response) HasUrlTags() bool`

HasUrlTags returns a boolean if a field has been set.

### SetUrlTagsNil

`func (o *GetAdTrackingTags200Response) SetUrlTagsNil(b bool)`

 SetUrlTagsNil sets the value for UrlTags to be an explicit nil

### UnsetUrlTags
`func (o *GetAdTrackingTags200Response) UnsetUrlTags()`

UnsetUrlTags ensures that no value is present for UrlTags, not even an explicit nil
### GetTemplateUrlSpec

`func (o *GetAdTrackingTags200Response) GetTemplateUrlSpec() map[string]interface{}`

GetTemplateUrlSpec returns the TemplateUrlSpec field if non-nil, zero value otherwise.

### GetTemplateUrlSpecOk

`func (o *GetAdTrackingTags200Response) GetTemplateUrlSpecOk() (*map[string]interface{}, bool)`

GetTemplateUrlSpecOk returns a tuple with the TemplateUrlSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateUrlSpec

`func (o *GetAdTrackingTags200Response) SetTemplateUrlSpec(v map[string]interface{})`

SetTemplateUrlSpec sets TemplateUrlSpec field to given value.

### HasTemplateUrlSpec

`func (o *GetAdTrackingTags200Response) HasTemplateUrlSpec() bool`

HasTemplateUrlSpec returns a boolean if a field has been set.

### SetTemplateUrlSpecNil

`func (o *GetAdTrackingTags200Response) SetTemplateUrlSpecNil(b bool)`

 SetTemplateUrlSpecNil sets the value for TemplateUrlSpec to be an explicit nil

### UnsetTemplateUrlSpec
`func (o *GetAdTrackingTags200Response) UnsetTemplateUrlSpec()`

UnsetTemplateUrlSpec ensures that no value is present for TemplateUrlSpec, not even an explicit nil
### GetTrackingUrlTemplate

`func (o *GetAdTrackingTags200Response) GetTrackingUrlTemplate() string`

GetTrackingUrlTemplate returns the TrackingUrlTemplate field if non-nil, zero value otherwise.

### GetTrackingUrlTemplateOk

`func (o *GetAdTrackingTags200Response) GetTrackingUrlTemplateOk() (*string, bool)`

GetTrackingUrlTemplateOk returns a tuple with the TrackingUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrlTemplate

`func (o *GetAdTrackingTags200Response) SetTrackingUrlTemplate(v string)`

SetTrackingUrlTemplate sets TrackingUrlTemplate field to given value.

### HasTrackingUrlTemplate

`func (o *GetAdTrackingTags200Response) HasTrackingUrlTemplate() bool`

HasTrackingUrlTemplate returns a boolean if a field has been set.

### SetTrackingUrlTemplateNil

`func (o *GetAdTrackingTags200Response) SetTrackingUrlTemplateNil(b bool)`

 SetTrackingUrlTemplateNil sets the value for TrackingUrlTemplate to be an explicit nil

### UnsetTrackingUrlTemplate
`func (o *GetAdTrackingTags200Response) UnsetTrackingUrlTemplate()`

UnsetTrackingUrlTemplate ensures that no value is present for TrackingUrlTemplate, not even an explicit nil
### GetFinalUrlSuffix

`func (o *GetAdTrackingTags200Response) GetFinalUrlSuffix() string`

GetFinalUrlSuffix returns the FinalUrlSuffix field if non-nil, zero value otherwise.

### GetFinalUrlSuffixOk

`func (o *GetAdTrackingTags200Response) GetFinalUrlSuffixOk() (*string, bool)`

GetFinalUrlSuffixOk returns a tuple with the FinalUrlSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalUrlSuffix

`func (o *GetAdTrackingTags200Response) SetFinalUrlSuffix(v string)`

SetFinalUrlSuffix sets FinalUrlSuffix field to given value.

### HasFinalUrlSuffix

`func (o *GetAdTrackingTags200Response) HasFinalUrlSuffix() bool`

HasFinalUrlSuffix returns a boolean if a field has been set.

### SetFinalUrlSuffixNil

`func (o *GetAdTrackingTags200Response) SetFinalUrlSuffixNil(b bool)`

 SetFinalUrlSuffixNil sets the value for FinalUrlSuffix to be an explicit nil

### UnsetFinalUrlSuffix
`func (o *GetAdTrackingTags200Response) UnsetFinalUrlSuffix()`

UnsetFinalUrlSuffix ensures that no value is present for FinalUrlSuffix, not even an explicit nil
### GetDynamicValueParameters

`func (o *GetAdTrackingTags200Response) GetDynamicValueParameters() map[string]interface{}`

GetDynamicValueParameters returns the DynamicValueParameters field if non-nil, zero value otherwise.

### GetDynamicValueParametersOk

`func (o *GetAdTrackingTags200Response) GetDynamicValueParametersOk() (*map[string]interface{}, bool)`

GetDynamicValueParametersOk returns a tuple with the DynamicValueParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicValueParameters

`func (o *GetAdTrackingTags200Response) SetDynamicValueParameters(v map[string]interface{})`

SetDynamicValueParameters sets DynamicValueParameters field to given value.

### HasDynamicValueParameters

`func (o *GetAdTrackingTags200Response) HasDynamicValueParameters() bool`

HasDynamicValueParameters returns a boolean if a field has been set.

### SetDynamicValueParametersNil

`func (o *GetAdTrackingTags200Response) SetDynamicValueParametersNil(b bool)`

 SetDynamicValueParametersNil sets the value for DynamicValueParameters to be an explicit nil

### UnsetDynamicValueParameters
`func (o *GetAdTrackingTags200Response) UnsetDynamicValueParameters()`

UnsetDynamicValueParameters ensures that no value is present for DynamicValueParameters, not even an explicit nil
### GetCustomValueParameters

`func (o *GetAdTrackingTags200Response) GetCustomValueParameters() map[string]interface{}`

GetCustomValueParameters returns the CustomValueParameters field if non-nil, zero value otherwise.

### GetCustomValueParametersOk

`func (o *GetAdTrackingTags200Response) GetCustomValueParametersOk() (*map[string]interface{}, bool)`

GetCustomValueParametersOk returns a tuple with the CustomValueParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValueParameters

`func (o *GetAdTrackingTags200Response) SetCustomValueParameters(v map[string]interface{})`

SetCustomValueParameters sets CustomValueParameters field to given value.

### HasCustomValueParameters

`func (o *GetAdTrackingTags200Response) HasCustomValueParameters() bool`

HasCustomValueParameters returns a boolean if a field has been set.

### SetCustomValueParametersNil

`func (o *GetAdTrackingTags200Response) SetCustomValueParametersNil(b bool)`

 SetCustomValueParametersNil sets the value for CustomValueParameters to be an explicit nil

### UnsetCustomValueParameters
`func (o *GetAdTrackingTags200Response) UnsetCustomValueParameters()`

UnsetCustomValueParameters ensures that no value is present for CustomValueParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


