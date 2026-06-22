# UpdateGoogleBusinessPlaceActionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Resource name of the place action link (e.g. locations/123/placeActionLinks/456) | 
**Uri** | Pointer to **string** | New action URL | [optional] 
**PlaceActionType** | Pointer to **string** | New action type | [optional] 

## Methods

### NewUpdateGoogleBusinessPlaceActionRequest

`func NewUpdateGoogleBusinessPlaceActionRequest(name string, ) *UpdateGoogleBusinessPlaceActionRequest`

NewUpdateGoogleBusinessPlaceActionRequest instantiates a new UpdateGoogleBusinessPlaceActionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoogleBusinessPlaceActionRequestWithDefaults

`func NewUpdateGoogleBusinessPlaceActionRequestWithDefaults() *UpdateGoogleBusinessPlaceActionRequest`

NewUpdateGoogleBusinessPlaceActionRequestWithDefaults instantiates a new UpdateGoogleBusinessPlaceActionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateGoogleBusinessPlaceActionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGoogleBusinessPlaceActionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGoogleBusinessPlaceActionRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUri

`func (o *UpdateGoogleBusinessPlaceActionRequest) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *UpdateGoogleBusinessPlaceActionRequest) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *UpdateGoogleBusinessPlaceActionRequest) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *UpdateGoogleBusinessPlaceActionRequest) HasUri() bool`

HasUri returns a boolean if a field has been set.

### GetPlaceActionType

`func (o *UpdateGoogleBusinessPlaceActionRequest) GetPlaceActionType() string`

GetPlaceActionType returns the PlaceActionType field if non-nil, zero value otherwise.

### GetPlaceActionTypeOk

`func (o *UpdateGoogleBusinessPlaceActionRequest) GetPlaceActionTypeOk() (*string, bool)`

GetPlaceActionTypeOk returns a tuple with the PlaceActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceActionType

`func (o *UpdateGoogleBusinessPlaceActionRequest) SetPlaceActionType(v string)`

SetPlaceActionType sets PlaceActionType field to given value.

### HasPlaceActionType

`func (o *UpdateGoogleBusinessPlaceActionRequest) HasPlaceActionType() bool`

HasPlaceActionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


