# ConversionEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventName** | **string** | Standard event name (Purchase, Lead, CompleteRegistration, AddToCart, InitiateCheckout, AddPaymentInfo, Subscribe, StartTrial, ViewContent, Search, Contact, SubmitApplication, Schedule) or a custom string (only supported on platforms that accept custom events — Meta).  Per-platform behavior: - Meta: free-form; standard names match Meta&#39;s built-ins. - Google: ignored — the conversion action&#39;s category determines the type. - LinkedIn: ignored — the conversion rule&#39;s &#x60;type&#x60; is locked to the destination.  | 
**EventTime** | **int32** | When the conversion happened, in unix seconds. | 
**EventId** | **string** | Unique dedup key. The same eventId must be used on pixel + CAPI to prevent double-counting. Mapped to event_id on Meta, transactionId on Google, eventId on LinkedIn (LinkedIn deduplicates against Insight Tag events with the same eventId; the Insight Tag event wins when both arrive).  | 
**Value** | Pointer to **float32** | Conversion value in the specified currency. | [optional] 
**Currency** | Pointer to **string** | ISO 4217 currency code. | [optional] 
**User** | [**ConversionEventUser**](ConversionEventUser.md) |  | 
**Items** | Pointer to [**[]ConversionEventItemsInner**](ConversionEventItemsInner.md) | Item-level detail for ecommerce events. | [optional] 
**SourceUrl** | Pointer to **string** | URL where the conversion originated (used by Meta). | [optional] 
**ActionSource** | Pointer to **string** | Where the conversion happened. Used by Meta; Google ignores. | [optional] 
**PlatformData** | Pointer to **map[string]interface{}** | Escape hatch for platform-specific fields we haven&#39;t normalized. Forwarded as-is. | [optional] 

## Methods

### NewConversionEvent

`func NewConversionEvent(eventName string, eventTime int32, eventId string, user ConversionEventUser, ) *ConversionEvent`

NewConversionEvent instantiates a new ConversionEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConversionEventWithDefaults

`func NewConversionEventWithDefaults() *ConversionEvent`

NewConversionEventWithDefaults instantiates a new ConversionEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventName

`func (o *ConversionEvent) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *ConversionEvent) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *ConversionEvent) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *ConversionEvent) GetEventTime() int32`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *ConversionEvent) GetEventTimeOk() (*int32, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *ConversionEvent) SetEventTime(v int32)`

SetEventTime sets EventTime field to given value.


### GetEventId

`func (o *ConversionEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *ConversionEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *ConversionEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetValue

`func (o *ConversionEvent) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConversionEvent) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConversionEvent) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *ConversionEvent) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrency

`func (o *ConversionEvent) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ConversionEvent) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ConversionEvent) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ConversionEvent) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUser

`func (o *ConversionEvent) GetUser() ConversionEventUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ConversionEvent) GetUserOk() (*ConversionEventUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ConversionEvent) SetUser(v ConversionEventUser)`

SetUser sets User field to given value.


### GetItems

`func (o *ConversionEvent) GetItems() []ConversionEventItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConversionEvent) GetItemsOk() (*[]ConversionEventItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConversionEvent) SetItems(v []ConversionEventItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConversionEvent) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ConversionEvent) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ConversionEvent) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ConversionEvent) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ConversionEvent) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetActionSource

`func (o *ConversionEvent) GetActionSource() string`

GetActionSource returns the ActionSource field if non-nil, zero value otherwise.

### GetActionSourceOk

`func (o *ConversionEvent) GetActionSourceOk() (*string, bool)`

GetActionSourceOk returns a tuple with the ActionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionSource

`func (o *ConversionEvent) SetActionSource(v string)`

SetActionSource sets ActionSource field to given value.

### HasActionSource

`func (o *ConversionEvent) HasActionSource() bool`

HasActionSource returns a boolean if a field has been set.

### GetPlatformData

`func (o *ConversionEvent) GetPlatformData() map[string]interface{}`

GetPlatformData returns the PlatformData field if non-nil, zero value otherwise.

### GetPlatformDataOk

`func (o *ConversionEvent) GetPlatformDataOk() (*map[string]interface{}, bool)`

GetPlatformDataOk returns a tuple with the PlatformData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformData

`func (o *ConversionEvent) SetPlatformData(v map[string]interface{})`

SetPlatformData sets PlatformData field to given value.

### HasPlatformData

`func (o *ConversionEvent) HasPlatformData() bool`

HasPlatformData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


