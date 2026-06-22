# GetInboxTopAccounts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Accounts** | Pointer to [**[]GetInboxTopAccounts200ResponseAccountsInner**](GetInboxTopAccounts200ResponseAccountsInner.md) |  | [optional] 

## Methods

### NewGetInboxTopAccounts200Response

`func NewGetInboxTopAccounts200Response() *GetInboxTopAccounts200Response`

NewGetInboxTopAccounts200Response instantiates a new GetInboxTopAccounts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxTopAccounts200ResponseWithDefaults

`func NewGetInboxTopAccounts200ResponseWithDefaults() *GetInboxTopAccounts200Response`

NewGetInboxTopAccounts200ResponseWithDefaults instantiates a new GetInboxTopAccounts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetInboxTopAccounts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetInboxTopAccounts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetInboxTopAccounts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetInboxTopAccounts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFrom

`func (o *GetInboxTopAccounts200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxTopAccounts200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxTopAccounts200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxTopAccounts200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetInboxTopAccounts200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetInboxTopAccounts200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetInboxTopAccounts200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetInboxTopAccounts200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *GetInboxTopAccounts200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetInboxTopAccounts200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetAccounts

`func (o *GetInboxTopAccounts200Response) GetAccounts() []GetInboxTopAccounts200ResponseAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetInboxTopAccounts200Response) GetAccountsOk() (*[]GetInboxTopAccounts200ResponseAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetInboxTopAccounts200Response) SetAccounts(v []GetInboxTopAccounts200ResponseAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetInboxTopAccounts200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


