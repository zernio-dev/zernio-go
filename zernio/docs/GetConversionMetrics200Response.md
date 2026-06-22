# GetConversionMetrics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Granularity** | Pointer to **string** |  | [optional] 
**Rows** | Pointer to [**[]GetConversionMetrics200ResponseRowsInner**](GetConversionMetrics200ResponseRowsInner.md) |  | [optional] 

## Methods

### NewGetConversionMetrics200Response

`func NewGetConversionMetrics200Response() *GetConversionMetrics200Response`

NewGetConversionMetrics200Response instantiates a new GetConversionMetrics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConversionMetrics200ResponseWithDefaults

`func NewGetConversionMetrics200ResponseWithDefaults() *GetConversionMetrics200Response`

NewGetConversionMetrics200ResponseWithDefaults instantiates a new GetConversionMetrics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetConversionMetrics200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetConversionMetrics200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetConversionMetrics200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetConversionMetrics200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetGranularity

`func (o *GetConversionMetrics200Response) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *GetConversionMetrics200Response) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *GetConversionMetrics200Response) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.

### HasGranularity

`func (o *GetConversionMetrics200Response) HasGranularity() bool`

HasGranularity returns a boolean if a field has been set.

### GetRows

`func (o *GetConversionMetrics200Response) GetRows() []GetConversionMetrics200ResponseRowsInner`

GetRows returns the Rows field if non-nil, zero value otherwise.

### GetRowsOk

`func (o *GetConversionMetrics200Response) GetRowsOk() (*[]GetConversionMetrics200ResponseRowsInner, bool)`

GetRowsOk returns a tuple with the Rows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRows

`func (o *GetConversionMetrics200Response) SetRows(v []GetConversionMetrics200ResponseRowsInner)`

SetRows sets Rows field to given value.

### HasRows

`func (o *GetConversionMetrics200Response) HasRows() bool`

HasRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


