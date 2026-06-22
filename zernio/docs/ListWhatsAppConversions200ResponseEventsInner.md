# ListWhatsAppConversions200ResponseEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamp** | Pointer to **time.Time** | When the event was sent to Meta. | [optional] 
**EventName** | Pointer to **string** | One of LeadSubmitted, Purchase, AddToCart, InitiateCheckout, ViewContent. | [optional] 
**ConversationId** | Pointer to **NullableString** |  | [optional] 
**EventsReceived** | Pointer to **NullableInt32** | Number of events Meta accepted on this send (usually 1). | [optional] 
**EventsFailed** | Pointer to **NullableInt32** | Number of events Meta rejected (usually 0). | [optional] 
**TraceId** | Pointer to **NullableString** | Meta fbtrace_id for cross-referencing in Events Manager. | [optional] 
**DurationMs** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewListWhatsAppConversions200ResponseEventsInner

`func NewListWhatsAppConversions200ResponseEventsInner() *ListWhatsAppConversions200ResponseEventsInner`

NewListWhatsAppConversions200ResponseEventsInner instantiates a new ListWhatsAppConversions200ResponseEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppConversions200ResponseEventsInnerWithDefaults

`func NewListWhatsAppConversions200ResponseEventsInnerWithDefaults() *ListWhatsAppConversions200ResponseEventsInner`

NewListWhatsAppConversions200ResponseEventsInnerWithDefaults instantiates a new ListWhatsAppConversions200ResponseEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamp

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetEventName

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetConversationId

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *ListWhatsAppConversions200ResponseEventsInner) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetEventsReceived

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetEventsReceived() int32`

GetEventsReceived returns the EventsReceived field if non-nil, zero value otherwise.

### GetEventsReceivedOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetEventsReceivedOk() (*int32, bool)`

GetEventsReceivedOk returns a tuple with the EventsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsReceived

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetEventsReceived(v int32)`

SetEventsReceived sets EventsReceived field to given value.

### HasEventsReceived

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasEventsReceived() bool`

HasEventsReceived returns a boolean if a field has been set.

### SetEventsReceivedNil

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetEventsReceivedNil(b bool)`

 SetEventsReceivedNil sets the value for EventsReceived to be an explicit nil

### UnsetEventsReceived
`func (o *ListWhatsAppConversions200ResponseEventsInner) UnsetEventsReceived()`

UnsetEventsReceived ensures that no value is present for EventsReceived, not even an explicit nil
### GetEventsFailed

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetEventsFailed() int32`

GetEventsFailed returns the EventsFailed field if non-nil, zero value otherwise.

### GetEventsFailedOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetEventsFailedOk() (*int32, bool)`

GetEventsFailedOk returns a tuple with the EventsFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventsFailed

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetEventsFailed(v int32)`

SetEventsFailed sets EventsFailed field to given value.

### HasEventsFailed

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasEventsFailed() bool`

HasEventsFailed returns a boolean if a field has been set.

### SetEventsFailedNil

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetEventsFailedNil(b bool)`

 SetEventsFailedNil sets the value for EventsFailed to be an explicit nil

### UnsetEventsFailed
`func (o *ListWhatsAppConversions200ResponseEventsInner) UnsetEventsFailed()`

UnsetEventsFailed ensures that no value is present for EventsFailed, not even an explicit nil
### GetTraceId

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### SetTraceIdNil

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetTraceIdNil(b bool)`

 SetTraceIdNil sets the value for TraceId to be an explicit nil

### UnsetTraceId
`func (o *ListWhatsAppConversions200ResponseEventsInner) UnsetTraceId()`

UnsetTraceId ensures that no value is present for TraceId, not even an explicit nil
### GetDurationMs

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ListWhatsAppConversions200ResponseEventsInner) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ListWhatsAppConversions200ResponseEventsInner) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *ListWhatsAppConversions200ResponseEventsInner) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *ListWhatsAppConversions200ResponseEventsInner) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


