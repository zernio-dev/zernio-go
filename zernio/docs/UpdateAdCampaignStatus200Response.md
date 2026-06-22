# UpdateAdCampaignStatus200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | Pointer to **int32** | Number of ads updated | [optional] 
**Skipped** | Pointer to **int32** | Number of ads skipped | [optional] 
**SkippedReasons** | Pointer to **[]string** |  | [optional] 
**Message** | Pointer to **string** | Human-readable summary (present when no ads were actionable) | [optional] 

## Methods

### NewUpdateAdCampaignStatus200Response

`func NewUpdateAdCampaignStatus200Response() *UpdateAdCampaignStatus200Response`

NewUpdateAdCampaignStatus200Response instantiates a new UpdateAdCampaignStatus200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdCampaignStatus200ResponseWithDefaults

`func NewUpdateAdCampaignStatus200ResponseWithDefaults() *UpdateAdCampaignStatus200Response`

NewUpdateAdCampaignStatus200ResponseWithDefaults instantiates a new UpdateAdCampaignStatus200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *UpdateAdCampaignStatus200Response) GetUpdated() int32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *UpdateAdCampaignStatus200Response) GetUpdatedOk() (*int32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *UpdateAdCampaignStatus200Response) SetUpdated(v int32)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *UpdateAdCampaignStatus200Response) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetSkipped

`func (o *UpdateAdCampaignStatus200Response) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *UpdateAdCampaignStatus200Response) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *UpdateAdCampaignStatus200Response) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *UpdateAdCampaignStatus200Response) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetSkippedReasons

`func (o *UpdateAdCampaignStatus200Response) GetSkippedReasons() []string`

GetSkippedReasons returns the SkippedReasons field if non-nil, zero value otherwise.

### GetSkippedReasonsOk

`func (o *UpdateAdCampaignStatus200Response) GetSkippedReasonsOk() (*[]string, bool)`

GetSkippedReasonsOk returns a tuple with the SkippedReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedReasons

`func (o *UpdateAdCampaignStatus200Response) SetSkippedReasons(v []string)`

SetSkippedReasons sets SkippedReasons field to given value.

### HasSkippedReasons

`func (o *UpdateAdCampaignStatus200Response) HasSkippedReasons() bool`

HasSkippedReasons returns a boolean if a field has been set.

### GetMessage

`func (o *UpdateAdCampaignStatus200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateAdCampaignStatus200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateAdCampaignStatus200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateAdCampaignStatus200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


