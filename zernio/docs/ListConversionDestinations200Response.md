# ListConversionDestinations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Destinations** | Pointer to [**[]ListConversionDestinations200ResponseDestinationsInner**](ListConversionDestinations200ResponseDestinationsInner.md) |  | [optional] 

## Methods

### NewListConversionDestinations200Response

`func NewListConversionDestinations200Response() *ListConversionDestinations200Response`

NewListConversionDestinations200Response instantiates a new ListConversionDestinations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConversionDestinations200ResponseWithDefaults

`func NewListConversionDestinations200ResponseWithDefaults() *ListConversionDestinations200Response`

NewListConversionDestinations200ResponseWithDefaults instantiates a new ListConversionDestinations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *ListConversionDestinations200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListConversionDestinations200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListConversionDestinations200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListConversionDestinations200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDestinations

`func (o *ListConversionDestinations200Response) GetDestinations() []ListConversionDestinations200ResponseDestinationsInner`

GetDestinations returns the Destinations field if non-nil, zero value otherwise.

### GetDestinationsOk

`func (o *ListConversionDestinations200Response) GetDestinationsOk() (*[]ListConversionDestinations200ResponseDestinationsInner, bool)`

GetDestinationsOk returns a tuple with the Destinations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinations

`func (o *ListConversionDestinations200Response) SetDestinations(v []ListConversionDestinations200ResponseDestinationsInner)`

SetDestinations sets Destinations field to given value.

### HasDestinations

`func (o *ListConversionDestinations200Response) HasDestinations() bool`

HasDestinations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


