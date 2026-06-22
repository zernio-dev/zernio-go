# GetGoogleBusinessVerifications200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**VoiceOfMerchantState** | Pointer to [**GetGoogleBusinessVerifications200ResponseVoiceOfMerchantState**](GetGoogleBusinessVerifications200ResponseVoiceOfMerchantState.md) |  | [optional] 
**Verifications** | Pointer to [**[]GetGoogleBusinessVerifications200ResponseVerificationsInner**](GetGoogleBusinessVerifications200ResponseVerificationsInner.md) | Verification history, newest first. Empty when none exist. | [optional] 

## Methods

### NewGetGoogleBusinessVerifications200Response

`func NewGetGoogleBusinessVerifications200Response() *GetGoogleBusinessVerifications200Response`

NewGetGoogleBusinessVerifications200Response instantiates a new GetGoogleBusinessVerifications200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessVerifications200ResponseWithDefaults

`func NewGetGoogleBusinessVerifications200ResponseWithDefaults() *GetGoogleBusinessVerifications200Response`

NewGetGoogleBusinessVerifications200ResponseWithDefaults instantiates a new GetGoogleBusinessVerifications200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGoogleBusinessVerifications200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGoogleBusinessVerifications200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGoogleBusinessVerifications200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGoogleBusinessVerifications200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGoogleBusinessVerifications200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGoogleBusinessVerifications200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGoogleBusinessVerifications200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGoogleBusinessVerifications200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetGoogleBusinessVerifications200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetGoogleBusinessVerifications200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetGoogleBusinessVerifications200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetGoogleBusinessVerifications200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetVoiceOfMerchantState

`func (o *GetGoogleBusinessVerifications200Response) GetVoiceOfMerchantState() GetGoogleBusinessVerifications200ResponseVoiceOfMerchantState`

GetVoiceOfMerchantState returns the VoiceOfMerchantState field if non-nil, zero value otherwise.

### GetVoiceOfMerchantStateOk

`func (o *GetGoogleBusinessVerifications200Response) GetVoiceOfMerchantStateOk() (*GetGoogleBusinessVerifications200ResponseVoiceOfMerchantState, bool)`

GetVoiceOfMerchantStateOk returns a tuple with the VoiceOfMerchantState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceOfMerchantState

`func (o *GetGoogleBusinessVerifications200Response) SetVoiceOfMerchantState(v GetGoogleBusinessVerifications200ResponseVoiceOfMerchantState)`

SetVoiceOfMerchantState sets VoiceOfMerchantState field to given value.

### HasVoiceOfMerchantState

`func (o *GetGoogleBusinessVerifications200Response) HasVoiceOfMerchantState() bool`

HasVoiceOfMerchantState returns a boolean if a field has been set.

### GetVerifications

`func (o *GetGoogleBusinessVerifications200Response) GetVerifications() []GetGoogleBusinessVerifications200ResponseVerificationsInner`

GetVerifications returns the Verifications field if non-nil, zero value otherwise.

### GetVerificationsOk

`func (o *GetGoogleBusinessVerifications200Response) GetVerificationsOk() (*[]GetGoogleBusinessVerifications200ResponseVerificationsInner, bool)`

GetVerificationsOk returns a tuple with the Verifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifications

`func (o *GetGoogleBusinessVerifications200Response) SetVerifications(v []GetGoogleBusinessVerifications200ResponseVerificationsInner)`

SetVerifications sets Verifications field to given value.

### HasVerifications

`func (o *GetGoogleBusinessVerifications200Response) HasVerifications() bool`

HasVerifications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


