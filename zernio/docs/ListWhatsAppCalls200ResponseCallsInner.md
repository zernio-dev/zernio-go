# ListWhatsAppCalls200ResponseCallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**EndedAt** | Pointer to **time.Time** |  | [optional] 
**DurationSeconds** | Pointer to **int32** |  | [optional] 
**EndReason** | Pointer to **string** |  | [optional] 
**RecordingUrl** | Pointer to **string** |  | [optional] 
**Billing** | Pointer to [**ListWhatsAppCalls200ResponseCallsInnerBilling**](ListWhatsAppCalls200ResponseCallsInnerBilling.md) |  | [optional] 

## Methods

### NewListWhatsAppCalls200ResponseCallsInner

`func NewListWhatsAppCalls200ResponseCallsInner() *ListWhatsAppCalls200ResponseCallsInner`

NewListWhatsAppCalls200ResponseCallsInner instantiates a new ListWhatsAppCalls200ResponseCallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppCalls200ResponseCallsInnerWithDefaults

`func NewListWhatsAppCalls200ResponseCallsInnerWithDefaults() *ListWhatsAppCalls200ResponseCallsInner`

NewListWhatsAppCalls200ResponseCallsInnerWithDefaults instantiates a new ListWhatsAppCalls200ResponseCallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDirection

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFrom

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetStatus

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartedAt

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetDurationSeconds

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetEndReason

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetEndReason() string`

GetEndReason returns the EndReason field if non-nil, zero value otherwise.

### GetEndReasonOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetEndReasonOk() (*string, bool)`

GetEndReasonOk returns a tuple with the EndReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndReason

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetEndReason(v string)`

SetEndReason sets EndReason field to given value.

### HasEndReason

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasEndReason() bool`

HasEndReason returns a boolean if a field has been set.

### GetRecordingUrl

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetRecordingUrl() string`

GetRecordingUrl returns the RecordingUrl field if non-nil, zero value otherwise.

### GetRecordingUrlOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetRecordingUrlOk() (*string, bool)`

GetRecordingUrlOk returns a tuple with the RecordingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingUrl

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetRecordingUrl(v string)`

SetRecordingUrl sets RecordingUrl field to given value.

### HasRecordingUrl

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasRecordingUrl() bool`

HasRecordingUrl returns a boolean if a field has been set.

### GetBilling

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetBilling() ListWhatsAppCalls200ResponseCallsInnerBilling`

GetBilling returns the Billing field if non-nil, zero value otherwise.

### GetBillingOk

`func (o *ListWhatsAppCalls200ResponseCallsInner) GetBillingOk() (*ListWhatsAppCalls200ResponseCallsInnerBilling, bool)`

GetBillingOk returns a tuple with the Billing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilling

`func (o *ListWhatsAppCalls200ResponseCallsInner) SetBilling(v ListWhatsAppCalls200ResponseCallsInnerBilling)`

SetBilling sets Billing field to given value.

### HasBilling

`func (o *ListWhatsAppCalls200ResponseCallsInner) HasBilling() bool`

HasBilling returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


