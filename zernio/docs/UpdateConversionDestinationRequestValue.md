# UpdateConversionDestinationRequestValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyCode** | Pointer to **string** | ISO 4217. | [optional] 
**Amount** | Pointer to **string** | Decimal string (e.g. \&quot;49.99\&quot;). | [optional] 

## Methods

### NewUpdateConversionDestinationRequestValue

`func NewUpdateConversionDestinationRequestValue() *UpdateConversionDestinationRequestValue`

NewUpdateConversionDestinationRequestValue instantiates a new UpdateConversionDestinationRequestValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConversionDestinationRequestValueWithDefaults

`func NewUpdateConversionDestinationRequestValueWithDefaults() *UpdateConversionDestinationRequestValue`

NewUpdateConversionDestinationRequestValueWithDefaults instantiates a new UpdateConversionDestinationRequestValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyCode

`func (o *UpdateConversionDestinationRequestValue) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *UpdateConversionDestinationRequestValue) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *UpdateConversionDestinationRequestValue) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *UpdateConversionDestinationRequestValue) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetAmount

`func (o *UpdateConversionDestinationRequestValue) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *UpdateConversionDestinationRequestValue) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *UpdateConversionDestinationRequestValue) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *UpdateConversionDestinationRequestValue) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


