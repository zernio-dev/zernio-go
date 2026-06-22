# GetInboxResponseTime200ResponseHistogramInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bucket** | Pointer to **string** | Human label (0-1m, 1-5m, 5-15m, 15-60m, 1-4h, 4-24h, 1d+) | [optional] 
**LowerSeconds** | Pointer to **int32** |  | [optional] 
**UpperSeconds** | Pointer to **NullableInt32** | null on the open-ended last bucket | [optional] 
**Count** | Pointer to **int32** |  | [optional] 

## Methods

### NewGetInboxResponseTime200ResponseHistogramInner

`func NewGetInboxResponseTime200ResponseHistogramInner() *GetInboxResponseTime200ResponseHistogramInner`

NewGetInboxResponseTime200ResponseHistogramInner instantiates a new GetInboxResponseTime200ResponseHistogramInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxResponseTime200ResponseHistogramInnerWithDefaults

`func NewGetInboxResponseTime200ResponseHistogramInnerWithDefaults() *GetInboxResponseTime200ResponseHistogramInner`

NewGetInboxResponseTime200ResponseHistogramInnerWithDefaults instantiates a new GetInboxResponseTime200ResponseHistogramInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucket

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *GetInboxResponseTime200ResponseHistogramInner) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *GetInboxResponseTime200ResponseHistogramInner) HasBucket() bool`

HasBucket returns a boolean if a field has been set.

### GetLowerSeconds

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetLowerSeconds() int32`

GetLowerSeconds returns the LowerSeconds field if non-nil, zero value otherwise.

### GetLowerSecondsOk

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetLowerSecondsOk() (*int32, bool)`

GetLowerSecondsOk returns a tuple with the LowerSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowerSeconds

`func (o *GetInboxResponseTime200ResponseHistogramInner) SetLowerSeconds(v int32)`

SetLowerSeconds sets LowerSeconds field to given value.

### HasLowerSeconds

`func (o *GetInboxResponseTime200ResponseHistogramInner) HasLowerSeconds() bool`

HasLowerSeconds returns a boolean if a field has been set.

### GetUpperSeconds

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetUpperSeconds() int32`

GetUpperSeconds returns the UpperSeconds field if non-nil, zero value otherwise.

### GetUpperSecondsOk

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetUpperSecondsOk() (*int32, bool)`

GetUpperSecondsOk returns a tuple with the UpperSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperSeconds

`func (o *GetInboxResponseTime200ResponseHistogramInner) SetUpperSeconds(v int32)`

SetUpperSeconds sets UpperSeconds field to given value.

### HasUpperSeconds

`func (o *GetInboxResponseTime200ResponseHistogramInner) HasUpperSeconds() bool`

HasUpperSeconds returns a boolean if a field has been set.

### SetUpperSecondsNil

`func (o *GetInboxResponseTime200ResponseHistogramInner) SetUpperSecondsNil(b bool)`

 SetUpperSecondsNil sets the value for UpperSeconds to be an explicit nil

### UnsetUpperSeconds
`func (o *GetInboxResponseTime200ResponseHistogramInner) UnsetUpperSeconds()`

UnsetUpperSeconds ensures that no value is present for UpperSeconds, not even an explicit nil
### GetCount

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *GetInboxResponseTime200ResponseHistogramInner) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *GetInboxResponseTime200ResponseHistogramInner) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *GetInboxResponseTime200ResponseHistogramInner) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


