# UpdateGoogleBusinessFoodMenusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Menus** | [**[]FoodMenu**](FoodMenu.md) | Array of food menus to set | 
**UpdateMask** | Pointer to **string** | Field mask for partial updates (e.g. \&quot;menus\&quot;) | [optional] 

## Methods

### NewUpdateGoogleBusinessFoodMenusRequest

`func NewUpdateGoogleBusinessFoodMenusRequest(menus []FoodMenu, ) *UpdateGoogleBusinessFoodMenusRequest`

NewUpdateGoogleBusinessFoodMenusRequest instantiates a new UpdateGoogleBusinessFoodMenusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoogleBusinessFoodMenusRequestWithDefaults

`func NewUpdateGoogleBusinessFoodMenusRequestWithDefaults() *UpdateGoogleBusinessFoodMenusRequest`

NewUpdateGoogleBusinessFoodMenusRequestWithDefaults instantiates a new UpdateGoogleBusinessFoodMenusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMenus

`func (o *UpdateGoogleBusinessFoodMenusRequest) GetMenus() []FoodMenu`

GetMenus returns the Menus field if non-nil, zero value otherwise.

### GetMenusOk

`func (o *UpdateGoogleBusinessFoodMenusRequest) GetMenusOk() (*[]FoodMenu, bool)`

GetMenusOk returns a tuple with the Menus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMenus

`func (o *UpdateGoogleBusinessFoodMenusRequest) SetMenus(v []FoodMenu)`

SetMenus sets Menus field to given value.


### GetUpdateMask

`func (o *UpdateGoogleBusinessFoodMenusRequest) GetUpdateMask() string`

GetUpdateMask returns the UpdateMask field if non-nil, zero value otherwise.

### GetUpdateMaskOk

`func (o *UpdateGoogleBusinessFoodMenusRequest) GetUpdateMaskOk() (*string, bool)`

GetUpdateMaskOk returns a tuple with the UpdateMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMask

`func (o *UpdateGoogleBusinessFoodMenusRequest) SetUpdateMask(v string)`

SetUpdateMask sets UpdateMask field to given value.

### HasUpdateMask

`func (o *UpdateGoogleBusinessFoodMenusRequest) HasUpdateMask() bool`

HasUpdateMask returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


