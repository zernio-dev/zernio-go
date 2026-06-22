# GetGoogleBusinessAttributes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**Attributes** | Pointer to [**[]GetGoogleBusinessAttributes200ResponseAttributesInner**](GetGoogleBusinessAttributes200ResponseAttributesInner.md) |  | [optional] 

## Methods

### NewGetGoogleBusinessAttributes200Response

`func NewGetGoogleBusinessAttributes200Response() *GetGoogleBusinessAttributes200Response`

NewGetGoogleBusinessAttributes200Response instantiates a new GetGoogleBusinessAttributes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessAttributes200ResponseWithDefaults

`func NewGetGoogleBusinessAttributes200ResponseWithDefaults() *GetGoogleBusinessAttributes200Response`

NewGetGoogleBusinessAttributes200ResponseWithDefaults instantiates a new GetGoogleBusinessAttributes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGoogleBusinessAttributes200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGoogleBusinessAttributes200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGoogleBusinessAttributes200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGoogleBusinessAttributes200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGoogleBusinessAttributes200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGoogleBusinessAttributes200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGoogleBusinessAttributes200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGoogleBusinessAttributes200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetGoogleBusinessAttributes200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetGoogleBusinessAttributes200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetGoogleBusinessAttributes200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetGoogleBusinessAttributes200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetAttributes

`func (o *GetGoogleBusinessAttributes200Response) GetAttributes() []GetGoogleBusinessAttributes200ResponseAttributesInner`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GetGoogleBusinessAttributes200Response) GetAttributesOk() (*[]GetGoogleBusinessAttributes200ResponseAttributesInner, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GetGoogleBusinessAttributes200Response) SetAttributes(v []GetGoogleBusinessAttributes200ResponseAttributesInner)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GetGoogleBusinessAttributes200Response) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


