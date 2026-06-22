# GetInboxVolume200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**To** | Pointer to **NullableString** |  | [optional] 
**Summary** | Pointer to [**GetInboxVolume200ResponseSummary**](GetInboxVolume200ResponseSummary.md) |  | [optional] 
**Timeseries** | Pointer to [**[]GetInboxVolume200ResponseTimeseriesInner**](GetInboxVolume200ResponseTimeseriesInner.md) |  | [optional] 
**ByPlatform** | Pointer to [**[]GetInboxVolume200ResponseByPlatformInner**](GetInboxVolume200ResponseByPlatformInner.md) |  | [optional] 

## Methods

### NewGetInboxVolume200Response

`func NewGetInboxVolume200Response() *GetInboxVolume200Response`

NewGetInboxVolume200Response instantiates a new GetInboxVolume200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxVolume200ResponseWithDefaults

`func NewGetInboxVolume200ResponseWithDefaults() *GetInboxVolume200Response`

NewGetInboxVolume200ResponseWithDefaults instantiates a new GetInboxVolume200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetInboxVolume200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetInboxVolume200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetInboxVolume200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetInboxVolume200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFrom

`func (o *GetInboxVolume200Response) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GetInboxVolume200Response) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GetInboxVolume200Response) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GetInboxVolume200Response) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *GetInboxVolume200Response) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GetInboxVolume200Response) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GetInboxVolume200Response) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GetInboxVolume200Response) HasTo() bool`

HasTo returns a boolean if a field has been set.

### SetToNil

`func (o *GetInboxVolume200Response) SetToNil(b bool)`

 SetToNil sets the value for To to be an explicit nil

### UnsetTo
`func (o *GetInboxVolume200Response) UnsetTo()`

UnsetTo ensures that no value is present for To, not even an explicit nil
### GetSummary

`func (o *GetInboxVolume200Response) GetSummary() GetInboxVolume200ResponseSummary`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetInboxVolume200Response) GetSummaryOk() (*GetInboxVolume200ResponseSummary, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetInboxVolume200Response) SetSummary(v GetInboxVolume200ResponseSummary)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetInboxVolume200Response) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTimeseries

`func (o *GetInboxVolume200Response) GetTimeseries() []GetInboxVolume200ResponseTimeseriesInner`

GetTimeseries returns the Timeseries field if non-nil, zero value otherwise.

### GetTimeseriesOk

`func (o *GetInboxVolume200Response) GetTimeseriesOk() (*[]GetInboxVolume200ResponseTimeseriesInner, bool)`

GetTimeseriesOk returns a tuple with the Timeseries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeseries

`func (o *GetInboxVolume200Response) SetTimeseries(v []GetInboxVolume200ResponseTimeseriesInner)`

SetTimeseries sets Timeseries field to given value.

### HasTimeseries

`func (o *GetInboxVolume200Response) HasTimeseries() bool`

HasTimeseries returns a boolean if a field has been set.

### GetByPlatform

`func (o *GetInboxVolume200Response) GetByPlatform() []GetInboxVolume200ResponseByPlatformInner`

GetByPlatform returns the ByPlatform field if non-nil, zero value otherwise.

### GetByPlatformOk

`func (o *GetInboxVolume200Response) GetByPlatformOk() (*[]GetInboxVolume200ResponseByPlatformInner, bool)`

GetByPlatformOk returns a tuple with the ByPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByPlatform

`func (o *GetInboxVolume200Response) SetByPlatform(v []GetInboxVolume200ResponseByPlatformInner)`

SetByPlatform sets ByPlatform field to given value.

### HasByPlatform

`func (o *GetInboxVolume200Response) HasByPlatform() bool`

HasByPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


