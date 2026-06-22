# UpdateAdCampaignRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**Budget** | Pointer to [**UpdateAdCampaignRequestBudget**](UpdateAdCampaignRequestBudget.md) |  | [optional] 
**BidStrategy** | Pointer to [**BidStrategy**](BidStrategy.md) | Campaign-level default. Ad sets inherit this unless they override. | [optional] 
**Name** | Pointer to **string** | Rename the campaign (Meta only; other platforms return 501). At least one of budget/bidStrategy/name is required. | [optional] 

## Methods

### NewUpdateAdCampaignRequest

`func NewUpdateAdCampaignRequest(platform string, ) *UpdateAdCampaignRequest`

NewUpdateAdCampaignRequest instantiates a new UpdateAdCampaignRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdCampaignRequestWithDefaults

`func NewUpdateAdCampaignRequestWithDefaults() *UpdateAdCampaignRequest`

NewUpdateAdCampaignRequestWithDefaults instantiates a new UpdateAdCampaignRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *UpdateAdCampaignRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateAdCampaignRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateAdCampaignRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetBudget

`func (o *UpdateAdCampaignRequest) GetBudget() UpdateAdCampaignRequestBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdateAdCampaignRequest) GetBudgetOk() (*UpdateAdCampaignRequestBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdateAdCampaignRequest) SetBudget(v UpdateAdCampaignRequestBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdateAdCampaignRequest) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetBidStrategy

`func (o *UpdateAdCampaignRequest) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *UpdateAdCampaignRequest) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *UpdateAdCampaignRequest) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *UpdateAdCampaignRequest) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### GetName

`func (o *UpdateAdCampaignRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAdCampaignRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAdCampaignRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAdCampaignRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


