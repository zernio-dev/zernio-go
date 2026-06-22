# GetGoogleBusinessServices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**Services** | Pointer to [**[]GetGoogleBusinessServices200ResponseServicesInner**](GetGoogleBusinessServices200ResponseServicesInner.md) |  | [optional] 

## Methods

### NewGetGoogleBusinessServices200Response

`func NewGetGoogleBusinessServices200Response() *GetGoogleBusinessServices200Response`

NewGetGoogleBusinessServices200Response instantiates a new GetGoogleBusinessServices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessServices200ResponseWithDefaults

`func NewGetGoogleBusinessServices200ResponseWithDefaults() *GetGoogleBusinessServices200Response`

NewGetGoogleBusinessServices200ResponseWithDefaults instantiates a new GetGoogleBusinessServices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGoogleBusinessServices200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGoogleBusinessServices200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGoogleBusinessServices200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGoogleBusinessServices200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGoogleBusinessServices200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGoogleBusinessServices200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGoogleBusinessServices200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGoogleBusinessServices200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetGoogleBusinessServices200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetGoogleBusinessServices200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetGoogleBusinessServices200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetGoogleBusinessServices200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetServices

`func (o *GetGoogleBusinessServices200Response) GetServices() []GetGoogleBusinessServices200ResponseServicesInner`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *GetGoogleBusinessServices200Response) GetServicesOk() (*[]GetGoogleBusinessServices200ResponseServicesInner, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *GetGoogleBusinessServices200Response) SetServices(v []GetGoogleBusinessServices200ResponseServicesInner)`

SetServices sets Services field to given value.

### HasServices

`func (o *GetGoogleBusinessServices200Response) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


