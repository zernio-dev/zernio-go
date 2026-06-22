# GetGmbLocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]GetGmbLocations200ResponseLocationsInner**](GetGmbLocations200ResponseLocationsInner.md) |  | [optional] 
**HasMore** | Pointer to **bool** | True when more locations exist than were returned (use search to narrow down). | [optional] 
**SelectedLocationId** | Pointer to **string** |  | [optional] 
**Cached** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetGmbLocations200Response

`func NewGetGmbLocations200Response() *GetGmbLocations200Response`

NewGetGmbLocations200Response instantiates a new GetGmbLocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGmbLocations200ResponseWithDefaults

`func NewGetGmbLocations200ResponseWithDefaults() *GetGmbLocations200Response`

NewGetGmbLocations200ResponseWithDefaults instantiates a new GetGmbLocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *GetGmbLocations200Response) GetLocations() []GetGmbLocations200ResponseLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *GetGmbLocations200Response) GetLocationsOk() (*[]GetGmbLocations200ResponseLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *GetGmbLocations200Response) SetLocations(v []GetGmbLocations200ResponseLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *GetGmbLocations200Response) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetHasMore

`func (o *GetGmbLocations200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *GetGmbLocations200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *GetGmbLocations200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *GetGmbLocations200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetSelectedLocationId

`func (o *GetGmbLocations200Response) GetSelectedLocationId() string`

GetSelectedLocationId returns the SelectedLocationId field if non-nil, zero value otherwise.

### GetSelectedLocationIdOk

`func (o *GetGmbLocations200Response) GetSelectedLocationIdOk() (*string, bool)`

GetSelectedLocationIdOk returns a tuple with the SelectedLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedLocationId

`func (o *GetGmbLocations200Response) SetSelectedLocationId(v string)`

SetSelectedLocationId sets SelectedLocationId field to given value.

### HasSelectedLocationId

`func (o *GetGmbLocations200Response) HasSelectedLocationId() bool`

HasSelectedLocationId returns a boolean if a field has been set.

### GetCached

`func (o *GetGmbLocations200Response) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *GetGmbLocations200Response) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *GetGmbLocations200Response) SetCached(v bool)`

SetCached sets Cached field to given value.

### HasCached

`func (o *GetGmbLocations200Response) HasCached() bool`

HasCached returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


