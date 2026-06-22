# ListGoogleBusinessLocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Locations** | Pointer to [**[]ListGoogleBusinessLocations200ResponseLocationsInner**](ListGoogleBusinessLocations200ResponseLocationsInner.md) |  | [optional] 
**HasMore** | Pointer to **bool** | True when more locations exist than were returned (the list is bounded). Prompt the user to narrow the result set with search.  | [optional] 

## Methods

### NewListGoogleBusinessLocations200Response

`func NewListGoogleBusinessLocations200Response() *ListGoogleBusinessLocations200Response`

NewListGoogleBusinessLocations200Response instantiates a new ListGoogleBusinessLocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGoogleBusinessLocations200ResponseWithDefaults

`func NewListGoogleBusinessLocations200ResponseWithDefaults() *ListGoogleBusinessLocations200Response`

NewListGoogleBusinessLocations200ResponseWithDefaults instantiates a new ListGoogleBusinessLocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocations

`func (o *ListGoogleBusinessLocations200Response) GetLocations() []ListGoogleBusinessLocations200ResponseLocationsInner`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *ListGoogleBusinessLocations200Response) GetLocationsOk() (*[]ListGoogleBusinessLocations200ResponseLocationsInner, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *ListGoogleBusinessLocations200Response) SetLocations(v []ListGoogleBusinessLocations200ResponseLocationsInner)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *ListGoogleBusinessLocations200Response) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### GetHasMore

`func (o *ListGoogleBusinessLocations200Response) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ListGoogleBusinessLocations200Response) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ListGoogleBusinessLocations200Response) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ListGoogleBusinessLocations200Response) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


