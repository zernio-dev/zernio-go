# PreviewQueue200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | Pointer to **string** |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Slots** | Pointer to [**[]time.Time**](time.Time.md) |  | [optional] 

## Methods

### NewPreviewQueue200Response

`func NewPreviewQueue200Response() *PreviewQueue200Response`

NewPreviewQueue200Response instantiates a new PreviewQueue200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewQueue200ResponseWithDefaults

`func NewPreviewQueue200ResponseWithDefaults() *PreviewQueue200Response`

NewPreviewQueue200ResponseWithDefaults instantiates a new PreviewQueue200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *PreviewQueue200Response) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PreviewQueue200Response) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PreviewQueue200Response) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *PreviewQueue200Response) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetCount

`func (o *PreviewQueue200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PreviewQueue200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PreviewQueue200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PreviewQueue200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSlots

`func (o *PreviewQueue200Response) GetSlots() []time.Time`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *PreviewQueue200Response) GetSlotsOk() (*[]time.Time, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *PreviewQueue200Response) SetSlots(v []time.Time)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *PreviewQueue200Response) HasSlots() bool`

HasSlots returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


