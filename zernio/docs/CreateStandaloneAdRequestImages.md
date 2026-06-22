# CreateStandaloneAdRequestImages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Landscape** | Pointer to **string** | Landscape 1.91:1 marketing image URL (e.g. 1200x628). Also accepted via the top-level &#x60;imageUrl&#x60; for backward compatibility. | [optional] 
**Square** | Pointer to **string** | Square 1:1 marketing image URL (e.g. 1080x1080). Required for Google Display. | [optional] 

## Methods

### NewCreateStandaloneAdRequestImages

`func NewCreateStandaloneAdRequestImages() *CreateStandaloneAdRequestImages`

NewCreateStandaloneAdRequestImages instantiates a new CreateStandaloneAdRequestImages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestImagesWithDefaults

`func NewCreateStandaloneAdRequestImagesWithDefaults() *CreateStandaloneAdRequestImages`

NewCreateStandaloneAdRequestImagesWithDefaults instantiates a new CreateStandaloneAdRequestImages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLandscape

`func (o *CreateStandaloneAdRequestImages) GetLandscape() string`

GetLandscape returns the Landscape field if non-nil, zero value otherwise.

### GetLandscapeOk

`func (o *CreateStandaloneAdRequestImages) GetLandscapeOk() (*string, bool)`

GetLandscapeOk returns a tuple with the Landscape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandscape

`func (o *CreateStandaloneAdRequestImages) SetLandscape(v string)`

SetLandscape sets Landscape field to given value.

### HasLandscape

`func (o *CreateStandaloneAdRequestImages) HasLandscape() bool`

HasLandscape returns a boolean if a field has been set.

### GetSquare

`func (o *CreateStandaloneAdRequestImages) GetSquare() string`

GetSquare returns the Square field if non-nil, zero value otherwise.

### GetSquareOk

`func (o *CreateStandaloneAdRequestImages) GetSquareOk() (*string, bool)`

GetSquareOk returns a tuple with the Square field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSquare

`func (o *CreateStandaloneAdRequestImages) SetSquare(v string)`

SetSquare sets Square field to given value.

### HasSquare

`func (o *CreateStandaloneAdRequestImages) HasSquare() bool`

HasSquare returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


