# AdjustConversions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**AdjustmentsReceived** | Pointer to **int32** | Adjustments accepted by Google. | [optional] 
**AdjustmentsFailed** | Pointer to **int32** | Adjustments rejected (see failures). | [optional] 
**Failures** | Pointer to [**[]AdjustConversions200ResponseFailuresInner**](AdjustConversions200ResponseFailuresInner.md) |  | [optional] 
**TraceId** | Pointer to **string** |  | [optional] 

## Methods

### NewAdjustConversions200Response

`func NewAdjustConversions200Response() *AdjustConversions200Response`

NewAdjustConversions200Response instantiates a new AdjustConversions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustConversions200ResponseWithDefaults

`func NewAdjustConversions200ResponseWithDefaults() *AdjustConversions200Response`

NewAdjustConversions200ResponseWithDefaults instantiates a new AdjustConversions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *AdjustConversions200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AdjustConversions200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AdjustConversions200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AdjustConversions200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAdjustmentsReceived

`func (o *AdjustConversions200Response) GetAdjustmentsReceived() int32`

GetAdjustmentsReceived returns the AdjustmentsReceived field if non-nil, zero value otherwise.

### GetAdjustmentsReceivedOk

`func (o *AdjustConversions200Response) GetAdjustmentsReceivedOk() (*int32, bool)`

GetAdjustmentsReceivedOk returns a tuple with the AdjustmentsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentsReceived

`func (o *AdjustConversions200Response) SetAdjustmentsReceived(v int32)`

SetAdjustmentsReceived sets AdjustmentsReceived field to given value.

### HasAdjustmentsReceived

`func (o *AdjustConversions200Response) HasAdjustmentsReceived() bool`

HasAdjustmentsReceived returns a boolean if a field has been set.

### GetAdjustmentsFailed

`func (o *AdjustConversions200Response) GetAdjustmentsFailed() int32`

GetAdjustmentsFailed returns the AdjustmentsFailed field if non-nil, zero value otherwise.

### GetAdjustmentsFailedOk

`func (o *AdjustConversions200Response) GetAdjustmentsFailedOk() (*int32, bool)`

GetAdjustmentsFailedOk returns a tuple with the AdjustmentsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentsFailed

`func (o *AdjustConversions200Response) SetAdjustmentsFailed(v int32)`

SetAdjustmentsFailed sets AdjustmentsFailed field to given value.

### HasAdjustmentsFailed

`func (o *AdjustConversions200Response) HasAdjustmentsFailed() bool`

HasAdjustmentsFailed returns a boolean if a field has been set.

### GetFailures

`func (o *AdjustConversions200Response) GetFailures() []AdjustConversions200ResponseFailuresInner`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *AdjustConversions200Response) GetFailuresOk() (*[]AdjustConversions200ResponseFailuresInner, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *AdjustConversions200Response) SetFailures(v []AdjustConversions200ResponseFailuresInner)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *AdjustConversions200Response) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetTraceId

`func (o *AdjustConversions200Response) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *AdjustConversions200Response) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *AdjustConversions200Response) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *AdjustConversions200Response) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


