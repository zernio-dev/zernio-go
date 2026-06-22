# FoodMenuItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [**[]FoodMenuLabel**](FoodMenuLabel.md) |  | 
**Attributes** | Pointer to [**FoodMenuItemAttributes**](FoodMenuItemAttributes.md) |  | [optional] 
**Options** | Pointer to [**[]FoodMenuItemOptionsInner**](FoodMenuItemOptionsInner.md) | Item variants/options (e.g. sizes, preparations) | [optional] 

## Methods

### NewFoodMenuItem

`func NewFoodMenuItem(labels []FoodMenuLabel, ) *FoodMenuItem`

NewFoodMenuItem instantiates a new FoodMenuItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoodMenuItemWithDefaults

`func NewFoodMenuItemWithDefaults() *FoodMenuItem`

NewFoodMenuItemWithDefaults instantiates a new FoodMenuItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *FoodMenuItem) GetLabels() []FoodMenuLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FoodMenuItem) GetLabelsOk() (*[]FoodMenuLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FoodMenuItem) SetLabels(v []FoodMenuLabel)`

SetLabels sets Labels field to given value.


### GetAttributes

`func (o *FoodMenuItem) GetAttributes() FoodMenuItemAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *FoodMenuItem) GetAttributesOk() (*FoodMenuItemAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *FoodMenuItem) SetAttributes(v FoodMenuItemAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *FoodMenuItem) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetOptions

`func (o *FoodMenuItem) GetOptions() []FoodMenuItemOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FoodMenuItem) GetOptionsOk() (*[]FoodMenuItemOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FoodMenuItem) SetOptions(v []FoodMenuItemOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FoodMenuItem) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


