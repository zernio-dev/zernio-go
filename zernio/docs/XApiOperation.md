# XApiOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Operation** | Pointer to **string** | Internal operation key. Matches keys in &#x60;xApiCallsByOperation&#x60;. | [optional] 
**EventType** | Pointer to **string** | Metronome &#x60;event_type&#x60; emitted when this operation runs. | [optional] 
**DisplayName** | Pointer to **string** | Human-readable label shown on Metronome invoices. | [optional] 
**PricePerCallUsd** | Pointer to **float32** |  | [optional] 
**PricePerCallCents** | Pointer to **float32** | Per-call price in cents. Fractional values are intentional. | [optional] 
**Tier** | Pointer to **string** | Tier key derived from &#x60;pricePerCallUsd&#x60; (e.g. &#x60;x_api_005&#x60; for $0.005, &#x60;x_api_200&#x60; for $0.200). Useful for grouping operations by price in dashboards.  | [optional] 
**TriggeredBy** | Pointer to [**[]XApiOperationTriggeredByInner**](XApiOperationTriggeredByInner.md) | Zernio platform methods that emit this operation, with their metering rule. | [optional] 

## Methods

### NewXApiOperation

`func NewXApiOperation() *XApiOperation`

NewXApiOperation instantiates a new XApiOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXApiOperationWithDefaults

`func NewXApiOperationWithDefaults() *XApiOperation`

NewXApiOperationWithDefaults instantiates a new XApiOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperation

`func (o *XApiOperation) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *XApiOperation) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *XApiOperation) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *XApiOperation) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetEventType

`func (o *XApiOperation) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *XApiOperation) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *XApiOperation) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *XApiOperation) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDisplayName

`func (o *XApiOperation) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *XApiOperation) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *XApiOperation) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *XApiOperation) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetPricePerCallUsd

`func (o *XApiOperation) GetPricePerCallUsd() float32`

GetPricePerCallUsd returns the PricePerCallUsd field if non-nil, zero value otherwise.

### GetPricePerCallUsdOk

`func (o *XApiOperation) GetPricePerCallUsdOk() (*float32, bool)`

GetPricePerCallUsdOk returns a tuple with the PricePerCallUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerCallUsd

`func (o *XApiOperation) SetPricePerCallUsd(v float32)`

SetPricePerCallUsd sets PricePerCallUsd field to given value.

### HasPricePerCallUsd

`func (o *XApiOperation) HasPricePerCallUsd() bool`

HasPricePerCallUsd returns a boolean if a field has been set.

### GetPricePerCallCents

`func (o *XApiOperation) GetPricePerCallCents() float32`

GetPricePerCallCents returns the PricePerCallCents field if non-nil, zero value otherwise.

### GetPricePerCallCentsOk

`func (o *XApiOperation) GetPricePerCallCentsOk() (*float32, bool)`

GetPricePerCallCentsOk returns a tuple with the PricePerCallCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerCallCents

`func (o *XApiOperation) SetPricePerCallCents(v float32)`

SetPricePerCallCents sets PricePerCallCents field to given value.

### HasPricePerCallCents

`func (o *XApiOperation) HasPricePerCallCents() bool`

HasPricePerCallCents returns a boolean if a field has been set.

### GetTier

`func (o *XApiOperation) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *XApiOperation) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *XApiOperation) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *XApiOperation) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTriggeredBy

`func (o *XApiOperation) GetTriggeredBy() []XApiOperationTriggeredByInner`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *XApiOperation) GetTriggeredByOk() (*[]XApiOperationTriggeredByInner, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *XApiOperation) SetTriggeredBy(v []XApiOperationTriggeredByInner)`

SetTriggeredBy sets TriggeredBy field to given value.

### HasTriggeredBy

`func (o *XApiOperation) HasTriggeredBy() bool`

HasTriggeredBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


