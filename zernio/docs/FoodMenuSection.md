# FoodMenuSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [**[]FoodMenuLabel**](FoodMenuLabel.md) |  | 
**Items** | Pointer to [**[]FoodMenuItem**](FoodMenuItem.md) |  | [optional] 

## Methods

### NewFoodMenuSection

`func NewFoodMenuSection(labels []FoodMenuLabel, ) *FoodMenuSection`

NewFoodMenuSection instantiates a new FoodMenuSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoodMenuSectionWithDefaults

`func NewFoodMenuSectionWithDefaults() *FoodMenuSection`

NewFoodMenuSectionWithDefaults instantiates a new FoodMenuSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *FoodMenuSection) GetLabels() []FoodMenuLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FoodMenuSection) GetLabelsOk() (*[]FoodMenuLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FoodMenuSection) SetLabels(v []FoodMenuLabel)`

SetLabels sets Labels field to given value.


### GetItems

`func (o *FoodMenuSection) GetItems() []FoodMenuItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FoodMenuSection) GetItemsOk() (*[]FoodMenuItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FoodMenuSection) SetItems(v []FoodMenuItem)`

SetItems sets Items field to given value.

### HasItems

`func (o *FoodMenuSection) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


