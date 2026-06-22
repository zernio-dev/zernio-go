# CtwaMultiResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdType** | **string** |  | 
**Ads** | **[]map[string]interface{}** | The persisted Ad documents (one per creative), all sharing the same &#x60;platformCampaignId&#x60; and &#x60;platformAdSetId&#x60;.  | 
**PlatformCampaignId** | **string** |  | 
**PlatformAdSetId** | **string** |  | 
**Message** | **string** |  | 

## Methods

### NewCtwaMultiResponse

`func NewCtwaMultiResponse(adType string, ads []map[string]interface{}, platformCampaignId string, platformAdSetId string, message string, ) *CtwaMultiResponse`

NewCtwaMultiResponse instantiates a new CtwaMultiResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCtwaMultiResponseWithDefaults

`func NewCtwaMultiResponseWithDefaults() *CtwaMultiResponse`

NewCtwaMultiResponseWithDefaults instantiates a new CtwaMultiResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdType

`func (o *CtwaMultiResponse) GetAdType() string`

GetAdType returns the AdType field if non-nil, zero value otherwise.

### GetAdTypeOk

`func (o *CtwaMultiResponse) GetAdTypeOk() (*string, bool)`

GetAdTypeOk returns a tuple with the AdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdType

`func (o *CtwaMultiResponse) SetAdType(v string)`

SetAdType sets AdType field to given value.


### GetAds

`func (o *CtwaMultiResponse) GetAds() []map[string]interface{}`

GetAds returns the Ads field if non-nil, zero value otherwise.

### GetAdsOk

`func (o *CtwaMultiResponse) GetAdsOk() (*[]map[string]interface{}, bool)`

GetAdsOk returns a tuple with the Ads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAds

`func (o *CtwaMultiResponse) SetAds(v []map[string]interface{})`

SetAds sets Ads field to given value.


### GetPlatformCampaignId

`func (o *CtwaMultiResponse) GetPlatformCampaignId() string`

GetPlatformCampaignId returns the PlatformCampaignId field if non-nil, zero value otherwise.

### GetPlatformCampaignIdOk

`func (o *CtwaMultiResponse) GetPlatformCampaignIdOk() (*string, bool)`

GetPlatformCampaignIdOk returns a tuple with the PlatformCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignId

`func (o *CtwaMultiResponse) SetPlatformCampaignId(v string)`

SetPlatformCampaignId sets PlatformCampaignId field to given value.


### GetPlatformAdSetId

`func (o *CtwaMultiResponse) GetPlatformAdSetId() string`

GetPlatformAdSetId returns the PlatformAdSetId field if non-nil, zero value otherwise.

### GetPlatformAdSetIdOk

`func (o *CtwaMultiResponse) GetPlatformAdSetIdOk() (*string, bool)`

GetPlatformAdSetIdOk returns a tuple with the PlatformAdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdSetId

`func (o *CtwaMultiResponse) SetPlatformAdSetId(v string)`

SetPlatformAdSetId sets PlatformAdSetId field to given value.


### GetMessage

`func (o *CtwaMultiResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CtwaMultiResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CtwaMultiResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


