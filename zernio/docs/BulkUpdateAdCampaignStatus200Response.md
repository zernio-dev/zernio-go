# BulkUpdateAdCampaignStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Totals** | Pointer to [**BulkUpdateAdCampaignStatus200ResponseTotals**](BulkUpdateAdCampaignStatus200ResponseTotals.md) |  | [optional] 
**Results** | Pointer to [**[]BulkUpdateAdCampaignStatus200ResponseResultsInner**](BulkUpdateAdCampaignStatus200ResponseResultsInner.md) |  | [optional] 

## Methods

### NewBulkUpdateAdCampaignStatus200Response

`func NewBulkUpdateAdCampaignStatus200Response() *BulkUpdateAdCampaignStatus200Response`

NewBulkUpdateAdCampaignStatus200Response instantiates a new BulkUpdateAdCampaignStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUpdateAdCampaignStatus200ResponseWithDefaults

`func NewBulkUpdateAdCampaignStatus200ResponseWithDefaults() *BulkUpdateAdCampaignStatus200Response`

NewBulkUpdateAdCampaignStatus200ResponseWithDefaults instantiates a new BulkUpdateAdCampaignStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BulkUpdateAdCampaignStatus200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkUpdateAdCampaignStatus200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkUpdateAdCampaignStatus200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkUpdateAdCampaignStatus200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotals

`func (o *BulkUpdateAdCampaignStatus200Response) GetTotals() BulkUpdateAdCampaignStatus200ResponseTotals`

GetTotals returns the Totals field if non-nil, zero value otherwise.

### GetTotalsOk

`func (o *BulkUpdateAdCampaignStatus200Response) GetTotalsOk() (*BulkUpdateAdCampaignStatus200ResponseTotals, bool)`

GetTotalsOk returns a tuple with the Totals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotals

`func (o *BulkUpdateAdCampaignStatus200Response) SetTotals(v BulkUpdateAdCampaignStatus200ResponseTotals)`

SetTotals sets Totals field to given value.

### HasTotals

`func (o *BulkUpdateAdCampaignStatus200Response) HasTotals() bool`

HasTotals returns a boolean if a field has been set.

### GetResults

`func (o *BulkUpdateAdCampaignStatus200Response) GetResults() []BulkUpdateAdCampaignStatus200ResponseResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkUpdateAdCampaignStatus200Response) GetResultsOk() (*[]BulkUpdateAdCampaignStatus200ResponseResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkUpdateAdCampaignStatus200Response) SetResults(v []BulkUpdateAdCampaignStatus200ResponseResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkUpdateAdCampaignStatus200Response) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


