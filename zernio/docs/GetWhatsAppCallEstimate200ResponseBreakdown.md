# GetWhatsAppCallEstimate200ResponseBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetaMinutes** | Pointer to **int32** |  | [optional] 
**MetaCostUSD** | Pointer to **float32** | Estimated Meta per-minute charge, billed by Meta directly to your WABA. Display only; not billed by Zernio. | [optional] 
**TelnyxCostUSD** | Pointer to **float32** |  | [optional] 
**RecordingCostUSD** | Pointer to **float32** |  | [optional] 
**BillableCostUSD** | Pointer to **float32** | Estimated amount Zernio bills you &#x3D; Telnyx leg + recording (excludes Meta). | [optional] 
**TotalCostUSD** | Pointer to **float32** | Estimated full cost incl. the Meta portion you pay directly. Display only. | [optional] 

## Methods

### NewGetWhatsAppCallEstimate200ResponseBreakdown

`func NewGetWhatsAppCallEstimate200ResponseBreakdown() *GetWhatsAppCallEstimate200ResponseBreakdown`

NewGetWhatsAppCallEstimate200ResponseBreakdown instantiates a new GetWhatsAppCallEstimate200ResponseBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppCallEstimate200ResponseBreakdownWithDefaults

`func NewGetWhatsAppCallEstimate200ResponseBreakdownWithDefaults() *GetWhatsAppCallEstimate200ResponseBreakdown`

NewGetWhatsAppCallEstimate200ResponseBreakdownWithDefaults instantiates a new GetWhatsAppCallEstimate200ResponseBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetaMinutes

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetMetaMinutes() int32`

GetMetaMinutes returns the MetaMinutes field if non-nil, zero value otherwise.

### GetMetaMinutesOk

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetMetaMinutesOk() (*int32, bool)`

GetMetaMinutesOk returns a tuple with the MetaMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaMinutes

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) SetMetaMinutes(v int32)`

SetMetaMinutes sets MetaMinutes field to given value.

### HasMetaMinutes

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) HasMetaMinutes() bool`

HasMetaMinutes returns a boolean if a field has been set.

### GetMetaCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetMetaCostUSD() float32`

GetMetaCostUSD returns the MetaCostUSD field if non-nil, zero value otherwise.

### GetMetaCostUSDOk

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetMetaCostUSDOk() (*float32, bool)`

GetMetaCostUSDOk returns a tuple with the MetaCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetaCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) SetMetaCostUSD(v float32)`

SetMetaCostUSD sets MetaCostUSD field to given value.

### HasMetaCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) HasMetaCostUSD() bool`

HasMetaCostUSD returns a boolean if a field has been set.

### GetTelnyxCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetTelnyxCostUSD() float32`

GetTelnyxCostUSD returns the TelnyxCostUSD field if non-nil, zero value otherwise.

### GetTelnyxCostUSDOk

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetTelnyxCostUSDOk() (*float32, bool)`

GetTelnyxCostUSDOk returns a tuple with the TelnyxCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelnyxCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) SetTelnyxCostUSD(v float32)`

SetTelnyxCostUSD sets TelnyxCostUSD field to given value.

### HasTelnyxCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) HasTelnyxCostUSD() bool`

HasTelnyxCostUSD returns a boolean if a field has been set.

### GetRecordingCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetRecordingCostUSD() float32`

GetRecordingCostUSD returns the RecordingCostUSD field if non-nil, zero value otherwise.

### GetRecordingCostUSDOk

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetRecordingCostUSDOk() (*float32, bool)`

GetRecordingCostUSDOk returns a tuple with the RecordingCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordingCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) SetRecordingCostUSD(v float32)`

SetRecordingCostUSD sets RecordingCostUSD field to given value.

### HasRecordingCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) HasRecordingCostUSD() bool`

HasRecordingCostUSD returns a boolean if a field has been set.

### GetBillableCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetBillableCostUSD() float32`

GetBillableCostUSD returns the BillableCostUSD field if non-nil, zero value otherwise.

### GetBillableCostUSDOk

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetBillableCostUSDOk() (*float32, bool)`

GetBillableCostUSDOk returns a tuple with the BillableCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillableCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) SetBillableCostUSD(v float32)`

SetBillableCostUSD sets BillableCostUSD field to given value.

### HasBillableCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) HasBillableCostUSD() bool`

HasBillableCostUSD returns a boolean if a field has been set.

### GetTotalCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetTotalCostUSD() float32`

GetTotalCostUSD returns the TotalCostUSD field if non-nil, zero value otherwise.

### GetTotalCostUSDOk

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) GetTotalCostUSDOk() (*float32, bool)`

GetTotalCostUSDOk returns a tuple with the TotalCostUSD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) SetTotalCostUSD(v float32)`

SetTotalCostUSD sets TotalCostUSD field to given value.

### HasTotalCostUSD

`func (o *GetWhatsAppCallEstimate200ResponseBreakdown) HasTotalCostUSD() bool`

HasTotalCostUSD returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


