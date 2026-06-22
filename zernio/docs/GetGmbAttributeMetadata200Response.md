# GetGmbAttributeMetadata200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** | Only present in location mode. | [optional] 
**AttributeMetadata** | Pointer to [**[]GetGmbAttributeMetadata200ResponseAttributeMetadataInner**](GetGmbAttributeMetadata200ResponseAttributeMetadataInner.md) |  | [optional] 
**NextPageToken** | Pointer to **string** | Present when additional pages of results are available. | [optional] 

## Methods

### NewGetGmbAttributeMetadata200Response

`func NewGetGmbAttributeMetadata200Response() *GetGmbAttributeMetadata200Response`

NewGetGmbAttributeMetadata200Response instantiates a new GetGmbAttributeMetadata200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGmbAttributeMetadata200ResponseWithDefaults

`func NewGetGmbAttributeMetadata200ResponseWithDefaults() *GetGmbAttributeMetadata200Response`

NewGetGmbAttributeMetadata200ResponseWithDefaults instantiates a new GetGmbAttributeMetadata200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGmbAttributeMetadata200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGmbAttributeMetadata200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGmbAttributeMetadata200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGmbAttributeMetadata200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGmbAttributeMetadata200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGmbAttributeMetadata200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGmbAttributeMetadata200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGmbAttributeMetadata200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetGmbAttributeMetadata200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetGmbAttributeMetadata200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetGmbAttributeMetadata200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetGmbAttributeMetadata200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetAttributeMetadata

`func (o *GetGmbAttributeMetadata200Response) GetAttributeMetadata() []GetGmbAttributeMetadata200ResponseAttributeMetadataInner`

GetAttributeMetadata returns the AttributeMetadata field if non-nil, zero value otherwise.

### GetAttributeMetadataOk

`func (o *GetGmbAttributeMetadata200Response) GetAttributeMetadataOk() (*[]GetGmbAttributeMetadata200ResponseAttributeMetadataInner, bool)`

GetAttributeMetadataOk returns a tuple with the AttributeMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMetadata

`func (o *GetGmbAttributeMetadata200Response) SetAttributeMetadata(v []GetGmbAttributeMetadata200ResponseAttributeMetadataInner)`

SetAttributeMetadata sets AttributeMetadata field to given value.

### HasAttributeMetadata

`func (o *GetGmbAttributeMetadata200Response) HasAttributeMetadata() bool`

HasAttributeMetadata returns a boolean if a field has been set.

### GetNextPageToken

`func (o *GetGmbAttributeMetadata200Response) GetNextPageToken() string`

GetNextPageToken returns the NextPageToken field if non-nil, zero value otherwise.

### GetNextPageTokenOk

`func (o *GetGmbAttributeMetadata200Response) GetNextPageTokenOk() (*string, bool)`

GetNextPageTokenOk returns a tuple with the NextPageToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPageToken

`func (o *GetGmbAttributeMetadata200Response) SetNextPageToken(v string)`

SetNextPageToken sets NextPageToken field to given value.

### HasNextPageToken

`func (o *GetGmbAttributeMetadata200Response) HasNextPageToken() bool`

HasNextPageToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


