# BusinessCenter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BcId** | Pointer to **string** | Business Center ID | [optional] 
**Name** | Pointer to **string** | Display name set by the BC owner | [optional] 
**AdvertiserCount** | Pointer to **NullableInt32** | Number of advertisers reachable under this BC for the calling token. &#x60;null&#x60; when the BC asset walk returned empty or failed (typical for agency apps without full BC asset read scope) — distinct from &#x60;0&#x60;, which would imply the BC genuinely has no advertisers.  | [optional] 

## Methods

### NewBusinessCenter

`func NewBusinessCenter() *BusinessCenter`

NewBusinessCenter instantiates a new BusinessCenter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessCenterWithDefaults

`func NewBusinessCenterWithDefaults() *BusinessCenter`

NewBusinessCenterWithDefaults instantiates a new BusinessCenter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcId

`func (o *BusinessCenter) GetBcId() string`

GetBcId returns the BcId field if non-nil, zero value otherwise.

### GetBcIdOk

`func (o *BusinessCenter) GetBcIdOk() (*string, bool)`

GetBcIdOk returns a tuple with the BcId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcId

`func (o *BusinessCenter) SetBcId(v string)`

SetBcId sets BcId field to given value.

### HasBcId

`func (o *BusinessCenter) HasBcId() bool`

HasBcId returns a boolean if a field has been set.

### GetName

`func (o *BusinessCenter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BusinessCenter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BusinessCenter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BusinessCenter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdvertiserCount

`func (o *BusinessCenter) GetAdvertiserCount() int32`

GetAdvertiserCount returns the AdvertiserCount field if non-nil, zero value otherwise.

### GetAdvertiserCountOk

`func (o *BusinessCenter) GetAdvertiserCountOk() (*int32, bool)`

GetAdvertiserCountOk returns a tuple with the AdvertiserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiserCount

`func (o *BusinessCenter) SetAdvertiserCount(v int32)`

SetAdvertiserCount sets AdvertiserCount field to given value.

### HasAdvertiserCount

`func (o *BusinessCenter) HasAdvertiserCount() bool`

HasAdvertiserCount returns a boolean if a field has been set.

### SetAdvertiserCountNil

`func (o *BusinessCenter) SetAdvertiserCountNil(b bool)`

 SetAdvertiserCountNil sets the value for AdvertiserCount to be an explicit nil

### UnsetAdvertiserCount
`func (o *BusinessCenter) UnsetAdvertiserCount()`

UnsetAdvertiserCount ensures that no value is present for AdvertiserCount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


