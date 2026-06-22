# BulkUploadResultResultsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowIndex** | Pointer to **int32** | 1-based index of the CSV data row (header excluded) | [optional] 
**Ok** | Pointer to **bool** | Whether the row was created successfully | [optional] 
**CreatedPostId** | Pointer to **string** | ID of the created post. Present only when &#x60;ok&#x60; is true and not a dry run. | [optional] 
**Errors** | Pointer to **[]string** | Machine-readable failure codes for this row. Present only when &#x60;ok&#x60; is false. Examples: &#x60;unknown_profile:&lt;id&gt;&#x60;, &#x60;no_account_for_platform:&lt;platform&gt;&#x60;, &#x60;schedule_time_missing&#x60;, &#x60;rate_limited:&lt;platform&gt;:@&lt;username&gt;:&lt;remaining&gt;&#x60;.  | [optional] 

## Methods

### NewBulkUploadResultResultsInner

`func NewBulkUploadResultResultsInner() *BulkUploadResultResultsInner`

NewBulkUploadResultResultsInner instantiates a new BulkUploadResultResultsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkUploadResultResultsInnerWithDefaults

`func NewBulkUploadResultResultsInnerWithDefaults() *BulkUploadResultResultsInner`

NewBulkUploadResultResultsInnerWithDefaults instantiates a new BulkUploadResultResultsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowIndex

`func (o *BulkUploadResultResultsInner) GetRowIndex() int32`

GetRowIndex returns the RowIndex field if non-nil, zero value otherwise.

### GetRowIndexOk

`func (o *BulkUploadResultResultsInner) GetRowIndexOk() (*int32, bool)`

GetRowIndexOk returns a tuple with the RowIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowIndex

`func (o *BulkUploadResultResultsInner) SetRowIndex(v int32)`

SetRowIndex sets RowIndex field to given value.

### HasRowIndex

`func (o *BulkUploadResultResultsInner) HasRowIndex() bool`

HasRowIndex returns a boolean if a field has been set.

### GetOk

`func (o *BulkUploadResultResultsInner) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *BulkUploadResultResultsInner) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *BulkUploadResultResultsInner) SetOk(v bool)`

SetOk sets Ok field to given value.

### HasOk

`func (o *BulkUploadResultResultsInner) HasOk() bool`

HasOk returns a boolean if a field has been set.

### GetCreatedPostId

`func (o *BulkUploadResultResultsInner) GetCreatedPostId() string`

GetCreatedPostId returns the CreatedPostId field if non-nil, zero value otherwise.

### GetCreatedPostIdOk

`func (o *BulkUploadResultResultsInner) GetCreatedPostIdOk() (*string, bool)`

GetCreatedPostIdOk returns a tuple with the CreatedPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedPostId

`func (o *BulkUploadResultResultsInner) SetCreatedPostId(v string)`

SetCreatedPostId sets CreatedPostId field to given value.

### HasCreatedPostId

`func (o *BulkUploadResultResultsInner) HasCreatedPostId() bool`

HasCreatedPostId returns a boolean if a field has been set.

### GetErrors

`func (o *BulkUploadResultResultsInner) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkUploadResultResultsInner) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkUploadResultResultsInner) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkUploadResultResultsInner) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


