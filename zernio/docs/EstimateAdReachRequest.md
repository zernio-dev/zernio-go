# EstimateAdReachRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Zernio social account ID on the target ad platform (the estimate runs against its platform). | 
**AdAccountId** | **string** | Required. The platform ad-account ID the reach call runs against (Meta act_..., LinkedIn numeric sponsoredAccount ID, Pinterest ad-account ID, X account ID) - every backing reach API is scoped to one ad account. Get it from GET /v1/ads/accounts. | 
**Spec** | [**TargetingSpec**](TargetingSpec.md) | The targeting spec to estimate. Same shape used by POST /v1/ads/create. | 
**OptimizationGoal** | Pointer to **string** | Optional. The optimization goal the estimate should assume (platform&#39;s own vocabulary, e.g. Meta &#x60;REACH&#x60;, &#x60;LINK_CLICKS&#x60;, &#x60;OFFSITE_CONVERSIONS&#x60;). Some platforms vary the estimate by goal; omit to use the platform default.  | [optional] 

## Methods

### NewEstimateAdReachRequest

`func NewEstimateAdReachRequest(accountId string, adAccountId string, spec TargetingSpec, ) *EstimateAdReachRequest`

NewEstimateAdReachRequest instantiates a new EstimateAdReachRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateAdReachRequestWithDefaults

`func NewEstimateAdReachRequestWithDefaults() *EstimateAdReachRequest`

NewEstimateAdReachRequestWithDefaults instantiates a new EstimateAdReachRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *EstimateAdReachRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *EstimateAdReachRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *EstimateAdReachRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAdAccountId

`func (o *EstimateAdReachRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *EstimateAdReachRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *EstimateAdReachRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetSpec

`func (o *EstimateAdReachRequest) GetSpec() TargetingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *EstimateAdReachRequest) GetSpecOk() (*TargetingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *EstimateAdReachRequest) SetSpec(v TargetingSpec)`

SetSpec sets Spec field to given value.


### GetOptimizationGoal

`func (o *EstimateAdReachRequest) GetOptimizationGoal() string`

GetOptimizationGoal returns the OptimizationGoal field if non-nil, zero value otherwise.

### GetOptimizationGoalOk

`func (o *EstimateAdReachRequest) GetOptimizationGoalOk() (*string, bool)`

GetOptimizationGoalOk returns a tuple with the OptimizationGoal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizationGoal

`func (o *EstimateAdReachRequest) SetOptimizationGoal(v string)`

SetOptimizationGoal sets OptimizationGoal field to given value.

### HasOptimizationGoal

`func (o *EstimateAdReachRequest) HasOptimizationGoal() bool`

HasOptimizationGoal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


