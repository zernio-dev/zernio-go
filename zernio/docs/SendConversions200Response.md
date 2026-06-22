# SendConversions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**EventsReceived** | Pointer to **int32** | Events accepted by the platform. | [optional] 
**EventsFailed** | Pointer to **int32** | Events rejected (see failures). | [optional] 
**Failures** | Pointer to [**[]SendConversions200ResponseFailuresInner**](SendConversions200ResponseFailuresInner.md) |  | [optional] 
**TraceId** | Pointer to **string** | Platform trace ID for debugging. fbtrace_id for Meta, requestId for Google. Absent for LinkedIn (LinkedIn&#39;s conversionEvents endpoint does not surface a trace ID).  | [optional] 

## Methods

### NewSendConversions200Response

`func NewSendConversions200Response() *SendConversions200Response`

NewSendConversions200Response instantiates a new SendConversions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendConversions200ResponseWithDefaults

`func NewSendConversions200ResponseWithDefaults() *SendConversions200Response`

NewSendConversions200ResponseWithDefaults instantiates a new SendConversions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *SendConversions200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SendConversions200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SendConversions200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SendConversions200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEventsReceived

`func (o *SendConversions200Response) GetEventsReceived() int32`

GetEventsReceived returns the EventsReceived field if non-nil, zero value otherwise.

### GetEventsReceivedOk

`func (o *SendConversions200Response) GetEventsReceivedOk() (*int32, bool)`

GetEventsReceivedOk returns a tuple with the EventsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsReceived

`func (o *SendConversions200Response) SetEventsReceived(v int32)`

SetEventsReceived sets EventsReceived field to given value.

### HasEventsReceived

`func (o *SendConversions200Response) HasEventsReceived() bool`

HasEventsReceived returns a boolean if a field has been set.

### GetEventsFailed

`func (o *SendConversions200Response) GetEventsFailed() int32`

GetEventsFailed returns the EventsFailed field if non-nil, zero value otherwise.

### GetEventsFailedOk

`func (o *SendConversions200Response) GetEventsFailedOk() (*int32, bool)`

GetEventsFailedOk returns a tuple with the EventsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsFailed

`func (o *SendConversions200Response) SetEventsFailed(v int32)`

SetEventsFailed sets EventsFailed field to given value.

### HasEventsFailed

`func (o *SendConversions200Response) HasEventsFailed() bool`

HasEventsFailed returns a boolean if a field has been set.

### GetFailures

`func (o *SendConversions200Response) GetFailures() []SendConversions200ResponseFailuresInner`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *SendConversions200Response) GetFailuresOk() (*[]SendConversions200ResponseFailuresInner, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *SendConversions200Response) SetFailures(v []SendConversions200ResponseFailuresInner)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *SendConversions200Response) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetTraceId

`func (o *SendConversions200Response) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SendConversions200Response) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SendConversions200Response) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *SendConversions200Response) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


