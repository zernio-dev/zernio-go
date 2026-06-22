# InitiateWhatsAppCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**To** | **string** | Consumer wa_id (E.164, leading + optional) | 
**Action** | Pointer to **string** | Omit to place a call. Set to send the consent prompt instead. | [optional] 
**BodyText** | Pointer to **string** | Body text shown with the consent prompt (send_call_permission_request only). | [optional] 
**ForwardTo** | Pointer to **string** | Per-call destination override. Same accepted shape as the number&#39;s stored forwardTo (tel:+E164, sip:..., wss://...).  | [optional] 
**RecordOverride** | Pointer to **bool** |  | [optional] 
**BizOpaqueCallbackData** | Pointer to **string** | Accepted for forward compatibility. Not currently echoed back in webhook payloads (SIP-first flow does not pass through Meta&#39;s Graph API where Meta would echo this).  | [optional] 

## Methods

### NewInitiateWhatsAppCallRequest

`func NewInitiateWhatsAppCallRequest(accountId string, to string, ) *InitiateWhatsAppCallRequest`

NewInitiateWhatsAppCallRequest instantiates a new InitiateWhatsAppCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitiateWhatsAppCallRequestWithDefaults

`func NewInitiateWhatsAppCallRequestWithDefaults() *InitiateWhatsAppCallRequest`

NewInitiateWhatsAppCallRequestWithDefaults instantiates a new InitiateWhatsAppCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *InitiateWhatsAppCallRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InitiateWhatsAppCallRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InitiateWhatsAppCallRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTo

`func (o *InitiateWhatsAppCallRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *InitiateWhatsAppCallRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *InitiateWhatsAppCallRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetAction

`func (o *InitiateWhatsAppCallRequest) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InitiateWhatsAppCallRequest) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InitiateWhatsAppCallRequest) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *InitiateWhatsAppCallRequest) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBodyText

`func (o *InitiateWhatsAppCallRequest) GetBodyText() string`

GetBodyText returns the BodyText field if non-nil, zero value otherwise.

### GetBodyTextOk

`func (o *InitiateWhatsAppCallRequest) GetBodyTextOk() (*string, bool)`

GetBodyTextOk returns a tuple with the BodyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyText

`func (o *InitiateWhatsAppCallRequest) SetBodyText(v string)`

SetBodyText sets BodyText field to given value.

### HasBodyText

`func (o *InitiateWhatsAppCallRequest) HasBodyText() bool`

HasBodyText returns a boolean if a field has been set.

### GetForwardTo

`func (o *InitiateWhatsAppCallRequest) GetForwardTo() string`

GetForwardTo returns the ForwardTo field if non-nil, zero value otherwise.

### GetForwardToOk

`func (o *InitiateWhatsAppCallRequest) GetForwardToOk() (*string, bool)`

GetForwardToOk returns a tuple with the ForwardTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwardTo

`func (o *InitiateWhatsAppCallRequest) SetForwardTo(v string)`

SetForwardTo sets ForwardTo field to given value.

### HasForwardTo

`func (o *InitiateWhatsAppCallRequest) HasForwardTo() bool`

HasForwardTo returns a boolean if a field has been set.

### GetRecordOverride

`func (o *InitiateWhatsAppCallRequest) GetRecordOverride() bool`

GetRecordOverride returns the RecordOverride field if non-nil, zero value otherwise.

### GetRecordOverrideOk

`func (o *InitiateWhatsAppCallRequest) GetRecordOverrideOk() (*bool, bool)`

GetRecordOverrideOk returns a tuple with the RecordOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordOverride

`func (o *InitiateWhatsAppCallRequest) SetRecordOverride(v bool)`

SetRecordOverride sets RecordOverride field to given value.

### HasRecordOverride

`func (o *InitiateWhatsAppCallRequest) HasRecordOverride() bool`

HasRecordOverride returns a boolean if a field has been set.

### GetBizOpaqueCallbackData

`func (o *InitiateWhatsAppCallRequest) GetBizOpaqueCallbackData() string`

GetBizOpaqueCallbackData returns the BizOpaqueCallbackData field if non-nil, zero value otherwise.

### GetBizOpaqueCallbackDataOk

`func (o *InitiateWhatsAppCallRequest) GetBizOpaqueCallbackDataOk() (*string, bool)`

GetBizOpaqueCallbackDataOk returns a tuple with the BizOpaqueCallbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBizOpaqueCallbackData

`func (o *InitiateWhatsAppCallRequest) SetBizOpaqueCallbackData(v string)`

SetBizOpaqueCallbackData sets BizOpaqueCallbackData field to given value.

### HasBizOpaqueCallbackData

`func (o *InitiateWhatsAppCallRequest) HasBizOpaqueCallbackData() bool`

HasBizOpaqueCallbackData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


