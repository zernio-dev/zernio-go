# FoodMenuItemAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Price** | Pointer to [**Money**](Money.md) |  | [optional] 
**Spiciness** | Pointer to **string** | Spiciness level (e.g. MILD, MEDIUM, HOT) | [optional] 
**Allergen** | Pointer to **[]string** | Allergens (e.g. DAIRY, GLUTEN, SHELLFISH) | [optional] 
**DietaryRestriction** | Pointer to **[]string** | Dietary labels (e.g. VEGETARIAN, VEGAN, GLUTEN_FREE) | [optional] 
**ServesNumPeople** | Pointer to **int32** | Number of people the item serves | [optional] 
**PreparationMethods** | Pointer to **[]string** | Preparation methods (e.g. GRILLED, FRIED) | [optional] 
**MediaKeys** | Pointer to **[]string** | Media references for item photos | [optional] 

## Methods

### NewFoodMenuItemAttributes

`func NewFoodMenuItemAttributes() *FoodMenuItemAttributes`

NewFoodMenuItemAttributes instantiates a new FoodMenuItemAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoodMenuItemAttributesWithDefaults

`func NewFoodMenuItemAttributesWithDefaults() *FoodMenuItemAttributes`

NewFoodMenuItemAttributesWithDefaults instantiates a new FoodMenuItemAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrice

`func (o *FoodMenuItemAttributes) GetPrice() Money`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *FoodMenuItemAttributes) GetPriceOk() (*Money, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *FoodMenuItemAttributes) SetPrice(v Money)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *FoodMenuItemAttributes) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetSpiciness

`func (o *FoodMenuItemAttributes) GetSpiciness() string`

GetSpiciness returns the Spiciness field if non-nil, zero value otherwise.

### GetSpicinessOk

`func (o *FoodMenuItemAttributes) GetSpicinessOk() (*string, bool)`

GetSpicinessOk returns a tuple with the Spiciness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpiciness

`func (o *FoodMenuItemAttributes) SetSpiciness(v string)`

SetSpiciness sets Spiciness field to given value.

### HasSpiciness

`func (o *FoodMenuItemAttributes) HasSpiciness() bool`

HasSpiciness returns a boolean if a field has been set.

### GetAllergen

`func (o *FoodMenuItemAttributes) GetAllergen() []string`

GetAllergen returns the Allergen field if non-nil, zero value otherwise.

### GetAllergenOk

`func (o *FoodMenuItemAttributes) GetAllergenOk() (*[]string, bool)`

GetAllergenOk returns a tuple with the Allergen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllergen

`func (o *FoodMenuItemAttributes) SetAllergen(v []string)`

SetAllergen sets Allergen field to given value.

### HasAllergen

`func (o *FoodMenuItemAttributes) HasAllergen() bool`

HasAllergen returns a boolean if a field has been set.

### GetDietaryRestriction

`func (o *FoodMenuItemAttributes) GetDietaryRestriction() []string`

GetDietaryRestriction returns the DietaryRestriction field if non-nil, zero value otherwise.

### GetDietaryRestrictionOk

`func (o *FoodMenuItemAttributes) GetDietaryRestrictionOk() (*[]string, bool)`

GetDietaryRestrictionOk returns a tuple with the DietaryRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDietaryRestriction

`func (o *FoodMenuItemAttributes) SetDietaryRestriction(v []string)`

SetDietaryRestriction sets DietaryRestriction field to given value.

### HasDietaryRestriction

`func (o *FoodMenuItemAttributes) HasDietaryRestriction() bool`

HasDietaryRestriction returns a boolean if a field has been set.

### GetServesNumPeople

`func (o *FoodMenuItemAttributes) GetServesNumPeople() int32`

GetServesNumPeople returns the ServesNumPeople field if non-nil, zero value otherwise.

### GetServesNumPeopleOk

`func (o *FoodMenuItemAttributes) GetServesNumPeopleOk() (*int32, bool)`

GetServesNumPeopleOk returns a tuple with the ServesNumPeople field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServesNumPeople

`func (o *FoodMenuItemAttributes) SetServesNumPeople(v int32)`

SetServesNumPeople sets ServesNumPeople field to given value.

### HasServesNumPeople

`func (o *FoodMenuItemAttributes) HasServesNumPeople() bool`

HasServesNumPeople returns a boolean if a field has been set.

### GetPreparationMethods

`func (o *FoodMenuItemAttributes) GetPreparationMethods() []string`

GetPreparationMethods returns the PreparationMethods field if non-nil, zero value otherwise.

### GetPreparationMethodsOk

`func (o *FoodMenuItemAttributes) GetPreparationMethodsOk() (*[]string, bool)`

GetPreparationMethodsOk returns a tuple with the PreparationMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparationMethods

`func (o *FoodMenuItemAttributes) SetPreparationMethods(v []string)`

SetPreparationMethods sets PreparationMethods field to given value.

### HasPreparationMethods

`func (o *FoodMenuItemAttributes) HasPreparationMethods() bool`

HasPreparationMethods returns a boolean if a field has been set.

### GetMediaKeys

`func (o *FoodMenuItemAttributes) GetMediaKeys() []string`

GetMediaKeys returns the MediaKeys field if non-nil, zero value otherwise.

### GetMediaKeysOk

`func (o *FoodMenuItemAttributes) GetMediaKeysOk() (*[]string, bool)`

GetMediaKeysOk returns a tuple with the MediaKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaKeys

`func (o *FoodMenuItemAttributes) SetMediaKeys(v []string)`

SetMediaKeys sets MediaKeys field to given value.

### HasMediaKeys

`func (o *FoodMenuItemAttributes) HasMediaKeys() bool`

HasMediaKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


