# GetAllAccountsHealth200ResponseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int32** | Total number of accounts | [optional] 
**Healthy** | Pointer to **int32** | Number of healthy accounts | [optional] 
**Warning** | Pointer to **int32** | Number of accounts with warnings | [optional] 
**Error** | Pointer to **int32** | Number of accounts with errors | [optional] 
**NeedsReconnect** | Pointer to **int32** | Number of accounts needing reconnection | [optional] 

## Methods

### NewGetAllAccountsHealth200ResponseSummary

`func NewGetAllAccountsHealth200ResponseSummary() *GetAllAccountsHealth200ResponseSummary`

NewGetAllAccountsHealth200ResponseSummary instantiates a new GetAllAccountsHealth200ResponseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllAccountsHealth200ResponseSummaryWithDefaults

`func NewGetAllAccountsHealth200ResponseSummaryWithDefaults() *GetAllAccountsHealth200ResponseSummary`

NewGetAllAccountsHealth200ResponseSummaryWithDefaults instantiates a new GetAllAccountsHealth200ResponseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetAllAccountsHealth200ResponseSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAllAccountsHealth200ResponseSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAllAccountsHealth200ResponseSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAllAccountsHealth200ResponseSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetHealthy

`func (o *GetAllAccountsHealth200ResponseSummary) GetHealthy() int32`

GetHealthy returns the Healthy field if non-nil, zero value otherwise.

### GetHealthyOk

`func (o *GetAllAccountsHealth200ResponseSummary) GetHealthyOk() (*int32, bool)`

GetHealthyOk returns a tuple with the Healthy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthy

`func (o *GetAllAccountsHealth200ResponseSummary) SetHealthy(v int32)`

SetHealthy sets Healthy field to given value.

### HasHealthy

`func (o *GetAllAccountsHealth200ResponseSummary) HasHealthy() bool`

HasHealthy returns a boolean if a field has been set.

### GetWarning

`func (o *GetAllAccountsHealth200ResponseSummary) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *GetAllAccountsHealth200ResponseSummary) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *GetAllAccountsHealth200ResponseSummary) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *GetAllAccountsHealth200ResponseSummary) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetError

`func (o *GetAllAccountsHealth200ResponseSummary) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetAllAccountsHealth200ResponseSummary) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetAllAccountsHealth200ResponseSummary) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *GetAllAccountsHealth200ResponseSummary) HasError() bool`

HasError returns a boolean if a field has been set.

### GetNeedsReconnect

`func (o *GetAllAccountsHealth200ResponseSummary) GetNeedsReconnect() int32`

GetNeedsReconnect returns the NeedsReconnect field if non-nil, zero value otherwise.

### GetNeedsReconnectOk

`func (o *GetAllAccountsHealth200ResponseSummary) GetNeedsReconnectOk() (*int32, bool)`

GetNeedsReconnectOk returns a tuple with the NeedsReconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsReconnect

`func (o *GetAllAccountsHealth200ResponseSummary) SetNeedsReconnect(v int32)`

SetNeedsReconnect sets NeedsReconnect field to given value.

### HasNeedsReconnect

`func (o *GetAllAccountsHealth200ResponseSummary) HasNeedsReconnect() bool`

HasNeedsReconnect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


