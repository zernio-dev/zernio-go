# DuplicateAdCampaign200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CopiedCampaignId** | Pointer to **string** | Platform ID of the new campaign | [optional] 
**Discovery** | Pointer to **string** |  | [optional] 
**Raw** | Pointer to **map[string]interface{}** | Platform-native response from the copy endpoint (Meta includes ad_object_ids for child copies) | [optional] 

## Methods

### NewDuplicateAdCampaign200Response

`func NewDuplicateAdCampaign200Response() *DuplicateAdCampaign200Response`

NewDuplicateAdCampaign200Response instantiates a new DuplicateAdCampaign200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicateAdCampaign200ResponseWithDefaults

`func NewDuplicateAdCampaign200ResponseWithDefaults() *DuplicateAdCampaign200Response`

NewDuplicateAdCampaign200ResponseWithDefaults instantiates a new DuplicateAdCampaign200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCopiedCampaignId

`func (o *DuplicateAdCampaign200Response) GetCopiedCampaignId() string`

GetCopiedCampaignId returns the CopiedCampaignId field if non-nil, zero value otherwise.

### GetCopiedCampaignIdOk

`func (o *DuplicateAdCampaign200Response) GetCopiedCampaignIdOk() (*string, bool)`

GetCopiedCampaignIdOk returns a tuple with the CopiedCampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopiedCampaignId

`func (o *DuplicateAdCampaign200Response) SetCopiedCampaignId(v string)`

SetCopiedCampaignId sets CopiedCampaignId field to given value.

### HasCopiedCampaignId

`func (o *DuplicateAdCampaign200Response) HasCopiedCampaignId() bool`

HasCopiedCampaignId returns a boolean if a field has been set.

### GetDiscovery

`func (o *DuplicateAdCampaign200Response) GetDiscovery() string`

GetDiscovery returns the Discovery field if non-nil, zero value otherwise.

### GetDiscoveryOk

`func (o *DuplicateAdCampaign200Response) GetDiscoveryOk() (*string, bool)`

GetDiscoveryOk returns a tuple with the Discovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscovery

`func (o *DuplicateAdCampaign200Response) SetDiscovery(v string)`

SetDiscovery sets Discovery field to given value.

### HasDiscovery

`func (o *DuplicateAdCampaign200Response) HasDiscovery() bool`

HasDiscovery returns a boolean if a field has been set.

### GetRaw

`func (o *DuplicateAdCampaign200Response) GetRaw() map[string]interface{}`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *DuplicateAdCampaign200Response) GetRawOk() (*map[string]interface{}, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *DuplicateAdCampaign200Response) SetRaw(v map[string]interface{})`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *DuplicateAdCampaign200Response) HasRaw() bool`

HasRaw returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


