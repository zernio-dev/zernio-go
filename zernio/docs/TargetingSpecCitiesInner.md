# TargetingSpecCitiesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Radius** | Pointer to **float32** | Radius around the city. Requires distanceUnit. | [optional] 
**DistanceUnit** | Pointer to **string** | Required if radius is set. | [optional] 

## Methods

### NewTargetingSpecCitiesInner

`func NewTargetingSpecCitiesInner(key string, ) *TargetingSpecCitiesInner`

NewTargetingSpecCitiesInner instantiates a new TargetingSpecCitiesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetingSpecCitiesInnerWithDefaults

`func NewTargetingSpecCitiesInnerWithDefaults() *TargetingSpecCitiesInner`

NewTargetingSpecCitiesInnerWithDefaults instantiates a new TargetingSpecCitiesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TargetingSpecCitiesInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetingSpecCitiesInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetingSpecCitiesInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *TargetingSpecCitiesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetingSpecCitiesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetingSpecCitiesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TargetingSpecCitiesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRadius

`func (o *TargetingSpecCitiesInner) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *TargetingSpecCitiesInner) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *TargetingSpecCitiesInner) SetRadius(v float32)`

SetRadius sets Radius field to given value.

### HasRadius

`func (o *TargetingSpecCitiesInner) HasRadius() bool`

HasRadius returns a boolean if a field has been set.

### GetDistanceUnit

`func (o *TargetingSpecCitiesInner) GetDistanceUnit() string`

GetDistanceUnit returns the DistanceUnit field if non-nil, zero value otherwise.

### GetDistanceUnitOk

`func (o *TargetingSpecCitiesInner) GetDistanceUnitOk() (*string, bool)`

GetDistanceUnitOk returns a tuple with the DistanceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistanceUnit

`func (o *TargetingSpecCitiesInner) SetDistanceUnit(v string)`

SetDistanceUnit sets DistanceUnit field to given value.

### HasDistanceUnit

`func (o *TargetingSpecCitiesInner) HasDistanceUnit() bool`

HasDistanceUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


