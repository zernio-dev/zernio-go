# XApiOperationTriggeredByInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | Pointer to **string** | Zernio platform method name. | [optional] 
**Metering** | Pointer to **string** | When the method actually bills the user:   * &#x60;always&#x60; — every call is metered   * &#x60;analytics_optin&#x60; — only when the X account has analytics enabled   * &#x60;inbox_optin&#x60; — only when the X account has inbox sync enabled   * &#x60;absorbed&#x60; — Zernio eats the cost, never billed  | [optional] 

## Methods

### NewXApiOperationTriggeredByInner

`func NewXApiOperationTriggeredByInner() *XApiOperationTriggeredByInner`

NewXApiOperationTriggeredByInner instantiates a new XApiOperationTriggeredByInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXApiOperationTriggeredByInnerWithDefaults

`func NewXApiOperationTriggeredByInnerWithDefaults() *XApiOperationTriggeredByInner`

NewXApiOperationTriggeredByInnerWithDefaults instantiates a new XApiOperationTriggeredByInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *XApiOperationTriggeredByInner) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *XApiOperationTriggeredByInner) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *XApiOperationTriggeredByInner) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *XApiOperationTriggeredByInner) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetMetering

`func (o *XApiOperationTriggeredByInner) GetMetering() string`

GetMetering returns the Metering field if non-nil, zero value otherwise.

### GetMeteringOk

`func (o *XApiOperationTriggeredByInner) GetMeteringOk() (*string, bool)`

GetMeteringOk returns a tuple with the Metering field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetering

`func (o *XApiOperationTriggeredByInner) SetMetering(v string)`

SetMetering sets Metering field to given value.

### HasMetering

`func (o *XApiOperationTriggeredByInner) HasMetering() bool`

HasMetering returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


