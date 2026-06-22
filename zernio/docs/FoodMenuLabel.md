# FoodMenuLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** | Display name of the item/section/menu | 
**Description** | Pointer to **string** | Optional description | [optional] 
**LanguageCode** | Pointer to **string** | BCP-47 language code (e.g. en, es) | [optional] 

## Methods

### NewFoodMenuLabel

`func NewFoodMenuLabel(displayName string, ) *FoodMenuLabel`

NewFoodMenuLabel instantiates a new FoodMenuLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoodMenuLabelWithDefaults

`func NewFoodMenuLabelWithDefaults() *FoodMenuLabel`

NewFoodMenuLabelWithDefaults instantiates a new FoodMenuLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *FoodMenuLabel) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FoodMenuLabel) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FoodMenuLabel) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *FoodMenuLabel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FoodMenuLabel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FoodMenuLabel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FoodMenuLabel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLanguageCode

`func (o *FoodMenuLabel) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *FoodMenuLabel) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *FoodMenuLabel) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *FoodMenuLabel) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


