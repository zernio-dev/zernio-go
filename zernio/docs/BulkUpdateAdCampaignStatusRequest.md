# BulkUpdateAdCampaignStatusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Campaigns** | [**[]BulkUpdateAdCampaignStatusRequestCampaignsInner**](BulkUpdateAdCampaignStatusRequestCampaignsInner.md) |  | 

## Methods

### NewBulkUpdateAdCampaignStatusRequest

`func NewBulkUpdateAdCampaignStatusRequest(status string, campaigns []BulkUpdateAdCampaignStatusRequestCampaignsInner, ) *BulkUpdateAdCampaignStatusRequest`

NewBulkUpdateAdCampaignStatusRequest instantiates a new BulkUpdateAdCampaignStatusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateAdCampaignStatusRequestWithDefaults

`func NewBulkUpdateAdCampaignStatusRequestWithDefaults() *BulkUpdateAdCampaignStatusRequest`

NewBulkUpdateAdCampaignStatusRequestWithDefaults instantiates a new BulkUpdateAdCampaignStatusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkUpdateAdCampaignStatusRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkUpdateAdCampaignStatusRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkUpdateAdCampaignStatusRequest) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCampaigns

`func (o *BulkUpdateAdCampaignStatusRequest) GetCampaigns() []BulkUpdateAdCampaignStatusRequestCampaignsInner`

GetCampaigns returns the Campaigns field if non-nil, zero value otherwise.

### GetCampaignsOk

`func (o *BulkUpdateAdCampaignStatusRequest) GetCampaignsOk() (*[]BulkUpdateAdCampaignStatusRequestCampaignsInner, bool)`

GetCampaignsOk returns a tuple with the Campaigns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaigns

`func (o *BulkUpdateAdCampaignStatusRequest) SetCampaigns(v []BulkUpdateAdCampaignStatusRequestCampaignsInner)`

SetCampaigns sets Campaigns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


