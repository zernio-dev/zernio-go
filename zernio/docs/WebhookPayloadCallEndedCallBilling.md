# WebhookPayloadCallEndedCallBilling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaCostUSD** | Pointer to **float32** | Meta per-minute charge. Billed by Meta DIRECTLY to your WhatsApp Business Account payment method (your separate Meta invoice). Zernio does NOT charge this. Display only. | [optional] 
**TelnyxCostUSD** | Pointer to **float32** |  | [optional] 
**RecordingCostUSD** | Pointer to **float32** |  | [optional] 
**BillableCostUSD** | Pointer to **float32** | The amount Zernio bills you &#x3D; Telnyx leg + recording. Excludes Meta (billed by Meta directly). | [optional] 
**TotalCostUSD** | Pointer to **float32** | Full economic cost incl. the Meta portion you pay directly (Meta + Telnyx + recording). Display only, not the Zernio-billed amount. | [optional] 

## Methods

### NewWebhookPayloadCallEndedCallBilling

`func NewWebhookPayloadCallEndedCallBilling() *WebhookPayloadCallEndedCallBilling`

NewWebhookPayloadCallEndedCallBilling instantiates a new WebhookPayloadCallEndedCallBilling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallEndedCallBillingWithDefaults

`func NewWebhookPayloadCallEndedCallBillingWithDefaults() *WebhookPayloadCallEndedCallBilling`

NewWebhookPayloadCallEndedCallBillingWithDefaults instantiates a new WebhookPayloadCallEndedCallBilling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) GetMetaCostUSD() float32`

GetMetaCostUSD returns the MetaCostUSD field if non-nil, zero value otherwise.

### GetMetaCostUSDOk

`func (o *WebhookPayloadCallEndedCallBilling) GetMetaCostUSDOk() (*float32, bool)`

GetMetaCostUSDOk returns a tuple with the MetaCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) SetMetaCostUSD(v float32)`

SetMetaCostUSD sets MetaCostUSD field to given value.

### HasMetaCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) HasMetaCostUSD() bool`

HasMetaCostUSD returns a boolean if a field has been set.

### GetTelnyxCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) GetTelnyxCostUSD() float32`

GetTelnyxCostUSD returns the TelnyxCostUSD field if non-nil, zero value otherwise.

### GetTelnyxCostUSDOk

`func (o *WebhookPayloadCallEndedCallBilling) GetTelnyxCostUSDOk() (*float32, bool)`

GetTelnyxCostUSDOk returns a tuple with the TelnyxCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) SetTelnyxCostUSD(v float32)`

SetTelnyxCostUSD sets TelnyxCostUSD field to given value.

### HasTelnyxCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) HasTelnyxCostUSD() bool`

HasTelnyxCostUSD returns a boolean if a field has been set.

### GetRecordingCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) GetRecordingCostUSD() float32`

GetRecordingCostUSD returns the RecordingCostUSD field if non-nil, zero value otherwise.

### GetRecordingCostUSDOk

`func (o *WebhookPayloadCallEndedCallBilling) GetRecordingCostUSDOk() (*float32, bool)`

GetRecordingCostUSDOk returns a tuple with the RecordingCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) SetRecordingCostUSD(v float32)`

SetRecordingCostUSD sets RecordingCostUSD field to given value.

### HasRecordingCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) HasRecordingCostUSD() bool`

HasRecordingCostUSD returns a boolean if a field has been set.

### GetBillableCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) GetBillableCostUSD() float32`

GetBillableCostUSD returns the BillableCostUSD field if non-nil, zero value otherwise.

### GetBillableCostUSDOk

`func (o *WebhookPayloadCallEndedCallBilling) GetBillableCostUSDOk() (*float32, bool)`

GetBillableCostUSDOk returns a tuple with the BillableCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) SetBillableCostUSD(v float32)`

SetBillableCostUSD sets BillableCostUSD field to given value.

### HasBillableCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) HasBillableCostUSD() bool`

HasBillableCostUSD returns a boolean if a field has been set.

### GetTotalCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) GetTotalCostUSD() float32`

GetTotalCostUSD returns the TotalCostUSD field if non-nil, zero value otherwise.

### GetTotalCostUSDOk

`func (o *WebhookPayloadCallEndedCallBilling) GetTotalCostUSDOk() (*float32, bool)`

GetTotalCostUSDOk returns a tuple with the TotalCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) SetTotalCostUSD(v float32)`

SetTotalCostUSD sets TotalCostUSD field to given value.

### HasTotalCostUSD

`func (o *WebhookPayloadCallEndedCallBilling) HasTotalCostUSD() bool`

HasTotalCostUSD returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


