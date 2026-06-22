# GetAccountHealth200ResponseTokenStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | Whether the token is valid | [optional] 
**ExpiresAt** | Pointer to **time.Time** |  | [optional] 
**ExpiresIn** | Pointer to **string** | Human-readable time until expiry | [optional] 
**NeedsRefresh** | Pointer to **bool** | Whether token expires within 24 hours | [optional] 

## Methods

### NewGetAccountHealth200ResponseTokenStatus

`func NewGetAccountHealth200ResponseTokenStatus() *GetAccountHealth200ResponseTokenStatus`

NewGetAccountHealth200ResponseTokenStatus instantiates a new GetAccountHealth200ResponseTokenStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountHealth200ResponseTokenStatusWithDefaults

`func NewGetAccountHealth200ResponseTokenStatusWithDefaults() *GetAccountHealth200ResponseTokenStatus`

NewGetAccountHealth200ResponseTokenStatusWithDefaults instantiates a new GetAccountHealth200ResponseTokenStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *GetAccountHealth200ResponseTokenStatus) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *GetAccountHealth200ResponseTokenStatus) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *GetAccountHealth200ResponseTokenStatus) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *GetAccountHealth200ResponseTokenStatus) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetExpiresAt

`func (o *GetAccountHealth200ResponseTokenStatus) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetAccountHealth200ResponseTokenStatus) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetAccountHealth200ResponseTokenStatus) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *GetAccountHealth200ResponseTokenStatus) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExpiresIn

`func (o *GetAccountHealth200ResponseTokenStatus) GetExpiresIn() string`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *GetAccountHealth200ResponseTokenStatus) GetExpiresInOk() (*string, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *GetAccountHealth200ResponseTokenStatus) SetExpiresIn(v string)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *GetAccountHealth200ResponseTokenStatus) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetNeedsRefresh

`func (o *GetAccountHealth200ResponseTokenStatus) GetNeedsRefresh() bool`

GetNeedsRefresh returns the NeedsRefresh field if non-nil, zero value otherwise.

### GetNeedsRefreshOk

`func (o *GetAccountHealth200ResponseTokenStatus) GetNeedsRefreshOk() (*bool, bool)`

GetNeedsRefreshOk returns a tuple with the NeedsRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsRefresh

`func (o *GetAccountHealth200ResponseTokenStatus) SetNeedsRefresh(v bool)`

SetNeedsRefresh sets NeedsRefresh field to given value.

### HasNeedsRefresh

`func (o *GetAccountHealth200ResponseTokenStatus) HasNeedsRefresh() bool`

HasNeedsRefresh returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


