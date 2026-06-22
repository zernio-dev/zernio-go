# UpdateGmbLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedLocationId** | **string** |  | 
**GoogleAccountId** | Pointer to **string** | Optional but recommended. The Google Business Account resource name (\&quot;accounts/123\&quot;) that owns the new location (from GET gmb-locations). When provided, the location is resolved directly instead of by enumerating the account, which is required for accounts with many locations. Named &#x60;googleAccountId&#x60; to disambiguate from the path &#x60;accountId&#x60; (the Zernio account). The legacy field name &#x60;accountId&#x60; is still accepted for backwards compatibility.  | [optional] 

## Methods

### NewUpdateGmbLocationRequest

`func NewUpdateGmbLocationRequest(selectedLocationId string, ) *UpdateGmbLocationRequest`

NewUpdateGmbLocationRequest instantiates a new UpdateGmbLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGmbLocationRequestWithDefaults

`func NewUpdateGmbLocationRequestWithDefaults() *UpdateGmbLocationRequest`

NewUpdateGmbLocationRequestWithDefaults instantiates a new UpdateGmbLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedLocationId

`func (o *UpdateGmbLocationRequest) GetSelectedLocationId() string`

GetSelectedLocationId returns the SelectedLocationId field if non-nil, zero value otherwise.

### GetSelectedLocationIdOk

`func (o *UpdateGmbLocationRequest) GetSelectedLocationIdOk() (*string, bool)`

GetSelectedLocationIdOk returns a tuple with the SelectedLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedLocationId

`func (o *UpdateGmbLocationRequest) SetSelectedLocationId(v string)`

SetSelectedLocationId sets SelectedLocationId field to given value.


### GetGoogleAccountId

`func (o *UpdateGmbLocationRequest) GetGoogleAccountId() string`

GetGoogleAccountId returns the GoogleAccountId field if non-nil, zero value otherwise.

### GetGoogleAccountIdOk

`func (o *UpdateGmbLocationRequest) GetGoogleAccountIdOk() (*string, bool)`

GetGoogleAccountIdOk returns a tuple with the GoogleAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleAccountId

`func (o *UpdateGmbLocationRequest) SetGoogleAccountId(v string)`

SetGoogleAccountId sets GoogleAccountId field to given value.

### HasGoogleAccountId

`func (o *UpdateGmbLocationRequest) HasGoogleAccountId() bool`

HasGoogleAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


