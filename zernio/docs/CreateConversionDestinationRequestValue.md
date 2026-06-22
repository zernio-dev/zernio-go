# CreateConversionDestinationRequestValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | **string** | ISO 4217 (e.g. \&quot;USD\&quot;). | 
**Amount** | **string** | Decimal string (e.g. \&quot;49.99\&quot;). | 

## Methods

### NewCreateConversionDestinationRequestValue

`func NewCreateConversionDestinationRequestValue(currencyCode string, amount string, ) *CreateConversionDestinationRequestValue`

NewCreateConversionDestinationRequestValue instantiates a new CreateConversionDestinationRequestValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConversionDestinationRequestValueWithDefaults

`func NewCreateConversionDestinationRequestValueWithDefaults() *CreateConversionDestinationRequestValue`

NewCreateConversionDestinationRequestValueWithDefaults instantiates a new CreateConversionDestinationRequestValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *CreateConversionDestinationRequestValue) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *CreateConversionDestinationRequestValue) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *CreateConversionDestinationRequestValue) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.


### GetAmount

`func (o *CreateConversionDestinationRequestValue) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateConversionDestinationRequestValue) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateConversionDestinationRequestValue) SetAmount(v string)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


