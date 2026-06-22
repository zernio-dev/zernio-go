# StartGoogleBusinessVerification200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**Verification** | Pointer to [**GetGoogleBusinessVerifications200ResponseVerificationsInner**](GetGoogleBusinessVerifications200ResponseVerificationsInner.md) |  | [optional] 

## Methods

### NewStartGoogleBusinessVerification200Response

`func NewStartGoogleBusinessVerification200Response() *StartGoogleBusinessVerification200Response`

NewStartGoogleBusinessVerification200Response instantiates a new StartGoogleBusinessVerification200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartGoogleBusinessVerification200ResponseWithDefaults

`func NewStartGoogleBusinessVerification200ResponseWithDefaults() *StartGoogleBusinessVerification200Response`

NewStartGoogleBusinessVerification200ResponseWithDefaults instantiates a new StartGoogleBusinessVerification200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *StartGoogleBusinessVerification200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *StartGoogleBusinessVerification200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *StartGoogleBusinessVerification200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *StartGoogleBusinessVerification200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *StartGoogleBusinessVerification200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *StartGoogleBusinessVerification200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *StartGoogleBusinessVerification200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *StartGoogleBusinessVerification200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *StartGoogleBusinessVerification200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *StartGoogleBusinessVerification200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *StartGoogleBusinessVerification200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *StartGoogleBusinessVerification200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetVerification

`func (o *StartGoogleBusinessVerification200Response) GetVerification() GetGoogleBusinessVerifications200ResponseVerificationsInner`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *StartGoogleBusinessVerification200Response) GetVerificationOk() (*GetGoogleBusinessVerifications200ResponseVerificationsInner, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *StartGoogleBusinessVerification200Response) SetVerification(v GetGoogleBusinessVerifications200ResponseVerificationsInner)`

SetVerification sets Verification field to given value.

### HasVerification

`func (o *StartGoogleBusinessVerification200Response) HasVerification() bool`

HasVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


