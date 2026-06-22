# GetWhatsAppCallEstimate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestinationCountry** | Pointer to **NullableString** |  | [optional] 
**PerMinuteUsd** | Pointer to **float32** |  | [optional] 
**Breakdown** | Pointer to [**GetWhatsAppCallEstimate200ResponseBreakdown**](GetWhatsAppCallEstimate200ResponseBreakdown.md) |  | [optional] 

## Methods

### NewGetWhatsAppCallEstimate200Response

`func NewGetWhatsAppCallEstimate200Response() *GetWhatsAppCallEstimate200Response`

NewGetWhatsAppCallEstimate200Response instantiates a new GetWhatsAppCallEstimate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppCallEstimate200ResponseWithDefaults

`func NewGetWhatsAppCallEstimate200ResponseWithDefaults() *GetWhatsAppCallEstimate200Response`

NewGetWhatsAppCallEstimate200ResponseWithDefaults instantiates a new GetWhatsAppCallEstimate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestinationCountry

`func (o *GetWhatsAppCallEstimate200Response) GetDestinationCountry() string`

GetDestinationCountry returns the DestinationCountry field if non-nil, zero value otherwise.

### GetDestinationCountryOk

`func (o *GetWhatsAppCallEstimate200Response) GetDestinationCountryOk() (*string, bool)`

GetDestinationCountryOk returns a tuple with the DestinationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationCountry

`func (o *GetWhatsAppCallEstimate200Response) SetDestinationCountry(v string)`

SetDestinationCountry sets DestinationCountry field to given value.

### HasDestinationCountry

`func (o *GetWhatsAppCallEstimate200Response) HasDestinationCountry() bool`

HasDestinationCountry returns a boolean if a field has been set.

### SetDestinationCountryNil

`func (o *GetWhatsAppCallEstimate200Response) SetDestinationCountryNil(b bool)`

 SetDestinationCountryNil sets the value for DestinationCountry to be an explicit nil

### UnsetDestinationCountry
`func (o *GetWhatsAppCallEstimate200Response) UnsetDestinationCountry()`

UnsetDestinationCountry ensures that no value is present for DestinationCountry, not even an explicit nil
### GetPerMinuteUsd

`func (o *GetWhatsAppCallEstimate200Response) GetPerMinuteUsd() float32`

GetPerMinuteUsd returns the PerMinuteUsd field if non-nil, zero value otherwise.

### GetPerMinuteUsdOk

`func (o *GetWhatsAppCallEstimate200Response) GetPerMinuteUsdOk() (*float32, bool)`

GetPerMinuteUsdOk returns a tuple with the PerMinuteUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerMinuteUsd

`func (o *GetWhatsAppCallEstimate200Response) SetPerMinuteUsd(v float32)`

SetPerMinuteUsd sets PerMinuteUsd field to given value.

### HasPerMinuteUsd

`func (o *GetWhatsAppCallEstimate200Response) HasPerMinuteUsd() bool`

HasPerMinuteUsd returns a boolean if a field has been set.

### GetBreakdown

`func (o *GetWhatsAppCallEstimate200Response) GetBreakdown() GetWhatsAppCallEstimate200ResponseBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *GetWhatsAppCallEstimate200Response) GetBreakdownOk() (*GetWhatsAppCallEstimate200ResponseBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *GetWhatsAppCallEstimate200Response) SetBreakdown(v GetWhatsAppCallEstimate200ResponseBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *GetWhatsAppCallEstimate200Response) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


