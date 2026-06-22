# GetAdAnalytics200ResponseAd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **NullableString** | ISO 4217 code of the ad account that owns this ad (e.g. USD, THB, INR). All money values in &#x60;summary&#x60; and &#x60;daily&#x60; are in this currency. Null only on legacy ads synced before currency was persisted. | [optional] 

## Methods

### NewGetAdAnalytics200ResponseAd

`func NewGetAdAnalytics200ResponseAd() *GetAdAnalytics200ResponseAd`

NewGetAdAnalytics200ResponseAd instantiates a new GetAdAnalytics200ResponseAd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdAnalytics200ResponseAdWithDefaults

`func NewGetAdAnalytics200ResponseAdWithDefaults() *GetAdAnalytics200ResponseAd`

NewGetAdAnalytics200ResponseAdWithDefaults instantiates a new GetAdAnalytics200ResponseAd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAdAnalytics200ResponseAd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAdAnalytics200ResponseAd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAdAnalytics200ResponseAd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetAdAnalytics200ResponseAd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetAdAnalytics200ResponseAd) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAdAnalytics200ResponseAd) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAdAnalytics200ResponseAd) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAdAnalytics200ResponseAd) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *GetAdAnalytics200ResponseAd) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAdAnalytics200ResponseAd) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAdAnalytics200ResponseAd) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetAdAnalytics200ResponseAd) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetTrigger

`func (o *GetAdAnalytics200ResponseAd) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *GetAdAnalytics200ResponseAd) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *GetAdAnalytics200ResponseAd) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *GetAdAnalytics200ResponseAd) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetStatus

`func (o *GetAdAnalytics200ResponseAd) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAdAnalytics200ResponseAd) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAdAnalytics200ResponseAd) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAdAnalytics200ResponseAd) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrency

`func (o *GetAdAnalytics200ResponseAd) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetAdAnalytics200ResponseAd) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetAdAnalytics200ResponseAd) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetAdAnalytics200ResponseAd) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *GetAdAnalytics200ResponseAd) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *GetAdAnalytics200ResponseAd) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


