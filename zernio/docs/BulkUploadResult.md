# BulkUploadResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Number of data rows processed from the CSV | [optional] 
**Valid** | Pointer to **int32** | Count of rows that succeeded (results[].ok &#x3D;&#x3D;&#x3D; true) | [optional] 
**Invalid** | Pointer to **int32** | Count of rows that failed (total - valid) | [optional] 
**Results** | Pointer to [**[]BulkUploadResultResultsInner**](BulkUploadResultResultsInner.md) | One entry per CSV data row, in row order. | [optional] 
**Warnings** | Pointer to **[]string** | Top-level advisory warnings (e.g. &#x60;rows_exceed_advisory_limit:500&#x60;). Empty when none. | [optional] 
**RateLimitedAccounts** | Pointer to [**[]BulkUploadResultRateLimitedAccountsInner**](BulkUploadResultRateLimitedAccountsInner.md) | Present only when one or more rows targeted an account currently in cooldown. Lets callers map &#x60;rate_limited:*&#x60; row errors back to structured metadata without parsing the error strings.  | [optional] 

## Methods

### NewBulkUploadResult

`func NewBulkUploadResult() *BulkUploadResult`

NewBulkUploadResult instantiates a new BulkUploadResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUploadResultWithDefaults

`func NewBulkUploadResultWithDefaults() *BulkUploadResult`

NewBulkUploadResultWithDefaults instantiates a new BulkUploadResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *BulkUploadResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkUploadResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkUploadResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BulkUploadResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValid

`func (o *BulkUploadResult) GetValid() int32`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *BulkUploadResult) GetValidOk() (*int32, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *BulkUploadResult) SetValid(v int32)`

SetValid sets Valid field to given value.

### HasValid

`func (o *BulkUploadResult) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetInvalid

`func (o *BulkUploadResult) GetInvalid() int32`

GetInvalid returns the Invalid field if non-nil, zero value otherwise.

### GetInvalidOk

`func (o *BulkUploadResult) GetInvalidOk() (*int32, bool)`

GetInvalidOk returns a tuple with the Invalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalid

`func (o *BulkUploadResult) SetInvalid(v int32)`

SetInvalid sets Invalid field to given value.

### HasInvalid

`func (o *BulkUploadResult) HasInvalid() bool`

HasInvalid returns a boolean if a field has been set.

### GetResults

`func (o *BulkUploadResult) GetResults() []BulkUploadResultResultsInner`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *BulkUploadResult) GetResultsOk() (*[]BulkUploadResultResultsInner, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *BulkUploadResult) SetResults(v []BulkUploadResultResultsInner)`

SetResults sets Results field to given value.

### HasResults

`func (o *BulkUploadResult) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetWarnings

`func (o *BulkUploadResult) GetWarnings() []string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *BulkUploadResult) GetWarningsOk() (*[]string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *BulkUploadResult) SetWarnings(v []string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *BulkUploadResult) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetRateLimitedAccounts

`func (o *BulkUploadResult) GetRateLimitedAccounts() []BulkUploadResultRateLimitedAccountsInner`

GetRateLimitedAccounts returns the RateLimitedAccounts field if non-nil, zero value otherwise.

### GetRateLimitedAccountsOk

`func (o *BulkUploadResult) GetRateLimitedAccountsOk() (*[]BulkUploadResultRateLimitedAccountsInner, bool)`

GetRateLimitedAccountsOk returns a tuple with the RateLimitedAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimitedAccounts

`func (o *BulkUploadResult) SetRateLimitedAccounts(v []BulkUploadResultRateLimitedAccountsInner)`

SetRateLimitedAccounts sets RateLimitedAccounts field to given value.

### HasRateLimitedAccounts

`func (o *BulkUploadResult) HasRateLimitedAccounts() bool`

HasRateLimitedAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


