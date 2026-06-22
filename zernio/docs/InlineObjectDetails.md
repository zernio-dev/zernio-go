# InlineObjectDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeTierAccountLimit** | Pointer to **int32** | How many accounts the free tier allows. Only set when reason&#x3D;free_tier_exceeded. | [optional] 
**CurrentAccountCount** | Pointer to **int32** | How many accounts the team currently has connected. Set when reason&#x3D;free_tier_exceeded or reason&#x3D;enterprise_required. | [optional] 
**HasPaymentMethod** | Pointer to **bool** | Whether the team currently has a card on file in Stripe. Set when reason&#x3D;free_tier_exceeded or reason&#x3D;twitter_passthrough. | [optional] 
**PublicAccountLimit** | Pointer to **int32** | Public pricing ceiling (the published cap beyond which an enterprise contract is required). Only set when reason&#x3D;enterprise_required. | [optional] 
**EffectiveAccountLimit** | Pointer to **int32** | The cap actually applied to this team. Equals &#x60;public_account_limit&#x60; for organic teams; for teams with a per-customer override (grandfathered legacy customers, signed enterprise contracts) this can be higher. Only set when reason&#x3D;enterprise_required.  | [optional] 

## Methods

### NewInlineObjectDetails

`func NewInlineObjectDetails() *InlineObjectDetails`

NewInlineObjectDetails instantiates a new InlineObjectDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObjectDetailsWithDefaults

`func NewInlineObjectDetailsWithDefaults() *InlineObjectDetails`

NewInlineObjectDetailsWithDefaults instantiates a new InlineObjectDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeTierAccountLimit

`func (o *InlineObjectDetails) GetFreeTierAccountLimit() int32`

GetFreeTierAccountLimit returns the FreeTierAccountLimit field if non-nil, zero value otherwise.

### GetFreeTierAccountLimitOk

`func (o *InlineObjectDetails) GetFreeTierAccountLimitOk() (*int32, bool)`

GetFreeTierAccountLimitOk returns a tuple with the FreeTierAccountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeTierAccountLimit

`func (o *InlineObjectDetails) SetFreeTierAccountLimit(v int32)`

SetFreeTierAccountLimit sets FreeTierAccountLimit field to given value.

### HasFreeTierAccountLimit

`func (o *InlineObjectDetails) HasFreeTierAccountLimit() bool`

HasFreeTierAccountLimit returns a boolean if a field has been set.

### GetCurrentAccountCount

`func (o *InlineObjectDetails) GetCurrentAccountCount() int32`

GetCurrentAccountCount returns the CurrentAccountCount field if non-nil, zero value otherwise.

### GetCurrentAccountCountOk

`func (o *InlineObjectDetails) GetCurrentAccountCountOk() (*int32, bool)`

GetCurrentAccountCountOk returns a tuple with the CurrentAccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAccountCount

`func (o *InlineObjectDetails) SetCurrentAccountCount(v int32)`

SetCurrentAccountCount sets CurrentAccountCount field to given value.

### HasCurrentAccountCount

`func (o *InlineObjectDetails) HasCurrentAccountCount() bool`

HasCurrentAccountCount returns a boolean if a field has been set.

### GetHasPaymentMethod

`func (o *InlineObjectDetails) GetHasPaymentMethod() bool`

GetHasPaymentMethod returns the HasPaymentMethod field if non-nil, zero value otherwise.

### GetHasPaymentMethodOk

`func (o *InlineObjectDetails) GetHasPaymentMethodOk() (*bool, bool)`

GetHasPaymentMethodOk returns a tuple with the HasPaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPaymentMethod

`func (o *InlineObjectDetails) SetHasPaymentMethod(v bool)`

SetHasPaymentMethod sets HasPaymentMethod field to given value.

### HasHasPaymentMethod

`func (o *InlineObjectDetails) HasHasPaymentMethod() bool`

HasHasPaymentMethod returns a boolean if a field has been set.

### GetPublicAccountLimit

`func (o *InlineObjectDetails) GetPublicAccountLimit() int32`

GetPublicAccountLimit returns the PublicAccountLimit field if non-nil, zero value otherwise.

### GetPublicAccountLimitOk

`func (o *InlineObjectDetails) GetPublicAccountLimitOk() (*int32, bool)`

GetPublicAccountLimitOk returns a tuple with the PublicAccountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicAccountLimit

`func (o *InlineObjectDetails) SetPublicAccountLimit(v int32)`

SetPublicAccountLimit sets PublicAccountLimit field to given value.

### HasPublicAccountLimit

`func (o *InlineObjectDetails) HasPublicAccountLimit() bool`

HasPublicAccountLimit returns a boolean if a field has been set.

### GetEffectiveAccountLimit

`func (o *InlineObjectDetails) GetEffectiveAccountLimit() int32`

GetEffectiveAccountLimit returns the EffectiveAccountLimit field if non-nil, zero value otherwise.

### GetEffectiveAccountLimitOk

`func (o *InlineObjectDetails) GetEffectiveAccountLimitOk() (*int32, bool)`

GetEffectiveAccountLimitOk returns a tuple with the EffectiveAccountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveAccountLimit

`func (o *InlineObjectDetails) SetEffectiveAccountLimit(v int32)`

SetEffectiveAccountLimit sets EffectiveAccountLimit field to given value.

### HasEffectiveAccountLimit

`func (o *InlineObjectDetails) HasEffectiveAccountLimit() bool`

HasEffectiveAccountLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


