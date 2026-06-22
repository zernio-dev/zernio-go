# CreateWhatsAppDataset200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | Pointer to **string** | Meta dataset ID linked to the WABA | [optional] 
**Created** | Pointer to **bool** | True if Meta created a new dataset on this call; false if one already existed | [optional] 

## Methods

### NewCreateWhatsAppDataset200Response

`func NewCreateWhatsAppDataset200Response() *CreateWhatsAppDataset200Response`

NewCreateWhatsAppDataset200Response instantiates a new CreateWhatsAppDataset200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhatsAppDataset200ResponseWithDefaults

`func NewCreateWhatsAppDataset200ResponseWithDefaults() *CreateWhatsAppDataset200Response`

NewCreateWhatsAppDataset200ResponseWithDefaults instantiates a new CreateWhatsAppDataset200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetId

`func (o *CreateWhatsAppDataset200Response) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *CreateWhatsAppDataset200Response) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *CreateWhatsAppDataset200Response) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.

### HasDatasetId

`func (o *CreateWhatsAppDataset200Response) HasDatasetId() bool`

HasDatasetId returns a boolean if a field has been set.

### GetCreated

`func (o *CreateWhatsAppDataset200Response) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateWhatsAppDataset200Response) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateWhatsAppDataset200Response) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CreateWhatsAppDataset200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


