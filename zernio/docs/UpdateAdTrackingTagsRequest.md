# UpdateAdTrackingTagsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UrlTags** | Pointer to [**[]UpdateAdTrackingTagsRequestUrlTagsInner**](UpdateAdTrackingTagsRequestUrlTagsInner.md) | Meta only. Click-URL params appended to a freshly-rebuilt creative. | [optional] 
**Creative** | Pointer to [**UpdateAdTrackingTagsRequestCreative**](UpdateAdTrackingTagsRequestCreative.md) |  | [optional] 
**TrackingUrlTemplate** | Pointer to **string** | Google only. Full tracking template (must contain {lpurl}). | [optional] 
**FinalUrlSuffix** | Pointer to **string** | Google only. Parse-only key&#x3D;value params. | [optional] 
**DynamicValueParameters** | Pointer to **map[string]string** | LinkedIn only. key -&gt; dynamic value enum (CAMPAIGN_ID, CAMPAIGN_NAME, CREATIVE_ID, ...). | [optional] 
**CustomValueParameters** | Pointer to **map[string]string** | LinkedIn only. key -&gt; static value. | [optional] 

## Methods

### NewUpdateAdTrackingTagsRequest

`func NewUpdateAdTrackingTagsRequest() *UpdateAdTrackingTagsRequest`

NewUpdateAdTrackingTagsRequest instantiates a new UpdateAdTrackingTagsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdTrackingTagsRequestWithDefaults

`func NewUpdateAdTrackingTagsRequestWithDefaults() *UpdateAdTrackingTagsRequest`

NewUpdateAdTrackingTagsRequestWithDefaults instantiates a new UpdateAdTrackingTagsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrlTags

`func (o *UpdateAdTrackingTagsRequest) GetUrlTags() []UpdateAdTrackingTagsRequestUrlTagsInner`

GetUrlTags returns the UrlTags field if non-nil, zero value otherwise.

### GetUrlTagsOk

`func (o *UpdateAdTrackingTagsRequest) GetUrlTagsOk() (*[]UpdateAdTrackingTagsRequestUrlTagsInner, bool)`

GetUrlTagsOk returns a tuple with the UrlTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlTags

`func (o *UpdateAdTrackingTagsRequest) SetUrlTags(v []UpdateAdTrackingTagsRequestUrlTagsInner)`

SetUrlTags sets UrlTags field to given value.

### HasUrlTags

`func (o *UpdateAdTrackingTagsRequest) HasUrlTags() bool`

HasUrlTags returns a boolean if a field has been set.

### GetCreative

`func (o *UpdateAdTrackingTagsRequest) GetCreative() UpdateAdTrackingTagsRequestCreative`

GetCreative returns the Creative field if non-nil, zero value otherwise.

### GetCreativeOk

`func (o *UpdateAdTrackingTagsRequest) GetCreativeOk() (*UpdateAdTrackingTagsRequestCreative, bool)`

GetCreativeOk returns a tuple with the Creative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreative

`func (o *UpdateAdTrackingTagsRequest) SetCreative(v UpdateAdTrackingTagsRequestCreative)`

SetCreative sets Creative field to given value.

### HasCreative

`func (o *UpdateAdTrackingTagsRequest) HasCreative() bool`

HasCreative returns a boolean if a field has been set.

### GetTrackingUrlTemplate

`func (o *UpdateAdTrackingTagsRequest) GetTrackingUrlTemplate() string`

GetTrackingUrlTemplate returns the TrackingUrlTemplate field if non-nil, zero value otherwise.

### GetTrackingUrlTemplateOk

`func (o *UpdateAdTrackingTagsRequest) GetTrackingUrlTemplateOk() (*string, bool)`

GetTrackingUrlTemplateOk returns a tuple with the TrackingUrlTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrlTemplate

`func (o *UpdateAdTrackingTagsRequest) SetTrackingUrlTemplate(v string)`

SetTrackingUrlTemplate sets TrackingUrlTemplate field to given value.

### HasTrackingUrlTemplate

`func (o *UpdateAdTrackingTagsRequest) HasTrackingUrlTemplate() bool`

HasTrackingUrlTemplate returns a boolean if a field has been set.

### GetFinalUrlSuffix

`func (o *UpdateAdTrackingTagsRequest) GetFinalUrlSuffix() string`

GetFinalUrlSuffix returns the FinalUrlSuffix field if non-nil, zero value otherwise.

### GetFinalUrlSuffixOk

`func (o *UpdateAdTrackingTagsRequest) GetFinalUrlSuffixOk() (*string, bool)`

GetFinalUrlSuffixOk returns a tuple with the FinalUrlSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalUrlSuffix

`func (o *UpdateAdTrackingTagsRequest) SetFinalUrlSuffix(v string)`

SetFinalUrlSuffix sets FinalUrlSuffix field to given value.

### HasFinalUrlSuffix

`func (o *UpdateAdTrackingTagsRequest) HasFinalUrlSuffix() bool`

HasFinalUrlSuffix returns a boolean if a field has been set.

### GetDynamicValueParameters

`func (o *UpdateAdTrackingTagsRequest) GetDynamicValueParameters() map[string]string`

GetDynamicValueParameters returns the DynamicValueParameters field if non-nil, zero value otherwise.

### GetDynamicValueParametersOk

`func (o *UpdateAdTrackingTagsRequest) GetDynamicValueParametersOk() (*map[string]string, bool)`

GetDynamicValueParametersOk returns a tuple with the DynamicValueParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicValueParameters

`func (o *UpdateAdTrackingTagsRequest) SetDynamicValueParameters(v map[string]string)`

SetDynamicValueParameters sets DynamicValueParameters field to given value.

### HasDynamicValueParameters

`func (o *UpdateAdTrackingTagsRequest) HasDynamicValueParameters() bool`

HasDynamicValueParameters returns a boolean if a field has been set.

### GetCustomValueParameters

`func (o *UpdateAdTrackingTagsRequest) GetCustomValueParameters() map[string]string`

GetCustomValueParameters returns the CustomValueParameters field if non-nil, zero value otherwise.

### GetCustomValueParametersOk

`func (o *UpdateAdTrackingTagsRequest) GetCustomValueParametersOk() (*map[string]string, bool)`

GetCustomValueParametersOk returns a tuple with the CustomValueParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomValueParameters

`func (o *UpdateAdTrackingTagsRequest) SetCustomValueParameters(v map[string]string)`

SetCustomValueParameters sets CustomValueParameters field to given value.

### HasCustomValueParameters

`func (o *UpdateAdTrackingTagsRequest) HasCustomValueParameters() bool`

HasCustomValueParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


