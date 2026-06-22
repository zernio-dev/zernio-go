# UpdateAdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Budget** | Pointer to [**UpdateAdSetRequestBudget**](UpdateAdSetRequestBudget.md) |  | [optional] 
**Targeting** | Pointer to [**UpdateAdRequestTargeting**](UpdateAdRequestTargeting.md) |  | [optional] 
**Creative** | Pointer to [**UpdateAdRequestCreative**](UpdateAdRequestCreative.md) |  | [optional] 
**Name** | Pointer to **string** | Rename the ad. Now propagated to Meta (POST /{ad-id}); non-Meta platforms return 501. | [optional] 

## Methods

### NewUpdateAdRequest

`func NewUpdateAdRequest() *UpdateAdRequest`

NewUpdateAdRequest instantiates a new UpdateAdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdRequestWithDefaults

`func NewUpdateAdRequestWithDefaults() *UpdateAdRequest`

NewUpdateAdRequestWithDefaults instantiates a new UpdateAdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *UpdateAdRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateAdRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateAdRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateAdRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetBudget

`func (o *UpdateAdRequest) GetBudget() UpdateAdSetRequestBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *UpdateAdRequest) GetBudgetOk() (*UpdateAdSetRequestBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *UpdateAdRequest) SetBudget(v UpdateAdSetRequestBudget)`

SetBudget sets Budget field to given value.

### HasBudget

`func (o *UpdateAdRequest) HasBudget() bool`

HasBudget returns a boolean if a field has been set.

### GetTargeting

`func (o *UpdateAdRequest) GetTargeting() UpdateAdRequestTargeting`

GetTargeting returns the Targeting field if non-nil, zero value otherwise.

### GetTargetingOk

`func (o *UpdateAdRequest) GetTargetingOk() (*UpdateAdRequestTargeting, bool)`

GetTargetingOk returns a tuple with the Targeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargeting

`func (o *UpdateAdRequest) SetTargeting(v UpdateAdRequestTargeting)`

SetTargeting sets Targeting field to given value.

### HasTargeting

`func (o *UpdateAdRequest) HasTargeting() bool`

HasTargeting returns a boolean if a field has been set.

### GetCreative

`func (o *UpdateAdRequest) GetCreative() UpdateAdRequestCreative`

GetCreative returns the Creative field if non-nil, zero value otherwise.

### GetCreativeOk

`func (o *UpdateAdRequest) GetCreativeOk() (*UpdateAdRequestCreative, bool)`

GetCreativeOk returns a tuple with the Creative field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreative

`func (o *UpdateAdRequest) SetCreative(v UpdateAdRequestCreative)`

SetCreative sets Creative field to given value.

### HasCreative

`func (o *UpdateAdRequest) HasCreative() bool`

HasCreative returns a boolean if a field has been set.

### GetName

`func (o *UpdateAdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateAdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateAdRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateAdRequest) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


