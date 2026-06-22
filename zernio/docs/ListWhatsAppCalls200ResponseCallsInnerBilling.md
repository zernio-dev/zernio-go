# ListWhatsAppCalls200ResponseCallsInnerBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaCostUSD** | Pointer to **float32** | Meta per-minute charge, billed by Meta directly to your WABA. Display only; not billed by Zernio. | [optional] 
**TelnyxCostUSD** | Pointer to **float32** |  | [optional] 
**RecordingCostUSD** | Pointer to **float32** |  | [optional] 
**BillableCostUSD** | Pointer to **float32** | Amount Zernio bills you &#x3D; Telnyx leg + recording (excludes Meta). | [optional] 
**TotalCostUSD** | Pointer to **float32** | Full cost incl. the Meta portion you pay directly. Display only. | [optional] 
**Currency** | Pointer to **string** |  | [optional] 

## Methods

### NewListWhatsAppCalls200ResponseCallsInnerBilling

`func NewListWhatsAppCalls200ResponseCallsInnerBilling() *ListWhatsAppCalls200ResponseCallsInnerBilling`

NewListWhatsAppCalls200ResponseCallsInnerBilling instantiates a new ListWhatsAppCalls200ResponseCallsInnerBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppCalls200ResponseCallsInnerBillingWithDefaults

`func NewListWhatsAppCalls200ResponseCallsInnerBillingWithDefaults() *ListWhatsAppCalls200ResponseCallsInnerBilling`

NewListWhatsAppCalls200ResponseCallsInnerBillingWithDefaults instantiates a new ListWhatsAppCalls200ResponseCallsInnerBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetMetaCostUSD() float32`

GetMetaCostUSD returns the MetaCostUSD field if non-nil, zero value otherwise.

### GetMetaCostUSDOk

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetMetaCostUSDOk() (*float32, bool)`

GetMetaCostUSDOk returns a tuple with the MetaCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) SetMetaCostUSD(v float32)`

SetMetaCostUSD sets MetaCostUSD field to given value.

### HasMetaCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) HasMetaCostUSD() bool`

HasMetaCostUSD returns a boolean if a field has been set.

### GetTelnyxCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetTelnyxCostUSD() float32`

GetTelnyxCostUSD returns the TelnyxCostUSD field if non-nil, zero value otherwise.

### GetTelnyxCostUSDOk

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetTelnyxCostUSDOk() (*float32, bool)`

GetTelnyxCostUSDOk returns a tuple with the TelnyxCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) SetTelnyxCostUSD(v float32)`

SetTelnyxCostUSD sets TelnyxCostUSD field to given value.

### HasTelnyxCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) HasTelnyxCostUSD() bool`

HasTelnyxCostUSD returns a boolean if a field has been set.

### GetRecordingCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetRecordingCostUSD() float32`

GetRecordingCostUSD returns the RecordingCostUSD field if non-nil, zero value otherwise.

### GetRecordingCostUSDOk

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetRecordingCostUSDOk() (*float32, bool)`

GetRecordingCostUSDOk returns a tuple with the RecordingCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) SetRecordingCostUSD(v float32)`

SetRecordingCostUSD sets RecordingCostUSD field to given value.

### HasRecordingCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) HasRecordingCostUSD() bool`

HasRecordingCostUSD returns a boolean if a field has been set.

### GetBillableCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetBillableCostUSD() float32`

GetBillableCostUSD returns the BillableCostUSD field if non-nil, zero value otherwise.

### GetBillableCostUSDOk

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetBillableCostUSDOk() (*float32, bool)`

GetBillableCostUSDOk returns a tuple with the BillableCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) SetBillableCostUSD(v float32)`

SetBillableCostUSD sets BillableCostUSD field to given value.

### HasBillableCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) HasBillableCostUSD() bool`

HasBillableCostUSD returns a boolean if a field has been set.

### GetTotalCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetTotalCostUSD() float32`

GetTotalCostUSD returns the TotalCostUSD field if non-nil, zero value otherwise.

### GetTotalCostUSDOk

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetTotalCostUSDOk() (*float32, bool)`

GetTotalCostUSDOk returns a tuple with the TotalCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) SetTotalCostUSD(v float32)`

SetTotalCostUSD sets TotalCostUSD field to given value.

### HasTotalCostUSD

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) HasTotalCostUSD() bool`

HasTotalCostUSD returns a boolean if a field has been set.

### GetCurrency

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListWhatsAppCalls200ResponseCallsInnerBilling) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


