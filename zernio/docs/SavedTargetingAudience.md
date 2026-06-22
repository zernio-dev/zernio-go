# SavedTargetingAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**AccountId** | **string** | Social account ID on the target ad platform. | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Spec** | [**NullableTargetingSpec**](TargetingSpec.md) | The targeting spec to store. | 

## Methods

### NewSavedTargetingAudience

`func NewSavedTargetingAudience(type_ string, accountId string, name string, spec NullableTargetingSpec, ) *SavedTargetingAudience`

NewSavedTargetingAudience instantiates a new SavedTargetingAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedTargetingAudienceWithDefaults

`func NewSavedTargetingAudienceWithDefaults() *SavedTargetingAudience`

NewSavedTargetingAudienceWithDefaults instantiates a new SavedTargetingAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SavedTargetingAudience) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SavedTargetingAudience) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SavedTargetingAudience) SetType(v string)`

SetType sets Type field to given value.


### GetAccountId

`func (o *SavedTargetingAudience) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SavedTargetingAudience) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SavedTargetingAudience) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *SavedTargetingAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SavedTargetingAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SavedTargetingAudience) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SavedTargetingAudience) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SavedTargetingAudience) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SavedTargetingAudience) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SavedTargetingAudience) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSpec

`func (o *SavedTargetingAudience) GetSpec() TargetingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SavedTargetingAudience) GetSpecOk() (*TargetingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SavedTargetingAudience) SetSpec(v TargetingSpec)`

SetSpec sets Spec field to given value.


### SetSpecNil

`func (o *SavedTargetingAudience) SetSpecNil(b bool)`

 SetSpecNil sets the value for Spec to be an explicit nil

### UnsetSpec
`func (o *SavedTargetingAudience) UnsetSpec()`

UnsetSpec ensures that no value is present for Spec, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


