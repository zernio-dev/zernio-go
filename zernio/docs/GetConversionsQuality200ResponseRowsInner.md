# GetConversionsQuality200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | Pointer to **string** |  | [optional] 
**CompositeScore** | Pointer to **float32** | Composite EMQ score, 0-10. | [optional] 
**MatchKeys** | Pointer to [**[]GetConversionsQuality200ResponseRowsInnerMatchKeysInner**](GetConversionsQuality200ResponseRowsInnerMatchKeysInner.md) |  | [optional] 
**EventCoveragePercentage** | Pointer to **float32** | Pixel↔CAPI coverage rate for this event. | [optional] 

## Methods

### NewGetConversionsQuality200ResponseRowsInner

`func NewGetConversionsQuality200ResponseRowsInner() *GetConversionsQuality200ResponseRowsInner`

NewGetConversionsQuality200ResponseRowsInner instantiates a new GetConversionsQuality200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConversionsQuality200ResponseRowsInnerWithDefaults

`func NewGetConversionsQuality200ResponseRowsInnerWithDefaults() *GetConversionsQuality200ResponseRowsInner`

NewGetConversionsQuality200ResponseRowsInnerWithDefaults instantiates a new GetConversionsQuality200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *GetConversionsQuality200ResponseRowsInner) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *GetConversionsQuality200ResponseRowsInner) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *GetConversionsQuality200ResponseRowsInner) SetEventName(v string)`

SetEventName sets EventName field to given value.

### HasEventName

`func (o *GetConversionsQuality200ResponseRowsInner) HasEventName() bool`

HasEventName returns a boolean if a field has been set.

### GetCompositeScore

`func (o *GetConversionsQuality200ResponseRowsInner) GetCompositeScore() float32`

GetCompositeScore returns the CompositeScore field if non-nil, zero value otherwise.

### GetCompositeScoreOk

`func (o *GetConversionsQuality200ResponseRowsInner) GetCompositeScoreOk() (*float32, bool)`

GetCompositeScoreOk returns a tuple with the CompositeScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompositeScore

`func (o *GetConversionsQuality200ResponseRowsInner) SetCompositeScore(v float32)`

SetCompositeScore sets CompositeScore field to given value.

### HasCompositeScore

`func (o *GetConversionsQuality200ResponseRowsInner) HasCompositeScore() bool`

HasCompositeScore returns a boolean if a field has been set.

### GetMatchKeys

`func (o *GetConversionsQuality200ResponseRowsInner) GetMatchKeys() []GetConversionsQuality200ResponseRowsInnerMatchKeysInner`

GetMatchKeys returns the MatchKeys field if non-nil, zero value otherwise.

### GetMatchKeysOk

`func (o *GetConversionsQuality200ResponseRowsInner) GetMatchKeysOk() (*[]GetConversionsQuality200ResponseRowsInnerMatchKeysInner, bool)`

GetMatchKeysOk returns a tuple with the MatchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchKeys

`func (o *GetConversionsQuality200ResponseRowsInner) SetMatchKeys(v []GetConversionsQuality200ResponseRowsInnerMatchKeysInner)`

SetMatchKeys sets MatchKeys field to given value.

### HasMatchKeys

`func (o *GetConversionsQuality200ResponseRowsInner) HasMatchKeys() bool`

HasMatchKeys returns a boolean if a field has been set.

### GetEventCoveragePercentage

`func (o *GetConversionsQuality200ResponseRowsInner) GetEventCoveragePercentage() float32`

GetEventCoveragePercentage returns the EventCoveragePercentage field if non-nil, zero value otherwise.

### GetEventCoveragePercentageOk

`func (o *GetConversionsQuality200ResponseRowsInner) GetEventCoveragePercentageOk() (*float32, bool)`

GetEventCoveragePercentageOk returns a tuple with the EventCoveragePercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCoveragePercentage

`func (o *GetConversionsQuality200ResponseRowsInner) SetEventCoveragePercentage(v float32)`

SetEventCoveragePercentage sets EventCoveragePercentage field to given value.

### HasEventCoveragePercentage

`func (o *GetConversionsQuality200ResponseRowsInner) HasEventCoveragePercentage() bool`

HasEventCoveragePercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


