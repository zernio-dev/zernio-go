# EstimateAdReach200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Whether a pre-flight estimate is available on this platform. False for Google and TikTok. | 
**Lower** | Pointer to **NullableInt32** | Lower bound of the estimated reachable audience. Present only when available. | [optional] 
**Upper** | Pointer to **NullableInt32** | Upper bound of the estimated reachable audience. Present only when available. | [optional] 
**Daily** | Pointer to **NullableInt32** | Optional estimated daily reach/results at the given budget, when the platform returns it. | [optional] 
**Currency** | Pointer to **NullableString** | Currency of any monetary fields in the estimate, when applicable. | [optional] 
**EstimateReady** | Pointer to **NullableBool** | Meta only. False when Meta is still computing the estimate (the audience is too new); retry shortly. | [optional] 

## Methods

### NewEstimateAdReach200Response

`func NewEstimateAdReach200Response(available bool, ) *EstimateAdReach200Response`

NewEstimateAdReach200Response instantiates a new EstimateAdReach200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimateAdReach200ResponseWithDefaults

`func NewEstimateAdReach200ResponseWithDefaults() *EstimateAdReach200Response`

NewEstimateAdReach200ResponseWithDefaults instantiates a new EstimateAdReach200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *EstimateAdReach200Response) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *EstimateAdReach200Response) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *EstimateAdReach200Response) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetLower

`func (o *EstimateAdReach200Response) GetLower() int32`

GetLower returns the Lower field if non-nil, zero value otherwise.

### GetLowerOk

`func (o *EstimateAdReach200Response) GetLowerOk() (*int32, bool)`

GetLowerOk returns a tuple with the Lower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLower

`func (o *EstimateAdReach200Response) SetLower(v int32)`

SetLower sets Lower field to given value.

### HasLower

`func (o *EstimateAdReach200Response) HasLower() bool`

HasLower returns a boolean if a field has been set.

### SetLowerNil

`func (o *EstimateAdReach200Response) SetLowerNil(b bool)`

 SetLowerNil sets the value for Lower to be an explicit nil

### UnsetLower
`func (o *EstimateAdReach200Response) UnsetLower()`

UnsetLower ensures that no value is present for Lower, not even an explicit nil
### GetUpper

`func (o *EstimateAdReach200Response) GetUpper() int32`

GetUpper returns the Upper field if non-nil, zero value otherwise.

### GetUpperOk

`func (o *EstimateAdReach200Response) GetUpperOk() (*int32, bool)`

GetUpperOk returns a tuple with the Upper field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpper

`func (o *EstimateAdReach200Response) SetUpper(v int32)`

SetUpper sets Upper field to given value.

### HasUpper

`func (o *EstimateAdReach200Response) HasUpper() bool`

HasUpper returns a boolean if a field has been set.

### SetUpperNil

`func (o *EstimateAdReach200Response) SetUpperNil(b bool)`

 SetUpperNil sets the value for Upper to be an explicit nil

### UnsetUpper
`func (o *EstimateAdReach200Response) UnsetUpper()`

UnsetUpper ensures that no value is present for Upper, not even an explicit nil
### GetDaily

`func (o *EstimateAdReach200Response) GetDaily() int32`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *EstimateAdReach200Response) GetDailyOk() (*int32, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *EstimateAdReach200Response) SetDaily(v int32)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *EstimateAdReach200Response) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### SetDailyNil

`func (o *EstimateAdReach200Response) SetDailyNil(b bool)`

 SetDailyNil sets the value for Daily to be an explicit nil

### UnsetDaily
`func (o *EstimateAdReach200Response) UnsetDaily()`

UnsetDaily ensures that no value is present for Daily, not even an explicit nil
### GetCurrency

`func (o *EstimateAdReach200Response) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EstimateAdReach200Response) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EstimateAdReach200Response) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EstimateAdReach200Response) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *EstimateAdReach200Response) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *EstimateAdReach200Response) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetEstimateReady

`func (o *EstimateAdReach200Response) GetEstimateReady() bool`

GetEstimateReady returns the EstimateReady field if non-nil, zero value otherwise.

### GetEstimateReadyOk

`func (o *EstimateAdReach200Response) GetEstimateReadyOk() (*bool, bool)`

GetEstimateReadyOk returns a tuple with the EstimateReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimateReady

`func (o *EstimateAdReach200Response) SetEstimateReady(v bool)`

SetEstimateReady sets EstimateReady field to given value.

### HasEstimateReady

`func (o *EstimateAdReach200Response) HasEstimateReady() bool`

HasEstimateReady returns a boolean if a field has been set.

### SetEstimateReadyNil

`func (o *EstimateAdReach200Response) SetEstimateReadyNil(b bool)`

 SetEstimateReadyNil sets the value for EstimateReady to be an explicit nil

### UnsetEstimateReady
`func (o *EstimateAdReach200Response) UnsetEstimateReady()`

UnsetEstimateReady ensures that no value is present for EstimateReady, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


