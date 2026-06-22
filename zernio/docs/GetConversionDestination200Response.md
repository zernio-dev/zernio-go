# GetConversionDestination200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to [**ConversionDestination**](ConversionDestination.md) |  | [optional] 

## Methods

### NewGetConversionDestination200Response

`func NewGetConversionDestination200Response() *GetConversionDestination200Response`

NewGetConversionDestination200Response instantiates a new GetConversionDestination200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConversionDestination200ResponseWithDefaults

`func NewGetConversionDestination200ResponseWithDefaults() *GetConversionDestination200Response`

NewGetConversionDestination200ResponseWithDefaults instantiates a new GetConversionDestination200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetConversionDestination200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetConversionDestination200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetConversionDestination200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetConversionDestination200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDestination

`func (o *GetConversionDestination200Response) GetDestination() ConversionDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *GetConversionDestination200Response) GetDestinationOk() (*ConversionDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *GetConversionDestination200Response) SetDestination(v ConversionDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *GetConversionDestination200Response) HasDestination() bool`

HasDestination returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


