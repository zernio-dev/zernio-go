# GetInboxHeatmap200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Buckets** | Pointer to [**[]GetInboxHeatmap200ResponseBucketsInner**](GetInboxHeatmap200ResponseBucketsInner.md) |  | [optional] 

## Methods

### NewGetInboxHeatmap200Response

`func NewGetInboxHeatmap200Response() *GetInboxHeatmap200Response`

NewGetInboxHeatmap200Response instantiates a new GetInboxHeatmap200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxHeatmap200ResponseWithDefaults

`func NewGetInboxHeatmap200ResponseWithDefaults() *GetInboxHeatmap200Response`

NewGetInboxHeatmap200ResponseWithDefaults instantiates a new GetInboxHeatmap200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetInboxHeatmap200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetInboxHeatmap200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetInboxHeatmap200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetInboxHeatmap200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFrom

`func (o *GetInboxHeatmap200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxHeatmap200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxHeatmap200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxHeatmap200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetInboxHeatmap200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetInboxHeatmap200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetInboxHeatmap200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetInboxHeatmap200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *GetInboxHeatmap200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetInboxHeatmap200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetBuckets

`func (o *GetInboxHeatmap200Response) GetBuckets() []GetInboxHeatmap200ResponseBucketsInner`

GetBuckets returns the Buckets field if non-nil, zero value otherwise.

### GetBucketsOk

`func (o *GetInboxHeatmap200Response) GetBucketsOk() (*[]GetInboxHeatmap200ResponseBucketsInner, bool)`

GetBucketsOk returns a tuple with the Buckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuckets

`func (o *GetInboxHeatmap200Response) SetBuckets(v []GetInboxHeatmap200ResponseBucketsInner)`

SetBuckets sets Buckets field to given value.

### HasBuckets

`func (o *GetInboxHeatmap200Response) HasBuckets() bool`

HasBuckets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


