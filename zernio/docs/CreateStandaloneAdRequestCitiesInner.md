# CreateStandaloneAdRequestCitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Meta city ID, from /v1/ads/targeting/search results. | 
**Radius** | Pointer to **float32** | Optional radius around the city. Must be set together with distance_unit. | [optional] 
**DistanceUnit** | Pointer to **string** | Unit for radius. Required if radius is set. | [optional] 

## Methods

### NewCreateStandaloneAdRequestCitiesInner

`func NewCreateStandaloneAdRequestCitiesInner(key string, ) *CreateStandaloneAdRequestCitiesInner`

NewCreateStandaloneAdRequestCitiesInner instantiates a new CreateStandaloneAdRequestCitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateStandaloneAdRequestCitiesInnerWithDefaults

`func NewCreateStandaloneAdRequestCitiesInnerWithDefaults() *CreateStandaloneAdRequestCitiesInner`

NewCreateStandaloneAdRequestCitiesInnerWithDefaults instantiates a new CreateStandaloneAdRequestCitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *CreateStandaloneAdRequestCitiesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateStandaloneAdRequestCitiesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateStandaloneAdRequestCitiesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetRadius

`func (o *CreateStandaloneAdRequestCitiesInner) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *CreateStandaloneAdRequestCitiesInner) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *CreateStandaloneAdRequestCitiesInner) SetRadius(v float32)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *CreateStandaloneAdRequestCitiesInner) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetDistanceUnit

`func (o *CreateStandaloneAdRequestCitiesInner) GetDistanceUnit() string`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *CreateStandaloneAdRequestCitiesInner) GetDistanceUnitOk() (*string, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *CreateStandaloneAdRequestCitiesInner) SetDistanceUnit(v string)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *CreateStandaloneAdRequestCitiesInner) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


