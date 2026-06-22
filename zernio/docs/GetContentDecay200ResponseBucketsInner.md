# GetContentDecay200ResponseBucketsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BucketOrder** | Pointer to **int32** | Sort order (0 &#x3D; earliest, 6 &#x3D; latest) | [optional] 
**BucketLabel** | Pointer to **string** | Human-readable label | [optional] 
**AvgPctOfFinal** | Pointer to **float32** | Average % of final engagement reached (0-100) | [optional] 
**PostCount** | Pointer to **int32** | Number of posts with data in this bucket | [optional] 

## Methods

### NewGetContentDecay200ResponseBucketsInner

`func NewGetContentDecay200ResponseBucketsInner() *GetContentDecay200ResponseBucketsInner`

NewGetContentDecay200ResponseBucketsInner instantiates a new GetContentDecay200ResponseBucketsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetContentDecay200ResponseBucketsInnerWithDefaults

`func NewGetContentDecay200ResponseBucketsInnerWithDefaults() *GetContentDecay200ResponseBucketsInner`

NewGetContentDecay200ResponseBucketsInnerWithDefaults instantiates a new GetContentDecay200ResponseBucketsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBucketOrder

`func (o *GetContentDecay200ResponseBucketsInner) GetBucketOrder() int32`

GetBucketOrder returns the BucketOrder field if non-nil, zero value otherwise.

### GetBucketOrderOk

`func (o *GetContentDecay200ResponseBucketsInner) GetBucketOrderOk() (*int32, bool)`

GetBucketOrderOk returns a tuple with the BucketOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketOrder

`func (o *GetContentDecay200ResponseBucketsInner) SetBucketOrder(v int32)`

SetBucketOrder sets BucketOrder field to given value.

### HasBucketOrder

`func (o *GetContentDecay200ResponseBucketsInner) HasBucketOrder() bool`

HasBucketOrder returns a boolean if a field has been set.

### GetBucketLabel

`func (o *GetContentDecay200ResponseBucketsInner) GetBucketLabel() string`

GetBucketLabel returns the BucketLabel field if non-nil, zero value otherwise.

### GetBucketLabelOk

`func (o *GetContentDecay200ResponseBucketsInner) GetBucketLabelOk() (*string, bool)`

GetBucketLabelOk returns a tuple with the BucketLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketLabel

`func (o *GetContentDecay200ResponseBucketsInner) SetBucketLabel(v string)`

SetBucketLabel sets BucketLabel field to given value.

### HasBucketLabel

`func (o *GetContentDecay200ResponseBucketsInner) HasBucketLabel() bool`

HasBucketLabel returns a boolean if a field has been set.

### GetAvgPctOfFinal

`func (o *GetContentDecay200ResponseBucketsInner) GetAvgPctOfFinal() float32`

GetAvgPctOfFinal returns the AvgPctOfFinal field if non-nil, zero value otherwise.

### GetAvgPctOfFinalOk

`func (o *GetContentDecay200ResponseBucketsInner) GetAvgPctOfFinalOk() (*float32, bool)`

GetAvgPctOfFinalOk returns a tuple with the AvgPctOfFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgPctOfFinal

`func (o *GetContentDecay200ResponseBucketsInner) SetAvgPctOfFinal(v float32)`

SetAvgPctOfFinal sets AvgPctOfFinal field to given value.

### HasAvgPctOfFinal

`func (o *GetContentDecay200ResponseBucketsInner) HasAvgPctOfFinal() bool`

HasAvgPctOfFinal returns a boolean if a field has been set.

### GetPostCount

`func (o *GetContentDecay200ResponseBucketsInner) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *GetContentDecay200ResponseBucketsInner) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *GetContentDecay200ResponseBucketsInner) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.

### HasPostCount

`func (o *GetContentDecay200ResponseBucketsInner) HasPostCount() bool`

HasPostCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


