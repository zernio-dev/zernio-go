# SendInboxMessageRequestLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float32** | Latitude in decimal degrees. | 
**Longitude** | **float32** | Longitude in decimal degrees. | 
**Name** | Pointer to **string** | Optional location name. | [optional] 
**Address** | Pointer to **string** | Optional street address. | [optional] 

## Methods

### NewSendInboxMessageRequestLocation

`func NewSendInboxMessageRequestLocation(latitude float32, longitude float32, ) *SendInboxMessageRequestLocation`

NewSendInboxMessageRequestLocation instantiates a new SendInboxMessageRequestLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestLocationWithDefaults

`func NewSendInboxMessageRequestLocationWithDefaults() *SendInboxMessageRequestLocation`

NewSendInboxMessageRequestLocationWithDefaults instantiates a new SendInboxMessageRequestLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *SendInboxMessageRequestLocation) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *SendInboxMessageRequestLocation) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *SendInboxMessageRequestLocation) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *SendInboxMessageRequestLocation) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *SendInboxMessageRequestLocation) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *SendInboxMessageRequestLocation) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.


### GetName

`func (o *SendInboxMessageRequestLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendInboxMessageRequestLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendInboxMessageRequestLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SendInboxMessageRequestLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *SendInboxMessageRequestLocation) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SendInboxMessageRequestLocation) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SendInboxMessageRequestLocation) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SendInboxMessageRequestLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


