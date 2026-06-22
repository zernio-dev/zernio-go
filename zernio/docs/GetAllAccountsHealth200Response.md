# GetAllAccountsHealth200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**GetAllAccountsHealth200ResponseSummary**](GetAllAccountsHealth200ResponseSummary.md) |  | [optional] 
**Accounts** | Pointer to [**[]GetAllAccountsHealth200ResponseAccountsInner**](GetAllAccountsHealth200ResponseAccountsInner.md) |  | [optional] 

## Methods

### NewGetAllAccountsHealth200Response

`func NewGetAllAccountsHealth200Response() *GetAllAccountsHealth200Response`

NewGetAllAccountsHealth200Response instantiates a new GetAllAccountsHealth200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllAccountsHealth200ResponseWithDefaults

`func NewGetAllAccountsHealth200ResponseWithDefaults() *GetAllAccountsHealth200Response`

NewGetAllAccountsHealth200ResponseWithDefaults instantiates a new GetAllAccountsHealth200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *GetAllAccountsHealth200Response) GetSummary() GetAllAccountsHealth200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetAllAccountsHealth200Response) GetSummaryOk() (*GetAllAccountsHealth200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetAllAccountsHealth200Response) SetSummary(v GetAllAccountsHealth200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetAllAccountsHealth200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetAccounts

`func (o *GetAllAccountsHealth200Response) GetAccounts() []GetAllAccountsHealth200ResponseAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetAllAccountsHealth200Response) GetAccountsOk() (*[]GetAllAccountsHealth200ResponseAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetAllAccountsHealth200Response) SetAccounts(v []GetAllAccountsHealth200ResponseAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetAllAccountsHealth200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


