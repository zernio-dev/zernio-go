# InitiateWhatsAppCall200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**CallId** | Pointer to **string** | Internal Call doc ID | [optional] 
**TelnyxCallControlId** | Pointer to **string** | Telnyx call_control_id of the outbound leg | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 
**ForwardTo** | Pointer to **NullableString** |  | [optional] 
**RecordingEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewInitiateWhatsAppCall200Response

`func NewInitiateWhatsAppCall200Response() *InitiateWhatsAppCall200Response`

NewInitiateWhatsAppCall200Response instantiates a new InitiateWhatsAppCall200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateWhatsAppCall200ResponseWithDefaults

`func NewInitiateWhatsAppCall200ResponseWithDefaults() *InitiateWhatsAppCall200Response`

NewInitiateWhatsAppCall200ResponseWithDefaults instantiates a new InitiateWhatsAppCall200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InitiateWhatsAppCall200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InitiateWhatsAppCall200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InitiateWhatsAppCall200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InitiateWhatsAppCall200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCallId

`func (o *InitiateWhatsAppCall200Response) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *InitiateWhatsAppCall200Response) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *InitiateWhatsAppCall200Response) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *InitiateWhatsAppCall200Response) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetTelnyxCallControlId

`func (o *InitiateWhatsAppCall200Response) GetTelnyxCallControlId() string`

GetTelnyxCallControlId returns the TelnyxCallControlId field if non-nil, zero value otherwise.

### GetTelnyxCallControlIdOk

`func (o *InitiateWhatsAppCall200Response) GetTelnyxCallControlIdOk() (*string, bool)`

GetTelnyxCallControlIdOk returns a tuple with the TelnyxCallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxCallControlId

`func (o *InitiateWhatsAppCall200Response) SetTelnyxCallControlId(v string)`

SetTelnyxCallControlId sets TelnyxCallControlId field to given value.

### HasTelnyxCallControlId

`func (o *InitiateWhatsAppCall200Response) HasTelnyxCallControlId() bool`

HasTelnyxCallControlId returns a boolean if a field has been set.

### GetStatus

`func (o *InitiateWhatsAppCall200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InitiateWhatsAppCall200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InitiateWhatsAppCall200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InitiateWhatsAppCall200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDirection

`func (o *InitiateWhatsAppCall200Response) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *InitiateWhatsAppCall200Response) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *InitiateWhatsAppCall200Response) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *InitiateWhatsAppCall200Response) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTo

`func (o *InitiateWhatsAppCall200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InitiateWhatsAppCall200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InitiateWhatsAppCall200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *InitiateWhatsAppCall200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetForwardTo

`func (o *InitiateWhatsAppCall200Response) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *InitiateWhatsAppCall200Response) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *InitiateWhatsAppCall200Response) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *InitiateWhatsAppCall200Response) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### SetForwardToNil

`func (o *InitiateWhatsAppCall200Response) SetForwardToNil(b bool)`

 SetForwardToNil sets the value for ForwardTo to be an explicit nil

### UnsetForwardTo
`func (o *InitiateWhatsAppCall200Response) UnsetForwardTo()`

UnsetForwardTo ensures that no value is present for ForwardTo, not even an explicit nil
### GetRecordingEnabled

`func (o *InitiateWhatsAppCall200Response) GetRecordingEnabled() bool`

GetRecordingEnabled returns the RecordingEnabled field if non-nil, zero value otherwise.

### GetRecordingEnabledOk

`func (o *InitiateWhatsAppCall200Response) GetRecordingEnabledOk() (*bool, bool)`

GetRecordingEnabledOk returns a tuple with the RecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingEnabled

`func (o *InitiateWhatsAppCall200Response) SetRecordingEnabled(v bool)`

SetRecordingEnabled sets RecordingEnabled field to given value.

### HasRecordingEnabled

`func (o *InitiateWhatsAppCall200Response) HasRecordingEnabled() bool`

HasRecordingEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


