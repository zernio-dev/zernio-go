# SendConversions200ResponseFailuresInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventIndex** | Pointer to **int32** | Index into the submitted events array. | [optional] 
**EventId** | Pointer to **string** | Echoes back the eventId of the failed event. | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to [**SendConversions200ResponseFailuresInnerCode**](SendConversions200ResponseFailuresInnerCode.md) |  | [optional] 

## Methods

### NewSendConversions200ResponseFailuresInner

`func NewSendConversions200ResponseFailuresInner() *SendConversions200ResponseFailuresInner`

NewSendConversions200ResponseFailuresInner instantiates a new SendConversions200ResponseFailuresInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendConversions200ResponseFailuresInnerWithDefaults

`func NewSendConversions200ResponseFailuresInnerWithDefaults() *SendConversions200ResponseFailuresInner`

NewSendConversions200ResponseFailuresInnerWithDefaults instantiates a new SendConversions200ResponseFailuresInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventIndex

`func (o *SendConversions200ResponseFailuresInner) GetEventIndex() int32`

GetEventIndex returns the EventIndex field if non-nil, zero value otherwise.

### GetEventIndexOk

`func (o *SendConversions200ResponseFailuresInner) GetEventIndexOk() (*int32, bool)`

GetEventIndexOk returns a tuple with the EventIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventIndex

`func (o *SendConversions200ResponseFailuresInner) SetEventIndex(v int32)`

SetEventIndex sets EventIndex field to given value.

### HasEventIndex

`func (o *SendConversions200ResponseFailuresInner) HasEventIndex() bool`

HasEventIndex returns a boolean if a field has been set.

### GetEventId

`func (o *SendConversions200ResponseFailuresInner) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SendConversions200ResponseFailuresInner) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SendConversions200ResponseFailuresInner) SetEventId(v string)`

SetEventId sets EventId field to given value.

### HasEventId

`func (o *SendConversions200ResponseFailuresInner) HasEventId() bool`

HasEventId returns a boolean if a field has been set.

### GetMessage

`func (o *SendConversions200ResponseFailuresInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendConversions200ResponseFailuresInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendConversions200ResponseFailuresInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SendConversions200ResponseFailuresInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *SendConversions200ResponseFailuresInner) GetCode() SendConversions200ResponseFailuresInnerCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SendConversions200ResponseFailuresInner) GetCodeOk() (*SendConversions200ResponseFailuresInnerCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SendConversions200ResponseFailuresInner) SetCode(v SendConversions200ResponseFailuresInnerCode)`

SetCode sets Code field to given value.

### HasCode

`func (o *SendConversions200ResponseFailuresInner) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


