# CreateCtwaAd201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdType** | **string** |  | 
**Ad** | **map[string]interface{}** | The persisted Ad document. | 
**Message** | **string** |  | 
**Ads** | **[]map[string]interface{}** | The persisted Ad documents (one per creative), all sharing the same &#x60;platformCampaignId&#x60; and &#x60;platformAdSetId&#x60;.  | 
**PlatformCampaignId** | **string** |  | 
**PlatformAdSetId** | **string** |  | 

## Methods

### NewCreateCtwaAd201Response

`func NewCreateCtwaAd201Response(adType string, ad map[string]interface{}, message string, ads []map[string]interface{}, platformCampaignId string, platformAdSetId string, ) *CreateCtwaAd201Response`

NewCreateCtwaAd201Response instantiates a new CreateCtwaAd201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCtwaAd201ResponseWithDefaults

`func NewCreateCtwaAd201ResponseWithDefaults() *CreateCtwaAd201Response`

NewCreateCtwaAd201ResponseWithDefaults instantiates a new CreateCtwaAd201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdType

`func (o *CreateCtwaAd201Response) GetAdType() string`

GetAdType returns the AdType field if non-nil, zero value otherwise.

### GetAdTypeOk

`func (o *CreateCtwaAd201Response) GetAdTypeOk() (*string, bool)`

GetAdTypeOk returns a tuple with the AdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdType

`func (o *CreateCtwaAd201Response) SetAdType(v string)`

SetAdType sets AdType field to given value.


### GetAd

`func (o *CreateCtwaAd201Response) GetAd() map[string]interface{}`

GetAd returns the Ad field if non-nil, zero value otherwise.

### GetAdOk

`func (o *CreateCtwaAd201Response) GetAdOk() (*map[string]interface{}, bool)`

GetAdOk returns a tuple with the Ad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAd

`func (o *CreateCtwaAd201Response) SetAd(v map[string]interface{})`

SetAd sets Ad field to given value.


### GetMessage

`func (o *CreateCtwaAd201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateCtwaAd201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateCtwaAd201Response) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetAds

`func (o *CreateCtwaAd201Response) GetAds() []map[string]interface{}`

GetAds returns the Ads field if non-nil, zero value otherwise.

### GetAdsOk

`func (o *CreateCtwaAd201Response) GetAdsOk() (*[]map[string]interface{}, bool)`

GetAdsOk returns a tuple with the Ads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAds

`func (o *CreateCtwaAd201Response) SetAds(v []map[string]interface{})`

SetAds sets Ads field to given value.


### GetPlatformCampaignId

`func (o *CreateCtwaAd201Response) GetPlatformCampaignId() string`

GetPlatformCampaignId returns the PlatformCampaignId field if non-nil, zero value otherwise.

### GetPlatformCampaignIdOk

`func (o *CreateCtwaAd201Response) GetPlatformCampaignIdOk() (*string, bool)`

GetPlatformCampaignIdOk returns a tuple with the PlatformCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCampaignId

`func (o *CreateCtwaAd201Response) SetPlatformCampaignId(v string)`

SetPlatformCampaignId sets PlatformCampaignId field to given value.


### GetPlatformAdSetId

`func (o *CreateCtwaAd201Response) GetPlatformAdSetId() string`

GetPlatformAdSetId returns the PlatformAdSetId field if non-nil, zero value otherwise.

### GetPlatformAdSetIdOk

`func (o *CreateCtwaAd201Response) GetPlatformAdSetIdOk() (*string, bool)`

GetPlatformAdSetIdOk returns a tuple with the PlatformAdSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdSetId

`func (o *CreateCtwaAd201Response) SetPlatformAdSetId(v string)`

SetPlatformAdSetId sets PlatformAdSetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


