# FoodMenu

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | [**[]FoodMenuLabel**](FoodMenuLabel.md) |  | 
**Sections** | Pointer to [**[]FoodMenuSection**](FoodMenuSection.md) |  | [optional] 
**Cuisines** | Pointer to **[]string** | Cuisine types (e.g. AMERICAN, ITALIAN, JAPANESE) | [optional] 
**SourceUrl** | Pointer to **string** | URL of the original menu source | [optional] 

## Methods

### NewFoodMenu

`func NewFoodMenu(labels []FoodMenuLabel, ) *FoodMenu`

NewFoodMenu instantiates a new FoodMenu object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoodMenuWithDefaults

`func NewFoodMenuWithDefaults() *FoodMenu`

NewFoodMenuWithDefaults instantiates a new FoodMenu object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *FoodMenu) GetLabels() []FoodMenuLabel`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *FoodMenu) GetLabelsOk() (*[]FoodMenuLabel, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *FoodMenu) SetLabels(v []FoodMenuLabel)`

SetLabels sets Labels field to given value.


### GetSections

`func (o *FoodMenu) GetSections() []FoodMenuSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *FoodMenu) GetSectionsOk() (*[]FoodMenuSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *FoodMenu) SetSections(v []FoodMenuSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *FoodMenu) HasSections() bool`

HasSections returns a boolean if a field has been set.

### GetCuisines

`func (o *FoodMenu) GetCuisines() []string`

GetCuisines returns the Cuisines field if non-nil, zero value otherwise.

### GetCuisinesOk

`func (o *FoodMenu) GetCuisinesOk() (*[]string, bool)`

GetCuisinesOk returns a tuple with the Cuisines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuisines

`func (o *FoodMenu) SetCuisines(v []string)`

SetCuisines sets Cuisines field to given value.

### HasCuisines

`func (o *FoodMenu) HasCuisines() bool`

HasCuisines returns a boolean if a field has been set.

### GetSourceUrl

`func (o *FoodMenu) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *FoodMenu) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *FoodMenu) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *FoodMenu) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


