# Money

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | ISO 4217 currency code (e.g. USD, EUR) | 
**Units** | **string** | Whole units of the amount | 
**Nanos** | Pointer to **int32** | Nano units (10^-9) of the amount | [optional] 

## Methods

### NewMoney

`func NewMoney(currencyCode string, units string, ) *Money`

NewMoney instantiates a new Money object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyWithDefaults

`func NewMoneyWithDefaults() *Money`

NewMoneyWithDefaults instantiates a new Money object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *Money) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *Money) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *Money) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetUnits

`func (o *Money) GetUnits() string`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *Money) GetUnitsOk() (*string, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *Money) SetUnits(v string)`

SetUnits sets Units field to given value.


### GetNanos

`func (o *Money) GetNanos() int32`

GetNanos returns the Nanos field if non-nil, zero value otherwise.

### GetNanosOk

`func (o *Money) GetNanosOk() (*int32, bool)`

GetNanosOk returns a tuple with the Nanos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNanos

`func (o *Money) SetNanos(v int32)`

SetNanos sets Nanos field to given value.

### HasNanos

`func (o *Money) HasNanos() bool`

HasNanos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


