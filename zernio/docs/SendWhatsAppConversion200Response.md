# SendWhatsAppConversion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**EventsReceived** | Pointer to **int32** | Events accepted by Meta. | [optional] 
**EventsFailed** | Pointer to **int32** | Events rejected by Meta (see failures). | [optional] 
**Failures** | Pointer to [**[]SendConversions200ResponseFailuresInner**](SendConversions200ResponseFailuresInner.md) | Per-event failure detail. Empty when all events were accepted.  | [optional] 
**TraceId** | Pointer to **string** | Meta &#x60;fbtrace_id&#x60; for debugging. Surface in support tickets.  | [optional] 

## Methods

### NewSendWhatsAppConversion200Response

`func NewSendWhatsAppConversion200Response() *SendWhatsAppConversion200Response`

NewSendWhatsAppConversion200Response instantiates a new SendWhatsAppConversion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWhatsAppConversion200ResponseWithDefaults

`func NewSendWhatsAppConversion200ResponseWithDefaults() *SendWhatsAppConversion200Response`

NewSendWhatsAppConversion200ResponseWithDefaults instantiates a new SendWhatsAppConversion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *SendWhatsAppConversion200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *SendWhatsAppConversion200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *SendWhatsAppConversion200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *SendWhatsAppConversion200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetEventsReceived

`func (o *SendWhatsAppConversion200Response) GetEventsReceived() int32`

GetEventsReceived returns the EventsReceived field if non-nil, zero value otherwise.

### GetEventsReceivedOk

`func (o *SendWhatsAppConversion200Response) GetEventsReceivedOk() (*int32, bool)`

GetEventsReceivedOk returns a tuple with the EventsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsReceived

`func (o *SendWhatsAppConversion200Response) SetEventsReceived(v int32)`

SetEventsReceived sets EventsReceived field to given value.

### HasEventsReceived

`func (o *SendWhatsAppConversion200Response) HasEventsReceived() bool`

HasEventsReceived returns a boolean if a field has been set.

### GetEventsFailed

`func (o *SendWhatsAppConversion200Response) GetEventsFailed() int32`

GetEventsFailed returns the EventsFailed field if non-nil, zero value otherwise.

### GetEventsFailedOk

`func (o *SendWhatsAppConversion200Response) GetEventsFailedOk() (*int32, bool)`

GetEventsFailedOk returns a tuple with the EventsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsFailed

`func (o *SendWhatsAppConversion200Response) SetEventsFailed(v int32)`

SetEventsFailed sets EventsFailed field to given value.

### HasEventsFailed

`func (o *SendWhatsAppConversion200Response) HasEventsFailed() bool`

HasEventsFailed returns a boolean if a field has been set.

### GetFailures

`func (o *SendWhatsAppConversion200Response) GetFailures() []SendConversions200ResponseFailuresInner`

GetFailures returns the Failures field if non-nil, zero value otherwise.

### GetFailuresOk

`func (o *SendWhatsAppConversion200Response) GetFailuresOk() (*[]SendConversions200ResponseFailuresInner, bool)`

GetFailuresOk returns a tuple with the Failures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailures

`func (o *SendWhatsAppConversion200Response) SetFailures(v []SendConversions200ResponseFailuresInner)`

SetFailures sets Failures field to given value.

### HasFailures

`func (o *SendWhatsAppConversion200Response) HasFailures() bool`

HasFailures returns a boolean if a field has been set.

### GetTraceId

`func (o *SendWhatsAppConversion200Response) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *SendWhatsAppConversion200Response) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *SendWhatsAppConversion200Response) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *SendWhatsAppConversion200Response) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


